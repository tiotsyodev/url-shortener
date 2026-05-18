package url_repo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

func (r *URLRepository) UpdateUrl(ctx context.Context, dom core_domain.UpdateUrlDomain) (core_domain.UrlDomain, error) {
	qctx, cancel := context.WithTimeout(ctx, r.Pool.TimeoutOperation)
	defer cancel()

	i := 1
	args := []any{}
	clauses := []string{}

	if dom.Alias != nil {
		clauses = append(clauses, fmt.Sprintf("alias=$%d", i))
		args = append(args, *dom.Alias)
		i++
	}

	if dom.Url != nil {
		clauses = append(clauses, fmt.Sprintf("url=$%d", i))
		args = append(args, *dom.Url)
		i++
	}


	if len(clauses) == 0 {
		return core_domain.UrlDomain{}, fmt.Errorf("update url: %w", core_domain.ErrInvalidInput)
	}

	args = append(args, dom.Id)

	query := fmt.Sprintf("UPDATE url SET %s WHERE id=$%d RETURNING id, alias, url", strings.Join(clauses, ", "), i)

	var urlModel UrlModel
	err := r.Pool.Pool.QueryRow(qctx, query, args...).Scan(&urlModel.Id, &urlModel.Alias, &urlModel.Url)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
            return core_domain.UrlDomain{}, fmt.Errorf("update url: %w", core_domain.ErrNotFound)
        }
        if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
            return core_domain.UrlDomain{}, fmt.Errorf("update url: %w", core_domain.ErrTimeout)
        }
        var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) && pgErr.Code == "23505" {
            return core_domain.UrlDomain{}, fmt.Errorf("update url: %w", core_domain.ErrAlreadyExists)
        }
        return core_domain.UrlDomain{}, fmt.Errorf("update url: %w", core_domain.ErrDatabase)
	}

	return core_domain.NewUrl(urlModel.Id, urlModel.Alias, urlModel.Url), nil

}