// nuxt.config.ts
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  css: ['~/assets/css/main.css'], // Add your Tailwind CSS file
  postcss: {
    plugins: {
      '@tailwindcss/postcss': {}, // Tailwind CSS v4
      autoprefixer: {}, // Autoprefixer
    },
  },
});