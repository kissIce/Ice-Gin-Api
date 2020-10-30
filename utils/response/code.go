package response

const (
	// 全局请求
	ApiSuccess = 10000
	ApiError   = 99999

	// token
	TokenMiss         = 10001
	SignError         = 10002
	SignBefore        = 10003
	SignExpire        = 10004
	SignOther         = 10005
	SignRefresh       = 10006
	PermissionsDenied = 10010

	// ROUTER
	RouteNofound  = 10020
	MethodNoAllow = 10021

	// DB
	DbError = 10030

	// VALIDATE
	ValidateError = 10040

	CaptchaError = 10050

	CaptchaVerifyFail = 10051

	// PAY
	PayError = 10070

	// THIRD
	ThirdError = 10090
)
