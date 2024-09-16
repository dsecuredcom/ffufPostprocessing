package results

import (
	"bufio"
	"github.com/dsecuredcom/ffufPostprocessing/pkg/general"
	_struct "github.com/dsecuredcom/ffufPostprocessing/pkg/struct"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

const maxReadSize = 786432

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
	numWorkers := runtime.NumCPU() // You can adjust this number based on your needs
	jobs := make(chan int, len(*Entries))
	var wg sync.WaitGroup

	// Create worker pool
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg, FfufBodiesFolder, Entries)
	}

	// Send jobs to the pool
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
