package config

type Environment struct {
	AzcapchaApiKey   string `env:"AZCAPCHA_API_KEY,required=true"`
	CorsAllowOrigins string `env:"CORS_ALLOW_ORIGINS,required=true"`
	Port             int    `env:"PORT,required=true"`
}
