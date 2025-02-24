package design

import (
	"fmt"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"lr1/internal/app/parser"
	"lr1/internal/constants"
)

func ShowSymptomQuestion(symptom string, w fyne.Window, callback func(bool)) {
	dialog.ShowConfirm(
		constants.SymptomQuestionTitle.String(),
		fmt.Sprintf(constants.SymptomQuestionMsg.String(), symptom),
		func(answer bool) {
			callback(answer)
		},
		w,
	)
}

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
		if selectedDisease == "" || len(checkedSymptoms) == 0 {
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
			return
		}

		diagnoseLabel.SetText(constants.DiagnoseIsCorrect.String())
		historyLabel.SetText(constants.HistoryLabel.String())

		p.Backward.Parse(
			checkedSymptoms,
			selectedDisease,
			func(symptom string, callback func(bool)) {
				ShowSymptomQuestion(symptom, w, callback)
			},
			func(diagnose bool, history []string) {
				if diagnose {
					diagnoseLabel.SetText(constants.DiagnoseIsCorrect.String() + constants.Yes.String())
				} else {
					diagnoseLabel.SetText(constants.DiagnoseIsCorrect.String() + constants.No.String())
				}
				historyLabel.SetText(historyLabel.Text + strings.Join(history, "\n"))
			},
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
