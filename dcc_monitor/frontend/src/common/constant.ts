export const responseCode = {
  Success: 0,
  Fail: 1,
  Unauth: 2, // http status 401
  Exception: 3, // http status 404
  Local: 4, // http status 500
  Jwt: 5,
}

export const errorMessages = {
  required: (fieldName: string) => `${fieldName} 栏位必须输入`,
  maxLength: (fieldName: string, maxLength: number) => `${fieldName} 最大长度为 ${maxLength} 字符`,
  minLength: (fieldName: string, minLength: number) => `${fieldName} 最小长度为 ${minLength} 字符`,
  stringLength: (fieldName: string, length: number) => `${fieldName} 长度必须为 ${length} 字符`,
  lowercaseEnglishAndNumber4To16: (fieldName: string) => `${fieldName} 格式为 4~16 小写及英文数字字符`,
  englishAndNumber8To16: (fieldName: string) => `${fieldName} 格式为 8~16 大小写英文及数字字符`,
  expired: (fieldName: string) => `${fieldName} 已过期`,
}
