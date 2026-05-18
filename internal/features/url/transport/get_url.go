package url_transport

import (
	"fmt"
	"net/http"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
	core_response_handler "github.com/tiotsyodev/url-shortener.git/internal/core/transport/http/response"
)

func (h *UrlHandler) GetUrl(rw http.ResponseWriter, r *http.Request) {
	resHandler := core_response_handler.NewResponseHandler(h.Log, rw)

	alias := r.PathValue("alias")

	if alias == "" {
		resHandler.RespondError(fmt.Errorf("alias is required: %w", core_domain.ErrInvalidInput), "invalid id")
		return
	}

	if len([]rune(alias)) < 3 || len([]rune(alias)) > 100 {
		resHandler.RespondError(fmt.Errorf("invalid alias: %w", core_domain.ErrInvalidInput), "alias must be a greate than 3 and lower than 100")
		return
	}

	urlDomain, err := h.UserService.GetUrl(r.Context(), alias)
	if err != nil {
		resHandler.RespondError(err, "failed to get url")
		return
	}


	resHandler.ResponseJSON(http.StatusOK, urlDomain)
}