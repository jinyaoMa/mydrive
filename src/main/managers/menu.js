import { app, Menu } from 'electron'

class MenuManager {
  constructor(appManager) {
    this.appManager = appManager
    this.windowManager = appManager.windowManager
    this.translator = appManager.translator
  }

  AppTrayMenu() {
    const $t = this.translator.get()

    // Menu template
    const template = [
      {
        label: $t('trayMenu.reset'),
        click: () => {
          if (!this.windowManager.mainWindow.win) {
            this.windowManager.mainWindow.createWindow()
          }

          // Execute electron window method
          this.windowManager.mainWindow.win.restore()
          this.windowManager.mainWindow.win.moveTop()
        }
      },
      {
        label: $t('trayMenu.settings'),
        click: () => {}
      },
      {
        type: 'separator'
      },
      {
        label: $t('trayMenu.language.title'),
        enabled: false
      },
      {
        label: $t('trayMenu.language.zh'),
        type: 'checkbox',
        checked: process.env.VUE_APP_DEFAULT_LANGUAGE === 'zh'
      },
      {
        label: $t('trayMenu.language.en'),
        type: 'checkbox',
        checked: process.env.VUE_APP_DEFAULT_LANGUAGE === 'en'
      },
      {
        type: 'separator'
      },
      {
        label: $t('trayMenu.exit'),
        click: () => {
          app.exit()
        }
      }
    ]
    return Menu.buildFromTemplate(template)
  }
}

export default MenuManager
