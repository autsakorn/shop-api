# shop-api
An example web API application built with Beego and Postgres.
Most of this application built by the standard library.

# Get Started
- `docker-compose build`
- `docker-compose up`

# Migrate
After running the docker, you need to migrate the database.
You can migrate by executing to docker and run migrate by command 2 steps.
1. `docker-compose exec api-shop sh`
2. `bee migrate -driver=$driver -conn=$sqlconn`

If you need to roll back the database
- `bee rollback migrate -driver=$driver -conn=$sqlconn`

# Swagger
Open your browser and go to http://localhost:8080

# Test
You can run unit test by the command
- `go test ./... -cover`
  
If you want to see test coverage detail 3 step

1. `go test -coverprofile=coverage/cover.out ./... -cover`
2. `go tool cover -html=coverage/cover.out -o coverage/index.html`
3. Open your browser and go to http://localhost:8081/coverage

# Description
This is an example of the implementation of web API in Golang, by separating the application into layers. you will create a testable system. When any of the external parts of the system become obsolete, like the database, you can swap, your business rules are not bound to the database

This project has 4 layers:
- Controllers layer
- Services layer
- Storage layer
- Models layer

## Controllers layer
The controller layer is the user interface. This is the software user sees and interacts with. They enter the needed information. This layer also acts as a go-between for the data layer and the user, passing on the user’s different actions to the logic layer.

## Services layer
The service layer is where all the “thinking” happens, and it knows what is allowed by your application and what is possible, and it makes other decisions. This logic layer is also the one that writes and reads data into the data layer.

## Storage layer
The storage will store any database handler. Querying, Inserting into any database will store here. This layer will act for CRUD to the database only. No business rule happens here. Only plain function to Database.

## Models layer
The model layer is the same as entities, This layer will store entities Object’s Struct