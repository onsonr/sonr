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
    'guide/introduction',

    'guide/how-it-works',

    {
      type: 'category',
      label: 'Advanced',
      items: ['guide/advanced/identifiers', 'guide/advanced/privacy', 'guide/advanced/security', 'guide/advanced/token'],
      collapsible: true,
      collapsed: true,
    },
  ],

  modulesSidebar: [
    'build-apps/why-sonr',
    'build-apps/sonr-stack',
    'build-apps/installation',
    {
      type: 'category',
      label: 'Peer to Peer',
      items: ['build-apps/motor/access-authentication', 'build-apps/motor/discovery', 'build-apps/motor/transmission'],
      collapsible: true,
      collapsed: false,
    },
    {
      type: 'category',
      label: 'List of Modules',
      items: [{
        type: 'category',
        label: 'Registry',
        items: ['build-apps/modules/registry/overview', 'build-apps/modules/registry/state', 'build-apps/modules/registry/keepers', 'build-apps/modules/registry/messages', 'build-apps/modules/registry/events', 'build-apps/modules/registry/client'],
        collapsible: true,
        collapsed: true,
      },
      {
        type: 'category',
        label: 'Schema',
        items: [
          'build-apps/modules/schema/overview',
          'build-apps/modules/schema/protocol',
          'build-apps/modules/schema/usage',
        ],
        collapsible: true,
        collapsed: true,
      },
      {
        type: 'category',
        label: 'Buckets',
        items: ['build-apps/modules/buckets/overview', 'build-apps/modules/buckets/state', 'build-apps/modules/buckets/keepers', 'build-apps/modules/buckets/messages', 'build-apps/modules/buckets/events', 'build-apps/modules/buckets/client'],
        collapsible: true,
        collapsed: true,
      },
      {
        type: 'category',
        label: 'Channel',
        items: ['build-apps/modules/channel/overview', 'build-apps/modules/channel/state', 'build-apps/modules/channel/keepers', 'build-apps/modules/channel/messages', 'build-apps/modules/channel/events', 'build-apps/modules/channel/client'],
        collapsible: true,
        collapsed: true,
      },],
      collapsible: true,
      collapsed: true,
    },

  ],
  // But you can create a sidebar manually
  runSidebar: [
    'run-nodes/blockchain/setup-validator',
  ],

  // But you can create a sidebar manually
  resourcesSidebar: [
    'architecture/adr-overview',
    'architecture/adr-001',
    'architecture/adr-002',
    'architecture/adr-003',
    'architecture/adr-004',
    'architecture/adr-005',
    'architecture/adr-006'
  ],
};

module.exports = sidebars;
