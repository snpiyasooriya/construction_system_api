package interfaces

import "gorm.io/gorm"

type Database interface {
	GetDb() *gorm.DB
}
