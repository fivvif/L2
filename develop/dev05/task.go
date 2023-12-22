package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(grep())
}

func grep() interface{} {
	after := flag.Int("A", 0, "print +N lines after a match")
	before := flag.Int("B", 0, "print +N lines before a match")
	context := flag.Int("C", 0, "print +-N lines around a match")
	count := flag.Bool("c", false, "number of lines")
	ignoreCase := flag.Bool("i", false, "ignore register")
	invert := flag.Bool("v", false, "Invert matching(exclude matching lines")
	fixed := flag.Bool("F", false, "Treat the pattern as a fixed string (not a regular expression)")
	lineShow := flag.Bool("n", false, "Print line numbers")
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)

	if len(args) < 2 {
		flag.PrintDefaults()
		log.Fatalln("Usage:  [OPTIONS] PATTERN FILE")

	}
	pattern := args[0]
	fileName := args[1]

	var regex *regexp.Regexp
	if *ignoreCase {
		regex = regexp.MustCompile("(?i)" + pattern)
	} else {
		regex = regexp.MustCompile(pattern)
	}
	if *fixed {
		regex = regexp.MustCompile("\\b" + pattern + "\\b")
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Error while opening file :", err.Error())
	}
	defer file.Close()

	if *context > 0 {
		*after = *context
		*before = *context
	}

	scanner := bufio.NewScanner(file)
	var buffer []string
	var lines []string
	var matchedLines int
	flagAfter := 0
	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		if *lineShow {
			buffer = append(buffer, fmt.Sprintf("%d:%s", lineNumber, scanner.Text()))
		} else {
			buffer = append(buffer, scanner.Text())
		}

		line := scanner.Text()
		if regex.MatchString(line) != *invert {

			matchedLines++
			if *count {
				continue
			}
			if *before > 0 {
				if lineNumber-*before-1 < 0 {
					lines = append(lines, buffer[:lineNumber-1]...)
				} else {
					lines = append(lines, buffer[lineNumber-*before-1:lineNumber-1]...)
				}
			}

			if *lineShow {
				line = fmt.Sprintf("%d:%s", lineNumber, line)
			}

			lines = append(lines, line)

			if *after > 0 && regex.MatchString(line) != *invert {
				flagAfter = *after + 1
			}
		}
		if regex.MatchString(line) == *invert && flagAfter > 0 {
			if *lineShow {
				line = fmt.Sprintf("%d:%s", lineNumber, line)
			}

			lines = append(lines, line)
		}
		flagAfter--

	}
	if *count {
		return matchedLines
	} else {
		lines = removeDuplicates(lines)
		return strings.Join(lines, "\n") + "\n"
	}

}

func removeDuplicates(slice []string) []string {
	uniqueElements := make(map[string]bool)
	result := []string{}

	for _, element := range slice {
		if _, exists := uniqueElements[element]; !exists {
			uniqueElements[element] = true
			result = append(result, element)
		}
	}

	return result
}
