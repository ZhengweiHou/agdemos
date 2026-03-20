// DO NOT EDIT
// DO NOT EDIT
// DO NOT EDIT

package dao
const MYSQL_Student_FindByAge = "SELECT ID,STUNO,NAME,AGE FROM student WHERE (AGE = @Age AND STUNO > @Stuno)  "

const MYSQL_Student_FindByWithPage = "SELECT ID,STUNO,NAME,AGE,VERSION,FCT,LCT FROM student WHERE AGE = @Age  "


func InitStudentMYSQL() {
    StudentNamingSqlMap["MYSQL_Student_FindByAge"]  = MYSQL_Student_FindByAge
    
    StudentNamingSqlMap["MYSQL_Student_FindByWithPage"]  = MYSQL_Student_FindByWithPage
              
}