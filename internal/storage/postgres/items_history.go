package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Vladimirmoscow84/Warehouse_Control/internal/model"
)

func (p *Postgres) ListItemHistory(ctx context.Context, itemID int) ([]*model.ItemHistory, error) {
	var histories []*model.ItemHistory

	query := `
		SELECT 
			id, item_id, action_type, old_value, new_value, changed_by, changed_at
		FROM item_history
		WHERE item_id=$1
		ORDER BY changed_at ASC
`

	err := p.DB.SelectContext(ctx, &histories, query, itemID)
	if err != nil {
		log.Println("[postgres-history] error geting list of histories from DB")
		return nil, err
	}

	return histories, nil

}

func (p *Postgres) FilterItemHistory(ctx context.Context, itemID int, userID *int, actionType *string, from, to *time.Time) ([]*model.ItemHistory, error) {
	var histories []*model.ItemHistory

	query := `
		SELECT 
			id, item_id, action_type, old_value, new_value, changed_by, changed_at
		FROM item_history
		WHERE item_id=$1
`
	args := []interface{}{itemID}
	argIdx := 2

	if userID != nil {
		query += fmt.Sprintf(" AND changed_by=$%d", argIdx)
		args = append(args, *userID)
		argIdx++
	}

	if actionType != nil {
		query += fmt.Sprintf(" AND action_type=$%d", argIdx)
		args = append(args, *actionType)
		argIdx++
	}

	if from != nil {
		query += fmt.Sprintf(" AND changed_at >= $%d", argIdx)
		args = append(args, *from)
		argIdx++
	}

	if to != nil {
		query += fmt.Sprintf(" AND changed_at <= $%d", argIdx)
		args = append(args, *to)
		argIdx++
	}

	query += " ORDER BY changed_at ASC"

	err := p.DB.SelectContext(ctx, &histories, query, args...)
	if err != nil {
		log.Println("[postgres-history] error filtering histories in DB")
		return nil, err
	}

	return histories, nil

}
