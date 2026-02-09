import storage from '@/utils/storage'

export function get() {
  return storage.local.get(storage.keys.ACCESS_TOKEN)
}

export function set(token: string, expiresAt: number) {
  const expiresTime = new Date().getTime() + expiresAt * 1000
  storage.local.set(storage.keys.ACCESS_TOKEN, token, expiresTime)
  storage.local.set(storage.keys.EXPIRES_TIME, expiresTime.toString())
}

export function update(token: string, expiresAt: number) {
  const expiresTime = new Date().getTime() + expiresAt * 1000
  storage.local.set(storage.keys.ACCESS_TOKEN, token, expiresTime)
  storage.local.set(storage.keys.EXPIRES_TIME, expiresTime.toString())
}

export function clear() {
  storage.local.remove(storage.keys.ACCESS_TOKEN)
  storage.local.remove(storage.keys.EXPIRES_TIME)
}
