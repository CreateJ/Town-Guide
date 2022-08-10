package model

import (
	"errors"
	"town-guide/dao"
)

type CategoryInfoDTO struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	IconActive string `json:"icon_active"`
}

func AddCategory(categoryInfo *CategoryInfoDTO) (*CategoryInfoDTO, error) {
	info := CategoryInfoDTOToTbCategory(categoryInfo)
	info.ID = 0
	categoryDao := dao.GetCategoryDao()
	id, err := categoryDao.Insert(info)
	if err != nil {
		return nil, err
	}
	categoryInfo.ID = id
	return categoryInfo, nil
}

func QueryAllCategory() *[]CategoryInfoDTO {
	categoryDao := dao.GetCategoryDao()
	categoryInfos := categoryDao.QueryAll()
	if categoryInfos == nil || len(*categoryInfos) <= 0 {
		return nil
	}

	result := make([]CategoryInfoDTO, 0, len(*categoryInfos))
	for _, info := range *categoryInfos {
		temp := TbCategoryToCategoryInfoDTO(&info)
		result = append(result, *temp)
	}
	return &result
}

func QueryCategoryByID(id int64) *CategoryInfoDTO {
	if id <= 0 {
		return nil
	}

	categoryDao := dao.GetCategoryDao()
	info := categoryDao.QueryOne(id)
	if info == nil {
		return nil
	}
	return TbCategoryToCategoryInfoDTO(info)
}

func DeleteCategoryByID(id int64) error {
	if id <= 0 {
		return nil
	}
	categoryDao := dao.GetCategoryDao()
	return categoryDao.DeleteOne(id)
}

func EditCategory(categoryInfo *CategoryInfoDTO) error {
	if categoryInfo.ID <= 0 {
		return errors.New("id err")
	}

	categoryDao := dao.GetCategoryDao()
	info := CategoryInfoDTOToTbCategory(categoryInfo)
	if _, err := categoryDao.Edit(info); err != nil {
		return err
	}
	return nil
}

func CategoryInfoDTOToTbCategory(categoryInfo *CategoryInfoDTO) *dao.TbCategory {
	return &dao.TbCategory{
		ID:         categoryInfo.ID,
		Name:       categoryInfo.Name,
		Icon:       categoryInfo.Icon,
		IconActive: categoryInfo.IconActive,
	}
}

func TbCategoryToCategoryInfoDTO(categoryInfo *dao.TbCategory) *CategoryInfoDTO {
	return &CategoryInfoDTO{
		ID:         categoryInfo.ID,
		Name:       categoryInfo.Name,
		Icon:       categoryInfo.Icon,
		IconActive: categoryInfo.IconActive,
	}
}
