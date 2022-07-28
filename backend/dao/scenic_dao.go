package dao

import (
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"town-guide/base"
)

type TbScenicInfo struct {
	ID           int64  `db:"id,omitempty"`
	Name         string `db:"name"`
	LocationDesc string `db:"location_desc"`
	Description  string `db:"description"`
	Intro        string `db:"intro"`
	PicUrl       string `db:"pic_url"`
	Icon         string `db:"icon"`
	VideoUrl     string `db:"video_url"`
	Tag          string `db:"tag"`
	OpenTime     string `db:"open_time"`
	CheckNum     int64  `db:"check_num"`
	CreateTime   int64  `db:"create_time"`
}
type ScenicDao struct {
	runner *dbx.Database
}

func GetScenicDao() ScenicDao {
	return ScenicDao{
		runner: base.DbxDatabase(),
	}
}

func (dao *ScenicDao) GetOne(id int64) *TbScenicInfo {
	a := &TbScenicInfo{ID: id}
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

func (dao *ScenicDao) Insert(a *TbScenicInfo) (id int64, err error) {
	rs, err := dao.runner.Insert(a)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}
