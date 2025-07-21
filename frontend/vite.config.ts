import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

// Use defineConfig to ensure proper TypeScript types are applied
export default defineConfig({
  plugins: [sveltekit()],
  server: {
    // Proxy to Go backend
    proxy: {
      '/api': 'http://localhost:8080'
    }
  },
  optimizeDeps: {
    exclude: ['fsevents']
  }
});
