import { theme } from "@saas-ui/pro";

export default {
  colors: {
    primary: theme.colors.blue,
    secondary: theme.colors.orange,
  },
  colorMode: "dark",
  grayTint: "#25241E",
  tokens: {
    "app-background": {
      default: "#FFFFFF",
      _dark: "#25241E",
    },
    "app-text": {
      default: "#37352F",
      _dark: "#FFFFFF",
    },
    "sidebar-background": {
      default: "#F5F1F1",
      _dark: "#1C1B17",
    },
    "sidebar-text": {
      default: "#2E2C2C",
      _dark: "#FAF7F7",
    },
  },
};
