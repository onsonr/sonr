import Head from 'next/head'
import { NextStudio } from 'next-sanity/studio'
import { metadata } from 'next-sanity/studio/metadata'

import config from '../../sanity.config'

export default function StudioPage() {
    return (
        <>
            <Head>
                {Object.entries(metadata).map(([key, value]) => (
                    <meta key={key} name={key} content={value} />
                ))}
            </Head>
            <div className="studio-page">
                <NextStudio config={config} />
            </div>
        </>
    )
}
