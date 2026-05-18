package dao

import (
 db "ag-core/contribute/agdb/gormdb"
 "agdemoB/internal/repository/model"
)

// StudentNamingSqlMap 命名SQL映射
var StudentNamingSqlMap = map[string]string{}

// excludeStudentZeroColNames 插入忽略空值时标记哪些字段需要排除在外
var excludeStudentZeroColNames = map[string]int{}



var FindByAgeNamingInfo = &db.NameingSqlArgInfo{
	SqlName:  "FindByAge",
	ReqType:  (*model.StudentFindByAgeArg)(nil),
	RespType: ([]*model.StudentFindByAgeRes)(nil),
}


var FindByWithPageNamingInfo = &db.NameingSqlArgInfo{
	SqlName:  "FindByWithPage",
	ReqType:  (*model.StudentFindByWithPageArg)(nil),
	RespType: (*model.StudentFindByWithPagePageRes)(nil),
}
