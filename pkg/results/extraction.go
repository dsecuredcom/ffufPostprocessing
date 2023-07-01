package results

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var (
	reHTMLTags   = regexp.MustCompile(`<(.*?)>`)
	reXMLTags    = regexp.MustCompile(`<(.*?)>`)
	reJSONToken  = regexp.MustCompile(`("|')(\\s|):(\\s|)("|'|)`)
	reCSSExt     = regexp.MustCompile(`\.css(\?|)`)
	reJSExt      = regexp.MustCompile(`\.js(\?|)`)
	reHTMLTitle  = regexp.MustCompile(`(?mi)<title>(.*?)</title>`)
	reHTTPHeader = regexp.MustCompile(`"(.*):(.*)`)
)

func CountTags(ContentType string, Body string) string {
	// I could have used switch/case but this reduces the amount of code
	if strings.Contains(ContentType, "html") {
		data := reHTMLTags.FindAllString(Body, -1)
		return strconv.Itoa(len(data))
	}

	// Should work the same as html, but since I am not sure - keep it seperate
	if strings.Contains(ContentType, "xml") {
		data := reXMLTags.FindAllString(Body, -1)
		return strconv.Itoa(len(data))
	}

	// JSON has obviously no tags, but we can count the number '":" and similar json related syntax'
	// Obviously this is not perfect, but it should work for most cases
	// Parsing json is not an option, since it is not a trivial task and costs a lot of performance
	if strings.Contains(ContentType, "json") {
		data := reJSONToken.FindAllString(Body, -1)
		return strconv.Itoa(len(data))
	}

	//@TODO: Implement other content types?

	return "0"
}

func CountCssFiles(Body string) string {
	data := reCSSExt.FindAllStringIndex(Body, -1)
	return strconv.Itoa(len(data))
}

func CountJsFiles(Body string) string {
	data := reJSExt.FindAllStringIndex(Body, -1)
	return strconv.Itoa(len(data))
}

func CalculateTitleLength(Body string) string {
	// Checks only the first occurenc of a title tag
	data := reHTMLTitle.FindStringSubmatch(Body)

	// Usually we have an array of two elements, first is the complete match including
	// tags, the second one contains only the string between the tags
	if len(data) != 2 {
		return "0"
	}

	return strconv.Itoa(len(data[1]))
}

func CalculateTitleWords(Body string) string {
	// Checks only the first occurenc of a title tag
	data := reHTMLTitle.FindStringSubmatch(Body)

	// Usually we have an array of two elements, first is the complete match including
	// tags, the second one contains only the string between the tags
	if len(data) != 2 {
		return "0"
	}

	Title := data[1]
	SplittedBySpace := strings.Split(Title, " ")

	return strconv.Itoa(len(SplittedBySpace))
}

func ExtractRedirectDomain(Url string) string {
	ParsedUrl, err := url.Parse(Url)
	if err != nil {
		return ""
	}

	return ParsedUrl.Host
}

func CountRedirectParameters(Url string) string {
	ParsedUrl, err := url.Parse(Url)
	if err != nil {
		return "0"
	}

	m, _ := url.ParseQuery(ParsedUrl.RawQuery)
	return strconv.Itoa(len(m))
}

func CountHeaders(HeaderString string) string {
	data := reHTTPHeader.FindAllString(HeaderString, -1)
	return strconv.Itoa(len(data))
}
