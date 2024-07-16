package global

import "gorm.io/gorm"

var (
	DB *gorm.DB
)

func GetGlobalDB() *gorm.DB {
	return DB
}
