package v2

import (
	"fmt"
)

type HArg struct{}
type HResult struct{}
type HArg2 struct{}
type HResult2 struct{}

// NamingInfo 泛型结构体
type NamingInfo[R, T any] struct {
	Name  string
	AType R
	RType T
}

// INamingExecutor 接口 - 类型擦除
type INamingExecutor interface {
	Execute() any
}

// NamingExecutor 泛型实现
type NamingExecutor[R, T any] struct {
	info NamingInfo[R, T]
}

func (e *NamingExecutor[R, T]) Execute() any {
	// 这里可以实现具体的逻辑
	// 返回 []*T 类型
	var result []*T
	return result
}

type HzwDao struct{}

// NamingSql 方法 - 非泛型，返回接口
func (h *HzwDao) NamingSql(info any) (any, error) {
	// 根据 info 的类型创建对应的执行器
	switch v := info.(type) {
	case NamingInfo[HArg, HResult]:
		executor := &NamingExecutor[HArg, HResult]{info: v}
		return executor.Execute(), nil
	case NamingInfo[HArg2, HResult2]:
		executor := &NamingExecutor[HArg2, HResult2]{info: v}
		return executor.Execute(), nil
	default:
		return nil, fmt.Errorf("unsupported naming info type: %T", info)
	}
}

// 或者使用注册模式
func (h *HzwDao) NamingSql2(info NamingInfo[any, any]) any {
	// 这里需要使用反射来处理
	// 但 API 更简洁
	return nil
}
