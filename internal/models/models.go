package models

type ConfigViper struct {
	Service struct {
		Name    string   `yaml:"name"`
		Enabled bool     `yaml:"enabled"`
		Number  int      `yaml:"number"`
		List    []string `yaml:"list"`
	} `yaml:"service"`
}
