package config

import (
	"fmt"
	"github.com/spf13/viper"
	gLogger "gorm.io/gorm/logger"
	"log"
)

var Configs Config

type Config struct {
	Logger   LoggerConfig   `mapstructure:"log"`
	Server   ServerConfig   `mapstructure:"server"`
	DataBase DataBaseConfig `mapstructure:"database"`
}

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

type DataBaseConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Port     string `mapstructure:"port"`
	SSLMode  string `mapstructure:"ssl_mode"`
	Timezone string `mapstructure:"timezone"`
	Logger   struct {
		SlowThreshold             int    `mapstructure:"slow_threshold"`
		LogLevel                  string `mapstructure:"log_level"`
		IgnoreRecordNotFoundError bool   `mapstructure:"ignore_record_not_found_error"`
		ParameterizedQueries      bool   `mapstructure:"parameterized_queries"`
	} `mapstructure:"logger"`
}

func (dbc DataBaseConfig) GetLogLevel() gLogger.LogLevel {
	switch dbc.Logger.LogLevel {
	case "silent":
		return gLogger.Silent
	case "error":
		return gLogger.Error
	case "warn":
		return gLogger.Warn
	default:
		return gLogger.Info
	}
}

func (dbc DataBaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbc.Host, dbc.User, dbc.Password, dbc.Name, dbc.Port, dbc.SSLMode, dbc.Timezone,
	)
}

func InitConfig(configFile string) {
	v := viper.New()

	v.SetDefault("log", map[string]string{"level": "debug"})
	v.SetDefault("server", map[string]string{"port": "8080", "host": "127.0.0.1"})
	v.SetDefault("database", map[string]any{
		"host": "localhost", "user": "postgres", "password": "postgres",
		"db_name": "user", "port": 5432, "ssl_mode": "disable", "timezone": "utc"},
	)

	v.SetConfigFile(configFile)

	log.Printf("Parsing Config File")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config file, err:%s", err)
	}

	err = v.Unmarshal(&Configs)
	if err != nil {
		log.Printf("error parsing configs to Config struct:%s", err)
	}
}
