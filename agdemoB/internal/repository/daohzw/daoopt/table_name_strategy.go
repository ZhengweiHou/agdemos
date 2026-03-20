package daoopt

import (
	"context"
	"fmt"
)

type HzwTableNameStrategy struct{}

// 入参为原始表名，返回值为修改后的表名
func (s *HzwTableNameStrategy) TableName(ctx context.Context, tableName string) string {
	newtbname := fmt.Sprintf("%s_%s", tableName, "hzw")
	return newtbname
}
