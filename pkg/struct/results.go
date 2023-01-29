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

	//@TODO: "json":false,
	//		 "matchers":{"IsCalibrated":false,"Mutex":{},"Matchers":{"status":{"value":"200,301,302,400,401,402,403,405,429,500,502,503,504"}},
	//		 "Filters":{"status":{"value":"404"}},
	//		 "PerDomainFilters":{}},
	//		 "mmode":"or",
	//		 "maxtime":1200,
	//		 "maxtime_job":600,
	//		 "method":"GET",
	//		 "noninteractive":false,
	//		 "outputdirectory":"/tmp/ffuf/bodies1",
	//		 "outputfile":"/tmp/ffuf/results.json",
	//		 "outputformat":"json",
	//		 "OutputSkipEmptyFile":false,
	//		 "proxyurl":"",
	//		 "quiet":false,
	//		 "rate":175,
	//		 "recursion":false,
	//		 "recursion_depth":0,
	//		 "recursion_strategy":
	//		 "default",
	//		 "replayproxyurl":"",
	//		 "sni":"",
	//		 "stop_429":true,
	//		 "stop_403":false,
	//		 "stop_all":false,
	//		 "stop_errors":false,
	//		 "threads":15,
	//		 "timeout":5,
	//		 "url":"https://hirejp.indeed.com/FUZZ",
	//		 "verbose":true,
	//		 "http2":false
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
