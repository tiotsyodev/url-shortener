package url_service

import (
	"context"
	"fmt"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

func (s *URLService) DeleteUrl(ctx context.Context, id int) (core_domain.UrlDomain, error) {
	urlDomain,err := s.Repo.DeleteUrl(ctx, id)
	if err != nil {
		return core_domain.UrlDomain{}, fmt.Errorf("delete url: %w", err)
	}

	return urlDomain, nil
}