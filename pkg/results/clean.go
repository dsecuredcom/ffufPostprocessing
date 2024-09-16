package results

import (
	"fmt"
	_struct "github.com/dsecuredcom/ffufPostprocessing/pkg/struct"
	"strconv"
)

func MinimizeOriginalResults(Entries *[]_struct.Result) []_struct.Result {
	// Pre-allocate maps with an estimated capacity
	estimatedCapacity := len(*Entries) / 2
	UniqueStatusMd5 := make(map[string]int, estimatedCapacity)
	UniqueStatusLengthMd5 := make(map[string]int, estimatedCapacity)
	UniqueStatusWordsMd5 := make(map[string]int, estimatedCapacity)
	UniqueStatusLinesMd5 := make(map[string]int, estimatedCapacity)
	UniqueStatusContentTypeMd5 := make(map[string]int, estimatedCapacity)
	UniqueWordsContentTypeMd5 := make(map[string]int, estimatedCapacity)
	UniqueStatusRedirectAndParameters := make(map[string]int, estimatedCapacity)
	UniqueTitleLengthMd5 := make(map[string]int, estimatedCapacity)
	UniqueTitleWordsMd5 := make(map[string]int, estimatedCapacity)
	UniqueTitleLinesWordsMd5 := make(map[string]int, estimatedCapacity)
	UniqueCssFilesMd5 := make(map[string]int, estimatedCapacity)
	UniqueJsFilesMd5 := make(map[string]int, estimatedCapacity)
	UniqueStatusJsCssFilesMd5 := make(map[string]int, estimatedCapacity)
	UniqueTagsMd5 := make(map[string]int, estimatedCapacity)
	UniqueHttpStatusHeaderCountMd5 := make(map[string]int, estimatedCapacity)
	UniqueLengthSumPerHttpStatus := make(map[string]float64, estimatedCapacity)
	UniqueMeanLengthPerHttpStatus := make(map[string]float64, estimatedCapacity)

	// Use a single loop for all analyses
	for i := range *Entries {
		entry := &(*Entries)[i]
		AnalyzeByHttpStatus(entry, &UniqueStatusMd5)
		AnalyzeByHttpStatusAndLength(entry, &UniqueStatusLengthMd5)
		AnalyzeByHttpStatusAndWords(entry, &UniqueStatusWordsMd5)
		AnalyzeByHttpStatusAndLines(entry, &UniqueStatusLinesMd5)
		AnalyzeByHttpStatusAndContentType(entry, &UniqueStatusContentTypeMd5)
		AnalyzeByWordsAndContentType(entry, &UniqueWordsContentTypeMd5)
		AnalyzeByHttpStatusAndRedirectData(entry, &UniqueStatusRedirectAndParameters)
		AnalyzeByTitleLength(entry, &UniqueTitleLengthMd5)
		AnalyzeByTitleWords(entry, &UniqueTitleWordsMd5)
		AnalyzeByTitleLengthWords(entry, &UniqueTitleLinesWordsMd5)
		AnalyzeByCssFiles(entry, &UniqueCssFilesMd5)
		AnalyzeByJsFiles(entry, &UniqueJsFilesMd5)
		AnalyzeByHttpStatusJsCssFiles(entry, &UniqueStatusJsCssFilesMd5)
		AnalyzeByTags(entry, &UniqueTagsMd5)
		AnalyzeByHttpStatusAndHeadersCount(entry, &UniqueHttpStatusHeaderCountMd5)
		CalculateLengthSumPerHttpStatus(entry, &UniqueLengthSumPerHttpStatus)
	}

	CalculateMeanLengthPerHttpStatus(&UniqueStatusMd5, &UniqueLengthSumPerHttpStatus, &UniqueMeanLengthPerHttpStatus)

	// Pre-allocate TemporaryCleanedResults with an estimated capacity
	TemporaryCleanedResults := make([]_struct.Result, 0, len(*Entries)/2)
	PositionsDone := make(map[int]bool, len(*Entries))
	ContentCounterMap := make(map[string]int, len(*Entries))

	// Combine all filtering loops into a single loop
	for i, entry := range *Entries {
		if PositionsDone[i] {
			continue
		}

		statusStr := strconv.Itoa(entry.Status)
		lengthStr := strconv.Itoa(entry.Length)
		wordsStr := strconv.Itoa(entry.Words)
		linesStr := strconv.Itoa(entry.Lines)

		MeanOfCurrentHttpStatus := UniqueMeanLengthPerHttpStatus["status-mean-length:"+statusStr]
		DevFloat := float64(entry.Length) / MeanOfCurrentHttpStatus
		Dev := fmt.Sprintf("%f", DevFloat)
		Content := "dev:" + statusStr + Dev

		if shouldKeepEntry(DevFloat, ContentCounterMap, Content, 2) {
			entry.KeepReason = "deviation (" + Dev + ")"
			TemporaryCleanedResults = append(TemporaryCleanedResults, entry)
			PositionsDone[i] = true
			ContentCounterMap[Content]++
			continue
		}

		Content = "status-length:" + statusStr + ":" + lengthStr
		if UniqueStatusLengthMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			entry.KeepReason = "http status + length"
			TemporaryCleanedResults = append(TemporaryCleanedResults, entry)
			PositionsDone[i] = true
			ContentCounterMap[Content]++
			continue
		}

		Content = "status-words:" + statusStr + ":" + wordsStr
		if UniqueStatusWordsMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			entry.KeepReason = "http status + words"
			TemporaryCleanedResults = append(TemporaryCleanedResults, entry)
			PositionsDone[i] = true
			ContentCounterMap[Content]++
			continue
		}

		Content = "status-lines:" + statusStr + ":" + linesStr
		if UniqueStatusLinesMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			entry.KeepReason = "http status + lines"
			TemporaryCleanedResults = append(TemporaryCleanedResults, entry)
			PositionsDone[i] = true
			ContentCounterMap[Content]++
			continue
		}

		Content = "words-content-type:" + wordsStr + ":" + entry.ContentType
		if UniqueWordsContentTypeMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			entry.KeepReason = "words + content type"
			TemporaryCleanedResults = append(TemporaryCleanedResults, entry)
			PositionsDone[i] = true
			ContentCounterMap[Content]++
			continue
		}

		Content = "status-content type:" + statusStr + ":" + entry.ContentType
		if UniqueStatusContentTypeMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			entry.KeepReason = "http status + content type"
			TemporaryCleanedResults = append(TemporaryCleanedResults, entry)
			PositionsDone[i] = true
			ContentCounterMap[Content]++
			continue
		}

		Content = "status+js+css:" + statusStr + ":" + entry.CountJsFiles + ":" + entry.CountCssFiles
		if UniqueStatusJsCssFilesMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			entry.KeepReason = "status+js+css"
			TemporaryCleanedResults = append(TemporaryCleanedResults, entry)
			PositionsDone[i] = true
			ContentCounterMap[Content]++
			continue
		}

		Content = "status-redirect:" + statusStr + ":" + entry.RedirectDomain + ":" + entry.CountRedirectParameters
		if UniqueStatusRedirectAndParameters[Content] < 5 && ContentCounterMap[Content] < 2 {
			entry.KeepReason = "http status + redirect"
			TemporaryCleanedResults = append(TemporaryCleanedResults, entry)
			PositionsDone[i] = true
			ContentCounterMap[Content]++
			continue
		}

		Content = "status-header-count:" + statusStr + ":" + entry.CountHeaders
		if UniqueHttpStatusHeaderCountMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			entry.KeepReason = "http status + header count"
			TemporaryCleanedResults = append(TemporaryCleanedResults, entry)
			PositionsDone[i] = true
			ContentCounterMap[Content]++
			continue
		}

		Content = "title-length:" + entry.LengthTitle
		if UniqueTitleLengthMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			entry.KeepReason = "title length"
			TemporaryCleanedResults = append(TemporaryCleanedResults, entry)
			PositionsDone[i] = true
			ContentCounterMap[Content]++
			continue
		}

		Content = "title-words:" + entry.WordsTitle
		if UniqueTitleWordsMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			entry.KeepReason = "title words"
			TemporaryCleanedResults = append(TemporaryCleanedResults, entry)
			PositionsDone[i] = true
			ContentCounterMap[Content]++
			continue
		}

		Content = "title-length-words:" + entry.WordsTitle + ":" + entry.LengthTitle
		if UniqueTitleLinesWordsMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			entry.KeepReason = "title length + words"
			TemporaryCleanedResults = append(TemporaryCleanedResults, entry)
			PositionsDone[i] = true
			ContentCounterMap[Content]++
			continue
		}

		Content = "css:" + entry.CountCssFiles
		if UniqueCssFilesMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			entry.KeepReason = "css files"
			TemporaryCleanedResults = append(TemporaryCleanedResults, entry)
			PositionsDone[i] = true
			ContentCounterMap[Content]++
			continue
		}

		Content = "js:" + entry.CountJsFiles
		if UniqueJsFilesMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			entry.KeepReason = "js files"
			TemporaryCleanedResults = append(TemporaryCleanedResults, entry)
			PositionsDone[i] = true
			ContentCounterMap[Content]++
			continue
		}

		Content = "tags:" + entry.CountTags
		if UniqueTagsMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			entry.KeepReason = "tags"
			TemporaryCleanedResults = append(TemporaryCleanedResults, entry)
			PositionsDone[i] = true
			ContentCounterMap[Content]++
			continue
		}

		Content = "status:" + statusStr
		if UniqueStatusMd5[Content] > 0 && ContentCounterMap[Content] < 2 {
			entry.KeepReason = "http status"
			TemporaryCleanedResults = append(TemporaryCleanedResults, entry)
			PositionsDone[i] = true
			ContentCounterMap[Content]++
			continue
		}
	}

	return TemporaryCleanedResults
}

func CalculateMeanLengthPerHttpStatus(HttpStatusData *map[string]int, LengthSum *map[string]float64, LengthMean *map[string]float64) {
	for id, value := range *LengthSum {
		(*LengthMean)["status-mean-length:"+id] = value / float64((*HttpStatusData)["status:"+id])
	}
}

func CalculateLengthSumPerHttpStatus(entry *_struct.Result, Hashes *map[string]float64) {
	Content := strconv.Itoa(entry.Status)
	if (*Hashes)[Content] == 0 {
		(*Hashes)[Content] = float64(entry.Length)
	} else {
		(*Hashes)[Content] = float64(entry.Length) + (*Hashes)[Content]
	}
}

func AnalyzeByHttpStatusJsCssFiles(entry *_struct.Result, Hashes *map[string]int) {
	Content := "status+js+css:" + strconv.Itoa(entry.Status) + ":" + entry.CountJsFiles + ":" + entry.CountCssFiles

	if (*Hashes)[Content] == 0 {
		(*Hashes)[Content] = 1
	} else {
		(*Hashes)[Content]++
	}
}

func AnalyzeByWordsAndContentType(entry *_struct.Result, Hashes *map[string]int) {

	Content := "words-content-type:" + strconv.Itoa(entry.Words) + ":" + entry.ContentType

	if (*Hashes)[Content] == 0 {
		(*Hashes)[Content] = 1
	} else {
		(*Hashes)[Content]++
	}
}

func AnalyzeByHttpStatus(entry *_struct.Result, StatusMd5 *map[string]int) {

	Content := "status:" + strconv.Itoa(entry.Status)

	if (*StatusMd5)[Content] == 0 {
		(*StatusMd5)[Content] = 1
	} else {
		(*StatusMd5)[Content]++
	}

}

func AnalyzeByHttpStatusAndLength(entry *_struct.Result, StatusLengthMd5 *map[string]int) {
	Content := "status-length:" + strconv.Itoa(entry.Status) + ":" + strconv.Itoa(entry.Length)

	if (*StatusLengthMd5)[Content] == 0 {
		(*StatusLengthMd5)[Content] = 1
	} else {
		(*StatusLengthMd5)[Content]++
	}

}

func AnalyzeByHttpStatusAndHeadersCount(entry *_struct.Result, countMd5 *map[string]int) {
	Content := "status-header-count:" + strconv.Itoa(entry.Status) + ":" + entry.CountHeaders

	if (*countMd5)[Content] == 0 {
		(*countMd5)[Content] = 1
	} else {
		(*countMd5)[Content]++
	}

}

func AnalyzeByTags(entry *_struct.Result, tagsMd5 *map[string]int) {
	Content := "tags:" + entry.CountTags

	if (*tagsMd5)[Content] == 0 {
		(*tagsMd5)[Content] = 1
	} else {
		(*tagsMd5)[Content]++
	}

}

func AnalyzeByJsFiles(entry *_struct.Result, filesMd5 *map[string]int) {
	Content := "js:" + entry.CountJsFiles

	if (*filesMd5)[Content] == 0 {
		(*filesMd5)[Content] = 1
	} else {
		(*filesMd5)[Content]++
	}
}

func AnalyzeByCssFiles(entry *_struct.Result, filesMd5 *map[string]int) {
	Content := "css:" + entry.CountCssFiles

	if (*filesMd5)[Content] == 0 {
		(*filesMd5)[Content] = 1
	} else {
		(*filesMd5)[Content]++
	}
}

func AnalyzeByTitleLengthWords(entry *_struct.Result, wordsMd5 *map[string]int) {
	Content := "title-length-words:" + entry.WordsTitle + ":" + entry.LengthTitle

	if (*wordsMd5)[Content] == 0 {
		(*wordsMd5)[Content] = 1
	} else {
		(*wordsMd5)[Content]++
	}
}

func AnalyzeByTitleWords(entry *_struct.Result, wordsMd5 *map[string]int) {
	Content := "title-words:" + entry.WordsTitle

	if (*wordsMd5)[Content] == 0 {
		(*wordsMd5)[Content] = 1
	} else {
		(*wordsMd5)[Content]++
	}
}

func AnalyzeByTitleLength(entry *_struct.Result, lengthMd5 *map[string]int) {
	Content := "title-length:" + entry.LengthTitle

	if (*lengthMd5)[Content] == 0 {
		(*lengthMd5)[Content] = 1
	} else {
		(*lengthMd5)[Content]++
	}
}

func AnalyzeByHttpStatusAndRedirectData(entry *_struct.Result, parameters *map[string]int) {
	Content := "status-redirect:" + strconv.Itoa(entry.Status) + ":" + entry.RedirectDomain + ":" + entry.CountRedirectParameters

	if (*parameters)[Content] == 0 {
		(*parameters)[Content] = 1
	} else {
		(*parameters)[Content]++
	}
}

func AnalyzeByHttpStatusAndContentType(entry *_struct.Result, StatusCTMd5 *map[string]int) {
	Content := "status-content type:" + strconv.Itoa(entry.Status) + ":" + entry.ContentType

	if (*StatusCTMd5)[Content] == 0 {
		(*StatusCTMd5)[Content] = 1
	} else {
		(*StatusCTMd5)[Content]++
	}
}

func AnalyzeByHttpStatusAndLines(entry *_struct.Result, StatusLinesMd5 *map[string]int) {
	Content := "status-lines:" + strconv.Itoa(entry.Status) + ":" + strconv.Itoa(entry.Lines)

	if (*StatusLinesMd5)[Content] == 0 {
		(*StatusLinesMd5)[Content] = 1
	} else {
		(*StatusLinesMd5)[Content]++
	}

}

func AnalyzeByHttpStatusAndWords(entry *_struct.Result, StatusWordsMd5 *map[string]int) {
	Content := "status-words:" + strconv.Itoa(entry.Status) + ":" + strconv.Itoa(entry.Words)

	if (*StatusWordsMd5)[Content] == 0 {
		(*StatusWordsMd5)[Content] = 1
	} else {
		(*StatusWordsMd5)[Content]++
	}

}

func shouldKeepEntry(devFloat float64, contentCounterMap map[string]int, content string, threshold int) bool {
	return devFloat != 1.0 && contentCounterMap[content] <= threshold
}
