package paginator

import (
	"fmt"
	"testing"
)

func TestPaginator(t *testing.T) {
	p := InitPages()
	p.SetBaseUrl("http://daiem.com")
	p.SetTotalRows(50)
	p.SetCurrentPage(5)
	html := p.Create()
	fmt.Println(html)
}