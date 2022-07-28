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
	ctx.JSON(Response{SuccessCode, "", nil})
}

type GetOpenIDDTO struct {
	Code string `json:"code"`
}

//func (u *UserServiceApi) GetUserOpenID(ctx iris.Context) {
//	dto := GetOpenIDDTO{}
//	ctx.ReadQuery(&dto)
//	if dto.Code == "" {
//		ctx.JSON(Response{ErrorCode, "参数错误", nil})
//	}
//
//	openID := model.GetUserOpenID(dto.Code)
//	if openID == "" {
//		ctx.JSON(Response{ErrorCode, "获取不到openID,请检查参数是否有效", nil})
//		return
//	}
//
//	result := map[string]string{
//		"open_id": openID,
//	}
//	ctx.JSON(Response{SuccessCode, "", result})
//}

func (u *UserServiceApi) Register(ctx iris.Context) {
	userInfoDTO := model.UserInfoDTO{}
	ctx.ReadJSON(&userInfoDTO)
	if userInfoDTO.NickName == "" || userInfoDTO.Avatar == "" || userInfoDTO.OpenID == "" {
		ctx.JSON(Response{ErrorCode, "参数错误", nil})
	}

	model.Register(userInfoDTO)
	ctx.JSON(Response{SuccessCode, "", nil})
}
