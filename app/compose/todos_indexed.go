package compose

import (
	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/components/todo"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

func TodosIndexed(cfg config.Config, tpls []string) (page.Modules, error) {
	index := todo.NewIndex(cfg.Year, cfg.Layout.Numbers.TodosOnPage, cfg.Layout.Numbers.TodosIndexPages)
	year := cal.NewYear(cfg.WeekStart, cfg.Year)
	modules := make(page.Modules, 0, 1)

	for idx, indexPage := range index.Pages {
		modules = append(modules, page.Module{
			Cfg: cfg,
			Tpl: tpls[0],
			Body: map[string]interface{}{
				"Todos":        indexPage,
				"Breadcrumb":   indexPage.Breadcrumb(cfg.Year, idx),
				"HeadingMOS":   indexPage.HeadingMOS(idx+1, len(index.Pages)),
				"SideQuarters": year.SideQuarters(0),
				"SideMonths":   year.SideMonths(0),
				"Extra":        index.PrevNext(idx).WithTopRightCorner(cfg.ClearTopRightCorner),
				"Extra2":       extra2(cfg.ClearTopRightCorner, false, true, nil, 0),
			},
		})
	}

	for idxPage, todos := range index.Pages {
		for _, nt := range todos {
			modules = append(modules, page.Module{
				Cfg: cfg,
				Tpl: tpls[1],
				Body: map[string]interface{}{
					"Todo":         nt,
					"Breadcrumb":   nt.Breadcrumb(),
					"HeadingMOS":   nt.HeadingMOS(idxPage),
					"SideQuarters": year.SideQuarters(0),
					"SideMonths":   year.SideMonths(0),
					"Extra": nt.
						PrevNext(cfg.Layout.Numbers.TodosOnPage * cfg.Layout.Numbers.TodosIndexPages).
						WithTopRightCorner(cfg.ClearTopRightCorner),
					"Extra2": extra2(cfg.ClearTopRightCorner, false, false, nil, idxPage+1),
				},
			})
		}
	}

	return modules, nil
}
