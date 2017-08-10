package models

// QuantifiedItem adds a single quantity field to an item (for storage purpose)
type QuantifiedItem struct {
	Item     `yaml:",inline"`
	Quantity int `yaml:"quantity"`
}

// Item represents a single item
type Item struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Effect      string `yaml:"effect"`
	Unit        string `yaml:"unit"`
}
