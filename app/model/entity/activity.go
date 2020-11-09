package entity

type Activity struct {
	Model
	Title         string `json:"title" form:"title"`
	StartTime     int64  `json:"start_time" form:"start_time"`
	EndTime       int64  `json:"end_time" form:"end_time"`
	PagePath      string `json:"page_path" form:"page_path"`
	ShareTitle    string `json:"share_title" form:"share_title"`
	ShareSubtitle string `json:"share_subtitle" form:"share_subtitle"`
	ShareImg      string `json:"share_img" form:"share_img"`
	ShareUrl      string `json:"share_url" form:"share_url"`
	Status        int64  `json:"status" form:"status"`
	AdminId       int64  `json:"admin_id" form:"admin_id"`
}
