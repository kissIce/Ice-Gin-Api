package admin

type AddApi struct {
	Group  string `json:"group" form:"group"`
	Path   string `json:"path" form:"path"`
	Method string `form:"method" form:"method"`
}

type EditApi struct {
	IdReq
	AddApi
}
