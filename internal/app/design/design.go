package design

import (
	"lr1/internal/app/parser"
	"lr1/internal/constants"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func MustLoad(p *parser.ChainParser, w fyne.Window) *container.AppTabs {
	tab1 := Tab1(p, w)
	tab2 := Tab2(p, w)

	tabs := container.NewAppTabs(
		container.NewTabItem(constants.Tab1Name.String(), tab1),
		container.NewTabItem(constants.Tab2Name.String(), tab2),
	)
	return tabs
}
