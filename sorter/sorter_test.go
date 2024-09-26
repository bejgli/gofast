package sorter

import (
	"testing"
)

func TestCheckPatternsGoodPattern(t *testing.T) {
	pattern := "goodpattern"
	expected := ""

	c := newConfig("", []rule{{Pattern: pattern, Target: ""}})

	p := c.CheckPatterns()
	if p != expected {
		t.Fatalf("Expected no pattern, got: %s", p)
	}
}
func TestCheckPatternsBadPattern(t *testing.T) {
	pattern := ")" // Unmatched parenthesis
	expected := ")"

	c := newConfig("", []rule{{Pattern: pattern, Target: ""}})

	p := c.CheckPatterns()
	if p != expected {
		t.Fatalf("Expected pattern: %s, got: %s", pattern, p)
	}
}

func TestCheckDirsGoodDirs(t *testing.T) {
	// TODO
	t.Skip()
}

func TestCheckDirsBadDirs(t *testing.T) {
	// TODO
	t.Skip()
}
