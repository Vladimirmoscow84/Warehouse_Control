package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/Vladimirmoscow84/Warehouse_Control/internal/model"
)

func (p *Postgres) CreateRole(ctx context.Context, r *model.Role) (int, error) {
	if r == nil {
		log.Println("[postgres-role] role is nil")
		return 0, errors.New("role is nil")
	}
	var id int
	query := `
		INSERT INTO roles 
			(role_name,description)
		VALUES
			($1, $2)
		RETURNING id
	`
	row := p.DB.QueryRowContext(ctx, query, r.RoleName, r.Description)
	err := row.Scan(&id)
	if err != nil {
		log.Printf("[postgres-role] error adding role in DB: %v", err)
		return 0, err
	}
	r.ID = id
	return id, nil
}

func (p *Postgres) GetRole(ctx context.Context, id int) (*model.Role, error) {
	var r model.Role
	query := `
		SELECT 
			id, role_name, description, created_at, updated_at
		FROM roles
		WHERE id=$1
	`
	err := p.DB.GetContext(ctx, &r, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		log.Printf("[postgres-role] error getting role in DB by id=%d: %v", id, err)
		return nil, err
	}
	return &r, nil
}

func (p *Postgres) ListRoles(ctx context.Context) ([]*model.Role, error) {
	var roles []*model.Role
	query := `
		SELECT
			id, role_name, description, created_at, updated_at
		FROM roles
		ORDER by id
	`
	err := p.DB.SelectContext(ctx, &roles, query)
	if err != nil {
		log.Println("[postgres-role] error listing roles")
		return nil, err
	}
	return roles, nil
}
