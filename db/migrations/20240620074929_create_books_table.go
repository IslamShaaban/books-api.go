package migrations

import (
	"books-api/app/models"
	database "books-api/config"
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateBooksTable, downCreateBooksTable)
}

func upCreateBooksTable(ctx context.Context, tx *sql.Tx) error {
	return database.DB_MIGRATOR.CreateTable(&models.Book{})
}

func downCreateBooksTable(ctx context.Context, tx *sql.Tx) error {
	return database.DB_MIGRATOR.DropTable(&models.Book{})
}
