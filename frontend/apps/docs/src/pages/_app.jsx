import Head from 'next/head'
import { Router, useRouter } from 'next/router'
import { MDXProvider } from '@mdx-js/react'

import { Layout } from '@/components/Layout'
import * as mdxComponents from '@/components/mdx'
import { useMobileNavigationStore } from '@/components/MobileNavigation'
import { PlainProvider, Chat } from '@team-plain/react-chat-ui';
import '@/styles/tailwind.css'
import 'focus-visible'

function onRouteChange() {
  useMobileNavigationStore.getState().close()
}

Router.events.on('routeChangeStart', onRouteChange)
Router.events.on('hashChangeStart', onRouteChange)

export default function App({ Component, pageProps }) {
  let router = useRouter()

  return (
    <>
      <Head>
        {router.pathname === '/' ? (
          <title>Sonr API Reference</title>
        ) : (
          <title>{`${pageProps.title} - Sonr API Reference`}</title>
        )}
        <meta name="description" content={pageProps.description} />
              <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png" />
        <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png" />
        <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png" />
        <link rel="manifest" href="/site.webmanifest" />

        <link rel="mask-icon" href="/safari-pinned-tab.svg" color="#5bbad5" />
        <meta name="msapplication-TileColor" content="#da532c" />
        <meta name="theme-color" content="#ffffff" />

      </Head>
      <PlainProvider appKey="appKey_uk_01GXV4STAY3XG3JX46CT9S0FDD">
        <MDXProvider components={mdxComponents}>
          <Layout {...pageProps}>
            <Component {...pageProps} />

            {/* <Chat /> */}
          </Layout>
        </MDXProvider>

      </PlainProvider>

    </>
  )
}
