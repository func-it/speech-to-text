package config

import (
	"log"

	"github.com/spf13/pflag"
	viperlib "github.com/spf13/viper"
)

var viper *viperlib.Viper

func init() {
	v := viperlib.NewWithOptions(viperlib.KeyDelimiter("_"))
	v.AutomaticEnv()
	viper = v
}

func BindPFlag(key string, flag *pflag.Flag) error {
	return viper.BindPFlag(key, flag)
}

func (c *Conf) Load() error {
	configName := c.Config
	// Initialize Viper settings

	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./") // Look for config file in the current directory
	viper.AddConfigPath("./.config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config file, %v", err)
	}

	// Read the env file
	viper.SetConfigType("env")
	viper.SetConfigName(configName + ".env")

	err = viper.MergeInConfig()
	if err != nil {
		log.Fatalf("error merging env %s", err)
	}

	err = viper.Unmarshal(c)
	if err != nil {
		log.Fatalf("error unmarshaling conf %s", err)
	}

	return nil
}
