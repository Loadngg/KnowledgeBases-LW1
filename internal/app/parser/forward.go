package parser

import (
	"lr1/internal/app/repository"
	"lr1/internal/constants"
	"lr1/internal/utils"
)

type ForwardParser struct {
	r *repository.Repository
}

func NewForwardParser(r *repository.Repository) *ForwardParser {
	return &ForwardParser{
		r: r,
	}
}

func (p *ForwardParser) Parse(checkedSymptoms []string) (string, []string) {
	var facts = make(map[string]bool)
	for _, symptom := range checkedSymptoms {
		facts[symptom] = true
	}

	rules, err := p.r.GetRules()
	if err != nil {
		panic(err)
	}

	var history []string
	for {
		newFactAdded := false
		for _, rule := range *rules {
			if utils.RuleTriggered(rule.Conditions, facts) && !facts[rule.Result] {
				history = append(history, rule.RuleStr)
				facts[rule.Result] = true
				newFactAdded = true

				if rule.IsDisease {
					return rule.Result, utils.HumanReadableRules(history)
				}
			}
		}

		if !newFactAdded {
			break
		}
	}

	return constants.DiseaseNotFound.String(), nil
}
