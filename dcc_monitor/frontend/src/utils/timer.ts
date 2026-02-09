import time from '@/utils/time'

class timer {
  private _id: number
  private _handler: () => void

  constructor(id: number, handler: () => void) {
    this._id = id
    this._handler = handler
  }

  get id(): number {
    return this._id
  }

  handler() {
    this._handler()
  }
}

class timerHandler {
  private list: timer[]
  private requestId: number | null
  private timeoutId: number | undefined
  private interval: number
  private support = window.requestAnimationFrame !== undefined
  private lastTimestamp: number

  constructor(options: { interval: number }) {
    this.list = []
    this.requestId = null
    this.timeoutId = undefined
    this.interval = options.interval
    this.lastTimestamp = 0
  }

  get(id: number) {
    return this.list.find((c) => c.id === id)
  }

  add(timer: timer) {
    this.list.push(timer)
    if (this.list.length == 1) {
      this.run()
    }
  }

  remove(id: number) {
    const index = this.list.findIndex((c) => c.id === id)
    if (index >= 0) {
      this.list.splice(index, 1)
      if (this.list.length <= 0) {
        this.stop()
      }
    }
  }

  removeAll() {
    this.list = []
    this.stop()
  }

  run() {
    if (this.support) {
      this.lastTimestamp = 0

      const step = (timestamp: number) => {
        if (this.lastTimestamp === 0) {
          this.lastTimestamp = timestamp
        }

        if (timestamp - this.lastTimestamp < this.interval) {
          this.requestId = requestAnimationFrame(step)
        } else {
          this.execute()
        }
      }

      this.requestId = requestAnimationFrame(step)
    } else {
      clearTimeout(this.timeoutId)

      if (this.lastTimestamp > 0 && time.now().getTime() - this.lastTimestamp >= this.interval) {
        this.execute()
      }

      this.timeoutId = window.setTimeout(() => {
        this.lastTimestamp = time.now().getTime()
        this.execute()
      }, this.interval)
    }
  }

  private execute() {
    for (let i = this.list.length - 1; i >= 0; i--) {
      this.list[i].handler()
    }
    if (this.list.length > 0) {
      this.run()
    } else {
      this.stop()
    }
  }

  stop() {
    if (this.support) {
      cancelAnimationFrame(this.requestId as number)
    } else {
      clearTimeout(this.timeoutId)
    }
  }
}

class timerController {
  private id: number
  private timerHandlers: Record<number, timerHandler>

  constructor() {
    this.id = 0
    this.timerHandlers = {}
  }

  register(interval: number, handler: () => void): number {
    if (this.timerHandlers[interval] === undefined) {
      this.timerHandlers[interval] = new timerHandler({ interval })
    }

    const t = new timer(this.id++, () => {
      handler()
    })
    this.timerHandlers[interval].add(t)

    return t.id
  }

  unregister(id: number) {
    Object.values(this.timerHandlers)
      .find((handler) => handler.get(id))
      ?.remove(id)
  }

  removeAll() {
    Object.values(this.timerHandlers).forEach((handler) => handler.removeAll())
  }
}

export default new timerController()
