package config

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

var DefaultConfig = Config{
	Data: Data{
		MongoDB: MongoDB{
			ConnectionString: "mongodb://localhost:27017",
			Database:         "payment",
			Collection:       "transactions",
		},
	},
	Logger: Logger{
		Env: "dev",
	},
}

type Config struct {
	Server Server `koanf:"server"`
	Data   Data   `koanf:"data"`
	Logger Logger `koanf:"logger"`
}
type GRPC struct {
	Addr string `koanf:"addr"`
}
type HTTP struct {
	Addr string `koanf:"addr"`
}
type Server struct {
	GRPC GRPC `koanf:"grpc"`
	HTTP HTTP `koanf:"http"`
}
type MongoDB struct {
	ConnectionString string `koanf:"connection_string"`
	Database         string `koanf:"database"`
	Collection       string `koanf:"collection"`
}
type Data struct {
	MongoDB MongoDB `koanf:"mongodb"`
}
type Logger struct {
	Env string `koanf:"env"`
}

var (
	k      = koanf.New(".")
	parser = toml.Parser()
)

func Configure(p string) (Config, error) {
	c := DefaultConfig
	if p != "" {
		if err := k.Load(file.Provider(p), parser); err != nil {
			return c, fmt.Errorf("loading config: %w", err)
		}
	}
	if err := k.Load(env.Provider("PAYMENT_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "PAYMENT_")), "__", ".", -1)
	}), nil); err != nil {
		return c, fmt.Errorf("loading config: %w", err)
	}

	err := k.Unmarshal("", &c)
	if err != nil {
		return c, fmt.Errorf("unmarshalling config: %w", err)
	}
	return c, nil
}
