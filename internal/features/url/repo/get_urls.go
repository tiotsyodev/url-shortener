package url_repo

import (
	"context"
	"errors"
	"fmt"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

func (r *URLRepository) GetUrls(ctx context.Context, limit int, offset int) ([]core_domain.UrlDomain, error) {
	qctx, cancel := context.WithTimeout(ctx, r.Pool.TimeoutOperation)
	defer cancel()

	query := `
		select id, alias, url from url LIMIT $1 OFFSET 1
	`

	rows, err := r.Pool.Pool.Query(qctx, query)
	
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
            return nil, fmt.Errorf("get urls: %w", core_domain.ErrTimeout)
        }
        return nil, fmt.Errorf("get urls: %w", core_domain.ErrDatabase)
	}
	defer rows.Close()
	var urls []core_domain.UrlDomain

	for rows.Next() {
		var url core_domain.UrlDomain

		if err := rows.Scan(&url.Id, &url.Alias, &url.Url); err != nil {
			return nil, fmt.Errorf("get urls: %w", core_domain.ErrDatabase)
		}

		urls = append(urls, url)
	}

	return urls, nil
}