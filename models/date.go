package models

import "app/services/time"

// Date ...
type Date struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"` 
	DeletedAt *time.Time `json:"-"`
}
