package admin

type AddRole struct {
	Pid uint64 `json:"pid" form:"pid"`
	Name string `json:"name" form:"name"`
}

type EditRole struct {
	IdReq
	Name string `json:"name" form:"name"`
}