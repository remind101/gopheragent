package gopheragent

import (
	"fmt"
	"regexp"
	"strings"
)

// Browsers
const (
	Konqueror   = "konqueror"
	Chrome      = "chrome"
	Safari      = "safari"
	Opera       = "opera"
	PS3         = "ps3"
	PSP         = "psp"
	Firefox     = "firefox"
	Lotus       = "lotus"
	Netscape    = "netscape"
	SeaMonkey   = "seamonkey"
	Thunderbird = "thunderbird"
	Outlook     = "outlook"
	Evolution   = "evolution"
	IEMobile    = "iemobile"
	IE          = "ie"
)

// Engines
const (
	Webkit = "webkit"
	Khtml  = "khtml"
	Presto = "presto"
	Gecko  = "gecko"
	Msie   = "msie"
)

// Platforms
const (
	Windows      = "windows"
	Mac          = "macintosh"
	Linux        = "linux"
	Wii          = "wii"
	Playstation  = "playstation"
	Ipad         = "ipad"
	Ipod         = "ipod"
	Iphone       = "iphone"
	Android      = "android"
	Blackberry   = "blackberry"
	WindowsPhone = "windows_phone"
	Symbian      = "symbian"
)

// Unknown is returned when a result cannot be extracted
const Unknown = "unknown"

type regexpTest struct {
	Result  string
	Pattern *regexp.Regexp
	Expand  bool
}

var browsers,
	engines,
	oses,
	platforms []regexpTest
var browserVersions map[string]*regexp.Regexp
var mobilePlatforms []string

// UserAgent provides methods for extracting UA details
type UserAgent struct {
	s string
}

// New returns a UserAgent for the given UA string
func New(ua string) *UserAgent {

	result := UserAgent{
		s: strings.TrimSpace(ua),
	}

	return &result
}

// BrowserName returns the name of the browser from the user agent
func (ua *UserAgent) BrowserName() string {

	return matchFirst(browsers, ua.s)

}

// BrowserVersion returns the version of the browser from the user agent
func (ua *UserAgent) BrowserVersion() string {

	r, err := browserVersionRegexp(ua.BrowserName())
	if err == nil {
		matches := r.FindStringSubmatch(ua.s)

		if len(matches) > 1 {
			return matches[1]
		}
	}

	return ""
}

// Engine returns the rendering engine from the user agent
func (ua *UserAgent) Engine() string {

	return matchFirst(engines, ua.s)

}

// EngineVersion returns the version of the rendering engine used
func (ua *UserAgent) EngineVersion() string {

	engineVersion, err := regexp.Compile(`(?i:` + ua.Engine() + `[\/ ]([\d\w\.\-]+))`)

	if err == nil {
		matches := engineVersion.FindStringSubmatch(ua.s)

		if len(matches) > 1 {
			return matches[1]
		}

	}

	return ""

}

// OS returns the operating system from the user agent
func (ua *UserAgent) OS() string {

	return matchFirst(oses, ua.s)

}

// Platform returns the platform from the user agent
func (ua *UserAgent) Platform() string {

	return matchFirst(platforms, ua.s)

}

// Mobile returns true if the user agent represents a mobile client
func (ua *UserAgent) Mobile() bool {

	platform := ua.Platform()

	for _, t := range mobilePlatforms {
		if t == platform {
			return true
		}
	}

	return ua.BrowserName() == PSP

}

func browserVersionRegexp(b string) (r *regexp.Regexp, err error) {

	r, ok := browserVersions[b]
	if !ok {
		r, err = regexp.Compile(
			`(?i:` + b + `[\/ ]([\d\w\.\-]+)`,
		)
	}

	return r, err
}

func matchFirst(pp []regexpTest, ua string) string {

	for _, test := range pp {
		if m := test.Pattern.FindStringSubmatch(ua); m != nil {

			// see if we need to expand the result
			if test.Expand && len(m) > 1 {
				substrings := make([]interface{}, len(m)-1)
				for i := 1; i < len(m); i++ {
					substrings[i-1] = interface{}(m[i])
				}
				return fmt.Sprintf(test.Result, substrings...)
			}

			return test.Result
		}
	}

	return Unknown
}

func init() {

	browsers = []regexpTest{
		regexpTest{
			Result:  Konqueror,
			Pattern: regexp.MustCompile(`(?i:konqueror)`),
		},
		regexpTest{
			Result:  Chrome,
			Pattern: regexp.MustCompile(`(?i:chrome)`),
		},
		regexpTest{
			Result:  Safari,
			Pattern: regexp.MustCompile(`(?i:safari)`),
		},
		regexpTest{
			Result:  IE,
			Pattern: regexp.MustCompile(`(?i:msie)`)},
		regexpTest{
			Result:  Opera,
			Pattern: regexp.MustCompile(`(?i:opera)`),
		},
		regexpTest{
			Result:  PS3,
			Pattern: regexp.MustCompile(`(?i:playstation 3)`),
		},
		regexpTest{
			Result:  PSP,
			Pattern: regexp.MustCompile(`(?i:playstation portable)`),
		},
		regexpTest{
			Result:  Firefox,
			Pattern: regexp.MustCompile(`(?i:firefox)`),
		},
		regexpTest{
			Result:  Lotus,
			Pattern: regexp.MustCompile(`(?i:lotus.notes)`),
		},
		regexpTest{
			Result:  Netscape,
			Pattern: regexp.MustCompile(`(?i:netscape)`),
		},
		regexpTest{
			Result:  SeaMonkey,
			Pattern: regexp.MustCompile(`(?i:seamonkey)`),
		},
		regexpTest{
			Result:  Thunderbird,
			Pattern: regexp.MustCompile(`(?i:thunderbird)`),
		},
		regexpTest{
			Result:  Outlook,
			Pattern: regexp.MustCompile(`(?i:microsoft.outlook)`),
		},
		regexpTest{
			Result:  Evolution,
			Pattern: regexp.MustCompile(`(?i:evolution)`),
		},
		regexpTest{
			Result:  IEMobile,
			Pattern: regexp.MustCompile(`(?i:iemobile|windows phone)`),
		},
	}

	browserVersions = map[string]*regexp.Regexp{
		Chrome: regexp.MustCompile(`(?i:chrome\/([\d\w\.\-]+))`),
		Safari: regexp.MustCompile(`(?i:version\/([\d\w\.\-]+))`),
		PS3:    regexp.MustCompile(`(?i:([\d\w\.\-]+)\)\s*$)`),
		PSP:    regexp.MustCompile(`(?i:([\d\w\.\-]+)\)?\s*$)`),
		Lotus:  regexp.MustCompile(`(?i:Lotus-Notes\/([\w.]+))`),
	}

	engines = []regexpTest{
		regexpTest{
			Result:  Webkit,
			Pattern: regexp.MustCompile(`(?i:webkit)`),
		},

		regexpTest{
			Result:  Khtml,
			Pattern: regexp.MustCompile(`(?i:khtml)`),
		},
		regexpTest{
			Result:  Konqueror,
			Pattern: regexp.MustCompile(`(?i:konqueror)`),
		},
		regexpTest{
			Result:  Chrome,
			Pattern: regexp.MustCompile(`(?i:chrome)`),
		},
		regexpTest{
			Result:  Presto,
			Pattern: regexp.MustCompile(`(?i:presto)`),
		},
		regexpTest{
			Result:  Gecko,
			Pattern: regexp.MustCompile(`(?i:gecko)`),
		},
		regexpTest{
			Result:  Unknown,
			Pattern: regexp.MustCompile(`(?i:opera)`),
		},
		regexpTest{
			Result:  Msie,
			Pattern: regexp.MustCompile(`(?i:msie)`),
		},
	}

	oses = []regexpTest{
		regexpTest{
			Result:  "Windows Phone",
			Pattern: regexp.MustCompile(`(?i:windows (ce|phone|mobile)( os)?)`),
		},
		regexpTest{
			Result:  "Windows Vista",
			Pattern: regexp.MustCompile(`(?i:windows nt 6\.0)`),
		},
		regexpTest{
			Result:  "Windows 7",
			Pattern: regexp.MustCompile(`(?i:windows nt 6\.\d+)`),
		},
		regexpTest{
			Result:  "Windows 2003",
			Pattern: regexp.MustCompile(`(?i:windows nt 5\.2)`),
		},
		regexpTest{
			Result:  "Windows XP",
			Pattern: regexp.MustCompile(`(?i:windows nt 5\.1)`),
		},
		regexpTest{
			Result:  "Windows 2000",
			Pattern: regexp.MustCompile(`(?i:windows nt 5\.0)`),
		},
		regexpTest{
			Result:  "Windows",
			Pattern: regexp.MustCompile(`(?i:windows)`),
		},
		regexpTest{
			Result:  "OS X %s.%s",
			Pattern: regexp.MustCompile(`(?i:os x (\d+)[._](\d+))`),
			Expand:  true,
		},
		regexpTest{
			Result:  "Linux",
			Pattern: regexp.MustCompile(`(?i:linux)`),
		},
		regexpTest{
			Result:  "Wii",
			Pattern: regexp.MustCompile(`(?i:wii)`),
		},
		regexpTest{
			Result:  "Playstation",
			Pattern: regexp.MustCompile(`(?i:playstation 3)`),
		},
		regexpTest{
			Result:  "Playstation",
			Pattern: regexp.MustCompile(`(?i:playstation portable)`),
		},
		regexpTest{
			Result:  "iPad OS %s.%s",
			Pattern: regexp.MustCompile(`(?i:\(iPad.*os (\d+)[._](\d+))`),
			Expand:  true,
		},
		regexpTest{
			Result:  "iPhone OS %s.%s",
			Pattern: regexp.MustCompile(`(?i:\(iPhone.*os (\d+)[._](\d+))`),
			Expand:  true,
		},
		regexpTest{
			Result:  "Symbian OS",
			Pattern: regexp.MustCompile(`(?i:symbian(os)?)`),
		},
	}

	platforms = []regexpTest{
		regexpTest{
			Result:  WindowsPhone,
			Pattern: regexp.MustCompile(`(?i:windows (ce|phone|mobile)( os)?)`),
		},
		regexpTest{
			Result:  Windows,
			Pattern: regexp.MustCompile(`(?i:windows)`),
		},
		regexpTest{
			Result:  Mac,
			Pattern: regexp.MustCompile(`(?i:macintosh)`),
		},
		regexpTest{
			Result:  Android,
			Pattern: regexp.MustCompile(`(?i:android)`),
		},
		regexpTest{
			Result:  Blackberry,
			Pattern: regexp.MustCompile(`(?i:blackberry)`),
		},
		regexpTest{
			Result:  Linux,
			Pattern: regexp.MustCompile(`(?i:linux)`),
		},
		regexpTest{
			Result:  Wii,
			Pattern: regexp.MustCompile(`(?i:wii)`),
		},
		regexpTest{
			Result:  Playstation,
			Pattern: regexp.MustCompile(`(?i:playstation)`),
		},
		regexpTest{
			Result:  Ipad,
			Pattern: regexp.MustCompile(`(?i:ipad)`),
		},
		regexpTest{
			Result:  Ipod,
			Pattern: regexp.MustCompile(`(?i:ipod)`),
		},
		regexpTest{
			Result:  Iphone,
			Pattern: regexp.MustCompile(`(?i:iphone)`),
		},
		regexpTest{
			Result:  Symbian,
			Pattern: regexp.MustCompile(`(?i:symbian(os)?)`),
		},
	}

	mobilePlatforms = []string{
		Android,
		Blackberry,
		Ipad,
		Ipod,
		Iphone,
		Symbian,
		WindowsPhone,
	}
}
