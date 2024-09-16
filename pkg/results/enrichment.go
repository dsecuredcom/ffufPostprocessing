package results

import (
	"github.com/dsecuredcom/ffufPostprocessing/pkg/general"
	_struct "github.com/dsecuredcom/ffufPostprocessing/pkg/struct"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

const maxReadSize = 786432 // 3/4 of 1MB

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

func EnrichResults(FfufBodiesFolder string, Entries *[]_struct.Result) {
	numWorkers := runtime.NumCPU()
	jobs := make(chan int, len(*Entries))
	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg, FfufBodiesFolder, Entries)
	}

	for i := range *Entries {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup, FfufBodiesFolder string, Entries *[]_struct.Result) {
	defer wg.Done()
	for i := range jobs {
		enrichEntry(i, FfufBodiesFolder, Entries)
	}
}

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

	content := make([]byte, maxReadSize)
	n, err := io.ReadFull(file, content)
	if err != nil && err != io.ErrUnexpectedEOF {
		return
	}
	content = content[:n]

	Headers, Body := SeperateContentIntoHeadersAndBody(string(content))

	(*Entries)[i].CountHeaders = CountHeaders(Headers)
	(*Entries)[i].LengthTitle = CalculateTitleLength(Body)
	(*Entries)[i].WordsTitle = CalculateTitleWords(Body)
	(*Entries)[i].CountCssFiles = CountCssFiles(Body)
	(*Entries)[i].CountJsFiles = CountJsFiles(Body)
	(*Entries)[i].CountTags = CountTags((*Entries)[i].ContentType, Body)
}

func SeperateContentIntoHeadersAndBody(Content string) (string, string) {
	parts := strings.SplitN(Content, "---- ↑ Request ---- Response ↓ ----", 2)
	if len(parts) < 2 {
		return "", ""
	}

	EntireResponse := strings.TrimSpace(parts[1])

	var HeaderBuilder, BodyBuilder strings.Builder
	inHeaders := true

	lines := strings.Split(EntireResponse, "\n")

	// Start from index 1, skipping the first line
	for i := 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

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

	HeaderString := strings.TrimSpace(HeaderBuilder.String())
	BodyString := strings.TrimSpace(BodyBuilder.String())
	return HeaderString, BodyString
}
