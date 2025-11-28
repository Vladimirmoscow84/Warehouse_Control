package postgres

import (
	"context"
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
			(user_name, password_hash, email, role_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	err := p.DB.QueryRowContext(ctx, query, u.Username, u.PasswordHash, u.Email, u.RoleID).Scan(&id)
	if err != nil {
		log.Printf("[posgres-user] error add user in DB: %v", err)
		return 0, err
	}
	return id, nil
}
