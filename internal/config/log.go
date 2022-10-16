package config

type Log struct {
	Prefix    string `mapstructure:"prefix" json:"prefix"`
	Directory string `mapstructure:"directory" json:"directory"`
	Level     string `mapstructure:"level" json:"level"`
}
