package response

var RetMsg = map[int]map[string]string{
	ApiSuccess: {
		"zh" : "请求成功",
		"en" : "Request Success",
	},
	ApiError: {
		"zh" : "请求失败",
		"en" : "Request Fail",
	},

	TokenMiss: {
		"zh" : "token缺失",
		"en" : "Token Lose",
	},

	SignError: {
		"zh" : "token验证失败",
		"en" : "Token Error",
	},

	SignBefore: {
		"zh" : "token未到可用时间",
		"en" : "Token Before Use Time",
	},

	SignExpire: {
		"zh" : "token已过期",
		"en" : "Token Expire",
	},

	SignOther: {
		"zh" : "token错误",
		"en" : "Token Error",
	},

	SignRefresh: {
		"zh" : "token刷新",
		"en" : "Token Refresh",
	},

	PermissionsDenied: {
		"zh" : "权限不足",
		"en" : "Permissions Denied",
	},

	RouteNofound: {
		"zh" : "请求地址不存在",
		"en" : "Address Not Exists",
	},

	MethodNoAllow: {
		"zh" : "请求方法不允许",
		"en" : "Request Method Now Allow",
	},

	DbError: {
		"zh" : "数据库错误",
		"en" : "DataBase Error",
	},

	ValidateError: {
		"zh" : "参数错误",
		"en" : "Param Error",
	},

	CaptchaError: {
		"zh" : "验证码生成失败",
		"en" : "Captcha Make Fail",
	},

	CaptchaVerifyFail: {
		"zh" : "验证码错误",
		"en" : "Captcha Verify Fail",
	},

	PayError: {
		"zh" : "支付错误",
		"en" : "Pay Error",
	},

	ThirdError: {
		"zh" : "第三方错误",
		"en" : "Third Error",
	},
}
