package app

type Config struct {
	Port string `env:"PORT,required"`
	ConfigDatabase
}

type ConfigDatabase struct {
	Host     string `env:"DATABASE_HOST,required"`
	Port     string `env:"DATABASE_PORT,required"`
	Username string `env:"DATABASE_USERNAME,required"`
	Password string `env:"DATABASE_PASSWORD,required"`
	Name     string `env:"DATABASE_NAME,required"`
}
