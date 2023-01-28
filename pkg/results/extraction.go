package results

import (
	"net/url"
	"regexp"
	"strings"
)

func CountCssFiles(Body string) int {
	r, _ := regexp.Compile(`\.css(\?|)`)
	data := r.FindAllStringIndex(Body, -1)
	return len(data)
}

func CountJsFiles(Body string) int {
	r, _ := regexp.Compile(`\.js(\?|)`)
	data := r.FindAllStringIndex(Body, -1)
	return len(data)
}

func CalculateTitleLength(Body string) int {

	r, _ := regexp.Compile("(?mi)<title>(.*?)</title>")
	// Checks only the first occurenc of a title tag
	data := r.FindStringSubmatch(Body)

	// Usually we have an array of two elements, first is the complete match including
	// tags, the second one contains only the string between the tags
	if len(data) != 2 {
		return 0
	}

	return len(data[1])
}

func CalculateTitleWords(Body string) int {

	r, _ := regexp.Compile("(?mi)<title>(.*?)</title>")
	// Checks only the first occurenc of a title tag
	data := r.FindStringSubmatch(Body)

	// Usually we have an array of two elements, first is the complete match including
	// tags, the second one contains only the string between the tags
	if len(data) != 2 {
		return 0
	}

	Title := data[1]
	SplittedBySpace := strings.Split(Title, " ")

	return len(SplittedBySpace)
}

func ExtractRedirectDomain(Url string) string {
	ParsedUrl, err := url.Parse(Url)
	if err != nil {
		return ""
	}

	return ParsedUrl.Host
}

func CountRedirectParameters(Url string) int {
	ParsedUrl, err := url.Parse(Url)
	if err != nil {
		return 0
	}

	m, _ := url.ParseQuery(ParsedUrl.RawQuery)
	return len(m)
}

func CountHeaders(HeaderString string) int {

	r, _ := regexp.Compile("(.*):(.*)")
	data := r.FindAllString(HeaderString, -1)

	return len(data)
}
