package entity

import (
	"github.com/rafaelbreno/go-api-template/api/internal/states"
	"gorm.io/gorm"
)

type List struct {
	gorm.Model

	Title       string `db:"title" json:"title"`             // Title, required
	Description string `db:"description" json:"description"` // Description, not required

	// 0 - ListNotStarted ListStatus = iota
	// 1 - ListInProgress
	// 2 - ListCompleted
	// 3 - ListCancelled
	Status states.ListStatus `gorm:"default:0" db:"status" json:"status"`
}
