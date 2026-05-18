// TODO 这里需要根据业务场景，自行修改实现

package service

import (
	"ag-core/ag/ag_common/agmetadata"
	"ag-core/contribute/agredis"
	demob "agdemoB/api/demob"
	"agdemoB/internal/repository/dao"
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"
)

// DemobImpl defines the service implementation for DemobImpl.
type DemobImpl struct {
	studao   dao.IStudentDao
	rediscli agredis.AgRedisClient
}

// NewDemobImpl creates and returns a new DemobImpl instance.
// @param TODO inject dependencies
// @return *DemobImpl
func NewDemobImpl(studao dao.IStudentDao, rediscli agredis.AgRedisClient) *DemobImpl {
	return &DemobImpl{
		studao:   studao,
		rediscli: rediscli,
	}
}

// Hello TODO:DESCRIBE
func (c *DemobImpl) Hello(ctx context.Context, req *demob.BRequest) (*demob.BReply, error) {
	var resp *demob.BReply
	resp = &demob.BReply{}
	// TODO 添加业务处理逻辑

	return resp, nil
}

// NewDemobImplI new DemobImpl instance, and return interface demob.Demob.
// @param TODO inject dependencies
// @return demob.Demob
// func NewDemobImplI( /* TODO inject dependencies */ ) demob.Demob {
// 	return NewDemobImpl( /* TODO inject dependencies */ )
// }

// Calldemob TODO:DESCRIBE
func (c *DemobImpl) Calldemob(ctx context.Context, req *demob.BRequest) (*demob.BReply, error) {
	slog.Info("[I] agdemoB Calldemob")
	defer slog.Info("[O] agdemoB Calldemob")

	md := agmetadata.GetMdFromContext(ctx)
	slog.Info(fmt.Sprintf("md: %v", md))

	name := req.Name

	// 截取name后两位
	dtime := name[len(name)-2:]
	dt, err := strconv.Atoi(dtime)
	if err != nil {
		dt = 0
	}

	stu, err := c.studao.FindByStuno(ctx, 1)
	if err != nil {
		slog.Error(fmt.Sprintf("FindByStuno err: %v", err))
	}
	slog.Info(fmt.Sprintf("根据学号查询: %v", stu))

	stu, err = c.studao.FindByPrimaryKey(ctx, 1)
	if err != nil {
		slog.Error(fmt.Sprintf("FindByPrimaryKey err: %v", err))
	}
	slog.Info(fmt.Sprintf("根据主键查询: %v", stu))

	// arg := &model.StudentFindByWithPageArg{
	// 	Age:      18,
	// 	PageNum:  1,
	// 	PageSize: 5,
	// }
	// stuPageRes, err := c.studao.FindByWithPage(ctx, arg)
	// if err != nil {
	// 	slog.Error(fmt.Sprintf("FindByWithPage err: %v", err))
	// }
	// slog.Info(fmt.Sprintf("分页查询: %v", stuPageRes))

	intcmd := c.rediscli.Incr(ctx, "demobflag")
	flag, err := intcmd.Result()
	if err != nil {
		slog.Error(fmt.Sprintf("Incr err: %v", err))
	}

	msg := fmt.Sprintf("hello demob, flag:%d", flag)

	resp := &demob.BReply{
		Msg: msg,
	}

	// TODO 添加业务处理逻辑
	for i := 1; i < dt+1; i++ {
		slog.Info(fmt.Sprintf("sleep count:%d", i))
		time.Sleep(time.Second)
	}

	return resp, nil
}

// // ClientSideStreaming TODO:DESCRIBE
// // ClientStreaming
// func (c *DemobImpl) ClientSideStreaming(stream demob.Demob_ClientSideStreamingServer) error {
// 	stream.Recv()
// 	// TODO
// 	return nil
// }

// // ServerSideStreaming TODO:DESCRIBE
// // ServerStreaming
// func (c *DemobImpl) ServerSideStreaming(req *demob.BRequest, stream demob.Demob_ServerSideStreamingServer) error {
// 	// TODO
// 	return nil
// }

// // BidiSideStreaming TODO:DESCRIBE
// // ClientStreaming and ServerStreaming
// func (c *DemobImpl) BidiSideStreaming(stream demob.Demob_BidiSideStreamingServer) error {

// 	req, err := stream.Recv()
// 	if err != nil {
// 		return err
// 	}

// 	resp := &demob.BReply{
// 		Msg: req.Name,
// 	}
// 	stream.Send(resp)
// 	// TODO
// 	return nil
// }
