package models

import "fmt"

// DiceThrow defines the way dice throws are represented
// For exemple 10D10 → Throws: 10, Type: 10
type DiceThrow struct {
	Throws         int  `yaml:"throws"`
	Type           int  `yaml:"type"`
	Ranked         bool `yaml:"ranked"`
	Modifier       int  `yaml:"modifier"`
	RankedModifier bool `yaml:"ranked_modifier"`
	FmtModifier    string
}

// Format formats the dice throw
func (d *DiceThrow) Format() {
	if d.Modifier > 0 {
		d.FmtModifier = fmt.Sprintf("+ %d", d.Modifier)
	} else if d.Modifier < 0 {
		d.FmtModifier = fmt.Sprintf("- %d", d.Modifier)
	}
	if d.RankedModifier {
		d.FmtModifier = d.FmtModifier + " per rank"
	}
}
