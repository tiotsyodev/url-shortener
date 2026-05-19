package stat_transport

import (
	"context"
	"log/slog"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
	core_http_server "github.com/tiotsyodev/url-shortener.git/internal/core/transport/http/server"
)

type StatService interface {
	GetStats(ctx context.Context, alias string) (core_domain.StatsByAlias, error)
}

type StatHandler struct {
	StatService StatService
	Log *slog.Logger
}

func NewStatHandler(log *slog.Logger, statService StatService) StatHandler {
	return StatHandler{
		StatService: statService,
		Log: log,
	}
}

func (h *StatHandler) GetRoutes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			HandleFunc: h.GetStatsByAlias,
			Method: "GET",
			Url: "/stat/{alias}",
		},
	}
}
