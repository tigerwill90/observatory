[![Build Status](https://www.travis-ci.com/tigerwill90/observatory.svg?token=vig1YafE12hJt82nEwq2&branch=master)](https://www.travis-ci.com/tigerwill90/observatory)

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
````