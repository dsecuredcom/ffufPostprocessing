package results

import (
	"crypto/md5"
	"encoding/hex"
	_struct "github.com/Damian89/ffufPostprocessing/pkg/struct"
	"strconv"
)

func MinimizeOriginalResults(Entries *[]_struct.Result) {
	UniqueMd5Hashes := map[string]int{}

	for i := 0; i < len(*Entries); i++ {
		(*Entries)[i].DropEntry = false

		Content := strconv.Itoa((*Entries)[i].Status) + ":"
		Content += strconv.Itoa((*Entries)[i].Length) + ":"
		Content += strconv.Itoa((*Entries)[i].Words) + ":"
		Content += strconv.Itoa((*Entries)[i].Lines) + ":"
		Content += (*Entries)[i].ContentType + ":"
		Content += (*Entries)[i].RedirectDomain + ":"
		Content += (*Entries)[i].CountRedirectParameters + ":"

		// @TODO Those won't exist in the original results file
		// When bodies are not present, we way run into errors here
		Content += (*Entries)[i].CountHeaders + ":"
		Content += (*Entries)[i].LengthTitle + ":"
		Content += (*Entries)[i].WordsTitle + ":"
		Content += (*Entries)[i].CountCssFiles + ":"
		Content += (*Entries)[i].CountJsFiles + ":"
		Content += (*Entries)[i].CountTags

		hash := md5.Sum([]byte(Content))
		md5Hash := hex.EncodeToString(hash[:])

		if UniqueMd5Hashes[md5Hash] == 0 {
			UniqueMd5Hashes[md5Hash] = 1
		} else {
			UniqueMd5Hashes[md5Hash]++
		}

		if UniqueMd5Hashes[md5Hash] >= 5 {
			(*Entries)[i].DropEntry = true
		}
	}
	//
	//for hash, occurences := range UniqueMd5Hashes {
	//	fmt.Printf("Hash: %s, occurences: %d\n", hash, occurences)
	//}
}
