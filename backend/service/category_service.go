package service

import (
	"github.com/kataras/iris/v12"
	"town-guide/model"
)

type CategoryServiceApi struct{}

func NewCategoryService() *CategoryServiceApi {
	return &CategoryServiceApi{}
}

func (u *CategoryServiceApi) GetAllCategory(ctx iris.Context) {
	info := model.QueryAllCategory()
	_, _ = ctx.JSON(Response{SuccessCode, "", info})
}

func (u *CategoryServiceApi) AddCategory(ctx iris.Context) {
	dto := model.CategoryInfoDTO{}
	err := ctx.ReadJSON(&dto)
	if err != nil {
		_, _ = ctx.JSON(Response{ErrorCode, err.Error(), nil})
		return
	}

	info, err := model.AddCategory(&dto)
	if err != nil {
		_, _ = ctx.JSON(Response{ErrorCode, err.Error(), nil})
		return
	}

	_, _ = ctx.JSON(Response{SuccessCode, "", info})
}

func (u *CategoryServiceApi) DeleteCategory(ctx iris.Context) {
	id, err := ctx.Params().GetInt64("category_id")
	if err != nil {
		_, _ = ctx.JSON(Response{ErrorCode, "参数错误", nil})
		return
	}

	err = model.DeleteCategoryByID(id)
	if err != nil {
		_, _ = ctx.JSON(Response{ErrorCode, err.Error(), nil})
		return
	}
	_, _ = ctx.JSON(Response{SuccessCode, "", nil})
}

func (u *CategoryServiceApi) EditCategory(ctx iris.Context) {
	dto := model.CategoryInfoDTO{}
	_ = ctx.ReadJSON(&dto)
	if dto.ID <= 0 {
		_, _ = ctx.JSON(Response{ErrorCode, "参数错误", nil})
		return
	}
	err := model.EditCategory(&dto)
	if err != nil {
		_, _ = ctx.JSON(Response{ErrorCode, "修改失败", nil})
		return
	}

	_, _ = ctx.JSON(Response{SuccessCode, "", nil})
}
