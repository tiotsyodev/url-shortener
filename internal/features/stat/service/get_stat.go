package stats_service

import (
	"context"
	"fmt"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

func (s *StatService) GetStats(ctx context.Context, alias string) (core_domain.StatsByAlias, error) {
	stats, err := s.StatRepo.GetStats(ctx, alias)
	if err != nil {
		return core_domain.StatsByAlias{}, fmt.Errorf("get stats: %w", err)
	}
	return stats, nil
}