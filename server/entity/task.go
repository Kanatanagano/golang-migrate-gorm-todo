package entity

import "time"

type Priority string

const (
	Low Priority = "Low"
	Middle Priority = "Middle"
	High Priority = "High"
)

type Status string
const (
	Pending Status = "Pending"
	InProgress Status = "In_Progress"
	Completed Status = "Completed"
)

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey;column:id"`
	Title       string    `json:"title" gorm:"not null;column:title"`
	Description string    `json:"description" gorm:"not null;column:description"`
	Priority    Priority  `json:"priority" gorm:"not null;column:priority;type:priority"`
	Status      Status    `json:"status" gorm:"not null;column:status;type:status"`
	Labels      []Label   `json:"labels" gorm:"many2many:task_labels;"`
	StartedAt   time.Time `json:"started_at" gorm:"not null;column:started_at"`
	FinishedAt  time.Time `json:"finished_at" gorm:"not null;column:finished_at"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null;default:now();column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"not null;default:now();column:updated_at"`
}
