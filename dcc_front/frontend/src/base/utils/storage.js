const keys = {
  // auth
  ACCESS_TOKEN: 'ACCESS_TOKEN',
  EXPIRES_TIME: 'EXPIRES_TIME',
  // user
  LOCALE: 'LOCALE',
  INFO: 'INFO',
}

class Storage {
  constructor(storage) {
    this.support = null
    this.storage = storage
  }

  _getSupport() {
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

  get(key) {
    if (!this._getSupport()) {
      return null
    }
    let val = this.storage[key]
    if (!val) {
      return null
    }
    val = _tryParseJson(val)
    if (typeof val === 'string') {
      return val
    }
    // value有可能是boolean，需特別判斷
    if (val.expire && (typeof val.value === 'boolean' || val.value)) {
      const expireDate = new Date(parseInt(val.expire))
      const valid = expireDate instanceof Date && !isNaN(expireDate)
      if (valid && new Date() > expireDate) {
        return null
      }
      return val.value
    }
    return val
  }

  set(key, val, expire) {
    if (!this._getSupport()) {
      return
    }
    if (expire) {
      val = { value: val, expire: expire }
    }
    if (typeof val === 'object') {
      val = JSON.stringify(val)
    }
    this.storage[key] = val
  }

  remove(key) {
    this.storage.removeItem(key)
  }
}

function _tryParseJson(str) {
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
