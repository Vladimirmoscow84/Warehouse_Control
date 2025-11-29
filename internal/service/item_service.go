package service

import (
	"context"
	"errors"

	"github.com/Vladimirmoscow84/Warehouse_Control/internal/model"
)

func (s *Service) CreateItem(ctx context.Context, i *model.Item, userID int) (int, error) {
	if i == nil {
		return 0, errors.New("[service-item] item is nil")
	}
	if i.SKU == "" {
		return 0, errors.New("[service-item] SKU is empty")
	}
	if i.Title == "" {
		return 0, errors.New("[service-item] title is empty")
	}
	if userID <= 0 {
		return 0, errors.New("[service-item] invalid userID")
	}

	//Получение роли пользователя для проверки права
	user, err := s.users.GetUser(ctx, userID)
	if err != nil {
		return 0, err
	}
	if user == nil {
		return 0, errors.New("[service-item] user not found")
	}
	role, err := s.roles.GetRole(ctx, user.RoleID)
	if err != nil {
		return 0, err
	}
	if role == nil {
		return 0, errors.New("[service-item] role not found")
	}

	//Проверка права
	if role.RoleName != RoleAdmin && role.RoleName != RoleManager {
		return 0, errors.New("[service-item] no rules to create item")
	}

	//Создание Item
	id, err := s.items.CreateItem(ctx, i, userID)
	if err != nil {
		return 0, err
	}
	i.ID = id
	return id, nil
}

func (s *Service) GetItem(ctx context.Context, id int) (*model.Item, error) {
	if id <= 0 {
		return nil, errors.New("[service-item] invalid item id")
	}
	return s.items.GetItem(ctx, id)
}

func (s *Service) ListItems(ctx context.Context) ([]*model.Item, error) {
	return s.items.ListItems(ctx)
}

func (s *Service) UpdateItem(ctx context.Context, i *model.Item, userID int) error {
	if i == nil {
		return errors.New("[service-item] item is nil")
	}
	if i.SKU == "" {
		return errors.New("[service-item] SKU is empty")
	}
	if i.Title == "" {
		return errors.New("[service-item] title is empty")
	}
	if userID <= 0 {
		return errors.New("[service-item] invalid userID")
	}

	//Получение роли пользователя для проверки права
	user, err := s.users.GetUser(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("[service-item] user not found")
	}
	role, err := s.roles.GetRole(ctx, user.RoleID)
	if err != nil {
		return err
	}
	if role == nil {
		return errors.New("[service-item] role not found")
	}

	//Проверка права
	if role.RoleName != RoleAdmin && role.RoleName != RoleManager {
		return errors.New("[service-item] no rules to update item")
	}
	return s.items.UpdateItem(ctx, i, userID)
}

func (s *Service) DeleteItem(ctx context.Context, id, userID int) error {
	//Проверка роли пользователя
	user, err := s.users.GetUser(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("[service-item] user not found")
	}
	role, err := s.roles.GetRole(ctx, user.RoleID)
	if err != nil {
		return err
	}
	if role == nil {
		return errors.New("[service-item] role not found")
	}

	//Проверка права
	if role.RoleName != RoleAdmin {
		return errors.New("[service-item] no rules to delete item")
	}
	return s.items.DeleteItem(ctx, id, userID)

}
