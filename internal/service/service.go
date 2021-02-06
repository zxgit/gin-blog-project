package service

import (
	"context"
	"github.com/zxgit/gin-blog-project/internal/dao"
)

type Service struct {
	Ctx context.Context
	Dao *dao.Dao
}
func New(ctx context.Context) Service {
 return Service{Ctx: ctx,Dao: dao.New()}
}