import { consola } from 'consola'

const instance = consola.create({
  formatOptions: {
    colors: true,
    columns: 4,
    compact: false,
    date: false
  }
})

export const logInfo = instance.info
export const logWarn = instance.warn
export const logError = instance.error
export const logSuccess = instance.success
