package gopheragent

import (
	"fmt"
	"regexp"
	"strings"
)

// Browsers
const (
	Electron    = "desktop"
	Chrome      = "chrome"
	Safari      = "safari"
	Opera       = "opera"
	Firefox     = "firefox"
	IEMobile    = "iemobile"
	IE          = "ie"
)

// Engines
const (
	Webkit = "webkit"
	Presto = "presto"
	Gecko  = "gecko"
	Msie   = "msie"
)

// Platforms
const (
	Windows      = "windows"
	Mac          = "macintosh"
	Linux        = "linux"
	Ipad         = "ipad"
	Ipod         = "ipod"
	Iphone       = "iphone"
	Android      = "android"
	Blackberry   = "blackberry"
	WindowsPhone = "windows_phone"
)

// Unknown is returned when a result cannot be extracted
const Unknown = "unknown"

// Our apps send user-agent strings that start with this.
const RemindPrefix = "Remind"

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
	iosBrowsers,
	iosEngines,
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
		if ua.IsRemind() {
			ua.engine = Unknown
			ua.browser = Unknown
		} else if ua.IsIos() {
			ua.browser = matchFirst(iosBrowsers, ua.s)
		} else {
			ua.browser = matchFirst(browsers, ua.s)
		}
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
		if ua.IsRemind() {
			ua.engine = Unknown
			ua.browser = Unknown
		} else if ua.IsIos() {
			ua.engine = matchFirst(iosEngines, ua.s)
		} else {
			ua.engine = matchFirst(engines, ua.s)
		}
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

func (ua *UserAgent) IsIos() bool {
	platform := ua.Platform()
	return platform == Iphone || platform == Ipad || platform == Ipod
}

func (ua *UserAgent) IsRemind() bool {
	return strings.HasPrefix(ua.s, RemindPrefix)
}

// Mobile returns true if the user agent represents a mobile client
func (ua *UserAgent) Mobile() bool {

	platform := ua.Platform()

	for _, t := range mobilePlatforms {
		if t == platform {
			return true
		}
	}

	return false

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
		if test.Expand {
			if m := test.Pattern.FindStringSubmatch(ua); m != nil {

				// see if we need to expand the result
				if len(m) > 1 {
					submatches := m[1:]
					args := make([]interface{}, len(submatches))

					for i, v := range submatches {
						args[i] = interface{}(v)
					}

					return fmt.Sprintf(test.Result, args...)
				} else {
					return test.Result
				}
			}

		} else if test.Pattern.MatchString(ua) {
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
			newSimpleTest(Chrome, `(?i:chrome)`),
			newSimpleTest(Safari, `(?i:safari)`),
			newSimpleTest(Opera, `(?i:opera)`),
			newSimpleTest(Firefox, `(?i:firefox)`),
			newSimpleTest(IEMobile, `(?i:iemobile|windows phone)`),
			newSimpleTest(IE, `(?i:msie)`),
		},
		fallback: Unknown,
	}

	iosBrowsers = regexpTestChain{
		tests: []*regexpTest{
			newSimpleTest(Safari, `(?i:safari)`),
		},
		fallback: Unknown,
	}

	browserVersions = map[string]*regexp.Regexp{
		Electron: regexp.MustCompile(`(?i:electron\/([\d\w\.\-]+))`),
		Chrome:   regexp.MustCompile(`(?i:chrome\/([\d\w\.\-]+))`),
		Safari:   regexp.MustCompile(`(?i:version\/([\d\w\.\-]+))`),
	}

	engines = regexpTestChain{
		tests: []*regexpTest{
			newSimpleTest(Webkit, `(?i:webkit)`),
			newSimpleTest(Chrome, `(?i:chrome)`),
			newSimpleTest(Gecko, `(?i:gecko)`),
			newSimpleTest(Msie, `(?i:msie)`),
			newSimpleTest(Presto, `(?i:presto)`),
			newSimpleTest(Opera, `(?i:opera)`),
		},
		fallback: Unknown,
	}

	iosEngines = regexpTestChain{
		tests: []*regexpTest{
			newSimpleTest(Webkit, `(?i:webkit)`),
		},
		fallback: Unknown,
	}

	oses = regexpTestChain{
		tests: []*regexpTest{
			newRegexpTest("iPad OS %s.%s", `(?i:\(iPad.*os (\d+)[._](\d+))`, true),
			newRegexpTest("iPhone OS %s.%s", `(?i:\(iPhone.*os (\d+)[._](\d+))`, true),
			newSimpleTest("Windows Phone", `(?i:windows (ce|phone|mobile)( os)?)`),
			newSimpleTest("Windows Vista", `(?i:windows nt 6\.0)`),
			newSimpleTest("Windows 7", `(?i:windows nt 6\.\d+)`),
			newSimpleTest("Windows 2003", `(?i:windows nt 5\.2)`),
			newSimpleTest("Windows XP", `(?i:windows nt 5\.1)`),
			newSimpleTest("Windows 2000", `(?i:windows nt 5\.0)`),
			newSimpleTest("Windows", `(?i:windows)`),
			newRegexpTest("OS X %s.%s", `(?i:os x (\d+)[._](\d+))`, true),
			newSimpleTest("Linux", `(?i:linux)`),
		},
		fallback: "Unknown",
	}

	platforms = regexpTestChain{
		tests: []*regexpTest{
			newSimpleTest(Ipad, `(?i:ipad)`),
			newSimpleTest(Ipod, `(?i:ipod)`),
			newSimpleTest(Iphone, `(?i:iphone)`),
			newSimpleTest(WindowsPhone, `(?i:windows (ce|phone|mobile)( os)?)`),
			newSimpleTest(Windows, `(?i:windows)`),
			newSimpleTest(Mac, `(?i:macintosh)`),
			newSimpleTest(Android, `(?i:android)`),
			newSimpleTest(Linux, `(?i:linux)`),
			newSimpleTest(Blackberry, `(?i:blackberry)`),
		},
		fallback: Unknown,
	}

	mobilePlatforms = []string{
		Android,
		Blackberry,
		Ipad,
		Ipod,
		Iphone,
		WindowsPhone,
	}
}
