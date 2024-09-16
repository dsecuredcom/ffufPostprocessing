package results

import (
	"bufio"
	"github.com/dsecuredcom/ffufPostprocessing/pkg/general"
	_struct "github.com/dsecuredcom/ffufPostprocessing/pkg/struct"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

// EnrichResultsWithRedirectData enriches results with redirect data
func EnrichResultsWithRedirectData(Entries *[]_struct.Result) {
	var wg sync.WaitGroup
	for i := range *Entries {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			(*Entries)[i].RedirectDomain = ExtractRedirectDomain((*Entries)[i].RedirectLocation)
			(*Entries)[i].CountRedirectParameters = CountRedirectParameters((*Entries)[i].RedirectLocation)
		}(i)
	}
	wg.Wait()
}

// EnrichResults enriches results with additional data from body files
func EnrichResults(FfufBodiesFolder string, Entries *[]_struct.Result) {
	numWorkers := runtime.NumCPU()
	semaphore := make(chan struct{}, numWorkers)
	var wg sync.WaitGroup

	for i := range *Entries {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()
			enrichEntry(i, FfufBodiesFolder, Entries)
		}(i)
	}

	wg.Wait()
}

// enrichEntry enriches a single entry with data from its body file
func enrichEntry(i int, FfufBodiesFolder string, Entries *[]_struct.Result) {
	FfufBodiesFolder = strings.TrimRight(FfufBodiesFolder, "/\\")
	BodyFilePath := filepath.Join(FfufBodiesFolder, (*Entries)[i].Resultfile)

	if !general.FileExists(BodyFilePath) {
		return
	}

	file, err := os.Open(BodyFilePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var headers, body strings.Builder
	inHeaders := true

	for scanner.Scan() {
		line := scanner.Text()
		if inHeaders {
			if line == "" {
				inHeaders = false
				continue
			}
			headers.WriteString(line)
			headers.WriteByte('\n')
		} else {
			body.WriteString(line)
			body.WriteByte('\n')
		}
	}

	// Process headers and body
	(*Entries)[i].CountHeaders = CountHeaders(headers.String())
	bodyStr := body.String()
	(*Entries)[i].LengthTitle = CalculateTitleLength(bodyStr)
	(*Entries)[i].WordsTitle = CalculateTitleWords(bodyStr)
	(*Entries)[i].CountCssFiles = CountCssFiles(bodyStr)
	(*Entries)[i].CountJsFiles = CountJsFiles(bodyStr)
	(*Entries)[i].CountTags = CountTags((*Entries)[i].ContentType, bodyStr)
}

// SeperateContentIntoHeadersAndBody separates the content into headers and body
func SeperateContentIntoHeadersAndBody(Content string) (string, string) {
	parts := strings.SplitN(Content, "---- ↑ Request ---- Response ↓ ----", 2)
	if len(parts) < 2 {
		return "", ""
	}

	EntireResponse := strings.TrimSpace(parts[1])

	var HeaderBuilder, BodyBuilder strings.Builder
	inHeaders := true

	scanner := bufio.NewScanner(strings.NewReader(EntireResponse))
	scanner.Split(bufio.ScanLines)

	// Skip the first line
	if scanner.Scan() {
		for scanner.Scan() {
			line := scanner.Text()
			if inHeaders {
				if line == "" {
					inHeaders = false
					continue
				}
				HeaderBuilder.WriteString(line)
				HeaderBuilder.WriteByte('\n')
			} else {
				BodyBuilder.WriteString(line)
				BodyBuilder.WriteByte('\n')
			}
		}
	}

	HeaderString := strings.TrimSpace(HeaderBuilder.String())
	BodyString := strings.TrimSpace(BodyBuilder.String())
	return HeaderString, BodyString
}
