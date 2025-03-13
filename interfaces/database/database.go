package database

import "gorm.io/gorm"

type DataBaseInterface interface {
	GetDb() *gorm.DB
}
