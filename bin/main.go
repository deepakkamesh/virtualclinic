package main

import (
	"time"

	"github.com/deepakkamesh/virtualclinic/sysagent"
	"github.com/go-rod/rod/lib/proto"
	"github.com/golang/glog"
)

func main() {
	glog.Info("Starting Virtual Clinic")
	bpath := "/home/drguru/.config/google-chrome/"
	bin := "/usr/bin/google-chrome"
	gvcid := "pym-jphe-rwg"
	joinN := "#yDmH0d > c-wiz > div > div > div.TKU8Od > div.crqnQb > div > div.gAGjv > div.vgJExf > div > div > div.d7iDfe.NONs6c > div.shTJQe > div.jtn8y > div.XCoPyb"

	b := sysagent.NewBrowser(bin, bpath,
		&proto.BrowserBounds{WindowState: proto.BrowserWindowStateMaximized})
	if err := b.Browser(); err != nil {
		glog.Fatalf("Failed to open Browser: %v", err)
	}
	if err := b.GVC(gvcid, joinN); err != nil {
		glog.Fatalf("Failed to start GVC %v", err)
	}
	time.Sleep(4 * time.Second)

	if err := b.InfoPage("https://example.com"); err != nil {
		glog.Errorf("failed info %v", err)
	}
	time.Sleep(4 * time.Second)
	b.FocusGVCPage()
	time.Sleep(2 * time.Second)
	b.SetInfoPageUrl("https://google.com")
	time.Sleep(2 * time.Second)
	//	b.Close()

	time.Sleep(time.Hour)
}
