package main

import (
	"encoding/json"
	"fmt"
	"github.com/Damian89/ffufPostprocessing/pkg/general"
	_struct "github.com/Damian89/ffufPostprocessing/pkg/struct"
	"io"
)

func main() {

	Configuration := general.GetArguments()
	fmt.Printf("Original result file: %s\n", Configuration.OriginalFfufResultFile)
	fmt.Printf("New result file: %s\n", Configuration.NewFfufResultFile)
	fmt.Printf("Bodies folder: %s\n", Configuration.FfufBodiesFolder)
	fmt.Printf("Delete unnecessary body files: %t\n", Configuration.DeleteUnnecessaryBodyFiles)

	if !general.FileExists(Configuration.OriginalFfufResultFile) {
		fmt.Printf("Original result file does not exist: %s\n", Configuration.OriginalFfufResultFile)
		return
	}

	jsonFile := general.LoadJsonFile(Configuration.OriginalFfufResultFile)

	if jsonFile == nil {
		fmt.Printf("Could not load original result file: %s\n", Configuration.OriginalFfufResultFile)
		return
	}

	defer jsonFile.Close()

	jsonByteValue, _ := io.ReadAll(jsonFile)

	var results _struct.Results

	json.Unmarshal(jsonByteValue, &results)

	for i := 0; i < len(results.Results); i++ {
		fmt.Println(results.Results[i].Fuzz)
	}

	// case 1: body path exists
	// enrich loaded json with more meta data
	// case 2: body path does not exist
	// do nothing

	// determine which metadata type has no uniqueness
	// make json file entries unique

}
