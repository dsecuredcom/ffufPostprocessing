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
	CountHeaders            any    `json:"count-headers"`
	RedirectDomain          any    `json:"redirect-domain"`
	CountRedirectParameters any    `json:"count-redirect-parameters"`
	LengthTitle             any    `json:"length-title"`
	WordsTitle              any    `json:"words-title"`
	CountCssFiles           any    `json:"count-css-files"`
	CountJsFiles            any    `json:"count-js-files"`
}
