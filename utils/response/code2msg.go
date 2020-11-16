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

	RouteNotfound: {
		"zh" : "请求地址不存在",
		"en" : "Address Not Exists",
	},

	MethodNotAllow: {
		"zh" : "请求方法不允许",
		"en" : "Request Method Now Allow",
	},

	DbError: {
		"zh" : "数据库错误",
		"en" : "DataBase Error",
	},

	ParamSignMiss: {
		"zh" : "参数签名缺失",
		"en" : "param sign miss",
	},

	ParamSignLenErr: {
		"zh" : "参数签名长度错误",
		"en" : "param sign len error",
	},

	ParamSignNotEq: {
		"zh" : "参数签名错误",
		"en" : "param sign Error",
	},

	ParamTimeMiss: {
		"zh" : "参数timestamp缺失",
		"en" : "param timestamp miss",
	},

	ParamTimeOutUse: {
		"zh" : "参数timestamp超出安全阈值",
		"en" : "param timestamp out of use",
	},

	ParamNonceStrMiss: {
		"zh" : "参数nonce_str缺失",
		"en" : "param nonce_str miss",
	},

	ParamNonceStrLenErr: {
		"zh" : "参数nonce_str长度错误",
		"en" : "param nonce_str len error",
	},

	ParamNonceStrRepeat: {
		"zh" : "请求重复",
		"en" : "request repeat",
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
		"zh" : "图文验证码错误",
		"en" : "Captcha Verify Fail",
	},

	SmsVerifyFail: {
		"zh" : "短信验证码错误",
		"en" : "Sms Verify Fail",
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
