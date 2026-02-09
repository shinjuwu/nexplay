import constant from '@/base/common/constant'

export function columnSorting(direction, aColumn, bColumn) {
  const aColumnType = typeof aColumn
  const bColumnType = typeof bColumn

  if (aColumnType !== bColumnType) return 0

  switch (bColumnType) {
    case 'string':
      return direction === constant.TableSortDirection.Asc
        ? aColumn.localeCompare(bColumn)
        : bColumn.localeCompare(aColumn)
    default:
      return direction === constant.TableSortDirection.Asc ? aColumn - bColumn : bColumn - aColumn
  }
}
