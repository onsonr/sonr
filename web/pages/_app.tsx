import type { AppProps } from "next/app";
import { ModalsProvider, SaasProvider, LinkProps } from "@saas-ui/react";
import { extendTheme } from "@chakra-ui/react";
import { theme as baseTheme } from "@saas-ui/pro";
import { theme as glassTheme } from "@saas-ui/theme-glass";
import { withThemeColors } from "@saas-ui/pro";
import "@fontsource/inter/variable.css";

const colors = {
  black: "#151221",
  gray: {
    "50": "#faf9fa",
    "100": "#f2f1f3",
    "200": "#e8e7eb",
    "300": "#d5d3da",
    "400": "#aeabb8",
    "500": "#817c90",
    "600": "#57516b",
    "700": "#3a3251",
    "800": "#221c34",
    "900": "#1b172a",
  },
  blue: {
    "50": "#eef7ff",
    "100": "#c0e2ff",
    "200": "#90cbff",
    "300": "#59b2ff",
    "400": "#1e96ff",
    "500": "#1481df",
    "600": "#116cbb",
    "700": "#0d538f",
    "800": "#0b4475",
    "900": "#093760",
  },
  purple: {
    "50": "#f8f5ff",
    "100": "#e5d9ff",
    "200": "#d1bdff",
    "300": "#b493ff",
    "400": "#9f75ff",
    "500": "#804aff",
    "600": "#6624ff",
    "700": "#5114de",
    "800": "#4311b7",
    "900": "#320c8a",
  },
  pink: {
    "50": "#fff5fa",
    "100": "#ffd6ec",
    "200": "#ffb2db",
    "300": "#ff7ec3",
    "400": "#ff4fad",
    "500": "#e91586",
    "600": "#c81274",
    "700": "#a40f5f",
    "800": "#810c4a",
    "900": "#600937",
  },
  orange: {
    "50": "#fffaf5",
    "100": "#ffe9d7",
    "200": "#ffd0a7",
    "300": "#ffa85d",
    "400": "#fb8117",
    "500": "#d86f13",
    "600": "#b65e10",
    "700": "#914b0d",
    "800": "#733b0a",
    "900": "#5e3009",
  },
  yellow: {
    "50": "#fffefa",
    "100": "#fff9e1",
    "200": "#ffeda4",
    "300": "#ffdd54",
    "400": "#f3c716",
    "500": "#c8a412",
    "600": "#a0830e",
    "700": "#7d660b",
    "800": "#5e4d08",
    "900": "#4d3f07",
  },
  green: {
    "50": "#effff7",
    "100": "#9bffca",
    "200": "#16f57e",
    "300": "#14da70",
    "400": "#11bf62",
    "500": "#0fa454",
    "600": "#0c8846",
    "700": "#0a6a36",
    "800": "#08572d",
    "900": "#064725",
  },
  teal: {
    "50": "#e4fffe",
    "100": "#5dfff9",
    "200": "#15ede5",
    "300": "#13d3cd",
    "400": "#10b4ae",
    "500": "#0e9994",
    "600": "#0b7c78",
    "700": "#09615e",
    "800": "#07514e",
    "900": "#064240",
  },
  cyan: {
    "50": "#edfdff",
    "100": "#b0f4ff",
    "200": "#86efff",
    "300": "#4de7ff",
    "400": "#14c6e1",
    "500": "#13b6cf",
    "600": "#11a4bb",
    "700": "#0e889a",
    "800": "#0b6f7f",
    "900": "#095662",
  },
  primary: {
    "50": "#f1f8ff",
    "100": "#c5e4ff",
    "200": "#91cbff",
    "300": "#4dabff",
    "400": "#2197ff",
    "500": "#147fdd",
    "600": "#116bbb",
    "700": "#0e5696",
    "800": "#0b497f",
    "900": "#08355c",
  },
};

// 2. Extend your theme
const theme = extendTheme({ colors }, glassTheme, baseTheme);

function App({ Component, pageProps }: AppProps) {
  return (
    <SaasProvider theme={theme}>
      <ModalsProvider>
        <Component {...pageProps} />
      </ModalsProvider>
    </SaasProvider>
  );
}

export default App;
