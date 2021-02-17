package database

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	entsql "entgo.io/ent/dialect/sql"

	"moocss.com/gaea/internal/data/ent"
	"moocss.com/gaea/internal/data/ent/migrate"
	"moocss.com/gaea/pkg/conf"
	logx "moocss.com/gaea/pkg/log"
)

// AsDefault alias for "default"
const AsDefault = "default"

var (
	defaultSQL *ent.Client
	sqlMap     sync.Map
)

// 初始化数据库
func InitDatabase(logger logx.Logger) (client *ent.Client, err error) {
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
	db.SetMaxIdleConns(maxIdleConns)       // 用于设置连接池中空闲连接的最大数量。
	db.SetMaxOpenConns(maxOpenConns)       // 设置打开数据库连接的最大数量。
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
		opts = append(opts, ent.Log(logger.Print))
	}
	client = ent.NewClient(opts...)

	// Run Database Setup/Migrations
	ctx := context.Background()
	mode := conf.Get("app.mode")
	AutoMigration(ctx, client, logger)
	if mode == "debug" {
		DebugMode(ctx, client, logger)
	}

	// 全局数据库客户端服务
	defaultSQL = client
	sqlMap.Store("default", client)

	return
}

// AutoMigration .
func AutoMigration(ctx context.Context, client *ent.Client, logger logx.Logger) error {
	err := client.Schema.Create(context.Background())
	if err != nil {
		logger.Printf("failed creating schema resources: %s", err.Error())
		return err
	}
	return nil
}

// DebugMode .
func DebugMode(ctx context.Context, client *ent.Client, logger logx.Logger) error {
	err := client.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		logger.Printf("failed creating schema resources: %s", err.Error())
		return err
	}
	return nil
}

// GetDatabase 全局数据库客户端服务
func GetDatabase(name ...string) *ent.Client {
	if len(name) == 0 || name[0] == AsDefault {
		if defaultSQL == nil {
			log.Panicf("Invalid db `%s` \n", AsDefault)
		}
		return defaultSQL
	}

	v, ok := sqlMap.Load(name[0])
	if !ok {
		log.Panicf("Invalid db `%s` \n", name[0])
	}

	return v.(*ent.Client)
}
