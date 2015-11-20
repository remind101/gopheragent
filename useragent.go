package gopheragent

import (
	"fmt"
	"regexp"
	"strings"
)

// Browsers
const (
	Electron    = "desktop"
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
type regexpTestChain struct {
	tests    []*regexpTest
	fallback string
}

var browsers,
	engines,
	oses,
	platforms regexpTestChain
var browserVersions map[string]*regexp.Regexp
var mobilePlatforms []string

// UserAgent provides methods for extracting UA details
type UserAgent struct {
	s,
	browser,
	engine,
	os,
	platform string
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

	if ua.browser == "" {
		ua.browser = matchFirst(browsers, ua.s)
	}

	return ua.browser

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

	if ua.engine == "" {
		ua.engine = matchFirst(engines, ua.s)
	}

	return ua.engine

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

	if ua.os == "" {
		ua.os = matchFirst(oses, ua.s)
	}

	return ua.os

}

// Platform returns the platform from the user agent
func (ua *UserAgent) Platform() string {

	if ua.platform == "" {
		ua.platform = matchFirst(platforms, ua.s)
	}

	return ua.platform

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
			`(?i:` + b + `[\/ ]([\d\w\.\-]+))`,
		)
	}

	return r, err
}

func matchFirst(tt regexpTestChain, ua string) string {

	for _, test := range tt.tests {
		if m := test.Pattern.FindStringSubmatch(ua); m != nil {

			// see if we need to expand the result
			if test.Expand && len(m) > 1 {
				submatches := m[1:]
				args := make([]interface{}, len(submatches))

				for i, v := range submatches {
					args[i] = interface{}(v)
				}

				return fmt.Sprintf(test.Result, args...)
			}

			return test.Result
		}
	}

	return tt.fallback
}

func newRegexpTest(result, pattern string, expand bool) *regexpTest {
	return &regexpTest{
		Result:  result,
		Pattern: regexp.MustCompile(pattern),
		Expand:  expand,
	}
}

func newSimpleTest(result, pattern string) *regexpTest {
	return newRegexpTest(result, pattern, false)
}

func init() {

	browsers = regexpTestChain{
		tests: []*regexpTest{
			newSimpleTest(Electron, `(?i:electron)`),
			newSimpleTest(Konqueror, `(?i:konqueror)`),
			newSimpleTest(Chrome, `(?i:chrome)`),
			newSimpleTest(Safari, `(?i:safari)`),
			newSimpleTest(Opera, `(?i:opera)`),
			newSimpleTest(PS3, `(?i:playstation 3)`),
			newSimpleTest(PSP, `(?i:playstation portable)`),
			newSimpleTest(Firefox, `(?i:firefox)`),
			newSimpleTest(Lotus, `(?i:lotus.notes)`),
			newSimpleTest(Netscape, `(?i:netscape)`),
			newSimpleTest(SeaMonkey, `(?i:seamonkey)`),
			newSimpleTest(Thunderbird, `(?i:thunderbird)`),
			newSimpleTest(Outlook, `(?i:microsoft.outlook)`),
			newSimpleTest(Evolution, `(?i:evolution)`),
			newSimpleTest(IEMobile, `(?i:iemobile|windows phone)`),
			newSimpleTest(IE, `(?i:msie)`),
		},
		fallback: Unknown,
	}

	browserVersions = map[string]*regexp.Regexp{
		Electron: regexp.MustCompile(`(?i:electron\/([\d\w\.\-]+))`),
		Chrome:   regexp.MustCompile(`(?i:chrome\/([\d\w\.\-]+))`),
		Safari:   regexp.MustCompile(`(?i:version\/([\d\w\.\-]+))`),
		PS3:      regexp.MustCompile(`(?i:([\d\w\.\-]+)\)\s*$)`),
		PSP:      regexp.MustCompile(`(?i:([\d\w\.\-]+)\)?\s*$)`),
		Lotus:    regexp.MustCompile(`(?i:Lotus-Notes\/([\w.]+))`),
	}

	engines = regexpTestChain{
		tests: []*regexpTest{
			newSimpleTest(Webkit, `(?i:webkit)`),
			newSimpleTest(Khtml, `(?i:khtml)`),
			newSimpleTest(Konqueror, `(?i:konqueror)`),
			newSimpleTest(Chrome, `(?i:chrome)`),
			newSimpleTest(Presto, `(?i:presto)`),
			newSimpleTest(Gecko, `(?i:gecko)`),
			newSimpleTest(Unknown, `(?i:opera)`),
			newSimpleTest(Msie, `(?i:msie)`),
		},
		fallback: Unknown,
	}

	oses = regexpTestChain{
		tests: []*regexpTest{
			newSimpleTest("Windows Phone", `(?i:windows (ce|phone|mobile)( os)?)`),
			newSimpleTest("Windows Vista", `(?i:windows nt 6\.0)`),
			newSimpleTest("Windows 7", `(?i:windows nt 6\.\d+)`),
			newSimpleTest("Windows 2003", `(?i:windows nt 5\.2)`),
			newSimpleTest("Windows XP", `(?i:windows nt 5\.1)`),
			newSimpleTest("Windows 2000", `(?i:windows nt 5\.0)`),
			newSimpleTest("Windows", `(?i:windows)`),
			newRegexpTest("OS X %s.%s", `(?i:os x (\d+)[._](\d+))`, true),
			newSimpleTest("Linux", `(?i:linux)`),
			newSimpleTest("Wii", `(?i:wii)`),
			newSimpleTest("Playstation", `(?i:playstation 3)`),
			newSimpleTest("Playstation", `(?i:playstation portable)`),
			newRegexpTest("iPad OS %s.%s", `(?i:\(iPad.*os (\d+)[._](\d+))`, true),
			newRegexpTest("iPhone OS %s.%s", `(?i:\(iPhone.*os (\d+)[._](\d+))`, true),
			newSimpleTest("Symbian OS", `(?i:symbian(os)?)`),
		},
		fallback: "Unknown",
	}

	platforms = regexpTestChain{
		tests: []*regexpTest{
			newSimpleTest(WindowsPhone, `(?i:windows (ce|phone|mobile)( os)?)`),
			newSimpleTest(Windows, `(?i:windows)`),
			newSimpleTest(Mac, `(?i:macintosh)`),
			newSimpleTest(Android, `(?i:android)`),
			newSimpleTest(Blackberry, `(?i:blackberry)`),
			newSimpleTest(Linux, `(?i:linux)`),
			newSimpleTest(Wii, `(?i:wii)`),
			newSimpleTest(Playstation, `(?i:playstation)`),
			newSimpleTest(Ipad, `(?i:ipad)`),
			newSimpleTest(Ipod, `(?i:ipod)`),
			newSimpleTest(Iphone, `(?i:iphone)`),
			newSimpleTest(Symbian, `(?i:symbian(os)?)`),
		},
		fallback: Unknown,
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
