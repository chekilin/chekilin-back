package db

import (
	"cheki-back/config"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, cfg *config.Config) (*sqlx.DB, func(), error) {
	var db *sql.DB
	count := 0

	for {
		log.Printf("This is %d", count)
		dbTemp, err := sql.Open("mysql",
			fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?parseTime=true",
				cfg.DBUser, cfg.DBPassword,
				cfg.DBHost, cfg.DBPort,
				cfg.DBName,
			),
		)
		if err != nil {
			log.Printf("count: %d  Failed to open db: %v", count, err)
			if count > 10 {
				return nil, nil, err
			}
			count += 1
			continue
		}
		db = dbTemp
		fmt.Println("Succeeded to connect DB")
		break
	}
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	time.Sleep(5 * time.Second)

	xdb := sqlx.NewDb(db, "mysql")
	return xdb, func() { _ = db.Close() }, nil
}
