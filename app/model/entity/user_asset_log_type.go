package entity

type UserAssetLogType struct {
	Type uint16 `json:"type" form:"type"`
	Val  string `json:"val" form:"val"`
}
