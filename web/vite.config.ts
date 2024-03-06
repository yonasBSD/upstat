import path from "path";
import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";
import dotenv from 'dotenv'

dotenv.config()

const BACKEND_URL = "http://localhost:" + process.env.PORT;

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      "/api/": {
        // setting origin in headers is important
        // becuase it returns csrf failed (does not trust the origin error)
        target: BACKEND_URL,
        changeOrigin: true,
        headers: {
          origin: BACKEND_URL,
        },
      },
    },
  },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
});
