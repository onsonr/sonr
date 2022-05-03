import React from 'react';
import ApiDoc from '@theme/ApiDoc';
import useBaseUrl from '@docusaurus/useBaseUrl';

const STATIC_SPEC = '/highway.yaml';

function CustomPage() {
    return (
        <ApiDoc
            layoutProps={{
                title: 'Highway API',
                description: `Open API Reference Docs Highway for API`,
            }}
            spec={{
                type: 'url',
                content: useBaseUrl(STATIC_SPEC),
            }}
        />
    );
}

export default CustomPage;
