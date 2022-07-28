package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"town-guide/dao"
)

func login(openID string) *dao.UserInfo {
	userDao := dao.GetUserDao()
	userInfo := userDao.GetOne(openID)
	if userInfo == nil {
		return nil
	}
	return nil
}

//func register(openID string, userName string, url string, gender int8) error {
//	info := &dao.UserInfo{
//		OpenID:     openID,
//		Url:        url,
//		Gender:     gender,
//		CreateTime: time.Now().Unix(),
//		NickName:   userName,
//	}
//	userDao := dao.GetUserDao()
//	insert, err := userDao.Insert(info)
//	if err != nil {
//
//	}
//	return nil
//}

//	POST
//https: //api.weixin.qq.com/wxa/getpluginopenpid?access_token=ACCESS_TOKEN
//	{
//		"errcode": 0,
//		"errmsg": "ok",
//		"openpid": "GACo74wkDIkDzEhkwRwgjGt1pqlk",
//	}
//	GET
//https: //api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET

//}

//type UserInfo struct {
//	`id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
//	`open_id`     varchar(225)        NULL     DEFAULT '',
//	`nick_name`   varchar(225)        NULL     DEFAULT '',
//	`url`         varchar(225)        NULL     DEFAULT '',
//	`gender`      tinyint(4)          NOT NULL DEFAULT '0',
//	`create_time` i
//}

func getUserInfoByOpenID() {

}

type TokenResult struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
type Result struct {
	ErrCode int    `json:"session_key"`
	OpenID  string `json:"openid"`
}

func GetUserOpenID(code string) string {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=wx942fe69ede395b15&secret=499b3a99cec40c32a8c0f0796c2a0aa2&js_code=" + code + "&grant_type=authorization_code"
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var res Result
	json.Unmarshal(body, &res)
	return res.OpenID
}
