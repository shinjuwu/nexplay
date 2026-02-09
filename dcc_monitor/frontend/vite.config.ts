import { fileURLToPath, URL } from 'node:url'
import path from 'node:path'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  build: {
    outDir: path.join(__dirname, '/dist/rtp_monitor'),
  },
  server: {
    open: true,
    port: 17784,
    proxy: {
      '/monitor': {
        target: 'http://127.0.0.1:17782',
        changeOrigin: true,
      },
    },
  },
})
