package service

import (
	"context"
	"errors"

	"github.com/Vladimirmoscow84/Warehouse_Control/internal/auth"
	"github.com/Vladimirmoscow84/Warehouse_Control/internal/model"
)

func (s *Service) CreateUser(ctx context.Context, u *model.User) (int, error) {
	if u == nil {
		return 0, errors.New("[service-user] user is nil")
	}
	if u.Username == "" {
		return 0, errors.New("[service-user] username is empty")
	}
	if u.PasswordHash == "" {
		return 0, errors.New("[service-user] password_hash is empty")
	}
	if u.RoleID <= 0 {
		return 0, errors.New("[service-user] invalid role_id")
	}

	// Проверка валидности роли
	role, err := s.roles.GetRole(ctx, u.RoleID)
	if err != nil {
		return 0, err
	}
	if role == nil {
		return 0, errors.New("[service-user] invalid role_id")
	}

	id, err := s.users.CreateUser(ctx, u)
	if err != nil {
		return 0, err
	}
	u.ID = id
	return id, nil

}

func (s *Service) GetUser(ctx context.Context, id int) (*model.User, error) {
	if id <= 0 {
		return nil, errors.New("[service-user] invalid user id")
	}
	return s.users.GetUser(ctx, id)
}

func (s *Service) ListUsers(ctx context.Context) ([]*model.User, error) {
	return s.users.ListUsers(ctx)
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	if username == "" {
		return nil, errors.New("[service] username is empty")
	}
	return s.users.GetUserByUsername(ctx, username)
}

func (s *Service) AuthenticateUser(ctx context.Context, username, password string) (int, string, error) {
	if username == "" || password == "" {
		return 0, "", errors.New("[service] empty credentials")
	}

	u, err := s.users.GetUserByUsername(ctx, username)
	if err != nil {
		return 0, "", err
	}
	if u == nil {
		return 0, "", errors.New("[service] invalid credentials")
	}

	// Проверка пароля (auth.CheckPassword)
	if !auth.CheckPassword(password, u.PasswordHash) {
		return 0, "", errors.New("[service] invalid credentials")
	}

	// получение роли и пароля
	role, err := s.roles.GetRole(ctx, u.RoleID)
	if err != nil {
		return 0, "", err
	}
	if role == nil {
		return 0, "", errors.New("[service] role not found")
	}

	return u.ID, role.RoleName, nil
}
