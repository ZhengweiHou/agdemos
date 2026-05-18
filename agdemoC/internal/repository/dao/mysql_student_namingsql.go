package dao

// DO NOT EDIT
// DO NOT EDIT
// DO NOT EDIT

const MYSQL_Student_FindByAge = "SELECT ID,STUNO,NAME,AGE FROM STUDENT WHERE (AGE = @Age AND STUNO > @Stuno)"

const MYSQL_Student_FindByAgeWithPage = "SELECT * FROM STUDENT WHERE (AGE = @Age) LIMIT @Start, @End"

const MYSQL_Student_FindByAgeWithPage_Count = "SELECT COUNT(*) FROM STUDENT WHERE (AGE = @Age)"

func InitStudentMYSQL() {
	StudentNamingSqlMap["MYSQL_Student_FindByAge"] = MYSQL_Student_FindByAge
	StudentNamingSqlMap["MYSQL_Student_FindByAgeWithPage"] = MYSQL_Student_FindByAgeWithPage
	StudentNamingSqlMap["MYSQL_Student_FindByAgeWithPage_Count"] = MYSQL_Student_FindByAgeWithPage_Count
}
