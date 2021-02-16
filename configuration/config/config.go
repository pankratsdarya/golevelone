package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config is configuration structure
type Config struct {
	MainPort    int    `envconfig:"PORT" default:"8080" required:"true"`
	DBURL       string `envconfig:"DB_URL" default:"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable" required:"true"`
	JaegerURL   string `envconfig:"JAEGER_URL" default:"http://jaeger:16686" required:"true"`
	SentryURL   string `envconfig:"SENTRY_URL" default:"http://sentry:9000" required:"true"`
	KafkaBroker string `envconfig:"KAFKA_BROKER" default:"kafka:9092" required:"true"`
	SomeAppID   string `envconfig:"SOME_APP_ID" default:"IDtest" required:"true"`
	SomeAppKEY  string `envconfig:"SOME_APP_KEY" default:"KEYtest" required:"true"`
}

// GetConfig gets base configuration
func GetConfig() (*Config, error) {
	conf := Config{}
	err := envconfig.Process("", &conf)
	if err != nil {
		return nil, fmt.Errorf("can't process the config: %w", err)
	}
	return &conf, err
}
