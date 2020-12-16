package admin

type AddAdCate struct {
	Name      string `json:"name" form:"name"`
	Width     uint16 `json:"width" form:"width"`
	Height    uint16 `json:"height" form:"height"`
	Desc      string `json:"desc" form:"desc"`
	Status    int8   `json:"status" form:"status"`
	AllowType string `json:"allow_type" form:"allow_type"`
}

type EditAdCate struct {
	IdReq
	AddAdCate
}
