package entity

type UserVip struct {
	Uid     uint32 `json:"uid" form:"uid"`
	Level   uint8  `json:"level" form:"level"`
	EndTime string `json:"end_time" form:"end_time"`
}
