// DO NOT EDIT
// DO NOT EDIT
// DO NOT EDIT

package daohzw

const MYSQL_Student_FindByAge = "SELECT ID,STUNO,NAME,AGE FROM student WHERE (AGE = @Age AND STUNO > @Stuno)  "

const MYSQL_Student_FindByWithPage = "SELECT VERSION,FCT,LCT,ID,STUNO,NAME,AGE FROM student WHERE AGE = @Age ORDER BY id limit @StartNum,@EndNum"

const MYSQL_Student_FindByWithPage_Count = "SELECT COUNT(*) FROM student WHERE AGE = @Age"

func InitStudentMYSQL() {
	// StudentNamingSqlMap["MYSQL_Student_FindByAge"]  = MYSQL_Student_FindByAge

	// StudentNamingSqlMap["MYSQL_Student_FindByWithPage"]  = MYSQL_Student_FindByWithPage

	// StudentNamingSqlMap["MYSQL_Student_FindByWithPage_Count"]  = MYSQL_Student_FindByWithPage_Count
}
