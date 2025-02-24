package constants

type TextEnum int

const (
	WindowName TextEnum = iota
	Tab1Name
	Tab2Name
	Yes
	No
	FindBtn
	DiseaseNotFound
	DiseaseSelect
	NotCheckedSymptoms
	NotSelectedDisease
	ErrorTitle
	DiagnoseLabel
	DiagnoseIsCorrect
	HistoryLabel
	SymptomQuestionTitle
	SymptomQuestionMsg
	IncorrectRuleFormat
)

var textLabels = map[TextEnum]string{
	WindowName:           "Базы знаний Лр1",
	Tab1Name:             "Найти болезнь",
	Tab2Name:             "Подтвердить гипотезу",
	Yes:                  "Да",
	No:                   "Нет",
	DiseaseNotFound:      "Болезнь не найдена",
	FindBtn:              "Найти",
	DiseaseSelect:        "Выберите болезнь",
	NotCheckedSymptoms:   "Вы не выбрали ни один из симптомов",
	NotSelectedDisease:   "Вы не выбрали болезнь",
	ErrorTitle:           "Ошибка",
	DiagnoseLabel:        "Диагностированная болезнь: ",
	DiagnoseIsCorrect:    "Корректная ли гипотеза?: ",
	HistoryLabel:         "Цепочка гипотезы:\n",
	SymptomQuestionTitle: "Наличие симптома",
	SymptomQuestionMsg:   "Есть ли симптом '%s'?",
	IncorrectRuleFormat:  "Неверный формат правила: %s",
}

func (e TextEnum) String() string {
	if val, ok := textLabels[e]; ok {
		return val
	}
	return "Неизвестный ключ"
}
