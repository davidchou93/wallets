package db

import (
	"fmt"
	"github.com/davidchou93/wallets/internal/logging"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"time"

	"github.com/davidchou93/wallets/config"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	Debug    bool
}

func NewDBConfig(config *config.Config) *DBConfig {
	opt := DBConfig{}
	opt.Host = config.DB.Host
	opt.Port = config.DB.Port
	opt.User = config.DB.User
	opt.DBName = config.DB.DBName
	opt.Password = config.DB.Password
	return &opt
}

func NewPgConn(dbConf *DBConfig) (*sqlx.DB, error) {
	var (
		db  *sqlx.DB
		err error
	)
	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=5",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.DBName)

	// TODO: use stdlib.RegisterConnConfig() to inject more db Option
	connConfig, err := pgx.ParseConfig(connectionStr)
	if err != nil {
		return nil, err
	}
	// connConfig.PreferSimpleProtocol = true
	connConfig.RuntimeParams["timezone"] = "UTC"
	if dbConf.Debug {
		log := logging.NewLogger()
		connConfig.Logger = logrusadapter.NewLogger(log)
		connConfig.LogLevel = pgx.LogLevelTrace
	}
	connectionStr = stdlib.RegisterConnConfig(connConfig)

	// TODO: retry or reconnect
	if db, err = sqlx.Open("pgx", connectionStr); err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(500)
	return db, nil
}
