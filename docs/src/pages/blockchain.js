import React from 'react';
import ApiDoc from '@theme/ApiDoc';
import useBaseUrl from '@docusaurus/useBaseUrl';

const STATIC_SPEC = '/blockchain.yaml';

function CustomPage() {
    return (
        <ApiDoc
            layoutProps={{
                title: 'Blockchain API',
                description: `Open API Reference Docs Blockchain for API`,
            }}
            spec={{
                type: 'url',
                content: useBaseUrl(STATIC_SPEC),
            }}
        />
    );
}

export default CustomPage;
