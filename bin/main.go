package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/deepakkamesh/virtualclinic"
	"github.com/golang/glog"
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

	/* Volume Management
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
	*/

	// SSH see sshtun.go in temp folder.

	/* Config File*/
	var conf virtualclinic.Config

	_, err := toml.DecodeFile("../virtualclinic.conf.toml", &conf)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v\n", conf)
}
