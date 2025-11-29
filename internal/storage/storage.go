package storage

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Vladimirmoscow84/Warehouse_Control/internal/model"
)

type roleStorage interface {
	CreateRole(ctx context.Context, r *model.Role) (int, error)
	GetRole(ctx context.Context, id int) (*model.Role, error)
	ListRoles(ctx context.Context) ([]*model.Role, error)
}

type userStorage interface {
	CreateUser(ctx context.Context, u *model.User) (int, error)
	GetUser(ctx context.Context, id int) (*model.User, error)
	ListUsers(ctx context.Context) ([]*model.User, error)
}

type itemStorage interface {
	CreateItem(ctx context.Context, i *model.Item, userID int) (int, error)
	GetItem(ctx context.Context, id int) (*model.Item, error)
	ListItems(ctx context.Context) ([]*model.Item, error)
	UpdateItem(ctx context.Context, i *model.Item, userID int) error
	DeleteItem(ctx context.Context, id, userID int) error
}

type itemsHistoryStorage interface {
	ListItemHistory(ctx context.Context, itemID int) ([]*model.ItemHistory, error)
	FilterItemHistory(ctx context.Context, itemID int, userID *int, actionType *string, from, to *time.Time) ([]*model.ItemHistory, error)
}

type Storage struct {
	roles        roleStorage
	users        userStorage
	items        itemStorage
	itemsHistory itemsHistoryStorage
}

func New(r roleStorage, u userStorage, i itemStorage, iH itemsHistoryStorage) (*Storage, error) {
	if r == nil || u == nil || i == nil || iH == nil {
		log.Println("[storage] one of the repositories is nil")
		return nil, errors.New("[storage] one of the repositories is nil")
	}
	return &Storage{
		roles:        r,
		users:        u,
		items:        i,
		itemsHistory: iH,
	}, nil
}

func (s *Storage) CreateRole(ctx context.Context, r *model.Role) (int, error) {
	return s.roles.CreateRole(ctx, r)
}

func (s *Storage) GetRole(ctx context.Context, id int) (*model.Role, error) {
	return s.roles.GetRole(ctx, id)
}

func (s *Storage) ListRoles(ctx context.Context) ([]*model.Role, error) {
	return s.roles.ListRoles(ctx)
}

func (s *Storage) CreateUser(ctx context.Context, u *model.User) (int, error) {
	return s.users.CreateUser(ctx, u)
}

func (s *Storage) GetUser(ctx context.Context, id int) (*model.User, error) {
	return s.users.GetUser(ctx, id)
}

func (s *Storage) ListUsers(ctx context.Context) ([]*model.User, error) {
	return s.users.ListUsers(ctx)
}

func (s *Storage) CreateItem(ctx context.Context, i *model.Item, userID int) (int, error) {
	return s.items.CreateItem(ctx, i, userID)
}

func (s *Storage) GetItem(ctx context.Context, id int) (*model.Item, error) {
	return s.items.GetItem(ctx, id)
}

func (s *Storage) ListItems(ctx context.Context) ([]*model.Item, error) {
	return s.items.ListItems(ctx)
}

func (s *Storage) UpdateItem(ctx context.Context, i *model.Item, userID int) error {
	return s.items.UpdateItem(ctx, i, userID)
}

func (s *Storage) DeleteItem(ctx context.Context, id, userID int) error {
	return s.items.DeleteItem(ctx, id, userID)
}

func (s *Storage) ListItemHistory(ctx context.Context, itemID int) ([]*model.ItemHistory, error) {
	return s.itemsHistory.ListItemHistory(ctx, itemID)
}

func (s *Storage) FilterItemHistory(ctx context.Context, itemID int, userID *int, actionType *string, from, to *time.Time) ([]*model.ItemHistory, error) {
	return s.itemsHistory.FilterItemHistory(ctx, itemID, userID, actionType, from, to)
}
