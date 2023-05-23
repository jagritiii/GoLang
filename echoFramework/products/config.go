package products

type configDatabase struct {
	AppName  string `env:"APP_NAME" env-default:"PRODUCTS"`
	Appenv   string `env:"APP_ENV" env-default:"DEV"`
	Port     string `env:"MY_APP_PORT" env-default:"8080"`
	Host     string `env:"HOST" env-default:"Localhost"`
	LogLevel string `env:"APP_NAME" env-default:"\ERROR"`
}

var cfg configDatabase
