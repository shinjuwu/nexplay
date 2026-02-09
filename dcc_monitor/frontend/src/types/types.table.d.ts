export interface TableProps<T> {
  items?: T[]

  lengthMenu?: number[]
  pageLength?: number

  fastPagingContainer?: HTMLElement | string | null
  infoContainer?: HTMLElement | string | null
  lengthMenuContainer?: HTMLElement | string | null
  paginationContainer?: HTMLElement | string | null
  processingContainer?: HTMLElement | string | null
  searchButtonContainer?: HTMLElement | string | null
}

export interface TableEmitSearchPayload {
  start: number
  length: number

  showProcessing: () => void
  closeProcessing: () => void
  resetTable: () => void
}
