// DO NOT EDIT
// DO NOT EDIT
// DO NOT EDIT

package dao
const DB2_Student_FindByAge = "SELECT ID,STUNO,NAME,AGE FROM student WHERE (AGE = @Age AND STUNO > @Stuno)  "

const DB2_Student_FindByWithPage = "SELECT ID,STUNO,NAME,AGE,VERSION,FCT,LCT FROM student WHERE AGE = @Age  "


func InitStudentDB2() {
    StudentNamingSqlMap["DB2_Student_FindByAge"]  = DB2_Student_FindByAge
    
    StudentNamingSqlMap["DB2_Student_FindByWithPage"]  = DB2_Student_FindByWithPage
              
}