package _struct

type Configuration struct {
	OriginalFfufResultFile     string
	NewFfufResultFile          string
	FfufBodiesFolder           string
	DeleteUnnecessaryBodyFiles bool
	DeleteAllBodyFiles         bool
	OverwriteResultFile        bool
	GenerateHtmlReport         bool
}
