
# 通过表格生成yaml IDL描述
gen-go-db yaml -i ./daoidl/student.xlsx

# 根据yaml IDL描述生成数据库BaseDao
# gen-go-db db -i ./repository/yaml/student.yaml -p agdemoB/internal -d mysql
gen-go-db db -i ./internal/repository/yaml/student.yaml -o ./internal -p agdemoB/internal -d mysql

gen-go-db db -i /home/houzw/document/git-rep/HOUZW/golang/aic-gospring/AG_WS/agdemoB/internal/repository/yaml/student.yaml -o /home/houzw/document/git-rep/HOUZW/golang/aic-gospring/AG_WS/agdemoB/internal -p agdemoB/internal -d mysql



/home/houzw/document/git-rep/HOUZW/golang/aic-gospring/AG_WS/agdemoB