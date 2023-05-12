const { fontFamily } = require("tailwindcss/defaultTheme")

/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: ["class"],
  content: ["pages/**/*.{ts,tsx}", "components/**/*.{ts,tsx}", "pages/*.{ts,tsx}"],
  theme: {
    fontSize: {
      xs: "1.3rem",
      sm: "1.4rem",
      base: "1.5rem",
      md: "1.6rem",
      lg: "1.8rem",
      xl: ["2.2rem", "1.3"],
      "2xl": "2.4rem",
      "3xl": "2.6rem",
      "4xl": "3.2rem",
      "5xl": "4rem",
      "6xl": ["4.4rem", "1.1"],
      "7xl": ["4.8rem", "1.1"],
      "8xl": ["8rem", "1.1"],
    },
    container: {
      center: true,
      padding: "2rem",
      screens: {
        "2xl": "1400px",
      },
    },
    extend: {

      colors: {
        ethereum: "#KC3996B",
        dark: "#0E0D0D",
        gandalf: "#1A1A2E",
        bitcoin: "#F7931A",
        sonr: "#2BC0F6",
        border: "hsl(var(--border))",
        input: "hsl(var(--input))",
        ring: "hsl(var(--ring))",
        background: "hsl(var(--background))",
        foreground: "hsl(var(--foreground))",
        transparent: "transparent",
        white: "#fff",
        "off-white": "#f7f8f8",
        "transparent-white": "rgba(255, 255, 255, 0.08)",
        background: "#000212",
        grey: "#858699",
        "grey-dark": "#222326",
        "primary-text": "#b4bcd0",
        primary: {
          DEFAULT: "hsl(var(--primary))",
          foreground: "hsl(var(--primary-foreground))",
        },
        secondary: {
          DEFAULT: "hsl(var(--secondary))",
          foreground: "hsl(var(--secondary-foreground))",
        },
        destructive: {
          DEFAULT: "hsl(var(--destructive))",
          foreground: "hsl(var(--destructive-foreground))",
        },
        muted: {
          DEFAULT: "hsl(var(--muted))",
          foreground: "hsl(var(--muted-foreground))",
        },
        accent: {
          DEFAULT: "hsl(var(--accent))",
          foreground: "hsl(var(--accent-foreground))",
        },
        popover: {
          DEFAULT: "hsl(var(--popover))",
          foreground: "hsl(var(--popover-foreground))",
        },
        card: {
          DEFAULT: "hsl(var(--card))",
          foreground: "hsl(var(--card-foreground))",
        },
      },
      spacing: {
        0: "0",
        1: "0.4rem",
        2: "0.8rem",
        3: "1.2rem",
        4: "1.6rem",
        5: "2rem",
        6: "2.4rem",
        7: "2.8rem",
        8: "3.2rem",
        9: "3.6rem",
        10: "4rem",
        11: "4.4rem",
        12: "4.8rem",
        13: "5.2rem",
        14: "5.6rem",
        15: "6rem",
        16: "6.4rem",
        "navigation-height": "var(--navigation-height)",
      },
      fontFamily: {
        sans: ["var(--font-sans)", ...fontFamily.sans],
      },
      keyframes: {
        "accordion-down": {
          from: { height: 0 },
          to: { height: "var(--radix-accordion-content-height)" },
        },
        "accordion-up": {
          from: { height: "var(--radix-accordion-content-height)" },
          to: { height: 0 },
        },
        "fade-in": {
          from: { opacity: 0, transform: "translateY(-10px)" },
          to: { opacity: 1, transform: "none" },
        },
        "image-rotate": {
          "0%": { transform: "rotateX(25deg)" },
          "25%": { transform: "rotateX(25deg) scale(0.9)" },
          "60%": { transform: "none" },
          "100%": { transform: "none" },
        },
        "image-glow": {
          "0%": {
            opacity: 0,
            "animation-timing-function": "cubic-bezier(0.74,0.25,0.76,1)",
          },
          "10%": {
            opacity: 1,
            "animation-timing-function": "cubic-bezier(0.12,0.01,0.08,0.99)",
          },
          "100%": {
            opacity: 0.2,
          },
        },
        "sketch-lines": {
          "0%": { "stroke-dashoffset": 1 },
          "50%": { "stroke-dashoffset": 0 },
          "99%": { "stroke-dashoffset": 0 },
          "100%": { visiblity: "hidden" },
        },
        "glow-line-horizontal": {
          "0%": { opacity: 0, transform: "translateX(0)" },
          "5%": { opacity: 1, transform: "translateX(0)" },
          "90%": { opacity: 1 },
          "100%": { opacity: 0, transform: "translateX(min(60vw, 45rem))" },
        },
        "glow-line-vertical": {
          "0%": { opacity: 0, transform: "translateY(0)" },
          "5%": { opacity: 1, transform: "translateY(0)" },
          "90%": { opacity: 1 },
          "100%": { opacity: 0, transform: "translateY(min(21vw, 45rem))" },
        },
        zap: {
          "0%, 9%, 11%, 100% ": {
            fill: "transparent",
          },
          "10%": {
            fill: "white",
          },
        },
        bounce: {
          "50%": {
            transform: "scale(0.98)",
          },
        },
        carouselSlide: {
          '0%': { transform: 'translateX(0%)' },
          '50%': { transform: 'translateX(-150%)' },
          '100%': { transform: 'translateX(0%)' },
        }
      },
      animation: {
        'carousel-slide-slow': 'carouselSlide 12s linear infinite',
        'carousel-slide-medium': 'carouselSlide 10s linear infinite',
        'carousel-slide-fast': 'carouselSlide 8s linear infinite',
        "accordion-down": "accordion-down 0.2s ease-out",
        "accordion-up": "accordion-up 0.2s ease-out",
        "fade-in": "fade-in 1000ms var(--animation-delay, 0ms) ease forwards",
        "image-rotate": "image-rotate 1400ms ease forwards",
        "image-glow": "image-glow 4100ms 600ms ease-out forwards",
        "sketch-lines": "sketch-lines 1200ms ease-out forwards",
        "glow-line-horizontal":
          "glow-line-horizontal var(--animation-duration) ease-in forwards",
        "glow-line-vertical":
          "glow-line-vertical var(--animation-duration) ease-in forwards",
        zap: "zap 2250ms calc(var(--index) * 20ms) linear infinite",
        bounce: "240ms ease 0s 1 running bounce",
      },
    },
  },
  plugins: [require("tailwindcss-animate")],
}
