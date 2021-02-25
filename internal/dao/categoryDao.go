package dao

import (
	"github.com/zxgit/gin-blog-project/internal/model"
)

type Category struct {
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
func (d *Dao) GetCategoryList(page,pageSize int) (list []model.Category,err error){
	a := model.Category{}
	return  a.GetCategoryList(page,pageSize)
}
