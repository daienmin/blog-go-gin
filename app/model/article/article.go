package article

import (
	"fmt"
	"blog-go-gin/lib/db"
	"blog-go-gin/lib/helper"
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

type ArtTags struct {
	ArticleId uint32 `db:"article_id" json:"article_id"`
	TagId     uint32 `db:"tag_id" json:"tag_id"`
	CreatedAt  string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

func GetList(offset, pageSize int) []Article {
	Db := db.GetDb()
	selectSql := fmt.Sprintf("SELECT * FROM articles ORDER BY id desc LIMIT %d,%d", offset, pageSize)
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
	countSql := "SELECT COUNT(*) as num FROM articles ORDER BY id desc"
	totalRows = 0
	err := Db.Get(&totalRows, countSql)
	if err != nil {
		fmt.Printf("count article err:%#v\n", err)
		return
	}
	return
}

func GetOne(id int) (art Article) {
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
	sqlStr = "DELETE FROM article_tags WHERE article_id=?"
	_, err = Db.Exec(sqlStr, id)
	return err
}

func InsertData(art *Article, tagIds []string) error {
	sqlStr := `INSERT INTO articles(
				category_id,title,author,markdown,html,description,
				keywords,cover,is_top,is_recommend,status,created_at,updated_at) 
				VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)`
	createAt := helper.GetDateTime()
	Db := db.GetDb()
	res, err := Db.Exec(
		sqlStr, art.CategoryId, art.Title, art.Author, art.MarkDown, art.Html, art.Description, art.KeyWords, art.Cover,
		art.IsTop, art.IsRecommend, art.Status, createAt, createAt)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	sqlStr = "INSERT INTO article_tags(article_id,tag_id,created_at,updated_at) VALUES(?,?,?,?)"
	for _, v := range tagIds {
		Db.Exec(sqlStr, id, v, createAt, createAt)
	}
	return nil
}

func UpdateData(art *Article, tagIds []string, id string) error {
	sqlStr := `UPDATE articles SET
				category_id=?,title=?,author=?,markdown=?,html=?,description=?,
				keywords=?,cover=?,is_top=?,is_recommend=?,status=?,updated_at=? 
				WHERE id=?`
	updateAt := helper.GetDateTime()
	Db := db.GetDb()
	_, err := Db.Exec(
		sqlStr, art.CategoryId, art.Title, art.Author, art.MarkDown, art.Html, art.Description, art.KeyWords, art.Cover,
		art.IsTop, art.IsRecommend, art.Status, updateAt, id)
	if err != nil {
		return err
	}
	sqlStr = "DELETE FROM article_tags WHERE article_id=?"
	Db.Exec(sqlStr, id)
	sqlStr = "INSERT INTO article_tags(article_id,tag_id,created_at,updated_at) VALUES(?,?,?,?)"
	for _, v := range tagIds {
		Db.Exec(sqlStr, id, v, updateAt, updateAt)
	}
	return nil
}

func GetTags(aid uint32) []ArtTags {
	Db := db.GetDb()
	var aTags []ArtTags
	sqlStr := "SELECT * FROM article_tags WHERE article_id=?"
	err := Db.Select(&aTags, sqlStr, aid)
	if err != nil {
		fmt.Printf("select from article_tags err:%#v\n", err)
		return []ArtTags{}
	}
	return aTags
}