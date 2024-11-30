package main

import (
	"fmt"
	"log"

	"github.com/golang/glog"
	"github.com/itchyny/volume-go"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	glog.Info("Starting Virtual Clinic")

	/* Test for Browser
	bpath := "/home/drguru/.config/google-chrome/"
	bin := "/usr/bin/google-chrome"
	gvcid := "pym-jphe-rwg"
	joinN := "#yDmH0d > c-wiz > div > div > div.TKU8Od > div.crqnQb > div > div.gAGjv > div.vgJExf > div > div > div.d7iDfe.NONs6c > div.shTJQe > div.jtn8y > div.XCoPyb"

	b := sysagent.NewBrowser(bin, bpath,
		&proto.BrowserBounds{WindowState: proto.BrowserWindowStateMaximized})

	if err := b.GVC(gvcid, joinN); err != nil {
		glog.Fatalf("Failed to start GVC %v", err)
	}
	time.Sleep(4 * time.Second)

	if err := b.InfoPage("https://example.com"); err != nil {
		glog.Errorf("failed info %v", err)
	}
	time.Sleep(4 * time.Second)
	b.FocusPage(sysagent.GVCPage)
	time.Sleep(2 * time.Second)
	b.InfoPage("https://google.com")
	time.Sleep(2 * time.Second)
	b.ClosePage(sysagent.InfoPage)
	time.Sleep(2 * time.Second)
	time.Sleep(time.Hour) */

	/* Test for WhatsApp
	wa := sysagent.NewWhatsApp()

	qrChan, err := wa.Login(false)
	if err != nil {
		fmt.Printf("Login failed %v", err)
		return
	}
	if qrChan == nil {
		fmt.Println("Login Success")
	} else {
		fmt.Printf("not logged in. Need to auth")
		for evt := range qrChan {
			switch evt.Event {
			case "code":
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			case "error":
				fmt.Printf("error login %v\n", evt.Error)
			default:
				if evt == whatsmeow.QRChannelSuccess {
					fmt.Println("Auth success & Login")
					continue
				}
				fmt.Printf("Something else %v", evt.Event)
			}
		}
	}

	time.Sleep(3 * time.Second)

	if err := wa.SendMessage("16024050044", "msg 1"); err != nil {
		fmt.Printf("Send error %v", err)
	}
	time.Sleep(3 * time.Second)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	if err := wa.Logout(); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("done")*/

	vol, err := volume.GetVolume()
	if err != nil {
		log.Fatalf("get volume failed: %+v", err)
	}
	fmt.Printf("current volume: %d\n", vol)

	err = volume.SetVolume(10)
	if err != nil {
		log.Fatalf("set volume failed: %+v", err)
	}
	fmt.Printf("set volume success\n")

	err = volume.Mute()
	if err != nil {
		log.Fatalf("mute failed: %+v", err)
	}

	err = volume.Unmute()
	if err != nil {
		log.Fatalf("unmute failed: %+v", err)
	}
}
