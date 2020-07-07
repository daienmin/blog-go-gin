package article

import (
	"fmt"
	"blog-go-gin/lib/db"
)

type Article struct {
	Id          uint32 `db:"id" json:"id"`
	CategoryId  uint8  `db:"category_id" json:"category_id"`
	Title       string `db:"title" json:"title"`
	Author      string `db:"author" json:"author"`
	MarkDown    string `db:"markdown" json:"markdown"`
	Html        string `db:"html" json:"html"`
	Description string `db:"description" json:"description"`
	KeyWords    string `db:"keywords" json:"keywords"`
	Cover       string `db:"cover" json:"cover"`
	IsTop       uint8  `db:"is_top" json:"is_top"`
	IsRecommend uint8  `db:"is_recommend" json:"is_recommend"`
	Views       uint32 `db:"views" json:"views"`
	Status      uint8  `db:"status" json:"status"`
	CreatedAt   string `db:"created_at" json:"created_at"`
	UpdatedAt   string `db:"updated_at" json:"updated_at"`
}

func GetList(offset, pageSize int) []Article {
	Db := db.GetDb()
	selectSql := fmt.Sprintf("SELECT * FROM articles LIMIT %d,%d", offset, pageSize)
	var arts []Article
	err := Db.Select(&arts, selectSql)
	if err != nil {
		fmt.Printf("select articles err:%#v\n", err)
		return []Article{}
	}
	return arts
}

func GetCount() (totalRows int) {
	Db := db.GetDb()
	countSql := "SELECT COUNT(*) as num FROM articles"
	totalRows = 0
	err := Db.Get(&totalRows, countSql)
	if err != nil {
		fmt.Printf("count article err:%#v\n", err)
		return
	}
	return
}

func GetOne(id int) (art *Article) {
	sqlStr := "SELECT * FROM articles WHERE id=?"
	Db := db.GetDb()
	err := Db.Get(&art, sqlStr, id)
	if err != nil {
		fmt.Printf("query err:%#v\n", err)
		return
	}
	return
}

func DeleteData(id int) error {
	sqlStr := "DELETE FROM articles WHERE id=?"
	Db := db.GetDb()
	_, err := Db.Exec(sqlStr, id)
	return err
}