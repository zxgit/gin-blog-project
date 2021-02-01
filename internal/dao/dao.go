package dao

import "github.com/jinzhu/gorm"

type Dao struct {
	Mdb *gorm.DB
}
