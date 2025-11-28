package model

import (
	"encoding/json"
	"time"
)

type Role struct {
	ID          int       `json:"-" db:"id"`
	RoleName    string    `json:"role_name" db:"role_name"`
	Description string    `json:"description,omitempty" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type User struct {
	ID           int       `json:"-" db:"id"`
	Username     string    `json:"username" db:"username"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Email        string    `json:"email,omitempty" db:"email"`
	RoleID       int       `json:"role_id" db:"role_id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type Item struct {
	ID        int       `json:"id" db:"id"`
	SKU       string    `json:"sku" db:"sku"`
	Title     string    `json:"title" db:"title"`
	Quantity  int       `json:"quantity" db:"quantity"`
	Price     float64   `json:"price" db:"price"`
	Version   int       `json:"version" db:"version"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ItemHistory struct {
	ID         int             `json:"id" db:"id"`
	ItemID     int             `json:"item_id" db:"item_id"`
	ActionType string          `json:"action_type" db:"action_type"`
	OldValue   json.RawMessage `json:"old_value" db:"old_value"`
	NewValue   json.RawMessage `json:"new_value" db:"new_value"`
	ChangedBy  int             `json:"changed_by" db:"changed_by"`
	ChangedAt  time.Time       `json:"changed_at" db:"changed_at"`
}
