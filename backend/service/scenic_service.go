package service

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"town-guide/model"
)

type ScenicServiceApi struct{}

func NewScenicService() *ScenicServiceApi {
	return &ScenicServiceApi{}
}

type UserOpenIDDTO struct {
	OpenID string `url:"open_id"`
}

func (u *ScenicServiceApi) GetScenic(ctx iris.Context) {
	scenicID, err := ctx.Params().GetInt64("scenic_id")
	if err != nil || scenicID <= 0 {
		_, _ = ctx.JSON(Response{ErrorCode, "参数错误", nil})
		return
	}

	info := model.QueryScenicByID(scenicID)
	if info == nil {
		_, _ = ctx.JSON(Response{ErrorCode, "获取景区信息失败", nil})
		return
	}

	dto := UserOpenIDDTO{}
	err = ctx.ReadQuery(&dto)
	if err != nil || dto.OpenID == "" {
		_, _ = ctx.JSON(Response{SuccessCode, "", info})
		return
	}

	if model.GetUserClockState(dto.OpenID, scenicID) {
		info.UserClockState = 2
	}

	if model.GetUserCollectionState(dto.OpenID, scenicID) {
		info.UserCollectionState = 2
	}
	_, _ = ctx.JSON(Response{SuccessCode, "", info})
}

func (u *ScenicServiceApi) AddScenic(ctx iris.Context) {
	dto := model.ScenicInfoDTO{}
	err := ctx.ReadJSON(&dto)
	if err != nil {
		_, _ = ctx.JSON(Response{ErrorCode, err.Error(), nil})
		return
	}

	info, err := model.AddScenic(&dto)
	if err != nil {
		_, _ = ctx.JSON(Response{ErrorCode, err.Error(), nil})
		return
	}

	_, _ = ctx.JSON(Response{SuccessCode, "", info})
}

func (u *ScenicServiceApi) DeleteScenic(ctx iris.Context) {
	id, err := ctx.Params().GetInt64("scenic_id")
	if err != nil {
		_, _ = ctx.JSON(Response{ErrorCode, "参数错误", nil})
		return
	}

	err = model.DeleteScenicByID(id)
	if err != nil {
		_, _ = ctx.JSON(Response{ErrorCode, err.Error(), nil})
		return
	}
	_, _ = ctx.JSON(Response{SuccessCode, "", nil})
}

func (u *ScenicServiceApi) EditScenic(ctx iris.Context) {
	dto := model.ScenicInfoDTO{}
	err := ctx.ReadJSON(&dto)
	if err != nil || dto.ID <= 0 {
		_, _ = ctx.JSON(Response{ErrorCode, "参数错误", nil})
		return
	}
	err = model.EditScenic(&dto)
	if err != nil {
		_, _ = ctx.JSON(Response{ErrorCode, "修改失败", nil})
		return
	}

	_, _ = ctx.JSON(Response{SuccessCode, "修改成功", nil})
}

func (u *ScenicServiceApi) GetAllScenic(ctx iris.Context) {
	info := model.QueryAllScenic()
	dto := UserOpenIDDTO{}
	var result []model.ScenicInfoDTO
	err := ctx.ReadQuery(&dto)
	if err != nil || dto.OpenID == "" {
		_, _ = ctx.JSON(Response{SuccessCode, "", info})
		return
	}

	for _, v := range *info {
		temp := v
		if model.GetUserClockState(dto.OpenID, temp.ID) {
			temp.UserClockState = 2
		}
		if model.GetUserCollectionState(dto.OpenID, temp.ID) {
			temp.UserCollectionState = 2
		}
		result = append(result, temp)
	}
	_, _ = ctx.JSON(Response{SuccessCode, "", result})
}

func (u *ScenicServiceApi) GetScenicByCategoryID(ctx iris.Context) {
	id, err := ctx.Params().GetInt64("category_id")
	info := new([]model.ScenicInfoDTO)
	if err != nil {
		info = model.QueryAllScenic()
	} else {
		info = model.QueryScenicByCategoryID(id)
	}

	dto := UserOpenIDDTO{}
	var result []model.ScenicInfoDTO
	err = ctx.ReadQuery(&dto)
	if err != nil || dto.OpenID == "" {
		_, _ = ctx.JSON(Response{SuccessCode, "", info})
		return
	}

	for _, v := range *info {
		temp := v

		if model.GetUserClockState(dto.OpenID, temp.ID) {
			fmt.Println(temp)
			fmt.Println(dto)
			temp.UserClockState = 2
		}
		if model.GetUserCollectionState(dto.OpenID, temp.ID) {
			temp.UserCollectionState = 2
		}
		result = append(result, temp)
	}
	_, _ = ctx.JSON(Response{SuccessCode, "", result})
}
