package server

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const EnvFileName = ".env"

type IENV interface{}

type ENVConfig struct {
	LogLevel logrus.Level

	Host    string `mapstructure:"host"`
	ENV     string `mapstructure:"env"`
	Service string `mapstructure:"service"`

	DBEngine   string `mapstructure:"DB_ENGINE"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBPort     string `mapstructure:"DB_PORT"`
}

type ENVType struct {
	config *ENVConfig
}

func NewEnv() IENV {
	return NewENVPath(".")
}

func NewENVPath(path string) IENV {
	viper.SetConfigName(EnvFileName)
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	envKeys := []string{
		"LOG_HOST",
		"HOST", "ENV", "SERVICE",
		"SENTRY_DSN", "DB_ENGINE", "DB_HOST",
	}

	for _, key := range envKeys {
		viper.BindEnv(key)
	}

	env := &ENVConfig{}
	err := viper.Unmarshal(env)
	if err != nil {
	}

	env.LogLevel, _ = logrus.ParseLevel(viper.GetString("log_level"))
	return &ENVType{
		config: env,
	}
}
