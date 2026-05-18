package internal

import (
	"ag-core/ag/ag_common/agmetadata"
	"ag-core/contribute/agdb"
	"agdemoB/internal/svcgen"
)

func init() {
	// 初始化headdev
	agmetadata.RegMdKey("ORG")
	agmetadata.RegMdKey("Tradeid")
	agmetadata.RegMdKey("ORG")

	// 注册事务
	// agdb.TransactionRegistry.RegTxForCall(svcgen.DemobCalldemobCallInfo)
	// tagCallInfo(svcgen.DemobCalldemobCallInfo, "transaction", true) // 等价于 spring  @Transactional
	// svcgen.DemobCalldemobCallInfo.AddTag("transaction", true)
	// svcgen.DemobCalldemobCallInfo.AddTag("transaction", 6)
	svcgen.DemobCalldemobCallInfo.AddTag(agdb.TransactionTag, agdb.TRANSACTION_PROPAGATION_REQUIRED)
}
