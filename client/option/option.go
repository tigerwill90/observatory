package option

func DefaultOption() *AnalyzeOption {
	return &AnalyzeOption{
		Hidden:       true,
		Rescan:       true,
		WaitFinished: true,
	}
}

type Option interface {
	Apply(*AnalyzeOption)
}

type AnalyzeOption struct {
	Hidden       bool
	Rescan       bool
	WaitFinished bool
}
