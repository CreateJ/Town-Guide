package dao

import (
	"github.com/tietang/dbx"
	"town-guide/base"
)

type TbUserScenicCollection struct {
	Id         int64  `db:"id,omitempty"`
	OpenID     string `db:"open_id"`
	ScenicID   string `db:"scenic_id"`
	UpdateTime int8   `db:"update_time"`
	CreateTime int64  `db:"create_time"`
}
type TbUserScenicClock struct {
	Id         int64  `db:"id,omitempty"`
	OpenID     string `db:"open_id"`
	ScenicID   string `db:"scenic_id"`
	UpdateTime int8   `db:"update_time"`
	CreateTime int64  `db:"create_time"`
}
type UserActionDao struct {
	runner *dbx.Database
}

func GetUserActionDao() UserActionDao {
	return UserActionDao{
		runner: base.DbxDatabase(),
	}
}

func (dao *UserActionDao) GetOne(openID string) *TbUserInfo {
	a := &TbUserInfo{OpenID: openID}
	ok, err := dao.runner.GetOne(a)

	if err != nil {
		return nil
	}
	if !ok {
		return nil
	}
	return a
}

func (dao *UserActionDao) InsertCollection(a *TbUserScenicCollection) (id int64, err error) {
	rs, err := dao.runner.Insert(a)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}

func (dao *UserActionDao) InsertClock(a *TbUserScenicClock) (id int64, err error) {
	rs, err := dao.runner.Insert(a)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}

func (dao *UserActionDao) QueryUserClock(openID string) *[]TbUserScenicClock {
	query := &TbUserScenicClock{OpenID: openID}
	var dst []TbUserScenicClock
	err := dao.runner.FindExample(query, &dst)
	if err != nil {
		return nil
	}
	return &dst
}

func (dao *UserActionDao) QueryUserCollection(openID string) *[]TbUserScenicCollection {
	query := &TbUserScenicCollection{OpenID: openID}
	var dst []TbUserScenicCollection
	err := dao.runner.FindExample(query, &dst)
	if err != nil {
		return nil
	}
	return &dst
}

