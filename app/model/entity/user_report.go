package entity

type UserReport struct {
	Id        uint64 `json:"id" form:"id" gorm:"primarykey"`
	Uid       uint64 `json:"uid" form:"uid"`
	Sid       uint64 `json:"sid" form:"sid"`
	Suid      uint64 `json:"suid" form:"suid"`
	Stable    string `json:"stable" form:"stable"`
	Scontent  string `json:"scontent" form:"scontent"`
	Simg      string `json:"simg" form:"simg"`
	Svideo    string `json:"svideo" form:"svideo"`
	Sdomain   string `json:"sdomain" form:"sdomain"`
	Domain    string `json:"domain" form:"domain"`
	Img       string `json:"img" form:"img"`
	Content   string `json:"content" form:"content"`
	Status    int8   `json:"status" form:"status"`
	Do        int8   `json:"do" form:"do"`
	CreatedAt uint32 `json:"created_at" form:"created_at" gorm:"autoCreateTime;default:0;comment:添加时间"`
	UpdatedAt uint32 `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime;default:0;comment:添加时间"`
	DeletedAt uint32 `json:"deleted_at" form:"deleted_at" gorm:"default:0;comment:删除时间"`
}
