process.env.BABEL_ENV = 'production';
process.env.NODE_ENV = 'production';

const fs = require('fs');
const path = require('path');

const copyManifest = () =>
  fs.copyFileSync(
    path.join(__dirname, '../manifest.json'),
    path.join(__dirname, '../build/manifest.json')
  );

const copyBackgroundScripts = () =>
  fs.copyFileSync(
    path.join(__dirname, '../background.js'),
    path.join(__dirname, '../build/background.js')
  );

const build = async () => {
  console.log('Copying remaining files...');
  await copyManifest();
  await copyBackgroundScripts();
  console.log('Build complete!');
};

build();
