package main

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	//u := launcher.New().Bin("/usr/bin/google-chrome").MustLaunch()
	u := launcher.New().Headless(false).RemoteDebuggingPort(37712).Delete("enable-automation").Bin("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome").MustLaunch()

	/*page := rod.New().ControlURL(u).MustConnect().MustPage("https://www.wikipedia.com")
	page.MustElement("#searchInput").MustInput("earth")
	page.MustElement("#search-form > fieldset > button > i").MustClick()*/

	page := rod.New().ControlURL(u).MustConnect().MustPage("https://meet.google.com")
	//page.MustElement("#searchInput").MustInput("earth")
	page.MustElement("#m2 > div > div > div._column_vqvmf_125._content_djoel_40._fixedRightPadding_djoel_91 > div._container_1lrvv_287._body_1lrvv_2._lg_1lrvv_38._body_djoel_100 > div._buttonGroup_dqfj5_2._buttonGroupHorizontal_dqfj5_14.button-group.horizontal > gws-button.breakpoints--desktop").MustClick()

	page.MustElement("#identifierId").MustInput("drguruswamyclinic@gmail.com")

	page.MustElement("#identifierNext > div > button > span").MustClick()

	page.MustElement("#password > div.aCsJod.oJeWuf > div > div.Xb9hP > input").MustInput("GuruSasi@4656")

	page.MustElement("#passwordNext > div > button > span").MustClick()

	page.MustElement("#yDmH0d > c-wiz > div > div.eEJIWe > div.x3toNe > div > div > div.UdVxgf.rcpVMd > div > div.bxxkd > div.sZjBXe > div > div.ix4A2e > div.Ufn6O.eFN2Jd > label").MustInput("pym-jphe-rwg")

	//page.MustElement("#yDmH0d > c-wiz > div > div.eEJIWe > div.x3toNe > div.cIzKaf > div > div.UdVxgf.rcpVMd > div > div.bxxkd > div.sZjBXe > div > div.ix4A2e > div.VfPpkd-dgl2Hf-ppHlrf-sM5MNb > button > span").MustClick()
	//	page.MustElement("#yDmH0d > c-wiz > div > div.eEJIWe > div.x3toNe > div.cIzKaf > div > div.UdVxgf.rcpVMd > div > div.bxxkd > div.sZjBXe > div > div.ix4A2e > div.VfPpkd-dgl2Hf-ppHlrf-sM5MNb > button > div.VfPpkd-RLmnJb").MustClick()
	page.MustElement("#yDmH0d > c-wiz > div > div.eEJIWe > div.x3toNe > div.cIzKaf > div > div.UdVxgf.rcpVMd > div > div.XG3Kfe > div.lpxrTc > div > div.VdLOD.Q2pHlf.JxfZTd > div > div.d5kC8b > div > div > div > div > div.taFJS").MustClick()

	page.MustElement("#yDmH0d > div.uW2Fw-Sx9Kwc.uW2Fw-Sx9Kwc-OWXEXe-n2to0e.uW2Fw-Sx9Kwc-OWXEXe-wdeprb-MD85tf-DKzjMe.UcM4Rc.YqDQLb.tTaNEd.uW2Fw-Sx9Kwc-OWXEXe-FNFY6c > div.uW2Fw-wzTsW > div > div > div > div > div.VlHPz > div > div:nth-child(1)").MustClick()
	//page.MustElement("#passwordNext > div > button > span").MustClick()
	page.MustWaitStable().MustScreenshot("a.png")
	_ = page
	time.Sleep(time.Hour)
}

//
