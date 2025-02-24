package design

import (
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"lr1/internal/app/parser"
	"lr1/internal/constants"
)

func Tab1(p *parser.ChainParser, w fyne.Window) *fyne.Container {
	symptoms, err := p.Repository.GetSymptomsList()
	if err != nil {
		log.Fatal(err)
	}
	var checkedSymptoms []string
	symptomsBox := widget.NewCheckGroup(*symptoms, func(s []string) {
		checkedSymptoms = s
	})

	diagnoseLabel := widget.NewLabel(constants.DiagnoseLabel.String())
	historyLabel := widget.NewLabel(constants.HistoryLabel.String())
	applyBtn := widget.NewButton(constants.FindBtn.String(), func() {
		if len(checkedSymptoms) == 0 {
			dialog.ShowInformation(
				constants.ErrorTitle.String(),
				constants.NotCheckedSymptoms.String(),
				w,
			)
			return
		}
		diagnoseLabel.SetText(constants.DiagnoseLabel.String())
		historyLabel.SetText(constants.HistoryLabel.String())

		diagnose, history := p.Forward.Parse(checkedSymptoms)

		diagnoseLabel.SetText(diagnoseLabel.Text + diagnose)
		historyLabel.SetText(historyLabel.Text + strings.Join(history, "\n"))
	})

	content := container.NewVBox(
		symptomsBox,
		applyBtn,
		diagnoseLabel,
		historyLabel,
	)
	return content
}
