package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ice/app/model/data"
	"ice/app/model/entity"
	apiResponse "ice/app/model/response/api"
	"ice/app/service"
	"ice/helper"
	"ice/utils/response"
	"strconv"
	"time"
)

var err error

func LoginByPhone(u *entity.User) *response.Resp {
	u, err = data.GetUserByPhone(u.Phone, []string{"phone", "id", "username", "avatar", "im_acc", "im_pwd", "deleted_at", "status"})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return response.InitErrCode(response.DbError)
		}
		u.Username = "用户" + strconv.FormatInt(helper.RandRangeInt(1000,9999), 10) + strconv.FormatInt(time.Now().Unix(), 10)
		if err = data.CreateUser(u); err != nil {
			return response.InitErrCode(response.DbError)
		}
		// 注册im账号
		go func() {
			bool, _ := service.ImReg(&service.ImInfo{
				Identifier: strconv.FormatUint(u.Id, 10) + "_" + strconv.FormatInt(time.Now().Unix(), 10),
				Nick:       u.Username,
				FaceUrl:    u.Avatar,
			})
			if bool {

			}
		}()
	}
	if u.DeletedAt > 0 || u.Status > 0 {
		return response.InitErrMsg("用户禁止登陆")
	}
	// 创建JWT
	var jwt service.JWT
	rsp := &apiResponse.LoginUserByPhone{
		Uid:      u.Id,
		Phone:    u.Phone,
		Avatar:   u.Avatar,
		Username: u.Username,
		ImAcc:    u.ImAcc,
		ImPwd:    u.ImPwd,
	}
	rsp.Token, _ = jwt.CreateToken(service.InitClaims(gin.H{"id": u.Id}, 3*24*3600))
	return response.InitSucc(rsp)
}
