/*
Packages
github.com/kenshaw/escpos - Images dont work, text good
github.com/augustopimenta/escpos - Works (no image)

github.com/cloudinn/escpos - Compilation problem (fork kenshaw)
github.com/epsimatic/escpos - Compilation problem (fork cloundinn)
github.com/conejoninja/go-escpos - Compilation issue
github.com/hennedo/escpos - compilation problem
*/
package main

import (
	"os"
	"time"

	"github.com/kenshaw/escpos"
)

func main() {
	/*f, err := os.OpenFile("/dev/usb/lp0", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	p := escpos.New(f)

	p.Init()
	p.SetSmooth(1)
	p.SetFontSize(2, 3)
	p.SetFont("A")
	p.Write("test ")
	p.SetFont("B")
	p.Write("test2 ")
	p.SetFont("C")
	p.Write("test3 ")
	p.Formfeed()

	p.SetFont("B")
	p.SetFontSize(1, 1)

	p.SetEmphasize(1)
	p.Write("halle")
	p.Formfeed()

	p.SetUnderline(1)
	p.SetFontSize(4, 4)
	p.Write("halle")

	p.SetReverse(1)
	p.SetFontSize(2, 4)
	p.Write("halle")
	p.Formfeed()

	p.SetFont("C")
	p.SetFontSize(8, 8)
	p.Write("halle")
	p.FormfeedN(5)

	p.Cut()
	p.End()*/
	if err := PrintScript("/dev/usb/lp0", "9840084500", "dr.guruswamy@gmail.com", "Muthukumaran", "20/M"); err != nil {
		panic(err)
	}
}

func PrintScript(devPrinter, drPhone, drEmail, patientName, sexAge string) error {
	f, err := os.OpenFile(devPrinter, os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	p := escpos.New(f)
	p.Init()
	p.SetSmooth(1)

	// Printer Header.
	p.SetFontSize(2, 2)
	p.SetFont("A")
	p.SetAlign("left")
	//p.SetEmphasize(2)
	p.Write("Dr. R. Guruswamy")
	p.Formfeed()

	p.SetFontSize(1, 1)
	p.SetAlign("center")
	p.Write("Ph: +919840084500")
	p.Write("  dr.guruswamy@gmail.com")
	p.Formfeed()

	p.SetFontSize(1, 1)
	p.SetAlign("center")
	p.Write("______________________________________________")
	p.Formfeed()

	loc, _ := time.LoadLocation("Asia/Kolkata") // Always print date/time in India time.
	now := time.Now().In(loc)
	date := now.Format("2 Jan 2006  3:04 pm")
	p.SetAlign("right")
	p.SetFontSize(1, 1)
	p.Write(date)
	p.Formfeed()

	p.SetAlign("left")
	p.SetFontSize(1, 1)
	p.Write(patientName)
	p.Write(sexAge)
	p.FormfeedN(2)

	p.SetFontSize(1, 1)
	p.SetAlign("left")
	p.Write("Patient presents with cold and sore throat. Needs to do throat cultuer and report back in 10 days. Take time off work.")
	p.FormfeedN(2)

	p.SetFontSize(1, 1)
	p.SetAlign("left")
	p.Write("Paracetamol - 3 times a day after food")
	p.Formfeed()
	p.Write("5 days ")
	p.FormfeedN(2)

	p.SetFontSize(1, 1)
	p.SetAlign("left")
	p.Write("Amoxcillin - 5 times a day after food")
	p.Formfeed()
	p.Write("5 days ")
	p.Formfeed()

	p.Cut()
	p.End()
	return nil
}
