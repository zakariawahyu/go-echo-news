package config

import "github.com/spf13/viper"

type LoggerConfig struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

func LoadLoggerConfig() LoggerConfig {
	return LoggerConfig{
		Development:       viper.GetBool("LOGGER_DEVELOPMENT"),
		DisableCaller:     viper.GetBool("LOGGER_DISABLE_CALLER"),
		DisableStacktrace: viper.GetBool("LOGGER_DISABLE_STACKTRACE"),
		Encoding:          viper.GetString("LOGGER_ENCODING"),
		Level:             viper.GetString("LOGGER_LEVEL"),
	}
}
