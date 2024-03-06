package db

import (
	"cheki-back/config"
	"context"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, cfg *config.Config) (*sqlx.DB, func(), error) {
	//DB, err := sqlx.Open("mysql", "cheki:cheki@(172.17.0.2:3306)/cheki")

	//if err != nil {
	//	return DB, nil, nil
	//}
	//err = DB.Ping()
	//if err != nil {
	//	return nil, nil, nil
	//}

	return nil, nil, nil
}
