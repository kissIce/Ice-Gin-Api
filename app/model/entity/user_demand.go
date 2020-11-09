package entity

type UserDemand struct {
  Model
  Uid int64 `json:"uid" form:"uid"`
  Tid string `json:"tid" form:"tid"`
  WorkYear int64 `json:"work_year" form:"work_year"`
  Area int64 `json:"area" form:"area"`
  Budget int64 `json:"budget" form:"budget"`
  Demand string `json:"demand" form:"demand"`
  Status int64 `json:"status" form:"status"`
}