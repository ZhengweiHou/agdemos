package internal

import (
	"ag-core/ag/ag_common/agmetadata"
)

func init() {
	// 初始化headdev
	agmetadata.RegMdKey("ORG")
	agmetadata.RegMdKey("Tradeid")

	// 注册事务
	// agdb.TransactionRegistry.RegTxForCall(svcgen.StudentServiceCreateStudentCallInfo)
	// svcgen.StudentServiceCreateStudentCallInfo.AddTag(agdb.TransactionTag, true)
}
