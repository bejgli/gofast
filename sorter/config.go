package sorter

import (
	"os"
	"regexp"
)

type rule struct {
	Pattern string
	Target  string
	Overwrite bool
}

type Config struct {
	Source string
	Rules  []rule
}

// Check regex patterns for syntax errors.
func (c *Config) CheckPatterns() (p string) {
	for _, p = range c.getPatterns() {
		_, err := regexp.Compile(p)
		if err != nil {
			return
		}
	}

	return ""
}

// Check if source and target directories exist. Return bad dir name if not.
func (c *Config) CheckDirs() (dir string) {
	dirs := c.getTargets()
	dirs = append(dirs, c.Source)
	for _, dir = range dirs {
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			return
		}
	}

	return ""
}

// Config constructor used for easier initialization of test configs.
func newConfig(src string, rules []rule) *Config {
	return &Config{Source: src, Rules: rules}
}

func (c *Config) getPatterns() (patterns []string) {
	for _, r := range c.Rules {
		patterns = append(patterns, r.Pattern)
	}

	return
}

func (c *Config) getTargets() (targets []string) {
	for _, r := range c.Rules {
		targets = append(targets, r.Target)
	}

	return
}
