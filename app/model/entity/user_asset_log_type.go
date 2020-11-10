package entity

type UserAssetLogType struct {
  Model
  Type int16 `json:"type" form:"type"`
  Val string `json:"val" form:"val"`
}