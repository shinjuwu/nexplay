import storage from '@/base/utils/storage'

export function get() {
  return storage.local.get(storage.keys.ACCESS_TOKEN)
}

export function set(result) {
  const expiresTime = new Date().getTime() + result.expires_in * 1000
  storage.local.set(storage.keys.ACCESS_TOKEN, result.access_token, expiresTime)
  storage.local.set(storage.keys.EXPIRES_TIME, expiresTime.toString())
}

export function update(result) {
  const expiresTime = new Date().getTime() + result.expires_in * 1000
  storage.local.set(storage.keys.ACCESS_TOKEN, result.access_token, expiresTime)
  storage.local.set(storage.keys.EXPIRES_TIME, expiresTime.toString())
}

export function clear() {
  storage.local.remove(storage.keys.ACCESS_TOKEN)
  storage.local.remove(storage.keys.EXPIRES_TIME)
}
