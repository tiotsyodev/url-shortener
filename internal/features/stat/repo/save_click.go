package stats_repo

import (
	"context"
	"errors"
	"fmt"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

func (r *StatRepo) SaveClick(ctx context.Context, dom core_domain.StatDomain) error {
	qctx, cancel := context.WithTimeout(ctx, r.Pool.TimeoutOperation)
	defer cancel()

	query := `
		INSERT INTO click (url_id, ip, user_agent, device) VALUES($1, $2, $3, $4)
	`

	_, err := r.Pool.Pool.Exec(qctx, query, dom.UrlId, dom.Ip, dom.UserAgent, dom.Device) 
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
			return fmt.Errorf("save click: %w", core_domain.ErrTimeout)
		}
		return fmt.Errorf("save click: %w", core_domain.ErrDatabase)
	}

	return nil

}