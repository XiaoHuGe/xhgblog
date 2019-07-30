package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	SUCCESS_REGISTER = 1001
	SUCCESS_LOGIN    = 1002
	SUCCESS_LOGOUT   = 1003
	SUCCESS_GETME    = 1004

	ERROR_NOT_LOGIN       = 1101
	ERROR_NOT_ENAIL       = 1107
	ERROR_ENAIL_OR_PASS   = 1102
	ERROR_PASSWORD_DIFFER = 1103
	ERROR_ENCRYPT         = 1104
	ERROR_CREATE_SQL      = 1105
	ERROR_EXIST_EMAIL     = 1106

	ERROR_EXIST_TAG       = 2001
	ERROR_EXIST_TAG_FAIL  = 2002
	ERROR_NOT_EXIST_TAG   = 2003
	ERROR_GET_TAGS_FAIL   = 2004
	ERROR_COUNT_TAG_FAIL  = 2005
	ERROR_ADD_TAG_FAIL    = 2006
	ERROR_EDIT_TAG_FAIL   = 2007
	ERROR_DELETE_TAG_FAIL = 2008
)

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	SUCCESS_REGISTER: "注册成功",
	SUCCESS_LOGIN:    "登录成功",
	SUCCESS_LOGOUT:   "退出成功",
	SUCCESS_GETME:    "成功获取个人信息",

	ERROR_NOT_LOGIN:       "用户未登录",
	ERROR_NOT_ENAIL:       "用户不存在",
	ERROR_ENAIL_OR_PASS:   "邮箱或密码错误",
	ERROR_PASSWORD_DIFFER: "上下密码输入不一致",
	ERROR_ENCRYPT:         "密码加密失败",
	ERROR_CREATE_SQL:      "数据库创建失败",
	ERROR_EXIST_EMAIL:     "已存在该邮箱",

	ERROR_EXIST_TAG:       "已存在该标签名称",
	ERROR_EXIST_TAG_FAIL:  "获取已存在标签失败",
	ERROR_NOT_EXIST_TAG:   "该标签不存在",
	ERROR_GET_TAGS_FAIL:   "获取所有标签失败",
	ERROR_COUNT_TAG_FAIL:  "统计标签失败",
	ERROR_ADD_TAG_FAIL:    "新增标签失败",
	ERROR_EDIT_TAG_FAIL:   "修改标签失败",
	ERROR_DELETE_TAG_FAIL: "删除标签失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
