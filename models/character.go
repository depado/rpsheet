package models

// Character represents a single player character
type Character struct {
	Name         string  `yaml:"name"`
	Age          int     `yaml:"age"`
	FakeID       string  `yaml:"fake_id"`
	Race         string  `yaml:"race"`
	Height       float64 `yaml:"height"`
	Weight       float64 `yaml:"weight"`
	Origin       string  `yaml:"origin"`
	PreferedHand string  `yaml:"prefered_hand"`
	XPs          int     `yaml:"xps"`
	SpentXPs     int     `yaml:"spent_xps"`

	HitPoints    int `yaml:"hit_points"`
	Fatigue      int `yaml:"fatigue"`
	ForcePoints  int `yaml:"force_points"`
	BaseMovement int `yaml:"base_movement"`

	Wealth Wealth `yaml:"wealth"`

	Formations []Formation `yaml:"formations"`
	Talents    []Talent    `yaml:"talents"`
	MainStats  []Stat      `yaml:"main_stats"`
	ExtraStats []Stat      `yaml:"extra_stats"`

	Equipment Equipment        `yaml:"equipment"`
	Inventory []QuantifiedItem `yaml:"inventory"`
}

// DiceThrow defines the way dice throws are represented
// For exemple 10D10 → Throws: 10, Type: 10
type DiceThrow struct {
	Throws int `yaml:"throws"`
	Type   int `yaml:"type"`
}

// Stat can be a main stat or an extra stat (in which case, rank and base
// are not taken into account)
type Stat struct {
	Name     string `yaml:"name"`
	Base     int    `yaml:"base"`
	Rank     string `yaml:"rank"`
	Modifier int    `yaml:"modifier"`
}

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
}
