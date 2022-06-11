package configs

import (
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var (
	AuthorizationHeader = "Authorization"
	defaultDotEnvKey    = "DOTENV_FILE"
	Config              *ConfigType
	origConfig          ConfigType
)

type ConfigType struct {
	Port       int    `env:"PORT" envDefault:"8080"`
	JWTSecret  string `env:"JWT_SECRET" envDefault:"test"`
	JWTCookies string `env:"JWT_COOKIES" envDefault:"authToken"`

	DATABASE_URL string `env:"DATABASE_URL" envDefault:""`
	REDIS_URL    string `env:"REDIS_URL" envDefault:"localhost:3379"`

	LineLoginCallback      string `env:"LINE_CALLBACK_URL" envDefault:"http://localhost:8080/api/v1/auth/line/callback"`
	LineClientID           string `env:"LINE_CLIENT_ID" envDefault:""`
	LineSecretID           string `env:"LINE_SECRET_ID" envDefault:""`
	LINE_WEB_CALLBACK_PATH string `env:"LINE_WEB_CALLBACK_PATH" envDefault:"/user/line/callback"`

	LogFilePath       string `env:"LOG_PATH" envDefault:"/var/log/cjtim-backend-go.log"`
	GCLOUD_CREDENTIAL string `env:"GCLOUD_CREDENTIAL" envDefault:"./configs/serviceAcc.json"`
}

func init() {
	log.Default().Println("Initial config...")
	fp, err := os.Create("/var/log/cjtim-backend-go.log")
	if err != nil {
		os.Setenv("LOG_PATH", "./log/cjtim-backend.go.log")
	}
	defer fp.Close()

	cfg := ConfigType{}
	envFile := os.Getenv(defaultDotEnvKey)
	if envFile == "" {
		envFile = ".env"
	}
	_ = godotenv.Load(envFile)
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	Config = &cfg
	origConfig = cfg
}

func RestoreConfigMock() {
	Config = &origConfig
}
