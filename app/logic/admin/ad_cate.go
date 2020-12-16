package admin

import (
	"ice/app/model/data"
	"ice/app/model/entity"
	adminRequest "ice/app/model/request/admin"
	"ice/utils/response"
)

func GetAdCateList() *response.Resp {
	rsp, _ := data.GetAdCateList()
	return response.InitSucc(rsp)
}

func AddAdCate(r *adminRequest.AddAdCate) *response.Resp {
	adCate := &entity.AdCate{
		Name: r.Name,
		Width: r.Width,
		Height: r.Height,
		Desc: r.Desc,
		Status: r.Status,
		AllowType: r.AllowType,
	}
	err := data.AddAdCate(adCate)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(adCate)
}

func EditAdCate(r *adminRequest.EditAdCate) *response.Resp {
	adCate := &entity.AdCate{
		Id:   r.Id,
		Name: r.Name,
		Width: r.Width,
		Height: r.Height,
		Desc: r.Desc,
		Status: r.Status,
		AllowType: r.AllowType,
	}
	err := data.EditAdCate(adCate)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(adCate)
}

func DelAdCate(r *adminRequest.IdReq) *response.Resp {
	api := &entity.AdCate{
		Id:   r.Id,
	}
	err := data.DelAdCate(api)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(api)
}
