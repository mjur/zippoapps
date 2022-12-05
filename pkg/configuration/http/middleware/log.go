package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mjur/zippo/pkg/configuration"
	handlers "github.com/mjur/zippo/pkg/configuration/http"
	"github.com/mjur/zippo/pkg/configuration/log"
)

type logMiddleware struct {
	config *configuration.Config
	next   handlers.Handler
	logger *log.Logger
}

// NewLogMiddleware returns a new logging middleware for the template handler.
func NewLogMiddleware(config *configuration.Config, next handlers.Handler, logger *log.Logger) handlers.Handler {
	m := logMiddleware{
		next:   next,
		config: config,
		logger: logger,
	}
	return &m
}

// GetMainSku logs the request for the GetMainSku handler.
func (m *logMiddleware) GetMainSku(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	m.logger.Log = m.logger.Log.With().
		Str("useCase", "GetMainSku").
		Interface("requestHeaders", r.Header).Logger()
	m.logger.Info("Received get sku request")

	m.next.GetMainSku(w, r, p)
}
