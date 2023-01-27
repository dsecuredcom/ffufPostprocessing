package general

import (
	"os"
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
