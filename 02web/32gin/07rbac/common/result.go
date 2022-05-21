package result

type Result struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (errcode *Result) Error() string {
	return errcode.Msg
}

func NewResult(Code int32, Msg string, Data interface{}) *Result {

	return &Result{Code, Msg, Data}
}

var (
	SUCCESS             = NewResult(2000, "操作成功", nil)
	FORMATE_ERROR       = NewResult(4001, "格式转换异常", nil)
	USER_IS_EXISTED     = NewResult(4002, "用户已存在", nil)
	PASSWORD_ERROR      = NewResult(4003, "密码错误", nil)
	UNAUTHORIZED        = NewResult(4004, "未授权", nil)
	NO_PERMISSION       = NewResult(4005, "无权限", nil)
	NOT_FOUND_TOKEN     = NewResult(4006, "未找到TOKEN", nil)
	TOKEN_FORMATE_ERROR = NewResult(4007, "token格式有误", nil)
	INVALID_TOKEN       = NewResult(4008, "不合法的token", nil)

	UNKNOW_ERROR = NewResult(5000, "未知异常", nil)
)
