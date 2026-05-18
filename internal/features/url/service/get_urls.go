package url_service

import (
	"context"
	"fmt"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

func (s *URLService) GetUrls(ctx context.Context, limit int, offset int) ([]core_domain.UrlDomain, error) {
	urls, err := s.Repo.GetUrls(ctx, limit, offset)
	if err != nil {
		return []core_domain.UrlDomain{}, fmt.Errorf("unabled to get urls: %w", err)
	}
	return urls, nil
}