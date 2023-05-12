const fs = require('fs');
const path = require('path');


const projectRootDir = path.join(__dirname, '..');
const cryptoIconsDir = path.join(projectRootDir, 'public', 'img', 'crypto');
const outputFile = path.join(cryptoIconsDir, 'icons.json');

fs.readdir(cryptoIconsDir, (err, files) => {
    if (err) {
        console.error('Error reading directory:', err);
        process.exit(1);
    }

    const icons = files
        .filter(file => path.extname(file) === '.png')
        .map(file => `/img/crypto/${file}`);

    fs.writeFile(outputFile, JSON.stringify(icons, null, 2), (err) => {
        if (err) {
            console.error('Error writing JSON file:', err);
            process.exit(1);
        }
        console.log('icons.json file generated successfully.');
    });
});
