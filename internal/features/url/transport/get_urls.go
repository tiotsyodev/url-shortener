package url_transport

import (
	"net/http"
	"strconv"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
	core_response_handler "github.com/tiotsyodev/url-shortener.git/internal/core/transport/http/response"
)

func (h *UrlHandler) GetUrls(rw http.ResponseWriter, r *http.Request) {
	handler := core_response_handler.NewResponseHandler(h.Log, rw)

	limitStr := r.URL.Query().Get("limit")
	pageStr := r.URL.Query().Get("page")

	limit := 10
	page := 1

	if limitStr != "" {
		v, err := strconv.Atoi(limitStr)
		if err != nil {
			handler.RespondError(core_domain.ErrInvalidInput, "limit must be a number")
			return
		}
		limit = v
	}

	if pageStr != "" {
		v, err := strconv.Atoi(pageStr)
		if err != nil {
			handler.RespondError(core_domain.ErrInvalidInput, "page must be a number")
			return
		}
		page = v
	}

	offset := (page - 1) * limit

	urls, err := h.UserService.GetUrls(r.Context(), limit, offset)
	if err != nil { 
		handler.RespondError(err, "failed to get urls")
		return
	}

	handler.ResponseJSON(http.StatusOK, urls)
}