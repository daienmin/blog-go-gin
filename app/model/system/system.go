package system

import (
	"blog-go-gin/lib/db"
	"fmt"
	"blog-go-gin/lib/helper"
)

type Config struct {
	Id        uint32 `db:"id"`
	Name      string `db:"name"`
	Value     string `db:"value"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func GetConfig() map[string]string {
	sqlStr := "SELECT * FROM configs"
	Db := db.GetDb()
	res := map[string]string{
		"c_title":       "",
		"c_keywords":    "",
		"c_description": "",
		"c_copy_icp":    "",
		"c_stat_code":   "",
	}
	var conf []Config
	err := Db.Select(&conf, sqlStr)
	if err != nil {
		fmt.Printf("select configs err: %#v\n", err)
		return res
	}
	for _, v := range conf {
		res[v.Name] = v.Value
	}
	return res
}

func SaveData(data map[string]string) error {
	Db := db.GetDb()
	sqlStr := "INSERT INTO configs(name,value,created_at,updated_at) VALUES(?,?,?,?) ON DUPLICATE KEY UPDATE value=?,updated_at=?;"
	dateTime := helper.GetDateTime()
	for i, v := range data {
		Db.Exec(sqlStr, i, v, dateTime, dateTime, v, dateTime)
	}
	return nil
}
