package dao

import (
	"fmt"
	"github.com/tietang/dbx"
	"town-guide/base"
)

type TbUserInfo struct {
	Id         int64  `db:"id,omitempty"`
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
	fmt.Println(*a)
	fmt.Print(ok)
	fmt.Print(err)

	if err != nil {
		return nil
	}
	if !ok {
		return nil
	}
	return a
}

func (dao *userDao) Insert(a *TbUserInfo) (id int64, err error) {
	rs, err := dao.runner.Insert(a)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}
