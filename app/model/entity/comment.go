package entity

type Comment struct {
	Id        uint64 `json:"id" form:"id" gorm:"primarykey"`
	Uid       uint64  `json:"uid" form:"uid"`
	Content   string `json:"content" form:"content"`
	Sid       uint64  `json:"sid" form:"sid"`
	Suid      uint64  `json:"suid" form:"suid"`
	Scontent  string `json:"scontent" form:"scontent"`
	Simg      string `json:"simg" form:"simg"`
	Stable    string `json:"stable" form:"stable"`
	Tcid      uint64  `json:"tcid" form:"tcid"`
	Tcuid     uint64  `json:"tcuid" form:"tcuid"`
	Tccontent string `json:"tccontent" form:"tccontent"`
	Cid       uint64  `json:"cid" form:"cid"`
	Cuid      uint64  `json:"cuid" form:"cuid"`
	Ccontent  string `json:"ccontent" form:"ccontent"`
	IsRead    uint8   `json:"is_read" form:"is_read"`
	Support   uint32  `json:"support" form:"support"`
	Comment   uint32  `json:"comment" form:"comment"`
	CreatedAt int32  `json:"created_at" form:"created_at"`
	UpdatedAt int32  `json:"updated_at" form:"updated_at"`
	DeletedAt int32  `json:"deleted_at" form:"deleted_at"`
}
