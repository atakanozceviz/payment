package logger

import (
	"fmt"
	"payment/internal/config"
	"strings"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(New)

var (
	dev  = "dev"
	prod = "prod"
)

func New(c config.Logger) (log logr.Logger, err error) {
	var zapLog *zap.Logger

	switch strings.ToLower(c.Env) {
	case dev:
		zapLog, err = zap.NewDevelopment()
		if err != nil {
			err = fmt.Errorf("building development logger: %w", err)
			return
		}
	case prod:
		zapLog, err = zap.NewProduction()
		if err != nil {
			err = fmt.Errorf("building production logger: %w", err)
			return
		}
	default:
		err = fmt.Errorf("invalid env '%s', must be one of '%s' or '%s'", c.Env, dev, prod)
		return
	}
	log = zapr.NewLogger(zapLog)
	return
}
