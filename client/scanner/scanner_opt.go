package scanner

import "github.com/tigerwill90/observatory/client/option"

type scanOptionImpl struct {
	f func(*option.AnalyzeOption)
}

func (s *scanOptionImpl) Apply(o *option.AnalyzeOption) {
	s.f(o)
}

func newScanOptionImpl(f func(analyzeOption *option.AnalyzeOption)) *scanOptionImpl {
	return &scanOptionImpl{f: f}
}

func HideResult(hidden bool) option.Option {
	return newScanOptionImpl(func(o *option.AnalyzeOption) {
		o.Hidden = hidden
	})
}

func ForceRescan(rescan bool) option.Option {
	return newScanOptionImpl(func(o *option.AnalyzeOption) {
		o.Rescan = rescan
	})
}

func WaitFinished(wait bool) option.Option {
	return newScanOptionImpl(func(o *option.AnalyzeOption) {
		o.WaitFinished = wait
	})
}
