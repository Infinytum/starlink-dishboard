import { defineNuxtConfig } from "nuxt";

// https://v3.nuxtjs.org/docs/directory-structure/nuxt.config
export default defineNuxtConfig({
    ssr: false,
    buildModules: ["@nuxtjs/tailwindcss"],
    typescript: {
        strict: true,
    },
    build: {
        transpile: ["@heroicons/vue", "chart.js"],
    }
});
