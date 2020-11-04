package entity

type Api struct {
	Model
	Group  string `json:"group" gorm:"type:varchar(50);default:'';comment:api组;"`
	Path   string `json:"path" gorm:"type:varchar(150);default:'';comment:路径"`
	Method string `json:"method" gorm:"type:varchar(30);default:'';comment:方法"`
}
