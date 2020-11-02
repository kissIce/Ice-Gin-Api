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

	TokenCreateFail: {
		"zh" : "token签发失败",
		"en" : "Token Create Fail",
	},

	TokenBefore: {
		"zh" : "token未到可用时间",
		"en" : "Token Before Use Time",
	},

	TokenExpire: {
		"zh" : "token已过期",
		"en" : "Token Expire",
	},


	TokenRefresh: {
		"zh" : "token刷新",
		"en" : "Token Refresh",
	},

	TokenBlack: {
		"zh" : "token无效",
		"en" : "Token Useless",
	},

	TokenMalformed: {
		"zh" : "token格式错误",
		"en" : "Token Format Error",
	},

	TokenInvalid: {
		"zh" : "token验证失败",
		"en" : "Token Invalid",
	},

	TokenOther: {
		"zh" : "token其他错误",
		"en" : "Token Other Error",
	},

	PermissionDenied: {
		"zh" : "权限不足",
		"en" : "Permission Denied",
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
