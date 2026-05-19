package redirect_service

import (
	"context"

	core_cache "github.com/tiotsyodev/url-shortener.git/internal/core/cache"
	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
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
}

func NewRedirectService(urlRepo UrlRepo, statRepo StatRepo, rds core_cache.Cache) *RedirectService {
	return &RedirectService{
		UrlRepo: urlRepo,
		StatRepo: statRepo,
		CacheClient: rds,
	}
}