package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

// Open a browser in user context and Automatically join Google Meet by clicking the button.
func main() {

	/*mac := false

	uD := "/home/drguru/.config/google-chrome/"
	bin := "/usr/bin/google-chrome"
	if mac {
		uD = "/Users/dkg/Library/Application Support/Google/Chrome"
		bin = "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	}
	//JoinNow := "#yDmH0d > c-wiz > div > div > div.TKU8Od > div.crqnQb > div > div.gAGjv > div.vgJExf > div > div > div.d7iDfe.NONs6c > div.shTJQe > div.jtn8y > div.XCoPyb > div > div > button > span.UywwFc-RLmnJb"
	//JoinNow = "#yDmH0d > c-wiz > div > div > div.TKU8Od > div.crqnQb > div > div.gAGjv > div.vgJExf > div > div > div.d7iDfe.NONs6c > div.shTJQe > div.jtn8y > div.XCoPyb > div > div > button"
	//JoinNow = "#yDmH0d > c-wiz > div > div > div.TKU8Od > div.crqnQb > div > div.gAGjv > div.vgJExf > div > div > div.d7iDfe.NONs6c > div.shTJQe > div.jtn8y > div.XCoPyb"
	*/
	// Launch your local browser first:
	//
	//     chrome  --remote-debugging-port=9222
	//

	u := "ws://127.0.0.1:8008/devtools/browser/e7bf2135-b964-4297-a0e7-8fe6906d00b3"

	b := rod.New().ControlURL(u).MustConnect()
	p := b.MustPages().MustFindByURL("meet.google.com")
	//p.MustActivate()
	pgs := []string{}
	// burger menu.
	pgs = append(pgs, "//*[@id=\"yDmH0d\"]/c-wiz/div/div/div[34]/div[4]/div[10]/div/div/div[2]/div/div[7]/div[4]/div[1]/span/button")
	// Settings button
	/*pgs = append(pgs, "/html/body/div[3]/div/div/ul/li[12]")
	//pgs = append(pgs, "/html/body/div[2]/div[4]/div[2]/div/div[2]/div/div/div[1]/div/div[2]/span/button")
	pgs = append(pgs, "//*[@id=\"yDmH0d\"]/div[4]/div[2]/div/div[2]/div/div/div[1]/div/div[2]/span/button")
	pgs = append(pgs, "/html/body/div[2]/div[4]/div[2]/div/div[2]/div/div/div[2]/div[2]/div/div[2]/div[2]/div[1]/div/div/span/span/div/div[1]/div/button")
	pgs = append(pgs, "/html/body/div[2]/div[4]/div[2]/div/div[2]/div/div/div[2]/div[2]/div/div[2]/div[2]/div[1]/div/div/span/span/div/div[2]/div/ul/li/ul/li[1]")
	pgs = append(pgs, "/html/body/div[2]/div[4]/div[2]/div/div[2]/div/div/div[1]/div/div[2]/span/button")
	*/
	cameras := []string{
		//	"/html/body/div[2]/div[4]/div[2]/div/div[2]/div/div/div[2]/div[2]/div/div[2]/div[2]/div[1]/div/div/span/span/div/div[2]/div/ul/li/ul/li[1]",
		//	"/html/body/div[2]/div[4]/div[2]/div/div[2]/div/div/div[2]/div[2]/div/div[2]/div[2]/div[1]/div/div/span/span/div/div[2]/div/ul/li/ul/li[2]",
		"//*[@id=\"yDmH0d\"]/div[4]/div[2]/div/div[2]/div/div/div[2]/div[2]/div/div[2]/div[2]/div[1]/div/div/span/span/div/div[2]/div/ul/li/ul/li[1]",
		"//*[@id=\"yDmH0d\"]/div[4]/div[2]/div/div[2]/div/div/div[2]/div[2]/div/div[2]/div[2]/div[1]/div/div/span/span/div/div[2]/div/ul/li/ul/li[2]",
	}
	//	/html/body/div[2]/div[3]/div[2]/div/div[2]/div/div/div[2]/div[2]/div/div[2]/div[2]/div[1]/div/div/span/span/div/div[2]/div/ul/li/ul/li[1]
	//	/html/body/div[2]/div[3]/div[2]/div/div[2]/div/div/div[2]/div[2]/div/div[2]/div[2]/div[1]/div/div/span/span/div/div[2]/div/ul/li/ul/li[1]
	//	/html/body/div[2]/div[3]/div[2]/div/div[2]/div/div/div[2]/div[2]/div/div[2]/div[2]/div[1]/div/div/span/span/div/div[2]/div/ul/li/ul/li[1]

	/*[@id=\"yDmH0d\"]/div[3]/div[2]/div/div[2]/div/div/div[2]/div[2]/div/div[2]/div[2]/div[1]/div/div/span/span/div/div[2]/div/ul/li/ul/li[1]
		/*for _, pg := range pgs {
		p.MustElementX(pg).MustClick()
		_ = pg
	}
	p.Activate()
	p.KeyActions().Press(input.Escape).MustDo()*/
	//p.MustElementR("span", "VYBDae-Bz112c-kBDsod-Rtc0Jf").MustClick()
	camera := 2

	p.MustElementX(pgs[0]).MustClick()
	p.MustElementR("li", "Settings").MustClick()
	p.MustElementR("button", "Video").MustClick()
	p.MustElementR("span", "Webcam").MustClick().Focus()
	for i := 0; i < camera; i++ {
		p.KeyActions().Press(input.ArrowDown).MustDo()
	}
	//p.KeyActions().Press(input.ArrowDown).MustDo()
	//p.KeyActions().Press(input.ArrowDown).MustDo()
	p.KeyActions().Press(input.Enter).MustDo()
	p.MustElementR("button", "Video").MustClick().Focus()

	p.KeyActions().Press(input.Escape).MustDo()
	//time.Sleep(time.Minute)
	_ = cameras[0]
}
