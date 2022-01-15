package config

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

type Config struct {
	Server  Server  `koanf:"server"`
	MongoDB MongoDB `koanf:"mongodb"`
	Logger  Logger  `koanf:"logger"`
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
}
type Logger struct {
	Env string `koanf:"env"`
}

var (
	k      = koanf.New(".")
	parser = toml.Parser()
)

func Configure(p string) (*Config, error) {
	if err := k.Load(file.Provider(p), parser); err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}
	if err := k.Load(env.Provider("PAYMENT_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "PAYMENT_")), "_", ".", -1)
	}), nil); err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}

	c := new(Config)
	err := k.Unmarshal("", c)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling config: %w", err)
	}
	return c, nil
}
