// tailwind.config.js
module.exports = {
  content: [
    "./styles/button/**/*.{templ,html}",
    "./styles/form/**/*.{templ,html}",
    "./styles/icon/**/*.{templ,html}",
    "./styles/layout/**/*.{templ,html}",
    "./styles/text/**/*.{templ,html}",
    "./styles/view/**/*.{templ,html}",

    "../gateway/front/**/*.{templ,html}",
    "../vault/front/**/*.{templ,html}",
  ],
  theme: {
    extend: {
      fontFamily: {
        inter: ["Inter", "sans-serif"],
        "inter-tight": ["Inter Tight", "sans-serif"],
        fancy: ['"ZT Bros Oskon 90s"', "sans-serif"],
      },
      fontSize: {
        xs: ["0.75rem", { lineHeight: "1.5" }],
        sm: ["0.875rem", { lineHeight: "1.5715" }],
        base: ["1rem", { lineHeight: "1.5", letterSpacing: "-0.017em" }],
        lg: ["1.125rem", { lineHeight: "1.5", letterSpacing: "-0.017em" }],
        xl: ["1.25rem", { lineHeight: "1.5", letterSpacing: "-0.017em" }],
        "2xl": ["1.5rem", { lineHeight: "1.415", letterSpacing: "-0.017em" }],
        "3xl": ["2rem", { lineHeight: "1.3125", letterSpacing: "-0.017em" }],
        "4xl": ["2.5rem", { lineHeight: "1.25", letterSpacing: "-0.017em" }],
        "5xl": ["3.25rem", { lineHeight: "1.2", letterSpacing: "-0.017em" }],
        "6xl": ["3.75rem", { lineHeight: "1.1666", letterSpacing: "-0.017em" }],
        "7xl": ["4.5rem", { lineHeight: "1.1666", letterSpacing: "-0.017em" }],
      },
      animation: {
        "infinite-scroll": "infinite-scroll 60s linear infinite",
        "infinite-scroll-inverse":
          "infinite-scroll-inverse 60s linear infinite",
      },
      keyframes: {
        "infinite-scroll": {
          from: { transform: "translateX(0)" },
          to: { transform: "translateX(-100%)" },
        },
        "infinite-scroll-inverse": {
          from: { transform: "translateX(-100%)" },
          to: { transform: "translateX(0)" },
        },
      },
    },
  },
  plugins: [require("@tailwindcss/typography"), require("@tailwindcss/forms"), require('tailwindcss-motion')],
};
