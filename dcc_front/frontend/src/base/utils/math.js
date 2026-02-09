export function round(n, p) {
  return parseFloat(n.toFixed(p))
}

export function roundDown(n) {
  return Math.floor(n * 100) / 100
}
