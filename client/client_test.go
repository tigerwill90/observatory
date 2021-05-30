package client

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestClientAnalyze(t *testing.T) {
	c := New(nil)
	sr, err := c.Analyze(context.TODO(), "sylvainmuller.ch")
	require.Nil(t, err)
	fmt.Printf("%+v\n", sr)

	tests, err := c.GetTestResults(context.TODO(), sr.ScanID)
	require.Nil(t, err)
	fmt.Printf("%+v\n", tests)

}

func TestClientGetScannerState(t *testing.T) {
	c := New(new(http.Client))
	state, err := c.GetScannerState(context.TODO())
	require.Nil(t, err)
	fmt.Printf("%+v\n", state)
}

func TestClientGetGradeDistribution(t *testing.T) {
	c := New(nil)
	distribution, err := c.GetGradeDistribution(context.TODO())
	require.Nil(t, err)
	fmt.Printf("%+v\n", distribution)
}
