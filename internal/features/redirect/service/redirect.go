package redirect_service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

func (s *RedirectService) Redirect(ctx context.Context, alias string) (core_domain.UrlDomain, error) {
	cached, err := s.CacheClient.Get(ctx, alias)
	if err == nil  {
		var cachedDomain core_domain.UrlDomain
		if err := json.Unmarshal([]byte(cached), &cachedDomain); err == nil {
			s.Metrics.RedirectCounter.WithLabelValues("cache").Inc()
			return cachedDomain, nil
		}
	}

	url, err := s.UrlRepo.GetUrl(ctx, alias)
	if err != nil {
		return core_domain.UrlDomain{}, fmt.Errorf("failed to get url for redirect: %w", err)
	}

	if data, err := json.Marshal(url); err == nil {
        s.CacheClient.Set(ctx, alias, string(data), time.Hour)
    }
	s.Metrics.RedirectCounter.WithLabelValues("database").Inc()
	return url, nil
}