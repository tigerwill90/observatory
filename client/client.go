package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const Endpoint = "https://http-observatory.security.mozilla.org/api/v1"

const (
	Aborted  = "ABORTED"
	Failed   = "FAILED"
	Finished = "FINISHED"
	Pending  = "PENDING"
	Starting = "STARTING"
	Running  = "RUNNING"
)

const (
	ApiCallGetGradeDistribution = "getGradeDistribution"
	ApiCallAnalyze              = "analyze"
	ApiCallGetScanResults       = "getScanResults"
	ApiCallGetScannerStates     = "getScannerStates"
	ApiCallHostHistory          = "getHostHistory"
	ApiCallRecentScans          = "getRecentScans"
)

var (
	ErrScannerAborted = errors.New("scan aborted")
	ErrScannerFailed  = errors.New("scan failed")
)

type Client struct {
	client *http.Client
	url    string
}

func New(c *http.Client) *Client {
	if c == nil {
		c = &http.Client{
			Transport: &http.Transport{
				TLSHandshakeTimeout: 5 * time.Second,
			},
			Timeout: 10 * time.Second,
		}
	}
	return &Client{
		client: c,
		url:    Endpoint,
	}
}

func (c *Client) Analyze(ctx context.Context, host string, opts ...Option) (*ScannerResult, error) {
	analyseOpt := defaultOption()
	for _, opt := range opts {
		opt.apply(analyseOpt)
	}

	result, err := c.analyze(ctx, host, analyseOpt)
	if err != nil {
		return nil, fmt.Errorf("invoke assessment failed: %w", err)
	}

	if result.State == "" {
		return c.GetAssessment(ctx, host)
	}

	if result.State == Finished || !analyseOpt.waitFinished {
		return result, nil
	}

	timer := time.NewTicker(10 * time.Second)
	defer timer.Stop()
STOP:
	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("retrieve assessment aborted: %w", ctx.Err())
		case <-timer.C:
			result, err = c.GetAssessment(ctx, host)
			if err != nil {
				return nil, err
			}

			if result.State == Finished {
				break STOP
			}
		}
	}

	return result, nil
}

func (c *Client) analyze(ctx context.Context, host string, opt *analyzeOption) (*ScannerResult, error) {
	data := url.Values{}
	data.Set("hidden", fmt.Sprintf("%t", opt.hidden))
	data.Set("rescan", fmt.Sprintf("%t", opt.rescan))

	queryParams := url.Values{}
	queryParams.Set("host", host)

	reqConfig := request{
		method:      "POST",
		apiCall:     ApiCallAnalyze,
		queryParams: queryParams,
		data:        data,
	}
	result := new(ScannerResult)
	if err := c.doRequest(ctx, reqConfig, result); err != nil {
		return nil, err
	}

	if err := checkScannerError(result.State); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetAssessment(ctx context.Context, host string) (*ScannerResult, error) {
	queryParams := url.Values{}
	queryParams.Set("host", host)

	reqConfig := request{
		method:      "GET",
		apiCall:     ApiCallAnalyze,
		queryParams: queryParams,
	}

	result := new(ScannerResult)
	if err := c.doRequest(ctx, reqConfig, result); err != nil {
		return nil, fmt.Errorf("retrieve assessment failed: %w", err)
	}

	if err := checkScannerError(result.State); err != nil {
		return nil, fmt.Errorf("retrieve assessment failed: %w", err)
	}

	return result, nil
}

func (c *Client) GetScannerState(ctx context.Context) (*ScannerStates, error) {
	reqConfig := request{
		method:  "GET",
		apiCall: ApiCallGetScannerStates,
	}

	states := new(ScannerStates)
	if err := c.doRequest(ctx, reqConfig, states); err != nil {
		return nil, fmt.Errorf("retrieve scanner states failed: %w", err)
	}

	return states, nil
}

func (c *Client) GetTestResults(ctx context.Context, scanID int) (*ScannerTestResult, error) {
	data := url.Values{}
	data.Set("scan", fmt.Sprintf("%d", scanID))

	reqConfig := request{
		method:      "GET",
		apiCall:     ApiCallGetScanResults,
		queryParams: data,
	}

	testResult := new(ScannerTestResult)
	if err := c.doRequest(ctx, reqConfig, testResult); err != nil {
		return nil, fmt.Errorf("retrieve test results failed: %w", err)
	}

	return testResult, nil
}

func (c *Client) GetGradeDistribution(ctx context.Context) (*ScannerGradeDistribution, error) {
	reqConfig := request{
		method:  "GET",
		apiCall: ApiCallGetGradeDistribution,
	}

	gradeDistribution := new(ScannerGradeDistribution)
	if err := c.doRequest(ctx, reqConfig, gradeDistribution); err != nil {
		return nil, fmt.Errorf("retrieve overall grade distribution failed: %w", err)
	}

	return gradeDistribution, nil
}

func (c *Client) GetScanHistory(ctx context.Context, host string) ([]*ScannerHostHistory, error) {
	data := url.Values{}
	data.Set("host", host)
	reqConfig := request{
		method:      "GET",
		apiCall:     ApiCallHostHistory,
		queryParams: data,
	}

	histories := make([]*ScannerHostHistory, 0)
	if err := c.doRequest(ctx, reqConfig, &histories); err != nil {
		return nil, fmt.Errorf("retrieve host's scan history failed: %w", err)
	}

	return histories, nil
}

func (c *Client) GetRecentScans(ctx context.Context, opt ScanOption) (ScannerRecentScans, error) {
	o := defaultScanOption()
	opt.apply(o)
	data := url.Values{}
	if o.max != 0 {
		data.Set("max", fmt.Sprintf("%d", o.max))
	} else {
		data.Set("min", fmt.Sprintf("%d", o.min))
	}

	recentScans := make(ScannerRecentScans)

	reqConfig := request{
		method:      "GET",
		apiCall:     ApiCallRecentScans,
		queryParams: data,
	}

	if err := c.doRequest(ctx, reqConfig, &recentScans); err != nil {
		return nil, fmt.Errorf("retrieve recent scans failed: %w", err)
	}

	return recentScans, nil
}

func checkScannerError(state string) error {
	if state == Aborted {
		return ErrScannerAborted
	}
	if state == Failed {
		return ErrScannerFailed
	}
	return nil
}
