package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"user/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	//TODO 模拟db查询操作
	return &biz.User{
		Id:       123456,
		UserName: "张三",
		Age:      22,
		Sex:      false,
	}, nil
}

func (u *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	//TODO 模拟CreateUser操作
	log.Info("CreateUser %v", user)
	return nil
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
