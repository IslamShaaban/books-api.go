Book APIs GoLang Project
========================

Book APIs is a Go JSON API framework that is built on [Gin](https://github.com/gin-gonic/gin),
[Gorm](https://gorm.io) and [Goose](https://github.com/pressly/goose). 

Current Features
----------------
- Uses Gorm as an ORM
- Database migrations via Goose
- Simple routing via Gin
- File watcher to auto reboot server during development

Setup
------
Create a `.env` in the root directory of the project. The required environment variables are:

## Environment Variables

The following environment variables are required for your application to connect to the database and execute migrations:

| Variable         | Description                                                             |
|-------------------|----------------------------------------------------------------------|
| DB_ADAPTER       | Database adapter to use (e.g., `mysql`, `postgres`)                    |
| DB_USERNAME       | Username for database authentication                                  |
| DB_PASSWORD       | Password for database authentication      |
| DB_HOST           | Hostname or IP address of the database server                        |
| DB_PORT           | Port number on which the database server listens                      |
| DB_NAME           | Name of the database to connect to                                     |
| SERVER_HOSTNAME   | Hostname or IP address of your application server (optional)           |
| SERVER_PORT       | Port number on which your application server listens (optional)       |
| GOOSE_DRIVER      | Same value as `DB_ADAPTER` (used by Goose for database connection)      |
| GOOSE_DBSTRING     | Database connection string for Goose (constructed using other variables) |
| GOOSE_MIGRATION_DIR | Path to the directory containing your database migration files       |

**Important Note:**

- Replace the placeholder values (USERNAME, PASSWORD, etc.) with your actual database credentials and configuration before running your application or migrations.

## Example Database Connection String
Here's an example MySQL connection string for reference:
- `user@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local`

**Explanation:**
- `user`: Username for authenticating with the MySQL server.
- `@tcp(127.0.0.1:3306)`: Connection protocol (`tcp`) and host (`127.0.0.1` for localhost). Replace with the actual hostname or IP address if different. The port number is `3306`, the default for MySQL.
- `/dbname`: The name of the database to connect to.

**Optional Parameters:**

- `charset=utf8mb4`: Sets the character encoding for the connection (recommended for Unicode support).
- `parseTime=True`: Instructs the driver to parse timestamps returned by the server.
- `loc=Local`: Sets the location for time zone handling (optional, adjust based on your needs).

**Important Notes:**

- Replace `user`, `password`, and `dbname` with your actual database credentials and database name.
- Adjust `charset`, `parseTime`, and `loc` parameters if necessary for your specific application requirements.

Commands
--------

- `go mod tidy`
  - Install all packages required

- `go build`
  - Build the project

- `go run .` or `go run . serve`
  - Runs the webserver
  - Also sets up a watcher and will auto restart webserver when files are changed

- `GO_ENV=env go run .` or `GO_ENV=env go run . serve`
  - Runs the webserver on Speicfic Environment by Default will be Run in `development` Mode

- `go run . migrate [status|up|down|create|etc...]`
  - Runs goose migrations, see goose's documentation for more details

- `go run . generator -type [model|controller|service] -name [ex:user] -fields [name:type,name:type,...]`
  - `-fields` is only used with model type
  - Use singular names, the name will automatically be pluralized when needed (example: `user` not `users`)
  - Example: `go run . generator -type model -name book -fields title:string,category_id:uint,pages:uint`
    - Creates a single file `book.go` in the `/models` directory
  - Example: `go run . generator -type controller -name book` or `go run . g -type c -name book`
    - Creates controller file inside the `app/controllers` directory
      - `books_controllers.go`