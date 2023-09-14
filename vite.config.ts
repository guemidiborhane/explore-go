import { defineConfig, loadEnv, splitVendorChunkPlugin } from 'vite'
import preact from '@preact/preset-vite'
import { resolve } from 'path'
import { visualizer } from "rollup-plugin-visualizer";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
    const env = loadEnv(mode, process.cwd(), '')

    return {
        build: {
            outDir: 'build/',
            rollupOptions: {
                output: {
                    manualChunks(id: string) {
                        // creating a chunk to @open-ish deps. Reducing the vendor chunk size
                        if (id.includes('preact')) {
                            return '@base';
                        }
                        // creating a chunk to react routes deps. Reducing the vendor chunk size
                        if (
                            id.includes('react-router-dom') ||
                            id.includes('@remix-run') ||
                            id.includes('react-router')
                        ) {
                            return '@router';
                        }
                    },
                },
            },
        },
        resolve: {
            alias: {
                "~": resolve(__dirname, 'static'),
                "@": resolve(__dirname, 'pkg')
            }
        },
        plugins: [
            preact({
                devtoolsInProd: env.NODE_ENV != 'production'
            }),
            splitVendorChunkPlugin(),
            visualizer()
        ],
    }
})
