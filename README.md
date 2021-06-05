[![Build Status](https://www.travis-ci.com/tigerwill90/observatory.svg?token=vig1YafE12hJt82nEwq2&branch=master)](https://www.travis-ci.com/tigerwill90/observatory)
[![codecov](https://codecov.io/gh/tigerwill90/observatory/branch/master/graph/badge.svg?token=rNMJHifM0i)](https://codecov.io/gh/tigerwill90/observatory)
[![Go Report Card](https://goreportcard.com/badge/github.com/tigerwill90/observatory)](https://goreportcard.com/report/github.com/tigerwill90/observatory)
[![Go Reference](https://pkg.go.dev/badge/github.com/tigerwill90/observatory.svg)](https://pkg.go.dev/github.com/tigerwill90/observatory)

# Observatory
A small and convenient Go client to consume [HTTP Observatory](https://observatory.mozilla.org/) [api](https://github.com/mozilla/http-observatory/blob/master/httpobs/docs/api.md).

**Install:**
````
go get -u github.com/tigerwill90/observatory
````

### Example
````go
c := observatory.NewClient()
result, err := c.Analyze(context.TODO(), "observatory.mozilla.org", option.ForceRescan(true), option.WaitFinished(true, 5*time.Second))
if err != nil {
    panic(err)
}
fmt.Printf("Grade: %s, Score: %d\n", result.Grade, result.Score)
// Grade: A+, Score: 125

detail, err := c.GetTestResults(context.TODO(), result.ScanID)
if err != nil {
    panic(err)
}

fmt.Printf(
    "Name: %s, Pass: %t, Expectation: %s\n",
    detail.ContentSecurityPolicy.Name,
    detail.ContentSecurityPolicy.Pass,
    detail.ContentSecurityPolicy.Expectation,
)
// Name: content-security-policy, Pass: true, Expectation: csp-implemented-with-no-unsafe
````

### Disclaimer
Breaking change may happen before `v1.0.0`.