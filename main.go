package main

import (
	"embed"
)

//go:embed frontend/dist
var assets embed.FS

var tray *Tray
var app *App

func main() {
	app = NewApp()
	tray = NewTray()
	tray.RunWith(app)
}
