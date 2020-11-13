package entity

type User struct {
	Id        uint64 `json:"id" form:"id" gorm:"primarykey"`
	Username  string `json:"username" form:"username"`
	TzoneImg  string `json:"tzone_img" form:"tzone_img"`
	Avatar    string `json:"avatar" form:"avatar"`
	Sex       int8   `json:"sex" form:"sex"`
	Age       int8   `json:"age" form:"age"`
	Phone     string `json:"phone" form:"phone"`
	AreaId    int32  `json:"area_id" form:"area_id"`
	Area      string `json:"area" form:"area"`
	Signature string `json:"signature" form:"signature"`
	ImAcc     string `json:"im_acc" form:"im_acc"`
	ImPwd     string `json:"im_pwd" form:"im_pwd"`
	Identity  int16  `json:"identity" form:"identity"`
	Vip       int8   `json:"vip" form:"vip"`
	RecomId   int64  `json:"recom_id" form:"recom_id"`
	Status    int8   `json:"status" form:"status"`
	Channel   int32  `json:"channel" form:"channel"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
