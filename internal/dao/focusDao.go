package dao

import (
	"github.com/zxgit/gin-blog-project/internal/model"
)

type Focus struct {
	ID            uint32 `json:"id"`
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Summary       string `json:"summary"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

/**
获取列表
 * @Author ZhangXin
 * @Description
 * @Date 15:28 2021/2/22
 * @Param
 * @return
 **/
func (d *Dao) GetFocusList(page,pageSize int) (list []model.Focus,err error){
	a := model.Focus{}
	return  a.GetFocusList(page,pageSize)
}
