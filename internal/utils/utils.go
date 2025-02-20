package utils

import (
	"bufio"
	"os"
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
