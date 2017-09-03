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
	Talent
	Score DiceThrow `yaml:"throw"`
	Test  string    `yaml:"test"`
}
