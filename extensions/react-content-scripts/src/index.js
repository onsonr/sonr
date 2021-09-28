/* global chrome */
import React from 'react';
import ReactDOM from 'react-dom';

import App from './app';

// Listen for toggle message
chrome.runtime.onMessage.addListener(function (request) {
  if (request.message === 'toggle_extension') {
    const extension = document.createElement('div');
    extension.id = 'my-chrome-extension';
    document.body.appendChild(extension);
    ReactDOM.render(
      <React.StrictMode>
        <App />
      </React.StrictMode>,
      extension
    );
  }
});
