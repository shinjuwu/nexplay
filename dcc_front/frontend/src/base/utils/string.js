export function ellipsisString(str, max = 20) {
  return str.length > max ? str.slice(0, max) + '...' : str
}
