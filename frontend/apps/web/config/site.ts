

interface SiteConfig {
  name: string
  description: string
  mainNav: NavItem[]
  links: {
    twitter: string
    github: string
    docs: string
  }
}

interface NavItem {
  label: string
  path: string
}

export const siteConfig: SiteConfig = {
  name: "Sonr",
  description:
    "Beautifully designed components built with Radix UI and Tailwind CSS.",
  mainNav: [
  ],
  links: {
    twitter: "https://twitter.com/sonr_io",
    github: "https://github.com/sonrhq/id",
    docs: "https://docs.sonr.io",
  },
}
