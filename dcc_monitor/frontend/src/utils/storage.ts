const keys = {
  // auth
  ACCESS_TOKEN: 'ACCESS_TOKEN',
  EXPIRES_TIME: 'EXPIRES_TIME',
  // user
  LOCALE: 'LOCALE',
  INFO: 'INFO',
}

class Storage {
  support: boolean | null
  storage: globalThis.Storage

  constructor(storage: globalThis.Storage) {
    this.support = null
    this.storage = storage
  }

  private getSupport() {
    if (this.support === null) {
      const test = 'test'
      try {
        this.storage.setItem(test, test)
        this.storage.removeItem(test)
        this.support = true
      } catch (e) {
        this.support = false
      }
    }
    return this.support
  }

  get(key: string) {
    if (!this.getSupport()) {
      return null
    }

    const val = this.storage.getItem(key)
    if (val === null) {
      return null
    }

    const parseVal = tryParseJson(val)
    if (typeof parseVal === 'string') {
      return parseVal
    }

    if (typeof parseVal === 'object') {
      // value有可能是boolean，需特別判斷
      if (parseVal.expire && (typeof parseVal.value === 'boolean' || parseVal.value)) {
        const expireDate = new Date(parseVal.expire)
        if (new Date() > expireDate) {
          return null
        }
        return parseVal.value
      }
    }
    return parseVal
  }

  set(key: string, val: object | string, expire?: number) {
    if (!this.getSupport()) {
      return
    }
    if (expire) {
      val = { value: val, expire: expire }
    }
    if (typeof val === 'object') {
      val = JSON.stringify(val)
    }
    this.storage.setItem(key, val)
  }

  remove(key: string) {
    this.storage.removeItem(key)
  }
}

function tryParseJson(str: string) {
  try {
    return JSON.parse(str)
  } catch (e) {
    return str
  }
}

export default {
  keys: keys,
  local: new Storage(localStorage),
  session: new Storage(sessionStorage),
}
