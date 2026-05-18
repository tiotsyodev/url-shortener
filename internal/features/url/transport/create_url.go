package url_transport

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
	core_response_handler "github.com/tiotsyodev/url-shortener.git/internal/core/transport/http/response"
)

type CreateURLRequest struct {
	Alias string `json:"alias" validate:"required,min=3,max=100"`
	Url string	`json:"url" validate:"required,url"`
}

func (h *UrlHandler) CreateURL(rw http.ResponseWriter, r *http.Request) {
	responseHandler := core_response_handler.NewResponseHandler(h.Log, rw)

	var req CreateURLRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseHandler.RespondError(err, "invalid json")
		return
	}

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(req); err != nil {
		responseHandler.RespondValidationError(err)
		return
	}

	domainURL := newUrlFromDto(req.Alias, req.Url)

	createdDomain, err := h.UserService.CreateURL(r.Context(), domainURL)
	if err != nil {
		responseHandler.RespondError(err, "failed to create url")
        return
	}
	h.Log.Info("url created", slog.String("alias", createdDomain.Alias))
	responseHandler.ResponseJSON(http.StatusCreated, createdDomain)
}

func newUrlFromDto(alias string, url string) core_domain.UrlDomain {
	return core_domain.NewUrl(-1, alias, url )
}