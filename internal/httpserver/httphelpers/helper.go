package httphelpers

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go-labs-game-platform/internal/models"
)

type Err struct {
	Message string `json:"message"`
} //@name Error

func toHTTPError(err error) (int, string) {
	switch {
	case errors.Is(err, models.ErrNotFound):
		return http.StatusNotFound, err.Error()
	case errors.Is(err, models.ErrUnauthorized):
		return http.StatusUnauthorized, err.Error()
	case errors.Is(err, models.ErrForbidden):
		return http.StatusForbidden, err.Error()
	case errors.Is(err, models.ErrConflict):
		return http.StatusConflict, err.Error()
	case errors.Is(err, models.ErrBadRequest):
		return http.StatusBadRequest, err.Error()
	default:
		return http.StatusInternalServerError, "internal server error"
	}
}

func WriteError(w http.ResponseWriter, r *http.Request, statusOverride int, err error) {
	var msg string
	if statusOverride == 0 {
		statusOverride, msg = toHTTPError(err)
	} else {
		msg = err.Error()
	}

	if statusOverride == http.StatusInternalServerError {
		WriteJSON(w, r, statusOverride, Err{Message: "Internal server error"})
		logrus.Errorf("Internal server error: %s", err)
		return
	}

	WriteJSON(w, r, statusOverride, Err{Message: msg})
}

func WriteJSON(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		logrus.Errorf("Marshal %v: %s", data, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err = w.Write(b); err != nil {
		logrus.Errorf("Write HTTP response %s", err)
	}
}
