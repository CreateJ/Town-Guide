package service

import (
	"github.com/kataras/iris/v12"
	"town-guide/model"
)

type ScenicServiceApi struct {
}

func NewScenicService() *ScenicServiceApi {
	return &ScenicServiceApi{}
}


func (u *ScenicServiceApi) GetScenicInfoByID(ctx iris.Context) {
	dto := GetOpenIDDTO{}
	ctx.ReadQuery(&dto)
	if dto.Code == "" {
		ctx.JSON(Response{ErrorCode, "参数错误", nil})
	}

	info := model.GetUserInfo(dto.Code)
	if info == nil {
		ctx.JSON(Response{ErrorCode, "获取不到openID,请检查参数是否有效", nil})
		return
	}

	ctx.JSON(Response{SuccessCode, "", info})
}
