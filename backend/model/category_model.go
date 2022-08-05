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
	info := &dao.TbCategory{
		Name:       categoryInfo.Name,
		Icon:       categoryInfo.Icon,
		IconActive: categoryInfo.IconActive,
	}
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
		temp := CategoryInfoDTO{
			ID:         info.ID,
			Name:       info.Name,
			Icon:       info.Icon,
			IconActive: info.IconActive,
		}
		result = append(result, temp)
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

	return &CategoryInfoDTO{
		ID:         info.ID,
		Name:       info.Name,
		Icon:       info.Icon,
		IconActive: info.IconActive,
	}
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
	info := &dao.TbCategory{
		ID:         categoryInfo.ID,
		Name:       categoryInfo.Name,
		Icon:       categoryInfo.Icon,
		IconActive: categoryInfo.IconActive,
	}
	_, err := categoryDao.Edit(info)
	if err != nil {
		return err
	}
	return nil
}
