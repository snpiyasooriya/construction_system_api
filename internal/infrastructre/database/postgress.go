package database

import (
	"fmt"
	"github.com/snpiyasooriya/construction_design_api/config"
	"github.com/snpiyasooriya/construction_design_api/internal/interfaces"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

type postgresDb struct {
	Db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *postgresDb
)

func NewPostgres(conf *config.Config) interfaces.Database {
	dbConf := conf.Db
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", dbConf.Host, dbConf.User, dbConf.Password, dbConf.Db, dbConf.Port)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("failed to connect database: %v", err))
		}
		dbInstance = &postgresDb{
			Db: db,
		}
	})
	return dbInstance
}

func (db *postgresDb) GetDb() *gorm.DB {
	return dbInstance.Db
}
