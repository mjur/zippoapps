package http

import (
	"encoding/json"
	"net/http"

	"github.com/mjur/zippo/pkg/configuration/log"

	"github.com/pkg/errors"
)

// WriteError writes the given error as the response body.
func WriteError(w http.ResponseWriter, log *log.Logger, statusCode int, err string) {
	log.Errorf("Response %d: %v", statusCode, err)
	w.Header().Set("Content-Type", "application/json")
	Write(w, log, statusCode, err)
}

// Write writes out the response body.
func Write(w http.ResponseWriter, log *log.Logger, statusCode int, returnBody interface{}) {
	// marshal response
	bytes, err := json.Marshal(returnBody)
	if err != nil {
		WriteError(w, log, http.StatusInternalServerError, errors.Wrapf(err, "unable to marshal response").Error())

		return
	}

	// log raw return
	if statusCode == http.StatusOK {
		log.Debugf("Response %d: %s", statusCode, bytes)
	}

	// write header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// write return body
	_, err = w.Write(bytes)
	if err != nil {
		log.Errorf("write error %v", err)
	}
}
