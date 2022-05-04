// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Sonr',
  tagline: 'The Internet rebuilt for you',
  url: 'https://sonr-io.github.io',
  organizationName: 'sonr-io', // Usually your GitHub org/user name.
  projectName: 'sonr', // Usually your repo name.
  trailingSlash: false,
  staticDirectories: ['static'],
  baseUrl: '/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',

  plugins: [
    [
      // Search
      require.resolve("@easyops-cn/docusaurus-search-local"),
      {
        language: ["en"],
        hashed: true,
        indexBlog: false,
        docsRouteBasePath: ["/articles"]
      },
    ],
  ],
  presets: [
    [
      "@docusaurus/preset-classic",
      {
        docs: {
          path: 'articles',
          routeBasePath: 'articles',
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          editUrl: 'https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/',
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          editUrl:
            'https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
          colorMode: {
            defaultMode: 'dark',
            disableSwitch: false,
            respectPrefersColorScheme: true,
          },
        },
      },
    ],
    [
      // REST
      "redocusaurus",
      {
      },
    ],
  ],
  themes: ["@saucelabs/theme-github-codeblock"],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      disableSwitch: true,
      navbar: {
        style: 'dark',
        logo: {
          alt: 'Sonr Docs Logo',
          src: 'img/logo.png',
        },
        items: [
          {
            type: 'doc',
            docId: 'introduction',
            position: 'left',
            label: 'Learn the Basics',
          },
          {
            type: 'doc',
            docId: 'sonr-stack',
            position: 'left',
            label: 'Build Apps',
          },
          {
            type: 'doc',
            docId: 'run-nodes/setup-validator',
            position: 'left',
            label: 'Run Nodes',
          },
          {
            position: 'left',
            label: 'API Reference',
            type: 'dropdown',
            items: [
              { href: '/highway', label: 'Highway API' },
              { href: '/blockchain', label: 'Blockchain API' },
            ],
          },
          { to: '/blog', label: 'Guides', position: 'left' },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          // {
          //   title: 'Docs',
          //   items: [
          //     {
          //       label: 'Fundamentals',
          //       to: '/docs/fundamentals/identifiers',
          //     },
          //     {
          //       label: 'Node',
          //       to: '/docs/node/comparison',
          //     },
          //   ],
          // },
          {
            title: 'Resources',
            items: [
              {
                label: 'Stack Overflow',
                href: 'https://stackoverflow.com/questions/tagged/sonr',
              },
              {
                label: 'Community',
                href: 'https://sonr.buzz',
              },
              {
                label: 'Twitter',
                href: 'https://twitter.com/SonrProtocol',
              },
            ],
          },
          {
            title: 'Specifications',
            items: [
              {
                label: 'DID Method',
                href: 'https://github.com/sonr-io/sonr/wiki/ADR-001:-Sonr-DID-Specification',
              },
              {
                label: 'Network Architecture',
                href: 'https://www.figma.com/file/wUulpz1zcRVcj1f2KMQjWB/Sonr-Architecture?node-id=0%3A1',
              },
            ],
          },
          {
            title: 'More',
            items: [
              {
                label: 'Guides',
                to: '/blog',
              },
              {
                label: 'GitHub',
                href: 'https://github.com/sonr-io/sonr',
              },
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} Sonr Inc.`,
      },
    }),
};

module.exports = config;
