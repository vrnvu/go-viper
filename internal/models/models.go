package models

type ConfigViper struct {
	Service struct {
		Name    string   `yaml:"name"`
		Enabled bool     `yaml:"enabled"`
		Number  int      `yaml:"number"`
		List    []string `yaml:"list"`
	} `yaml:"service"`
}

type ValidConfigViper struct {
	Service struct {
		Name    string   `yaml:"name" validate:"required"`
		Enabled bool     `yaml:"enabled" validate:"required"`
		Number  int      `yaml:"number" validate:"required"`
		List    []string `yaml:"list" validate:"required"`
	} `yaml:"service" validate:"required"`
}
