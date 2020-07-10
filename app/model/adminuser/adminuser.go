package adminuser

import (
	"blog-go-gin/lib/db"
	"fmt"
	"blog-go-gin/lib/helper"
)

type AdminUser struct {
	Id        uint32 `db:"id" json:"id"`
	UserName  string `db:"username" json:"username"`
	Password  string `db:"password" json:"password"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

func GetList() []AdminUser {
	sqlStr := "SELECT * FROM admin_users ORDER BY id DESC"
	Db := db.GetDb()
	var list []AdminUser
	err := Db.Select(&list, sqlStr)
	if err != nil {
		fmt.Printf("select from links err:%#v\n", err)
		return []AdminUser{}
	}
	return list
}

func InsertData(username, password string) error {
	createdAt := helper.GetDateTime()
	updatedAt := createdAt
	Db := db.GetDb()
	sqlStr := "INSERT INTO admin_users(username,password,created_at,updated_at) VALUES(?,?,?,?)"
	_, err := Db.Exec(sqlStr, username, password, createdAt, updatedAt)
	return err
}

func GetOne(id int) AdminUser {
	sqlStr := "SELECT * FROM admin_users WHERE id=?"
	Db := db.GetDb()
	var data AdminUser
	err := Db.Get(&data, sqlStr, id)
	if err != nil {
		fmt.Printf("query err:%#v\n", err)
		return AdminUser{}
	}
	return data
}

func UpdateData(username, password, id string) error {
	updatedAt := helper.GetDateTime()
	Db := db.GetDb()
	sqlStr := "UPDATE admin_users SET username=?,password=?,updated_at=? WHERE id=?"
	_, err := Db.Exec(sqlStr, username, password, updatedAt, id)
	return err
}

func DeleteData(id int) error {
	sqlStr := "DELETE FROM admin_users WHERE id=?"
	Db := db.GetDb()
	_, err := Db.Exec(sqlStr, id)
	return err
}
