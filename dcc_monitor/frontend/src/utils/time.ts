import { format } from 'date-fns'

class Time {
  format(date: Date, layout = 'yyyy-MM-dd HH:mm:ss') {
    return format(date, layout)
  }

  now(): Date {
    return new Date()
  }
}

export default new Time()
