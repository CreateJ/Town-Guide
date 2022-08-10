package dao

import (
	"fmt"
	"github.com/tietang/dbx"
	"town-guide/base"
)

type TbUserScenicCollection struct {
	ID         int64  `db:"id,omitempty,unique"`
	OpenID     string `db:"open_id,unique"`
	ScenicID   int64  `db:"scenic_id,unique"`
	UpdateTime int64  `db:"update_time"`
	CreateTime int64  `db:"create_time"`
}

type TbUserScenicClock struct {
	ID         int64  `db:"id,omitempty,unique"`
	OpenID     string `db:"open_id,unique"`
	ScenicID   int64  `db:"scenic_id,unique"`
	UpdateTime int64  `db:"update_time"`
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

func (dao *UserActionDao) GetUserCollectionState(openID string, scenicID int64) bool {
	a := &TbUserScenicCollection{OpenID: openID, ScenicID: scenicID}
	ok, err := dao.runner.GetOne(a)
	if err != nil || !ok {
		fmt.Println("UserAction GetUserCollectionState:", err)
		return false
	}
	return a.ID > 0
}

func (dao *UserActionDao) GetUserClockState(openID string, scenicID int64) bool {
	a := &TbUserScenicClock{OpenID: openID, ScenicID: scenicID}
	ok, err := dao.runner.GetOne(a)
	if err != nil || !ok {
		fmt.Println("UserAction GetUserClockState:", err)
		return false
	}
	return a.ID > 0
}

func (dao *UserActionDao) InsertCollection(a *TbUserScenicCollection) (id int64, err error) {
	rs, err := dao.runner.Insert(a)
	if err != nil {
		fmt.Println("UserAction InsertCollection:", err)
		return 0, err
	}
	return rs.LastInsertId()
}

func (dao *UserActionDao) InsertClock(a *TbUserScenicClock) (id int64, err error) {
	rs, err := dao.runner.Insert(a)
	if err != nil {
		fmt.Println("UserAction InsertClock:", err)
		return 0, err
	}
	return rs.LastInsertId()
}

func (dao *UserActionDao) QueryUserClock(openID string) *[]TbUserScenicClock {
	query := &TbUserScenicClock{OpenID: openID}
	var dst []TbUserScenicClock
	err := dao.runner.FindExample(query, &dst)
	if err != nil {
		fmt.Println("UserAction QueryUserClock:", err)
		return nil
	}
	return &dst
}

func (dao *UserActionDao) QueryUserCollection(openID string) *[]TbUserScenicCollection {
	query := &TbUserScenicCollection{OpenID: openID}
	var dst []TbUserScenicCollection
	err := dao.runner.FindExample(query, &dst)
	if err != nil {
		fmt.Println("UserAction QueryUserCollection:", err)
		return nil
	}
	return &dst
}

func (dao *UserActionDao) QueryUserScenicCollection(openID string, scenicID int64) *TbUserScenicCollection {
	query := &TbUserScenicCollection{OpenID: openID, ScenicID: scenicID}
	ok, err := dao.runner.GetOne(query)
	if err != nil || !ok {
		fmt.Println("UserAction QueryUserScenicCollection:", err)
		return nil
	}
	return query
}

func (dao *UserActionDao) QueryUserScenicClock(openID string, scenicID int64) *TbUserScenicClock {
	query := &TbUserScenicClock{OpenID: openID, ScenicID: scenicID}
	ok, err := dao.runner.GetOne(query)
	if err != nil || !ok {
		fmt.Println("UserAction QueryUserScenicClock:", err)
		return nil
	}
	return query
}

func (dao *UserActionDao) DeleteUserCollection(openID string, scenicID int64) error {
	sql := "DELETE FROM tb_user_scenic_collection WHERE open_id='" + openID + "' and scenic_id = " + fmt.Sprintf("%d", scenicID)
	_, err := dao.runner.Exec(sql)
	if err != nil {
		fmt.Println("UserAction DeleteUserCollection:", err)
		return err
	}
	return nil
}
