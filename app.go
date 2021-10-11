package main

import (
	"context"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx  context.Context
	tray *Tray
}

// NewApp creates a new App application struct
// NewApp 创建一个新的 App 应用程序
func NewApp() *App {
	return &App{}
}

func (a *App) bindTray(tray *Tray) {
	a.tray = tray
}

func (a *App) run() {
	// Create application with options
	// 使用选项创建应用
	err := wails.Run(&options.App{
		Title:             getMessage("title"),
		Width:             1024,
		Height:            640,
		MinWidth:          1024,
		MinHeight:         640,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         true,
		StartHidden:       false,
		HideWindowOnClose: false,
		RGBA: &options.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 0,
		},
		Assets:     assets,
		LogLevel:   logger.DEBUG,
		OnStartup:  a.startup,
		OnDomReady: a.domReady,
		OnShutdown: a.shutdown,
		Bind: []interface{}{
			a,
		},
		// Windows platform specific options
		// Windows平台特定选项
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			DisableWindowIcon:    false,
		},
		// Mac: &mac.Options{
		// 	WebviewIsTransparent:          true,
		// 	WindowBackgroundIsTranslucent: true,
		// 	TitleBar:                      mac.TitleBarHiddenInset(),
		// },
	})

	if err != nil {
		log.Fatal(err)
	}
}

// startup is called at application startup
// startup 在应用程序启动时调用
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	// 在这里执行初始化设置
	a.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
// domReady 在前端Dom加载完毕后调用
func (a *App) domReady(ctx context.Context) {
	// Add your action here
	// 在这里添加你的操作
	a.ctx = ctx
}

// shutdown is called at application termination
// 在应用程序终止时被调用
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	// 在此处做一些资源释放的操作
	a.ctx = ctx
}

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
