package redirect_service

import (
	"context"

	core_cache "github.com/tiotsyodev/url-shortener.git/internal/core/cache"
	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
	core_metrics "github.com/tiotsyodev/url-shortener.git/internal/core/metrics"
)

type UrlRepo interface {
    GetUrl(ctx context.Context, alias string) (core_domain.UrlDomain, error)
}
type StatRepo interface {
    SaveClick(ctx context.Context, dom core_domain.StatDomain) error
}

type RedirectService struct {
	UrlRepo UrlRepo
	StatRepo StatRepo
	CacheClient core_cache.Cache
	Metrics core_metrics.Metrics
}

func NewRedirectService(urlRepo UrlRepo, statRepo StatRepo, rds core_cache.Cache, metrics core_metrics.Metrics) *RedirectService {
	return &RedirectService{
		UrlRepo: urlRepo,
		StatRepo: statRepo,
		CacheClient: rds,
		Metrics: metrics,
	}
}