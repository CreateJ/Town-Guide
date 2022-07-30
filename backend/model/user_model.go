package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"town-guide/dao"
)

type UserBaseInfoDTO struct {
	Gender   int8   `json:"gender"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	OpenID   string `json:"open_id"`
}

type UserInfoDTO struct {
	Gender               int8   `json:"gender"`
	NickName             string `json:"nick_name"`
	Avatar               string `json:"avatar"`
	OpenID               string `json:"open_id"`
	CollectionScenicInfo *[]ScenicInfoDTO
	ClockScenicInfo      *[]ScenicInfoDTO
	ClockNum             int64 `json:"clock_num"`
}

type TokenResult struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type OpenIDResult struct {
	ErrCode int    `json:"session_key"`
	OpenID  string `json:"openid"`
}

func GetUserBaseInfo(code string) *UserBaseInfoDTO {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=wx942fe69ede395b15&secret=499b3a99cec40c32a8c0f0796c2a0aa2&js_code=" + code + "&grant_type=authorization_code"
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var res OpenIDResult
	json.Unmarshal(body, &res)
	if res.OpenID == "" {
		return nil
	}
	userdao := dao.GetUserDao()
	one := userdao.GetOne(res.OpenID)
	result := &UserBaseInfoDTO{
		OpenID: res.OpenID,
	}
	if one == nil || one.NickName == "" {
		return result
	}

	result.OpenID = res.OpenID
	result.Gender = one.Gender
	result.Avatar = one.Url
	result.NickName = one.NickName
	return result
}

func GetUserDetail(openID string) *UserInfoDTO {
	if openID == "" {
		return nil
	}
	userDao := dao.GetUserDao()
	userBaseInfo := userDao.GetOne(openID)
	result := &UserInfoDTO{
		OpenID:   userBaseInfo.OpenID,
		NickName: userBaseInfo.NickName,
		Avatar:   userBaseInfo.Url,
		Gender:   userBaseInfo.Gender,
	}

	userActionDao := dao.GetUserActionDao()
	scenicDao := dao.GetScenicDao()
	userClock := userActionDao.QueryUserClock(openID)
	if userClock != nil && len(*userClock) > 0 {
		clockList := make([]ScenicInfoDTO, 0, len(*userClock))

		for _, v := range *userClock {
			scenicInfo := scenicDao.QueryOne(v.ScenicID)
			if scenicInfo != nil {
				temp := ScenicInfoDTO{
					ID:           scenicInfo.ID,
					Name:         scenicInfo.Name,
					LocationDesc: scenicInfo.LocationDesc,
					Description:  scenicInfo.Description,
					Intro:        scenicInfo.Intro,
					PicName:       scenicInfo.PicName,
					Icon:         scenicInfo.Icon,
					VideoName:     scenicInfo.VideoName,
					Tag:          scenicInfo.Tag,
					OpenTime:     scenicInfo.OpenTime,
					CheckNum:     scenicInfo.CheckNum,
				}
				clockList = append(clockList, temp)
			}
			result.ClockScenicInfo = &clockList
			result.ClockNum = int64(len(clockList))
		}
	}
	userCollection := userActionDao.QueryUserCollection(openID)
	if userCollection != nil && len(*userCollection) > 0 {
		collectionList := make([]ScenicInfoDTO, 0, len(*userCollection))
		for _, v := range *userCollection {
			scenicInfo := scenicDao.QueryOne(v.ScenicID)
			if scenicInfo != nil {
				temp := ScenicInfoDTO{
					ID:           scenicInfo.ID,
					Name:         scenicInfo.Name,
					LocationDesc: scenicInfo.LocationDesc,
					Description:  scenicInfo.Description,
					Intro:        scenicInfo.Intro,
					PicName:       scenicInfo.PicName,
					Icon:         scenicInfo.Icon,
					VideoName:     scenicInfo.VideoName,
					Tag:          scenicInfo.Tag,
					OpenTime:     scenicInfo.OpenTime,
					CheckNum:     scenicInfo.CheckNum,
				}
				collectionList = append(collectionList, temp)
			}
			result.ClockScenicInfo = &collectionList
		}

	}
	return result
}

func Register(userInfo UserBaseInfoDTO) {
	userDao := dao.GetUserDao()
	info := userDao.GetOne(userInfo.OpenID)
	if info != nil {
		return
	}

	insertInfo := &dao.TbUserInfo{
		OpenID:   userInfo.OpenID,
		NickName: userInfo.NickName,
		Url:      userInfo.Avatar,
		Gender:   userInfo.Gender,
	}
	_, _ = userDao.Insert(insertInfo)
}
