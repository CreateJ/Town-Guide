package model

import (
	"errors"
	"time"
	"town-guide/dao"
)

type ScenicInfoDTO struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	LocationDesc string `json:"location_desc"`
	Description  string `json:"description"`
	Intro        string `json:"intro"`
	PicUrl       string `json:"pic_url"`
	Icon         string `json:"icon"`
	VideoUrl     string `json:"video_url"`
	Tag          string `json:"tag"`
	OpenTime     string `json:"open_time"`
	CheckNum     int64  `json:"check_num"`
	CreateTime   int64  `json:"create_time"`
}

func AddScenic(scenicInfo *ScenicInfoDTO) (*ScenicInfoDTO, error) {
	info := &dao.TbScenicInfo{
		Name:         scenicInfo.Name,
		LocationDesc: scenicInfo.LocationDesc,
		Description:  scenicInfo.Description,
		Intro:        scenicInfo.Intro,
		PicUrl:       scenicInfo.PicUrl,
		Icon:         scenicInfo.Icon,
		VideoUrl:     scenicInfo.VideoUrl,
		Tag:          scenicInfo.Tag,
		OpenTime:     scenicInfo.OpenTime,
		CheckNum:     0,
		CreateTime:   time.Now().Unix(),
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
	for _, v := range *scenicInfos {
		temp := ScenicInfoDTO{
			Name:         v.Name,
			LocationDesc: v.LocationDesc,
			Description:  v.Description,
			Intro:        v.Intro,
			PicUrl:       v.PicUrl,
			Icon:         v.Icon,
			VideoUrl:     v.VideoUrl,
			Tag:          v.Tag,
			OpenTime:     v.OpenTime,
			CheckNum:     v.CheckNum,
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
		Name:         scenicInfo.Name,
		LocationDesc: scenicInfo.LocationDesc,
		Description:  scenicInfo.Description,
		Intro:        scenicInfo.Intro,
		PicUrl:       scenicInfo.PicUrl,
		Icon:         scenicInfo.Icon,
		VideoUrl:     scenicInfo.VideoUrl,
		Tag:          scenicInfo.Tag,
		OpenTime:     scenicInfo.OpenTime,
		CheckNum:     scenicInfo.CheckNum,
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
		Name:         scenicInfo.Name,
		LocationDesc: scenicInfo.LocationDesc,
		Description:  scenicInfo.Description,
		Intro:        scenicInfo.Intro,
		PicUrl:       scenicInfo.PicUrl,
		Icon:         scenicInfo.Icon,
		VideoUrl:     scenicInfo.VideoUrl,
		Tag:          scenicInfo.Tag,
		OpenTime:     scenicInfo.OpenTime,
	}
	_, err := scenicDao.Edit(info)
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
