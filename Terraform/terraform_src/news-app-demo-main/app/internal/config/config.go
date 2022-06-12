package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type config struct {
	BindAddr string `env:"NEWS_APP_BIND_ADDR" env-default:"0.0.0.0:8080"`
	NewsLang string `env:"NEWS_APP_NEWS_LANG" env-default:"ru"`
	ApiKey string `env:"NEWS_APP_API_KEY" env-default:""`
}

var once sync.Once

var instance *config

func GetConfig() *config {
	once.Do(func() {
		log.Print("Read application config...")
		instance = &config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Fatal(help)
		}
	})
	return instance
}