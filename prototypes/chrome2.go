package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
)

// Open a browser in user context and Automatically join Google Meet by clicking the button.
func main() {

	u := "ws://127.0.0.1:9222/devtools/browser/33d29406-19dd-4735-8a75-ae460647d130"
	//gvcID := "ggq-etha-ewq"
	//joinNowSelector := "#yDmH0d > c-wiz > div > div > div.TKU8Od > div.crqnQb > div > div.gAGjv > div.vgJExf > div > div > div.d7iDfe.NONs6c > div.shTJQe > div.jtn8y > div.XCoPyb"
	//	optionsXpath := "//*[@id=\"yDmH0d\"]/c-wiz/div/div/div[64]/div[3]/div/div[7]/div/div/div[2]/div/div[7]/div[3]/div[1]/span/button/div"
	optionsSel := "#yDmH0d > c-wiz > div > div > div.TKU8Od > div.crqnQb > div > div:nth-child(8) > div > div > div.Tmb7Fd > div > div.UvDM0e > div.tB5Jxf-xl07Ob-XxIAqe-OWXEXe-oYxtQd > div:nth-child(1) > span > button"
	//optionsSel := "#ow19 > div.tB5Jxf-xl07Ob-XxIAqe-OWXEXe-oYxtQd" // Outer div of more options.
	//#yDmH0d > c-wiz > div > div > div.TKU8Od > div.crqnQb > div > div:nth-child(8) > div > div > div.Tmb7Fd > div > div.UvDM0e > div.tB5Jxf-xl07Ob-XxIAqe-OWXEXe-oYxtQd > div:nth-child(1) > span > button > div

	/*p, e := openBrowser(u, gvcID)
	if e != nil {
		log.Printf("%v", e)
	}*/

	p, e := getPage(u)
	if e != nil {
		log.Printf("Failed to get page: %v", e)
		return
	}

	/*if err := GVC(p); err != nil {
		log.Printf("Failed to join GVC: %v", err)
		return
	}*/

	if err := switchCam(p, optionsSel); err != nil {
		log.Printf("Failed to join switch: %v", err)
		return
	}

	/*if err := selCam(p); err != nil {
		log.Printf("Failed to select camera: %v", err)
		return
	}*/

	time.Sleep(time.Minute)
}

func openBrowser(u, gvcID string) (*rod.Page, error) {
	// Launch a new browser instance.
	browser := rod.New().ControlURL(u).NoDefaultDevice().SlowMotion(500 * time.Millisecond)
	if err := browser.Connect(); err != nil {
		return nil, err
	}
	p, err := browser.Page(proto.TargetCreateTarget{
		URL: "https://meet.google.com/" + gvcID,
	})
	if err != nil {
		return nil, err
	}
	return p, nil
}

func getPage(u string) (*rod.Page, error) {
	// Connect to the existing browser instance.
	browser := rod.New().ControlURL(u).NoDefaultDevice().SlowMotion(500 * time.Millisecond)
	if err := browser.Connect(); err != nil {
		return nil, err
	}

	pgs, err := browser.Pages()
	if err != nil {
		return nil, err
	}
	p, err := pgs.FindByURL("meet.google.com")
	if err != nil {
		return nil, fmt.Errorf("page not found") // No page found.
	}
	return p, nil
}

// GVC opens a Google Meet session with the given gvcID and clicks the Join Now button.
func GVC(p *rod.Page) error {
	time.Sleep(3 * time.Second)
	if err := rod.Try(func() {
		p.Timeout(2*time.Second).MustElementR("button", "Join now").MustClick()
	}); err != nil {
		return err
	}
	return nil
}

func switchCam(p *rod.Page, x string) error {
	// Click the camera switch button.
	//time.Sleep(2 * time.Second)
	if err := rod.Try(func() {
		//p.Timeout(2*time.Second).MustElementR("span", "More options").MustClick()
		p.Timeout(2 * time.Second).MustElement(x).MustClick()
		p.Timeout(2*time.Second).MustElementR("li", "Settings").MustClick() // Click Settings.
		p.MustElementR("button", "Video").MustClick()                       // Click Video.
		p.MustElementR("button", "Webcam").MustClick()
		p.KeyActions().Press(input.ArrowDown).MustDo()
		p.KeyActions().Press(input.ArrowDown).MustDo()

		p.KeyActions().Press(input.Enter).MustDo()

	}); err != nil {
		return err
	}
	return nil
}

func selCam(p *rod.Page) error {
	if err := rod.Try(func() {
		p.KeyActions().Press(input.ArrowDown).MustDo()
	}); err != nil {
		return err
	}
	return nil
}
