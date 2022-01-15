package logger

import (
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

var (
	dev  = "dev"
	prod = "prod"
)

func New(env string) (log logr.Logger, err error) {
	var zapLog *zap.Logger

	switch strings.ToLower(env) {
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
		err = fmt.Errorf("invalid env '%s', must be one of '%s' or '%s'", env, dev, prod)
		return
	}
	log = zapr.NewLogger(zapLog)
	return
}
