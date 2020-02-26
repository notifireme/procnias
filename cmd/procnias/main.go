package main

import (
	"github.com/notifireme/procnias"
)

func main() {
	app, err := procnias.NewApp()
	if err != nil {
		panic(err)
	}
	app.Run()
}
