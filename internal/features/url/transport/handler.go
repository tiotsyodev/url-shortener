package url_transport

import (
	"context"
	"log/slog"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
	core_http_server "github.com/tiotsyodev/url-shortener.git/internal/core/transport/http/server"
)

type UrlService interface {
	CreateURL(ctx context.Context, d core_domain.UrlDomain) (core_domain.UrlDomain, error)
	GetUrl(ctx context.Context, alias string) (core_domain.UrlDomain, error)
	GetUrls(ctx context.Context, limit int, offset int) ([]core_domain.UrlDomain, error)
	DeleteUrl(ctx context.Context, id int) (core_domain.UrlDomain, error)
	UpdateUrl(ctx context.Context, dom core_domain.UpdateUrlDomain) (core_domain.UrlDomain, error)
}

type UrlHandler struct {
	UserService UrlService
	Log *slog.Logger
}

func NewUrlHandler(log *slog.Logger, userService UrlService) UrlHandler {
	return UrlHandler{
		UserService: userService,
		Log: log,
	}
}

func (h *UrlHandler) GetRoutes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			HandleFunc: h.CreateURL,
			Method: "POST",
			Url: "/url",
		},
		{
			HandleFunc: h.GetUrls,
			Method: "GET",
			Url: "/url",
		},
		{
			HandleFunc: h.GetUrl,
			Method: "GET",
			Url: "/url/{alias}",
		},
		{
			HandleFunc: h.DeleteUrl,
			Method: "DELETE",
			Url: "/url/{id}",
		},
		{
			HandleFunc: h.UpdateUrl,
			Method: "PATCH",
			Url: "/url/{id}",
		},
	}
}
