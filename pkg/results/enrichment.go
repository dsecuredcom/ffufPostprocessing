package results

import (
	"fmt"
	_struct "github.com/Damian89/ffufPostprocessing/pkg/struct"
	"os"
	"strings"
)

func EnrichResultsWithRedirectData(Entries *[]_struct.Result) {

	for i := 0; i < len(*Entries); i++ {
		(*Entries)[i].RedirectDomain = ExtractRedirectDomain((*Entries)[i].RedirectLocation)
		(*Entries)[i].CountRedirectParameters = CountRedirectParameters((*Entries)[i].RedirectLocation)
	}
}
func EnrichResults(FfufBodiesFolder string, Entries *[]_struct.Result) {

	for i := 0; i < len(*Entries); i++ {

		FfufBodiesFolder = strings.TrimRight(FfufBodiesFolder, "/") + "/"
		BodyFilePath := fmt.Sprintf("%s/%s", FfufBodiesFolder, (*Entries)[i].Resultfile)
		ContentFile, err := os.ReadFile(BodyFilePath)

		if err != nil {
			fmt.Printf("\u001B[31m[x]\u001B[0m Could not load body file: %s\n", BodyFilePath)
			continue
		}

		Content := string(ContentFile)
		Headers, Body := SeperateContentIntoHeadersAndBody(Content)

		(*Entries)[i].CountHeaders = CountHeaders(Headers)
		(*Entries)[i].LengthTitle = CalculateTitleLength(Body)
		(*Entries)[i].WordsTitle = CalculateTitleWords(Body)
		(*Entries)[i].CountCssFiles = CountCssFiles(Body)
		(*Entries)[i].CountJsFiles = CountJsFiles(Body)
		(*Entries)[i].CountTags = CountTags((*Entries)[i].ContentType, Body)
	}
}

func SeperateContentIntoHeadersAndBody(Content string) (string, string) {

	/**
	* This could be done WAY better with a solid regular expression I guess, but I had no time to test - so here is the
	* stupid guys solution. If you read this, please feel free to improve this function!
	 */
	EntireResponseArray := strings.Split(Content, "---- ↑ Request ---- Response ↓ ----")

	EntireResponse := strings.Trim(EntireResponseArray[1], "\r\n")

	HeaderString := ""
	BodyString := ""

	EntireResponseByLine := strings.Split(EntireResponse, "\n")
	var line string
	stringToAddLineTo := "header"

	for i := 0; i < len(EntireResponseByLine); i++ {
		// First line is something like: HTTP/1.1 200 OK
		if i == 0 {
			continue
		}

		// Removes whitespaces at the end of current line, drops basically \r and \n at the end
		line = strings.TrimRight(strings.TrimRight(EntireResponseByLine[i], "\n"), "\r")

		if stringToAddLineTo == "header" {
			HeaderString += line + "\n"
		} else {
			BodyString += line + "\n"
		}

		if line == "" {
			stringToAddLineTo = "body"
		}
	}

	HeaderString = (strings.Trim(HeaderString, "\r\n"))
	BodyString = (strings.Trim(BodyString, "\r\n"))
	return HeaderString, BodyString
}
