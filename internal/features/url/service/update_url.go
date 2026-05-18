package url_service

import (
	"context"
	"fmt"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

func (s *URLService) UpdateUrl(ctx context.Context, dom core_domain.UpdateUrlDomain) (core_domain.UrlDomain, error) {
	domain, err := s.Repo.UpdateUrl(ctx, dom)
	if err != nil {
		return core_domain.UrlDomain{}, fmt.Errorf("update url id - %d: %w", dom.Id, err)
	}

	return domain, nil

}