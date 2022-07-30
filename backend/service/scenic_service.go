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

func (u *ScenicServiceApi) GetScenic(ctx iris.Context) {
	scenicID, err := ctx.Params().GetInt64("scenic_id")
	if err != nil || scenicID <= 0 {
		ctx.JSON(Response{ErrorCode, "参数错误", nil})
		return
	}

	info := model.QueryScenicByID(scenicID)
	if info == nil {
		ctx.JSON(Response{ErrorCode, "获取景区信息失败", nil})
		return
	}

	ctx.JSON(Response{SuccessCode, "", info})
}

func (u *ScenicServiceApi) AddScenic(ctx iris.Context) {
	dto := model.ScenicInfoDTO{}
	ctx.ReadQuery(&dto)
	info, err := model.AddScenic(&dto)
	if err != nil {
		ctx.JSON(Response{ErrorCode, err.Error(), nil})
		return
	}

	ctx.JSON(Response{SuccessCode, "", info})
}

func (u *ScenicServiceApi) DeleteScenic(ctx iris.Context) {
	id, err := ctx.Params().GetInt64("id")
	if err != nil {
		_, _ = ctx.JSON(Response{ErrorCode, "参数错误", nil})
		return
	}

	err = model.DeleteScenicByID(id)
	if err != nil {
		ctx.JSON(Response{ErrorCode, err.Error(), nil})
		return
	}

	ctx.JSON(Response{SuccessCode, "", nil})
}

func (u *ScenicServiceApi) EditScenic(ctx iris.Context) {
	dto := model.ScenicInfoDTO{}
	ctx.ReadJSON(&dto)
	if dto.ID <= 0 {
		ctx.JSON(Response{ErrorCode, "参数错误", nil})
		return
	}

	err := model.EditScenic(&dto)
	if err != nil {
		ctx.JSON(Response{ErrorCode, "修改失败", nil})
		return
	}

	ctx.JSON(Response{SuccessCode, "", nil})
}

func (u *ScenicServiceApi) GetAllScenic(ctx iris.Context) {
	info := model.QueryAllScenic()
	_, _ = ctx.JSON(Response{SuccessCode, "", info})
}
