package stats_service

import (
	"context"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

type StatRepo interface {
	SaveClick(ctx context.Context, dom core_domain.StatDomain) error 
}

type StatService struct {
	StatRepo StatRepo
}

func NewStatService(statRepo StatRepo) *StatService {
	return &StatService{
		StatRepo: statRepo,
	}
}