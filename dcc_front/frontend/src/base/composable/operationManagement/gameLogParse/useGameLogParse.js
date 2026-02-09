import { computed, defineAsyncComponent, inject, onErrorCaptured, onMounted, reactive, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import * as api from '@/base/api/sysRecord'
import constant from '@/base/common/constant'
import { roomTypeNameIndex } from '@/base/utils/room'

const BaccaratGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Baccarat/BaccaratGameLog.vue')
)
const FantanGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Fantan/FantanGameLog.vue')
)
const ColordiscGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Colordisc/ColordiscGameLog.vue')
)
const PrawncrabGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Prawncrab/PrawncrabGameLog.vue')
)
const HundredsicboGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Hundredsicbo/HundredsicboGameLog.vue')
)
const CockfightGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Cockfight/CockfightGameLog.vue')
)
const DogracingGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Dogracing/DogracingGameLog.vue')
)
const RocketGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Rocket/RocketGameLog.vue')
)
const AndarbaharGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Andarbahar/AndarbaharGameLog.vue')
)
const RouletteGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Roulette/RouletteGameLog.vue')
)
const BlackjackGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Blackjack/BlackjackGameLog.vue')
)
const SangongGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Sangong/SangongGameLog.vue')
)
const BullbullGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Bullbull/BullbullGameLog.vue')
)
const TexasGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Texas/TexasGameLog.vue')
)
const RummyGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Rummy/RummyGameLog.vue')
)
const GoldenflowerGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Goldenflower/GoldenflowerGameLog.vue')
)
const PokdengGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Pokdeng/PokdengGameLog.vue')
)
const CatteGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Catte/CatteGameLog.vue')
)
const ChinesepokerGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Chinesepoker/ChinesepokerGameLog.vue')
)
const OkeyGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Okey/OkeyGameLog.vue')
)
const TeenpattiGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Teenpatti/TeenpattiGameLog.vue')
)
const FruitslotGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Fruitslot/FruitslotGameLog.vue')
)
const RcfishingGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Rcfishing/RcfishingGameLog.vue')
)
const PlinkoGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Plinko/PlinkoGameLog.vue')
)
const HappyfishingGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Happyfishing/HappyfishingGameLog.vue')
)
const Fruit777slotGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Fruit777slot/Fruit777slotGameLog.vue')
)
const MegsharkslotGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Megsharkslot/MegsharkslotGameLog.vue')
)
const MidasslotGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Midasslot/MidasslotGameLog.vue')
)
const WildgemslotGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Wildgemslot/WildgemslotGameLog.vue')
)
const JumphighslotGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Jumphighslot/JumphighslotGameLog.vue')
)
const PyrtreasureslotGameLog = defineAsyncComponent(() =>
  import('@/base/views/OperationManagement/GameLogParse/Pyrtreasureslot/PyrtreasureslotGameLog.vue')
)

export function useGameLogParse(props) {
  const { t } = useI18n()

  const warn = inject('warn')

  const formInput = reactive({
    logNumber: props.logNumber,
    userName: props.userName,
    showProcessing: false,
  })
  const gameLog = reactive({
    userName: '',
    gameId: constant.Game.All,
    gameCode: '',
    gameName: computed(() => {
      let gameName = ''
      if (gameLog.gameId > constant.Game.All) {
        gameName += t(`game__${gameLog.gameId}`)
      }
      if (gameLog.gameCode !== '') {
        gameName += `( ${gameLog.gameCode} )`
      }
      return gameName
    }),
    roomType: constant.RoomType.All,
    roomTypeName: computed(() =>
      gameLog.roomType > constant.RoomType.All
        ? t(`roomType__${roomTypeNameIndex(gameLog.gameId, gameLog.roomType)}`)
        : ''
    ),
    logNumber: '',
    playLog: null,
    betTime: null,
  })

  const playLogComponent = computed(() => {
    switch (gameLog.gameId) {
      case constant.Game.Baccarat:
        return BaccaratGameLog
      case constant.Game.Fantan:
        return FantanGameLog
      case constant.Game.Colordisc:
        return ColordiscGameLog
      case constant.Game.Prawncrab:
        return PrawncrabGameLog
      case constant.Game.Hundredsicbo:
        return HundredsicboGameLog
      case constant.Game.Cockfight:
        return CockfightGameLog
      case constant.Game.Dogracing:
        return DogracingGameLog
      case constant.Game.Rocket:
        return RocketGameLog
      case constant.Game.Andarbahar:
        return AndarbaharGameLog
      case constant.Game.Roulette:
        return RouletteGameLog
      case constant.Game.Blackjack:
        return BlackjackGameLog
      case constant.Game.Sangong:
        return SangongGameLog
      case constant.Game.Bullbull:
        return BullbullGameLog
      case constant.Game.Texas:
      case constant.Game.Friendstexas:
        return TexasGameLog
      case constant.Game.Rummy:
        return RummyGameLog
      case constant.Game.Goldenflower:
        return GoldenflowerGameLog
      case constant.Game.Pokdeng:
        return PokdengGameLog
      case constant.Game.Catte:
        return CatteGameLog
      case constant.Game.Chinesepoker:
        return ChinesepokerGameLog
      case constant.Game.Okey:
        return OkeyGameLog
      case constant.Game.Teenpatti:
        return TeenpattiGameLog
      case constant.Game.Fruitslot:
        return FruitslotGameLog
      case constant.Game.Rcfishing:
        return RcfishingGameLog
      case constant.Game.Plinko:
        return PlinkoGameLog
      case constant.Game.Happyfishing:
        return HappyfishingGameLog
      case constant.Game.Fruit777slot:
        return Fruit777slotGameLog
      case constant.Game.Megsharkslot:
        return MegsharkslotGameLog
      case constant.Game.Midasslot:
        return MidasslotGameLog
      case constant.Game.Wildgemslot:
        return WildgemslotGameLog
      case constant.Game.Jumphighslot:
        return JumphighslotGameLog
      case constant.Game.Pyrtreasureslot:
        return PyrtreasureslotGameLog
      default:
        return null
    }
  })

  async function search() {
    const searchLogNumber = formInput.logNumber
    if (searchLogNumber === undefined || searchLogNumber === null || searchLogNumber === '') {
      warn(t('textRoundIdRequired'))
      return
    }

    formInput.showProcessing = true

    try {
      const resp = await api.getPlayLogCommon({
        lognumber: searchLogNumber,
      })

      if (resp.data.code !== constant.ErrorCode.Success) {
        warn(t(`errorCode__${resp.data.code}`))
        return
      }

      if (!resp.data.data) {
        resetGameLog(searchLogNumber)
        return
      }

      const data = resp.data.data
      gameLog.playLog = data.play_log
      gameLog.gameId = data.game_id
      gameLog.gameCode = data.game_code
      gameLog.roomType = data.room_type
      gameLog.logNumber = data.lognumber
      gameLog.betTime = data.bet_time
      gameLog.userName = formInput.userName
    } catch (err) {
      console.error(err)

      if (axios.isAxiosError(err)) {
        const errorCode = err.response.data
          ? err.response.data.code || constant.ErrorCode.ErrorLocal
          : constant.ErrorCode.ErrorLocal
        warn(t(`errorCode__${errorCode}`))
      }
    } finally {
      formInput.showProcessing = false
    }
  }

  function resetGameLog(searchLogNumber) {
    gameLog.gameId = constant.Game.All
    gameLog.gameCode = ''
    gameLog.roomType = constant.RoomType.All
    gameLog.logNumber = searchLogNumber
    gameLog.userName = ''
    gameLog.playLog = null
    gameLog.betTime = null
  }

  onErrorCaptured((err) => {
    console.error(err)
    warn(t('textGameLogParseError'))
    return false
  })

  onMounted(async () => {
    if (formInput.logNumber !== '' && formInput.userName !== '') {
      await search()
    }
  })

  watch(props, async () => {
    formInput.logNumber = props.logNumber
    formInput.userName = props.userName

    if (formInput.logNumber !== '' && formInput.userName !== '') {
      await search()
    }
  })

  return {
    formInput,
    gameLog,
    playLogComponent,
    search,
  }
}
