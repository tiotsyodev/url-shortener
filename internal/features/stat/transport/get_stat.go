package stat_transport

import (
	"net/http"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
	core_response_handler "github.com/tiotsyodev/url-shortener.git/internal/core/transport/http/response"
)

func (h *StatHandler) GetStatsByAlias(rw http.ResponseWriter, r *http.Request) {
	handler := core_response_handler.NewResponseHandler(h.Log, rw)

	alias := r.PathValue("alias")
	if alias == "" {
		handler.RespondError(core_domain.ErrInvalidInput, "alias is required")
		return
	}

	stats, err := h.StatService.GetStats(r.Context(), alias)
	if err != nil {
		handler.RespondError(err, "failed to get stats")
		return
	}

	handler.ResponseJSON(http.StatusOK, stats)
}