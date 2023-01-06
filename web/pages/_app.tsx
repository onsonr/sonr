import type { AppProps } from "next/app";
import { ModalsProvider, SaasProvider, LinkProps } from "@saas-ui/react";
import { extendTheme, type ThemeConfig } from "@chakra-ui/react";
import { theme as baseTheme } from "@saas-ui/pro";
import { theme as glassTheme } from "@saas-ui/theme-glass";
import { withThemeColors } from "@saas-ui/pro";
import colorScheme from "../styles/theme";
import "@fontsource/inter/variable.css";

const config: ThemeConfig = {
  initialColorMode: "dark", // 'dark' | 'light'
  useSystemColorMode: false,
};

// 2. Extend your theme
const theme = extendTheme(
  {
    config,
  },
  glassTheme,
  withThemeColors(colorScheme),
  baseTheme
);

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
