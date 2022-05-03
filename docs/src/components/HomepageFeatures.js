import React from 'react';
import clsx from 'clsx';
import styles from './HomepageFeatures.module.css';

const FeatureList = [
  {
    title: 'Web3 Integrations',
    Svg: require('../../static/img/illustrations/integrations.svg').default,
    description: (
      <>
        Sonr provides easy to use methods to interact with IPFS, Ethereum, and other blockchain networks.
      </>
    ),
  },
  {
    title: 'Universal SDK',
    Svg: require('../../static/img/illustrations/extensive-sdk.svg').default,
    description: (
      <>
        The Universal SDK is a set of libraries and tools that allow developers to easily interact with Sonr's blockchain.
      </>
    ),
  },
  {
    title: 'Cross Platform',
    Svg: require('../../static/img/illustrations/cross-platform.svg').default,
    description: (
      <>
        Sonr is available on desktop, mobile, and web. It's easy to use and easy to integrate with your existing applications.
      </>
    ),
  },
];

function Feature({ Svg, title, description }) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} alt={title} />
      </div>
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
