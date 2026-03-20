package test

import (
	"agdemoB/internal/repository/model"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var conhzwmysql = "root:root@tcp(localhost:3306)/houzw?parseTime=True"
var hzwlog = logger.New(log.New(log.Writer(), "\r\n", log.LstdFlags), logger.Config{LogLevel: logger.Info})

func TestStu1(t *testing.T) {
	db := getHzwMysqlDb()
	stu := model.Student{
		Stuno: 1,
		Name:  "sirius",
		Age:   18,
	}

	result := db.Create(&stu)
	if result.Error != nil {
		log.Fatalf("Create failed: %v", result.Error)
	}
	jstr, err := json.Marshal(stu)
	if err != nil {
		log.Fatalf("json marshal failed: %v", err)
	}
	fmt.Printf("Affected: %v Created: %s\n", result.RowsAffected, jstr)
}

func getHzwMysqlDb() *gorm.DB {
	dialector := mysql.Open(conhzwmysql)
	conf := &gorm.Config{
		PrepareStmt: true,
		//		DryRun:      true, // 生成SQL而不执行
		Logger: hzwlog,
	}
	db, err := gorm.Open(dialector, conf)
	if err != nil {
		log.Fatalf("open err: %v", err)
	}
	return db
}
