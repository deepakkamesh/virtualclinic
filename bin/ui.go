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
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()

	mainWin := app.NewWindow("Virtual Clinic")

	nameLabel := canvas.NewText("Name", color.White)
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}
	name := widget.NewEntry()
	name.PlaceHolder = "Enter name"
	findPatient := widget.NewButtonWithIcon("Find", theme.SearchIcon(), func() {
		FindPatientD(app, mainWin)

	})

	nameRow := container.NewGridWithColumns(2,
		container.New(layout.NewFormLayout(), nameLabel, name), findPatient)

	age := widget.NewLabel("12\n123\n343")
	sex := widget.NewLabel("M")
	phone := widget.NewLabel("23444")
	detailsRow := container.New(layout.NewHBoxLayout(), widget.NewLabel("Age:"), age, widget.NewLabel("Sex"), sex, widget.NewLabel("Phone"), phone)

	mainWin.SetContent(container.New(layout.NewVBoxLayout(), nameRow, detailsRow))

	mainWin.Show()
	mainWin.Resize(fyne.NewSize(1000, 800))
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
