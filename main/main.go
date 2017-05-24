package main

import (
	"github.com/remind101/gopheragent"
)

// the 100 most common user_agent strings in our event data.
var agents = []string{
	"Faraday v0.8.11",
	"Remind101/1236706 (iPhone8,1; iOS 10.3.1; Scale/2.00)",
	"Remind101/1236706 (iPhone7,2; iOS 10.3.1; Scale/2.00)",
	"Remind101/1236706 (iPhone9,1; iOS 10.3.1; Scale/2.00)",
	"Remind101/1236706 (iPhone8,2; iOS 10.3.1; Scale/3.00)",
	"Remind101/1236706 (iPhone9,2; iOS 10.3.1; Scale/3.00)",
	"Remind101/1236706 (iPhone9,3; iOS 10.3.1; Scale/2.00)",
	"Remind101/1236706 (iPhone6,1; iOS 10.3.1; Scale/2.00)",
	"Remind101/1236706 (iPhone9,4; iOS 10.3.1; Scale/3.00)",
	"Remind101/1236706 (iPhone7,1; iOS 10.3.1; Scale/3.00)",
	"Remind101/1236706 (iPhone7,2; iOS 10.2.1; Scale/2.00)",
	"Remind101/1236706 (iPhone8,1; iOS 10.2.1; Scale/2.00)",
	"Remind101/1236706 (iPhone8,4; iOS 10.3.1; Scale/2.00)",
	"Remind101/1236706 (iPhone6,1; iOS 10.2.1; Scale/2.00)",
	"Remind101/1236706 (iPhone8,1; iOS 10.2; Scale/2.00)",
	"Remind101/1236706 (iPhone8,2; iOS 10.2.1; Scale/3.00)",
	"Remind101/1236706 (iPhone9,1; iOS 10.2.1; Scale/2.00)",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
	"Remind101/1236706 (iPhone8,1; iOS 10.3.2; Scale/2.00)",
	"Remind101/1216345 (iPhone8,1; iOS 10.3.1; Scale/2.00)",
	"Remind101/1236706 (iPhone7,1; iOS 10.2.1; Scale/3.00)",
	"Remind101/1236706 (iPhone7,2; iOS 10.2; Scale/2.00)",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
	"Remind101/1236706 (iPhone7,2; iOS 10.3.2; Scale/2.00)",
	"Remind101/1236706 (iPhone9,3; iOS 10.2.1; Scale/2.00)",
	"Remind101/1216345 (iPhone7,2; iOS 10.3.1; Scale/2.00)",
	"Remind101/1236706 (iPhone9,2; iOS 10.2.1; Scale/3.00)",
	"Go-http-client/1.1",
	"Remind101/1236706 (iPhone9,4; iOS 10.2.1; Scale/3.00)",
	"Remind101/1236706 (iPhone5,3; iOS 10.2.1; Scale/2.00)",
	"Remind101/1236706 (iPhone7,2; iOS 10.1.1; Scale/2.00)",
	"Remind101/1236706 (iPhone5,3; iOS 10.3.1; Scale/2.00)",
	"Remind101/1236706 (iPhone8,4; iOS 10.2.1; Scale/2.00)",
	"Remind101/1236706 (iPhone8,2; iOS 10.3.2; Scale/3.00)",
	"Remind101/1236706 (iPhone7,2; iOS 10.0.2; Scale/2.00)",
	"Remind101/1236706 (iPhone8,1; iOS 10.1.1; Scale/2.00)",
	"Remind101/1236706 (iPhone9,1; iOS 10.2; Scale/2.00)",
	"Remind101/1216345 (iPhone9,1; iOS 10.3.1; Scale/2.00)",
	"Remind101/1196288 (iPhone8,1; iOS 10.3.1; Scale/2.00)",
	"Remind101/1216345 (iPhone8,2; iOS 10.3.1; Scale/3.00)",
	"Remind101/1236706 (iPhone6,1; iOS 10.2; Scale/2.00)",
	"Remind101/1236706 (iPhone8,1; iOS 10.0.2; Scale/2.00)",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1",
	"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
	"Remind101/1236706 (iPhone8,2; iOS 10.2; Scale/3.00)",
	"Remind101/1236706 (iPhone9,1; iOS 10.3.2; Scale/2.00)",
	"Remind101/1236706 (iPhone9,4; iOS 10.3.2; Scale/3.00)",
	"Remind101/1236706 (iPhone5,1; iOS 10.3.1; Scale/2.00)",
	"Remind101/1236706 (iPhone9,2; iOS 10.3.2; Scale/3.00)",
	"Remind101/1216345 (iPhone9,2; iOS 10.3.1; Scale/3.00)",
	"Remind101/1196288 (iPhone7,2; iOS 10.3.1; Scale/2.00)",
	"Remind101/1236706 (iPhone9,3; iOS 10.2; Scale/2.00)",
	"Remind101/1236706 (iPhone7,1; iOS 10.3.2; Scale/3.00)",
	"Remind101/1236706 (iPad5,3; iOS 10.3.1; Scale/2.00)",
	"Remind101/1175977 (iPhone8,1; iOS 10.3.1; Scale/2.00)",
	"Remind101/1236706 (iPhone9,3; iOS 10.3.2; Scale/2.00)",
	"Remind101/1137190 (iPhone8,1; iOS 10.3.1; Scale/2.00)",
	"Remind101/1216345 (iPhone8,1; iOS 10.2.1; Scale/2.00)",
	"Remind101/1216345 (iPhone9,3; iOS 10.3.1; Scale/2.00)",
	"Remind101/1236706 (iPhone4,1; iOS 9.3.5; Scale/2.00)",
	"Remind101/1236706 (iPhone5,2; iOS 10.3.1; Scale/2.00)",
	"Remind101/1216345 (iPhone9,4; iOS 10.3.1; Scale/3.00)",
	"Remind101/1236706 (iPod7,1; iOS 10.3.1; Scale/2.00)",
	"Remind101/1236706 (iPad4,1; iOS 10.3.1; Scale/2.00)",
	"Remind101/1216345 (iPhone7,2; iOS 10.2.1; Scale/2.00)",
	"Remind101/1236706 (iPhone9,2; iOS 10.2; Scale/3.00)",
	"Remind101/1236706 (iPod5,1; iOS 9.3.5; Scale/2.00)",
	"Remind101/1236706 (iPhone7,2; iOS 10.0.1; Scale/2.00)",
	"Remind101/1236706 (iPhone8,4; iOS 10.2; Scale/2.00)",
	"Remind101/1236706 (iPhone7,1; iOS 10.2; Scale/3.00)",
	"Remind101/1236706 (iPhone6,1; iOS 10.1.1; Scale/2.00)",
	"Remind101/1236706 (iPhone6,1; iOS 10.3.2; Scale/2.00)",
	"Remind101/1216345 (iPhone7,1; iOS 10.3.1; Scale/3.00)",
	"Remind101/1236706 (iPhone9,1; iOS 10.1.1; Scale/2.00)",
	"Remind101/1196288 (iPhone9,1; iOS 10.3.1; Scale/2.00)",
	"Remind101/1236706 (iPhone8,1; iOS 10.0.1; Scale/2.00)",
	"Remind101/1175977 (iPhone7,2; iOS 10.3.1; Scale/2.00)",
	"Remind101/1137190 (iPhone7,2; iOS 10.3.1; Scale/2.00)",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
	"Remind101/1216345 (iPhone6,1; iOS 10.3.1; Scale/2.00)",
	"Dalvik/2.1.0 (Linux; U; Android 7.0; SM-G930V Build/NRD90M)",
	"Remind101/1236706 (iPhone9,4; iOS 10.2; Scale/3.00)",
	"Remind101/1236706 (iPhone7,1; iOS 10.1.1; Scale/3.00)",
	"Remind101/1236706 (iPhone7,2; iOS 9.3.5; Scale/2.00)",
	"Remind101/1216345 (iPhone8,4; iOS 10.3.1; Scale/2.00)",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.1 Safari/603.1.30",
	"Remind101/1196288 (iPhone8,2; iOS 10.3.1; Scale/3.00)",
	"Remind101/1236706 (iPhone8,2; iOS 10.1.1; Scale/3.00)",
	"Remind101/1236706 (iPhone6,1; iOS 10.0.2; Scale/2.00)",
	"Remind101/1236706 (iPhone8,4; iOS 10.3.2; Scale/2.00)",
	"Remind101/1236706 (iPhone9,3; iOS 10.1.1; Scale/2.00)",
	"Remind101/1236706 (iPhone7,1; iOS 10.0.2; Scale/3.00)",
	"Remind101/1175977 (iPhone9,1; iOS 10.3.1; Scale/2.00)",
	"Remind101/1196288 (iPhone7,2; iOS 10.2.1; Scale/2.00)",
	"Remind101/1162141 (iPhone8,1; iOS 10.3.1; Scale/2.00)",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
	"Dalvik/2.1.0 (Linux; U; Android 7.0; SM-G920V Build/NRD90M)",
	"Remind101/1196288 (iPhone8,1; iOS 10.2.1; Scale/2.00)",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
}

// An unrealistic benchmark.
func test_gopheragent() {
	for i := 0; i < 1000; i++ {
		for _, agent := range agents {
			ua := gopheragent.New(agent)
			ua.BrowserName()
			ua.BrowserVersion()
			ua.Engine()
			ua.EngineVersion()
			ua.Mobile()
			ua.OS()
			ua.Platform()
		}
	}
}

func main() {
	test_gopheragent()
}
