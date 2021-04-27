package entity

import (
	"time"

	"github.com/rafaelbreno/go-api-template/api/internal/states"
)

type Task struct {
	ID          uint   `db:"id" json:"task_id"`              // ID auto-incremented by DB
	Title       string `db:"title" json:"title"`             // Title, required
	Description string `db:"description" json:"description"` // Description, not required

	// 0 - TaskIncomplete
	// 1 - TaskComplete
	// 2 - TaskCancelled
	Status states.TaskStatus `db:"status" json:"status"`

	// When task were created
	// auto-generated
	CreatedOn time.Time `db:"created_on" json:"created_on"`

	// When task were updated
	// auto-generated
	UpdatedOn time.Time `db:"updated_on" json:"updated_on"`
}
