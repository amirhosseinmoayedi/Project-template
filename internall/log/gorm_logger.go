package log

import (
	"github.com/amirhosseinmoayedi/Project-template/internall/config"
	gLogger "gorm.io/gorm/logger"
	"time"
)

func GetGormLogger() gLogger.Interface {
	configs := config.Configs.DataBase
	return gLogger.New(
		Logger,
		gLogger.Config{
			SlowThreshold:             time.Duration(configs.Logger.SlowThreshold) * time.Second,
			LogLevel:                  configs.GetLogLevel(),
			IgnoreRecordNotFoundError: configs.Logger.IgnoreRecordNotFoundError,
			ParameterizedQueries:      configs.Logger.ParameterizedQueries,
		},
	)
}
