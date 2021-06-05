package types

type ScanID int

// ScannerResult is a summarized result of a scan.
type ScannerResult struct {
	// timestamp for when the scan completed
	EndTime string `json:"end_time"`
	// final grade assessed upon a completed scan
	Grade string `json:"grade"`
	// whether the scan results are unlisted on the recent results page
	Hidden bool `json:"hidden"`
	// the entirety of the HTTP response headers
	ResponseHeaders map[string]string `json:"response_headers"`
	// unique ID number assigned to the scan
	ScanID ScanID `json:"scan_id"`
	// final score assessed upon a completed (FINISHED) scan
	Score int `json:"score"`
	// Mozilla risk likelihood indicator that is the equivalent of the grade
	LikelihoodIndicator string `json:"likelihood_indicator"`
	// timestamp for when the scan was first request
	StartTime string `json:"start_time"`
	// the current state of the scan
	State string `json:"state"`
	// the number of subtests that were assigned a fail result
	TestsFailed int `json:"tests_failed"`
	// the number of subtests that were assigned a passing result
	TestsPassed int `json:"tests_passed"`
	// the total number of tests available and assessed at the time of the scan
	TestsQuantity int `json:"tests_quantity"`
}

// ScannerStates hold statistics on HTTP Observatory usage.
// Pending, starting and running state are good indicator of
// the current api load.
type ScannerStates struct {
	// aborted for internal technical reasons
	Aborted int `json:"ABORTED"`
	// failed to complete, typically due to the site being unavailable or timing out
	Failed int `json:"FAILED"`
	// completed successfully
	Finished int `json:"FINISHED"`
	// issued by the API but not yet picked up by a scanner instance
	Pending int `json:"PENDING"`
	// assigned to a scanning instance
	Starting int `json:"STARTING"`
	// currently in the process of scanning a website
	Running int `json:"RUNNING"`
}

// ScannerTestResult hold the detailed result of each test of a scan.
type ScannerTestResult struct {
	ContentSecurityPolicy struct {
		Expectation string `json:"expectation"`
		Name        string `json:"name"`
		Output      struct {
			Data struct {
				ConnectSrc []string `json:"connect-src"`
				DefaultSrc []string `json:"default-src"`
				FontSrc    []string `json:"font-src"`
				FrameSrc   []string `json:"frame-src"`
				ImgSrc     []string `json:"img-src"`
				MediaSrc   []string `json:"media-src"`
				ObjectSrc  []string `json:"object-src"`
				ReportUri  []string `json:"report-uri"`
				ScriptSrc  []string `json:"script-src"`
				StyleSrc   []string `json:"style-src"`
			} `json:"data"`
		} `json:"output"`
		Pass             bool   `json:"pass"`
		Result           string `json:"result"`
		ScoreDescription string `json:"score_description"`
		ScoreModifier    int    `json:"score_modifier"`
	} `json:"content-security-policy"`
	Contribute struct {
		Expectation string `json:"expectation"`
		Name        string `json:"name"`
		Output      struct {
			Data struct {
				Bugs struct {
					List   string `json:"list"`
					Report string `json:"report"`
				} `json:"bugs"`
				Description string `json:"description"`
				Name        string `json:"name"`
				Participate struct {
					Docs        string   `json:"docs"`
					Home        string   `json:"home"`
					Irc         string   `json:"irc"`
					IrcContacts []string `json:"irc-contacts"`
				} `json:"participate"`
				Urls struct {
					Dev   string `json:"dev"`
					Prod  string `json:"prod"`
					Stage string `json:"stage"`
				} `json:"urls"`
			} `json:"data"`
		} `json:"output"`
		Pass             bool   `json:"pass"`
		Result           string `json:"result"`
		ScoreDescription string `json:"score_description"`
		ScoreModifier    int    `json:"score_modifier"`
	} `json:"contribute"`
	Cookies struct {
		Expectation string `json:"expectation"`
		Name        string `json:"name"`
		Output      struct {
			Data struct {
				SessionId struct {
					Domain   string      `json:"domain"`
					Expires  interface{} `json:"expires"`
					Httponly bool        `json:"httponly"`
					MaxAge   interface{} `json:"max-age"`
					Path     string      `json:"path"`
					Port     interface{} `json:"port"`
					Secure   bool        `json:"secure"`
				} `json:"sessionid"`
			} `json:"data"`
		} `json:"output"`
		Pass             bool   `json:"pass"`
		Result           string `json:"result"`
		ScoreDescription string `json:"score_description"`
		ScoreModifier    int    `json:"score_modifier"`
	} `json:"cookies"`
	CrossOriginResourceSharing struct {
		Expectation string `json:"expectation"`
		Name        string `json:"name"`
		Output      struct {
			Data struct {
				Acao               interface{} `json:"acao"`
				ClientAccessPolicy interface{} `json:"clientaccesspolicy"`
				Crossdomain        interface{} `json:"crossdomain"`
			} `json:"data"`
		} `json:"output"`
		Pass             bool   `json:"pass"`
		Result           string `json:"result"`
		ScoreDescription string `json:"score_description"`
		ScoreModifier    int    `json:"score_modifier"`
	} `json:"cross-origin-resource-sharing"`
	PublicKeyPinning struct {
		Expectation string `json:"expectation"`
		Name        string `json:"name"`
		Output      struct {
			Data              interface{} `json:"data"`
			IncludeSubDomains bool        `json:"includeSubDomains"`
			MaxAge            interface{} `json:"max-age"`
			NumPins           interface{} `json:"numPins"`
			Preloaded         bool        `json:"preloaded"`
		} `json:"output"`
		Pass             bool   `json:"pass"`
		Result           string `json:"result"`
		ScoreDescription string `json:"score_description"`
		ScoreModifier    int    `json:"score_modifier"`
	} `json:"public-key-pinning"`
	Redirection struct {
		Expectation string `json:"expectation"`
		Name        string `json:"name"`
		Output      struct {
			Destination string   `json:"destination"`
			Redirects   bool     `json:"redirects"`
			Route       []string `json:"route"`
			StatusCode  int      `json:"status_code"`
		} `json:"output"`
		Pass             bool   `json:"pass"`
		Result           string `json:"result"`
		ScoreDescription string `json:"score_description"`
		ScoreModifier    int    `json:"score_modifier"`
	} `json:"redirection"`
	StrictTransportSecurity struct {
		Expectation string `json:"expectation"`
		Name        string `json:"name"`
		Output      struct {
			Data              string `json:"data"`
			IncludeSubDomains bool   `json:"includeSubDomains"`
			MaxAge            int    `json:"max-age"`
			Preload           bool   `json:"preload"`
			Preloaded         bool   `json:"preloaded"`
		} `json:"output"`
		Pass             bool   `json:"pass"`
		Result           string `json:"result"`
		ScoreDescription string `json:"score_description"`
		ScoreModifier    int    `json:"score_modifier"`
	} `json:"strict-transport-security"`
	SubresourceIntegrity struct {
		Expectation string `json:"expectation"`
		Name        string `json:"name"`
		Output      struct {
			Data struct {
				HttpsAddonsCdnMozillaNetStaticJsImpalaMinJsBuild552Decc56Eadb2F struct {
					CrossOrigin interface{} `json:"crossorigin"`
					Integrity   interface{} `json:"integrity"`
				} `json:"https://addons.cdn.mozilla.net/static/js/impala-min.js?build=552decc-56eadb2f"`
				HttpsAddonsCdnMozillaNetStaticJsPreloadMinJsBuild552Decc56Eadb2F struct {
					CrossOrigin interface{} `json:"crossorigin"`
					Integrity   interface{} `json:"integrity"`
				} `json:"https://addons.cdn.mozilla.net/static/js/preload-min.js?build=552decc-56eadb2f"`
			} `json:"data"`
		} `json:"output"`
		Pass             bool   `json:"pass"`
		Result           string `json:"result"`
		ScoreDescription string `json:"score_description"`
		ScoreModifier    int    `json:"score_modifier"`
	} `json:"subresource-integrity"`
	XContentTypeOptions struct {
		Expectation string `json:"expectation"`
		Name        string `json:"name"`
		Output      struct {
			Data string `json:"data"`
		} `json:"output"`
		Pass             bool   `json:"pass"`
		Result           string `json:"result"`
		ScoreDescription string `json:"score_description"`
		ScoreModifier    int    `json:"score_modifier"`
	} `json:"x-content-type-options"`
	XFrameOptions struct {
		Expectation string `json:"expectation"`
		Name        string `json:"name"`
		Output      struct {
			Data string `json:"data"`
		} `json:"output"`
		Pass             bool   `json:"pass"`
		Result           string `json:"result"`
		ScoreDescription string `json:"score_description"`
		ScoreModifier    int    `json:"score_modifier"`
	} `json:"x-frame-options"`
	XXssProtection struct {
		Expectation string `json:"expectation"`
		Name        string `json:"name"`
		Output      struct {
			Data string `json:"data"`
		} `json:"output"`
		Pass             bool   `json:"pass"`
		Result           string `json:"result"`
		ScoreDescription string `json:"score_description"`
		ScoreModifier    int    `json:"score_modifier"`
	} `json:"x-xss-protection"`
}

// ScannerGradeDistribution hold statistics on "Grade" repartition
// of each public site web scanned by HTTP Observatory.
type ScannerGradeDistribution struct {
	A  int `json:"A+"`
	A1 int `json:"A"`
	A2 int `json:"A-"`
	B  int `json:"B+"`
	B1 int `json:"B"`
	B2 int `json:"B-"`
	C  int `json:"C+"`
	C1 int `json:"C"`
	C2 int `json:"C-"`
	D  int `json:"D+"`
	D1 int `json:"D"`
	D2 int `json:"D-"`
	F  int `json:"F"`
}

// ScannerHostHistory hold a short summary of the result of a past scan.
type ScannerHostHistory struct {
	EndTime              string `json:"end_time"`
	EndTimeUnixTimestamp int    `json:"end_time_unix_timestamp"`
	Grade                string `json:"grade"`
	ScanId               ScanID `json:"scan_id"`
	Score                int    `json:"score"`
}

// ScannerRecentScans hold the grade result of maximum ten last scans.
type ScannerRecentScans map[string]string
