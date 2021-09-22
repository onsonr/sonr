import React, { useEffect, useState } from 'react';

import graphicSrc from './images/graphic.png';
import './css/index.css';

const App = () => {
  const [loading, setLoading] = useState(true);
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    fetch('https://reddit.com/r/webdev.json')
      .then((response) => {
        if (response.status !== 200) {
          console.log('An error has occured.');
          setLoading(false);
          return;
        }

        response.json().then((data) => {
          setLoading(false);
          setPosts(data.data.children);
        });
      })
      .catch((err) => {
        console.log(err);
        setLoading(false);
      });
  }, []);

  return (
    <div className="App">
      <section className="container">
        <section className="content">
          <img src={graphicSrc} alt="Welcome Graphic" />
          <h1>Data Fetching Starter Template</h1>
          <p>
            This is a basic data fetching Chrome extension (with some basic
            styling added) that pulls the latest posts from r/webdev on Reddit.
            Please read the <b>README.md</b> for more information.
          </p>
          <div className="example">
            {loading ? (
              <div>Loading Posts...</div>
            ) : (
              <div>
                {posts.length ? (
                  <div>
                    <h2>Latest Reddit Posts on r/WebDev</h2>
                    <ul>
                      {posts.map((post) => (
                        <li key={post.data.url}>
                          <a href={post.data.url}>{post.data.title}</a>
                        </li>
                      ))}
                    </ul>
                  </div>
                ) : (
                  <div>No posts found.</div>
                )}
              </div>
            )}
          </div>
        </section>
      </section>
    </div>
  );
};

export default App;
