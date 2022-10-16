package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/0x00-ketsu/lazypip/internal/utils"
	"github.com/spf13/viper"
)

//go:embed default.yaml
var defaultConfig []byte

// Config is a structure contains configuratons data
type Config struct {
	App App `mapstrcture:"app" json:"app"`
	Pip Pip `mapstrcture:"pip" json:"pip"`
	Log Log `mapstrcture:"log" json:"log"`
}

// Load reads config file
// Create default config file if is not exist in OS user config dir
func Load() (*Config, error) {
	var configContent []byte
	s := Config{}

	v := viper.New()
	v.SetConfigType("yaml")

	filePath, err := userConfigFilepath()
	if err != nil {
		return nil, err
	}

	if !utils.FilePathExist(filePath) {
		configContent = defaultConfig
		if err = dumpConfigFile(filePath); err != nil {
			fmt.Printf("Dumps default config file to user config dir failed: %v\n", err.Error())
		}
	} else {
		configContent, err = os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
	}

	if err := v.ReadConfig(bytes.NewBuffer(configContent)); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(&s); err != nil {
		return nil, err
	}

	return &s, nil
}

// Return config file path of user config dir
func userConfigFilepath() (string, error) {
	confPath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	confPath = filepath.Join(confPath, "lazypip")
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		os.MkdirAll(confPath, 0755)
	} else if err != nil {
		return "", err
	}

	// copy default config file to OS user config dir
	filePath := filepath.Join(confPath, "config.yaml")
	return filePath, nil
}

// Copy default config file to OS user config dir
func dumpConfigFile(filePath string) error {
	if err := os.WriteFile(filePath, defaultConfig, 0644); err != nil {
		return err
	}

	fmt.Println("Dumps default config file to user config dir success.")
	return nil
}
