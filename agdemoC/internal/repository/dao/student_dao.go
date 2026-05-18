package dao

import (
	"ag-core/contribute/agdb/gormdb"
	"agdemoC/internal/repository/model"
	"context"
	"reflect"
	"errors"
	"ag-core/contribute/agdb/conditonwhere"

	agdao "ag-core/contribute/agdb/agdao"
	"strings"

	"gorm.io/gorm"
)

// StudentDao STUDENT DAO
// DO NOT EDIT
// DO NOT EDIT
// DO NOT EDIT
type StudentDao struct {
	*gormdb.Repository
	info    agdao.TableInfo
	baseDao agdao.BaseDao
}

// IStudentDao Student DAO接口
type IStudentDao interface {
	InsertOne(ctx context.Context, entity *model.Student) (int64, error)
	InsertOneIgnoreZeroValCols(ctx context.Context, entity *model.Student) (int64, error)
	UpdateByPrimaryKey(ctx context.Context, entity *model.Student) (int64, error)
	UpdateByPrimaryKeyIngoreZeroValCols(ctx context.Context, entity *model.Student) (int64, error)
	FindByPrimaryKey(ctx context.Context, id model.StudentPrimaryKey) (*model.Student, error)
	FindByStruct(ctx context.Context, entity *model.Student) ([]*model.Student, error)
	FindByCustomerRule(ctx context.Context, namingInfo *gormdb.NameingSqlArgInfo, args any) (any, error)
	FindByCondition(ctx context.Context, condition *conditonwhere.WhereClauseBuilder, orderBuilder *gormdb.OrderBuilder, page *gormdb.Page) ([]*model.Student, *gormdb.PageResult, error)
	FindFirstOneByCondition(ctx context.Context, condition *conditonwhere.WhereClauseBuilder, orderBuilder *gormdb.OrderBuilder) (*model.Student, error)
}

// NewStudentDao get dao instance
func NewStudentDao(repository *gormdb.Repository, baseDao agdao.BaseDao) IStudentDao {
	InitStudentNamingSql()
	return &StudentDao{
		Repository: repository,
		baseDao:    baseDao,
		info: agdao.TableInfo{
			TableName: "STUDENT",
		},
	}
}

// insertOne 插入一条数据库数据
func (dao *StudentDao) InsertOne(ctx context.Context, entity *model.Student) (int64, error) {
	db, err := dao.newDB(ctx)
	if err != nil {
		return 0, err
	}

	result := db.Create(entity)
	return result.RowsAffected, result.Error
}

// InsertOneIgnorenNullCols 插入数据时，自动剔除零值的列
func (dao *StudentDao) InsertOneIgnoreZeroValCols(ctx context.Context, entity *model.Student) (int64, error) {
	// 1. 剔除结构体中除主键和索引以及特殊列之外的零值列
	colnames,_,err:=entity.ListZeroValueCols(true, true, false, true)
	if err!= nil{
		return 0, err	
	}
	db, err := dao.newDB(ctx)
	if err != nil {
		return 0, err
	}

	result := db.Omit(colnames...).Create(entity)
	return result.RowsAffected, result.Error
}

// UpdateByPrimaryKey 根据主键或者唯一键更新，该操作只适合从数据库查询原实体修改值之后使用
func (dao *StudentDao) UpdateByPrimaryKey(ctx context.Context, entity *model.Student) (int64, error) {
	db, err := dao.newDB(ctx)
	if err != nil {
		return 0, err
	}

	// 4. 更新条件（主键）
	where := make(map[string]any)
	// 检查主键是否为空，如果为空继续检查唯一键
	if ((entity.Id == 0)) {
		return 0, errors.New("when update,primary key or unique key is required")
	} else {
		where["ID"] = entity.Id
	}

	if len(where) == 0 {
		return 0, errors.New("when update,primary key or unique key is required")
	}
	// 5. 使用支持更新的列
	result := db.Model(&model.Student{}).Where(where).Save(entity)
	return result.RowsAffected, result.Error
}

// UpdateByPrimaryKeyIngoreZeroValCols 根据主键或者唯一键更新，自动剔除参数中的零值列
func (dao *StudentDao) UpdateByPrimaryKeyIngoreZeroValCols(ctx context.Context, entity *model.Student) (int64, error) {
	db, err := dao.newDB(ctx)
	if err != nil {
		return 0, err
	}	
	// 4. 更新条件（主键）
	where := make(map[string]any)
	// 检查主键是否为空，如果为空继续检查唯一键
	if ((entity.Id == 0)) {
		return 0, errors.New("when update,primary key or unique key is required")
	} else {
		where["ID"] = entity.Id
	}

	if len(where) == 0 {
		return 0, errors.New("when update,primary key or unique key is required")
	}
	// 使用支持更新的列
	result := db.Model(&model.Student{}).Where(where).Updates(entity)
	return result.RowsAffected, result.Error
}

// FindByPrimaryKey 根据主键查询
func (dao *StudentDao) FindByPrimaryKey(ctx context.Context, id model.StudentPrimaryKey) (*model.Student, error) {
	db, err := dao.newDB(ctx)
	if err != nil {
		return nil, err
	}
	
	var entity model.Student
	result := db.Where("ID = ?", id).First(&entity)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &entity, result.Error
}

// FindByStruct 根据实体查询
func (dao *StudentDao) FindByStruct(ctx context.Context, entity *model.Student) ([]*model.Student, error) {
	var list []*model.Student
	db, err := dao.newDB(ctx)
	if err != nil {
		return nil, err
	}

	// 检查主键是否为空
	if entity.Id != 0 {
		db = db.Where("ID = ?", entity.Id)
		result := db.Find(&list)
		return list, result.Error
	}

	// 检查索引列，确保使用了索引
	indexUsed := false
	// 检查索引 INDEX_AGE
	if entity.Age != 0 {
		db = db.Where("AGE = ?", entity.Age)
		indexUsed = true
	}
	// 检查索引 STUNO_UNIUQE_1
	if entity.Stuno != 0 {
		db = db.Where("STUNO = ?", entity.Stuno)
		indexUsed = true
	}

	// 检查索引列，确保使用了索引
	if !indexUsed {
		return nil, errors.New("query not use any index")
	}

	// 除了主键和索引以外的其他列如果有值，也作为查询条件
	colnames, colvals, err := entity.ListZeroValueCols(true, true, true, false)
	if err != nil {
		return nil, err
	}
	if len(colnames) > 0 {
		for i, colname := range colnames {
			db = db.Where(colname+" = ?", colvals[i])
		}
	}

	// 执行查询
	result := db.Find(&list)
	return list, result.Error
}

// FindByCustomerRule 根据自定义规则查询
func (dao *StudentDao) FindByCustomerRule(ctx context.Context, namingInfo *gormdb.NameingSqlArgInfo, args any) (any, error) {

	if ctx == nil {
		return nil, errors.New("ctx is nil")
	}

	if namingInfo == nil {
		return nil, errors.New("namingInfo is nil")
	}

	if namingInfo.SqlName == "" {
		return nil, errors.New("namingInfo.SqlName is empty")
	}

	// 判断请求参数类型和实际类型是否一致
	reqType := reflect.TypeOf(namingInfo.ReqType)
	reqValue := reflect.ValueOf(args)
	if reqType != reqValue.Type() {
		return nil, errors.New("req type not match")
	}
	switch namingInfo.SqlName {
	case "FindByAge":
		return dao.doFindByAge(ctx, namingInfo, args)
	case "FindByAgeWithPage":
		return dao.doFindByAgeWithPage(ctx, namingInfo, args)
	default:
		return nil, errors.New("not found naming sql")
	}
}

// FindByCondition 根据条件构建器查询
func (dao *StudentDao) FindByCondition(ctx context.Context, condition *conditonwhere.WhereClauseBuilder, orderBuilder *gormdb.OrderBuilder, page *gormdb.Page) ([]*model.Student, *gormdb.PageResult, error) {
	var list []*model.Student
	db, err := dao.newDB(ctx)
	if err != nil {
		return nil, nil, err
	}

	// 主动使用where条件
	where, args, err := condition.Build()
	if err != nil {
		return nil, nil, err
	}
	// 主动拼接where条件
	db = db.Where(where, args...)

	var totalCount int64
	// 统计总数
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, nil, err
	}

	var pageResult *gormdb.PageResult
	// 如果需要分页
	if page != nil {
		start, _, totalPage, enablePage, err := gormdb.CalcPageStartRecord(page.PageNum, page.PageSize, totalCount, dao.DbType)
		if err != nil {
			return nil, nil, err
		}
		pageResult = &gormdb.PageResult{
			CurrentPage: page.PageNum,
			PageSize:    page.PageSize,
			TotalCount:  totalCount,
			TotalPage:   totalPage,
		}
		// 总记录数为0或者当前页码超过总页数时，不执行查询，直接返回空结果和分页信息
		if !enablePage {
			return nil, pageResult, nil
		}
		db = db.Limit(int(page.PageSize)).Offset(int(start))
	}

	// 主动拼排序条件
	if orderBuilder != nil {
		db = db.Order(orderBuilder.Build())
	}

	result := db.Find(&list)
	if result.Error != nil {
		return nil, nil, result.Error
	}

	return list, pageResult, nil
}

// FindFirstOneByCondition 根据条件构建器查询第一条记录
func (dao *StudentDao) FindFirstOneByCondition(ctx context.Context, condition *conditonwhere.WhereClauseBuilder, orderBuilder *gormdb.OrderBuilder) (*model.Student, error) {
	var entity model.Student
	db, err := dao.newDB(ctx)
	if err != nil {
		return nil, err
	}

	// 主动使用where条件
	where, args, err := condition.Build()
	if err != nil {
		return nil, err
	}
	// 主动拼接where条件
	db = db.Where(where, args...)

	// 主动拼排序条件
	if orderBuilder != nil {
		db = db.Order(orderBuilder.Build())
	}

	result := db.Limit(1).Find(&entity)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &entity, result.Error
}

// doFindByAge 执行FindByAge查询（非分页）
func (dao *StudentDao) doFindByAge(ctx context.Context, namingInfo *gormdb.NameingSqlArgInfo, args any) ([]*model.StudentFindByAgeRes, error) {

	queryArgs, ok := args.(*model.StudentFindByAgeArg)
	if !ok {
		return nil, errors.New("doFindByAge args type not match")
	}

	sqlName := dao.DbType + "_" + "Student" + "_" + namingInfo.SqlName
	execSql := StudentNamingSqlMap[sqlName]
	if execSql == "" {
		return nil, errors.New("not found naming sql")
	}

	oldwhere,_:=conditonwhere.ExtractWhereClauseByCut(execSql)
	newwhere,err:=queryArgs.FieldMask.BuildWhereFromConfig("FindByAge",model.StudentConditionMap)
	if err != nil {
		return nil, err
	}
	// 校验新的where条件是否使用了索引列，避免全表扫描
	check:=conditonwhere.ValidateLeadingCol(newwhere,model.StudentIndexLeadingCols)
	if !check {
		return nil, errors.New("query not use any index")
	}
	// 替换where条件
	execSql = strings.Replace(execSql, "WHERE ("+oldwhere+")", newwhere, 1)

	newTableName := dao.getApplyInfo(ctx).TableName
	if newTableName != "" {
		enity := &model.Student{}
		execSql = strings.ReplaceAll(execSql, "FROM "+enity.TableName()+" WHERE", "FROM "+newTableName+" WHERE")
	}

	argsMap := queryArgs.ConvertToMap()
	var list []*model.StudentFindByAgeRes
	result := dao.DB(ctx).Raw(execSql, argsMap).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}

// doFindByAgeWithPage 执行FindByAgeWithPage查询（分页）
func (dao *StudentDao) doFindByAgeWithPage(ctx context.Context, namingInfo *gormdb.NameingSqlArgInfo, args any) (*model.StudentFindByAgeWithPagePageRes, error) {

	queryArgs, ok := args.(*model.StudentFindByAgeWithPageArg)
	if !ok {
		return nil, errors.New("doFindByAgeWithPage args type not match")
	}

	sqlName := dao.DbType + "_" + "Student" + "_" + namingInfo.SqlName
	execSql := StudentNamingSqlMap[sqlName]
	if execSql == "" {
		return nil, errors.New("not found naming sql")
	}
	oldwhere,_:=conditonwhere.ExtractWhereClauseByCut(execSql)
	newwhere,err:=queryArgs.FieldMask.BuildWhereFromConfig("FindByAgeWithPage",model.StudentConditionMap)
	if err != nil {
		return nil, err
	}
	// 校验新的where条件是否使用了索引列，避免全表扫描
	check:=conditonwhere.ValidateLeadingCol(newwhere,model.StudentIndexLeadingCols)
	if !check {
		return nil, errors.New("query not use any index")
	}
	// 替换where条件
	execSql = strings.Replace(execSql, "WHERE ("+oldwhere+")", newwhere, 1)
	execCountSql := StudentNamingSqlMap[sqlName+"_Count"]
	execCountSql = strings.Replace(execCountSql, "WHERE ("+oldwhere+")", newwhere, 1)
	if execCountSql == "" {
		return nil, errors.New("not found naming sql count")
	}

	newTableName := dao.getApplyInfo(ctx).TableName
	if newTableName != "" {
		enity := &model.Student{}
		execSql = strings.ReplaceAll(execSql, "FROM "+enity.TableName()+" WHERE", "FROM "+newTableName+" WHERE")
		execCountSql = strings.ReplaceAll(execCountSql, "FROM "+enity.TableName()+" WHERE", "FROM "+newTableName+" WHERE")
	}

	argsMap := queryArgs.ConvertToMap()
	var totalCount int64
	result := dao.DB(ctx).Raw(execCountSql, argsMap).Scan(&totalCount)
	if result.Error != nil {
		return nil, result.Error
	}
	startRecord, endRecord, totalPage, enablePage, err := gormdb.CalcPageStartRecord(queryArgs.PageNum, queryArgs.PageSize, totalCount, dao.DbType)
	if err!= nil{
		return nil, err
	}
	if !enablePage {
		return &model.StudentFindByAgeWithPagePageRes{
			PageResult: gormdb.PageResult{
				CurrentPage: queryArgs.PageNum,
				PageSize:    queryArgs.PageSize,
				TotalCount:  totalCount,
				TotalPage:   totalPage,
			},
		}	, nil
	}
	argsMap["Start"] = startRecord
	argsMap["End"] = endRecord
	var list []*model.Student
	resultlist := dao.DB(ctx).Raw(execSql, argsMap).Find(&list)
	if resultlist.Error != nil {
		return nil, resultlist.Error
	}

	return &model.StudentFindByAgeWithPagePageRes{
		PageResult: gormdb.PageResult{
			CurrentPage: queryArgs.PageNum,
			PageSize:    queryArgs.PageSize,
			TotalCount:  totalCount,
			TotalPage:   totalPage,
		},
		ResultList: list,
	}, nil
}



// getInfo 获取表信息
func (dao *StudentDao) getInfo() agdao.TableInfo {
	return dao.info
}

// getApplyInfo 获取应用表信息
func (dao *StudentDao) getApplyInfo(ctx context.Context) agdao.TableInfo {
	info := dao.getInfo()
	dao.baseDao.ApplyTbInfoOpts(ctx, &info)
	return info
}

// newDB 创建一个新的DB实例
func (dao *StudentDao) newDB(ctx context.Context) (*gorm.DB, error) {
	db := dao.DB(ctx)
	info := dao.getApplyInfo(ctx)
	tbname := info.TableName
	if tbname == "" {
		return nil, errors.New("表名不能为空")
	}

	db = db.Table(tbname)
	return db, nil
}




