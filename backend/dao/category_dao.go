package dao

import (
	"fmt"
	"github.com/tietang/dbx"
	"town-guide/base"
)

type TbCategory struct {
	ID           int64  `db:"id,omitempty"`
	Name         string `db:"name"`
}

type CategoryDao struct {
	runner *dbx.Database
}

func GetCategoryDao() CategoryDao {
	return CategoryDao{
		runner: base.DbxDatabase(),
	}
}

func (dao *CategoryDao) QueryAll() *[]TbCategory {
	var dst []TbCategory
	err := dao.runner.Find(&dst, "SELECT * FROM tb_category")
	if err != nil {
		return nil
	}
	return &dst
}

func (dao *CategoryDao) QueryOne(id int64) *TbCategory {
	a := &TbCategory{ID: id}
	ok, err := dao.runner.GetOne(a)
	if err != nil || !ok {
		return nil
	}
	return a
}

func (dao *CategoryDao) DeleteOne(id int64) error {
	sql := "DELETE FROM tb_category WHERE id=" + fmt.Sprintf("%d", id)
	_, err := dao.runner.Exec(sql)
	if err != nil {
		return err
	}
	return nil

}

func (dao *CategoryDao) Insert(a *TbCategory) (id int64, err error) {
	rs, err := dao.runner.Insert(a)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}

