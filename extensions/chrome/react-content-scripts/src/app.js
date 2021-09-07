/* global chrome */
import React from 'react';

import './css/index.css';

/*
 * Note: since we are injecting the app into another web page, we need
 * to access the media via the chrome.runtime API otherwise it will try
 * to access it via the current web page
 */
const App = () => {
  return (
    <div className="my-extension">
      <section className="container">
        <section className="content">
          <img
            src={chrome.runtime.getURL('static/assets/graphic.png')}
            alt="Welcome Graphic"
          />
          <h1>Content Scripts Starter Template</h1>
          <p>
            This is a React-based content scripts Chrome extension (with some
            basic styling added) toggled by clicking the extension's icon.
            Please read the <b>README.md</b> for more information.
          </p>
        </section>
      </section>
    </div>
  );
};

export default App;
