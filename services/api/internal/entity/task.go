package entity

import (
	"github.com/rafaelbreno/go-api-template/api/internal/states"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `db:"title" json:"title"`             // Title, required
	Description string `db:"description" json:"description"` // Description, not required

	ListID uint `db:"list_id" json:"list_id"` // Foreign Key

	// 0 - TaskIncomplete
	// 1 - TaskComplete
	// 2 - TaskCancelled
	Status states.TaskStatus `gorm:"default:0" db:"status" json:"status"`
}
