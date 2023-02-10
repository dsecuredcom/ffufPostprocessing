package _struct

type Config struct {
	Autocalibration        bool   `json:"autocalibration"`
	AutocalibrationKeyword string `json:"autocalibration_keyword"`
	AutocalibrationPerHost bool   `json:"autocalibration_perhost"`
	AutocalibrationStratey string `json:"autocalibration_strategy"`
	AutocalibrationStrings any    `json:"autocalibration_strings"`
	Colors                 bool   `json:"colors"`
	ConfigFile             string `json:"configfile"`
	PostData               string `json:"postdata"`
	Delay                  any    `json:"delay"`
	DirsearchCompatibility bool   `json:"dirsearch_compatibility"`
	Extensions             any    `json:"extensions"`
	FMode                  string `json:"fmode"`
	FollowRedirects        bool   `json:"follow_redirects"`
	Headers                any    `json:"headers"`
	IgnoreBody             bool   `json:"ignorebody"`
	IgnoreWordlistComments bool   `json:"ignore_wordlist_comments"`
	InputMode              string `json:"inputmode"`
	CmdInput               string `json:"cmdinput"`
	InputProviders         any    `json:"inputproviders"`
	InputShell             string `json:"inputshell"`
	JsonOutput             bool   `json:"json"`
	Matchers               any    `json:"matchers"`
	Filters                any    `json:"Filters"`
	PerDomainFilters       any    `json:"PerDomainFilters"`
	MMode                  string `json:"mmode"`
	MaxTime                int    `json:"maxtime"`
	MaxTimeJob             int    `json:"maxtime_job"`
	Method                 string `json:"method"`
	NonInteractive         bool   `json:"noninteractive"`
	OutputDir              string `json:"outputdirectory"`
	OutputFile             string `json:"outputfile"`
	OutputFormat           string `json:"outputformat"`
	OutputSkipEmptyFiles   bool   `json:"OutputSkipEmptyFile"`
	ProxyURL               string `json:"proxyurl"`
	Quite                  bool   `json:"quite"`
	RateLimit              int    `json:"rate"`
	Recursion              bool   `json:"recursion"`
	RecursionDepth         int    `json:"recursion_depth"`
	RecursionStrategy      string `json:"recursion_strategy"`
	ReplayProxy            string `json:"replayproxyurl"`
	SNI                    string `json:"sni"`
	Stop429                bool   `json:"stop_429"`
	Stop403                bool   `json:"stop_403"`
	StopAll                bool   `json:"stop_all"`
	StopErrors             bool   `json:"stop_errors"`
	Threads                int    `json:"threads"`
	Timeout                int    `json:"timeout"`
	Url                    string `json:"url"`
	Verbose                bool   `json:"verbose"`
	Http2                  bool   `json:"http2"`
}

type Results struct {
	Commandline string   `json:"commandline"`
	Time        string   `json:"time"`
	Results     []Result `json:"results"`
	Config      Config   `json:"config"`
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
	KeepReason              string `json:"keepreason"`
}
