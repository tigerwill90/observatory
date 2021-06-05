package option

import (
	"github.com/tigerwill90/observatory/internal/option"
	"time"
)

type scanOptionImpl struct {
	f func(*option.AnalyzeOption)
}

func (s *scanOptionImpl) Apply(o *option.AnalyzeOption) {
	s.f(o)
}

func newScanOptionImpl(f func(*option.AnalyzeOption)) *scanOptionImpl {
	return &scanOptionImpl{f: f}
}

// HideResult will hide a scan from public results.
func HideResult(hidden bool) option.Option {
	return newScanOptionImpl(func(o *option.AnalyzeOption) {
		o.Hidden = hidden
	})
}

// ForceRescan forces a rescan of a site.
func ForceRescan(rescan bool) option.Option {
	return newScanOptionImpl(func(o *option.AnalyzeOption) {
		o.Rescan = rescan
	})
}

// WaitFinished wait until the scanner state reach FINISHED.
func WaitFinished(wait bool, pullInterval time.Duration) option.Option {
	return newScanOptionImpl(func(o *option.AnalyzeOption) {
		o.WaitFinished = wait
		o.PullInterval = pullInterval
	})
}

type recentScanOptionImpl struct {
	f func(*option.RecentScanOption)
}

func (r *recentScanOptionImpl) Apply(o *option.RecentScanOption) {
	r.f(o)
}

func newRecentScanOptionImpl(f func(*option.RecentScanOption)) *recentScanOptionImpl {
	return &recentScanOptionImpl{f: f}
}

// WithMinScore set minimum score to retrieve.
func WithMinScore(min uint) option.ScanOption {
	return newRecentScanOptionImpl(func(o *option.RecentScanOption) {
		o.Min = min
	})
}

// WithMaxScore set maximum score to retrieve.
func WithMaxScore(max uint) option.ScanOption {
	return newRecentScanOptionImpl(func(o *option.RecentScanOption) {
		o.Max = max
	})
}
