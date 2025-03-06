package entity

import "time"

type Label struct {
	ID          uint      `json:"id" gorm:"primaryKey;column:id"`
	Name        string    `json:"name" gorm:"not null;column:name"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null;default:now();column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"not null;default:now();column:updated_at"`
}


