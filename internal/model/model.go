package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID        int    `gorm:"primary_key" json:"id" form:"id"`
	CreatedAt uint32 `json:"createdAt"  form:"createdAt"`
	UpdatedAt uint32 `json:"updatedAt"  form:"updatedAt"`
	DeletedAt int    `json:"deletedAt"  form:"deletedAt"`
	IsDel     string `json:"is_del" form:"isDel"`
}
