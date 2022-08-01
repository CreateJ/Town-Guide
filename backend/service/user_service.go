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

	info := model.GetUserBaseInfo(dto.Code)
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
	userInfoDTO := model.UserBaseInfoDTO{}
	_ = ctx.ReadJSON(&userInfoDTO)
	if userInfoDTO.NickName == "" || userInfoDTO.Avatar == "" || userInfoDTO.OpenID == "" {
		_, _ = ctx.JSON(Response{ErrorCode, "参数错误", nil})
		return
	}

	model.Register(userInfoDTO)
	_, _ = ctx.JSON(Response{SuccessCode, "", nil})
}

func (u *UserServiceApi) GetUserDetail(ctx iris.Context) {
	openId := ctx.Params().Get("open_id")
	if openId == "" {
		_, _ = ctx.JSON(Response{ErrorCode, "参数错误", nil})
	}

	info := model.GetUserDetail(openId)
	if info == nil {
		_, _ = ctx.JSON(Response{ErrorCode, "获取不到该用户信息", nil})
		return
	}

	_, _ = ctx.JSON(Response{SuccessCode, "", info})
}

func (u *UserServiceApi) UserClockByScenicID(ctx iris.Context) {
	dto := model.UserActionDTO{}
	ctx.ReadJSON(&dto)
	if dto.OpenID == "" || dto.ScenicID == 0 || dto.ActionState == 0 {
		_, _ = ctx.JSON(Response{ErrorCode, "参数错误", nil})
		return
	}

	if dto.ActionState == 2 {
		if !model.UserClock(dto.OpenID, dto.ScenicID) {
			_, _ = ctx.JSON(Response{ErrorCode, "打卡失败", nil})
			return
		}
		_, _ = ctx.JSON(Response{SuccessCode, "打卡成功", nil})
		return
	}
	_, _ = ctx.JSON(Response{ErrorCode, "参数错误", nil})
	return
}

func (u *UserServiceApi) UserCollectionByScenicID(ctx iris.Context) {
	dto := model.UserActionDTO{}
	ctx.ReadJSON(&dto)
	if dto.OpenID == "" || dto.ScenicID == 0 || dto.ActionState == 0 {
		_, _ = ctx.JSON(Response{ErrorCode, "参数错误", nil})
		return
	}

	if dto.ActionState == 2 {
		if !model.UserCollection(dto.OpenID, dto.ScenicID) {
			_, _ = ctx.JSON(Response{ErrorCode, "收藏失败", nil})
			return
		}
		_, _ = ctx.JSON(Response{SuccessCode, "收藏成功", nil})
		return
	}
	if dto.ActionState == 1 {
		if !model.UserCancelCollection(dto.OpenID, dto.ScenicID) {
			_, _ = ctx.JSON(Response{ErrorCode, "取消收藏失败", nil})
			return
		}
		_, _ = ctx.JSON(Response{SuccessCode, "取消收藏成功", nil})
		return
	}
	_, _ = ctx.JSON(Response{ErrorCode, "参数错误", nil})
	return
}
