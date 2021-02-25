package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID         int    `gorm:"primary_key" json:"id" form:"id"`
	CreatedAt  uint32 `json:"createdAt"  form:"createdAt"`
	ModifiedAt uint32 `json:"modifiedAt" form:"modifiedAt"`
	CreatedBy  uint32 `json:"createdBy" form:"createdBy"`
	ModifiedBy uint32 `json:"modifiedBy" form:"modifiedBy"`
	DeletedAt  int    `json:"deletedAt"  form:"deletedAt"`
	IsDel      int    `json:"isDel"     form:"isDel"`
}
