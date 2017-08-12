package models

// Equipment groups all the stuff the player has equipped or can equip
type Equipment struct {
	// Weapons
	Weapons       []Weapon       `yaml:"weapons"`
	RangedWeapons []RangedWeapon `yaml:"ranged_weapons"`

	// Armors
	Head          []Armor `yaml:"head"`
	ShoulderStrap []Armor `yaml:"shoulder_strap"`
	Jacket        []Armor `yaml:"jacket"`
	Holster       []Armor `yaml:"holster"`
	Trouser       []Armor `yaml:"trouser"`
	Belt          []Armor `yaml:"belt"`
	Boots         []Armor `yaml:"boots"`
	Other         []Armor `yaml:"other"`
}

// Armor is a simple armor
type Armor struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

// RangedWeapon describes a ranged weapon
type RangedWeapon struct {
	Type   string    `yaml:"type"`
	Scope  int       `yaml:"scope"` // In meters
	Name   string    `yaml:"name"`
	Damage DiceThrow `yaml:"damage"`
}

// Weapon describes a weapon (not ranged)
type Weapon struct {
	Type   string    `yaml:"type"`
	Name   string    `yaml:"name"`
	Damage DiceThrow `yaml:"damage"`
}
