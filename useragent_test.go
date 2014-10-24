package gopheragent_test

import (
	"testing"

	"."
)

type UserAgentTestCase struct {
	UA,
	BrowserName,
	BrowserVersion,
	Engine,
	EngineVersion,
	OS,
	Platform string
	Mobile bool
}

var testCases []UserAgentTestCase

func Test_UserAgent_Parse(t *testing.T) {

	for _, test := range testCases {
		var ua = gopheragent.New(test.UA)

		// browser name
		if got := ua.BrowserName(); got != test.BrowserName {
			t.Errorf("UserAgent.BrowserName[%s] => %s; want %s",
				test.UA,
				got,
				test.BrowserName,
			)
		}

		// browser version
		if got := ua.BrowserVersion(); got != test.BrowserVersion {
			t.Errorf("UserAgent.BrowserVersion[%s] => %s; want %s",
				test.UA,
				got,
				test.BrowserVersion,
			)
		}

		// engine
		if got := ua.Engine(); got != test.Engine {
			t.Errorf("UserAgent.Engine[%s] => %s; want %s",
				test.UA,
				got,
				test.Engine,
			)
		}

		// engine version
		if got := ua.EngineVersion(); got != test.EngineVersion {
			t.Errorf("UserAgent.EngineVersion[%s] => %s; want %s",
				test.UA,
				got,
				test.EngineVersion,
			)
		}

		// operating system
		if got := ua.OS(); got != test.OS {
			t.Errorf("UserAgent.OS[%s] => %s; want %s",
				test.UA,
				got,
				test.OS,
			)
		}

		// platform
		if got := ua.Platform(); got != test.Platform {
			t.Errorf("UserAgent.Platform[%s] => %s; want %s",
				test.UA,
				got,
				test.Platform,
			)
		}

		// mobile
		if got := ua.Mobile(); got != test.Mobile {
			t.Errorf("UserAgent.Mobile[%s] => %t; want %t",
				test.UA,
				got,
				test.Mobile,
			)
		}

	}
}

func init() {

	testCases = []UserAgentTestCase{

		UserAgentTestCase{
			UA:             "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_5) AppleWebKit/537.75.14 (KHTML, like Gecko) Version/6.1.3 Safari/537.75.14",
			BrowserName:    "safari",
			BrowserVersion: "6.1.3",
			Engine:         "webkit",
			EngineVersion:  "537.75.14",
			OS:             "OS X 10.8",
			Platform:       "macintosh",
			Mobile:         false,
		},

		UserAgentTestCase{
			UA:             "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.143 Safari/537.36",
			BrowserName:    "chrome",
			BrowserVersion: "36.0.1985.143",
			Engine:         "webkit",
			EngineVersion:  "537.36",
			OS:             "Windows 7",
			Platform:       "windows",
			Mobile:         false,
		},

		UserAgentTestCase{
			UA:             "Mozilla/5.0 (Linux; U; Android 4.1.2; es-us; SGH-T599N Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
			BrowserName:    "safari",
			BrowserVersion: "4.0",
			Engine:         "webkit",
			EngineVersion:  "534.30",
			OS:             "Linux",
			Platform:       "android",
			Mobile:         true,
		},

		UserAgentTestCase{
			UA:             "Mozilla/5.0 (Linux; U; Android 4.1.2; en-us; SGH-T999 Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
			BrowserName:    "safari",
			BrowserVersion: "4.0",
			Engine:         "webkit",
			EngineVersion:  "534.30",
			OS:             "Linux",
			Platform:       "android",
			Mobile:         true,
		},

		UserAgentTestCase{
			UA:             "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36",
			BrowserName:    "chrome",
			BrowserVersion: "36.0.1985.125",
			Engine:         "webkit",
			EngineVersion:  "537.36",
			OS:             "Windows 7",
			Platform:       "windows",
			Mobile:         false,
		},

		UserAgentTestCase{
			UA:             "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_1 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D201 Safari/9537.53",
			BrowserName:    "safari",
			BrowserVersion: "7.0",
			Engine:         "webkit",
			EngineVersion:  "537.51.2",
			OS:             "iPhone OS 7.1",
			Platform:       "iphone",
			Mobile:         true,
		},

		UserAgentTestCase{
			UA:             "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D167 Safari/9537.53",
			BrowserName:    "safari",
			BrowserVersion: "7.0",
			Engine:         "webkit",
			EngineVersion:  "537.51.2",
			OS:             "iPhone OS 7.1",
			Platform:       "iphone",
			Mobile:         true,
		},

		UserAgentTestCase{
			UA:             "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.73.11 (KHTML, like Gecko) Version/7.0.1 Safari/537.73.11",
			BrowserName:    "safari",
			BrowserVersion: "7.0.1",
			Engine:         "webkit",
			EngineVersion:  "537.73.11",
			OS:             "OS X 10.9",
			Platform:       "macintosh",
			Mobile:         false,
		},

		UserAgentTestCase{
			UA:             "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D257 Safari/9537.53",
			BrowserName:    "safari",
			BrowserVersion: "7.0",
			Engine:         "webkit",
			EngineVersion:  "537.51.2",
			OS:             "iPhone OS 7.1",
			Platform:       "iphone",
			Mobile:         true,
		},

		UserAgentTestCase{
			UA:             "Mozilla/5.0 (iPad; CPU OS 7_1_2 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D257 Safari/9537.53",
			BrowserName:    "safari",
			BrowserVersion: "7.0",
			Engine:         "webkit",
			EngineVersion:  "537.51.2",
			OS:             "iPad OS 7.1",
			Platform:       "ipad",
			Mobile:         true,
		},
	}

}
