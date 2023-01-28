package general

import (
	"fmt"
	_struct "github.com/Damian89/ffufPostprocessing/pkg/struct"
)

func PrintEntry(Result _struct.Result) {
	fmt.Printf(
		"DROP: %v, Status: %d; Length: %d; Words: %d; Lines: %d; CT: %s; RF: %s; CH: %s; RDo: %s; CRP: %s; LT: %s; WT: %s; CSS: %s; JS: %s\n",
		Result.DropEntry,
		Result.Status,
		Result.Length,
		Result.Words,
		Result.Lines,
		Result.ContentType,
		Result.Resultfile,
		Result.CountHeaders,
		Result.RedirectDomain,
		Result.CountRedirectParameters,
		Result.LengthTitle,
		Result.WordsTitle,
		Result.CountCssFiles,
		Result.CountJsFiles,
	)
}
