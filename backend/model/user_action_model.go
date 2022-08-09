package model

import (
	"time"
	"town-guide/dao"
)

type UserActionDTO struct {
	OpenID      string `json:"open_id"`
	ScenicID    int64  `json:"scenic_id"`
	ActionState int8   `json:"action_state"`
}

func GetUserClockState(openID string, scenicID int64) bool {
	actionDao := dao.GetUserActionDao()
	clock := actionDao.QueryUserScenicClock(openID, scenicID)
	if clock == nil {
		return false
	}
	return true
}

func GetUserCollectionState(openID string, scenicID int64) bool {
	actionDao := dao.GetUserActionDao()
	collection := actionDao.QueryUserScenicCollection(openID, scenicID)
	if collection == nil {
		return false
	}
	return true
}

func UserClock(openID string, scenicID int64) bool {
	actionDao := dao.GetUserActionDao()
	clock := actionDao.QueryUserScenicClock(openID, scenicID)
	if clock != nil {
		return true
	}
	scenicClock := &dao.TbUserScenicClock{
		OpenID:     openID,
		ScenicID:   scenicID,
		UpdateTime: time.Now().Unix(),
		CreateTime: time.Now().Unix(),
	}
	_, _ = actionDao.InsertClock(scenicClock)
	_ = AddScenicClockNumByID(scenicID)
	return true
}

func UserCollection(openID string, scenicID int64) bool {
	actionDao := dao.GetUserActionDao()
	if actionDao.GetUserCollectionState(openID, scenicID) {
		return true
	}

	scenicCollection := &dao.TbUserScenicCollection{
		OpenID:     openID,
		ScenicID:   scenicID,
		UpdateTime: time.Now().Unix(),
		CreateTime: time.Now().Unix(),
	}
	_, _ = actionDao.InsertCollection(scenicCollection)
	return true
}

func UserCancelCollection(openID string, scenicID int64) bool {
	actionDao := dao.GetUserActionDao()
	if !actionDao.GetUserCollectionState(openID, scenicID) {
		return true
	}

	err := actionDao.DeleteUserCollection(openID, scenicID)
	if err != nil {
		return false
	}
	return true
}
