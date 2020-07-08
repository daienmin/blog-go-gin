package tag

import (
	"fmt"
	"blog-go-gin/lib/db"
	"blog-go-gin/lib/helper"
)

type Tag struct {
	Id          uint32 `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	KeyWords    string `db:"keywords" json:"keywords"`
	Description string `db:"description" json:"description"`
	CreatedAt   string `db:"created_at" json:"created_at"`
	UpdatedAt   string `db:"updated_at" json:"updated_at"`
}

func GetList() []Tag {
	sqlStr := "SELECT * FROM tags"
	Db := db.GetDb()
	var tags []Tag
	err := Db.Select(&tags, sqlStr)
	if err != nil {
		fmt.Printf("select err:%#v\n", err)
		return []Tag{}
	}
	return tags
}

func InsertData(name,keywords,description string) error {
	createdAt := helper.GetDateTime()
	updatedAt := createdAt
	Db := db.GetDb()
	sqlStr := "INSERT INTO tags(name,keywords,description,created_at,updated_at) VALUES(?,?,?,?,?)"
	_, err := Db.Exec(sqlStr, name, keywords, description, createdAt, updatedAt)
	return err
}

func GetOne(id int) Tag {
	sqlStr := "SELECT * FROM tags WHERE id=?"
	Db := db.GetDb()
	var tagData Tag
	err := Db.Get(&tagData, sqlStr, id)
	if err != nil {
		fmt.Printf("query err:%#v\n", err)
		return Tag{}
	}
	return tagData
}

func UpdateData(name, keywords, description, id string) error {
	updatedAt := helper.GetDateTime()
	Db := db.GetDb()
	sqlStr := "UPDATE tags SET name=?,keywords=?,description=?,updated_at=? WHERE id=?"
	_, err := Db.Exec(sqlStr, name, keywords, description, updatedAt, id)
	return err
}

func DeleteData(id int) error {
	sqlStr := "DELETE FROM tags WHERE id=?"
	Db := db.GetDb()
	_, err := Db.Exec(sqlStr, id)
	return err
}