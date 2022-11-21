package todo

import (
	"strconv"

	"github.com/kudrykv/latex-yearly-planner/app/components/header"
)

type Pages []Todos

type Index struct {
	Pages Pages
}

func NewIndex(year, todosOnPage, pages int) *Index {
	pgs := make(Pages, 0, pages)

	for pageNum := 1; pageNum <= pages; pageNum++ {
		pg := make(Todos, 0, todosOnPage)

		for todoNum := 1; todoNum <= todosOnPage; todoNum++ {
			pg = append(pg, NewTodo(year, pageNum, (pageNum-1)*todosOnPage+todoNum))
		}

		pgs = append(pgs, pg)
	}

	return &Index{Pages: pgs}
}

func (i Index) PrevNext(currIdx int) header.Items {
	if len(i.Pages) <= 1 {
		return header.Items{}
	}

	list := header.Items{}

	if currIdx > 0 {
		postfix := " " + strconv.Itoa(currIdx)
		if currIdx == 1 {
			postfix = ""
		}

		list = append(list, header.NewTextItem("Todos Index"+postfix))
	}

	if currIdx+1 < len(i.Pages) {
		list = append(list, header.NewTextItem("Todos Index "+strconv.Itoa(currIdx+2)))
	}

	return list
}
