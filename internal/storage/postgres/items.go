package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Vladimirmoscow84/Warehouse_Control/internal/model"
)

func (p *Postgres) CreateItem(ctx context.Context, i *model.Item, userID int) (int, error) {
	if i == nil {
		log.Println("[postgres-item] item is nil")
		return 0, errors.New("item is nil")
	}
	tx, err := p.DB.BeginTxx(ctx, nil)
	if err != nil {
		log.Println("[postgres-item] error of tx")
		return 0, err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "SELECT set_config('app.current_user',$1, true)", fmt.Sprint(userID))
	if err != nil {
		log.Println("[postgres-item] error set userID")
		return 0, err
	}
	var id int
	query := `
		INSERT INTO items
			(sku, title, quantity, price)
		VALUES
			($1, $2, $3, $4)
		RETURNING id
	`
	row := tx.QueryRowContext(ctx, query, i.SKU, i.Title, i.Quantity, i.Price)
	err = row.Scan(&id)
	if err != nil {
		log.Printf("[postgres-item] error adding item in DB: %v", err)
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		log.Println("[postgres-item] error of tx")
		return 0, err
	}

	i.ID = id
	return id, nil
}

func (p *Postgres) GetItem(ctx context.Context, id int) (*model.Item, error) {
	var i model.Item
	query := `
		SELECT
			id, sku, title, quantity, price, version, created_at, updated_at
		FROM items
		WHERE id=$1
	`
	err := p.DB.GetContext(ctx, &i, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		log.Printf("[postgres-item] error getting item in DB by id=%d: %v", id, err)
		return nil, err
	}
	return &i, nil
}

func (p *Postgres) ListItems(ctx context.Context) ([]*model.Item, error) {
	var items []*model.Item
	query := `
		SELECT 
			id, sku, title, quantity, price, version, created_at, updated_at
		FROM items
		ORDER BY id
	`
	err := p.DB.SelectContext(ctx, &items, query)
	if err != nil {
		log.Println("[postgres-item] error listing items")
		return nil, err
	}
	return items, nil
}

func (p *Postgres) UpdateItem(ctx context.Context, i *model.Item, userID int) error {
	if i == nil {
		log.Println("[postgres-item] item is nil")
		return errors.New("item is nil")
	}
	tx, err := p.DB.BeginTxx(ctx, nil)
	if err != nil {
		log.Println("[postgres-item] error of tx")
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "SELECT set_config('app.current_user', $1, true)", fmt.Sprint(userID))
	if err != nil {
		log.Println("[postgres-item] error set userID")
		return err
	}

	query := `
		UPDATE items
		SET sku=$1, title=$2, quantity=$3, price=$4
		WHERE id=$5
	`
	_, err = tx.ExecContext(ctx, query, i.SKU, i.Title, i.Quantity, i.Price, i.ID)
	if err != nil {
		log.Println("[postgres-item] error of UPDATE DB")
		return err
	}
	return tx.Commit()
}

func (p *Postgres) DeleteItem(ctx context.Context, id, userID int) error {
	tx, err := p.DB.BeginTxx(ctx, nil)
	if err != nil {
		log.Println("[postgres-item] error of tx")
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "SELECT set_config('app.current_user', $1, true)", fmt.Sprint(userID))
	if err != nil {
		log.Println("[postgres-item] error of DELETE DB")
		return err
	}

	query := `
		DELETE FROM items
		WHERE id=$1
	`
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("[postgres-item] error of DELETE DB")
		return err
	}
	return tx.Commit()
}
