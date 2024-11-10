import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  server: {
    host: 'localhost',  // Bind the server to localhost
  },
  define: {
    'process.env': process.env,
  },
});
