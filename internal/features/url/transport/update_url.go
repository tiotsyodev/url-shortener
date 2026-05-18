package url_transport

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
	core_response_handler "github.com/tiotsyodev/url-shortener.git/internal/core/transport/http/response"
)

type UpdateURLRequest struct {
	Alias *string `json:"alias" validate:"omitempty,min=3,max=100"`
	Url *string `json:"url" validate:"omitempty,url"`
}

func (h *UrlHandler) UpdateUrl(rw http.ResponseWriter, r *http.Request) {
	handler := core_response_handler.NewResponseHandler(h.Log, rw)

	id := r.PathValue("id")

	if id == "" {
		handler.RespondError(core_domain.ErrInvalidInput, "id is required")
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		handler.RespondError(core_domain.ErrInvalidInput, "id must be a number")
		return
	}

	var req UpdateURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.RespondError(err, "invalid json")
		return
	}

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(req); err != nil {
		handler.RespondValidationError(err)
		return
	}


	updatedUrl, err := h.UserService.UpdateUrl(r.Context(), core_domain.NewUpdateDomain(intId, req.Url, req.Alias))
	
	if err != nil {
		handler.RespondError(err, "failed to update url")
		return
	}

	handler.ResponseJSON(http.StatusOK, updatedUrl)

}
