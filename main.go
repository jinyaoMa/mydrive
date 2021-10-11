package main

var tray *Tray
var app *App

func main() {
	app = NewApp()
	tray = NewTray()
	tray.RunWith(app)
}
