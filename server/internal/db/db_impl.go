package db

import (
	"sync"

	"github.com/devarktini/nirix/server/common"
	"github.com/devarktini/nirix/server/config"
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

func (pdb *PGDB) Migrate() {}

func GetInstance() *PGDB {
	return pgInstance
}
