/* global chrome */
/*
 * This file includes all background scripts running while the extension
 * is active. React code should not be placed here.
 */

chrome.action.onClicked.addListener(function (tab) {
  // Send a message to the active tab
  chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
    const activeTab = tabs[0];
    chrome.tabs.sendMessage(activeTab.id, {
      message: 'toggle_extension',
    });
  });
});
