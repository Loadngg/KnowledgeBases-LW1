package repository

import (
	"fmt"
	"strings"

	"lr1/internal/utils"
)

type DiseaseRepository struct {
	rulesFilepath    string
	symptomsFilepath string
}

type Rule struct {
	Conditions []string
	Result     string
	IsDisease  bool
	RuleStr    string
}

func NewDiseaseRepository(rulesFilepath string, symptomsFilepath string) *DiseaseRepository {
	return &DiseaseRepository{
		rulesFilepath:    rulesFilepath,
		symptomsFilepath: symptomsFilepath,
	}
}

func (r *DiseaseRepository) GetSymptomsList() (*[]string, error) {
	symptoms, err := utils.ReadFileLines(r.symptomsFilepath)
	if err != nil {
		return nil, err
	}
	return &symptoms, nil
}

func (r *DiseaseRepository) GetRules() (*[]Rule, error) {
	fileLines, err := utils.ReadFileLines(r.rulesFilepath)
	if err != nil {
		return nil, err
	}

	var rules []Rule
	for _, line := range fileLines {
		rule := parseRule(line)
		rules = append(rules, rule)
	}

	return &rules, nil
}

func (r *DiseaseRepository) GetDiseases() (*[]string, error) {
	rules, err := r.GetRules()
	if err != nil {
		return nil, err
	}

	var diseases []string
	for _, rule := range *rules {
		if rule.IsDisease {
			diseases = append(diseases, rule.Result)
		}
	}

	return &diseases, nil
}

func parseRule(ruleStr string) Rule {
	parts := strings.Split(strings.TrimSpace(ruleStr), " ТО ")
	if len(parts) != 2 {
		panic(fmt.Sprintf("Неверный формат правила: %s", ruleStr))
	}

	conditionPart := strings.TrimPrefix(parts[0], "ЕСЛИ ")
	conditions := strings.Split(conditionPart, " И ")
	for i, condition := range conditions {
		conditions[i] = strings.TrimSuffix(condition, " = да")
	}

	resultPart := strings.TrimSpace(parts[1])
	isDisease := strings.HasPrefix(resultPart, "болезнь =")
	var result string
	if isDisease {
		result = strings.TrimPrefix(resultPart, "болезнь = ")
	} else {
		result = strings.TrimPrefix(resultPart, "факт = ")
	}

	return Rule{
		Conditions: conditions,
		Result:     result,
		IsDisease:  isDisease,
		RuleStr:    ruleStr,
	}
}
