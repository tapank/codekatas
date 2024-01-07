package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	valid := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return valid.Match([]byte(text))
}

func SplitLogLine(text string) []string {
	split := regexp.MustCompile(`<[~\*=-]*>`)
	return split.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	match := regexp.MustCompile(`(?i)".*password.*"`)
	count := 0
	for _, l := range lines {
		if match.Match([]byte(l)) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	r := regexp.MustCompile(`end-of-line\d+`)
	b := r.ReplaceAll([]byte(text), []byte(""))
	return string(b)
}

func TagWithUserName(lines []string) []string {
	r := regexp.MustCompile(`User +(\w+)`)
	logs := []string{}
	for _, l := range lines {
		match := r.FindStringSubmatch(l)
		if len(match) > 1 {
			l = "[USR] " + match[1] + " " + l
		}
		logs = append(logs, l)
	}
	return logs
}
