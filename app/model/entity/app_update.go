package entity

type AppUpdate struct {
  Model
  Type string `json:"type" form:"type"`
  Version string `json:"version" form:"version"`
  Title string `json:"title" form:"title"`
  Info string `json:"info" form:"info"`
  Force int8 `json:"force" form:"force"`
  DownloadUrl string `json:"download_url" form:"download_url"`
  Status int8 `json:"status" form:"status"`
}