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
    {
      type: 'category',
      label: 'Advanced',
      items: ['advanced/identifiers', 'advanced/privacy', 'advanced/security', 'advanced/token'],
      collapsible: true,
      collapsed: true,
    },
  ],
  // But you can create a sidebar manually
  buildSidebar: [
    'sonr-stack',
    {
      type: 'category',
      label: 'Motor Node',
      items: ['motor-node/access-authentication', 'motor-node/discovery', 'motor-node/transmission'],
      collapsible: true,
      collapsed: false,
    },
    {
      type: 'category',
      label: 'Highway SDK',
      items: ['highway-sdk/registry', 'highway-sdk/objects', 'highway-sdk/channels', 'highway-sdk/buckets'],
      collapsible: true,
      collapsed: false,
    },
  ],
  // But you can create a sidebar manually
  runSidebar: [
    'run-nodes/setup-validator',
  ],
};

module.exports = sidebars;
