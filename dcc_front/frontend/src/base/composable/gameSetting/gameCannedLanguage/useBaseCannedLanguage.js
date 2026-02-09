import { isEqual } from 'lodash'
import * as api from '@/base/api/sysGame'

export function useBaseCannedLanguage() {
  async function getCannedList(isDefault, agentId) {
    return await api.getCannedList({
      is_default: isDefault,
      agent_id: agentId,
    })
  }

  function getUpdateList(backupList, newList) {
    return newList.filter(
      (newItem) =>
        !isEqual(
          newItem,
          backupList.find((n) => n.serial === newItem.serial)
        )
    )
  }

  async function setCannedList(isDefault, agentId, updateList) {
    return await api.setCannedList({
      is_default: isDefault,
      agent_id: agentId,
      list: updateList.map((d) => ({
        serial: d.serial,
        canned_type: d.cannedType,
        emotion_type: d.emotionType,
        content_list: Array.from(d.langContentMap).map(([lang, content]) => {
          return {
            lang,
            content,
          }
        }),
        status: d.status,
      })),
    })
  }

  return {
    getCannedList,
    getUpdateList,
    setCannedList,
  }
}
