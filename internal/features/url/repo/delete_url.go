package url_repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

func (r *URLRepository) DeleteUrl(ctx context.Context, id int) (core_domain.UrlDomain, error) {
	qctx, cancel := context.WithTimeout(ctx, r.Pool.TimeoutOperation)
	defer cancel()

	query := `
		DELETE FROM url WHERE id=$1 RETURNING id, alias, url
	`

	var urlModel UrlModel
	err := r.Pool.Pool.QueryRow(qctx, query, id).Scan(&urlModel.Id, &urlModel.Alias, &urlModel.Url)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return core_domain.UrlDomain{}, fmt.Errorf("delete url: %w", core_domain.ErrTimeout)
		}	

		if errors.Is(err, pgx.ErrNoRows) {
			return core_domain.UrlDomain{}, fmt.Errorf("delete url: %w", core_domain.ErrNotFound)
		}

		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			switch pgErr.Code {
            	case "23502": 
                	return core_domain.UrlDomain{}, fmt.Errorf("get url invalid input: %w", core_domain.ErrInvalidInput)
			}
		}

        return core_domain.UrlDomain{}, fmt.Errorf("delete url: %w", core_domain.ErrDatabase)
	}

	urlDomain := core_domain.NewUrl(urlModel.Id, urlModel.Alias, urlModel.Url)

	return urlDomain, nil

}