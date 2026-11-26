package db

import (
	"context"
	"database/sql"
	"fmt"
	"hoc-golang/internal/config"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	connStr := config.NewConfig().DNS()

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", err)
	}

	DB.SetConnMaxLifetime(30 * time.Minute) // Đóng kết nối sau bao nhiêu phút
	DB.SetMaxIdleConns(3)                   //Idle là số kết nối nhàn rỗi (sau khi truy vấn xong nhưng vẫn còn) tối đa
	DB.SetMaxOpenConns(30)                  // Số kết nối tối đa
	DB.SetConnMaxIdleTime(5 * time.Minute)  // Đóng kết nối nhàn rỗi sau 5'

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := DB.PingContext(ctx); err != nil {
		DB.Close()
		return fmt.Errorf("DB ping error: %w", err)
	}

	log.Println("Connected")

	return nil
}
