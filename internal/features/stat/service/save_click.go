package stats_service

import (
	"context"
	"fmt"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

func (s *StatService) SaveClick(ctx context.Context, dom core_domain.StatDomain) error {
	if err := s.StatRepo.SaveClick(ctx, dom); err != nil {
		return fmt.Errorf("failed to save statistic: %w", err)
	}

	return  nil
}