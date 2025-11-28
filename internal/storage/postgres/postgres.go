package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Vladimirmoscow84/Warehouse_Control/internal/model"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type repoRoles interface {
	CreateRole(ctx context.Context, r *model.Role) (int, error)
	GetRole(ctx context.Context, id int) (*model.Role, error)
	ListRoles(ctx context.Context) ([]*model.Role, error)
}

type repoUsers interface {
	CreateUser(ctx context.Context, u *model.User) (int, error)
	GetUser(ctx context.Context, id int) (*model.User, error)
	ListUsers(ctx context.Context) ([]*model.User, error)
}

type repoItmes interface {
	CreateItem(ctx context.Context, i *model.Item, userID int) (int, error)
	GetItem(ctx context.Context, id int) (*model.Item, error)
	ListItems(ctx context.Context) ([]*model.Item, error)
	UpdateItem(ctx context.Context, i *model.Item, userID int) error
	DeleteItem(ctx context.Context, id, userID int) error
}

type repoItemHistory interface {
	ListItemHistory(ctx context.Context, itemID int) ([]*model.ItemHistory, error)
	FilterItemHistory(ctx context.Context, itemID int, userID *int, actionType *string, from, to *time.Time) ([]*model.ItemHistory, error)
}

type Postgres struct {
	DB *sqlx.DB
}

func New(databaseURI string) (*Postgres, error) {
	db, err := sqlx.Connect("sqlx", databaseURI)
	if err != nil {
		return nil, fmt.Errorf("[postgres] failed to connect to DB: %w", err)
	}
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(30)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("[postgres] ping is failed: %w", err)
	}
	log.Println("[postgres] successfull connect to DB")
	return &Postgres{
		DB: db,
	}, nil
}

func (p *Postgres) Close() error {
	err := p.DB.Close()
	if err != nil {
		log.Printf("[postgres] failed to close DB connection: %v", err)
		return err
	}
	log.Println("[postgres] connection closed")
	return nil
}
