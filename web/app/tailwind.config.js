module.exports = {
    content: [
        `components/**/*.{vue,js,ts}`,
        `layouts/**/*.vue`,
        `pages/**/*.vue`,
        `app.vue`,
        `plugins/**/*.{js,ts}`,
        `nuxt.config.{js,ts}`,
    ],
    theme: {
        extend: {
            colors: {
                clifford: '#da373d',
                "starlink-card": "#363636",
                "starlink-darkgray": "#5D5D5D",
                "starlink-lightgray": "#B2B2B2",
                "starlink-success": "#91D54E"
            },
            fontFamily: {
                'sans': ['Helvetica Neue', "sans-serif"],
            }
        }
    },
    darkMode: "media",
    plugins: [],
};
