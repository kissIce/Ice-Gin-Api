package admin

import (
	"ice/app/model/data"
	"ice/app/model/entity"
	adminRequest "ice/app/model/request/admin"
	"ice/utils/response"
)

func GetApiList() *response.Resp {
	rsp, _ := data.GetApiList()
	return response.InitSucc(rsp)
}

func AddApi(r *adminRequest.AddApi) *response.Resp {
	api := &entity.Api{
		Path: r.Path,
		Method: r.Method,
		Group: r.Group,
	}
	err := data.AddApi(api)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(api)
}

func EditApi(r *adminRequest.EditApi) *response.Resp {
	api := &entity.Api{
		Id:   r.Id,
		Path: r.Path,
		Method: r.Method,
		Group: r.Group,
	}
	err := data.EditApi(api)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(api)
}

func DelApi(r *adminRequest.IdReq) *response.Resp {
	api := &entity.Api{
		Id:   r.Id,
	}
	err := data.DelApi(api)
	if err != nil {
		return response.InitErrCode(response.DbError)
	}
	return response.InitSucc(api)
}