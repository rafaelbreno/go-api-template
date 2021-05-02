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
- `query/`
    - Store `constants` of _raw SQL_ queries (you can implement an ORM)
- `dto/`
    - Data transfer objects
