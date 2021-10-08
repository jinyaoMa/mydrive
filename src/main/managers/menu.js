import { app, Menu } from 'electron'

class MenuManager {
  constructor(appManager) {
    this.appManager = appManager
    this.windowManager = appManager.windowManager
    this.translator = appManager.translator
    this.contextMenu = null
  }

  AppTrayMenu() {
    const $t = this.translator.get()

    const handleLanguageClick = (locale) => {
      if (this.contextMenu) {
        const temp = ['zh', 'en']
        temp.forEach((code) => {
          if (code !== locale) {
            this.contextMenu.getMenuItemById(code).checked = false
          }
        })
        this.appManager.languageChange(locale)
        this.windowManager.mainWindow.win.destroy()
        this.windowManager.mainWindow.createWindow()
      }
    }

    const handleWindowReset = (path = '#/') => {
      if (this.windowManager.mainWindow.win) {
        this.windowManager.mainWindow.initBrowserPage(path)
      } else {
        this.windowManager.mainWindow.createWindow(path)
      }

      // Execute electron window method
      this.windowManager.mainWindow.win.restore()
      this.windowManager.mainWindow.win.moveTop()
    }

    // Menu template
    this.contextMenu = Menu.buildFromTemplate([
      {
        label: $t('trayMenu.reset'),
        click() {
          handleWindowReset()
        }
      },
      {
        label: $t('trayMenu.settings'),
        click() {
          handleWindowReset('#/settings')
        }
      },
      {
        type: 'separator'
      },
      {
        label: $t('trayMenu.language.title'),
        enabled: false
      },
      {
        id: 'zh',
        label: $t('trayMenu.language.zh'),
        type: 'checkbox',
        click() {
          handleLanguageClick('zh')
        }
      },
      {
        id: 'en',
        label: $t('trayMenu.language.en'),
        type: 'checkbox',
        click() {
          handleLanguageClick('en')
        }
      },
      {
        type: 'separator'
      },
      {
        label: $t('trayMenu.exit'),
        click() {
          app.exit()
        }
      }
    ])

    this.contextMenu.getMenuItemById(this.translator.locale).checked = true

    return this.contextMenu
  }
}

export default MenuManager
