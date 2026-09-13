/**
 * Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro
 *
 * vite.config.js
 * ==============
 * Vite + React. Dev server on port 5173.
 * Forwards /api to the Go agent at http://localhost:8080.
 */
import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [react()],
  server: {
    port: 5173,
    proxy: {
      "/api": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
    },
  },
});
