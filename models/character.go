package models

import (
	"fmt"
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Character represents a single player character
type Character struct {
	PortraitURL  string  `yaml:"portrait_url"`
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
	Stats      []Stat      `yaml:"stats"`
	Skills     []Stat      `yaml:"skills"`

	Equipment Equipment        `yaml:"equipment"`
	Inventory []QuantifiedItem `yaml:"inventory"`
}

// Load loads a character with the given yaml file
func (c *Character) Load(fp string) error {
	var err error
	var cr []byte

	if cr, err = ioutil.ReadFile(fp); err != nil {
		return err
	}
	if err = yaml.Unmarshal(cr, c); err != nil {
		return err
	}
	c.Enrich()
	return nil
}

// Enrich will execute many operations to try and enrich the character will
// computable data
func (c *Character) Enrich() {
	for i := range c.Stats {
		c.Stats[i].MatchIcon()
	}
	for i := range c.Skills {
		c.Skills[i].MatchIcon()
	}
	for i := range c.Equipment.Weapons {
		c.Equipment.Weapons[i].MatchIcon()
	}
	for i := range c.Equipment.RangedWeapons {
		c.Equipment.RangedWeapons[i].MatchIcon()
	}
}

// DiceThrow defines the way dice throws are represented
// For exemple 10D10 → Throws: 10, Type: 10
type DiceThrow struct {
	Throws  int `yaml:"throws"`
	Type    int `yaml:"type"`
	PerRank int `yaml:"per_rank"`
}

// Stat can be a main stat or an extra stat (in which case, rank and base
// are not taken into account)
type Stat struct {
	Name     string `yaml:"name"`
	Base     int    `yaml:"base"`
	Rank     string `yaml:"rank"`
	Modifier int    `yaml:"modifier"`
	Icon     string `yaml:"icon"`
}

// MatchIcon will assign a default icon if none is given
func (s *Stat) MatchIcon() {
	if s.Icon == "" {
		s.Icon = fmt.Sprintf("%s.svg", strings.ToLower(s.Name))
	}
}
