package service

import (
	"github.com/kataras/iris/v12"
	"town-guide/model"
)

const SuccessCode = 2
const ErrorCode = 1

type UserServiceApi struct {
}

func NewUserService() *UserServiceApi {
	return &UserServiceApi{}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (u *UserServiceApi) GetUserInfo(ctx iris.Context) {
	dto := GetOpenIDDTO{}
	_ = ctx.ReadQuery(&dto)
	if dto.Code == "" {
		_, _ = ctx.JSON(Response{ErrorCode, "参数错误", nil})
	}

	info := model.GetUserInfo(dto.Code)
	if info == nil {
		_, _ = ctx.JSON(Response{ErrorCode, "获取不到openID,请检查参数是否有效", nil})
		return
	}

	_, _ = ctx.JSON(Response{SuccessCode, "", info})
}

type GetOpenIDDTO struct {
	Code string `json:"code"`
}

func (u *UserServiceApi) Register(ctx iris.Context) {
	userInfoDTO := model.UserInfoDTO{}
	_ = ctx.ReadJSON(&userInfoDTO)
	if userInfoDTO.NickName == "" || userInfoDTO.Avatar == "" || userInfoDTO.OpenID == "" {
		_, _ = ctx.JSON(Response{ErrorCode, "参数错误", nil})
		return
	}

	model.Register(userInfoDTO)
	_, _ = ctx.JSON(Response{SuccessCode, "", nil})
}
