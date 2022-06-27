import React from 'react';
import clsx from 'clsx';
import Layout from '@theme/Layout';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import styles from './index.module.css';
import HomepageFeatures from '../components/HomepageFeatures';
import Head from '@docusaurus/Head';
import  ReactTypeformEmbed from 'react-typeform-embed';

function HomepageHeader() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className="container">
        <h1 className="typographySpecialTextDisplay2">Sonr Developer Portal</h1>
        <p className=".typographyRichTextParagraphDefault
">{siteConfig.tagline}</p>
        <div className={styles.buttons}>
          <Link
            className="button button--secondary button--lg typographyRichTextHeading"
            to="/articles/introduction">
            Get Started 🎉
          </Link>
        </div>
      </div>
    </header>
  );
}

export default function Home() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <Layout>
      <Head>
        <title>{siteConfig.title}</title>
        <meta name="description" content={siteConfig.tagline} />
        <meta property="og:image" content="img/open-graph.png" />
      </Head>
      <HomepageHeader />
      <main>
        <HomepageFeatures />
    
      <ReactTypeformEmbed url="https://rvhfyn9wf6h.typeform.com/to/xe8LXfoi#hubspot_utk=xxxxx&hubspot_page_name=xxxxx&hubspot_page_url=xxxxx"/> 
    
    
      </main>
    </Layout>

    
  );
}


