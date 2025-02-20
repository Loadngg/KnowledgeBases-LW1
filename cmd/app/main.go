package main

import (
	"lr1/internal/app/design"
	"lr1/internal/app/parser"
	"lr1/internal/app/repository"
	"lr1/internal/constants"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow(constants.WindowName.String())

	r := repository.New("Болезни кожи/Правила.txt")
	p := parser.NewChainParser(r)
	d := design.MustLoad(p, w)

	w.SetContent(d)
	w.ShowAndRun()
}
