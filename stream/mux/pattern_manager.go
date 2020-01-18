package mux

import (
	"log"
)

var std = newPatternManager()

func RegisterRegexp(name string, pattern string) {
	log.Printf("[INFO] Register regexp to pattern for stream mux: %s: %q", name, pattern)
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

func (h *patternManager) RegisterRegexp(name string, pattern string) {
	h.patterns[name] = pattern
}

func (h *patternManager) Get(name string) (string, bool) {
	pattern, ok := h.patterns[name]
	return pattern, ok
}
