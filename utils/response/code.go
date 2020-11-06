package response

const (
	// 全局请求
	ApiSuccess = 10000
	ApiError   = 99999

	// token
	TokenMiss        = 10001
	TokenCreateFail  = 10002
	TokenBefore      = 10003
	TokenExpire      = 10004
	TokenRefresh     = 10005
	TokenBlack       = 10006
	TokenMalformed   = 10007
	TokenInvalid     = 10008
	TokenOther       = 10009
	PermissionDenied = 10010

	// ROUTER
	RouteNotfound = 10020
	MethodNotAllow = 10021

	// DB
	DbError = 10030

	// VALIDATE
	ValidateError = 10050

	CaptchaError = 10060
	CaptchaVerifyFail = 10061
	SmsVerifyFail = 10062
	// PAY
	PayError = 10070

	// THIRD
	ThirdError = 10090
)
