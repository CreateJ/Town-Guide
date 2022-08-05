package model

import (
	"errors"
	"time"
	"town-guide/dao"
)

type ScenicInfoDTO struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	LocationDesc   string `json:"location_desc"`
	Description    string `json:"description"`
	Intro          string `json:"intro"`
	PicName        string `json:"pic_name"`
	Icon           string `json:"icon"`
	VideoName      string `json:"video_name"`
	Tag            string `json:"tag"`
	OpenTime       string `json:"open_time"`
	ClockNum       int64  `json:"clock_num"`
	ClockIcon      string `json:"clock_icon"`
	CategoryID     int64  `json:"category_id"`
	Category       string `json:"category"`
	Banner         string `json:"banner"` // 用｜分割
	Location       string `json:"location"`
	Audio          string `json:"audio"`
	UserClockState int8   `json:"user_clock_state"`
}

func AddScenic(scenicInfo *ScenicInfoDTO) (*ScenicInfoDTO, error) {
	info := &dao.TbScenicInfo{
		Name:         scenicInfo.Name,
		LocationDesc: scenicInfo.LocationDesc,
		Description:  scenicInfo.Description,
		Intro:        scenicInfo.Intro,
		PicName:      scenicInfo.PicName,
		Icon:         scenicInfo.Icon,
		VideoName:    scenicInfo.VideoName,
		Tag:          scenicInfo.Tag,
		OpenTime:     scenicInfo.OpenTime,
		CategoryID:   scenicInfo.CategoryID,
		CreateTime:   time.Now().Unix(),
		UpdateTime:   time.Now().Unix(),
		Banner:       scenicInfo.Banner,
		Location:     scenicInfo.Location,
		Audio:        scenicInfo.Audio,
		ClockIcon:    scenicInfo.ClockIcon,
	}

	scenicDao := dao.GetScenicDao()
	id, err := scenicDao.Insert(info)
	if err != nil {
		return nil, err
	}
	scenicInfo.ID = id
	return scenicInfo, nil
}

func QueryAllScenic() *[]ScenicInfoDTO {
	scenicDao := dao.GetScenicDao()
	scenicInfos := scenicDao.QueryAll()
	if scenicInfos == nil || len(*scenicInfos) <= 0 {
		return nil
	}
	categoryDao := dao.GetCategoryDao()
	allCategory := categoryDao.QueryAll()
	categoryMap := map[int64]string{}
	for _, v := range *allCategory {
		categoryMap[v.ID] = v.Name
	}
	result := make([]ScenicInfoDTO, 0, len(*scenicInfos))
	for _, scenicInfo := range *scenicInfos {
		categoryID := scenicInfo.CategoryID
		category, ok := categoryMap[categoryID]
		if !ok {
			categoryID = 0
			category = "其他"
		}
		temp := ScenicInfoDTO{
			ID:           scenicInfo.ID,
			Name:         scenicInfo.Name,
			LocationDesc: scenicInfo.LocationDesc,
			Description:  scenicInfo.Description,
			Intro:        scenicInfo.Intro,
			PicName:      scenicInfo.PicName,
			Icon:         scenicInfo.Icon,
			VideoName:    scenicInfo.VideoName,
			Tag:          scenicInfo.Tag,
			OpenTime:     scenicInfo.OpenTime,
			ClockNum:     scenicInfo.ClockNum,
			CategoryID:   categoryID,
			Banner:       scenicInfo.Banner,
			Category:     category,
			Location:     scenicInfo.Location,
			Audio:        scenicInfo.Audio,
			ClockIcon:    scenicInfo.ClockIcon,
		}
		result = append(result, temp)
	}
	return &result
}

func QueryScenicByCategoryID(categoryID int64) *[]ScenicInfoDTO {
	scenicDao := dao.GetScenicDao()
	scenicInfos := scenicDao.QueryByCategoryID(categoryID)
	if scenicInfos == nil || len(*scenicInfos) <= 0 {
		return nil
	}

	result := make([]ScenicInfoDTO, 0, len(*scenicInfos))
	for _, scenicInfo := range *scenicInfos {
		temp := ScenicInfoDTO{
			ID:           scenicInfo.ID,
			Name:         scenicInfo.Name,
			LocationDesc: scenicInfo.LocationDesc,
			Description:  scenicInfo.Description,
			Intro:        scenicInfo.Intro,
			PicName:      scenicInfo.PicName,
			Icon:         scenicInfo.Icon,
			VideoName:    scenicInfo.VideoName,
			Tag:          scenicInfo.Tag,
			OpenTime:     scenicInfo.OpenTime,
			ClockNum:     scenicInfo.ClockNum,
			CategoryID:   scenicInfo.CategoryID,
			Banner:       scenicInfo.Banner,
			Location:     scenicInfo.Location,
			Audio:        scenicInfo.Audio,
			ClockIcon:    scenicInfo.ClockIcon,
		}
		result = append(result, temp)
	}
	return &result
}

func QueryScenicByID(id int64) *ScenicInfoDTO {
	if id <= 0 {
		return nil
	}

	scenicDao := dao.GetScenicDao()
	scenicInfo := scenicDao.QueryOne(id)
	if scenicInfo == nil {
		return nil
	}

	categoryDao := dao.GetCategoryDao()
	allCategory := categoryDao.QueryOne(scenicInfo.CategoryID)
	categoryID := int64(0)
	category := "其他"
	if allCategory != nil {
		categoryID = scenicInfo.CategoryID
		category = allCategory.Name
	}
	return &ScenicInfoDTO{
		ID:           scenicInfo.ID,
		Name:         scenicInfo.Name,
		LocationDesc: scenicInfo.LocationDesc,
		Description:  scenicInfo.Description,
		Intro:        scenicInfo.Intro,
		PicName:      scenicInfo.PicName,
		Icon:         scenicInfo.Icon,
		VideoName:    scenicInfo.VideoName,
		Tag:          scenicInfo.Tag,
		OpenTime:     scenicInfo.OpenTime,
		ClockNum:     scenicInfo.ClockNum,
		CategoryID:   categoryID,
		Banner:       scenicInfo.Banner,
		Category:     category,
		Location:     scenicInfo.Location,
		Audio:        scenicInfo.Audio,
		ClockIcon:    scenicInfo.ClockIcon,
	}
}

func DeleteScenicByID(id int64) error {
	if id <= 0 {
		return nil
	}

	scenicDao := dao.GetScenicDao()
	return scenicDao.DeleteOne(id)
}

func EditScenic(scenicInfo *ScenicInfoDTO) error {
	if scenicInfo.ID <= 0 {
		return errors.New("id err")
	}

	scenicDao := dao.GetScenicDao()
	info := &dao.TbScenicInfo{
		ID:           scenicInfo.ID,
		Name:         scenicInfo.Name,
		LocationDesc: scenicInfo.LocationDesc,
		Description:  scenicInfo.Description,
		Intro:        scenicInfo.Intro,
		PicName:      scenicInfo.PicName,
		Icon:         scenicInfo.Icon,
		VideoName:    scenicInfo.VideoName,
		Tag:          scenicInfo.Tag,
		OpenTime:     scenicInfo.OpenTime,
		Banner:       scenicInfo.Banner,
		CategoryID:   scenicInfo.CategoryID,
		Location:     scenicInfo.Location,
		Audio:        scenicInfo.Audio,
		ClockIcon:    scenicInfo.ClockIcon,
	}
	_, err := scenicDao.Edit(info)
	if err != nil {
		return err
	}
	return nil
}

func AddScenicClockNumByID(id int64) error {
	if id <= 0 {
		return errors.New("id err")
	}

	scenicDao := dao.GetScenicDao()
	_, err := scenicDao.IncrClockNum(id)
	if err != nil {
		return err
	}
	return nil
}
