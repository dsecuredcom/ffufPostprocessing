package results

import (
	"fmt"
	_struct "github.com/Damian89/ffufPostprocessing/pkg/struct"
	"os"
	"regexp"
)

func EnrichResults(FfufBodiesFolder string, Entries *[]_struct.Result) {

	for i := 0; i < len(*Entries); i++ {

		BodyFilePath := fmt.Sprintf("%s/%s", FfufBodiesFolder, (*Entries)[i].Resultfile)
		ContentFile, err := os.ReadFile(BodyFilePath)

		if err != nil {
			fmt.Printf("\u001B[31m[x]\u001B[0m Could not load body file: %s\n", BodyFilePath)
			continue
		}

		Content := string(ContentFile)
		Headers, Body := SeperateContentIntoHeadersAndBody(Content)

		(*Entries)[i].CountHeaders = len(Headers)
		(*Entries)[i].RedirectDomain = len(Body)
		(*Entries)[i].CountRedirectParameters = "123"
		(*Entries)[i].LengthTitle = "123"
		(*Entries)[i].WordsTitle = "123"
		(*Entries)[i].CountCssFiles = "123"
		(*Entries)[i].CountJsFiles = "123"

	}
}

func SeperateContentIntoHeadersAndBody(Content string) (string, string) {
	// @ TODO: This is not working properly, it is not seperating headers and body correctly
	re := regexp.MustCompile("(?m)\\^.\\*$")
	match := re.FindAllString(Content, -1)
	panic(match[0])
	return "", ""
}
