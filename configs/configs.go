package configs

import (
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var (
	AuthorizationHeader = "Authorization"
	defaultDotEnvKey    = "DOTENV_FILE"
	Config              *ConfigType
	origConfig          ConfigType
)

type ConfigType struct {
	Port       int    `env:"PORT,required,notEmpty" envDefault:"8080"`
	JWTSecret  string `env:"JWT_SECRET,required,notEmpty" envDefault:"test"`
	JWTCookies string `env:"JWT_COOKIES,required,notEmpty" envDefault:"authToken"`

	DATABASE_URL string `env:"DATABASE_URL,required,notEmpty" envDefault:""`
	REDIS_URL    string `env:"REDIS_URL,required,notEmpty" envDefault:"localhost:3379"`

	LineLoginCallback      string `env:"LINE_CALLBACK_URL,required,notEmpty" envDefault:"http://localhost:8080/api/v1/auth/line/callback"`
	LineClientID           string `env:"LINE_CLIENT_ID,required,notEmpty" envDefault:""`
	LineSecretID           string `env:"LINE_SECRET_ID,required,notEmpty" envDefault:""`
	LINE_WEB_CALLBACK_PATH string `env:"LINE_WEB_CALLBACK_PATH,required,notEmpty" envDefault:"/user/line/callback"`

	BUCKET_NAME string `env:"BACKET_NAME,required,notEmpty" envDefault:""`

	LogFilePath       string `env:"LOG_PATH,required,notEmpty" envDefault:"/var/log/be-friends-api.log"`
	GCLOUD_CREDENTIAL string `env:"GCLOUD_CREDENTIAL,required,notEmpty" envDefault:"./configs/serviceAcc.json"`
}

func InitConfig() error {
	log.Default().Println("Initial config...")
	fp, err := os.Create("/var/log/be-friends-api.log")
	if err != nil {
		os.Setenv("LOG_PATH", "./log/be-friends-api.log")
	}
	defer fp.Close()

	cfg := ConfigType{}
	envFile := os.Getenv(defaultDotEnvKey)
	if envFile == "" {
		envFile = ".env"
	}
	// Load from .env file
	_ = godotenv.Load(envFile)

	// Parse env to struct
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
		return err
	}
	Config = &cfg
	origConfig = cfg
	return nil
}

func RestoreConfigMock() {
	Config = &origConfig
}
