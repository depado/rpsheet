package models

import (
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// CharMap is a map holding Characters based on the name of the yaml file they
// are stored in
type CharMap map[string]Character

// LoadGlob will load all the characters that matches the given unix shell
// pattern
func (cm CharMap) LoadGlob(pattern string) error {
	var err error
	var tl []string

	if tl, err = filepath.Glob(pattern); err != nil {
		return errors.Wrap(err, "glob error")
	}
	for _, v := range tl {
		var c Character
		cn := strings.TrimSuffix(filepath.Base(v), filepath.Ext(v))
		if err = c.Load(v); err != nil {
			return errors.Wrap(err, "couldn't load char")
		}
		cm[cn] = c
	}
	return err
}

// NewCharMapFromGlob creates a new CharMap and loads all the characters
// that matches the unix shell pattern inside it
func NewCharMapFromGlob(pattern string) (CharMap, error) {
	cm := make(CharMap)
	err := cm.LoadGlob(pattern)
	return cm, err
}