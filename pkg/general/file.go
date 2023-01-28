package general

import (
	"os"
	"strings"
)

func FileExists(File string) bool {

	if _, err := os.Stat(File); err == nil {
		return true
	}

	return false
}

func LoadJsonFile(File string) *os.File {
	jsonFile, err := os.Open(File)

	if err != nil {
		return nil
	}

	return jsonFile
}

func SaveJsonToFile(File string) *os.File {
	FileOs, err := os.Create(File)

	if err != nil {
		return nil
	}

	return FileOs
}

func DeleteFiles(folder string, Files []string) {
	NormalizedPath := strings.TrimRight(folder, "/") + "/"
	for i := 0; i < len(Files); i++ {
		CompletePath := NormalizedPath + Files[i]
		os.Remove(CompletePath)
	}
}
