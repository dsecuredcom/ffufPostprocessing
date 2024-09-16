package results

import (
	"fmt"
	_struct "github.com/dsecuredcom/ffufPostprocessing/pkg/struct"
	"strconv"
)

func MinimizeOriginalResults(Entries *[]_struct.Result) []_struct.Result {
	uniqueChecks := make(map[string]int)
	uniqueLengthSum := make(map[string]float64)
	uniqueMeanLength := make(map[string]float64)

	// First pass: Analyze entries and calculate means
	for _, entry := range *Entries {
		analyzeEntry(&entry, uniqueChecks, &uniqueLengthSum)
	}
	calculateMeanLength(uniqueChecks, uniqueLengthSum, uniqueMeanLength)

	// Second pass: Filter entries
	var cleanedResults []_struct.Result
	contentCounter := make(map[string]int)

	for _, entry := range *Entries {
		if keepEntry(&entry, uniqueChecks, uniqueMeanLength, contentCounter) {
			cleanedResults = append(cleanedResults, entry)
		}
	}

	return cleanedResults
}

func analyzeEntry(entry *_struct.Result, uniqueChecks map[string]int, uniqueLengthSum *map[string]float64) {
	status := strconv.Itoa(entry.Status)
	incrementUniqueCheck(uniqueChecks, "status:"+status)
	incrementUniqueCheck(uniqueChecks, "status-length:"+status+":"+strconv.Itoa(entry.Length))
	incrementUniqueCheck(uniqueChecks, "status-words:"+status+":"+strconv.Itoa(entry.Words))
	incrementUniqueCheck(uniqueChecks, "status-lines:"+status+":"+strconv.Itoa(entry.Lines))
	incrementUniqueCheck(uniqueChecks, "status-content-type:"+status+":"+entry.ContentType)
	incrementUniqueCheck(uniqueChecks, "words-content-type:"+strconv.Itoa(entry.Words)+":"+entry.ContentType)
	incrementUniqueCheck(uniqueChecks, "status-redirect:"+status+":"+entry.RedirectDomain+":"+entry.CountRedirectParameters)
	incrementUniqueCheck(uniqueChecks, "title-length:"+entry.LengthTitle)
	incrementUniqueCheck(uniqueChecks, "title-words:"+entry.WordsTitle)
	incrementUniqueCheck(uniqueChecks, "title-length-words:"+entry.WordsTitle+":"+entry.LengthTitle)
	incrementUniqueCheck(uniqueChecks, "css:"+entry.CountCssFiles)
	incrementUniqueCheck(uniqueChecks, "js:"+entry.CountJsFiles)
	incrementUniqueCheck(uniqueChecks, "status-js-css:"+status+":"+entry.CountJsFiles+":"+entry.CountCssFiles)
	incrementUniqueCheck(uniqueChecks, "tags:"+entry.CountTags)
	incrementUniqueCheck(uniqueChecks, "status-header-count:"+status+":"+entry.CountHeaders)

	(*uniqueLengthSum)[status] += float64(entry.Length)
}

func incrementUniqueCheck(uniqueChecks map[string]int, key string) {
	uniqueChecks[key]++
}

func calculateMeanLength(uniqueChecks map[string]int, uniqueLengthSum, uniqueMeanLength map[string]float64) {
	for status, sum := range uniqueLengthSum {
		uniqueMeanLength["status-mean-length:"+status] = sum / float64(uniqueChecks["status:"+status])
	}
}

func keepEntry(entry *_struct.Result, uniqueChecks map[string]int, uniqueMeanLength map[string]float64, contentCounter map[string]int) bool {
	status := strconv.Itoa(entry.Status)
	meanLength := uniqueMeanLength["status-mean-length:"+status]
	deviation := float64(entry.Length) / meanLength
	devContent := "dev:" + status + fmt.Sprintf("%f", deviation)

	if deviation != 1.0 && contentCounter[devContent] <= 2 {
		entry.KeepReason = "deviation (" + fmt.Sprintf("%f", deviation) + ")"
		contentCounter[devContent]++
		return true
	}

	checks := []struct {
		key    string
		reason string
		limit  int
	}{
		{"status-length:" + status + ":" + strconv.Itoa(entry.Length), "http status + length", 5},
		{"status-words:" + status + ":" + strconv.Itoa(entry.Words), "http status + words", 5},
		{"status-lines:" + status + ":" + strconv.Itoa(entry.Lines), "http status + lines", 5},
		{"words-content-type:" + strconv.Itoa(entry.Words) + ":" + entry.ContentType, "words + content type", 5},
		{"status-content-type:" + status + ":" + entry.ContentType, "http status + content type", 5},
		{"status-js-css:" + status + ":" + entry.CountJsFiles + ":" + entry.CountCssFiles, "status+js+css", 5},
		{"status-redirect:" + status + ":" + entry.RedirectDomain + ":" + entry.CountRedirectParameters, "http status + redirect", 5},
		{"status-header-count:" + status + ":" + entry.CountHeaders, "http status + header count", 5},
		{"title-length:" + entry.LengthTitle, "title length", 5},
		{"title-words:" + entry.WordsTitle, "title words", 5},
		{"title-length-words:" + entry.WordsTitle + ":" + entry.LengthTitle, "title length + words", 5},
		{"css:" + entry.CountCssFiles, "css files", 5},
		{"js:" + entry.CountJsFiles, "js files", 5},
		{"tags:" + entry.CountTags, "tags", 5},
		{"status:" + status, "http status", 1},
	}

	for _, check := range checks {
		if uniqueChecks[check.key] < check.limit && contentCounter[check.key] < 2 {
			entry.KeepReason = check.reason
			contentCounter[check.key]++
			return true
		}
	}

	return false
}
