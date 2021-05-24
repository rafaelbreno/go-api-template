package states

type ListStatus int

const (
	ListNotStarted ListStatus = iota

	ListInProgress

	ListCompleted

	ListCancelled
)
