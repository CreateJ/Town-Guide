package model

import (
	"errors"
	"fmt"
	"time"
	"town-guide/dao"
)

type ScenicInfoDTO struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	LocationDesc string `json:"location_desc"`
	Description  string `json:"description"`
	Intro        string `json:"intro"`
	PicName      string `json:"pic_name"`
	Icon         string `json:"icon"`
	VideoName    string `json:"video_name"`
	Tag          string `json:"tag"`
	OpenTime     string `json:"open_time"`
	CheckNum     int64  `json:"check_num"`
	CategoryID   int64  `json:"category_id"`
	Banner       string `json:"banner"` // 用｜分割
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
			CheckNum:     scenicInfo.CheckNum,
			CategoryID:   scenicInfo.CategoryID,
			Banner:       scenicInfo.Banner,
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
		CheckNum:     scenicInfo.CheckNum,
		CategoryID:   scenicInfo.CategoryID,
		Banner:       scenicInfo.Banner,
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
	}
	_, err := scenicDao.Edit(info)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func AddScenicCheckByID(id int64) error {
	if id <= 0 {
		return errors.New("id err")
	}

	scenicDao := dao.GetScenicDao()
	_, err := scenicDao.IncrCheckNum(id)
	if err != nil {
		return err
	}
	return nil
}
