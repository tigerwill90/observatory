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
	sr, err := c.Analyze(context.TODO(), "sylvainmuller.ch", ForceRescan(true), WaitFinished(true))
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

func TestClientGetScanHistory(t *testing.T) {
	c := New(nil)
	histories, err := c.GetScanHistory(context.TODO(), "sylvainmuller.ch")
	require.Nil(t, err)
	for _, history := range histories {
		fmt.Printf("%+v\n", history)
	}
}

func TestClientGetRecentScans(t *testing.T) {
	c := New(nil)
	recentScans, err := c.GetRecentScans(context.TODO(), WithMinScore(119))
	require.Nil(t, err)
	fmt.Println(recentScans)
}
