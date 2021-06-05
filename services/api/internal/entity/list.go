package entity

import (
	"github.com/rafaelbreno/go-api-template/api/internal/states"
	"gorm.io/gorm"
)

type List struct {
	gorm.Model

	UserID uint `db:"user_id" json:"user_id"` // User foreign key

	Title       string `db:"title" json:"title"`             // Title, required
	Description string `db:"description" json:"description"` // Description, not required

	Tasks []Task `json:"tasks"` // Tasks assigned to this list

	// 0 - ListNotStarted ListStatus = iota
	// 1 - ListInProgress
	// 2 - ListCompleted
	// 3 - ListCancelled
	Status states.ListStatus `gorm:"default:0" db:"status" json:"status"`
}
