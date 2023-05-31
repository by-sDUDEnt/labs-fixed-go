package config

import (
	"sync"
	"time"

	goconfig "github.com/Yalantis/go-config"
)

var (
	config Config
	once   sync.Once
)

type (
	Config struct {
		HTTP     HTTP     `json:"http"`
		Security Security `json:"security"`
		DB       DB       `json:"db"`
		Redis    Redis    `json:"redis"`
	}

	Security struct {
		AdminSecret         string        `json:"admin_secret" envconfig:"ADMIN_SECRET" default:"a"`
		UserAbsoluteTimeout time.Duration `json:"user_absolute_timeout" envconfig:"USER_ABSOLUTE_TIMEOUT" default:"24h"`
		UserIdleTimeout     time.Duration `json:"user_idle_timeout" envconfig:"USER_IDLE_TIMEOUT" default:"24h"`
	}

	HTTP struct {
		Host          string        `json:"host" envconfig:"HTTP_HOST" default:"localhost"`
		Port          uint          `json:"port" envconfig:"HTTP_POST" default:"8080"`
		ReadTimeout   time.Duration `json:"read_timeout" envconfig:"HTTP_READ_TIMEOUT" default:"15s"`
		WriteTimeout  time.Duration `json:"write_timeout" envconfig:"HTTP_WRITE_TIMEOUT" default:"15s"`
		URLPathPrefix string        `json:"url_path_prefix" envconfig:"HTTP_URL_PATH_PREFIX" default:"/api/v1"`
	}

	DB struct {
		Host     string `json:"host" envconfig:"DB_HOST" default:"localhost"`
		Port     string `json:"port" envconfig:"DB_PORT" default:"5432"`
		User     string `json:"users" envconfig:"DB_USER" default:"go-user"`
		Password string `json:"password" envconfig:"DB_PASSWORD" default:"go-password"`
		Database string `json:"database" envconfig:"DB_DATABASE" default:"go-db"`
		Debug    bool   `json:"debug" envconfig:"DB_DEBUG" default:"false"`
	}

	Redis struct {
		Host     string `json:"host" envconfig:"REDIS_HOST" default:"localhost"`
		Port     uint32 `json:"port" envconfig:"REDIS_PORT" default:"6379"`
		Password string `json:"password" envconfig:"REDIS_PASSWORD" default:""`
	}
)

func New() (Config, error) {
	var cfg Config

	if err := goconfig.Init(&cfg, "config.json"); err != nil {
		return Config{}, err
	}

	config = cfg

	return cfg, nil
}

func Get() Config {
	once.Do(func() {
		cfg, err := New()
		if err != nil {
			panic(err)
		}
		config = cfg
	})

	return config
}
