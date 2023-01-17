package example

//	@Name   week
//
// @Description  星期相关
const (
	Monday    = iota // 周一
	Tuesday          // 周二
	Wednesday        // 周三
	Thursday         // 周四
	Friday           // 周五
	Saturday         // 周六
	Sunday           // 周日
)

var week = map[int]string{
	Monday:    "周一",
	Tuesday:   "周二",
	Wednesday: "周三",
	Thursday:  "周四",
	Friday:    "周五",
	Saturday:  "周六",
	Sunday:    "周日",
}
