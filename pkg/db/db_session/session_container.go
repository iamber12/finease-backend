package db_session

import (
	"bitbucket.com/finease/backend/pkg/db"
	"bitbucket.com/finease/backend/pkg/environment/config"
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbSessionContainer struct {
	config *config.DbConfig
	g2     *gorm.DB
	db     *sql.DB
}

var _ db.SessionFactory = &DbSessionContainer{}

func NewSessionFactory(config *config.DbConfig) *DbSessionContainer {
	conn := &DbSessionContainer{}
	conn.Init(config)
	return conn
}

func (f *DbSessionContainer) Init(config *config.DbConfig) {
	var (
		dbx *sql.DB
		g2  *gorm.DB
		err error
	)

	// Open connection to DB via standard library
	dbx, err = sql.Open(config.Dialect, config.ConnectionString())
	if err != nil {
		dbx, err = sql.Open(config.Dialect, config.ConnectionString())
		if err != nil {
			log.Fatalf("SQL connection failed")
		}
	}
	dbx.SetMaxOpenConns(config.MaxOpenConnections)

	// Connect GORM to use the same connection
	conf := &gorm.Config{
		PrepareStmt:          false,
		FullSaveAssociations: false,
	}
	g2, err = gorm.Open(postgres.New(postgres.Config{
		Conn:                 dbx,
		PreferSimpleProtocol: true,
	}), conf)
	if err != nil {
		log.Fatalf("Connection establishment failed via GORM")
	}

	f.config = config
	f.g2 = g2
	f.db = dbx
}

func (f *DbSessionContainer) DirectDB() *sql.DB {
	return f.db
}

func (f *DbSessionContainer) New(ctx context.Context) *gorm.DB {
	conn := f.g2.Session(&gorm.Session{
		Context:              ctx,
		Logger:               f.g2.Logger.LogMode(logger.Silent),
		FullSaveAssociations: true,
	})
	if f.config.Debug {
		conn = conn.Debug()
	}
	return conn
}

func (f *DbSessionContainer) CheckConnection() error {
	return f.g2.Exec("SELECT 1").Error
}

func (f *DbSessionContainer) Close() error {
	return f.db.Close()
}
