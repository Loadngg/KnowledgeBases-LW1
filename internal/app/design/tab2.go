package design

import (
	"log"
	"strings"

	"lr1/internal/app/parser"
	"lr1/internal/constants"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func Tab2(p *parser.ChainParser, w fyne.Window) *fyne.Container {
	diseases, err := p.Repository.GetDiseases()
	if err != nil {
		log.Fatal(err)
	}

	var selectedDisease string
	combo := widget.NewSelect(*diseases, func(value string) {
		selectedDisease = value
	})
	combo.PlaceHolder = constants.DiseaseSelect.String()

	symptoms, err := p.Repository.GetSymptomsList()
	if err != nil {
		log.Fatal(err)
	}
	var checkedSymptoms []string
	symptomsBox := widget.NewCheckGroup(*symptoms, func(s []string) {
		checkedSymptoms = s
	})

	diagnoseLabel := widget.NewLabel(constants.DiagnoseIsCorrect.String())
	historyLabel := widget.NewLabel(constants.HistoryLabel.String())
	applyBtn := widget.NewButton(constants.FindBtn.String(), func() {
		if selectedDisease != "" && len(checkedSymptoms) > 0 {
			diagnoseLabel.SetText(constants.DiagnoseIsCorrect.String())
			historyLabel.SetText(constants.HistoryLabel.String())

			diagnose, history := p.Backward.Parse(checkedSymptoms, selectedDisease)

			var diagnoseStr string
			if diagnose {
				diagnoseStr = diagnoseLabel.Text + "Да"
			} else {
				diagnoseStr = diagnoseLabel.Text + "Нет"
			}
			diagnoseLabel.SetText(diagnoseStr)
			historyLabel.SetText(historyLabel.Text + strings.Join(history, "\n"))
			return
		}
		var errMsg string
		if selectedDisease == "" {
			errMsg = constants.NotSelectedDisease.String()
		}
		if len(checkedSymptoms) == 0 {
			errMsg = errMsg + "\n" + constants.NotCheckedSymptoms.String()
		}
		dialog.ShowInformation(
			constants.ErrorTitle.String(),
			errMsg,
			w,
		)
	})

	content := container.NewVBox(
		combo,
		symptomsBox,
		applyBtn,
		diagnoseLabel,
		historyLabel,
	)
	return content
}
