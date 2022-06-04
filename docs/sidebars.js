/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */

// @ts-check

/** @type {import('@docusaurus/plugin-content-docs').SidebarsConfig} */
const sidebars = {
  // But you can create a sidebar manually
  basicsSidebar: [
    'introduction',
    'why-sonr',
    'how-it-works',
    'sonr-stack',
    {
      type: 'category',
      label: 'Advanced',
      items: ['advanced/identifiers', 'advanced/privacy', 'advanced/security', 'advanced/token'],
      collapsible: true,
      collapsed: true,
    },
  ],

  motorSidebar: [
    'motor-node/access-authentication', 'motor-node/discovery', 'motor-node/transmission',
    {
      type: 'category',
      label: 'Getting Started',
      items: ['motor-node/installation'],
      collapsible: true,
      collapsed: false,
    },
  ],

  highwaySidebar: [
    'highway-sdk/using-cli',
    {
      type: 'category',
      label: 'Modules',
      items: ['highway-sdk/registry', 'highway-sdk/objects', 'highway-sdk/channels', 'highway-sdk/buckets'],
      collapsible: true,
      collapsed: false,
    },
  ],
  // But you can create a sidebar manually
  runSidebar: [
    'run-nodes/setup-validator',
  ],

  // But you can create a sidebar manually
  resourcesSidebar: [
    'reference/adr-overview',
    'reference/adr-001',
    'reference/adr-002',
    'reference/adr-003',
    'reference/adr-004',
    'reference/adr-005',
    'reference/adr-006'
  ],
};

module.exports = sidebars;
