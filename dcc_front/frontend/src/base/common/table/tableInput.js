import constant from '@/base/common/constant'

export class BaseTableInput {
  constructor(length, column, dir) {
    this.draw = 0
    this.start = 0
    this.length = length
    this.column = column
    this.dir = dir
    this.showProcessing = false
  }

  filterAdjust(filterTotal) {
    while (this.start > filterTotal) {
      this.start -= this.length
    }
  }

  pageChange(start) {
    this.start = start
    this.draw++
  }

  lengthChange(length) {
    this.start = this.start - (this.start % length)
    this.length = length
    this.draw++
  }

  sorting(column) {
    if (this.column !== column) {
      this.column = column
      this.dir = constant.TableSortDirection.Asc
    } else {
      this.dir =
        this.dir === constant.TableSortDirection.Asc
          ? constant.TableSortDirection.Desc
          : constant.TableSortDirection.Asc
    }
    this.draw++
  }
}

export class BaseServersideTableInput extends BaseTableInput {
  constructor(length, column, dir) {
    super(length, column, dir)
    this.totalRecords = 0
  }
}
