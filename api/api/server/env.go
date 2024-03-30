package server

const EnvFileName = ".env"

type ENVConfig struct {
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
