package parser

import (
	"lr1/internal/app/repository"
	"lr1/internal/utils"
)

type BackwardParser struct {
	r *repository.Repository
}

func NewBackwardParser(r *repository.Repository) *BackwardParser {
	return &BackwardParser{
		r: r,
	}
}

func (p *BackwardParser) Parse(checkedSymptoms []string, diagnose string) (bool, []string) {
	facts := make(map[string]bool)
	for _, symptom := range checkedSymptoms {
		facts[symptom] = true
	}

	rules, err := p.r.GetRules()
	if err != nil {
		panic(err)
	}

	var history []string

	result := p.resolveBackward(diagnose, facts, &history, *rules)
	return result, utils.ReverseArray(history)
}

func (p *BackwardParser) resolveBackward(
	goal string,
	facts map[string]bool,
	history *[]string,
	rules []repository.Rule,
) bool {
	if facts[goal] {
		return true
	}

	for _, rule := range rules {
		if rule.Result == goal {
			allConditionsMet := true
			for _, condition := range rule.Conditions {
				if !facts[condition] && !p.resolveBackward(condition, facts, history, rules) {
					allConditionsMet = false
					break
				}
			}

			if allConditionsMet {
				*history = append(*history, rule.RuleStr)
				facts[goal] = true
				return true
			}
		}
	}

	return false
}
