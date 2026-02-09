import { errorMessages } from '@/common/constant'

interface ValidateResult {
  isValid: boolean
  errorMessage: string
}

interface ValidateFn<T> {
  (source: T, fieldName: string): ValidateResult
}

export function validateRules<T>(...fns: ValidateFn<T>[]): ValidateFn<T> {
  return function (source: T, fieldName: string): ValidateResult {
    return fns.reduce(
      (vr, fn) => {
        if (!vr.isValid) {
          return vr
        }
        return fn(source, fieldName)
      },
      { isValid: true, errorMessage: '' } as ValidateResult
    )
  }
}

export function required(): ValidateFn<string> {
  return function (source: string, fieldName: string): ValidateResult {
    const isValid = source !== ''
    return {
      isValid,
      errorMessage: isValid ? '' : errorMessages.required(fieldName),
    }
  }
}

export function maxLength(maxLength: number): ValidateFn<string> {
  return function (source: string, fieldName: string): ValidateResult {
    const isValid = source.length <= maxLength
    return {
      isValid,
      errorMessage: isValid ? '' : errorMessages.maxLength(fieldName, maxLength),
    }
  }
}

export function minLength(minLength: number): ValidateFn<string> {
  return function (source: string, fieldName: string): ValidateResult {
    const isValid = source.length >= minLength
    return {
      isValid,
      errorMessage: isValid ? '' : errorMessages.minLength(fieldName, minLength),
    }
  }
}

export function stringLength(length: number): ValidateFn<string> {
  return function (source: string, fieldName: string): ValidateResult {
    const isValid = source.length === length
    return {
      isValid,
      errorMessage: isValid ? '' : errorMessages.stringLength(fieldName, length),
    }
  }
}

export function lowercaseEnglishAndNumber4To16(): ValidateFn<string> {
  return function (source: string, fieldName: string): ValidateResult {
    const isValid = /^[a-z0-9]{4,16}$/.test(source)
    return {
      isValid,
      errorMessage: isValid ? '' : errorMessages.lowercaseEnglishAndNumber4To16(fieldName),
    }
  }
}

export function englishAndNumber8To16(): ValidateFn<string> {
  return function (source: string, fieldName: string): ValidateResult {
    const isValid = /^[A-Za-z0-9]{8,16}$/.test(source)
    return {
      isValid,
      errorMessage: isValid ? '' : errorMessages.englishAndNumber8To16(fieldName),
    }
  }
}

export function expired(): ValidateFn<number | Date> {
  return function (source: number | Date, fieldName: string): ValidateResult {
    const isValid = new Date(source).getTime() >= new Date().getTime()
    return {
      isValid,
      errorMessage: isValid ? '' : errorMessages.expired(fieldName),
    }
  }
}
