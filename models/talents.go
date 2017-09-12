package models

import "fmt"

// Formation represent a single formation (and what it does)
type Formation struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Modifier    int    `yaml:"modifier"`
}

// Talent represent a single talent and what it does
type Talent struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Effect      string `yaml:"effect"`
}

// Spell is a talent with an associated test and a score
type Spell struct {
	Talent      `yaml:",inline"`
	Score       DiceThrow `yaml:"score"`
	Test        string    `yaml:"test"`
	Scope       int       `yaml:"scope"`
	Modifier    int       `yaml:"modifier"`
	FmtModifier string
}

// Enrich adds some formatted fields
func (s *Spell) Enrich() {
	if s.Modifier > 0 {
		s.FmtModifier = fmt.Sprintf("+ %d", s.Modifier)
	} else if s.Modifier < 0 {
		s.FmtModifier = fmt.Sprintf("- %d", s.Modifier)
	}
}
