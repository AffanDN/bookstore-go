package pkg

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload" // _ dibutuhkan namun tidak digunakan
)

func InitMySql() (*sqlx.DB, error) {
	// user := os.Getenv("DB_USER")
	// host := os.Getenv("DB_HOST")
	// dbname := os.Getenv("DB_NAME")
	// pass := os.Getenv("DB_PASS")
	// port := os.Getenv("DB_PORT")
	var options []any
	options = append(options, os.Getenv("DB_USER"))
	options = append(options, os.Getenv("DB_PASS"))
	options = append(options, os.Getenv("DB_HOST"))
	options = append(options, os.Getenv("DB_PORT"))
	options = append(options, os.Getenv("DB_NAME"))
	// username:password@tcp(host:port)/dbname
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", options...)
	return sqlx.Connect("mysql", connStr)
}
