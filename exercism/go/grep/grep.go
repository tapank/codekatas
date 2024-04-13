package grep

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Search(pattern string, flags, files []string) []string {
	// process flags
	var lineNum, namesOnly, ignoreCase, invertMatch, matchLine, appendName bool
	for _, flag := range flags {
		switch flag {
		case "-n":
			lineNum = true
		case "-l":
			namesOnly = true
		case "-i":
			ignoreCase = true
		case "-v":
			invertMatch = true
		case "-x":
			matchLine = true
		default:
			log.Fatalf("Unknown flag: %s", flag)
		}
	}
	if len(files) > 1 {
		appendName = true
	}

	// compile regex
	if matchLine {
		pattern = "^" + pattern + "$"
	}
	if ignoreCase {
		pattern = "(?i)" + pattern
	}
	var r = regexp.MustCompile(pattern)

	// search
	matched := []string{}
	for _, f := range files {
		// prepare file scanner
		file, err := os.Open(f)
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		lineCount := 0

		// scan file
		for scanner.Scan() {
			line := scanner.Text()
			lineCount++
			if r.MatchString(line) != invertMatch {
				if namesOnly {
					matched = append(matched, f)
					break
				} else {
					if lineNum {
						line = strconv.Itoa(lineCount) + ":" + line
					}
					if appendName {
						line = f + ":" + line
					}
					matched = append(matched, line)
				}
			}
		}
	}
	return matched
}
