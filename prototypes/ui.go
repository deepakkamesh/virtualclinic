package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()

	mainWin := app.NewWindow("Virtual Clinic")

	lblName := widget.NewLabel("Name:")
	lblSex := widget.NewLabel("Sex:")
	lblAge := widget.NewLabel("Age:")
	lblPhone := widget.NewLabel("Phone:")

	appName := canvas.NewText("Virtual Clinic", color.White)

	//appName.Alignment = fyne.TextAlignTrailing
	appName.TextStyle = fyne.TextStyle{Bold: true}
	appName.TextSize = 30
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 2

	name := canvas.NewText("Muthu", color.White)
	name.Alignment = fyne.TextAlignTrailing
	name.TextStyle = fyne.TextStyle{Bold: true}
	name.TextSize = 20
	sex := canvas.NewText("Male", color.White)
	sex.Alignment = fyne.TextAlignTrailing
	sex.TextStyle = fyne.TextStyle{Bold: true}
	age := canvas.NewText("45", color.White)
	age.Alignment = fyne.TextAlignTrailing
	age.TextStyle = fyne.TextStyle{Bold: true}
	phone := canvas.NewText("1234567890", color.White)
	phone.Alignment = fyne.TextAlignTrailing
	phone.TextStyle = fyne.TextStyle{Bold: true}
	find := widget.NewButton("Find", func() {
		log.Println("tapped")
	})
	newScript := widget.NewButton("New Prescription", func() {
		log.Println("tapped")
	})
	contentTop := container.NewHBox(lblName, name, lblSex, sex, lblAge, age, lblPhone, phone, layout.NewSpacer(), find, widget.NewSeparator(), newScript)
	c := container.NewVBox(appName, line, layout.NewSpacer(), contentTop, canvas.NewLine(color.White))

	//top := canvas.NewText("top bar", color.White)
	left := canvas.NewText("left", color.White)
	middle := canvas.NewText("content", color.White)
	content := container.NewBorder(c, nil, left, nil, middle)
	mainWin.SetContent(content)

	mainWin.Show()
	mainWin.Resize(fyne.NewSize(800, 600))
	app.Run()

}

type Patient struct {
	Name  string // Name of Patient (mandatory).
	Sex   string // Sex of patient (mandatory). Acceptable values M/F
	Phone string // Phone number.
}

func FindPatient(a fyne.App) {

	name := widget.NewEntry()
	name.PlaceHolder = "Enter name"
	w := a.NewWindow("Find")

	data := []Patient{
		Patient{"dkg", "M", "123"},
		Patient{"sup", "F", "423"},
		Patient{"mnv", "M", "523"},
	}
	lst := widget.NewTableWithHeaders(
		func() (int, int) {
			return len(data), 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			switch i.Col {
			case 0:
				o.(*widget.Label).SetText(data[i.Row].Name)
			case 1:
				o.(*widget.Label).SetText(data[i.Row].Sex)
			case 2:
				o.(*widget.Label).SetText(data[i.Row].Phone)
			}
		})
	lst.OnSelected = func(id widget.TableCellID) {
		log.Println(id.Row, id.Col)
	}

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Name", Widget: name}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", name.Text)
			data[0].Name = "ddd"
			lst.Hidden = false
			lst.Refresh()

		},
		OnCancel: func() {
			w.Close()
		},
	}

	lst.Hidden = true
	//lst.SetColumnWidth(2, 50)
	w.SetContent(container.NewBorder(
		container.NewVBox(form, widget.NewSeparator()), nil, nil, nil, layout.NewSpacer(), lst))
	w.Show()

	w.Resize(fyne.NewSize(500, 300))
}

func FindPatientD(a fyne.App, pw fyne.Window) {

	name := widget.NewEntry()
	name.PlaceHolder = "Enter name"

	data := []Patient{
		Patient{"dkg", "M", "123"},
		Patient{"sup", "F", "423"},
		Patient{"mnv", "M", "523"},
	}
	lst := widget.NewTableWithHeaders(
		func() (int, int) {
			return len(data), 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			switch i.Col {
			case 0:
				o.(*widget.Label).SetText(data[i.Row].Name)
			case 1:
				o.(*widget.Label).SetText(data[i.Row].Sex)
			case 2:
				o.(*widget.Label).SetText(data[i.Row].Phone)
			}
		})
	lst.OnSelected = func(id widget.TableCellID) {
		log.Println(id.Row, id.Col)
	}

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Name", Widget: name}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", name.Text)
			data[0].Name = "ddd"
			lst.Hidden = false
			lst.Refresh()

		},
		OnCancel: func() {
			//	w.Close()
		},
	}

	lst.Hidden = false
	cn := container.NewGridWithRows(2, container.NewVBox(form, widget.NewSeparator()), lst)
	d := dialog.NewCustom("test", "sdsd", cn, pw)

	d.Resize(fyne.NewSize(500, 400))
	d.Show()

}
