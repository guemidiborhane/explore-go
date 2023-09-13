import { defineConfig, loadEnv } from 'vite'
import preact from '@preact/preset-vite'
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
    const env = loadEnv(mode, process.cwd(), '')

    return {
        build: {
            outDir: 'build/'
        },
        resolve: {
            alias: {
                "~": resolve(__dirname, 'static'),
                "@": resolve(__dirname, 'pkg')
            }
        },
        plugins: [preact({
            devtoolsInProd: env.NODE_ENV !== 'production',
        })],
    }
})
