package states

type TaskStatus int

const (
	TaskIncomplete TaskStatus = iota

	TaskComplete

	TaskCancelled
)
