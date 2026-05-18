package url_repo

import core_repo "github.com/tiotsyodev/url-shortener.git/internal/core/repo"

type URLRepository struct {
	Pool core_repo.ConnectionPool
}

func NewUrlRepository(pool core_repo.ConnectionPool) *URLRepository {
	return &URLRepository{
		Pool: pool,
	}
}


