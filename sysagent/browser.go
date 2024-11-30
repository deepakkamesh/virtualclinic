package sysagent

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type PageType int

const (
	GVCPage PageType = iota
	InfoPage
)

// Browser manages a browser.
type Browser struct {
	browserPath  string                 // Path to chrome browser binary.
	userDataDir  string                 // Path to chrome user data for user to emulate.
	browser      *rod.Browser           // Handler to the browser.
	windowLayout *proto.BrowserBounds   // Dimensions, size of window
	pages        map[PageType]*rod.Page // Handle to open pages (tabs)
}

// NewBrowser returns an initialized GVC object.
func NewBrowser(browserPath, userDataDir string, windowLayout *proto.BrowserBounds) *Browser {
	return &Browser{
		browserPath:  browserPath,
		userDataDir:  userDataDir,
		browser:      nil,
		windowLayout: windowLayout,
		pages:        make(map[PageType]*rod.Page),
	}
}

// openBrowser opens a non headless browser.
func (b *Browser) openBrowser() error {

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
	if b.browser == nil {
		return nil
	}

	if err := b.browser.Close(); err != nil {
		return err
	}
	b.browser = nil
	return nil
}

// GVC Open a GVC session with gvcID and JoinNowElem which is selector path for the JoinNow button.
func (b *Browser) GVC(gvcID, JoinNowElem string) error {
	// Open a new browser if not open already.
	if b.browser == nil {
		if err := b.openBrowser(); err != nil {
			return err
		}
	}
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

	b.pages[GVCPage] = p
	return nil
}

// InfoPage  opens a new tab with the url.
func (b *Browser) InfoPage(url string) error {
	// Open a new browser is one is not open.
	if b.browser == nil {
		if err := b.openBrowser(); err != nil {
			return err
		}
	}
	// Info Page is already open. Navigate to new url and set focus.
	if b.pages[InfoPage] != nil {
		if _, err := b.pages[InfoPage].Activate(); err != nil {
			return err
		}
		return b.pages[InfoPage].Navigate(url)
	}
	// Otherwise open new tab.
	p, err := b.browser.Page(proto.TargetCreateTarget{
		URL: url,
	})
	if err != nil {
		return err
	}
	if err := p.SetWindow(b.windowLayout); err != nil {
		return err
	}
	b.pages[InfoPage] = p
	return nil
}

// FocusPage brings focus to page  Window.
func (b *Browser) FocusPage(page PageType) error {
	if b.pages[page] == nil {
		return fmt.Errorf("cannot focus on tab thats not open")
	}
	if _, err := b.pages[page].Activate(); err != nil {
		return err
	}
	return nil
}

// ClosePage closes the tab on the browser.
func (b *Browser) ClosePage(page PageType) error {
	p, ok := b.pages[page]
	if ok && p == nil {
		return nil
	}
	if err := p.Close(); err != nil {
		return err
	}
	b.pages[page] = nil
	return nil
}

// ToggleMuteGVC mutes the remote GVC.
func (b *Browser) ToggleMuteGVC() error {
	p, ok := b.pages[GVCPage]
	if ok && p == nil {
		return fmt.Errorf("GVC page not open")
	}
	if err := p.KeyActions().Press(input.ControlLeft).Type('d').Do(); err != nil {
		return err
	}

	return nil
}

// SendEscKey sends the escape key to the specified page.
func (b *Browser) SendEscKey(page PageType) error {
	p, ok := b.pages[page]
	if ok && p == nil {
		return nil
	}
	if err := p.KeyActions().Press(input.Escape).Do(); err != nil {
		return err
	}
	return nil
}

// TODO: If otoscope camera works via GVC then add camera selector in GVC settings.
// TODO get screenshot of page to see locally.
