package config

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DBConfig struct {
	Adapter   string `yaml:"adapter"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Dbname    string `yaml:"dbname"`
	Collation string `yaml:"collation"`
}

var (
	SQL_DB      *sql.DB
	DB_MIGRATOR gorm.Migrator
	DB_INSTANCE *gorm.DB
)

// Generate URL for the Database Connection
func (dbConfig *DBConfig) DbURL() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&collation=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Dbname,
		dbConfig.Collation,
	)
}

func initDB(dbConfig *DBConfig) {
	// Ensure MySQL driver is used
	dsn := dbConfig.DbURL()

	// Connecting to Database Server
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	// Set database connection pool parameters (optional)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(0)

	// Set DB Constants
	DB_INSTANCE = db
	SQL_DB = sqlDB
	DB_MIGRATOR = DB_INSTANCE.Migrator()

	// Set the dialect to MySQL for Goose Package
	goose.SetDialect(dbConfig.Adapter)

	fmt.Println("Database connection established")
}

// GetDBConnection returns the database connection instance
func GetDBConnection() *gorm.DB {
	if DB_INSTANCE == nil {
		fmt.Println("Database connection has not been initialized. Please call InitConfig first.")
		return nil
	}
	return DB_INSTANCE
}
