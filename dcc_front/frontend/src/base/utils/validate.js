export function numberRangeValidate(val, min, max) {
  return !isNaN(val) && val >= min && val <= max
}

export function numberMinValidate(val, min) {
  return !isNaN(val) && val >= min
}
