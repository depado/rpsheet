package models

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
	Score       *DiceThrow `yaml:"score"`
	Test        string     `yaml:"test"`
	Scope       int        `yaml:"scope"`
	Rank        int        `yaml:"rank"`
	Ingredients []string   `yaml:"ingredients"`
	Duration    string     `yaml:"duration"`
	Size        string     `yaml:"size"`
}

// Enrich adds some formatted fields
func (s *Spell) Enrich() {
	if s.Score != nil {
		s.Score.Format()
	}
}
