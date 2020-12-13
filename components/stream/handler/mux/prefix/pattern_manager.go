package prefix

var std = newPatternManager()

func RegisterRegexp(name string, pattern string) {
	std.RegisterRegexp(name, pattern)
}

func Get(name string) (string, bool) {
	return std.Get(name)
}

type patternManager struct {
	patterns map[string]string
}

func newPatternManager() *patternManager {
	return &patternManager{
		patterns: map[string]string{},
	}
}

func (p *patternManager) RegisterRegexp(name string, pattern string) error {
	p.patterns[name] = pattern
	return nil
}

func (p *patternManager) Get(name string) (string, bool) {
	pattern, ok := p.patterns[name]
	return pattern, ok
}
