package bv1

// import (
// 	"fmt"
// 	"testing"
// )

// type HArg struct{}
// type HResult struct{}
// type HArg2 struct{}
// type HResult2 struct{}

// type NamingInfo[R, T any] struct {
// 	Name  string
// 	AType R
// 	RType T
// }

// var HzwNaming1 = NamingInfo[HArg, HResult]{
// 	Name:  "HzwNaming1",
// 	AType: HArg{},
// 	RType: HResult{},
// }

// var HzwNaming2 = NamingInfo[HArg2, HResult2]{
// 	Name:  "HzwNaming2",
// 	AType: HArg2{},
// 	RType: HResult2{},
// }

// type HzwDao struct{}

// func (h *HzwDao) NamingSql[R, T any](info NamingInfo[R, T]) []*T {
//     switch info.Name {
//     case HzwNaming1.Name:
//         // ...
//         return nil
//     case HzwNaming2.Name:
//         // ...
//         return nil
//     default:
//         return nil
//     }
// }

// type NamingSqlQuery[R, T any] struct {
//     dao  *HzwDao
//     info NamingInfo[R, T]
// }

// func (q *NamingSqlQuery[R, T]) Execute() []*T {
//     // 实现逻辑
//     return nil
// }

// func (h *HzwDao) NewNamingSqlQuery[R, T any](info NamingInfo[R, T]) *NamingSqlQuery[R, T] {
//     return &NamingSqlQuery[R, T]{dao: h, info: info}
// }

// func TestHzeDao_NamingSql(t *testing.T) {
// 	hzwDao := &HzwDao{}
// 	res := hzwDao.NamingSql(HzwNaming1)
// 	fmt.Printf("res type: %T\n", res)

// 	res2 := hzwDao.NamingSql(HzwNaming2)
// 	fmt.Printf("res2 type: %T\n", res2)
// }
