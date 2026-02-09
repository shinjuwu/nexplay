/* eslint-env node */
import path from 'node:path'
import { createConfig } from './vite.config'

export default createConfig({
  root: path.join(__dirname, '/src/manager'),
  outDir: path.join(__dirname, '/dist/manager'),
  accHeader: 'M7NJpSXxh',
})
