package config

import (
	"fmt"
	"github.com/amirhosseinmoayedi/Project-template/internall/log"
	"github.com/spf13/viper"
)

var Configs = struct {
	Logger LoggerConfig `mapstructure:"log"`
	Server ServerConfig `mapstructure:"server"`
}{}

type LoggerConfig struct {
	Level string `mapstructure:"level"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

func (sc ServerConfig) Address() string {
	return fmt.Sprintf("%v:%v", sc.Host, sc.Port)
}

func InitConfig(configFile string) {
	v := viper.New()

	v.SetDefault("log", map[string]string{"level": "debug"})
	v.SetDefault("server", map[string]string{"port": "8080", "host": "127.0.0.1"})

	v.SetConfigFile(configFile)

	err := v.ReadInConfig()
	if err != nil {
		log.Logger.WithError(err).Fatal("error reading config file")
	}

	err = v.Unmarshal(Configs)
	if err != nil {
		log.Logger.WithError(err).Warn("error parsing configs to Config struct")
	}
}
