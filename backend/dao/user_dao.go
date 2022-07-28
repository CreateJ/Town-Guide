package dao

import (
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"town-guide/base"
)

type UserInfo struct {
	Id         int64  `db:"id,omitempty"`
	OpenID     string `db:"open_id,uni"`
	NickName   string `db:"nick_name"`
	Url        string `db:"url"`
	Gender     int8   `db:"gender"`
	CreateTime int64  `db:"create_time"`
}
type userDao struct {
	runner *dbx.Database
}

func GetUserDao() userDao {
	return userDao{
		runner: base.DbxDatabase(),
	}
}

func (dao *userDao) GetOne(openID string) *UserInfo {
	a := &UserInfo{OpenID: openID}
	ok, err := dao.runner.GetOne(a)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return a
}

func (dao *userDao) Insert(a *UserInfo) (id int64, err error) {
	rs, err := dao.runner.Insert(a)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}
