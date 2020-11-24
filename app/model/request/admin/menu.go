package admin

type AddMenu struct {
	Pid       uint64 `json:"pid" form:"pid"`
	Title     string `json:"title" form:"title"`
	Name      string `json:"name" form:"name"`
	Icon      string `json:"icon" form:"icon"`
	Path      string `json:"path" form:"path"`
	Component string `json:"component" form:"component"`
	KeepAlive int8   `json:"keep_alive" form:"keep_alive"`
	Hidden    int8   `json:"hidden" form:"hidden"`
	Sort      int16  `json:"sort" form:"sort"`
	Default   int8   `json:"default" form:"default"`
}

type EditMenu struct {
	IdReq
	AddMenu
}

type RoleMenu struct {
	RoleId uint64 `json:"role_id" form:"role_id"`
}