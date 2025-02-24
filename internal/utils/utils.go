package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadFileLines(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func RuleTriggered(conditions []string, facts map[string]bool) bool {
	for _, condition := range conditions {
		if !facts[condition] {
			return false
		}
	}
	return true
}

func ReverseArray(arr []string) []string {
	reversed := make([]string, len(arr))
	for i := 0; i < len(arr); i++ {
		reversed[i] = arr[len(arr)-1-i]
	}

	return reversed
}

func HumanReadableRules(rules []string) []string {
	var formatted []string
	for _, rule := range rules {
		parts := strings.SplitN(rule, " ТО ", 2)
		if len(parts) != 2 {
			continue
		}

		conditionsPart := strings.TrimPrefix(parts[0], "ЕСЛИ ")
		conditions := strings.Split(conditionsPart, " И ")

		var cleanConditions []string
		for _, cond := range conditions {
			cond = strings.ToLower(cond)
			cleanCond := strings.Replace(cond, " = да", "", 1)
			cleanCond = strings.ReplaceAll(cleanCond, "_", " ")
			cleanCond = strings.ReplaceAll(cleanCond, "(", " ")
			cleanCond = strings.ReplaceAll(cleanCond, ")", " ")
			cleanCond = strings.Join(strings.Fields(cleanCond), " ")
			cleanConditions = append(cleanConditions, cleanCond)
		}

		resultPart := parts[1]
		resultType, resultValue, _ := strings.Cut(resultPart, " = ")
		resultType = strings.ReplaceAll(resultType, "_", " ")
		resultValue = strings.ReplaceAll(resultValue, "_", " ")

		var sb strings.Builder
		sb.WriteString("Если ")
		sb.WriteString(strings.Join(cleanConditions, " и "))

		switch resultType {
		case "болезнь":
			sb.WriteString(", то диагностирована болезнь «")
		case "факт":
			sb.WriteString(", то имеем факт: «")
		default:
			sb.WriteString(", то ")
		}

		sb.WriteString(resultValue)
		sb.WriteString("»")

		formatted = append(formatted, sb.String())
	}
	return formatted
}
