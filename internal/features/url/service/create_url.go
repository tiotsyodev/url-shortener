package url_service

import (
	"context"
	"fmt"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

func (s *URLService) CreateURL(ctx context.Context, d core_domain.UrlDomain ) (core_domain.UrlDomain, error) {

	if err := d.Validate(); err != nil {
		return core_domain.UrlDomain{}, fmt.Errorf("validate url domain: %w", err)
	}
	
	dom, err := s.Repo.CreateURL(ctx, d)
	if err != nil {
		return core_domain.UrlDomain{}, fmt.Errorf("create url: %w", err)
	}

	return dom, nil
}