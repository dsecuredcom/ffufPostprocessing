package general

import (
	"flag"
	_struct "github.com/Damian89/ffufPostprocessing/pkg/struct"
)

func GetArguments() _struct.Configuration {

	OriginalFfufResultFile := flag.String("result-file", "", "Path to the original ffuf result file")
	NewFfufResultFile := flag.String("new-result-file", "", "Path to the new ffuf result file")
	FfufBodiesFolder := flag.String("bodies-folder", "", "Path to the ffuf bodies folder")
	DeleteUnnecessaryBodyFiles := flag.Bool("delete-bodies", false, "Delete unnecessary body files")
	DeleteAllBodyFiles := flag.Bool("delete-all-bodies", false, "Delete unnecessary body files")
	OverwriteResultFile := flag.Bool("overwrite-result-file", false, "Overwrite original result file")
	GenerateHtmlReport := flag.Bool("generate-html-report", false, "Generate HTML report")
	HtmlReportPath := flag.String("html-report-path", "", "Path to the HTML report")

	flag.Parse()

	return _struct.Configuration{
		OriginalFfufResultFile:     *OriginalFfufResultFile,
		NewFfufResultFile:          *NewFfufResultFile,
		FfufBodiesFolder:           *FfufBodiesFolder,
		DeleteUnnecessaryBodyFiles: *DeleteUnnecessaryBodyFiles,
		DeleteAllBodyFiles:         *DeleteAllBodyFiles,
		OverwriteResultFile:        *OverwriteResultFile,
		GenerateHtmlReport:         *GenerateHtmlReport,
		HtmlReportPath:             *HtmlReportPath,
	}
}
