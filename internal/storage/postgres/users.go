package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/Vladimirmoscow84/Warehouse_Control/internal/model"
)

func (p *Postgres) CreateUser(ctx context.Context, u *model.User) (int, error) {
	if u == nil {
		log.Println("[postgres-user] user is nil")
		return 0, errors.New("user is nil")
	}
	var id int
	query := `
		INSERT INTO users
			(username, password_hash, email, role_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	row := p.DB.QueryRowContext(ctx, query, u.Username, u.PasswordHash, u.Email, u.RoleID)
	err := row.Scan(&id)
	if err != nil {
		log.Printf("[postgres-user] error add user in DB: %v", err)
		return 0, err
	}
	u.ID = id
	return id, nil
}

func (p *Postgres) GetUser(ctx context.Context, id int) (*model.User, error) {
	var u model.User

	query := `
		SELECT 
			id, username, password_hash, email, role_id, created_at, updated_at
		FROM users
		WHERE id=$1
	`
	err := p.DB.GetContext(ctx, &u, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		log.Printf("[postgres-user] error getting user in DB by id=%d: %v", id, err)
		return nil, err

	}
	return &u, nil
}

func (p *Postgres) ListUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	query := `
		SELECT 
			id, username, password_hash, email, role_id, created_at, updated_at
		FROM users
		ORDER BY id
	`
	err := p.DB.SelectContext(ctx, &users, query)
	if err != nil {
		log.Println("[postgres-user] error listing users")
		return nil, err
	}
	return users, nil
}
