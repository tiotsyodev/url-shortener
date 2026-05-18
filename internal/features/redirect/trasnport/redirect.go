package redirect_transport

import (
	"log/slog"
	"net"
	"net/http"

	"github.com/mileusna/useragent"
	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
	core_response_handler "github.com/tiotsyodev/url-shortener.git/internal/core/transport/http/response"
)

func (h *RedirectHandler) Redirect(rw http.ResponseWriter, r *http.Request) {
	handler := core_response_handler.NewResponseHandler(h.Log, rw)

	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	reqUserAgent := r.UserAgent()
	device := parseDevice(reqUserAgent)

	alias := r.PathValue("alias")

	if len([]rune(alias)) < 3 || len([]rune(alias)) > 100 {
		handler.RespondError(core_domain.ErrInvalidInput, "invalid alias")
		return
	}

	url, err := h.RedirectService.Redirect(r.Context(), alias)	
	if err != nil {
		handler.RespondError(err, "failed to get url")
		return
	}

	statDom := core_domain.NewStatDomain(-1, url.Id, ip, device, reqUserAgent)

	err = h.StatService.SaveClick(r.Context(), statDom) 
	if err != nil {
		h.Log.Warn("failed to save click", slog.String("error", err.Error()))
	}

	http.Redirect(rw, r, url.Url, http.StatusPermanentRedirect)
}

func parseDevice(reqUserAgent string) string {
	ua := useragent.Parse(reqUserAgent)
	if ua.Device != "" {
		return ua.Device
	}

	switch {
	case ua.Mobile:
		return "mobile"
	case ua.Desktop:
		return "desktop"
	case ua.Bot:
		return "bot"
	case ua.Tablet:
		return "tablet"
	default: 
		return "unknown"
	}
	
}