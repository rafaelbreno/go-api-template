# API

### Examples
- [iDevoid/stygis](https://github.com/iDevoid/stygis)
    - There's mocks
- [err0r500/go-realworld-clean](https://github.com/err0r500/go-realworld-clean)
    - Here, inside `infra/` there's a nice configuration for Logging and Gin(router)
- [ThreeDotsLabs/wild-workouts-go-ddd-example](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example)
    - Here we've a repository organized with multiple services
- [More](https://github.com/topics/hexagonal-architecture?l=go)


### Project Structure

#### `main.go`
- Entrypoint, this will start everything

#### `cmd/`
- This folder will store packages related to our app infrastructure
    - Server
    - Logger
    - DB Connections
    - [...]
- `http/`
    - Here stores our HTTP configuration (`*gin.Engine`)

#### `routes/`
- Define the API's routes and handlers for:
    - > 404 - Not Found
    - > 405 - Method not Allowed

#### `internal/`
- Internal packages related to our data
    - Entities
    - Queries
- `entity/`
    - Represents our table
- `handler/`
    - This our apps handler
- `query/`
    - Store `constants` of _raw SQL_ queries (you can implement an ORM)
- `states/`
    - Store statuses as `constants` of entities
- `dto/`
    - Data transfer objects

### API Structure

#### List
- POST `/list/create`
    - Creates a new List
    - Payload:
        - ```json
            {
                "title": "List 2",
                "description": "This is the description"
            }
          ```
    - Response 200:
        - ```json
            {
              "data": {
                  "ID": 2,
                  "CreatedAt": "2021-05-24T13:19:44.58560553-03:00",
                  "UpdatedAt": "2021-05-24T13:19:44.58560553-03:00",
                  "DeletedAt": null,
                  "title": "List 2",
                  "description": "This is the description",
                  "status": 0
              }
            }
          ```
- GET `/list`
    - Retrieve __all__ Lists
    - Payload:
        - No need
    - Response 200:
        - ```json
            {
              "data": [
                {
                    "ID": 1,
                    "CreatedAt": "2021-05-24T13:19:44.585605-03:00",
                    "UpdatedAt": "2021-05-24T13:20:02.040603-03:00",
                    "DeletedAt": null,
                    "title": "Task 1",
                    "description": "This is the description",
                    "status": 1
                },
                {
                    "ID": 2,
                    "CreatedAt": "2021-05-24T13:19:44.585605-03:00",
                    "UpdatedAt": "2021-05-24T13:20:02.040603-03:00",
                    "DeletedAt": null,
                    "title": "Task 2",
                    "description": "This is the description",
                    "status": 2
                }
              ]
            }
          ```

- GET `/list/:id`
    - Retrieve list by ID
    - Payload:
        - No need
    - Response 200:
        - ```json
            {
                "data": {
                    "ID": 1,
                    "CreatedAt": "2021-05-24T13:19:44.585605-03:00",
                    "UpdatedAt": "2021-05-24T13:20:02.040603-03:00",
                    "DeletedAt": null,
                    "title": "Task 1",
                    "description": "This is the description",
                    "status": 1
                }
            }
          ```

- PUT/PATCH `/list/:id`
    - Update list by ID given new fields
    - Payload:
        - ```json
            {
                "data": {
                    "title": "Task 1 Updated",
                    "description": "This is the description",
                    "status": 1
                }
            }
          ```
    - Response 200:
        - ```json
            {
                "data": {
                    "ID": 1,
                    "CreatedAt": "2021-05-24T13:19:44.585605-03:00",
                    "UpdatedAt": "2021-05-24T13:20:02.040603-03:00",
                    "DeletedAt": null,
                    "title": "Task 1 Updated",
                    "description": "This is the description",
                    "status": 1
                }
            }
          ```

- DELETE `/list/:id`
    - Delete list by ID
    - Payload:
        - No need
    - Response 200:
        - ```json
              {
                  "data": {
                      "ID": 0,
                      "CreatedAt": "0001-01-01T00:00:00Z",
                      "UpdatedAt": "0001-01-01T00:00:00Z",
                      "DeletedAt": "2021-05-24T13:20:24.222376348-03:00",
                      "title": "",
                      "description": "",
                      "status": 0
                  }
              }
          ```
#### Task
- POST `/task/create`
    - Creates a new List
    - Payload:
        - ```json
            {
                "title": "Task 2",
                "description": "This is the description",
                "list_id": 1
            }
          ```
    - Response 200:
        - ```json
            {
              "data": {
                  "ID": 2,
                  "CreatedAt": "2021-05-24T13:19:44.58560553-03:00",
                  "UpdatedAt": "2021-05-24T13:19:44.58560553-03:00",
                  "DeletedAt": null,
                  "title": "Task 2",
                  "description": "This is the description",
                  "status": 0,
                  "list_id": 1
              }
            }
          ```
- GET `/task`
    - Retrieve tasks from list
    - Payload:
        - ```json
            {
                "list_id": 1
            }
          ```
    - Response 200:
        - ```json
            {
              "data": [
                {
                    "ID": 1,
                    "CreatedAt": "2021-05-24T13:19:44.585605-03:00",
                    "UpdatedAt": "2021-05-24T13:20:02.040603-03:00",
                    "DeletedAt": null,
                    "title": "Task 1",
                    "description": "This is the description",
                    "status": 1
                },
                {
                    "ID": 2,
                    "CreatedAt": "2021-05-24T13:19:44.585605-03:00",
                    "UpdatedAt": "2021-05-24T13:20:02.040603-03:00",
                    "DeletedAt": null,
                    "title": "Task 2",
                    "description": "This is the description",
                    "status": 2
                }
              ]
            }
          ```

- GET `/task/:id`
    - Retrieve list by ID
    - Payload:
        - No need
    - Response 200:
        - ```json
            {
                "data": {
                    "ID": 1,
                    "CreatedAt": "2021-05-24T13:19:44.585605-03:00",
                    "UpdatedAt": "2021-05-24T13:20:02.040603-03:00",
                    "DeletedAt": null,
                    "title": "Task 1",
                    "description": "This is the description",
                    "status": 1
                }
            }
          ```

- PUT/PATCH `/task/:id`
    - Update task by ID given new fields
    - Payload:
        - ```json
            {
                "data": {
                    "title": "Task 1 Updated",
                    "description": "This is the description",
                    "status": 1
                }
            }
          ```
    - Response 200:
        - ```json
            {
                "data": {
                    "ID": 1,
                    "CreatedAt": "2021-05-24T13:19:44.585605-03:00",
                    "UpdatedAt": "2021-05-24T13:20:02.040603-03:00",
                    "DeletedAt": null,
                    "title": "Task 1 Updated",
                    "description": "This is the description",
                    "status": 1
                }
            }
          ```

- DELETE `/task/:id`
    - Delete task by ID
    - Payload:
        - No need
    - Response 200:
        - ```json
              {
                  "data": {
                      "ID": 0,
                      "CreatedAt": "0001-01-01T00:00:00Z",
                      "UpdatedAt": "0001-01-01T00:00:00Z",
                      "DeletedAt": "2021-05-24T13:20:24.222376348-03:00",
                      "title": "",
                      "description": "",
                      "status": 0
                  }
              }
          ```
