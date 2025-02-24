package parser

import (
	"maps"

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

func (p *BackwardParser) Parse(
	checkedSymptoms []string,
	diagnose string,
	showConfirm func(string, func(bool)),
	onComplete func(bool, []string),
) {
	facts := make(map[string]bool)
	for _, symptom := range checkedSymptoms {
		facts[symptom] = true
	}

	rules, err := p.r.GetRules()
	if err != nil {
		panic(err)
	}
	rulesList := *rules

	var history []string
	initialFacts := copyMap(facts)

	result := p.resolveBackward(diagnose, initialFacts, &history, rulesList)
	if result {
		onComplete(true, utils.HumanReadableRules(utils.ReverseArray(history)))
		return
	}

	missingSymptoms := p.findMissingSymptoms(diagnose, facts, rulesList)
	p.askForSymptoms(missingSymptoms, facts, diagnose, rulesList, history, showConfirm, onComplete)
}

func (p *BackwardParser) askForSymptoms(
	missingSymptoms []string,
	facts map[string]bool,
	diagnose string,
	rulesList []repository.Rule,
	history []string,
	showConfirm func(string, func(bool)),
	onComplete func(bool, []string),
) {
	if len(missingSymptoms) == 0 {
		onComplete(false, history)
		return
	}

	symptom := missingSymptoms[0]

	showConfirm(symptom, func(userConfirmed bool) {
		if userConfirmed {
			facts[symptom] = true

			newHistory := []string{}
			newResult := p.resolveBackward(diagnose, copyMap(facts), &newHistory, rulesList)
			if newResult {
				finalHistory := append(history, newHistory...)
				onComplete(true, utils.HumanReadableRules(utils.ReverseArray(finalHistory)))
				return
			}
		}

		p.askForSymptoms(missingSymptoms[1:], facts, diagnose, rulesList, history, showConfirm, onComplete)
	})
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
			localFacts := copyMap(facts)
			localHistory := make([]string, len(*history))
			copy(localHistory, *history)

			allConditionsMet := true
			for _, condition := range rule.Conditions {
				if !localFacts[condition] && !p.resolveBackward(condition, localFacts, &localHistory, rules) {
					allConditionsMet = false
					break
				}
			}

			if allConditionsMet {
				*history = append(localHistory, rule.RuleStr)
				facts[goal] = true
				return true
			}
		}
	}

	return false
}

func copyMap(original map[string]bool) map[string]bool {
	copy := make(map[string]bool)
	maps.Copy(copy, original)
	return copy
}

func (p *BackwardParser) findMissingSymptoms(diagnose string, facts map[string]bool, rules []repository.Rule) []string {
	visited := make(map[string]bool)
	terminals := p.getTerminalSymptoms(diagnose, facts, rules, visited)
	var missing []string
	for term := range terminals {
		if !facts[term] {
			missing = append(missing, term)
		}
	}
	return missing
}

func (p *BackwardParser) getTerminalSymptoms(
	goal string,
	facts map[string]bool,
	rules []repository.Rule,
	visited map[string]bool,
) map[string]bool {
	terminals := make(map[string]bool)
	if visited[goal] {
		return terminals
	}
	visited[goal] = true

	for _, rule := range rules {
		if rule.Result == goal {
			for _, cond := range rule.Conditions {
				if facts[cond] {
					continue
				}

				hasRule := false
				for _, r := range rules {
					if r.Result == cond {
						hasRule = true
						break
					}
				}

				if !hasRule {
					terminals[cond] = true
				} else {
					subTerminals := p.getTerminalSymptoms(cond, facts, rules, visited)
					for st := range subTerminals {
						terminals[st] = true
					}
				}
			}
		}
	}

	return terminals
}
