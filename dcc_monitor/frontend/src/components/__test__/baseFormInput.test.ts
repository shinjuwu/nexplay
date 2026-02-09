import type { VueWrapper } from '@vue/test-utils'
import { beforeEach, describe, expect, test } from 'vitest'
import { mount } from '@vue/test-utils'

import BaseFormInput from '@/components/BaseFormInput.vue'

describe('base form input test', () => {
  let wrapper: VueWrapper

  beforeEach(() => {
    wrapper = mount(BaseFormInput, {
      attachTo: document.body,
      slots: {
        icon: 'icon slot',
      },
    })
  })

  test('scopeRef is binding', () => {
    expect(wrapper.vm.$refs.scopeRef).not.toBeUndefined()
  })

  test('inputRef is binding', () => {
    expect(wrapper.vm.$refs.inputRef).not.toBeUndefined()
  })

  test('input is focused by clicking none input part', async () => {
    // notice: 使用vitest觸發的事件不會冒泡，所以確保點擊任何區域都必須觸發的事件，綁在最上層進行測試即可
    await wrapper.find('div').trigger('click')
    expect(document.activeElement).toBe(wrapper.vm.$refs.inputRef)
  })

  test('input type is binding with props type property', async () => {
    const input = wrapper.find('input')

    await wrapper.setProps({ type: 'password' })
    expect(input.attributes('type')).toBe('password')

    await wrapper.setProps({ type: 'text' })
    expect(input.attributes('type')).toBe('text')
  })

  test('input value is binding with props modelValue property', async () => {
    const input = wrapper.find('input')

    await wrapper.setProps({ modelValue: '123456' })
    expect(input.element.value).toBe('123456')

    await wrapper.setProps({ modelValue: '' })
    expect(input.element.value).toBe('')
  })

  test('input update value is binding with emit update:modelValue', async () => {
    const input = wrapper.find('input')

    input.setValue('123456')
    const emit = wrapper.emitted('update:modelValue')
    expect(emit).toBeTruthy()
    expect(emit?.length).toBe(1)
  })

  test('input maxlength is binding with props maxlength property', async () => {
    const input = wrapper.find('input')

    await wrapper.setProps({ maxlength: 0 })
    expect(input.attributes('maxlength')).not.toBeDefined()

    await wrapper.setProps({ maxlength: 10 })
    expect(input.attributes('maxlength')).toBe('10')
  })

  test('props placeholder is binding', async () => {
    await wrapper.setProps({ placeholder: '123456' })

    let placeholderContainer = wrapper.find('div').find('div').findAll('div')[1]
    expect(placeholderContainer.exists()).toBeTruthy()
    expect(placeholderContainer.text()).toBe('123456')

    await wrapper.setProps({ placeholder: '654321' })
    placeholderContainer = wrapper.find('div').find('div').findAll('div')[1]
    expect(placeholderContainer.text()).toBe('654321')
  })

  test('error message render correctly with props errorMessage', async () => {
    await wrapper.setProps({ errorMessage: '' })
    let errorMessageContainer = wrapper.find('div').findAll('div')[2]
    expect(errorMessageContainer.element.style.display).toBe('none')

    await wrapper.setProps({ errorMessage: '654321' })
    errorMessageContainer = wrapper.find('div').findAll('div')[2]
    expect(errorMessageContainer.element.style.display).toBe('')
  })

  test('icon slot render correctly', async () => {
    await wrapper.setProps({ rightIcon: true })
    let iconContainer = wrapper.find('div').find('div').find('div').find('div')
    expect(iconContainer.exists()).toBeTruthy()
    expect(iconContainer.text()).toBe('icon slot')

    await wrapper.setProps({ rightIcon: false })
    iconContainer = wrapper.find('div').find('div').find('div').find('div')
    expect(iconContainer.exists()).toBeFalsy()
  })

  test('icon slot click event is binding with emit clickIcon', async () => {
    await wrapper.setProps({ rightIcon: true })

    const iconContainer = wrapper.find('div').find('div').find('div').find('div')
    await iconContainer.trigger('click')

    let clickIconEmit = wrapper.emitted('clickIcon')
    expect(clickIconEmit).toBeTruthy()
    expect(clickIconEmit?.length).toBe(1)

    await iconContainer.trigger('click')
    clickIconEmit = wrapper.emitted('clickIcon')
    expect(clickIconEmit?.length).toBe(2)
  })
})
