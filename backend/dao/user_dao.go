package dao

import (
	"fmt"
	"github.com/tietang/dbx"
	"town-guide/base"
)

type TbUserInfo struct {
	Id         int64  `db:"id,omitempty,unique"`
	OpenID     string `db:"open_id,unique"`
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

func (dao *userDao) GetOne(openID string) *TbUserInfo {
	a := &TbUserInfo{OpenID: openID}
	ok, err := dao.runner.GetOne(a)
	if err != nil || !ok {
		fmt.Println("User GetOne:", err)
		return nil
	}
	return a
}

func (dao *userDao) Insert(a *TbUserInfo) (id int64, err error) {
	rs, err := dao.runner.Insert(a)
	if err != nil {
		fmt.Println("User Insert:", err)
		return 0, err
	}
	return rs.LastInsertId()
}
