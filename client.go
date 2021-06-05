package observatory

import (
	"context"
	"errors"
	"fmt"
	"github.com/tigerwill90/observatory/internal/option"
	"github.com/tigerwill90/observatory/types"
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
	ApiCallGetHostHistory       = "getHostHistory"
	ApiCallGetRecentScans       = "getRecentScans"
)

var (
	ErrScannerAborted = errors.New("scan aborted")
	ErrScannerFailed  = errors.New("scan failed")
)

// Client is an http.Client wrapper for HTTP Observatory api.
type Client struct {
	client *http.Client
	url    string
}

// NewClient return a preconfigured HTTP Observatory client.
// Use NewCustomClient to use an existing http.Client.
func NewClient() *Client {
	return &Client{
		client: &http.Client{
			Transport: &http.Transport{
				TLSHandshakeTimeout: 5 * time.Second,
			},
			Timeout: 10 * time.Second,
		},
		url: Endpoint,
	}
}

// NewCustomClient return an HTTP Observatory client from an
// http.Client and an url.
func NewCustomClient(c *http.Client, url string) *Client {
	return &Client{
		client: c,
		url:    url,
	}
}

// Analyze is used to invoke a new scan of a website. By default, Analyze will return a cached site result if
// the site has been scanned anytime in the previous 24 hours. Use option.ForceRescan to ignore cached result and
// start a new scan. Regardless of the state of option.ForceRescan, HTTP Observatory can not be scanned at a
// frequency greater than every 3 minutes and Analyze will return a cached result if it is the case.
// https://github.com/mozilla/http-observatory/blob/master/httpobs/docs/api.md#invoke-assessment
func (c *Client) Analyze(ctx context.Context, host string, opts ...option.Option) (*types.ScannerResult, error) {
	analyseOpt := option.DefaultOption()
	for _, opt := range opts {
		opt.Apply(analyseOpt)
	}

	result, err := c.analyze(ctx, host, analyseOpt)
	if err != nil {
		return nil, fmt.Errorf("invoke assessment failed: %w", err)
	}

	if result.State == "" {
		return c.GetAssessment(ctx, host)
	}

	if result.State == Finished || !analyseOpt.WaitFinished {
		return result, nil
	}

	interval := 10 * time.Second
	if analyseOpt.PullInterval != 0 {
		interval = analyseOpt.PullInterval
	}

	timer := time.NewTicker(interval)
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

func (c *Client) analyze(ctx context.Context, host string, opt *option.AnalyzeOption) (*types.ScannerResult, error) {
	data := url.Values{}
	data.Set("hidden", fmt.Sprintf("%t", opt.Hidden))
	data.Set("rescan", fmt.Sprintf("%t", opt.Rescan))

	queryParams := url.Values{}
	queryParams.Set("host", host)

	reqConfig := request{
		method:      "POST",
		apiCall:     ApiCallAnalyze,
		queryParams: queryParams,
		data:        data,
	}
	result := new(types.ScannerResult)
	if err := c.doRequest(ctx, reqConfig, result); err != nil {
		return nil, err
	}

	if err := checkScannerError(result.State); err != nil {
		return nil, err
	}

	return result, nil
}

// GetAssessment is used to retrieve the results of an existing, ongoing or completed scan.
// https://github.com/mozilla/http-observatory/blob/master/httpobs/docs/api.md#retrieve-assessment
func (c *Client) GetAssessment(ctx context.Context, host string) (*types.ScannerResult, error) {
	queryParams := url.Values{}
	queryParams.Set("host", host)

	reqConfig := request{
		method:      "GET",
		apiCall:     ApiCallAnalyze,
		queryParams: queryParams,
	}

	result := new(types.ScannerResult)
	if err := c.doRequest(ctx, reqConfig, result); err != nil {
		return nil, fmt.Errorf("retrieve assessment failed: %w", err)
	}

	if err := checkScannerError(result.State); err != nil {
		return nil, fmt.Errorf("retrieve assessment failed: %w", err)
	}

	return result, nil
}

// GetScannerState returns the state of the scanner. It can be useful for determining
// how busy the HTTP Observatory is.
// https://github.com/mozilla/http-observatory/blob/master/httpobs/docs/api.md#retrieve-scanner-states
func (c *Client) GetScannerState(ctx context.Context) (*types.ScannerStates, error) {
	reqConfig := request{
		method:  "GET",
		apiCall: ApiCallGetScannerStates,
	}

	states := new(types.ScannerStates)
	if err := c.doRequest(ctx, reqConfig, states); err != nil {
		return nil, fmt.Errorf("retrieve scanner states failed: %w", err)
	}

	return states, nil
}

// GetTestResults returns the detailed test result of a scan. The results of all these tests can
// bet retrieved once the scan's state has been placed in the FINISHED state.
// https://github.com/mozilla/http-observatory/blob/master/httpobs/docs/api.md#retrieve-test-results
func (c *Client) GetTestResults(ctx context.Context, scanID types.ScanID) (*types.ScannerTestResult, error) {
	data := url.Values{}
	data.Set("scan", fmt.Sprintf("%d", scanID))

	reqConfig := request{
		method:      "GET",
		apiCall:     ApiCallGetScanResults,
		queryParams: data,
	}

	testResult := new(types.ScannerTestResult)
	if err := c.doRequest(ctx, reqConfig, testResult); err != nil {
		return nil, fmt.Errorf("retrieve test results failed: %w", err)
	}

	return testResult, nil
}

// GetGradeDistribution retrieve each possible grade in the HTTP Observatory, as well as how many scans
// have fallen into that grade.
// https://github.com/mozilla/http-observatory/blob/master/httpobs/docs/api.md#retrieve-overall-grade-distribution
func (c *Client) GetGradeDistribution(ctx context.Context) (*types.ScannerGradeDistribution, error) {
	reqConfig := request{
		method:  "GET",
		apiCall: ApiCallGetGradeDistribution,
	}

	gradeDistribution := new(types.ScannerGradeDistribution)
	if err := c.doRequest(ctx, reqConfig, gradeDistribution); err != nil {
		return nil, fmt.Errorf("retrieve overall grade distribution failed: %w", err)
	}

	return gradeDistribution, nil
}

// GetScanHistory retrieve the ten most recent scans for the given host.
// https://github.com/mozilla/http-observatory/blob/master/httpobs/docs/api.md#retrieve-hosts-scan-history
func (c *Client) GetScanHistory(ctx context.Context, host string) ([]*types.ScannerHostHistory, error) {
	data := url.Values{}
	data.Set("host", host)
	reqConfig := request{
		method:      "GET",
		apiCall:     ApiCallGetHostHistory,
		queryParams: data,
	}

	histories := make([]*types.ScannerHostHistory, 0)
	if err := c.doRequest(ctx, reqConfig, &histories); err != nil {
		return nil, fmt.Errorf("retrieve host's scan history failed: %w", err)
	}

	return histories, nil
}

// GetRecentScans retrieve the ten most recent scans that fall withing a given score range.
// Use option.WithMaxScore to set the upper limit or option.WithMinScore to set the lower limit.
// https://github.com/mozilla/http-observatory/blob/master/httpobs/docs/api.md#retrieve-recent-scans
func (c *Client) GetRecentScans(ctx context.Context, opt option.ScanOption) (types.ScannerRecentScans, error) {
	o := option.DefaultScanOption()
	opt.Apply(o)
	data := url.Values{}
	if o.Max != 0 {
		data.Set("max", fmt.Sprintf("%d", o.Max))
	} else {
		data.Set("min", fmt.Sprintf("%d", o.Min))
	}

	recentScans := make(types.ScannerRecentScans)
	reqConfig := request{
		method:      "GET",
		apiCall:     ApiCallGetRecentScans,
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
