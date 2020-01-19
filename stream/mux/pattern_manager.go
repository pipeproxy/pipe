package mux

import (
	"log"
)

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

func (h *patternManager) RegisterRegexp(name string, pattern string) error {
	log.Printf("[INFO] Register stream mux pattern: %s: %q", name, pattern)
	h.patterns[name] = pattern
	return nil
}

func (h *patternManager) Get(name string) (string, bool) {
	pattern, ok := h.patterns[name]
	return pattern, ok
}
