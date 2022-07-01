import React from 'react';
import clsx from 'clsx';
import Layout from '@theme/Layout';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import styles from './index.module.css';
import HomepageFeatures from '../components/HomepageFeatures';
import Head from '@docusaurus/Head';

// To Do 
// Change the buttonGap class's margin-right property to 20px to make it look more centered
// Rename the buttonGap class to mRight20 

function HomepageHeader() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className="container">
        <h1 className="typographySpecialTextDisplay2">Sonr Developer Portal</h1>
        <p className=".typographysRichTextParagraphDefault">{siteConfig.tagline}</p>
        <div className="container">
          <Link 
            className="button button--secondary button--lg typographyRichTextHeading buttonGap"
            to="/articles/introduction"> 
              Get Started
          </Link>
          <Link 
            className="button button--secondary button--lg typographyRichTextHeading"
            to="https://rvhfyn9wf6h.typeform.com/to/xe8LXfoi#hubspot_utk=xxxxx&hubspot_page_name=xxxxx&hubspot_page_url=xxxxx">
            Get Early Access 🎉
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
      
      </main>
    </Layout>


  );
}


