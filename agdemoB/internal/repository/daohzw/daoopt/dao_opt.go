package daoopt

import (
	"context"

	"gorm.io/gorm"
)

// type DbOpt func(ctx context.Context, db *gorm.DB) *gorm.DB // 应该区分Opt的类型
type GormScopeFunc func(db *gorm.DB) *gorm.DB

type ITableNameStrategy interface {
	TableName(ctx context.Context, tableName string) string
}

// func WithTableNameStrategy(strategy ITableNameStrategy) DbOpt {
// 	return func(ctx context.Context, db *gorm.DB) *gorm.DB {
// 		newtbname := strategy.TableName(ctx, db.Statement.Table)
// 		return db.Table(newtbname)
// 	}
// }
