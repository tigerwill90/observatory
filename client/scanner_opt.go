package client

type scanOptionImpl struct {
	f func(*analyzeOption)
}

func (s *scanOptionImpl) apply(o *analyzeOption) {
	s.f(o)
}

func newScanOptionImpl(f func(*analyzeOption)) *scanOptionImpl {
	return &scanOptionImpl{f: f}
}

func HideResult(hidden bool) Option {
	return newScanOptionImpl(func(o *analyzeOption) {
		o.hidden = hidden
	})
}

func ForceRescan(rescan bool) Option {
	return newScanOptionImpl(func(o *analyzeOption) {
		o.rescan = rescan
	})
}

func WaitFinished(wait bool) Option {
	return newScanOptionImpl(func(o *analyzeOption) {
		o.waitFinished = wait
	})
}

type recentScanOptionImpl struct {
	f func(*recentScanOption)
}

func (r *recentScanOptionImpl) apply(o *recentScanOption) {
	r.f(o)
}

func newRecentScanOptionImpl(f func(*recentScanOption)) *recentScanOptionImpl {
	return &recentScanOptionImpl{f: f}
}

func WithMinScore(min uint) ScanOption {
	return newRecentScanOptionImpl(func(o *recentScanOption) {
		o.min = min
	})
}

func WithMaxScore(max uint) ScanOption {
	return newRecentScanOptionImpl(func(o *recentScanOption) {
		o.max = max
	})
}
