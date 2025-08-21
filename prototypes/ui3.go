package main

import (
	"image/color"
	"log"
	"time"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	sc "github.com/deepakkamesh/virtualclinic/script"
)

type virtualclinic struct {
	app fyne.App
	win fyne.Window
}

func NewClinic() *virtualclinic {
	return &virtualclinic{
		app: app.New(),
	}
}

// Run starts the application and its a blocking call.
func (v *virtualclinic) Run() {
	//v.app.Settings().SetTheme()
	v.app.Run()
}

func (v *virtualclinic) Stop() {
	v.app.Quit()
}

func (v *virtualclinic) ShowMainWindow() {
	txtSz := float32(15)
	txtColor := color.White

	v.win = v.app.NewWindow("Virtual Clinic")

	// Top bar.
	patientInfo := container.NewVBox(
		newText("Muthu", fyne.TextStyle{Bold: true}, txtSz+5, txtColor),
		container.NewHBox(
			newLabel("Age:", fyne.TextStyle{Bold: true}),
			newText("45", fyne.TextStyle{Bold: true}, txtSz, txtColor),
			newLabel("Sex:", fyne.TextStyle{Bold: true}),
			newText("Male", fyne.TextStyle{Bold: true}, txtSz, txtColor),
			newLabel("Phone:", fyne.TextStyle{Bold: true}),
			newText("1234567890", fyne.TextStyle{Bold: true}, txtSz, txtColor),
		),
	)

	topBar := container.NewVBox(
		container.NewHBox(
			newText("Virtual Clinic", fyne.TextStyle{Bold: true}, 30, txtColor),
			layout.NewSpacer(),
			widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {}),
		),
		canvas.NewLine(txtColor),
		container.NewHBox(
			patientInfo,
			layout.NewSpacer(),
			widget.NewButton("Find Patient", func() { v.FindPatientWindow() }),
			widget.NewSeparator(),
			widget.NewButton("New Prescription", func() { v.ShowNewPrescriptionWindow() }),
			widget.NewSeparator(),
		),
	)
	patientInfo.Hide()
	// Left bar.
	leftBar := container.NewVBox(layout.NewSpacer())

	// Right bar.
	rightBar :=
		container.NewCenter(
			container.NewVBox(
				widget.NewButtonWithIcon("", theme.MediaFastRewindIcon(), func() {}),
				widget.NewSeparator(),
				widget.NewButtonWithIcon("", theme.MediaFastForwardIcon(), func() {}),
				widget.NewButton("Print", func() {}),
			),
		)

	// Content
	mainContent := container.NewPadded(
		canvas.NewRectangle(color.Black),
		widget.NewRichTextWithText("content \n kjkj"),
	)

	// Bottom bar.
	bottomBar := container.NewVBox(
		newText("(C) Mindfront Inc", fyne.TextStyle{Italic: true}, 10, color.RGBA{230, 210, 210, 255}),
		layout.NewSpacer(),
	)

	// Render the window.
	content := container.NewBorder(topBar, bottomBar, leftBar, rightBar, mainContent)
	v.win.SetContent(content)
	v.win.Resize(fyne.NewSize(800, 600))
	v.win.Show()
}

func newLabel(text string, style fyne.TextStyle) *widget.Label {
	l := widget.NewLabel(text)
	l.TextStyle = style
	return l
}

func newText(text string, style fyne.TextStyle, sz float32, color color.Color) *canvas.Text {
	l := canvas.NewText(text, color)
	l.TextStyle = style
	l.TextSize = sz
	return l
}

func (v *virtualclinic) FindPatientWindow() {
	name := widget.NewEntry()
	name.PlaceHolder = "Enter name"

	data := []sc.Patient{
		sc.Patient{"dkg", time.Now(), "M", "123"},
		sc.Patient{"sup", time.Now(), "F", "423"},
		sc.Patient{"mnv", time.Now(), "M", "523"},
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
	d := dialog.NewCustom("test", "sdsd", cn, v.win)

	d.Resize(fyne.NewSize(500, 400))
	d.Show()

}

func (v *virtualclinic) ShowNewPrescriptionWindow() {
	form := widget.NewForm()
	form.OnSubmit = func() {}
	w := widget.NewMultiLineEntry()
	w.Resize(fyne.NewSize(500, 600))

	form.Append("Name", widget.NewEntry())
	form.Append("Age", widget.NewEntry())
	form.Append("Sex", widget.NewEntry())
	form.Append("Script", w)

	d := dialog.NewCustom("New Prescription", "Enter details", form, v.win)
	d.Resize(fyne.NewSize(500, 400))
	d.Show()

}
