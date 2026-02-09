import { reactive, ref } from 'vue'
import { defineStore, storeToRefs } from 'pinia'
import axios from 'axios'
import * as api from '@/base/api/sysNotify'
import { useUserStore } from '@/base/store/userStore'
import time from '@/base/utils/time'

export const chatServerMessageSubject = {
  message: 'message',
  announcement: 'announcement',
}

export const useChatServiceStore = defineStore('chatService', () => {
  const { user } = storeToRefs(useUserStore())

  const chatServiceConnInfo = {
    apiKey: '',
    channel: '',
    domain: '',
    path: '',
    scheme: '',
    wsConnPath: '',
  }

  let chatServiceWS
  const chatServiceWSInfo = {
    platform: '',
    token: '',
    userId: '',
  }

  const unreadMessageCount = ref(0)
  const unreadMessages = reactive({
    [chatServerMessageSubject.message]: [],
    [chatServerMessageSubject.announcement]: [],
  })
  const readMessages = reactive({
    [chatServerMessageSubject.message]: [],
    [chatServerMessageSubject.announcement]: [],
  })

  async function init() {
    await getChatServiceConn()
    await getChatServiceWSToken()
  }

  /** 與後台取得chat service server 連接資訊 */
  async function getChatServiceConn() {
    try {
      const resp = await api.getChatServiceConnInfo()

      const data = resp.data.data
      chatServiceConnInfo.apiKey = data.api_key
      chatServiceConnInfo.channel = data.channel
      chatServiceConnInfo.domain = data.domain
      chatServiceConnInfo.path = data.path
      chatServiceConnInfo.scheme = data.scheme
      chatServiceConnInfo.wsConnPath = data.ws_conn_path
    } catch (err) {
      console.error(err)
      throw err
    }
  }

  /** 與chat service server取得 WS 連接資訊 */
  async function getChatServiceWSToken() {
    try {
      const url = `${chatServiceConnInfo.scheme}://${chatServiceConnInfo.domain}${chatServiceConnInfo.path}login?username=${user.value.name}&agent_id=${user.value.agentId}&platform=${chatServiceConnInfo.channel}`
      const authUserName = chatServiceConnInfo.apiKey

      const resp = await axios({
        method: 'get',
        url: url,
        auth: { username: authUserName },
      })

      const data = resp.data.data
      chatServiceWSInfo.platform = data.platform
      chatServiceWSInfo.token = data.token
      chatServiceWSInfo.userId = data.user_id
    } catch (err) {
      console.error(err)
      throw err
    }
  }

  /** 與chat service server進行ws連線 */
  function start() {
    if (chatServiceWS) {
      return
    }

    const scheme = window.location.protocol === 'http:' ? 'ws' : 'wss'

    chatServiceWS = new WebSocket(
      `${scheme}://${chatServiceConnInfo.domain}${chatServiceConnInfo.wsConnPath}?token=${chatServiceWSInfo.token}`
    )
    // 開啟連線
    chatServiceWS.onopen = () => {
      console.info('[ChatService]: ws connected.')
      sendGetUnreadMessages()
    }
    // 新訊息通知
    chatServiceWS.onmessage = (event) => {
      console.info(`[ChatService]: ws receive new message. msg: ${event.data}`)
      processMessage(event.data)
    }
    // 連線關閉時
    chatServiceWS.onclose = async () => {
      console.info('[ChatService]: ws closed.')
      await restart()
    }
    // 連線發生錯誤
    chatServiceWS.onerror = (event) => {
      console.info(`[ChatService]: ws occur error. err: ${JSON.stringify(event)}`)
    }
  }

  /** 與chat service server ws中斷連線 */
  function stop() {
    if (!chatServiceWS) {
      return
    }

    if (chatServiceWS.readyState === WebSocket.CLOSING || chatServiceWS.readyState === WebSocket.CLOSED) {
      chatServiceWS.close()
    }

    chatServiceWS = null
  }

  /** 與chat service server發送ws取得未讀取的資訊 */
  function sendGetUnreadMessages() {
    chatServiceWS.send(
      JSON.stringify({
        id: 'CWCT_UNREADCOUNT',
      })
    )
  }

  /** 與chat service serve重新進行ws連線 */
  async function restart() {
    // 每1秒重試一次
    do {
      stop()
      await init()
      start()
      await time.delay(1000)
    } while (!chatServiceWS || chatServiceWS.readyState !== WebSocket.OPEN)
  }

  /** 解析chat service server發送ws的新資訊 */
  function processMessage(data, createTime) {
    const msg = JSON.parse(data)

    switch (msg.id) {
      // 已讀全部訊息
      case 'cwct_readnotification':
        if (msg.code !== 0) {
          return
        }

        unreadMessageCount.value = 0

        Object.keys(chatServerMessageSubject).forEach((subject) => {
          readMessages[subject] = readMessages[subject].concat(unreadMessages[subject])
          unreadMessages[subject] = []
        })
        break
      // 未讀訊息
      case 'cwct_unreadcount': {
        const payload = JSON.parse(msg.payload)
        unreadMessageCount.value = payload.unread_count

        for (const item of payload.content_map) {
          processMessage(item.content, item.createTime)
        }
        break
      }
      // 廣播訊息
      case 'notifybroadcastmessage': {
        if (!createTime) {
          createTime = new Date().toISOString()
        }

        const subject = msg.subject

        unreadMessages[subject].push({
          payload: data.payload,
          createTime: createTime,
        })
        unreadMessageCount.value = unreadMessageCount.value + 1
        break
      }
    }
  }

  /** 與chat service server發送ws已讀所有的資訊 */
  function sendReadAllMessages() {
    chatServiceWS.send(
      JSON.stringify({
        id: 'CWCT_READNOTIFICATION',
      })
    )
  }

  /** 下面此區功能為預想先做起來的，目前沒有用到先將eslint關閉，未來有使用到eslint記得開啟檢查 */
  /* eslint-disable */

  /** 與chat service server發送ws已讀指定subject類型的所有資訊 */
  function sendReadAllSubjectMessages(subject) {
    // notice: 目前只有已讀全部訊息，沒有辦法針對單一類型訊息已讀，所以此處只能先用 readAllMessages，未來有實作再修改
    sendReadAllMessages()
  }

  /** 取得目前未讀的訊息數量 */
  function getUnreadMessageCount() {
    return unreadMessageCount.value
  }

  /** 取得所有未讀的訊息 */
  function getAllUnreadMessages() {
    return Object.keys(chatServerMessageSubject).reduce((messages, subject) => {
      return messages.concat(unreadMessages[subject])
    }, [])
  }

  /** 取得指定ubject所有未讀的訊息 */
  function getAllUnreadSubjectMessages(subject) {
    return unreadMessages[subject] || []
  }
  /* eslint-enable */

  /** 向後台發送特定的聊天訊息 */
  async function sendMessage(message) {
    if (!message) {
      return
    }

    const data = new FormData()
    data.append('message', message)

    try {
      return await api.notifyBroadcastMessage(data)
    } catch (err) {
      console.error(err)
      throw err
    }
  }

  return {
    readMessages,
    unreadMessageCount,
    unreadMessages,
    init,
    sendGetUnreadMessages,
    sendReadAllMessages,
    sendReadAllSubjectMessages,
    sendMessage,
    start,
    stop,
  }
})
