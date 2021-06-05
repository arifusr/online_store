package app

type Config struct {
	Name string `env:"APPLICATION_NAME,required"`
	Port string `env:"PORT,required"`
	ConfigDatabase
	ConfigJWT
}

type ConfigDatabase struct {
	Host     string `env:"DATABASE_HOST,required"`
	Port     string `env:"DATABASE_PORT,required"`
	Username string `env:"DATABASE_USERNAME,required"`
	Password string `env:"DATABASE_PASSWORD,required"`
	Name     string `env:"DATABASE_NAME,required"`
}

type ConfigJWT struct {
	Secret string `env:"JWT_SECRET,required"`
}
