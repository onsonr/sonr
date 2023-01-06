import type { AppProps } from "next/app";
import { ModalsProvider, SaasProvider, LinkProps } from "@saas-ui/react";
import { extendTheme } from "@chakra-ui/react";
import { theme as baseTheme } from "@saas-ui/pro";
import NextLink from "next/link";
import { theme as glassTheme } from "@saas-ui/theme-glass";
import "@fontsource/inter/variable.css";

const Link: React.FC<LinkProps> = (props) => {
  return <NextLink {...props} />;
};

// 2. Extend your theme
const theme = extendTheme(
  {
    // your custom theme
  },
  glassTheme,
  baseTheme
);

function App({ Component, pageProps }: AppProps) {
  return (
    <SaasProvider theme={theme} linkComponent={Link}>
      <ModalsProvider>
        <Component {...pageProps} />
      </ModalsProvider>
    </SaasProvider>
  );
}

export default App;
