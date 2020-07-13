package category

import (
	"fmt"
	"blog-go-gin/lib/db"
	"blog-go-gin/lib/helper"
)

type Category struct {
	Id          uint32 `db:"id" json:"id"`
	Pid         int32  `db:"pid" json:"pid"`
	Name        string `db:"name" json:"name"`
	KeyWords    string `db:"keywords" json:"keywords"`
	Description string `db:"description" json:"description"`
	Sort        uint8  `db:"sort" json:"sort"`
	Status      uint8  `db:"status" json:"status"`
	CreatedAt   string `db:"created_at" json:"created_at"`
	UpdatedAt   string `db:"updated_at" json:"updated_at"`
}

func GetList(visiable bool) []Category {
	sqlStr := "SELECT * FROM categories"
	if visiable {
		sqlStr += " WHERE status=1"
	}
	Db := db.GetDb()
	var categories []Category
	err := Db.Select(&categories, sqlStr)
	if err != nil {
		fmt.Printf("select err:%#v\n", err)
		return []Category{}
	}
	return categories
}

func InsertData(pid, name, keywords, description, sort, status string) error {
	createdAt := helper.GetDateTime()
	updatedAt := createdAt
	Db := db.GetDb()
	sqlStr := "INSERT INTO categories(pid,name,keywords,description,sort,status,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?)"
	_, err := Db.Exec(sqlStr, pid, name, keywords, description, sort, status, createdAt, updatedAt)
	return err
}

func GetOne(id int) Category {
	sqlStr := "SELECT * FROM categories WHERE id=?"
	Db := db.GetDb()
	var category Category
	err := Db.Get(&category, sqlStr, id)
	if err != nil {
		fmt.Printf("query err:%#v\n", err)
		return Category{}
	}
	return category
}

func UpdateData(pid, name, keywords, description, sort, status, id string) error {
	updatedAt := helper.GetDateTime()
	Db := db.GetDb()
	sqlStr := "UPDATE categories SET pid=?,name=?,keywords=?,description=?,sort=?,status=?,updated_at=? WHERE id=?"
	_, err := Db.Exec(sqlStr, pid, name, keywords, description, sort, status, updatedAt, id)
	return err
}

func DeleteData(id int) error {
	sqlStr := "DELETE FROM categories WHERE id=?"
	Db := db.GetDb()
	_, err := Db.Exec(sqlStr, id)
	return err
}

func GetArtCates(ids string) map[uint32]string {
	sqlStr := "SELECT * FROM categories WHERE id IN(?)"
	Db := db.GetDb()
	var cate []Category
	err := Db.Select(&cate, sqlStr, ids)
	if err != nil {
		fmt.Printf("query err:%#v\n", err)
		return make(map[uint32]string, 0)
	}
	res := make(map[uint32]string, 100)
	for _, v := range cate {
		res[v.Id] = v.Name
	}
	return res
}