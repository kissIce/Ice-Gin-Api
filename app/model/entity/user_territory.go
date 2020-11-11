package entity

type UserTerritory struct {
 
  Id int32 `json:"id" form:"id"`
  Uid int32 `json:"uid" form:"uid"`
  TopId int32 `json:"top_id" form:"top_id"`
  TopName string `json:"top_name" form:"top_name"`
  TerritoryId int32 `json:"territory_id" form:"territory_id"`
  SecondName string `json:"second_name" form:"second_name"`
  CreatedAt int32 `json:"created_at" form:"created_at"`
}