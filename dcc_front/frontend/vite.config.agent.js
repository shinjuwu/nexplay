/* eslint-env node */
import path from 'node:path'
import { createConfig } from './vite.config'

export default createConfig({
  root: path.join(__dirname, '/src/agent'),
  outDir: path.join(__dirname, '/dist/agent'),
  accHeader: 'aikDg57I1',
})
