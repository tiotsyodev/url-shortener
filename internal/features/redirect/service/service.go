package redirect_service

import (
	"context"

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
}

func NewRedirectService(urlRepo UrlRepo, statRepo StatRepo) *RedirectService {
	return &RedirectService{
		UrlRepo: urlRepo,
		StatRepo: statRepo,
	}
}