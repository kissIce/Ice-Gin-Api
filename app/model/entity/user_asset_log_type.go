package entity

type UserAssetLogType struct {
  Model
  Type int64 `json:"type" form:"type"`
  Val string `json:"val" form:"val"`
}