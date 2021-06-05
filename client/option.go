package client

func defaultOption() *analyzeOption {
	return &analyzeOption{}
}

type Option interface {
	apply(*analyzeOption)
}

type analyzeOption struct {
	hidden       bool
	rescan       bool
	waitFinished bool
}

func defaultScanOption() *recentScanOption {
	return &recentScanOption{}
}

type ScanOption interface {
	apply(*recentScanOption)
}

type recentScanOption struct {
	min uint
	max uint
}
