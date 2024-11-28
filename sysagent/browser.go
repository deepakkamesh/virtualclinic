package sysagent

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

// Browser manages a browser.
type Browser struct {
	browserPath  string               // Path to chrome browser binary.
	userDataDir  string               // Path to chrome user data for user to emulate.
	browser      *rod.Browser         // Handler to the browser.
	gvcPage      *rod.Page            // Handler to the tab for GVC.
	infoPage     *rod.Page            // Handler to the info page.
	windowLayout *proto.BrowserBounds // Dimensions, size of window

}

// NewBrowser returns an initialized GVC object.
func NewBrowser(browserPath, userDataDir string, windowLayout *proto.BrowserBounds) *Browser {
	return &Browser{
		browserPath:  browserPath,
		userDataDir:  userDataDir,
		browser:      nil,
		windowLayout: windowLayout,
	}
}

// Browser opens a non headless browser.
func (b *Browser) Browser() error {

	l := launcher.New().Headless(false).UserDataDir(b.userDataDir).RemoteDebuggingPort(0).Delete("enable-automation").Bin(b.browserPath)
	u, err := l.Launch()
	if err != nil {
		return err
	}

	browser := rod.New().ControlURL(u).NoDefaultDevice()
	if err := browser.Connect(); err != nil {
		return err
	}
	b.browser = browser

	return nil
}

// Close the browser session.
func (b *Browser) Close() error {
	if err := b.browser.Close(); err != nil {
		return err
	}
	b.browser = nil
	return nil
}

// GVC Open a GVC session with gvcID and JoinNowElem which is selector path for the JoinNow button.
func (b *Browser) GVC(gvcID, JoinNowElem string) error {
	p, err := b.browser.Page(proto.TargetCreateTarget{
		URL: "https://meet.google.com/" + gvcID,
	})
	if err != nil {
		return err
	}

	time.Sleep(3 * time.Second) // Small delay to let the page load fully.
	if err := p.SetWindow(b.windowLayout); err != nil {
		return err
	}

	// Find element for JoinNow button and click on it.
	elem, err := p.Element(JoinNowElem)
	if err != nil {
		return err
	}
	if err := elem.Click(proto.InputMouseButtonLeft, 1); err != nil {
		return err
	}
	if err := p.WaitStable(time.Millisecond * 100); err != nil {
		return err
	}

	b.gvcPage = p
	return nil
}

// CloseGVC closes the GVC tab.
func (b *Browser) CloseGVC() error {

	if err := b.gvcPage.Close(); err != nil {
		return err
	}
	b.gvcPage = nil
	return nil
}

// FocusGVCPage Focus the GVC Window
func (b *Browser) FocusGVCPage() error {
	if _, err := b.gvcPage.Activate(); err != nil {
		return err
	}
	return nil
}

// InfoPage  opens a new tab with the url.
func (b *Browser) InfoPage(url string) error {
	p, err := b.browser.Page(proto.TargetCreateTarget{
		URL: url,
	})
	if err != nil {
		return err
	}
	if err := p.SetWindow(b.windowLayout); err != nil {
		return err
	}
	b.infoPage = p
	return nil
}

// SetInfoPageUrl navigates info url to the url anf brings focus to tab.
func (b *Browser) SetInfoPageUrl(url string) error {
	if b.infoPage == nil {
		return fmt.Errorf("info page is not available. Call InfoPage() first")
	}
	if _, err := b.infoPage.Activate(); err != nil {
		return err
	}
	return b.infoPage.Navigate(url)
}

// CloseInfoPage closes the info tab.
func (b *Browser) CloseInfoPage() error {
	if err := b.infoPage.Close(); err != nil {
		return err
	}
	b.infoPage = nil
	return nil
}

// FocusInfoPage brings focus to info page  Window.
func (b *Browser) FocusInfoPage() error {
	if _, err := b.infoPage.Activate(); err != nil {
		return err
	}
	return nil
}
