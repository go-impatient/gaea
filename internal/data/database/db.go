package data

import (
	"context"
	"database/sql"
	"log"
	"time"

	entsql "github.com/facebook/ent/dialect/sql"
	"moocss.com/gaea/internal/data/ent"
	"moocss.com/gaea/pkg/conf"
)

// 初始化数据库
func InitDatabase() (client *ent.Client, err error) {
	dialect := conf.Get("database.dialect")
	dsn := conf.Get("database.url")
	var db *sql.DB
	db, err = sql.Open(dialect, dsn)
	if err != nil {
		return
	}

	maxOpenConns := conf.GetInt("database.max_open_conns")
	if maxOpenConns < 0 {
		maxOpenConns = 100
	}

	maxIdleConns := conf.GetInt("database.max_idle_conns")
	if maxIdleConns < 0 {
		maxIdleConns = 10
	}

	connMaxLifetime := conf.GetDuration("database.conn_max_lifetime")
	if connMaxLifetime < 0 {
		connMaxLifetime = 10 * time.Minute
	}

	// 数据库调优
	db.SetMaxIdleConns(maxIdleConns)                  // 用于设置连接池中空闲连接的最大数量。
	db.SetMaxOpenConns(maxOpenConns)                 // 设置打开数据库连接的最大数量。
	db.SetConnMaxLifetime(connMaxLifetime) // 设置了连接可复用的最大时间。

	// ping test
	err = db.Ping()
	if err != nil {
		return
	}

	drv := entsql.OpenDB(dialect, db)

	opts := []ent.Option{ent.Driver(drv)}
	logging := conf.GetBool("database.logging")
	if logging {
		opts = append(opts, ent.Debug())
		opts = append(opts, ent.Log(log.Print))
	}
	client = ent.NewClient(opts...)

	// Run Database Setup/Migrations
	if err = client.Schema.Create(context.Background()); err != nil {
		log.Print("failed creating schema resources")
		return nil, err
	}

	return
}
