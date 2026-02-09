import type { Component } from 'vue'
import type { VueWrapper } from '@vue/test-utils'
import { beforeEach, describe, expect, test, vi } from 'vitest'
import { mount } from '@vue/test-utils'

import HomeGroupMenuItem from '@/components/HomeGroupMenuItem.vue'

const mocks = vi.hoisted(() => {
  return {
    navigate: vi.fn(),
  }
})

const mockRouterLinkComponent: Component = {
  template: '<slot :navigate="navigate"></slot>',
  props: {
    custom: {
      type: Boolean,
    },
    to: {
      type: String,
    },
  },
  setup(props) {
    return {
      navigate: () => {
        mocks.navigate(props.to)
      },
    }
  },
}

describe('home group menu item test', () => {
  let wrapper: VueWrapper

  beforeEach(async () => {
    mocks.navigate.mockReset()

    wrapper = mount(HomeGroupMenuItem, {
      global: {
        components: {
          'router-link': mockRouterLinkComponent,
        },
      },
      slots: {
        default: 'Home Group Menu Item Default Slot',
      },
    })
  })

  test('default slot render correctly', () => {
    expect(wrapper.find('div').find('div').find('div:last-child').text()).toBe('Home Group Menu Item Default Slot')
  })

  test('default slot container class is binding with props iconContainerClass property', async () => {
    let iconContainerClass = 'testIconContainerClass1'
    await wrapper.setProps({ iconContainerClass })
    expect(wrapper.find('div').find('div').find('div:last-child').classes()).toContain(iconContainerClass)

    iconContainerClass = 'testIconContainerClass2'
    await wrapper.setProps({ iconContainerClass })
    expect(wrapper.find('div').find('div').find('div:last-child').classes()).toContain(iconContainerClass)
  })

  test('last div text is binding with props name property', async () => {
    let name = 'test name 1'
    await wrapper.setProps({ name })
    let divs = wrapper.findAll('div')
    expect(divs[divs.length - 1].text()).toBe(name)

    name = 'test name 2'
    await wrapper.setProps({ name })
    divs = wrapper.findAll('div')
    expect(divs[divs.length - 1].text()).toBe(name)
  })

  test('navigate path is binding with props linkTo property', async () => {
    let linkTo = '/test1'
    await wrapper.setProps({ linkTo })
    await wrapper.find('div').trigger('click')
    expect(mocks.navigate).toBeCalled()
    expect(mocks.navigate).toBeCalledTimes(1)
    expect(mocks.navigate).toBeCalledWith(linkTo)

    linkTo = '/test2'
    await wrapper.setProps({ linkTo })
    await wrapper.find('div').trigger('click')
    expect(mocks.navigate).toBeCalledTimes(2)
    expect(mocks.navigate).toBeCalledWith(linkTo)
  })
})
