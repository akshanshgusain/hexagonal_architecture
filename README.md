# Hexagonal Architecture
### Implementation in Go

##### Run `make run` to run the application

To run the application, you have to define the environment variables, default values of the variables are defined inside `start.sh`

- SERVER_ADDRESS    `[IP Address of the machine]`
- SERVER_PORT       `[Port of the machine]`
- DB_USER           `[Database username]`
- DB_PASSWD         `[Database password]`
- DB_ADDR           `[IP address of the database]`
- DB_PORT           `[Port of the database]`
- DB_NAME           `[Name of the database]`

## Postgres database
You can use any one of the following procedure to make a database instance.
1. Run `make start_db` to run an instance of postgres inside a docker container. The `docker-compose.yaml` file can be found inside 
`resources/postgres`


2. `resources/database.sql` this contains the SQL for generating the tables. In case you dont want to use the docker-compose file you can use this file to generate tables and insert the default data

# mocks generator
`make mock`

# run unit tests
`make test`