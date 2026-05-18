package dao

// DO NOT EDIT
// DO NOT EDIT
// DO NOT EDIT

const DB2_Student_FindByAge = "SELECT ID,STUNO,NAME,AGE FROM STUDENT WHERE (AGE = @Age AND STUNO > @Stuno)"

const DB2_Student_FindByAgeWithPage = "SELECT * FROM (SELECT *, ROW_NUMBER() OVER(ORDER BY ID) AS RN  FROM STUDENT WHERE (AGE = @Age)) AS T WHERE RN BETWEEN @Start AND @End"

const DB2_Student_FindByAgeWithPage_Count = "SELECT COUNT(*) FROM STUDENT WHERE (AGE = @Age)"

func InitStudentDB2() {
	StudentNamingSqlMap["DB2_Student_FindByAge"] = DB2_Student_FindByAge
	StudentNamingSqlMap["DB2_Student_FindByAgeWithPage"] = DB2_Student_FindByAgeWithPage
	StudentNamingSqlMap["DB2_Student_FindByAgeWithPage_Count"] = DB2_Student_FindByAgeWithPage_Count
}
