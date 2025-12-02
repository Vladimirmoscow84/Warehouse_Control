package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Vladimirmoscow84/Warehouse_Control/internal/model"
)

func (s *Service) ListItemHistory(ctx context.Context, itemID int) ([]*model.ItemHistory, error) {
	if itemID <= 0 {
		return nil, errors.New("[service-history] invalid itemID")
	}
	_, err := s.items.GetItem(ctx, itemID)
	if err != nil {
		return nil, fmt.Errorf("[service-history] item not found: %w", err)
	}
	return s.itemsHistory.ListItemHistory(ctx, itemID)
}

func (s *Service) FilterItemHistory(ctx context.Context, itemID int, userID *int, actionType *string, from, to *time.Time) ([]*model.ItemHistory, error) {
	if itemID <= 0 {
		return nil, errors.New("[service-history] invalid itemID")
	}
	_, err := s.items.GetItem(ctx, itemID)
	if err != nil {
		return nil, fmt.Errorf("[service-history] item not found: %w", err)
	}

	return s.itemsHistory.FilterItemHistory(ctx, itemID, userID, actionType, from, to)
}
