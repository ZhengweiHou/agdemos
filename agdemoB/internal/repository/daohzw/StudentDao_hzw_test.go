package daohzw

import (
	"agdemoB/internal/repository/model"
	"context"
	"fmt"
	"testing"
)

func TestHzwDao_FindByNaming(t *testing.T) {
	ctx := context.Background()
	dao := &StudentDao{}

	arg := &PageArg{} // 查询参数

	result, err := dao.FindByNaming2(ctx, FindByPageNamingInfo, arg)
	if err != nil {
		t.Errorf("FindByNaming2 failed: %v", err)
	}
	if result == nil {
		t.Errorf("FindByNaming2 result is nil")
	}

	students, ok := result.([]model.Student)
	if !ok {
		t.Errorf("FindByNaming2 result type assertion failed")
	}
	fmt.Printf("FindByNaming2 result: %v\n", students)

	arg2 := &model.Student{}
	result, err = dao.FindByNaming2(ctx, FindByPageNamingInfo, arg2)
	if err != nil {
		t.Errorf("FindByNaming2 failed: %v", err)
	}
	if result == nil {
		t.Errorf("FindByNaming2 result is nil")
	}

	students, ok = result.([]model.Student)
	if !ok {
		t.Errorf("FindByNaming2 result type assertion failed")
	}
	fmt.Printf("FindByNaming2 result: %v\n", students)
}
