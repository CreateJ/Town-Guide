package dao

import (
	"errors"
	"fmt"
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

func (dao *ScenicDao) QueryAll() *[]TbScenicInfo {
	query := &TbScenicInfo{}
	var dst []TbScenicInfo
	err := dao.runner.FindExample(query, &dst)
	if err != nil {
		return nil
	}
	return &dst
}

func (dao *ScenicDao) QueryOne(id int64) *TbScenicInfo {
	a := &TbScenicInfo{ID: id}
	ok, err := dao.runner.GetOne(a)
	if err != nil {
		return nil
	}
	if !ok {
		return nil
	}
	return a
}

func (dao *ScenicDao) DeleteOne(id int64) error {
	sql := "DELETE FROM tb_scenic_info WHERE id=" + fmt.Sprintf("%d", id)
	_, err := dao.runner.Exec(sql)
	if err != nil {
		return err
	}
	return nil

}

func (dao *ScenicDao) Insert(a *TbScenicInfo) (id int64, err error) {
	rs, err := dao.runner.Insert(a)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}

func (dao *ScenicDao) Edit(a *TbScenicInfo) (id int64, err error) {
	if a.ID <= 0 {
		return 0, errors.New("id err")
	}
	sql := " update tb_scenic_info set"
	if a.Name != "" {
		sql += " name=" + a.Name
	}

	if a.LocationDesc != "" {
		sql += " location_desc=" + a.LocationDesc
	}

	if a.Intro != "" {
		sql += " intro=" + a.Intro
	}

	if a.Description != "" {
		sql += " description=" + a.Description
	}

	if a.PicUrl != "" {
		sql += " pic_url=" + a.PicUrl
	}

	if a.Icon != "" {
		sql += " icon=" + a.Icon
	}
	if a.VideoUrl != "" {
		sql += " video_url=" + a.VideoUrl
	}
	if a.Tag != "" {
		sql += " tag=" + a.Tag
	}
	if a.OpenTime != "" {
		sql += " open_time=" + a.OpenTime
	}

	if a.VideoUrl != "" {
		sql += " video_url=" + a.VideoUrl
	}

	sql += " where id=" + fmt.Sprintf("%d", a.ID)
	rs, err := dao.runner.Exec(sql)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}

func (dao *ScenicDao) IncrCheckNum(id int64) (num int64, err error) {
	if id <= 0 {
		return 0, errors.New("id err")
	}
	sql := "UPDATE tb_scenic_info SET check_num=check_num+1" +
		" where id=" + fmt.Sprintf("%d", id)
	rs, err := dao.runner.Exec(sql)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}
