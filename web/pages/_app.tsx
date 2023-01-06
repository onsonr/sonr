import type { AppProps } from "next/app";
import { ModalsProvider, useModals, SaasProvider } from "@saas-ui/react";
import { extendTheme } from "@chakra-ui/react";
import { theme as baseTheme } from "@saas-ui/pro";
import { theme as glassTheme } from "@saas-ui/theme-glass";
import "@fontsource/inter/variable.css";

// 2. Extend your theme
const theme = extendTheme(
  {
    // your custom theme
  },
  glassTheme,
  baseTheme
);

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <SaasProvider theme={theme}>
      <ModalsProvider>
        <Component {...pageProps} />
      </ModalsProvider>
    </SaasProvider>
  );
}

export default MyApp;
