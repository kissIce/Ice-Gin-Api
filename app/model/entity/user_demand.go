package entity

type UserDemand struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  Tid string `json:"tid" form:"tid"`
  WorkYear int8 `json:"work_year" form:"work_year"`
  Area int32 `json:"area" form:"area"`
  Budget int32 `json:"budget" form:"budget"`
  Demand string `json:"demand" form:"demand"`
  Status int8 `json:"status" form:"status"`
}