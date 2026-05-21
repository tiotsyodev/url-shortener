package url_service

import (
	"context"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
	core_metrics "github.com/tiotsyodev/url-shortener.git/internal/core/metrics"
)

type URLRepo interface {
	CreateURL(ctx context.Context, d core_domain.UrlDomain) (core_domain.UrlDomain, error)
	GetUrl(ctx context.Context, alias string) (core_domain.UrlDomain, error)
	GetUrls(ctx context.Context, limit int, offset int) ([]core_domain.UrlDomain, error)
	DeleteUrl(ctx context.Context, id int) (core_domain.UrlDomain, error)
	UpdateUrl(ctx context.Context, dom core_domain.UpdateUrlDomain) (core_domain.UrlDomain, error) 
}

type URLService struct {
	Repo URLRepo
	Metrics core_metrics.Metrics
}

func NewUserService(repo URLRepo, metrics core_metrics.Metrics) *URLService {
	return &URLService{
		Repo: repo,
		Metrics: metrics,
	}
}