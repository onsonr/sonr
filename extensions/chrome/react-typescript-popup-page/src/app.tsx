import React from 'react';

import graphicSrc from './images/graphic.png';
import './css/index.css';

const App = () => {
  return (
    <div className="App">
      <section className="content">
        <img src={graphicSrc} alt="Welcome Graphic" />
        <h1>TypeScript Popup Starter Template</h1>
        <p>
          This is a React-based (+ TypeScript) popup Chrome extension (with some
          basic styling added) toggled by clicking the extension's icon or
          pressing <b>Ctrl + Shift + O</b>. Please read the <b>README.md</b> for
          more information.
        </p>
      </section>
    </div>
  );
};

export default App;
