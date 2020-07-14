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
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

type IndexList struct {
	Article Article
	Tags    []ArtTags
}

type TagStat struct {
	Num  uint32 `db:"num" json:"num"`
	Id   uint32 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type OptArt struct {
	PrefArt Article
	NextArt Article
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
	countSql := "SELECT COUNT(*) as num FROM articles"
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

// 获取文章标签
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

// 首页最新文章列表
func GetIndexList(nums int) []IndexList {
	Db := db.GetDb()
	// 获取文章列表
	selectSql := fmt.Sprintf("SELECT * FROM articles WHERE status=1 ORDER BY id desc LIMIT %d", nums)
	var arts []Article
	err := Db.Select(&arts, selectSql)
	if err != nil {
		fmt.Printf("select articles err:%#v\n", err)
		return []IndexList{}
	}
	var res []IndexList
	for _, v := range arts {
		artTags := GetTags(v.Id)
		tmp := IndexList{
			Article: v,
			Tags:    artTags,
		}
		res = append(res, tmp)
	}
	return res
}

// 标签云
func GetTagsStat() []TagStat {
	sqlStr := `SELECT count(a.tag_id) as num,c.id,c.name FROM article_tags a 
				LEFT JOIN articles b ON a.article_id = b.id 
				LEFT JOIN tags c ON a.tag_id = c.id 
				WHERE b.status=1 GROUP BY a.tag_id`
	Db := db.GetDb()
	var data []TagStat
	err := Db.Select(&data, sqlStr)
	if err != nil {
		fmt.Printf("get tags stat err:%#v\n", err)
		return []TagStat{}
	}
	return data
}

// 查看排行、置顶、推荐文章列表
func GetArtByParam(gType int) []Article {
	Db := db.GetDb()
	var selectSql string
	switch gType {
	case 0: // 热门查看排行
		selectSql = "SELECT * FROM articles WHERE status=1 ORDER BY views desc LIMIT 8"
	case 1: // 置顶文章
		selectSql = "SELECT * FROM articles WHERE status=1 AND is_top=1 ORDER BY id desc LIMIT 8"
	case 2: // 推荐文章
		selectSql = "SELECT * FROM articles WHERE status=1 AND is_recommend=1 ORDER BY id desc LIMIT 8"
	}
	var arts []Article
	err := Db.Select(&arts, selectSql)
	if err != nil {
		fmt.Printf("select articles err:%#v\n", err)
		return []Article{}
	}
	return arts
}

// 文章详情页获取文章内容
func GetArtInfo(id string) (art Article) {
	sqlStr := "SELECT * FROM articles WHERE id=? AND status=1"
	Db := db.GetDb()
	err := Db.Get(&art, sqlStr, id)
	if err != nil {
		fmt.Printf("query err:%#v\n", err)
		return
	}
	return
}

// 文章详情获取上一篇下一篇文章
func GetOptArt(id uint32) OptArt {
	sqlStr := "SELECT * FROM articles WHERE id > ? AND status=1 ORDER BY id ASC LIMIT 1"
	var prefArt Article
	Db := db.GetDb()
	err := Db.Get(&prefArt, sqlStr, id)
	if err != nil {
		fmt.Printf("get prefix article err: %#v\n", err)
	}
	var nextArt Article
	sqlStr = "SELECT * FROM articles WHERE id < ? AND status=1 ORDER BY id DESC LIMIT 1"
	err = Db.Get(&nextArt, sqlStr, id)
	if err != nil {
		fmt.Printf("get next article err: %#v\n", err)
	}
	return OptArt{
		PrefArt: prefArt,
		NextArt: nextArt,
	}
}

// 分类列表文章总数
func GetCateArtCount(id string) (totalRows int) {
	Db := db.GetDb()
	countSql := "SELECT COUNT(*) as num FROM articles WHERE category_id=? AND status=1"
	totalRows = 0
	err := Db.Get(&totalRows, countSql, id)
	if err != nil {
		fmt.Printf("count article err:%#v\n", err)
		return
	}
	return
}

// 分类列表页
func GetCateArtList(id string, offset, pageSize int) []IndexList {
	Db := db.GetDb()
	// 获取文章列表
	selectSql := fmt.Sprintf("SELECT * FROM articles WHERE category_id=%s AND status=1 ORDER BY id desc LIMIT %d,%d", id, offset, pageSize)
	var arts []Article
	err := Db.Select(&arts, selectSql)
	if err != nil {
		fmt.Printf("select articles err:%#v\n", err)
		return []IndexList{}
	}
	var res []IndexList
	for _, v := range arts {
		artTags := GetTags(v.Id)
		tmp := IndexList{
			Article: v,
			Tags:    artTags,
		}
		res = append(res, tmp)
	}
	return res
}

func GetTagArtCount(id string) (totalRows int) {
	Db := db.GetDb()
	countSql := "SELECT COUNT(*) as num FROM article_tags a LEFT JOIN articles b ON a.article_id=b.id WHERE tag_id=? AND b.status=1"
	totalRows = 0
	err := Db.Get(&totalRows, countSql, id)
	if err != nil {
		fmt.Printf("count tag article err:%#v\n", err)
		return
	}
	return
}

// 分类列表页
func GetTagArtList(id string, offset, pageSize int) []IndexList {
	Db := db.GetDb()
	// 获取文章列表
	selectSql := fmt.Sprintf("SELECT b.* FROM article_tags a LEFT JOIN articles b ON a.article_id=b.id WHERE a.tag_id=%s AND b.status=1 ORDER BY b.id desc LIMIT %d,%d", id, offset, pageSize)
	var arts []Article
	err := Db.Select(&arts, selectSql)
	if err != nil {
		fmt.Printf("select articles err:%#v\n", err)
		return []IndexList{}
	}
	var res []IndexList
	for _, v := range arts {
		artTags := GetTags(v.Id)
		tmp := IndexList{
			Article: v,
			Tags:    artTags,
		}
		res = append(res, tmp)
	}
	return res
}