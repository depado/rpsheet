package models

import (
	"fmt"
	"strings"
)

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

// Weapon describes a weapon (not ranged)
type Weapon struct {
	Type        string    `yaml:"type"`
	Icon        string    `yaml:"icon"`
	Name        string    `yaml:"name"`
	Notes       []string  `yaml:"notes"`
	Modifiers   []string  `yaml:"modifiers"`
	Damage      DiceThrow `yaml:"damage"`
	Description string    `yaml:"description"`
}

// MatchIcon will assign a default icon if none is set
func (w *Weapon) MatchIcon() {
	if w.Icon == "" {
		w.Icon = fmt.Sprintf("%s.svg", strings.ToLower(w.Type))
	}
}

// RangedWeapon describes a ranged weapon
type RangedWeapon struct {
	Weapon `yaml:",inline"`
	Scope  int `yaml:"scope"` // In meters
}

// Armor is a simple armor
type Armor struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}
