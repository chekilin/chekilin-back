package db

import (
	"context"
	"fmt"
	"log"

	"cheki-back/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, cfg *config.Config) (*sqlx.DB, func(), error) {
	var db *sqlx.DB
	count := 0

	for {
		log.Printf("This is %d", count)
		xdb, err := sqlx.Open("mysql",
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
		fmt.Println("Succeeded to connect DB")
		return xdb, func() { _ = db.Close() }, nil
	}
}
