package query

const (
	// Task base Table
	TaskSchema = `
		CREATE TABLE tasks
		(
			id 			SERIAL,
			title   	CHARACTER(64),
			description CHARACTER(256),
			status 		SMALLINT DEFAULT 0,
			created_on 	TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_on 	TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);
	`

	// Task create Query
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
