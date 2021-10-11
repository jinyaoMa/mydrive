package main

import "github.com/wailsapp/wails/v2/pkg/runtime"

// Bind frontend to response to system tray language change
// 绑定到前端，让前端可以响应系统托盘的语言切换事件
func (a *App) ChangeLanguage(lang string) {
	switch lang {
	case "en":
		if a.tray.ResetLocaleTo("en") {
			a.tray.ChooseLanguage(a.tray.languagesEn)
		}
	case "zh-Hans":
		if a.tray.ResetLocaleTo("zh-Hans") {
			a.tray.ChooseLanguage(a.tray.languagesZhHans)
		}
	}
	runtime.WindowSetTitle(a.ctx, getMessage("title"))
}
