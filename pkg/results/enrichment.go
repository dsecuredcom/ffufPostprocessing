package results

import (
	"github.com/Damian89/ffufPostprocessing/pkg/general"
	_struct "github.com/Damian89/ffufPostprocessing/pkg/struct"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func EnrichResultsWithRedirectData(Entries *[]_struct.Result) {

	for i := 0; i < len(*Entries); i++ {
		(*Entries)[i].RedirectDomain = ExtractRedirectDomain((*Entries)[i].RedirectLocation)
		(*Entries)[i].CountRedirectParameters = CountRedirectParameters((*Entries)[i].RedirectLocation)
	}
}
func EnrichResults(FfufBodiesFolder string, Entries *[]_struct.Result) {

	sem := make(chan struct{}, 25)
	var wg sync.WaitGroup
	for i := 0; i < len(*Entries); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			sem <- struct{}{}
			FfufBodiesFolder = strings.TrimRight(FfufBodiesFolder, "/\\")
			BodyFilePath := filepath.Join(FfufBodiesFolder, (*Entries)[i].Resultfile)
			ContentFile, _ := os.ReadFile(BodyFilePath)

			if general.FileExists(BodyFilePath) == false {
				<-sem
				return
			}

			Content := string(ContentFile)
			Headers, Body := SeperateContentIntoHeadersAndBody(Content)

			(*Entries)[i].CountHeaders = CountHeaders(Headers)
			(*Entries)[i].LengthTitle = CalculateTitleLength(Body)
			(*Entries)[i].WordsTitle = CalculateTitleWords(Body)
			(*Entries)[i].CountCssFiles = CountCssFiles(Body)
			(*Entries)[i].CountJsFiles = CountJsFiles(Body)
			(*Entries)[i].CountTags = CountTags((*Entries)[i].ContentType, Body)
			<-sem
		}(i)

	}
	wg.Wait()
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
