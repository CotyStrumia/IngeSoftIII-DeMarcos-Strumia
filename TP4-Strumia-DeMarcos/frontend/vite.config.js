import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  // lo dejamos explícito para el pipeline
  build: {
    outDir: 'dist',
    emptyOutDir: true
  },
  server: {
    host: true,
    port: 5173
  },
  preview: {
    host: true,
    port: 5173
  }
})