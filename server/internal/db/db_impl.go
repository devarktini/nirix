package db

import (
	"sync"

	"github.com/devarktini/nirix/server/common"
	"github.com/devarktini/nirix/server/config"
	"github.com/devarktini/nirix/server/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PGDB struct {
	db *gorm.DB
}

var pgInstance *PGDB

var once sync.Once

func Setup() {
	once.Do(func() {
		db, err := gorm.Open(postgres.Open(config.GetConfig().DatabaseURL), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			common.GetLogger().Log.Fatal("failed to setup db")
		}
		pgInstance = &PGDB{
			db: db,
		}
		common.GetLogger().Log.Info("DB connection established")
	})
}

// apply all models' migrations here
// not for production use
// only for development and testing purposes
func (pdb *PGDB) Migrate() {
	if err := pdb.db.AutoMigrate(models.Activities{}, models.User{}, models.Application{},
		models.Secret{}, models.NirixData{}, models.Policy{}); err != nil {
		common.GetLogger().Log.Sugar().Fatalf("failed to migrate db: %s", err.Error())
	}
	common.GetLogger().Log.Info("Database migration completed")
}

func GetInstance() *PGDB {
	return pgInstance
}
