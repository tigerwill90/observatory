package observatory

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tigerwill90/observatory/option"
	"github.com/tigerwill90/observatory/types"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"
	"time"
)

func TestClientAnalyzeAllOption(t *testing.T) {
	want := &types.ScannerResult{
		EndTime: "Tue, 22 Mar 2016 21:51:41 GMT",
		Grade:   "A",
		Hidden:  false,
		ResponseHeaders: map[string]string{
			"Content-Security-Policy": "Content-Security-Policy:default-src 'self'; style-src 'self' https://fonts.googleapis.com; font-src https://fonts.gstatic.com; frame-src 'none'; img-src 'self' https://cdn.jsdelivr.net data:; connect-src 'self'; frame-ancestors 'none'; script-src 'self' https://cdnjs.cloudflare.com",
			"Content-Type":            "text/html",
		},
		ScanID:              1,
		Score:               90,
		LikelihoodIndicator: "LOW",
		StartTime:           "Tue, 22 Mar 2016 21:51:40 GMT",
		State:               "FINISHED",
		TestsFailed:         2,
		TestsPassed:         9,
		TestsQuantity:       11,
	}

	var call uint32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wantQuery := url.Values{}
		wantQuery.Set("host", "observatory.mozilla.org")
		assert.Equal(t, wantQuery, r.URL.Query())
		assert.Equal(t, fmt.Sprintf("/%s", ApiCallAnalyze), r.URL.Path)
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				t.Fatal(err)
			}
			wantParam := url.Values{}
			wantParam.Set("hidden", "true")
			wantParam.Set("rescan", "true")
			assert.Equal(t, wantParam, r.PostForm)
			pendingResp := &types.ScannerResult{State: Pending}
			if err := json.NewEncoder(w).Encode(pendingResp); err != nil {
				t.Fatal(err)
			}
			return
		}

		nbCall := atomic.LoadUint32(&call)
		defer atomic.AddUint32(&call, 1)
		if nbCall <= 2 {
			pendingResp := &types.ScannerResult{State: Pending}
			if err := json.NewEncoder(w).Encode(pendingResp); err != nil {
				t.Fatal(err)
			}
			return
		}
		if nbCall <= 4 {
			startingResp := &types.ScannerResult{State: Starting}
			if err := json.NewEncoder(w).Encode(startingResp); err != nil {
				t.Fatal(err)
			}
			return
		}
		if nbCall <= 5 {
			runningResp := &types.ScannerResult{State: Running}
			if err := json.NewEncoder(w).Encode(runningResp); err != nil {
				t.Fatal(err)
			}
			return
		}

		if err := json.NewEncoder(w).Encode(want); err != nil {
			t.Fatal(err)
		}
	}))
	defer srv.Close()

	c := NewCustomClient(srv.Client(), srv.URL)
	got, err := c.Analyze(context.TODO(), "observatory.mozilla.org", option.ForceRescan(true), option.HideResult(true), option.WaitFinished(true, 1*time.Second))
	require.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestClientGetAssessment(t *testing.T) {
	want := &types.ScannerResult{
		EndTime: "Tue, 22 Mar 2016 21:51:41 GMT",
		Grade:   "A",
		Hidden:  false,
		ResponseHeaders: map[string]string{
			"Content-Security-Policy": "Content-Security-Policy:default-src 'self'; style-src 'self' https://fonts.googleapis.com; font-src https://fonts.gstatic.com; frame-src 'none'; img-src 'self' https://cdn.jsdelivr.net data:; connect-src 'self'; frame-ancestors 'none'; script-src 'self' https://cdnjs.cloudflare.com",
			"Content-Type":            "text/html",
		},
		ScanID:              1,
		Score:               90,
		LikelihoodIndicator: "LOW",
		StartTime:           "Tue, 22 Mar 2016 21:51:40 GMT",
		State:               "FINISHED",
		TestsFailed:         2,
		TestsPassed:         9,
		TestsQuantity:       11,
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wantQuery := url.Values{}
		wantQuery.Set("host", "observatory.mozilla.org")
		assert.Equal(t, wantQuery, r.URL.Query())
		assert.Equal(t, fmt.Sprintf("/%s", ApiCallAnalyze), r.URL.Path)
		if err := json.NewEncoder(w).Encode(want); err != nil {
			t.Fatal(err)
		}
	}))
	defer srv.Close()

	c := NewCustomClient(srv.Client(), srv.URL)

	got, err := c.GetAssessment(context.TODO(), "observatory.mozilla.org")
	require.Nil(t, err)
	require.Equal(t, want, got)
}

func TestClientGetTestResults(t *testing.T) {
	want := []byte("{\"content-security-policy\":{\"expectation\":\"csp-implemented-with-no-unsafe\",\"name\":\"content-security-policy\",\"output\":{\"data\":{\"connect-src\":[\"'self'\",\"https://sentry.prod.mozaws.net\"],\"default-src\":[\"'self'\"],\"font-src\":[\"'self'\",\"https://addons.cdn.mozilla.net\"],\"frame-src\":[\"'self'\",\"https://ic.paypal.com\",\"https://paypal.com\",\"https://www.google.com/recaptcha/\",\"https://www.paypal.com\"],\"img-src\":[\"'self'\",\"data:\",\"blob:\",\"https://www.paypal.com\",\"https://ssl.google-analytics.com\",\"https://addons.cdn.mozilla.net\",\"https://static.addons.mozilla.net\",\"https://ssl.gstatic.com/\",\"https://sentry.prod.mozaws.net\"],\"media-src\":[\"https://videos.cdn.mozilla.net\"],\"object-src\":[\"'none'\"],\"report-uri\":[\"/__cspreport__\"],\"script-src\":[\"'self'\",\"https://addons.mozilla.org\",\"https://www.paypalobjects.com\",\"https://apis.google.com\",\"https://www.google.com/recaptcha/\",\"https://www.gstatic.com/recaptcha/\",\"https://ssl.google-analytics.com\",\"https://addons.cdn.mozilla.net\"],\"style-src\":[\"'self'\",\"'unsafe-inline'\",\"https://addons.cdn.mozilla.net\"]}},\"pass\":false,\"result\":\"csp-implemented-with-unsafe-inline-in-style-src-only\",\"score_description\":\"Content Security Policy (CSP) implemented with unsafe-inline inside style-src directive\",\"score_modifier\":-5},\"contribute\":{\"expectation\":\"contribute-json-with-required-keys\",\"name\":\"contribute\",\"output\":{\"data\":{\"bugs\":{\"list\":\"https://github.com/mozilla/addons-server/issues\",\"report\":\"https://github.com/mozilla/addons-server/issues/new\"},\"description\":\"Mozilla's official site for add-ons to Mozilla software, such as Firefox, Thunderbird, and SeaMonkey.\",\"name\":\"Olympia\",\"participate\":{\"docs\":\"http://addons-server.readthedocs.org/\",\"home\":\"https://wiki.mozilla.org/Add-ons/Contribute/AMO/Code\",\"irc\":\"irc://irc.mozilla.org/#amo\",\"irc-contacts\":[\"andym\",\"cgrebs\",\"kumar\",\"magopian\",\"mstriemer\",\"muffinresearch\",\"tofumatt\"]},\"urls\":{\"dev\":\"https://addons-dev.allizom.org/\",\"prod\":\"https://addons.mozilla.org/\",\"stage\":\"https://addons.allizom.org/\"}}},\"pass\":true,\"result\":\"contribute-json-with-required-keys\",\"score_description\":\"Contribute.json implemented with the required contact information\",\"score_modifier\":0},\"cookies\":{\"expectation\":\"cookies-secure-with-httponly-sessions\",\"name\":\"cookies\",\"output\":{\"data\":{\"sessionid\":{\"domain\":\".addons.mozilla.org\",\"expires\":null,\"httponly\":true,\"max-age\":null,\"path\":\"/\",\"port\":null,\"secure\":true}}},\"pass\":true,\"result\":\"cookies-secure-with-httponly-sessions\",\"score_description\":\"All cookies use the Secure flag and all session cookies use the HttpOnly flag\",\"score_modifier\":0},\"cross-origin-resource-sharing\":{\"expectation\":\"cross-origin-resource-sharing-not-implemented\",\"name\":\"cross-origin-resource-sharing\",\"output\":{\"data\":{\"acao\":null,\"clientaccesspolicy\":null,\"crossdomain\":null}},\"pass\":true,\"result\":\"cross-origin-resource-sharing-not-implemented\",\"score_description\":\"Content is not visible via cross-origin resource sharing (CORS) files or headers\",\"score_modifier\":0},\"public-key-pinning\":{\"expectation\":\"hpkp-not-implemented\",\"name\":\"public-key-pinning\",\"output\":{\"data\":null,\"includeSubDomains\":false,\"max-age\":null,\"numPins\":null,\"preloaded\":false},\"pass\":true,\"result\":\"hpkp-not-implemented\",\"score_description\":\"HTTP Public Key Pinning (HPKP) header not implemented\",\"score_modifier\":0},\"redirection\":{\"expectation\":\"redirection-to-https\",\"name\":\"redirection\",\"output\":{\"destination\":\"https://addons.mozilla.org/en-US/firefox/\",\"redirects\":true,\"route\":[\"http://addons.mozilla.org/\",\"https://addons.mozilla.org/\",\"https://addons.mozilla.org/en-US/firefox/\"],\"status_code\":200},\"pass\":true,\"result\":\"redirection-to-https\",\"score_description\":\"Initial redirection is to https on same host, final destination is https\",\"score_modifier\":0},\"strict-transport-security\":{\"expectation\":\"hsts-implemented-max-age-at-least-six-months\",\"name\":\"strict-transport-security\",\"output\":{\"data\":\"max-age=31536000\",\"includeSubDomains\":false,\"max-age\":31536000,\"preload\":false,\"preloaded\":false},\"pass\":true,\"result\":\"hsts-implemented-max-age-at-least-six-months\",\"score_description\":\"HTTP Strict Transport Security (HSTS) header set to a minimum of six months (15768000)\",\"score_modifier\":0},\"subresource-integrity\":{\"expectation\":\"sri-implemented-and-external-scripts-loaded-securely\",\"name\":\"subresource-integrity\",\"output\":{\"data\":{\"https://addons.cdn.mozilla.net/static/js/impala-min.js?build=552decc-56eadb2f\":{\"crossorigin\":null,\"integrity\":null},\"https://addons.cdn.mozilla.net/static/js/preload-min.js?build=552decc-56eadb2f\":{\"crossorigin\":null,\"integrity\":null}}},\"pass\":false,\"result\":\"sri-not-implemented-but-external-scripts-loaded-securely\",\"score_description\":\"Subresource Integrity (SRI) not implemented, but all external scripts are loaded over https\",\"score_modifier\":-5},\"x-content-type-options\":{\"expectation\":\"x-content-type-options-nosniff\",\"name\":\"x-content-type-options\",\"output\":{\"data\":\"nosniff\"},\"pass\":true,\"result\":\"x-content-type-options-nosniff\",\"score_description\":\"X-Content-Type-Options header set to \\\"nosniff\\\"\",\"score_modifier\":0},\"x-frame-options\":{\"expectation\":\"x-frame-options-sameorigin-or-deny\",\"name\":\"x-frame-options\",\"output\":{\"data\":\"DENY\"},\"pass\":true,\"result\":\"x-frame-options-sameorigin-or-deny\",\"score_description\":\"X-Frame-Options (XFO) header set to SAMEORIGIN or DENY\",\"score_modifier\":0},\"x-xss-protection\":{\"expectation\":\"x-xss-protection-1-mode-block\",\"name\":\"x-xss-protection\",\"output\":{\"data\":\"1; mode=block\"},\"pass\":true,\"result\":\"x-xss-protection-enabled-mode-block\",\"score_description\":\"X-XSS-Protection header set to \\\"1; mode=block\\\"\",\"score_modifier\":0}}")

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wantQuery := url.Values{}
		wantQuery.Set("scan", "100")
		assert.Equal(t, wantQuery, r.URL.Query())
		assert.Equal(t, fmt.Sprintf("/%s", ApiCallGetScanResults), r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(want); err != nil {
			t.Fatal(err)
		}
	}))
	defer srv.Close()

	c := NewCustomClient(srv.Client(), srv.URL)
	result, err := c.GetTestResults(context.TODO(), 100)
	require.Nil(t, err)
	got, err := json.Marshal(result)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestClientGetScannerState(t *testing.T) {
	want := &types.ScannerStates{
		Aborted:  10,
		Failed:   281,
		Finished: 46240,
		Pending:  122,
		Starting: 96,
		Running:  128,
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, fmt.Sprintf("/%s", ApiCallGetScannerStates), r.URL.Path)
		if err := json.NewEncoder(w).Encode(want); err != nil {
			t.Fatal(err)
		}
	}))
	defer srv.Close()

	c := NewCustomClient(srv.Client(), srv.URL)
	got, err := c.GetScannerState(context.TODO())
	require.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestClientGetGradeDistribution(t *testing.T) {
	want := &types.ScannerGradeDistribution{
		A:  3,
		A1: 6,
		A2: 2,
		B:  8,
		B1: 76,
		B2: 79,
		C:  80,
		C1: 88,
		C2: 86,
		D:  60,
		D1: 110,
		D2: 215,
		F:  46770,
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, fmt.Sprintf("/%s", ApiCallGetGradeDistribution), r.URL.Path)
		if err := json.NewEncoder(w).Encode(want); err != nil {
			t.Fatal(err)
		}
	}))
	defer srv.Close()

	c := NewCustomClient(srv.Client(), srv.URL)
	got, err := c.GetGradeDistribution(context.TODO())
	require.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestClientGetScanHistory(t *testing.T) {
	want := []*types.ScannerHostHistory{
		{
			EndTime:              "Thu, 22 Sep 2016 23:24:28 GMT",
			EndTimeUnixTimestamp: 1474586668,
			Grade:                "C",
			ScanId:               1711106,
			Score:                50,
		},
		{
			EndTime:              "Thu, 09 Feb 2017 01:30:47 GMT",
			EndTimeUnixTimestamp: 1486603847,
			Grade:                "B+",
			ScanId:               3292839,
			Score:                80,
		},
		{
			EndTime:              "Fri, 10 Feb 2017 02:30:08 GMT",
			EndTimeUnixTimestamp: 1486693808,
			Grade:                "A",
			ScanId:               3302879,
			Score:                90,
		},
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wantQuery := url.Values{}
		wantQuery.Set("host", "observatory.mozilla.org")
		assert.Equal(t, wantQuery, r.URL.Query())
		assert.Equal(t, fmt.Sprintf("/%s", ApiCallGetHostHistory), r.URL.Path)
		if err := json.NewEncoder(w).Encode(&want); err != nil {
			t.Fatal(err)
		}
	}))
	defer srv.Close()

	c := NewCustomClient(srv.Client(), srv.URL)
	got, err := c.GetScanHistory(context.TODO(), "observatory.mozilla.org")
	require.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestClientGetRecentScans(t *testing.T) {
	want := types.ScannerRecentScans{
		"site1.mozilla.org": "A",
		"site2.mozilla.org": "B-",
		"site3.mozilla.org": "C+",
		"site4.mozilla.org": "F",
		"site5.mozilla.org": "F",
		"site6.mozilla.org": "B",
		"site7.mozilla.org": "F",
		"site8.mozilla.org": "B+",
		"site9.mozilla.org": "A+",
		"site0.mozilla.org": "A-",
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wantQuery := url.Values{}
		wantQuery.Set("min", "119")
		assert.Equal(t, wantQuery, r.URL.Query())
		assert.Equal(t, fmt.Sprintf("/%s", ApiCallGetRecentScans), r.URL.Path)
		if err := json.NewEncoder(w).Encode(&want); err != nil {
			t.Fatal(err)
		}
	}))
	defer srv.Close()

	c := NewCustomClient(srv.Client(), srv.URL)
	got, err := c.GetRecentScans(context.TODO(), option.WithMinScore(119))
	require.Nil(t, err)
	assert.Equal(t, want, got)
}
