<template>
  <ToggleHeader>
    <template #default="{ show: showSection }">
      <form v-show="showSection">
        <div class="flex flex-wrap items-center space-x-3">
          <label class="form-label ml-1">
            <input v-model="gameIconList.isDefault" type="radio" :value="false" :disabled="isAdmin || !isEditEnabled" />
            {{ t('textUseCustomSorting') }}
          </label>
          <label v-if="!isAdmin" class="form-label ml-1">
            <input v-model="gameIconList.isDefault" type="radio" :value="true" :disabled="!isEditEnabled" />
            {{ t('textUseSuperiorSorting') }}
          </label>
        </div>
        <div class="mt-4 flex flex-wrap">
          <div class="mb-1">
            <div v-for="(reminder, idx) in reminders" :key="`reminder__${idx}`" class="text-danger">
              {{ reminder }}
            </div>
          </div>
          <div class="flex flex-1 items-end justify-end">
            <LoadingButton
              v-show="isEditEnabled"
              :is-get-data="!tableInput.showProcessing"
              :parent-data="gameIconList"
              :button-click="() => setGameIconList()"
              class="btn btn-primary flex w-full md:w-24"
            >
              {{ t('textOnSave') }}
            </LoadingButton>
          </div>
        </div>
      </form>

      <hr class="my-2" />

      <ToggleHeader>
        <template #default="{ show: showNormals }">
          <div v-show="showNormals" class="tbl-container w-full overflow-visible">
            <table class="tbl tbl-hover">
              <thead>
                <tr>
                  <th>{{ t('textSortingNumber') }}</th>
                  <th>{{ t('textGameId') }}</th>
                  <th>{{ t('textGameCode') }}</th>
                  <th>{{ t('textGameName') }}</th>
                  <th>{{ t('textPromotePicture') }}</th>
                  <th>{{ t('textLabel') }}</th>
                </tr>
              </thead>
              <tbody>
                <template v-if="normalGameIconList.length === 0">
                  <tr class="no-hover">
                    <td colspan="6">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr
                    v-for="gameIcon in normalGameIconList"
                    :key="`gameIcon__${gameIcon.gameId}`"
                    :class="{ 'hover:shadow-md': isEditEnabled }"
                  >
                    <td>
                      <input v-model="gameIcon.rank" class="form-input text-center" type="number" min="0" max="9999" />
                    </td>
                    <td>{{ gameIcon.gameId }}</td>
                    <td>{{ gameIcon.gameCode }}</td>
                    <td>{{ t(`game__${gameIcon.gameId}`) }}</td>
                    <td>
                      <template v-if="isEditEnabled">
                        <div class="space-x-3">
                          <div v-for="push in 2" :key="`gameIconPush__${push}`" class="relative inline-flex">
                            <button
                              class="btn peer"
                              :class="gameIcon.push === push ? 'btn-success' : 'btn-secondary'"
                              @click="gameIcon.push = gameIcon.push === push ? 0 : push"
                            >
                              {{ t(`gameIconPush__${push}`) }}
                            </button>
                            <img
                              class="absolute bottom-11 hidden rounded border bg-gray-300 peer-hover:block"
                              :src="getImageUrl(gameIcon.gameId, push)"
                            />
                            <div
                              class="absolute -top-0.5 left-1/2 hidden h-0 w-0 -translate-x-1/2 -translate-y-1/2 border-4 border-b-0 border-gray-300 border-x-transparent peer-hover:block"
                            ></div>
                          </div>
                        </div>
                      </template>
                      <template v-else>
                        <span v-if="gameIcon.push > 0" class="text-success">
                          {{ t(`gameIconPush__${gameIcon.push}`) }}
                        </span>
                      </template>
                    </td>
                    <td :class="{ 'space-x-3': isEditEnabled }">
                      <template v-if="isEditEnabled">
                        <button
                          class="btn"
                          :class="gameIcon.isHot === 1 ? 'btn-danger' : 'btn-secondary'"
                          @click="changeIconLabel(gameIcon, 'isHot')"
                        >
                          {{ t('textIsHot') }}
                        </button>
                        <button
                          class="btn"
                          :class="gameIcon.isNewest === 1 ? 'btn-warning' : 'btn-secondary'"
                          @click="changeIconLabel(gameIcon, 'isNewest')"
                        >
                          {{ t('textIsNewest') }}
                        </button>
                      </template>
                      <template v-else>
                        <span v-if="gameIcon.isHot === 1" :class="{ 'text-danger': gameIcon.isHot === 1 }">
                          {{ t('textIsHot') }}
                        </span>
                        <span v-if="gameIcon.isNewest === 1" :class="{ 'text-warning': gameIcon.isNewest === 1 }">
                          {{ t('textIsNewest') }}
                        </span>
                      </template>
                    </td>
                  </tr>
                </template>
              </tbody>
            </table>
          </div>
        </template>
      </ToggleHeader>

      <hr class="my-2" />

      <ToggleHeader>
        <template #default="{ show: showFriendRooms }">
          <div v-show="showFriendRooms" class="tbl-container w-full overflow-visible">
            <table class="tbl tbl-hover">
              <thead>
                <tr>
                  <th>{{ t('textSortingNumber') }}</th>
                  <th>{{ t('textGameId') }}</th>
                  <th>{{ t('textGameCode') }}</th>
                  <th>{{ t('textGameName') }}</th>
                  <th>{{ t('textLabel') }}</th>
                </tr>
              </thead>
              <tbody>
                <template v-if="friendRoomGameIconList.length === 0">
                  <tr class="no-hover">
                    <td colspan="5">{{ t('textTableEmpty') }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr
                    v-for="gameIcon in friendRoomGameIconList"
                    :key="`gameIcon__${gameIcon.gameId}`"
                    :class="{ 'hover:shadow-md': isEditEnabled }"
                  >
                    <td>
                      <input v-model="gameIcon.rank" class="form-input text-center" type="number" min="0" max="9999" />
                    </td>
                    <td>{{ gameIcon.gameId }}</td>
                    <td>{{ gameIcon.gameCode }}</td>
                    <td>{{ t(`game__${gameIcon.gameId}`) }}</td>
                    <td :class="{ 'space-x-3': isEditEnabled }">
                      <template v-if="isEditEnabled">
                        <button
                          class="btn"
                          :class="gameIcon.isHot === 1 ? 'btn-danger' : 'btn-secondary'"
                          @click="changeIconLabel(gameIcon, 'isHot')"
                        >
                          {{ t('textIsHot') }}
                        </button>
                        <button
                          class="btn"
                          :class="gameIcon.isNewest === 1 ? 'btn-warning' : 'btn-secondary'"
                          @click="changeIconLabel(gameIcon, 'isNewest')"
                        >
                          {{ t('textIsNewest') }}
                        </button>
                      </template>
                      <template v-else>
                        <span v-if="gameIcon.isHot === 1" :class="{ 'text-danger': gameIcon.isHot === 1 }">
                          {{ t('textIsHot') }}
                        </span>
                        <span v-if="gameIcon.isNewest === 1" :class="{ 'text-warning': gameIcon.isNewest === 1 }">
                          {{ t('textIsNewest') }}
                        </span>
                      </template>
                    </td>
                  </tr>
                </template>
              </tbody>
            </table>
          </div>
        </template>
      </ToggleHeader>
    </template>
  </ToggleHeader>
</template>

<script setup>
import { useGameSorting } from '@/base/composable/gameSetting/gameSorting/userGameSorting'
import { useI18n } from 'vue-i18n'
import ToggleHeader from '@/base/components/Page/ToggleHeader.vue'
import LoadingButton from '@/base/components/Button/LoadingButton.vue'

const { t } = useI18n()

const {
  reminders,
  friendRoomGameIconList,
  gameIconList,
  normalGameIconList,
  tableInput,
  isAdmin,
  isEditEnabled,
  changeIconLabel,
  getImageUrl,
  setGameIconList,
} = useGameSorting()
</script>
