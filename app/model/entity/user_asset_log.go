package entity

type UserAssetLog struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  Tuid int64 `json:"tuid" form:"tuid"`
  SourceId int64 `json:"source_id" form:"source_id"`
  SourceTable string `json:"source_table" form:"source_table"`
  Op int64 `json:"op" form:"op"`
  Type int64 `json:"type" form:"type"`
  AssetField string `json:"asset_field" form:"asset_field"`
  BeforeVal float64 `json:"before_val" form:"before_val"`
  Money float64 `json:"money" form:"money"`
  Remark string `json:"remark" form:"remark"`
}