import { utils, writeFileXLSX } from 'xlsx'

class XlsxUtil {
  // headers [[hCol1, hCol2]]
  // rows [{col1: data11, col2: data21},{col1: data21, col2: data22}]
  // `fileName`.xlsx
  // hCol1  hCol2
  // data11 data21
  // data21 data22
  createAndDownloadFile(fileName, headers, rows) {
    /* generate worksheet and workbook */
    const worksheet = utils.json_to_sheet(rows)
    const workbook = utils.book_new()
    utils.book_append_sheet(workbook, worksheet, 'worksheet')

    /* fix headers */
    utils.sheet_add_aoa(worksheet, headers, { origin: 'A1' })

    /* create an XLSX file and try to save to `fileName`.xlsx */
    writeFileXLSX(workbook, `${fileName}.xlsx`, { compression: true })
  }
}

export default new XlsxUtil()
