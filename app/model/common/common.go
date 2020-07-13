package common

import (
	"blog-go-gin/app/model/category"
	"blog-go-gin/app/model/tag"
)

func GetArtCate(id uint8) string {
	cate := category.GetOne(int(id))
	return cate.Name
}

func GetArtTags(id uint32) string {
	tags := tag.GetOne(int(id))
	return tags.Name
}

func Mod(k, v int) int {
	return k % v
}