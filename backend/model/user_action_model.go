package model

import (
	"time"
	"town-guide/dao"
)

func GetUserClockState(openID string, scenicID int64) bool {
	actionDao := dao.GetUserActionDao()
	clock := actionDao.QueryUserScenicClock(openID, scenicID)
	if clock == nil {
		return false
	}
	return true
}

func UserClock(openID string, scenicID int64) bool {
	actionDao := dao.GetUserActionDao()
	scenicClock := &dao.TbUserScenicClock{
		OpenID:     openID,
		ScenicID:   scenicID,
		UpdateTime: time.Now().Unix(),
		CreateTime: time.Now().Unix(),
	}
	_, _ = actionDao.InsertClock(scenicClock)
	return true
}

func UserCollection(openID string, scenicID int64) bool {
	actionDao := dao.GetUserActionDao()
	scenicCollection := &dao.TbUserScenicCollection{
		OpenID:     openID,
		ScenicID:   scenicID,
		UpdateTime: time.Now().Unix(),
		CreateTime: time.Now().Unix(),
	}
	_, _ = actionDao.InsertCollection(scenicCollection)
	return true
}
