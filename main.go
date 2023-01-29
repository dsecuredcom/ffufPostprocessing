package main

import (
	"encoding/json"
	"fmt"
	"github.com/Damian89/ffufPostprocessing/pkg/general"
	"github.com/Damian89/ffufPostprocessing/pkg/results"
	_struct "github.com/Damian89/ffufPostprocessing/pkg/struct"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {

	Configuration := general.GetArguments()
	fmt.Printf("\033[34m[i]\033[0m Original result file: %s\n", Configuration.OriginalFfufResultFile)

	if !general.FileExists(Configuration.OriginalFfufResultFile) {
		fmt.Printf("\033[31m[x]\033[0m Original result file does not exist: %s\n", Configuration.OriginalFfufResultFile)
		return
	}

	if Configuration.OverwriteResultFile {
		fmt.Printf("\033[34m[i]\033[0m Original result file will be overwritten: %s\n", Configuration.OriginalFfufResultFile)
	}

	if Configuration.OverwriteResultFile == false && Configuration.NewFfufResultFile == "" {
		fmt.Printf("\033[31m[x]\033[0m New result file is not set\n")
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

	if Configuration.DeleteAllBodyFiles && Configuration.FfufBodiesFolder != "" {
		fmt.Printf("\033[34m[!]\033[0m ALL bodies \033[31mwill be deleted\033[0m after analysis\n")
	} else if Configuration.DeleteUnnecessaryBodyFiles && Configuration.FfufBodiesFolder != "" {
		fmt.Printf("\033[34m[!]\033[0m Unnecessary bodies \033[31mwill be deleted\033[0m after analysis\n")
	}

	if Configuration.GenerateHtmlReport && Configuration.HtmlReportPath != "" {
		fmt.Printf("\033[34m[i] HTML Datatables report will ge generated and saved to: %s\033[0m\n", Configuration.HtmlReportPath)
	}

	fmt.Printf("\033[34m[i]\033[0m Loading results file\n")

	jsonFile := general.LoadJsonFile(Configuration.OriginalFfufResultFile)

	if jsonFile == nil {
		fmt.Printf("\033[31m[x]\033[0m Could not load original result file: %s\n", Configuration.OriginalFfufResultFile)
		return
	}

	defer jsonFile.Close()

	jsonByteValue, _ := io.ReadAll(jsonFile)

	var ResultsData _struct.Results

	json.Unmarshal(jsonByteValue, &ResultsData)

	fmt.Printf("\033[34m[i]\033[0m ResultsData file successfully parsed:\n")
	fmt.Printf("\033[34m[i]\033[0m Entries: %d\n", len(ResultsData.Results))

	results.EnrichResultsWithRedirectData(&ResultsData.Results)

	if general.FileExists(Configuration.FfufBodiesFolder) {
		fmt.Printf("\033[32m[i]\033[0m Enriching result data based on header/body of each request!\n")
		results.EnrichResults(Configuration.FfufBodiesFolder, &ResultsData.Results)
	}

	fmt.Printf("\033[32m[i]\033[0m Filtering results!\n")

	// Copy original json to new one and clean the results
	NewResultsData := ResultsData
	NewResultsData.Results = results.MinimizeOriginalResults(&ResultsData.Results)

	EntriesToKeep := map[int]bool{}

	for i := 0; i < len(NewResultsData.Results); i++ {
		EntriesToKeep[NewResultsData.Results[i].Position] = true
	}

	ResultFileNamesToBeDeleted := []string{}

	for i := 0; i < len(ResultsData.Results); i++ {
		_, ok := EntriesToKeep[ResultsData.Results[i].Position]
		if !ok {
			ResultFileNamesToBeDeleted = append(ResultFileNamesToBeDeleted, ResultsData.Results[i].Resultfile)
		}
	}

	fmt.Printf("\033[32m[i]\033[0m Filtering completed\n")
	fmt.Printf("\033[34m[i]\033[0m Filtered result count: %d\n", len(NewResultsData.Results))
	fmt.Printf("\033[34m[!]\033[0m \033[31mDeletable\033[0m: %d\n", len(ResultFileNamesToBeDeleted))
	fmt.Printf("\033[34m[i]\033[0m Writing new results file\n")

	NewResultsDataJson, _ := json.Marshal(NewResultsData)

	var jsonFileWriter *os.File

	if Configuration.OverwriteResultFile {
		jsonFileWriter = general.SaveJsonToFile(Configuration.OriginalFfufResultFile)
	} else if Configuration.OverwriteResultFile == false && Configuration.NewFfufResultFile != "" {
		jsonFileWriter = general.SaveJsonToFile(Configuration.NewFfufResultFile)
	} else {
		fmt.Printf("\033[31m[x]\033[0m Instructions related to writing results are unclear, either overwrite original file or allow creating a new one but not both!\n")
		return
	}

	if jsonFileWriter == nil {
		fmt.Printf("\u001B[31m[x]\u001B[0m Could not create new result file: %s\n", Configuration.NewFfufResultFile)
		return
	}

	defer jsonFileWriter.Close()

	jsonFileWriter.WriteString(
		string(NewResultsDataJson),
	)

	if Configuration.Verbose {
		for i := 0; i < len(NewResultsData.Results); i++ {
			general.PrintEntry(NewResultsData.Results[i])
		}
	}

	if Configuration.DeleteUnnecessaryBodyFiles == false && Configuration.DeleteAllBodyFiles == false {
		return
	}

	if Configuration.FfufBodiesFolder == "" {
		fmt.Printf("\033[31m[x]\033[0m Bodies folder is not set, cannot delete unnecessary files!\n")
		return
	}

	if Configuration.DeleteUnnecessaryBodyFiles {
		fmt.Printf("\033[32m[i]\033[0m Deleting unnecessary body files\n")
	} else {
		fmt.Printf("\033[32m[i]\033[0m Deleting all body files\n")
	}

	sem := make(chan struct{}, 25)
	var wg sync.WaitGroup
	for _, Filename := range ResultFileNamesToBeDeleted {
		wg.Add(1)

		go func(Filename string) {
			defer wg.Done()
			sem <- struct{}{}

			NormalizedPath := strings.TrimRight(Configuration.FfufBodiesFolder, "/\\")
			NormalizedPath = filepath.Join(NormalizedPath, Filename)

			if general.FileExists(NormalizedPath) {
				os.Remove(NormalizedPath)
			}
			<-sem
		}(Filename)
	}
	wg.Wait()

}
