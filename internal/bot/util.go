package bot

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// Configuration holds config details
type Configuration struct {
	BotToken string `mapstructure:"bottoken"`
	DBUser   string `mapstructure:"dbuser"`
	DBPass   string `mapstructure:"dbpass"`
}

// Config reads a given config file
func Config() Configuration {
	v := viper.New()
	var config Configuration
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.ReadInConfig()

	if err := v.UnmarshalExact(&config); err != nil {
		log.Error(fmt.Sprintf("Unable to decode config into struct: %v", err))
	}

	return config
}
