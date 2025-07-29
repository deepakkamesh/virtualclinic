package main

import (
	vc "github.com/deepakkamesh/virtualclinic"
)

func main() {

	app := vc.NewClinic()
	app.ShowMainWindow()
	app.Run()
}
