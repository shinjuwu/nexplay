import type { VueWrapper } from '@vue/test-utils'
import { beforeEach, describe, expect, test } from 'vitest'
import { mount } from '@vue/test-utils'

import HomeGroupMenu from '@/components/HomeGroupMenu.vue'

describe('home group menu test', () => {
  let wrapper: VueWrapper

  beforeEach(() => {
    wrapper = mount(HomeGroupMenu, {
      slots: {
        default: 'Home Group Menu Default Slot',
      },
    })
  })

  test('default slot render correctly', () => {
    expect(wrapper.find('section').find('nav').text()).toBe('Home Group Menu Default Slot')
  })

  test('section header h2 text is binding with props groupName property', async () => {
    let groupName = 'test group name'
    await wrapper.setProps({ groupName })
    expect(wrapper.find('section').find('header').find('h2').text()).toBe(groupName)

    groupName = 'test group name2'
    await wrapper.setProps({ groupName })
    expect(wrapper.find('section').find('header').find('h2').text()).toBe(groupName)
  })
})
