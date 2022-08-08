package dao

import (
	"errors"
	"fmt"
	"github.com/tietang/dbx"
	"strings"
	"time"
	"town-guide/base"
)

type TbScenicInfo struct {
	ID           int64  `db:"id,omitempty"`
	Name         string `db:"name"`
	LocationDesc string `db:"location_desc"`
	Description  string `db:"description"`
	Intro        string `db:"intro"`
	PicName      string `db:"pic_name"`
	Icon         string `db:"icon"`
	VideoName    string `db:"video_name"`
	Tag          string `db:"tag"`
	OpenTime     string `db:"open_time"`
	ClockNum     int64  `db:"clock_num"`
	Banner       string `db:"banner"`
	Location     string `db:"location"`
	CreateTime   int64  `db:"create_time"`
	UpdateTime   int64  `db:"update_time"`
	CategoryID   int64  `db:"category_id"`
	AudioName    string `db:"audio_name"`
	ClockIcon    string `db:"clock_icon"`
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
	var dst []TbScenicInfo
	err := dao.runner.Find(&dst, "SELECT * FROM tb_scenic_info ")
	if err != nil {
		return nil
	}
	return &dst
}

func (dao *ScenicDao) QueryByCategoryID(scenicID int64) *[]TbScenicInfo {
	var dst []TbScenicInfo
	err := dao.runner.Find(&dst, "SELECT * FROM tb_scenic_info WHERE Category_id="+fmt.Sprintf("%d", scenicID))
	if err != nil {
		return nil
	}
	return &dst
}

func (dao *ScenicDao) QueryOne(id int64) *TbScenicInfo {
	a := &TbScenicInfo{ID: id}
	ok, err := dao.runner.GetOne(a)
	if err != nil || !ok {
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
	sql := " update tb_scenic_info set update_time=" + fmt.Sprintf("%d", time.Now().Unix())
	var params []interface{}
	if a.Name != "" {
		sql += " , name=?"
		params = append(params, a.Name)
	}

	if a.LocationDesc != "" {
		sql += " , location_desc=?"
		params = append(params, a.LocationDesc)
	}

	if a.Intro != "" {
		sql += " , intro=?"
		params = append(params, a.Intro)
	}
	if a.ClockIcon != "" {
		sql += " , clock_icon=?"
		params = append(params, a.ClockIcon)
	}

	if a.Description != "" {
		sql += " , description=?"
		params = append(params, a.Description)
	}

	if a.PicName != "" {
		sql += " , pic_name=?"
		params = append(params, a.PicName)
	}

	if a.Icon != "" {
		sql += " , icon=?"
		params = append(params, a.Icon)
	}

	if a.AudioName != "" {
		sql += " , audio_name=?"
		params = append(params, a.AudioName)
	}

	if a.Location != "" {
		sql += " , location=?"
		params = append(params, a.Location)
	}

	if a.VideoName != "" {
		sql += " , video_name=?"
		params = append(params, a.VideoName)
	}

	if a.Tag != "" {
		sql += " , tag=?"
		params = append(params, a.Tag)
	}

	if a.OpenTime != "" {
		sql += " , open_time=?"
		params = append(params, a.OpenTime)
	}

	if a.Banner != "" {
		sql += " , banner=?"
		params = append(params, a.Banner)
	}

	if a.CategoryID <= 0 {
		sql += " , category_id=" + fmt.Sprintf("%d", a.CategoryID)
	}

	sql += " where id=" + fmt.Sprintf("%d", a.ID)
	rs, err := dao.runner.Exec(sql, params...)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}

func (dao *ScenicDao) UpdateScenicCategoryID(scenicIds []int64, categoryID int64) (err error) {
	if len(scenicIds) <= 0 {
		return nil
	}

	sql := " update tb_scenic_info set update_time=" + fmt.Sprintf("%d", time.Now().Unix()) +
		" , category_id=" + fmt.Sprintf("%d", categoryID)

	temp := make([]string, len(scenicIds))
	for k, v := range scenicIds {
		temp[k] = fmt.Sprintf("%d", v)
	}
	result := "(" + strings.Join(temp, ",") + ")"

	sql += " where id in " + result
	_, err = dao.runner.Exec(sql)
	return err
}

func (dao *ScenicDao) IncrClockNum(id int64) (num int64, err error) {
	if id <= 0 {
		return 0, errors.New("id err")
	}
	sql := "UPDATE tb_scenic_info SET clock_num=clock_num+1" +
		" where id=" + fmt.Sprintf("%d", id)
	rs, err := dao.runner.Exec(sql)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}
