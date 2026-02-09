import type { VueWrapper } from '@vue/test-utils'
import { beforeEach, describe, expect, test, vi } from 'vitest'
import { mount } from '@vue/test-utils'

import LayoutRightAside from '@/components/LayoutRightAside.vue'

const mocks = vi.hoisted(() => {
  return {
    routerPuth: vi.fn(),
    userStoreSignOut: vi.fn(),
  }
})

vi.mock('vue-router', () => {
  return {
    useRouter: () => {
      return {
        push: mocks.routerPuth,
      }
    },
  }
})

vi.mock('@/store/userStore', () => {
  return {
    useUserStore: () => {
      return {
        signOut: mocks.userStoreSignOut,
      }
    },
  }
})

describe('layout right aside test', () => {
  let wrapper: VueWrapper

  const options = {
    global: {
      stubs: ['FontAwesomeIcon', 'RouterLink'],
    },
  }

  beforeEach(() => {
    mocks.routerPuth.mockReset()
    mocks.userStoreSignOut.mockReset()

    wrapper = mount(LayoutRightAside, options)
  })

  test('root aside class is binding with props isShow property correctly', async () => {
    const rootAside = wrapper.find('aside')
    await wrapper.setProps({ isShow: true })
    expect(rootAside.classes('right-0')).toBeTruthy()

    await wrapper.setProps({ isShow: false })
    expect(rootAside.classes('right-[-250px]')).toBeTruthy()
  })

  test('fontawsome arrow-right container click event is binding with emit toggleIsShow', async () => {
    const fontawsomeArrowRightContainer = wrapper.find('aside').find('div').find('div')
    await fontawsomeArrowRightContainer.trigger('click')

    let toggleIsShowEmit = wrapper.emitted('toggleIsShow')
    expect(toggleIsShowEmit).toBeTruthy()
    expect(toggleIsShowEmit?.length).toBe(1)
    expect(mocks.routerPuth).toBeCalledTimes(0)

    await fontawsomeArrowRightContainer.trigger('click')
    toggleIsShowEmit = wrapper.emitted('toggleIsShow')
    expect(toggleIsShowEmit?.length).toBe(2)
    expect(mocks.routerPuth).toBeCalledTimes(0)
  })

  test('fontawsome house container click event is binding with function redirectToHome', async () => {
    const fontawsomeHouseContainer = wrapper.find('aside').find('div').find('div:last-child')
    await fontawsomeHouseContainer.trigger('click')

    expect(mocks.routerPuth).toBeCalled()
    expect(mocks.routerPuth).toBeCalledTimes(1)

    await fontawsomeHouseContainer.trigger('click')
    expect(mocks.routerPuth).toBeCalledTimes(2)
  })
})
