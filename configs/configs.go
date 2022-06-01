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

	LineLoginCallback string `env:"LINE_CALLBACK_URL" envDefault:"http://localhost:8080/api/v1/auth/line/callback"`
	LineClientID      string `env:"LINE_CLIENT_ID" envDefault:""`
	LineSecretID      string `env:"LINE_SECRET_ID" envDefault:""`

	LoginSuccessURL         string `env:"LOGIN_SUCCESS_URL" envDefault:"http://localhost:3000/user"`
	LineAPIBroadcast        string `envDefault:"https://api.line.me/v2/bot/message/broadcast"`
	LineAPIReply            string `envDefault:"https://api.line.me/v2/bot/message/reply"`
	AirVisualAPINearestCity string `envDefault:"http://api.airvisual.com/v2/nearest_city"`
	AirVisualAPICity        string `envDefault:"http://api.airvisual.com/v2/city"`
	BinanceAccountAPI       string `envDefault:"https://api.binance.com/api/v3/account"`
	RebrandlyAPI            string `envDefault:"https://api.rebrandly.com/v1/links"`
	LogFilePath             string `env:"LOG_PATH" envDefault:"/var/log/cjtim-backend-go.log"`
	GCLOUD_CREDENTIAL       string `env:"GCLOUD_CREDENTIAL" envDefault:"./configs/serviceAcc.json"`
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
