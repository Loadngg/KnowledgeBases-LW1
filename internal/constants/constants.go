package constants

type TextEnum int

const (
	WindowName TextEnum = iota
	Tab1Name
	Tab2Name
	FindBtn
	DiseaseSelect
	NotCheckedSymptoms
	NotSelectedDisease
	ErrorTitle
	DiagnoseLabel
	DiagnoseIsCorrect
	HistoryLabel
)

var textLabels = map[TextEnum]string{
	WindowName:         "Базы знаний Лр1",
	Tab1Name:           "Найти болезнь",
	Tab2Name:           "Подтвердить гипотезу",
	FindBtn:            "Найти",
	DiseaseSelect:      "Выберите болезнь",
	NotCheckedSymptoms: "Вы не выбрали ни один из симптомов",
	NotSelectedDisease: "Вы не выбрали болезнь",
	ErrorTitle:         "Ошибка",
	DiagnoseLabel:      "Диагностированная болезнь: ",
	DiagnoseIsCorrect:  "Корректная ли гипотеза?: ",
	HistoryLabel:       "Цепочка гипотезы:\n",
}

func (e TextEnum) String() string {
	if val, ok := textLabels[e]; ok {
		return val
	}
	return "Неизвестный ключ"
}
