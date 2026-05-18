
```bash
# 拆分旧版本idl表格为新版本
gen-go-db sheet -i ./idl/student.xlsx -o ./idl
# 或
gen-go-db sheet -i ./idl/student.xlsx -o ./idl/student_v2.xlsx

# 生成yaml描述文件
gen-go-db yaml --input ./idl/student_v2.xlsx --output ../
## 会在../repository/yaml目录下生成[TableName].yaml文件

# 生成数据库代码
gen-go-db db --input ./yaml/STUDENT.yaml --output ../  -m agdemoC/internal 
## ../repository/[dao|model] 目录下生成相关数据库操作代码和model
```





