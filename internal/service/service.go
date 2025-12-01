package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Vladimirmoscow84/Warehouse_Control/internal/model"
)

// роли пользователей
const (
	RoleAdmin   = "admin"
	RoleManager = "manager"
	RoleViewer  = "viewer"
)

type roleService interface {
	CreateRole(ctx context.Context, r *model.Role) (int, error)
	GetRole(ctx context.Context, id int) (*model.Role, error)
	ListRoles(ctx context.Context) ([]*model.Role, error)
}

type userService interface {
	CreateUser(ctx context.Context, u *model.User) (int, error)
	GetUser(ctx context.Context, id int) (*model.User, error)
	ListUsers(ctx context.Context) ([]*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	AuthenticateUser(ctx context.Context, username, password string) (userID int, roleName string, err error)
}

type itemService interface {
	CreateItem(ctx context.Context, i *model.Item, userID int) (int, error)
	GetItem(ctx context.Context, id int) (*model.Item, error)
	ListItems(ctx context.Context) ([]*model.Item, error)
	UpdateItem(ctx context.Context, i *model.Item, userID int) error
	DeleteItem(ctx context.Context, id, userID int) error
}

type itemHistoryService interface {
	ListItemHistory(ctx context.Context, itemID int) ([]*model.ItemHistory, error)
	FilterItemHistory(ctx context.Context, itemID int, userID *int, actionType *string, from, to *time.Time) ([]*model.ItemHistory, error)
}

type Service struct {
	roles        roleService
	users        userService
	items        itemService
	itemsHistory itemHistoryService
}

func New(r roleService, u userService, i itemService, ih itemHistoryService) (*Service, error) {
	if r == nil || u == nil || i == nil || ih == nil {
		log.Println("[service] one of the services is nil")
		return nil, errors.New("[service] one of the services is nil")
	}
	return &Service{
		roles:        r,
		users:        u,
		items:        i,
		itemsHistory: ih,
	}, nil
}
