package e

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",

	ERROR_AUTH_NOT_LOGIN: "尚未登陆",
	ERROR_AUTH_STID_PSW: "用户名或密码错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL : "无权访问",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "登录超时",
	ERROR_AUTH_OTHER_ERROR: "未知错误",

	ERROR_NOT_EXIST_USER: "该用户不存在",
	ERROR_HAS_EXIST_USER: "已存在相同学号/工号成员",

	ERROR_UPLOAD_IMAGE_TOO_LARGE: "图片大小超过 5 MB",
	ERROR_UPLOAD_IMAGE_WRONG_TYPE: "仅允许 .jpg .jpeg .png 格式图片",
	ERROR_UPLOAD_IMAGE_FAIL: "上传图片失败",

	ERROR_BULLETIN_NOT_EXIST: "该公告不存在",

	ERROR_NEWS_NOT_EXIST: "该公告不存在",

	ERROR_CONTEST_NOT_EXIST: "该比赛不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}