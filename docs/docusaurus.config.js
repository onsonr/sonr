// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Sonr Docs',
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
        docsRouteBasePath: ["/posts"]
      },
    ],
  ],
  presets: [
    [
      "@docusaurus/preset-classic",
      {
        docs: {
          path: 'posts',
          routeBasePath: 'posts',
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
          disableSwitch: true,
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
            docId: 'guide/introduction',
            position: 'left',
            label: 'Learn the Basics',
          },

          {
            type: 'doc',
            docId: 'run-nodes/blockchain/setup-validator',
            position: 'left',
            label: 'Run Nodes',
          },
          {
            type: 'doc',
            position: 'left',
            label: 'Build Apps',
            docId: 'build-apps/why-sonr',
          },
          {
            position: 'left',
            label: 'Architecture',
            type: 'doc',
            docId: 'architecture/adr-overview',
          },
          {
            position: 'right',
            label: 'API Reference',
            type: 'dropdown',
            items: [
              { href: '/highway', label: '🛣 Highway API' },
              { href: '/blockchain', label: '⛓ Blockchain API' },
            ],
          },
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
                href: 'https://twitter.com/sonr_io',
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
                label: 'GitHub',
                href: 'https://github.com/sonr-io/sonr',
              },
              {
                label: 'Discord',
                href: 'https://discord.gg/fCsyz59h5s',
              }
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} Sonr Inc.`,
      },
    }),
};

module.exports = config;
