package dao

import (
	"errors"
	"fmt"
	"github.com/tietang/dbx"
	"strings"
	"town-guide/base"
)

type TbCategory struct {
	ID         int64  `db:"id,omitempty,unique"`
	Name       string `db:"name"`
	Icon       string `db:"icon"`
	IconActive string `db:"icon_active"`
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
		fmt.Println("Category QueryAll:", err)
		return nil
	}
	return &dst
}

func (dao *CategoryDao) QueryOne(id int64) *TbCategory {
	a := &TbCategory{ID: id}
	ok, err := dao.runner.GetOne(a)
	if err != nil || !ok {
		fmt.Println("Category QueryOne:", err)
		return nil
	}
	return a
}

func (dao *CategoryDao) DeleteOne(id int64) error {
	sql := "DELETE FROM tb_category WHERE id=" + fmt.Sprintf("%d", id)
	_, err := dao.runner.Exec(sql)
	if err != nil {
		fmt.Println("Category DeleteOne:", err)
		return err
	}
	return nil

}

func (dao *CategoryDao) Insert(a *TbCategory) (id int64, err error) {
	rs, err := dao.runner.Insert(a)
	if err != nil {
		fmt.Println("Category Insert:", err)
		return 0, err
	}
	return rs.LastInsertId()
}

func (dao *CategoryDao) Edit(a *TbCategory) (id int64, err error) {
	if a.ID <= 0 {
		return 0, errors.New("id err")
	}
	sql := "UPDATE tb_category SET "
	var params []interface{}
	var conditions []string
	if a.Name != "" {
		conditions = append(conditions, " name=?")
		params = append(params, a.Name)
	}

	if a.Icon != "" {
		conditions = append(conditions, " icon=?")
		params = append(params, a.Icon)
	}

	if len(conditions) == 0 {
		return 0, nil
	}
	strings.Join(conditions, ",")
	sql += strings.Join(conditions, ",") + " where id=" + fmt.Sprintf("%d", a.ID)
	rs, err := dao.runner.Exec(sql, params...)
	if err != nil {
		fmt.Println("Category Edit:", err)
		return 0, err
	}
	return rs.RowsAffected()
}
