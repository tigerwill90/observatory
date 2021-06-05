package option

import "time"

func DefaultOption() *AnalyzeOption {
	return &AnalyzeOption{
		PullInterval: 10 * time.Second,
	}
}

type Option interface {
	Apply(*AnalyzeOption)
}

type AnalyzeOption struct {
	Hidden       bool
	Rescan       bool
	WaitFinished bool
	PullInterval time.Duration
}

func DefaultScanOption() *RecentScanOption {
	return &RecentScanOption{}
}

type ScanOption interface {
	Apply(*RecentScanOption)
}

type RecentScanOption struct {
	Min uint
	Max uint
}
