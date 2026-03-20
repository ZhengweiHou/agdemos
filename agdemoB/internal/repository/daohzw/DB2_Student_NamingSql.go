// DO NOT EDIT
// DO NOT EDIT
// DO NOT EDIT

package daohzw

const DB2_Student_FindByAge = "SELECT ID,STUNO,NAME,AGE FROM student WHERE (AGE = @Age AND STUNO > @Stuno)  "

const DB2_Student_FindByWithPage = "SELECT VERSION,FCT,LCT,ID,STUNO,NAME,AGE FROM student WHERE AGE = @Age ORDER BY id limit @StartNum,@EndNum"

const DB2_Student_FindByWithPage_Count = "SELECT COUNT(*) FROM student WHERE AGE = @Age"

func InitStudentDB2() {
	// StudentNamingSqlMap["DB2_Student_FindByAge"]  = DB2_Student_FindByAge

	// StudentNamingSqlMap["DB2_Student_FindByWithPage"]  = DB2_Student_FindByWithPage

	// StudentNamingSqlMap["DB2_Student_FindByWithPage_Count"]  = DB2_Student_FindByWithPage_Count
}
