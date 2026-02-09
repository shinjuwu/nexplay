import type { VueWrapper } from '@vue/test-utils'
import { beforeEach, describe, expect, test, vi } from 'vitest'
import { mount } from '@vue/test-utils'

import LayoutTopBar from '@/components/LayoutTopBar.vue'

const mocks = vi.hoisted(() => {
  return {
    routePath: '',
    routerPush: vi.fn(),
    userStoreSignOut: vi.fn(),
  }
})

vi.mock('vue-router', () => {
  return {
    useRoute: () => {
      return {
        path: mocks.routePath,
      }
    },
    useRouter: () => {
      return {
        push: mocks.routerPush,
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

describe('layout top bar test', () => {
  let wrapper: VueWrapper

  const options = {
    global: {
      stubs: ['FontAwesomeIcon'],
    },
  }

  beforeEach(() => {
    mocks.routerPush.mockReset()
    mocks.userStoreSignOut.mockReset()

    wrapper = mount(LayoutTopBar, options)
  })

  test('root nav class translate-x-[-250px] is binding with props isShowRightAside property', async () => {
    const rootNav = wrapper.find('nav')
    await wrapper.setProps({ isShowRightAside: true })
    expect(rootNav.classes('translate-x-[-250px]')).toBe(true)

    await wrapper.setProps({ isShowRightAside: false })
    expect(rootNav.classes('translate-x-[-250px]')).toBe(false)
  })

  test('isHomePage computed property is render correctly', async () => {
    // notice: route path 必須在 mount 之前進行變更, isHomePage computed 值才會正確
    mocks.routePath = '/'
    wrapper = mount(LayoutTopBar, options)
    expect(wrapper.find('nav').find('ul').find('li').find('a').find('span').exists()).toBeTruthy()

    mocks.routePath = '/fake_path'
    wrapper = mount(LayoutTopBar, options)
    expect(wrapper.find('nav').find('ul').find('li').find('a').find('span').exists()).toBeFalsy()

    await wrapper.find('nav').find('ul').find('li').find('a').trigger('click')
    expect(mocks.routerPush).toBeCalledWith('/')
  })

  test('fontawsome bars container click event is binding with emit toggleRightAside', async () => {
    const fontawsomeBarsContainer = wrapper.find('nav').find('ul:last-child').find('li:last-child').find('a')
    await fontawsomeBarsContainer.trigger('click')

    let toggleRightAsideEmit = wrapper.emitted('toggleRightAside')
    expect(toggleRightAsideEmit).toBeTruthy()
    expect(toggleRightAsideEmit?.length).toBe(1)

    await fontawsomeBarsContainer.trigger('click')
    toggleRightAsideEmit = wrapper.emitted('toggleRightAside')
    expect(toggleRightAsideEmit?.length).toBe(2)
  })
})
