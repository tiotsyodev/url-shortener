package url_service

import (
	"context"
	"fmt"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

func (s *URLService) GetUrl(ctx context.Context, alias string) (core_domain.UrlDomain, error) {
	urlDomain, err := s.Repo.GetUrl(ctx, alias)
	if err != nil {
		return core_domain.UrlDomain{}, fmt.Errorf("get url service: %w", err)
	}

	return urlDomain, nil
}