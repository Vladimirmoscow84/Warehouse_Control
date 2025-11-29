package service

import (
	"context"
	"errors"

	"github.com/Vladimirmoscow84/Warehouse_Control/internal/model"
)

func (s *Service) CreateRole(ctx context.Context, r *model.Role) (int, error) {
	if r == nil {
		return 0, errors.New("[service-role] role is nil")
	}

	id, err := s.roles.CreateRole(ctx, r)
	if err != nil {
		return 0, err
	}
	return id, nil

}

func (s *Service) GetRole(ctx context.Context, id int) (*model.Role, error) {
	if id <= 0 {
		return nil, errors.New("[service-role] invalid role id")
	}
	return s.roles.GetRole(ctx, id)
}

func (s *Service) ListRoles(ctx context.Context) ([]*model.Role, error) {
	return s.roles.ListRoles(ctx)
}
