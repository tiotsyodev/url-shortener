package url_transport

import (
	"fmt"
	"net/http"
	"strconv"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
	core_response_handler "github.com/tiotsyodev/url-shortener.git/internal/core/transport/http/response"
)

func (h *UrlHandler) DeleteUrl(rw http.ResponseWriter, r *http.Request) {
	resHandler := core_response_handler.NewResponseHandler(h.Log, rw)

	id := r.PathValue("id")

	if id == "" {
		resHandler.RespondError(fmt.Errorf("id is required: %w", core_domain.ErrInvalidInput), "invalid id")
		return
	}

	intId ,err := strconv.Atoi(id)
	if err != nil {
		resHandler.RespondError(fmt.Errorf("invalid id: %w", core_domain.ErrInvalidInput), "id must be a number")
		return
	}

	urlDomain, err := h.UserService.DeleteUrl(r.Context(), intId)
	if err != nil {
		resHandler.RespondError(err, "failed to get url")
		return
	}


	resHandler.ResponseJSON(http.StatusOK, urlDomain)
}