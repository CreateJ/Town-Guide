package dao

import (
	"errors"
	"fmt"
	"github.com/tietang/dbx"
	"strings"
	"town-guide/base"
)

type TbCategory struct {
	ID   int64  `db:"id,omitempty"`
	Name string `db:"name"`
	Icon string `db:"icon"`
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

func (dao *CategoryDao) Edit(a *TbCategory) (id int64, err error) {
	if a.ID <= 0 {
		return 0, errors.New("id err")
	}
	sql := " update tb_category set "
	var params []interface{}
	var sqls[]string
	if a.Name != "" {
		sqls = append(sqls," name=?")
		params = append(params, a.Name)
	}

	if a.Icon != "" {
		sqls = append(sqls," icon=?")
		params = append(params, a.Icon)
	}

	if len(sqls)==0 {
		return 0, nil
	}
	strings.Join(sqls, ",")
	sql += strings.Join(sqls, ",")+" where id=" + fmt.Sprintf("%d", a.ID)
	rs, err := dao.runner.Exec(sql, params...)
	fmt.Println(err)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}
