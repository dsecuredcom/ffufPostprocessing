package main

import (
	"encoding/json"
	"fmt"
	"github.com/Damian89/ffufPostprocessing/pkg/general"
	"github.com/Damian89/ffufPostprocessing/pkg/results"
	_struct "github.com/Damian89/ffufPostprocessing/pkg/struct"
	"io"
)

func main() {

	Configuration := general.GetArguments()
	fmt.Printf("\033[34m[i]\033[0m Original result file: %s\n", Configuration.OriginalFfufResultFile)

	if !general.FileExists(Configuration.OriginalFfufResultFile) {
		fmt.Printf("\033[31m[x]\033[0m Original result file does not exist: %s\n", Configuration.OriginalFfufResultFile)
		return
	}

	fmt.Printf("\033[34m[i]\033[0m New result file: %s\n", Configuration.NewFfufResultFile)

	if Configuration.FfufBodiesFolder != "" {
		fmt.Printf("\033[34m[i]\033[0m Bodies folder: %s\n", Configuration.FfufBodiesFolder)
	}

	if Configuration.FfufBodiesFolder != "" && !general.FileExists(Configuration.FfufBodiesFolder) {
		fmt.Printf("\033[31[x]\033[0m Folder with bodies does not exist! Stopping here!\n")
		return
	}

	if Configuration.DeleteUnnecessaryBodyFiles {
		fmt.Printf("\033[34m[!]\033[0m Unnecessary bodies \033[31mwill be deleted\033[0m after analysis\n")
	}

	fmt.Printf("\033[34m[i]\033[0m Loading results file\n")

	jsonFile := general.LoadJsonFile(Configuration.OriginalFfufResultFile)

	if jsonFile == nil {
		fmt.Printf("\u001B[31m[x]\u001B[0m Could not load original result file: %s\n", Configuration.OriginalFfufResultFile)
		return
	}

	defer jsonFile.Close()

	jsonByteValue, _ := io.ReadAll(jsonFile)

	var ResultsData _struct.Results

	json.Unmarshal(jsonByteValue, &ResultsData)

	fmt.Printf("\033[34m[i]\033[0m ResultsData file successfully parsed:\n")
	fmt.Printf("\033[34m[i]\033[0m Entries: %d\n", len(ResultsData.Results))

	if general.FileExists(Configuration.FfufBodiesFolder) {
		fmt.Printf("\033[32m[i]\033[0m Enriching result data based on header/body of each request!\n")
		results.EnrichResults(Configuration.FfufBodiesFolder, &ResultsData.Results)
	}

	for i := 0; i < len(ResultsData.Results); i++ {
		fmt.Println(ResultsData.Results[i])
	}

	// determine which metadata type has no uniqueness
	// make json file entries unique

}
