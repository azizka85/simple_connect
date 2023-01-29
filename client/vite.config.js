import { defineConfig, loadEnv } from 'vite'
import solidPlugin from 'vite-plugin-solid'

export default defineConfig(({mode}) => {
  const env = loadEnv(mode, process.cwd())

  const port = +(env.PORT || '3000');

  return {
    plugins: [
      solidPlugin()    
    ],
    server: { port },
    preview: { port },
    build: {
      target: 'esnext'
    },
    envPrefix: ['VITE_', 'API_']
  }
})
