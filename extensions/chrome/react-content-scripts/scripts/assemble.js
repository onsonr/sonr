const archiver = require('archiver');
const fs = require('fs');
const path = require('path');

const archive = archiver('zip');
const output = fs.createWriteStream(path.join(__dirname, '../extension.zip'));

const zipContents = async () => {
  return new Promise((resolve, reject) => {
    archive
      .directory(path.join(__dirname, '../build'), false)
      .on('error', (err) => reject(err))
      .pipe(output);

    output.on('close', () => {
      console.log(
        'Done! Your file `extension.zip` is ready to be uploaded to the web store.'
      );
      resolve();
    });
    archive.finalize();
  });
};

const assemble = async () => {
  console.log('Compressing files...');
  await zipContents();
};

assemble();
