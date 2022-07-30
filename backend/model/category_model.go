package model

import (
	"town-guide/dao"
)

type CategoryInfoDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func AddCategory(categoryInfo *CategoryInfoDTO) (*CategoryInfoDTO, error) {
	info := &dao.TbCategory{
		Name: categoryInfo.Name,
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
			ID:   info.ID,
			Name: info.Name,
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
		ID:   info.ID,
		Name: info.Name,
	}
}

func DeleteCategoryByID(id int64) error {
	if id <= 0 {
		return nil
	}

	categoryDao := dao.GetCategoryDao()
	return categoryDao.DeleteOne(id)
}
