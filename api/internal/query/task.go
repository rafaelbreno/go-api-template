package query

const (
	TaskInsert = `
	INSERT INTO tasks (
		title,
		description
	) VALUES (
		:title,
		:description
	) RETURNING
		id,
		title,
		description,
		status,
		created_on,
		updated_on;
	`
)
