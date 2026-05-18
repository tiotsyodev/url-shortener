package url_repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

func (r *URLRepository) CreateURL(ctx context.Context, d core_domain.UrlDomain) (core_domain.UrlDomain, error) {
	qctx, cancel := context.WithTimeout(ctx, r.Pool.TimeoutOperation)
	defer cancel()

	query := `INSERT INTO url(alias, url) values ($1, $2) returning id, alias, url`

	var urlModel UrlModel
	if err := r.Pool.QueryRow(qctx, query, d.Alias, d.Url).Scan(&urlModel.Id, &urlModel.Alias, &urlModel.Url); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return core_domain.UrlDomain{}, fmt.Errorf("create url: %w", core_domain.ErrTimeout)
		}

		var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) {
            switch pgErr.Code {
            case "23505":
                return core_domain.UrlDomain{}, fmt.Errorf("create url: %w", core_domain.ErrAlreadyExists)
            case "23502": 
                return core_domain.UrlDomain{}, fmt.Errorf("create url: %w", core_domain.ErrInvalidInput)
            }
        }

        return core_domain.UrlDomain{}, fmt.Errorf("create url: %w", core_domain.ErrDatabase)
	}

	urlDomain := core_domain.NewUrl(urlModel.Id, urlModel.Alias, urlModel.Url)
	return urlDomain, nil
}