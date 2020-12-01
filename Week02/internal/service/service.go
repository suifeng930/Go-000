package service

import (
	"context"
	"github.com/Go-000/Week02/global"
	"github.com/Go-000/Week02/internal/dao"
)

type Service struct {
	ctx context.Context //上下文对象
	dao *dao.Dao        //dao 指针
}

func New(ctx context.Context) Service {

	svc := Service{ctx: ctx}           //初始化上下文
	svc.dao = dao.New(global.DBEngine) //获取dao 指针
	return svc

}


