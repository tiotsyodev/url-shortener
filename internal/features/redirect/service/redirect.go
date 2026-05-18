package redirect_service

import (
	"context"
	"fmt"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

func (s *RedirectService) Redirect(ctx context.Context, alias string) (core_domain.UrlDomain, error) {
	url, err := s.UrlRepo.GetUrl(ctx, alias)
	if err != nil {
		return core_domain.UrlDomain{}, fmt.Errorf("failed to get url for redirect: %w", err)
	}

	return url, nil
}