import { defineNuxtConfig } from "nuxt3";

// https://v3.nuxtjs.org/docs/directory-structure/nuxt.config
export default defineNuxtConfig({
    buildModules: ["@nuxtjs/tailwindcss"],
    typescript: {
        strict: true,
    },
    build: {
        transpile: ["@heroicons/vue", "chart.js"],
    },
});
