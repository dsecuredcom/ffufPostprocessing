package results

import (
	"fmt"
	_struct "github.com/Damian89/ffufPostprocessing/pkg/struct"
	"strconv"
)

func MinimizeOriginalResults(Entries *[]_struct.Result) []_struct.Result {

	UniqueStatusMd5 := map[string]int{}
	UniqueStatusLengthMd5 := map[string]int{}
	UniqueStatusWordsMd5 := map[string]int{}
	UniqueStatusLinesMd5 := map[string]int{}
	UniqueStatusContentTypeMd5 := map[string]int{}
	UniqueWordsContentTypeMd5 := map[string]int{}
	UniqueStatusRedirectAndParameters := map[string]int{}
	UniqueTitleLengthMd5 := map[string]int{}
	UniqueTitleWordsMd5 := map[string]int{}
	UniqueTitleLinesWordsMd5 := map[string]int{}
	UniqueCssFilesMd5 := map[string]int{}
	UniqueJsFilesMd5 := map[string]int{}
	UniqueTagsMd5 := map[string]int{}
	UniqueHttpStatusHeaderCountMd5 := map[string]int{}
	var MeanOfLength float64 = 0

	for i := 0; i < len(*Entries); i++ {
		AnalyzeByHttpStatus(Entries, i, &UniqueStatusMd5)
		AnalyzeByHttpStatusAndLength(Entries, i, &UniqueStatusLengthMd5)
		AnalyzeByHttpStatusAndWords(Entries, i, &UniqueStatusWordsMd5)
		AnalyzeByHttpStatusAndLines(Entries, i, &UniqueStatusLinesMd5)
		AnalyzeByHttpStatusAndContentType(Entries, i, &UniqueStatusContentTypeMd5)
		AnalyzeByWordsAndContentType(Entries, i, &UniqueWordsContentTypeMd5)
		AnalyzeByHttpStatusAndRedirectData(Entries, i, &UniqueStatusRedirectAndParameters)
		AnalyzeByTitleLength(Entries, i, &UniqueTitleLengthMd5)
		AnalyzeByTitleWords(Entries, i, &UniqueTitleWordsMd5)
		AnalyzeByTitleLengthWords(Entries, i, &UniqueTitleLinesWordsMd5)
		AnalyzeByCssFiles(Entries, i, &UniqueCssFilesMd5)
		AnalyzeByJsFiles(Entries, i, &UniqueJsFilesMd5)
		AnalyzeByTags(Entries, i, &UniqueTagsMd5)
		AnalyzeByHttpStatusAndHeadersCount(Entries, i, &UniqueHttpStatusHeaderCountMd5)
		MeanOfLength += float64((*Entries)[i].Length)
	}

	MeanOfLength = MeanOfLength / float64(len(*Entries))

	TemporaryCleanedResults := []_struct.Result{}
	PositionsDone := map[int]bool{}
	ContentCounterMap := map[string]int{}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		DevFloat := float64((*Entries)[i].Length) / MeanOfLength

		Dev := fmt.Sprintf("%f", DevFloat)
		Content := "dev:" + Dev

		if DevFloat > 2.0 && ContentCounterMap[Content] < 1 {
			(*Entries)[i].KeepReason = "deviation (" + Dev + ")"
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++

		}

	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "status-length:" + strconv.Itoa((*Entries)[i].Status) + ":" + strconv.Itoa((*Entries)[i].Length)

		if UniqueStatusLengthMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			(*Entries)[i].KeepReason = "http status + length"
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "status-words:" + strconv.Itoa((*Entries)[i].Status) + ":" + strconv.Itoa((*Entries)[i].Words)

		if UniqueStatusWordsMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			(*Entries)[i].KeepReason = "http status + words"
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "status-lines:" + strconv.Itoa((*Entries)[i].Status) + ":" + strconv.Itoa((*Entries)[i].Lines)

		if UniqueStatusLinesMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			(*Entries)[i].KeepReason = "http status + lines"
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "words-content-type:" + strconv.Itoa((*Entries)[i].Words) + ":" + (*Entries)[i].ContentType

		if UniqueWordsContentTypeMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			(*Entries)[i].KeepReason = "words + content type"
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "status-content type:" + strconv.Itoa((*Entries)[i].Status) + ":" + (*Entries)[i].ContentType

		if UniqueStatusContentTypeMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			(*Entries)[i].KeepReason = "http status + content type"
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "status-redirect:" + strconv.Itoa((*Entries)[i].Status) + ":" + (*Entries)[i].RedirectDomain + ":" + (*Entries)[i].CountRedirectParameters

		if UniqueStatusRedirectAndParameters[Content] < 5 && ContentCounterMap[Content] < 2 {
			(*Entries)[i].KeepReason = "http status + redirect"
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "status-header-count:" + strconv.Itoa((*Entries)[i].Status) + ":" + (*Entries)[i].CountHeaders

		if UniqueHttpStatusHeaderCountMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			(*Entries)[i].KeepReason = "http status + header count"
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "title-length:" + (*Entries)[i].LengthTitle

		if UniqueTitleLengthMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			(*Entries)[i].KeepReason = "title length"
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "title-words:" + (*Entries)[i].WordsTitle

		if UniqueTitleWordsMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			(*Entries)[i].KeepReason = "title words"
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "title-length-words:" + (*Entries)[i].WordsTitle + ":" + (*Entries)[i].LengthTitle

		if UniqueTitleLinesWordsMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			(*Entries)[i].KeepReason = "title length + words"
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "css:" + (*Entries)[i].CountCssFiles

		if UniqueCssFilesMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			(*Entries)[i].KeepReason = "css files"
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "js:" + (*Entries)[i].CountJsFiles

		if UniqueJsFilesMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			(*Entries)[i].KeepReason = "js files"
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "tags:" + (*Entries)[i].CountTags

		if UniqueTagsMd5[Content] < 5 && ContentCounterMap[Content] < 2 {
			(*Entries)[i].KeepReason = "tags"
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	for i := 0; i < len(*Entries); i++ {
		if PositionsDone[i] {
			continue
		}

		Content := "status:" + strconv.Itoa((*Entries)[i].Status)

		if UniqueStatusMd5[Content] > 0 && ContentCounterMap[Content] < 2 {
			(*Entries)[i].KeepReason = "http status"
			TemporaryCleanedResults = append(TemporaryCleanedResults, (*Entries)[i])
			PositionsDone[i] = true
			ContentCounterMap[Content]++
		}
	}

	return TemporaryCleanedResults
}

func AnalyzeByWordsAndContentType(Entries *[]_struct.Result, i int, Hashes *map[string]int) {

	Content := "words-content-type:" + strconv.Itoa((*Entries)[i].Words) + ":" + (*Entries)[i].ContentType

	if (*Hashes)[Content] == 0 {
		(*Hashes)[Content] = 1
	} else {
		(*Hashes)[Content]++
	}
}

func AnalyzeByHttpStatus(Entries *[]_struct.Result, i int, StatusMd5 *map[string]int) {

	Content := "status:" + strconv.Itoa((*Entries)[i].Status)

	if (*StatusMd5)[Content] == 0 {
		(*StatusMd5)[Content] = 1
	} else {
		(*StatusMd5)[Content]++
	}

}

func AnalyzeByHttpStatusAndLength(Entries *[]_struct.Result, i int, StatusLengthMd5 *map[string]int) {
	Content := "status-length:" + strconv.Itoa((*Entries)[i].Status) + ":" + strconv.Itoa((*Entries)[i].Length)

	if (*StatusLengthMd5)[Content] == 0 {
		(*StatusLengthMd5)[Content] = 1
	} else {
		(*StatusLengthMd5)[Content]++
	}

}

func AnalyzeByHttpStatusAndHeadersCount(Entries *[]_struct.Result, i int, countMd5 *map[string]int) {
	Content := "status-header-count:" + strconv.Itoa((*Entries)[i].Status) + ":" + (*Entries)[i].CountHeaders

	if (*countMd5)[Content] == 0 {
		(*countMd5)[Content] = 1
	} else {
		(*countMd5)[Content]++
	}

}

func AnalyzeByTags(Entries *[]_struct.Result, i int, tagsMd5 *map[string]int) {
	Content := "tags:" + (*Entries)[i].CountTags

	if (*tagsMd5)[Content] == 0 {
		(*tagsMd5)[Content] = 1
	} else {
		(*tagsMd5)[Content]++
	}

}

func AnalyzeByJsFiles(Entries *[]_struct.Result, i int, filesMd5 *map[string]int) {
	Content := "js:" + (*Entries)[i].CountJsFiles

	if (*filesMd5)[Content] == 0 {
		(*filesMd5)[Content] = 1
	} else {
		(*filesMd5)[Content]++
	}
}

func AnalyzeByCssFiles(Entries *[]_struct.Result, i int, filesMd5 *map[string]int) {
	Content := "css:" + (*Entries)[i].CountCssFiles

	if (*filesMd5)[Content] == 0 {
		(*filesMd5)[Content] = 1
	} else {
		(*filesMd5)[Content]++
	}
}

func AnalyzeByTitleLengthWords(Entries *[]_struct.Result, i int, wordsMd5 *map[string]int) {
	Content := "title-length-words:" + (*Entries)[i].WordsTitle + ":" + (*Entries)[i].LengthTitle

	if (*wordsMd5)[Content] == 0 {
		(*wordsMd5)[Content] = 1
	} else {
		(*wordsMd5)[Content]++
	}
}

func AnalyzeByTitleWords(Entries *[]_struct.Result, i int, wordsMd5 *map[string]int) {
	Content := "title-words:" + (*Entries)[i].WordsTitle

	if (*wordsMd5)[Content] == 0 {
		(*wordsMd5)[Content] = 1
	} else {
		(*wordsMd5)[Content]++
	}
}

func AnalyzeByTitleLength(Entries *[]_struct.Result, i int, lengthMd5 *map[string]int) {
	Content := "title-length:" + (*Entries)[i].LengthTitle

	if (*lengthMd5)[Content] == 0 {
		(*lengthMd5)[Content] = 1
	} else {
		(*lengthMd5)[Content]++
	}
}

func AnalyzeByHttpStatusAndRedirectData(Entries *[]_struct.Result, i int, parameters *map[string]int) {
	Content := "status-redirect:" + strconv.Itoa((*Entries)[i].Status) + ":" + (*Entries)[i].RedirectDomain + ":" + (*Entries)[i].CountRedirectParameters

	if (*parameters)[Content] == 0 {
		(*parameters)[Content] = 1
	} else {
		(*parameters)[Content]++
	}
}

func AnalyzeByHttpStatusAndContentType(Entries *[]_struct.Result, i int, StatusCTMd5 *map[string]int) {
	Content := "status-content type:" + strconv.Itoa((*Entries)[i].Status) + ":" + (*Entries)[i].ContentType

	if (*StatusCTMd5)[Content] == 0 {
		(*StatusCTMd5)[Content] = 1
	} else {
		(*StatusCTMd5)[Content]++
	}
}

func AnalyzeByHttpStatusAndLines(Entries *[]_struct.Result, i int, StatusLinesMd5 *map[string]int) {
	Content := "status-lines:" + strconv.Itoa((*Entries)[i].Status) + ":" + strconv.Itoa((*Entries)[i].Lines)

	if (*StatusLinesMd5)[Content] == 0 {
		(*StatusLinesMd5)[Content] = 1
	} else {
		(*StatusLinesMd5)[Content]++
	}

}

func AnalyzeByHttpStatusAndWords(Entries *[]_struct.Result, i int, StatusWordsMd5 *map[string]int) {
	Content := "status-words:" + strconv.Itoa((*Entries)[i].Status) + ":" + strconv.Itoa((*Entries)[i].Words)

	if (*StatusWordsMd5)[Content] == 0 {
		(*StatusWordsMd5)[Content] = 1
	} else {
		(*StatusWordsMd5)[Content]++
	}

}
