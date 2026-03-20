package headdev

// import "context"

// type agHeadKey struct{}

// var HeadKeys = []string{
// 	"ORG",
// 	"TraceId",
// }

// type MD map[string]string

// func AppendMdToContext(ctx context.Context, newmd MD) context.Context {
// 	// 从上下文提取已存在头信息
// 	var md MD
// 	aghead := ctx.Value(agHeadKey{})
// 	if aghead != nil {
// 		md, _ = aghead.(MD) // 是否可能类型失败
// 	}

// 	// 将新的头信息合并到已存在头信息中
// 	if md == nil {
// 		md = make(MD, len(newmd))
// 	}
// 	for k, v := range newmd {
// 		md[k] = v
// 	}

// 	// 将合并后的头信息设置到上下文
// 	rctx := context.WithValue(ctx, agHeadKey{}, md)
// 	return rctx
// }

// func GetMdFromContext(ctx context.Context) MD {
// 	aghead := ctx.Value(agHeadKey{})
// 	if aghead == nil {
// 		return nil
// 	}
// 	md, _ := aghead.(MD)
// 	return md
// }
