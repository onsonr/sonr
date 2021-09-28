import React, { useEffect, useState } from 'react';

import {
  getSyncStorage,
  setSyncStorage,
  removeSyncStorage,
} from './utils/sync-storage';
import graphicSrc from './images/graphic.png';
import './css/index.css';

const STORAGE_KEY = 'sampleName';

const App = () => {
  const [name, setName] = useState('');
  const [title, setTitle] = useState('what is your name?');

  useEffect(() => {
    checkStorage();
  }, []);

  const checkStorage = async () => {
    const result = await getSyncStorage([STORAGE_KEY]);
    if (result[STORAGE_KEY]) {
      setName(result[STORAGE_KEY]);
      setTitle(`${result[STORAGE_KEY]}!`);
    }
  };

  const handleAdd = async () => {
    if (name.length) {
      await setSyncStorage(STORAGE_KEY, name);
      setTitle(`${name}!`);
    }
  };

  const handleRemove = async () => {
    await removeSyncStorage([STORAGE_KEY]);
    setName('');
    setTitle('what is your name?');
  };

  return (
    <div className="App">
      <section className="container">
        <section className="content">
          <img src={graphicSrc} alt="Welcome Graphic" />
          <h1>Sync Storage Starter Template</h1>
          <p>
            This is a basic storage-based Chrome extension (with some basic
            styling added) that utilizes Chrome sync storage (Chrome API) to
            save data to the users Chrome account and supports syncing if they
            are logged into mulitple computers in Chrome. To test, enter your
            name below, close the page, and reopen again. Please read the{' '}
            <b>README.md</b> for more information.
          </p>
          <div className="example">
            <h2 id="example-title">Hello, {title}</h2>
            <input
              type="text"
              id="example-input"
              placeholder="Enter your name"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
            <button id="example-button" onClick={handleAdd}>
              Add Name
            </button>
            <button id="example-remove-button" onClick={handleRemove}>
              Remove Name
            </button>
          </div>
        </section>
      </section>
    </div>
  );
};

export default App;
