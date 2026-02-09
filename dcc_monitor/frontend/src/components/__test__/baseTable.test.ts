import type { TableEmitSearchPayload } from '@/types/types.table'
import type { DOMWrapper, VueWrapper } from '@vue/test-utils'
import { beforeEach, describe, expect, test } from 'vitest'
import { mount } from '@vue/test-utils'
import { h, nextTick } from 'vue'

import BaseTable from '@/components/BaseTable.vue'

describe('base table test', () => {
  // notice:
  // 1. 為了測試teleport功能，另外自定義的app component，
  //    因為teleport在mount的時候必須是已經存在的dom，
  //    無法在default slot裡同步進行render
  const app = {
    template: `
      <div id="info"></div>
      <div id="fast-paging"></div>
      <div id="length-menu"></div>
      <div id="pagination"></div>
      <div id="processing"></div>
      <div id="search-button"></div>
      <div id="table"></div>
    `,
  }

  const infoContainer = '#info'
  const fastPagingContainer = '#fast-paging'
  const lengthMenuContainer = '#length-menu'
  const paginationContainer = '#pagination'
  const processingContainer = '#processing'
  const searchButtonContainer = '#search-button'

  const appWrapper = mount(app, {
    attachTo: document.body,
  })

  let tableWrapper: VueWrapper

  beforeEach(() => {
    tableWrapper?.unmount()
    tableWrapper = mount(BaseTable, {
      attachTo: <HTMLElement>appWrapper.find('#table').element,
      global: {
        stubs: ['FontAwesomeIcon'],
      },
      slots: {
        default: (props: { currentItems: [] }) => {
          const children = []
          for (let i = 0; i < props.currentItems.length; i++) {
            children.push(h('div', { 'data-row': i }))
          }
          return children
        },
      },
    })
  })

  test('fast pagination is render correctly with props fastPagingContainer and props items', async () => {
    const items = Array.from({ length: 21 }, () => true)

    let input: DOMWrapper<HTMLInputElement>

    await tableWrapper.setProps({
      items,
      fastPagingContainer,
    })
    input = appWrapper.find(fastPagingContainer).find('input')
    expect(input.exists()).toBeTruthy()
    expect(input.attributes('max')).toBe('3')

    await tableWrapper.setProps({
      pageLength: 25,
    })
    input = appWrapper.find(fastPagingContainer).find('input')
    expect(input.attributes('max')).toBe('1')

    await tableWrapper.setProps({
      items: [],
      fastPagingContainer,
    })
    input = appWrapper.find(fastPagingContainer).find('input')
    expect(input.exists()).toBeFalsy()

    await tableWrapper.setProps({
      items,
      fastPagingContainer: '',
    })
    input = appWrapper.find(fastPagingContainer).find('input')
    expect(input.exists()).toBeFalsy()
  })

  test('length menu is render correctly with props lengthMenu and props lengthMenuContainer', async () => {
    const lengthMenu = [1, 2, 3]

    let select: DOMWrapper<HTMLSelectElement>

    await tableWrapper.setProps({
      lengthMenu,
      lengthMenuContainer,
    })
    select = appWrapper.find(lengthMenuContainer).find('select')
    const options = select.findAll('option')
    expect(select.exists()).toBeTruthy()
    expect(options.length).toBe(lengthMenu.length)
    for (let i = 0; i < lengthMenu.length; i++) {
      expect(options[i].attributes('value')).toBe(lengthMenu[i].toString())
    }

    await tableWrapper.setProps({
      lengthMenu,
      lengthMenuContainer: '',
    })
    select = appWrapper.find(lengthMenuContainer).find('select')
    expect(select.exists()).toBeFalsy()
  })

  test('processing is render correctly with props processingContainer', async () => {
    let overlap: DOMWrapper<HTMLDivElement>

    await tableWrapper.setProps({
      processingContainer,
    })
    overlap = appWrapper.find(processingContainer).find('div')
    expect(overlap.exists()).toBeTruthy()

    await tableWrapper.setProps({
      processingContainer: '',
    })
    overlap = appWrapper.find(processingContainer).find('div')
    expect(overlap.exists()).toBeFalsy()
  })

  test('pagination is functioning correctly with props paginationContainer, length menu and fast pagination', async () => {
    let buttons: DOMWrapper<HTMLButtonElement>[]
    let itemLength: number
    let itemTexts: string[]

    await tableWrapper.setProps({
      paginationContainer,
    })
    buttons = appWrapper.find(paginationContainer).findAll('button')
    expect(buttons.length).toBeGreaterThan(0)

    await tableWrapper.setProps({
      paginationContainer: '',
    })
    buttons = appWrapper.find(paginationContainer).findAll('button')
    expect(buttons.length).toBe(0)

    itemLength = 8
    await tableWrapper.setProps({
      items: new Array(itemLength),
      pageLength: 1,
      paginationContainer,
    })
    buttons = appWrapper.find(paginationContainer).findAll('button')
    expect(buttons.length).toBe(2 + itemLength)

    // 頁數最多就是9個指定頁數按鈕+往前一頁按鈕+往後一頁按鈕
    // 頁數在前面則省略後面
    // 頁數在後面則省略前面
    // 頁數在中間則省略前面及後面
    itemLength = 20
    await tableWrapper.setProps({
      items: new Array(20),
    })
    // []: disabled, (): bg-primary
    // [<] [(1)] 2 3 4 5 [...] 20 >
    buttons = appWrapper.find(paginationContainer).findAll('button')
    itemTexts = ['1', '2', '3', '4', '5', '...', '20']
    expect(buttons.length).toBe(9)
    for (let i = 0; i < 9; i++) {
      if (i === 0 || i === 1 || i === 6) {
        expect(buttons[i].attributes('disabled')).toBe('')
      } else {
        expect(buttons[i].attributes('disabled')).toBeUndefined()
      }

      if (i >= 1 && i <= 7) {
        expect(buttons[i].classes('bg-primary')).toBe(i === 1)
        expect(buttons[i].text()).toBe(itemTexts[i - 1])
      }
    }

    await buttons[2].trigger('click')
    // click page 2
    // < 1 [(2)] 3 4 5 [...] 20 >
    buttons = appWrapper.find(paginationContainer).findAll('button')
    expect(buttons.length).toBe(9)
    for (let i = 0; i < 9; i++) {
      if (i === 2 || i === 6) {
        expect(buttons[i].attributes('disabled')).toBe('')
      } else {
        expect(buttons[i].attributes('disabled')).toBeUndefined()
      }

      if (i >= 1 && i <= 7) {
        expect(buttons[i].classes('bg-primary')).toBe(i === 2)
        expect(buttons[i].text()).toBe(itemTexts[i - 1])
      }
    }

    await buttons[buttons.length - 1].trigger('click')
    // click >
    // < 1 2 [(3)] 4 5 [...] 20 >
    buttons = appWrapper.find(paginationContainer).findAll('button')
    expect(buttons.length).toBe(9)
    for (let i = 0; i < 9; i++) {
      if (i === 3 || i === 6) {
        expect(buttons[i].attributes('disabled')).toBe('')
      } else {
        expect(buttons[i].attributes('disabled')).toBeUndefined()
      }
      if (i >= 1 && i <= 7) {
        expect(buttons[i].classes('bg-primary')).toBe(i === 3)
        expect(buttons[i].text()).toBe(itemTexts[i - 1])
      }
    }

    await buttons[buttons.length - 2].trigger('click')
    // click 20
    // < 1 [...] 16 17 18 19 [(20)] [>]
    buttons = appWrapper.find(paginationContainer).findAll('button')
    itemTexts = ['1', '...', '16', '17', '18', '19', '20']
    expect(buttons.length).toBe(9)
    for (let i = 0; i < 9; i++) {
      if (i === 2 || i === 7 || i === 8) {
        expect(buttons[i].attributes('disabled')).toBe('')
      } else {
        expect(buttons[i].attributes('disabled')).toBeUndefined()
      }
      if (i >= 1 && i <= 7) {
        expect(buttons[i].classes('bg-primary')).toBe(i === 7)
        expect(buttons[i].text()).toBe(itemTexts[i - 1])
      }
    }

    await buttons[3].trigger('click')
    // click 16
    // < 1 [...] 15 [(16)] 17 [...] 20 >
    buttons = appWrapper.find(paginationContainer).findAll('button')
    itemTexts = ['1', '...', '15', '16', '17', '...', '20']
    expect(buttons.length).toBe(9)
    for (let i = 0; i < 9; i++) {
      if (i === 2 || i === 4 || i === 6) {
        expect(buttons[i].attributes('disabled')).toBe('')
      } else {
        expect(buttons[i].attributes('disabled')).toBeUndefined()
      }
      if (i >= 1 && i <= 7) {
        expect(buttons[i].classes('bg-primary')).toBe(i === 4)
        expect(buttons[i].text()).toBe(itemTexts[i - 1])
      }
    }

    await tableWrapper.setProps({
      lengthMenu: [1, 3, 10],
      lengthMenuContainer,
    })
    await appWrapper.find(lengthMenuContainer).find('select').findAll('option').at(1)?.setValue()
    // length menu
    // 1 3 10
    // change page length from 1 to 3
    // < 1 2 3 4 5 [(6)] 7 >
    buttons = appWrapper.find(paginationContainer).findAll('button')
    itemTexts = ['1', '2', '3', '4', '5', '6', '7']
    expect(buttons.length).toBe(9)
    for (let i = 0; i < 9; i++) {
      if (i === 6) {
        expect(buttons[i].attributes('disabled')).toBe('')
      } else {
        expect(buttons[i].attributes('disabled')).toBeUndefined()
      }

      if (i >= 1 && i <= 7) {
        expect(buttons[i].classes('bg-primary')).toBe(i === 6)
        expect(buttons[i].text()).toBe(itemTexts[i - 1])
      }
    }

    await tableWrapper.setProps({
      fastPagingContainer,
    })
    await appWrapper.find(fastPagingContainer).find('input').setValue(1)
    // change page from 6 to 1 by fast pagination input
    // [<] [(1)] 2 3 4 5 6 7 >
    buttons = appWrapper.find(paginationContainer).findAll('button')
    expect(buttons.length).toBe(9)
    for (let i = 0; i < 9; i++) {
      if (i <= 1) {
        expect(buttons[i].attributes('disabled')).toBe('')
      } else {
        expect(buttons[i].attributes('disabled')).toBeUndefined()
      }

      if (i >= 1 && i <= 7) {
        expect(buttons[i].classes('bg-primary')).toBe(i === 1)
        expect(buttons[i].text()).toBe(itemTexts[i - 1])
      }
    }
  })

  test('search button is functioning correctly with props searchButtonContainer', async () => {
    let button: DOMWrapper<HTMLButtonElement>

    await tableWrapper.setProps({
      searchButtonContainer: '',
    })
    button = appWrapper.find(searchButtonContainer).find('button')
    expect(button.exists()).toBeFalsy()

    await tableWrapper.setProps({
      searchButtonContainer,
      processingContainer,
    })
    button = appWrapper.find(searchButtonContainer).find('button')
    expect(button.exists()).toBeTruthy()

    await button.trigger('click')
    const searchEmit = tableWrapper.emitted('search')
    expect(searchEmit).toBeTruthy()
    expect(searchEmit?.length).toBe(1)

    // 前面已有測試searchEmit一定會存在，所以下方一定會執行，可進行測試processing顯示
    if (searchEmit) {
      const emit = <TableEmitSearchPayload>searchEmit[0][0]
      const processingDom = appWrapper.find(processingContainer).find('div')

      emit.showProcessing()
      await nextTick()
      expect(processingDom.element.style.display).toBe('')

      emit.closeProcessing()
      await nextTick()
      expect(processingDom.element.style.display).toBe('none')
    }
  })

  test('info is render correctly with pagination, fast pagination and length menu', async () => {
    const infoContainerDom = appWrapper.find(infoContainer)
    const items = Array.from({ length: 29 })

    let infoDom: DOMWrapper<HTMLDivElement>

    await tableWrapper.setProps({
      infoContainer: '',
    })
    infoDom = infoContainerDom.find('div')
    expect(infoDom.exists()).toBeFalsy()

    await tableWrapper.setProps({
      infoContainer,
    })
    infoDom = infoContainerDom.find('div')
    expect(infoDom.exists()).toBeTruthy()
    expect(infoDom.text()).toBe('显示第 0 到第 0 条纪录，总共 0 条纪录')

    await tableWrapper.setProps({
      items,
      lengthMenu: [1, 2, 3],
      pageLength: 1,
      fastPagingContainer,
      lengthMenuContainer,
      paginationContainer,
    })

    // click page 2
    await appWrapper.find(paginationContainer).findAll('button')[2].trigger('click')
    expect(infoDom.text()).toBe('显示第 2 到第 2 条纪录，总共 29 条纪录')

    // change page from 2 to 4 by fast pagination input
    await appWrapper.find(fastPagingContainer).find('input').setValue(4)
    expect(infoDom.text()).toBe('显示第 4 到第 4 条纪录，总共 29 条纪录')

    // change page length from 1 to 3
    await appWrapper.find(lengthMenuContainer).find('select').findAll('option').at(2)?.setValue()
    expect(infoDom.text()).toBe('显示第 4 到第 6 条纪录，总共 29 条纪录')
  })
})
