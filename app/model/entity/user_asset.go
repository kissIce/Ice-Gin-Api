package entity

type UserAsset struct {
	Uid          uint64  `json:"uid" form:"uid"`
	Amount       float64 `json:"amount" form:"amount"`
	FronzeAmount float64 `json:"fronze_amount" form:"fronze_amount"`
	Beans        float64 `json:"beans" form:"beans"`
	FronzeBeans  float64 `json:"fronze_beans" form:"fronze_beans"`
}
