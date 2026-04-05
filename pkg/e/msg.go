package e

var Msg = map[int]string{
	Success:       "success",
	Error:         "failed",
	InvalidParams: "参数错误",

	ErrorExist:          "已存在",
	ErrorNotFound:       "用户不存在",
	ErrorFailEncryption: "加密密码失败",
	ErrorFailCreat:      "创建用户失败",

	WrongPassword:         "密码错误",
	ErrorAuthToken:        "token认证失败",
	ErrorAuthTokenTimeOut: "token失效",
	ErrorAuthTokenEmpty:   "缺失token",
	ErrorUploadFail:       "图片上传失败",
	ErrorSendEmailFail:    "邮件发送失败",
}

func GetMsg(code int) string {
	msg, ok := Msg[code]
	if !ok {
		return Msg[Error]
	}
	return msg
}
