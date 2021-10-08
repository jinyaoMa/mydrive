import { ipcMain, dialog } from 'electron'

class IpcEvents {
  create(appManager) {
    this.appManager = appManager

    /* 翻译器函数
    Translator function */
    const translator = appManager.translator
    const $t = translator.get()

    // ipc通信示例 / ipc demo
    ipcMain.on('showDialog', (event, msg) => {
      dialog.showMessageBox({
        type: 'info',
        title: '收到消息！',

        // 在任何能调用翻译器函数的地方都能使用多语言
        // Multi-language support where translator functions are available
        message: $t('reciveFromRenderer'),
        detail: msg
      })
    })

    // renderer check for current locale
    ipcMain.on('getCurrentLocale', (event) => {
      event.reply('getCurrentLocale-reply', translator.locale)
    })

    // 语言变更事件 / language change event
    ipcMain.on('appLanguageChange', (event, lang) => {
      this.appManager.languageChange(lang)
    })
  }
}

export default new IpcEvents()
