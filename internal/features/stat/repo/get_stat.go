package stats_repo

import (
	"context"
	"fmt"
	"time"

	core_domain "github.com/tiotsyodev/url-shortener.git/internal/core/domain"
)

func (r *StatRepo) GetStats(ctx context.Context, alias string) (core_domain.StatsByAlias, error) {
	qctx, cancel := context.WithTimeout(ctx, r.Pool.TimeoutOperation)
	defer cancel()

	queryByDays := `
        SELECT DATE(created_at) as day, COUNT(*) as clicks
        FROM clicks
        WHERE url_id = (SELECT id FROM url WHERE alias = $1)
        GROUP BY day
        ORDER BY day DESC
    `

	queryByDevices := `
        SELECT device, COUNT(*) as clicks
        FROM clicks
        WHERE url_id = (SELECT id FROM url WHERE alias = $1)
        GROUP BY device
    `

	rows, err := r.Pool.Pool.Query(qctx, queryByDays, alias)
	if err != nil {
		return core_domain.StatsByAlias{}, fmt.Errorf("get stats by days: %w", core_domain.ErrDatabase)
	}
	defer rows.Close()

	byDays := map[string]int{}
	for rows.Next() {
		var day time.Time
		var clicks int
		if err := rows.Scan(&day, &clicks); err != nil {
			return core_domain.StatsByAlias{}, fmt.Errorf("scan stats by days: %w", core_domain.ErrDatabase)
		}
		byDays[day.Format("2006-01-02")] = clicks
	}
	if err := rows.Err(); err != nil {
		return core_domain.StatsByAlias{}, fmt.Errorf("rows error: %w", core_domain.ErrDatabase)
	}

	rows, err = r.Pool.Pool.Query(qctx, queryByDevices, alias)
	if err != nil {
		return core_domain.StatsByAlias{}, fmt.Errorf("get stats by devices: %w", core_domain.ErrDatabase)
	}
	defer rows.Close()

	byDevices := map[string]int{}
	for rows.Next() {
		var device string
		var clicks int
		if err := rows.Scan(&device, &clicks); err != nil {
			return core_domain.StatsByAlias{}, fmt.Errorf("scan stats by devices: %w", core_domain.ErrDatabase)
		}
		byDevices[device] = clicks
	}

	if err := rows.Err(); err != nil {
		return core_domain.StatsByAlias{}, fmt.Errorf("rows error: %w", core_domain.ErrDatabase)
	}	

	return core_domain.StatsByAlias{
		ByDays:    byDays,
		ByDevices: byDevices,
	}, nil
}