package types

type ScannerResult struct {
	EndTime             string            `json:"end_time"`
	Grade               string            `json:"grade"`
	Hidden              bool              `json:"hidden"`
	ResponseHeaders     map[string]string `json:"response_headers"`
	ScanID              int               `json:"scan_id"`
	Score               int               `json:"score"`
	LikelihoodIndicator string            `json:"likelihood_indicator"`
	StartTime           string            `json:"start_time"`
	State               string            `json:"state"`
	TestsFailed         int               `json:"tests_failed"`
	TestsPassed         int               `json:"tests_passed"`
	TestsQuantity       int               `json:"tests_quantity"`
}

type ScannerStates struct {
	Aborted  int `json:"ABORTED"`
	Failed   int `json:"FAILED"`
	Finished int `json:"FINISHED"`
	Pending  int `json:"PENDING"`
	Starting int `json:"STARTING"`
	Running  int `json:"RUNNING"`
}

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

type ScannerHostHistory struct {
	EndTime              string `json:"end_time"`
	EndTimeUnixTimestamp int    `json:"end_time_unix_timestamp"`
	Grade                string `json:"grade"`
	ScanId               int    `json:"scan_id"`
	Score                int    `json:"score"`
}

type ScannerRecentScans map[string]string
