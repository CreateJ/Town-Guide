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
	openId := ctx.Params().Get("open_id")
	scenicID, err := ctx.Params().GetInt64("scenic_id")

	if err != nil || openId == "" || scenicID <= 0 {
		_, _ = ctx.JSON(Response{ErrorCode, "参数错误", nil})
		return
	}

	userBaseInfo := model.GetUserBaseInfo(openId)
	if userBaseInfo == nil {
		_, _ = ctx.JSON(Response{ErrorCode, "获取不到该用户信息", nil})
		return
	}
	scenicInfo := model.QueryScenicByID(scenicID)
	if scenicInfo == nil {
		_, _ = ctx.JSON(Response{ErrorCode, "获取不到该景区信息", nil})
		return
	}

	state := model.GetUserClockState(openId, scenicID)
	if state {
		_, _ = ctx.JSON(Response{ErrorCode, "已打卡", nil})
		return
	}

	_ = model.UserClock(openId, scenicID)
	_, _ = ctx.JSON(Response{SuccessCode, "", nil})
}
