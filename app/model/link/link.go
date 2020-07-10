package link

import (
	"blog-go-gin/lib/db"
	"fmt"
	"blog-go-gin/lib/helper"
)

type Link struct {
	Id        uint32 `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Url       string `db:"url" json:"url"`
	Sort      uint8  `db:"sort" json:"sort"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

func GetList() []Link {
	sqlStr := "SELECT * FROM links ORDER BY id DESC"
	Db := db.GetDb()
	var links []Link
	err := Db.Select(&links, sqlStr)
	if err != nil {
		fmt.Printf("select from links err:%#v\n", err)
		return []Link{}
	}
	return links
}

func InsertData(name, url, sort string) error {
	createdAt := helper.GetDateTime()
	updatedAt := createdAt
	Db := db.GetDb()
	sqlStr := "INSERT INTO links(name,url,sort,created_at,updated_at) VALUES(?,?,?,?,?)"
	_, err := Db.Exec(sqlStr, name, url, sort, createdAt, updatedAt)
	return err
}

func GetOne(id int) Link {
	sqlStr := "SELECT * FROM links WHERE id=?"
	Db := db.GetDb()
	var data Link
	err := Db.Get(&data, sqlStr, id)
	if err != nil {
		fmt.Printf("query err:%#v\n", err)
		return Link{}
	}
	return data
}

func UpdateData(name, url, sort, id string) error {
	updatedAt := helper.GetDateTime()
	Db := db.GetDb()
	sqlStr := "UPDATE links SET name=?,url=?,sort=?,updated_at=? WHERE id=?"
	_, err := Db.Exec(sqlStr, name, url, sort, updatedAt, id)
	return err
}

func DeleteData(id int) error {
	sqlStr := "DELETE FROM links WHERE id=?"
	Db := db.GetDb()
	_, err := Db.Exec(sqlStr, id)
	return err
}