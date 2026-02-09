export function transpose2D(arr) {
  const row = arr.length
  const column = row > 0 ? arr[0].length : 0

  const newArr = Array.from({ length: column }, () => Array.from({ length: row }))
  for (let i = 0; i < row; i++) {
    for (let j = 0; j < column; j++) {
      newArr[j][i] = arr[i][j]
    }
  }

  return newArr
}

function baseEqualCheck(arr1, arr2) {
  return Array.isArray(arr1) && Array.isArray(arr2) && arr1.length === arr2.length
}

export function equal1D(arr1, arr2) {
  return baseEqualCheck(arr1, arr2) && arr1.every((a1, idx) => a1 === arr2[idx])
}

export function equal2D(arr1, arr2) {
  return Array.isArray(arr1) && arr1.every((subArr1, subIdx) => equal1D(subArr1, arr2[subIdx]))
}
