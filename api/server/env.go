package server

func NewENV() {
	panic("implement me")
}

type ENVConfig struct {
	Host       string `mapstructure:"host"`
	ENV        string `mapstructure:"env"`
	Service    string `mapstructure:"service"`
	DBDriver   string `mapstructure:"db_driver"`
	DBHost     string `mapstructure:"db_host"`
	DBName     string `mapstructure:"db_name"`
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBPort     string `mapstructure:"db_port"`
}
