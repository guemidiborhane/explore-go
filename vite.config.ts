import { defineConfig, loadEnv } from 'vite'
import preact from '@preact/preset-vite'
import path from 'path'

export default defineConfig(({ mode }) => {
    // Load env file based on `mode` in the current working directory.
    // Set the third parameter to '' to load all env regardless of the `VITE_` prefix.
    const env = loadEnv(mode, process.cwd(), '')
    return {
        build: {
            outDir: './modules/ui/build/',
        },
        resolve: {
            alias: {
                '@': path.resolve(__dirname, './modules/ui/src'),
                '@links': path.resolve(__dirname, './modules/links/ui/')
            }
        },
        plugins: [preact({
            devtoolsInProd: env.NNODE_ENV !== 'production'
        })],
    }
})
