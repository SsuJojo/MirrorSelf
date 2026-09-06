import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')
  const allowedHosts = (env.VITE_ALLOWED_HOSTS || '')
    .split(',')
    .map((host) => host.trim())
    .filter(Boolean)

  return {
    plugins: [vue()],
    resolve: {
      alias: {
        '@': path.resolve(__dirname, '/src'),
      },
    },
    server: {
      host: env.VITE_DEV_HOST || '0.0.0.0',
      port: Number(env.VITE_DEV_PORT || 5173),
      strictPort: true,
      cors: true,
      allowedHosts,
      hmr: {
        host: undefined
      },
      proxy: {
        '^/(api|db)': {
          target: env.VITE_BACKEND_URL || 'http://localhost:3001',
          changeOrigin: true,
        }
      }
    }
  }
})
