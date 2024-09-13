package results

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var (
	titleRegex  = regexp.MustCompile("(?mi)<title>(.*?)</title>")
	cssRegex    = regexp.MustCompile(`\.css(\?|)`)
	jsRegex     = regexp.MustCompile(`\.js(\?|)`)
	headerRegex = regexp.MustCompile("(.*):(.*)")
	tagsRegex   = regexp.MustCompile("<(.*?)>")
	jsonRegex   = regexp.MustCompile("(\"|')(\\s|):(\\s|)(\"|'|)")
)

func CountTags(contentType, body string) string {
	if strings.Contains(contentType, "html") || strings.Contains(contentType, "xml") {
		return strconv.Itoa(len(tagsRegex.FindAllString(body, -1)))
	}
	if strings.Contains(contentType, "json") {
		return strconv.Itoa(len(jsonRegex.FindAllString(body, -1)))
	}
	return "0"
}

func CountCssFiles(body string) string {
	return strconv.Itoa(len(cssRegex.FindAllString(body, -1)))
}

func CountJsFiles(body string) string {
	return strconv.Itoa(len(jsRegex.FindAllString(body, -1)))
}

func CalculateTitleLength(body string) string {
	matches := titleRegex.FindStringSubmatch(body)
	if len(matches) == 2 {
		return strconv.Itoa(len(matches[1]))
	}
	return "0"
}

func CalculateTitleWords(body string) string {
	matches := titleRegex.FindStringSubmatch(body)
	if len(matches) == 2 {
		return strconv.Itoa(len(strings.Fields(matches[1])))
	}
	return "0"
}

func ExtractRedirectDomain(urlStr string) string {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return ""
	}
	return parsedURL.Host
}

func CountRedirectParameters(urlStr string) string {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "0"
	}
	return strconv.Itoa(len(parsedURL.Query()))
}

func CountHeaders(headerString string) string {
	return strconv.Itoa(len(headerRegex.FindAllString(headerString, -1)))
}
