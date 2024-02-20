package db

import (
	"bitbucket.com/finease/backend/pkg/db/migrations"
	"bitbucket.com/finease/backend/pkg/environment/config"
	"context"
	"database/sql"
	"errors"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/golang/glog"

	"gorm.io/gorm"
)

func Migrate(g2 *gorm.DB) {
	m := newGormigrate(g2)

	if err := m.Migrate(); err != nil {
		if errors.Is(err, gormigrate.ErrNoMigrationDefined) {
			glog.Infof("no migration defined, exiting!")
			return
		}
		glog.Fatalf("Could not migrate: %v", err)
	}
}

func newGormigrate(g2 *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(g2, gormigrate.DefaultOptions, migrations.MigrationList)
}

type SessionFactory interface {
	Init(config *config.DbConfig)
	DirectDB() *sql.DB
	New(ctx context.Context) *gorm.DB
	CheckConnection() error
	Close() error
}
