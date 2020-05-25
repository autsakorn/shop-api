# shop-api
An example web API application built with Beego and Postgres.
Most of this application built by the standard library.

## Get Started
- `docker-compose build`
- `docker-compose up`

## Migrate
After running the docker, you need to migrate the database.
You can migrate by executing to docker and run migrate by command 2 steps.

1. `docker-compose exec api-shop sh`
2. `bee migrate -driver=$DRIVER -conn=$SQLCONN`

If you need to roll back the database
- `bee rollback migrate -driver=$DRIVER -conn=$SQLCONN`

## Swagger
- Open your browser and go to http://localhost:8080 or http://localhost:8081/swagger
- swagger/swagger.json and swagger/swagger.yml are auto-generate but you can manual generate by 

1. `docker-compose exec api-shop sh`
2. `bee generate docs`


## Test
You can run unit test by the command
- `go test ./services ./models ./utils -cover`
  
If you want to see test coverage detail 3 step

1. `go test -coverprofile=coverage/cover.out ./services ./models ./utils -cover`
2. `go tool cover -html=coverage/cover.out -o coverage/index.html`
3. Open your browser and go to http://localhost:8081/coverage

## Description layers
This is an example of the implementation of web API in Golang, by separating the application into layers. you will create a testable system. When any of the external parts of the system become obsolete, like the database, you can swap, your business rules are not bound to the database

This project has 4 layers:
- Controllers layer
- Services layer
- Storage layer
- Models layer

### Controllers layer
The controller layer is the user interface. This is the software user sees and interacts with. They enter the needed information. This layer passing on the user’s different actions to the service layer.

### Services layer
The service layer is where all the thinking happens, this layer contains application-specific business rules and makes decisions. This service layer is also the one that calls method writes and reads data into the storage layer.

### Storage layer
The storage will store any database handler. Querying, Inserting into any database will store here. This layer will act for CRUD to the database only. No business rule happens here. Only plain function to Database.

### Models layer
The model layer is the same as entities, This layer will store entities Object’s Struct

### Communications Between Layer

``` text
     +------------+         +------------+           2. GetByEmail       +---------------+
     |            |         |            |------------------------------>|    Storage    |
     |            |         |            |           3. return           |       A       |
     |            |         |            |<------------------------------|               |
     |            |         |            |                               +---------------+
     | Controller |         |  Service   |
     |            |         |            |           5. Add              +---------------+
     |    GET     |    1    |  4.logic   |------------------------------>|    Storage    |
     |    POST    | ------> | - Map data |           6. return           |       B       |
     |    PUT     |         | - Validate |<------------------------------|               |
     |    DELETE  |         | - Handle   |                               +---------------+
     |            |         |   status   |
     |            |         |            |              ...              +---------------+
     |            |         |            |------------------------------>|    Storage    |
     |            |         |            |                               |       C       |
     |            |         |            |<------------------------------|               |
     +------------+         +------------+                               +---------------+
```

Each layer will communicate through an interface. For example, The Service layer needs the Storage layer, The Storage will provide an interface to be their actions available or method. The service method can call more one the storage method

The service layer will communicate to the Storage layer using this method, and the Storage layer must implement this interface so it can be used by the Service layer.

Same with the Controller layer, the Service layer will provide an interface. And the Service layer must implement this interface.


## Testable
The business rules can be tested without the UI, Database, Web Server. So we will focus unit test the service layer but this layer depends on the storage layer, which means this layer needs the storage layer for testing. So we must make a mockup of the storage, based on the interface defined before.


## Developers guide

### How to generate migrate file and runs database migrations
This is a step to generate a migration file for making database schema updates.
We need to execute via docker, generate migration file, put SQL command and run migrate
1. `docker-compose exec api-shop sh`
2. `bee generate migration client -fields="name:string,api_key:string,created_at:datetime,updated_at:datetime"`
3. Find your generated file in `database/migrations` and recheck your SQL command and you can change it
4. Final step run the migrate file by `bee migrate -driver=$DRIVER -conn=$SQLCONN`

### Generate Mock 
Example generate mock
`mockgen -destination ./storage/client_mock/client_mock.go -source storage/client.go -package clientmock Client`

### Generate Test 
Example generate test
`gotests -all -w services/client.go`