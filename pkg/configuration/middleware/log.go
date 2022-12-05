package middleware

import (
	"context"

	"github.com/mjur/zippo/pkg/configuration"
	"github.com/mjur/zippo/pkg/configuration/log"
)

type logMiddleware struct {
	config *configuration.Config
	next   configuration.Service
	logger *log.Logger
}

// NewLogMiddleware returns a new logging middleware for the configuration service.
func NewLogMiddleware(config *configuration.Config, next configuration.Service, logger *log.Logger) configuration.Service {
	m := logMiddleware{
		next:   next,
		config: config,
		logger: logger,
	}

	return &m
}

// GetMainSku logs the inputs and outputs of the GetMainSku service method.
func (m *logMiddleware) GetMainSku(ctx context.Context, packageName, countryCode string) (*configuration.MainSku, error) {
	m.logger.Log = m.logger.Log.With().
		Str("useCase", "GetMainSku").
		Interface("packageName", packageName).
		Interface("countryCode", countryCode).Logger()

	m.logger.Info("Received get main sku request")

	res, err := m.next.GetMainSku(ctx, packageName, countryCode)

	m.logger.Log = m.logger.Log.With().
		Interface("result", res).
		Err(err).Logger()
	m.logger.Info("Returned response")

	return res, err
}
