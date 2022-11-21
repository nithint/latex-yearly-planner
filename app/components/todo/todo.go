package todo

import (
	"fmt"
	"strconv"

	"github.com/kudrykv/latex-yearly-planner/app/components/header"
	"github.com/kudrykv/latex-yearly-planner/app/components/hyper"
	"github.com/kudrykv/latex-yearly-planner/app/tex"
)

type Todos []*Todo
type Todo struct {
	Year   int
	Number int
	Page   int
}

func NewTodo(year, page, number int) *Todo {
	return &Todo{Year: year, Page: page, Number: number}
}

func (p Todos) Breadcrumb(year, idx int) string {
	postfix := ""
	if idx > 0 {
		postfix = " " + strconv.Itoa(idx+1)
	}

	return header.Items{
		header.NewIntItem(year),
		header.NewTextItem("Todos Index" + postfix).Ref(true),
	}.Table(true)
}

func (p Todos) HeadingMOS(page, pages int) string {
	var out string

	if page > 1 {
		out += tex.Hyperlink(p.ref(page-1), tex.ResizeBoxW(`\myLenHeaderResizeBox`, `$\langle$`)) + " "
	}

	out += tex.Hypertarget(p.ref(page), "") + tex.ResizeBoxW(`\myLenHeaderResizeBox`, `Index Todos`)

	if page < pages {
		out += " " + tex.Hyperlink(p.ref(page+1), tex.ResizeBoxW(`\myLenHeaderResizeBox`, `$\rangle$`))
	}

	return out
}

func (p Todos) ref(page int) string {
	var suffix string

	if page > 1 {
		suffix = " " + strconv.Itoa(page)
	}

	return "Todos Index" + suffix
}

func (n Todo) HyperLink() string {
	return hyper.Link(n.ref(), fmt.Sprintf("%02d", n.Number))
}

func (n Todo) Breadcrumb() string {
	page := ""

	if n.Page > 1 {
		page = " " + strconv.Itoa(n.Page)
	}

	return header.Items{
		header.NewIntItem(n.Year),
		header.NewTextItem("Todos Index" + page),
		header.NewTextItem(n.ref()).Ref(true),
	}.Table(true)
}

func (n Todo) PrevNext(todos int) header.Items {
	items := header.Items{}

	if n.Number > 1 {
		items = append(items, header.NewTextItem("Todo "+strconv.Itoa(n.Number-1)))
	}

	if n.Number < todos {
		items = append(items, header.NewTextItem("Todo "+strconv.Itoa(n.Number+1)))
	}

	return items
}

func (n Todo) HeadingMOS(page int) string {
	num := strconv.Itoa(n.Number)

	return tex.Hypertarget(n.ref(), "") + tex.ResizeBoxW(`\myLenHeaderResizeBox`, `Todo `+num+`\myDummyQ`)
}

func (n Todo) ref() string {
	return "Todo " + strconv.Itoa(n.Number)
}
