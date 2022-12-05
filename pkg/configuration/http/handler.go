package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mjur/zippo/pkg/configuration"
)

// New returns a new template handler.
func New(cfg *configuration.Config, service configuration.Service) Handler {
	return &handler{
		config:  cfg,
		service: service,
	}
}

//go:generate moq -out ./mocks/handler.go -pkg mocks  . Handler

// Handler represent an http handler.
type Handler interface {
	GetMainSku(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

type handler struct {
	config  *configuration.Config
	service configuration.Service
}

// GetMainSku is a http handler method used for retrieving main skus.
func (h *handler) GetMainSku(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	packageName := p.ByName("package")
	countryCode := r.Header.Get("X-Appengine-Country")
	if countryCode == "" {
		countryCode = configuration.UnassignredCountryCode
	}
	res, err := h.service.GetMainSku(r.Context(), packageName, countryCode)
	if err != nil {
		status, e := GetErrorResponse(err)
		WriteError(w, h.config.Log, status, e.Message)
		return
	}

	Write(w, h.config.Log, http.StatusOK, res.Sku)
}
