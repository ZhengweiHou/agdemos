package service

import (
	"ag-core/contribute/agdb/conditonwhere"
	"ag-core/contribute/agdb/gormdb"
	"ag-core/contribute/agredis"
	hello "agdemoC/api/hello"
	"agdemoC/internal/repository/dao"
	"agdemoC/internal/repository/model"
	"context"
	"fmt"
	"log/slog"
)

// HelloImpl defines the service implementation for HelloImpl.
type HelloImpl struct {
	studao   dao.IStudentDao
	rediscli agredis.AgRedisClient
}

// NewHelloImpl creates and returns a new HelloImpl instance.
func NewHelloImpl(
	studao dao.IStudentDao,
	rediscli agredis.AgRedisClient,
) *HelloImpl {

	return &HelloImpl{
		studao:   studao,
		rediscli: rediscli,
	}
}

// Pinghello
func (c *HelloImpl) Pinghello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloReply, error) {
	slog.Info("[I] agdemoB Calldemob")
	defer slog.Info("[O] agdemoB Calldemob")

	studentFindByAgeWithPageArg := &model.StudentFindByAgeWithPageArg{
		Page: gormdb.Page{
			PageNum:  1,
			PageSize: 3,
		},
		FieldMask: conditonwhere.NewFieldMask(),
	}
	studentFindByAgeWithPageArg.WithAge(18)

	result, err := c.studao.FindByCustomerRule(
		ctx,
		dao.FindByAgeWithPageNamingInfo,
		studentFindByAgeWithPageArg,
	)
	if err != nil {
		return nil, err
	}
	pageRes := result.(*model.StudentFindByAgeWithPagePageRes)
	stus := pageRes.ResultList

	slog.Info("stus", stus)

	msg := "hello world"

	msg += fmt.Sprintf("%s %s", msg, req.Name)

	var resp *hello.HelloReply
	resp = &hello.HelloReply{
		Msg: msg,
	}

	return resp, nil
}
