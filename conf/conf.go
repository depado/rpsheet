package conf

import (
	"io/ioutil"

	"github.com/imdario/mergo"
	yaml "gopkg.in/yaml.v2"
)

// Conf holds the configuration structure
type Conf struct {
	Host  string `yaml:"host"`
	Port  int    `yaml:"port"`
	Debug bool   `yaml:"debug"`
}

// C is the main configuration exported
var C Conf

// Defaults fills the configuration with some sane defaults
func (c *Conf) Defaults() error {
	return mergo.Merge(c, Conf{
		Host: "127.0.0.1",
		Port: 8080,
	})
}

// Load loads a yaml file into the C configuration struct
func Load(fp string) error {
	var err error
	var conf []byte

	if conf, err = ioutil.ReadFile(fp); err != nil {
		return err
	}
	if err = yaml.Unmarshal(conf, &C); err != nil {
		return err
	}
	return C.Defaults()
}
