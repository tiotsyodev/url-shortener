package redirect_transport

import (
	"context"
	"log/slog"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
	core_http_server "github.com/tiotsyodev/url-shortener.git/internal/core/transport/http/server"
)

type RedirectService interface {
	Redirect(ctx context.Context, alias string) (core_domain.UrlDomain, error)
}

type StatService interface {
	SaveClick(ctx context.Context, dom core_domain.StatDomain) error
}

type RedirectHandler struct {
	Log *slog.Logger
	RedirectService RedirectService
	StatService StatService
}

func NewRedirectHandler(log *slog.Logger, redirectService RedirectService, statService StatService) RedirectHandler {
	return RedirectHandler{
		Log: log,
		RedirectService: redirectService,
		StatService: statService,
	}
}

func (h *RedirectHandler) GetRoutes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			HandleFunc: h.Redirect,
			Method: "GET",
			Url: "/{alias}",
		},
	}
}