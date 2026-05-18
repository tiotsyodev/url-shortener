package stats_repo

import (
	core_repo "github.com/tiotsyodev/url-shortener.git/internal/core/repo"
)

type StatRepo struct {
	Pool core_repo.ConnectionPool
}

func NewStatRepo(pool core_repo.ConnectionPool) *StatRepo {
	return &StatRepo{
		Pool: pool,
	}
}