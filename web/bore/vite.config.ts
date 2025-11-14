import path from "node:path";
import { fileURLToPath } from "node:url";
import tailwindcss from "@tailwindcss/vite";
import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";
import checker from "vite-plugin-checker";

const __dirname = path.dirname(fileURLToPath(import.meta.url));

export default defineConfig({
  plugins: [react(), tailwindcss(), checker({ typescript: true })],
  server: { open: true },
  preview: { open: true },
  build: { outDir: "dist", emptyOutDir: true, chunkSizeWarningLimit: 3000 },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
});
