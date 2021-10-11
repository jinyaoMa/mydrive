package main

import (
	_ "embed"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const X = "by jinyaoMa"

//go:embed build/windows/icon.ico
var icon []byte

type Tray struct {
	app               *App
	openApp           *systray.MenuItem
	languages         *systray.MenuItem
	languagesEn       *systray.MenuItem
	languagesZhHans   *systray.MenuItem
	languagesSubItems []*systray.MenuItem
	quit              *systray.MenuItem
}

func NewTray() *Tray {
	return &Tray{}
}

func (t *Tray) RunWith(app *App) {
	t.app = app
	app.bindTray(t)
	systray.Register(t.onReady(), t.onExit())
	t.app.run()
}

func (t *Tray) ChooseLanguage(target *systray.MenuItem) {
	for _, item := range t.languagesSubItems {
		if item == target {
			target.Check()
		} else {
			item.Uncheck()
		}
	}
}

func (t *Tray) ResetLocaleTo(lang string) (ok bool) {
	ok = setLocale(lang)

	if ok {
		systray.SetTitle(getMessage("title"))
		systray.SetTooltip(getMessageWithParams("tooltip", X))

		t.openApp.SetTitle(getMessage("openApp"))
		t.openApp.SetTooltip(getMessage("openAppTooltip"))

		t.languages.SetTitle(getMessage("languages"))
		t.languages.SetTooltip(getMessage("languagesTooltip"))
		t.languagesEn.SetTitle(getMessage("languagesEn"))
		t.languagesEn.SetTooltip(getMessage("languagesEn"))
		t.languagesZhHans.SetTitle(getMessage("languagesZhHans"))
		t.languagesZhHans.SetTooltip(getMessage("languagesZhHans"))

		t.quit.SetTitle(getMessage("quit"))
		t.quit.SetTooltip(getMessage("quitTooltip"))
	}

	return
}

func (t *Tray) onReady() func() {
	return func() {
		systray.SetTemplateIcon(icon, icon)
		systray.SetTitle(getMessage("title"))
		systray.SetTooltip(getMessageWithParams("tooltip", X))

		t.openApp = systray.AddMenuItem(getMessage("openApp"), getMessage("openAppTooltip"))

		systray.AddSeparator()

		t.languages = systray.AddMenuItem(getMessage("languages"), getMessage("languagesTooltip"))
		t.languagesEn = t.languages.AddSubMenuItemCheckbox(getMessage("languagesEn"), getMessage("languagesEn"), locale == "en")
		t.languagesZhHans = t.languages.AddSubMenuItemCheckbox(getMessage("languagesZhHans"), getMessage("languagesZhHans"), locale == "zh-Hans")
		t.languagesSubItems = []*systray.MenuItem{
			t.languagesEn,
			t.languagesZhHans,
		}

		systray.AddSeparator()

		t.quit = systray.AddMenuItem(getMessage("quit"), getMessage("quitTooltip"))

		go func() {
			for {
				select {
				case <-t.openApp.ClickedCh:
					t.handleOpenApp()
				case <-t.languagesEn.ClickedCh:
					if t.ResetLocaleTo("en") {
						t.ChooseLanguage(t.languagesEn)
						t.handleLanguageChanged("en")
					}
				case <-t.languagesZhHans.ClickedCh:
					if t.ResetLocaleTo("zh-Hans") {
						t.ChooseLanguage(t.languagesZhHans)
						t.handleLanguageChanged("zh-Hans")
					}
				case <-t.quit.ClickedCh:
					systray.Quit()
					return
				}
			}
		}()
	}
}

func (t *Tray) onExit() func() {
	return func() {
		runtime.Quit(t.app.ctx)
	}
}

func (t *Tray) handleOpenApp() {
	runtime.WindowShow(t.app.ctx)
}

func (t *Tray) handleLanguageChanged(lang string) {
	runtime.EventsEmit(t.app.ctx, "languageChanged", lang)
	runtime.WindowSetTitle(t.app.ctx, getMessage("title"))
}
