package db

import (
	"bitbucket.com/finease/backend/pkg/environment/config"
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type SessionFactory interface {
	Init(config *config.DbConfig)
	DirectDB() *sql.DB
	New(ctx context.Context) *gorm.DB
	CheckConnection() error
	Close() error
}
