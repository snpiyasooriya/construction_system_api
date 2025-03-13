package utils

import (
	"gorm.io/gorm"
	"time"
)

func ConvertTimeToDeletedAt(t time.Time) gorm.DeletedAt {
	return gorm.DeletedAt{
		Time:  t,
		Valid: !t.IsZero(),
	}
}

func ConvertDeletedAtToTime(d gorm.DeletedAt) *time.Time {
	return &d.Time
}
