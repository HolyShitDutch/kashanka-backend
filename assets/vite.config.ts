import { defineConfig } from "vite";
import tailwindcss from "@tailwindcss/vite";
import { globSync } from "tinyglobby";

export default defineConfig({
  plugins: [tailwindcss()],
  build: {
    manifest: true,
    outDir: "dist/",
    rolldownOptions: {
      input: globSync("./src/*.ts"),
    },
  },
});
