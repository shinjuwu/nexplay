/* eslint-env node */
import { fileURLToPath, URL } from 'node:url'
import path from 'path'

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import { vueI18n } from '@intlify/vite-plugin-vue-i18n'

// https://vitejs.dev/config/
export function createConfig(input) {
  return defineConfig(({ mode }) => {
    const env = loadEnv(mode, path.resolve(process.cwd(), 'env'))

    return {
      root: input.root,
      publicDir: '../../public',
      build: {
        outDir: input.outDir,
      },
      define: {
        __APP_VERSION__: JSON.stringify(env.VITE_APP_VERSION ? env.VITE_APP_VERSION : 'V 0.0.0'),
      },
      plugins: [vue(), vueI18n()],
      resolve: {
        alias: {
          '@': fileURLToPath(new URL('./src', import.meta.url)),
        },
      },
      server: {
        open: true,
        port: env.VITE_SERVER_CLI_PORT,
        proxy: {
          [env.VITE_SERVER_PROXY_API_PATHNAME]: {
            target: `${env.VITE_SERVER_PROXY_CLI_ORIGIN}:${env.VITE_SERVER_PROXY_API_PORT}`,
            changeOrigin: true,
            headers: {
              'X-Dcc-Header': input.accHeader,
            },
          },
        },
      },
    }
  })
}
