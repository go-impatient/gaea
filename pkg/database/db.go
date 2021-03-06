package database

import (
	"context"
	"database/sql"
	logd "log"
	"sync"
	"time"

	entsql "entgo.io/ent/dialect/sql"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"

	"moocss.com/gaea/internal/data/ent"
	"moocss.com/gaea/internal/data/ent/migrate"
	"moocss.com/gaea/pkg/conf"
	"moocss.com/gaea/pkg/log"
)

// AsDefault alias for "default"
const AsDefault = "default"

var (
	defaultSQL *ent.Client
	sqlMap     sync.Map
)

// Init 初始化数据库
func Init(logger log.Logger) (client *ent.Client, err error) {
	log := log.NewHelper("database", logger)

	config := conf.File("config")
	dialect := config.Get("database.dialect")
	dsn := config.Get("database.dsn")

	var db *sql.DB
	db, err = sql.Open(dialect, dsn)
	if err != nil {
		return
	}

	maxOpenConns := config.GetInt("database.max_open_conns")
	if maxOpenConns < 0 {
		maxOpenConns = 100
	}

	maxIdleConns := config.GetInt("database.max_idle_conns")
	if maxIdleConns < 0 {
		maxIdleConns = 10
	}

	connMaxLifetime := config.GetDuration("database.conn_max_lifetime")
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
	logging := config.GetBool("database.logging")
	if logging {
		opts = append(opts, ent.Debug())
		opts = append(opts, ent.Log(logger.Print))
	}
	client = ent.NewClient(opts...)

	// Run Database Setup/Migrations
	ctx := context.Background()
	mode := config.Get("app.mode")
	AutoMigration(ctx, client, log)
	if mode == "debug" {
		DebugMode(ctx, client, log)
	}

	// 全局数据库客户端服务
	defaultSQL = client
	sqlMap.Store("default", client)

	return
}

// AutoMigration .
func AutoMigration(ctx context.Context, client *ent.Client, logger *log.Helper) error {
	err := client.Schema.Create(context.Background())
	if err != nil {
		logger.Errorf("failed creating schema resources: %s", err.Error())
		return err
	}
	return nil
}

// DebugMode .
func DebugMode(ctx context.Context, client *ent.Client, logger *log.Helper) error {
	err := client.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		logger.Errorf("failed creating schema resources: %s", err.Error())
		return err
	}
	return nil
}

// GetDatabase 全局数据库客户端服务
func GetDatabase(name ...string) *ent.Client {
	if len(name) == 0 || name[0] == AsDefault {
		if defaultSQL == nil {
			logd.Panicf("Invalid db `%s` \n", AsDefault)
		}
		return defaultSQL
	}

	v, ok := sqlMap.Load(name[0])
	if !ok {
		logd.Panicf("Invalid db `%s` \n", name[0])
	}

	return v.(*ent.Client)
}
