# React - Typescript Popup Page Starter Template

## About This Starter Template

This starter template provides a React (based on Create React App) + TypeScript Chrome extension that triggers a popup page. You simply click the extension icon in the top right or press `Ctrl+Shift+O` to toggle the page popup.

This template includes: React, TypeScript, Prettier, and Eslint.

## Running The Extension

1. Run `npm install` in the folder to install dependencies
2. Run `npm run dev` to start the development server (auto refreshes your extension on change)
3. Navigate to `chrome://extensions` in your browser
4. Ensure `Developer mode` is `enabled` (top right)
5. Click `Load unpacked` in the top left
6. Select the `/build` folder in this directory

After completing the above steps, you should see the developer, unpacked version appear in your extension list as well as the extension icon appear in your browser bar alongside the other extensions you may have installed.

To trigger the extension, simply click the extension icon.

### All Available Commands

- `npm run build` - create a production ready build
- `npm run postbuild` - copies required Chrome extension files after build has completed
- `npm run assemble` - creates a production ready build and zips all files needed to publish in the web store to `extension.zip`
- `npm run dev` - creates a development server that can be used for hot reloading inside the Google Chrome extension ecosystem (see steps 1-6 above)
- `npm run test` - runs any Jest tests you may have
- `npm run pretty` - runs prettier on the `/src` folder
- `npm run lint` - runs eslint on the `/src` folder

## Making Changes

Afer starting the dev server via `npm run dev`, changes will be automatically be rebuilt to `/build` via webpack and the unpacked Chrome extension will automatically be refreshed for you behind the scenes, meaning you don't need to press the refresh button on `chrome://extensions`. Note: you may need to re-toggle or refresh the popup / page to see actual UI changes reflected there after a rebuild (i.e. re-open it again by clicking the icon again).

## Extending The Template

Extending this template would be similar to working on any other Create React App React application. The core of the React app lives in the `/src` and is unopinionated in terms of how it's been setup, so have fun!

## Manifest Explained

There are 2 key sections of the manifest with this example template:

### Browser Action

```
  "browser_action": {
    "default_icon": {
      "32": "icon.png"
    },
    "default_popup": "index.html",
    "default_title": "Open Popup"
  },
```

This portion of the manifest is used to define how the browser action (extension icon) should behave. In this case, when clicked it will trigger a default popup displaying the page `index.html`. This file is the one found is the `/build` folder and as part of the build process, it gets copied with the manifest there as well.

### Custom Commands

```
  "commands": {
    "_execute_browser_action": {
      "suggested_key": {
        "default": "Ctrl+Shift+O",
        "mac": "MacCtrl+Shift+O"
      },
      "description": "Toggle Popup"
    }
  }
```

This portion of the manifest defines custom commands that execute the browser action mentioned above. For example, it defines `Ctrl+Shift+O` as the command that will trigger the browser action and the popup as a result. This can be changed, but keep in mind, commands may conflict with others elsewhere (For example, if you set the command to `Ctrl+S`).

## Preparing to Publish

To prepare for publish, simply run `npm run assemble` which will kick off a production-ready build step and then zip all the contents to `extension.zip` in this folder. This zip file will include all the files you need for your extension to be uploaded.

Note: if this isn't your first publish of this extension, make sure you bump up the verison number in the manifest (for example, `1.0.0` to `1.0.1`), as the web store will require new version in every update you upload.

## Need More help?

Check out the FAQs on https://ChromeExtensionKit.com/ or send an email to `Ryan@ChromeExtensionKit.com` and I will provide assistance as soon as possible.
