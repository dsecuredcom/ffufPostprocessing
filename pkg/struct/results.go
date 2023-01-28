package _struct

type Results struct {
	Commandline string   `json:"commandline"`
	Time        string   `json:"time"`
	Results     []Result `json:"results"`
}

type Fuzz struct {
	Fuzz string `json:"FUZZ"`
}

type Result struct {
	Fuzz                    Fuzz   `json:"input"`
	Position                int    `json:"position"`
	Status                  int    `json:"status"`
	Length                  int    `json:"length"`
	Words                   int    `json:"words"`
	Lines                   int    `json:"lines"`
	ContentType             string `json:"content-type"`
	RedirectLocation        string `json:"redirectlocation"`
	Resultfile              string `json:"resultfile"`
	Url                     string `json:"url"`
	Host                    string `json:"host"`
	CountHeaders            string `json:"count-headers"`
	RedirectDomain          string `json:"redirect-domain"`
	CountRedirectParameters string `json:"count-redirect-parameters"`
	LengthTitle             string `json:"length-title"`
	WordsTitle              string `json:"words-title"`
	CountCssFiles           string `json:"count-css-files"`
	CountJsFiles            string `json:"count-js-files"`
	CountTags               string `json:"count-tags"`
	DropEntry               bool   `json:"drop-entry"`
}
