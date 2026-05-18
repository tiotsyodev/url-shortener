package core_response_handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

type ResponseHandler struct {
	log *slog.Logger
	rw http.ResponseWriter
}

func NewResponseHandler(log *slog.Logger, rw http.ResponseWriter) ResponseHandler {
	return ResponseHandler{
		log: log,
		rw: rw,
	}
}

func (h *ResponseHandler) ResponseJSON(code int, data any) {
	h.rw.Header().Set("Content-Type", "application/json")
    h.rw.WriteHeader(code)
    if err := json.NewEncoder(h.rw).Encode(data); err != nil {
		h.log.Error("write HTTP response", "error", err)
	}
}

func (h *ResponseHandler) RespondError(err error, msg string) {
	h.log.Warn("request error", slog.String("error", err.Error()), slog.String("msg", msg))
    res := map[string]string{
        "error":   err.Error(),
        "message": msg,
    }

    switch {
    case errors.Is(err, core_domain.ErrNotFound):
        h.ResponseJSON(http.StatusNotFound, res)
    case errors.Is(err, core_domain.ErrAlreadyExists):
        h.ResponseJSON(http.StatusConflict, res)
    case errors.Is(err, core_domain.ErrUnauthorized):
        h.ResponseJSON(http.StatusUnauthorized, res)
    case errors.Is(err, core_domain.ErrForbidden):
        h.ResponseJSON(http.StatusForbidden, res)
    case errors.Is(err, core_domain.ErrTimeout):
        h.ResponseJSON(http.StatusGatewayTimeout, res)
    case errors.Is(err, core_domain.ErrInvalidInput):
        h.ResponseJSON(http.StatusBadRequest, res)
    default:
        h.log.Error("unhandled error", slog.String("error", err.Error()))
        h.ResponseJSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
    }
}

func (h *ResponseHandler) RespondValidationError(err error) {
    h.log.Warn("validation error", slog.String("error", err.Error()))
    h.ResponseJSON(http.StatusBadRequest, map[string]string{
        "error": formatValidationErrors(err),
    })
}

func formatValidationErrors(err error) string {
    var ve validator.ValidationErrors
    if !errors.As(err, &ve) {
        return err.Error()
    }

    msgs := make([]string, len(ve))
    for i, fe := range ve {
        msgs[i] = fmt.Sprintf("%s is invalid", strings.ToLower(fe.Field()))
    }
    return strings.Join(msgs, ", ")
}

