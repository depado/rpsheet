package models

// Wealth is the money carried and in bank of the player
type Wealth struct {
	Bank    WealthStorage `yaml:"bank"`
	Carried WealthStorage `yaml:"carried"`
}

// WealthStorage is a single place where wealth can be stored
type WealthStorage struct {
	Platinum int `yaml:"platinum"`
	Gold     int `yaml:"gold"`
	Silver   int `yaml:"silver"`
	Bronze   int `yaml:"bronze"`
}
