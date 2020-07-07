package paginator

import (
	"math"
	"fmt"
)

type paginator struct {
	BaseUrl       string
	TotalRows     int32
	ListRows      int32
	NumLinks      int32
	CurrentPage   int32
	FirstLinkName string
	LastLinkName  string
	NextLinkName  string
	PrevLinkName  string
}

func InitPages() *paginator {
	Pr := &paginator{
		ListRows:      10,
		NumLinks:      3,
		FirstLinkName: "首页",
		LastLinkName:  "末页",
		NextLinkName:  "下一页",
		PrevLinkName:  "上一页",
	}
	return Pr
}

func (p *paginator) SetBaseUrl(baseUrl string) {
	p.BaseUrl = baseUrl
}

func (p *paginator) SetTotalRows(total int32) {
	p.TotalRows = total
}

func (p *paginator) SetListRows(lRow int32) {
	p.ListRows = lRow
}

func (p *paginator) SetNumLinks(nLinks int32) {
	p.NumLinks = nLinks
}

func (p *paginator) SetCurrentPage(cPage int32) {
	p.CurrentPage = cPage
}

func (p *paginator) SetFirstLinkName(name string) {
	p.FirstLinkName = name
}

func (p *paginator) SetLastLinkName(name string) {
	p.LastLinkName = name
}

func (p *paginator) SetNextLinkName(name string) {
	p.NextLinkName = name
}

func (p *paginator) SetPrevLinkName(name string) {
	p.PrevLinkName = name
}

func (p *paginator) Create() string {
	// 总页数
	numPages := math.Ceil(float64(p.TotalRows / p.ListRows))

	var start int32 = 1
	if p.CurrentPage - p.NumLinks > 0 {
		start = p.CurrentPage - p.NumLinks - int32(1)
	}
	var end = int32(numPages)
	if p.CurrentPage + p.NumLinks < int32(numPages) {
		end = p.CurrentPage + p.NumLinks
	}

	html := fmt.Sprintf("<ul class='pagination'><li><a href='%s'>%s</a></li>", p.BaseUrl, p.FirstLinkName)
	if p.CurrentPage > 1 {
		html += fmt.Sprintf("<li><a href='%s/%d'>%s</a></li>", p.BaseUrl, p.CurrentPage-1, p.PrevLinkName)
	}

	for i := start - 1; i <= end; i++ {
		if i == 0 {
			continue
		}
		if i == p.CurrentPage {
			html += fmt.Sprintf("<li class='active'><span>%d</span></li>", i)
		} else {
			html += fmt.Sprintf("<li><a href='%s/%d'>%d</a></li>", p.BaseUrl, i, i)
		}
	}

	if float64(p.CurrentPage) < numPages {
		html += fmt.Sprintf("<li><a href='%s/%d'>%s</a></li>", p.BaseUrl, p.CurrentPage+1, p.NextLinkName)
	}
	html += fmt.Sprintf("<li><a href='%v/%d'>%s</a></li>", p.BaseUrl, int64(numPages), p.LastLinkName)
	html += "</ul>"
	return html
}
