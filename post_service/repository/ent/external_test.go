package ent

import (
	"context"
	"cw_post_service/conf"
	"fmt"
	"log"
	"testing"

	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var client *Client

func Init() {
	// init database connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		"127.0.0.1", "postgres", "123456", "postgres", 5432)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("db connection error: ", err)
	}
	if conf.Cfg.DB.MaxIdle > 0 {
		db.SetMaxIdleConns(conf.Cfg.DB.MaxIdle)
	}
	if conf.Cfg.DB.MaxOpen > 0 {
		db.SetMaxOpenConns(conf.Cfg.DB.MaxOpen)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	client = NewClient(Driver(drv))
}

func TestGetPosts(t *testing.T) {
	Init()

	posts, err := GetPosts(context.Background(), client, 1, 50)
	if err != nil {
		t.Error(err)
	}
	for _, post := range posts {
		t.Log(post)
	}
}
