package entity

type UserTerritory struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  TopId int64 `json:"top_id" form:"top_id"`
  TopName string `json:"top_name" form:"top_name"`
  TerritoryId int64 `json:"territory_id" form:"territory_id"`
  SecondName string `json:"second_name" form:"second_name"`
}