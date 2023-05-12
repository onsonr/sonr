import { defineConfig } from 'sanity';
import { deskTool } from 'sanity/desk';
import { visionTool } from '@sanity/vision';
import { singletonPlugin } from './studio/plugins/singletonPlugin';
import { previewDocumentNode } from './studio/plugins/preview';
import { schemasTypes } from './studio/schemas';
import { structure } from './studio/structure';

export default defineConfig({
    projectId: process.env.NEXT_PUBLIC_SANITY_PROJECT_ID as string,
    dataset: process.env.NEXT_PUBLIC_SANITY_DATASET as string,
    name: 'Studio',
    basePath: '/studio',
    schema: {
        types: schemasTypes
    },
    plugins: [
        deskTool({
            structure,
            defaultDocumentNode: previewDocumentNode()
        }),
        singletonPlugin({ types: ['siteSettings'] }),
        visionTool()
    ]
});
