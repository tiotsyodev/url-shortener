package url_repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

func (r *URLRepository) GetUrl(ctx context.Context, alias string) (core_domain.UrlDomain, error) {
	qctx, cancel := context.WithTimeout(ctx, r.Pool.TimeoutOperation)
	defer cancel()

	query := `
		SELECT id, alias, url FROM url WHERE alias= $1
	`
	var urlModel UrlModel
	if err := r.Pool.Pool.QueryRow(qctx, query, alias).Scan(
		&urlModel.Id, 
		&urlModel.Alias, 
		&urlModel.Url,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return core_domain.UrlDomain{}, fmt.Errorf("get url timeout: %w", core_domain.ErrTimeout)
		}

		if errors.Is(err, pgx.ErrNoRows) {
            return core_domain.UrlDomain{}, fmt.Errorf("unknown alias: %w", core_domain.ErrNotFound)
		}

		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			switch pgErr.Code {
            	case "23502": 
                	return core_domain.UrlDomain{}, fmt.Errorf("get url invalid input: %w", core_domain.ErrInvalidInput)
			}
		}

        return core_domain.UrlDomain{}, fmt.Errorf("get url: %w", core_domain.ErrDatabase)
	}

	urlDomain := core_domain.NewUrl(urlModel.Id, urlModel.Alias, urlModel.Url)

	return urlDomain, nil

}