export function strToNumber(value) {
  if (value) return parseFloat(value.replaceAll(',', ''))
}

export function numberToStr(number, fractionDigits) {
  if (!fractionDigits || isNaN(fractionDigits) || fractionDigits < 0) {
    fractionDigits = 2
  }

  const value = Math.pow(10, fractionDigits)

  return (parseInt((number * value).toFixed(2), 10) / value)
    .toFixed(fractionDigits)
    .replace(/\B(?<!\.\d*)(?=(\d{3})+(?!\d))/g, ',')
}
