/* global chrome */
// This file contains utility methods relating to sync storage (Chrome API) and returned as Promises

// This function is a utility used to get values from sync storage
// It takes in an array of keys from storage you want to retrieve
export const getSyncStorage = (keys = []) => {
  return new Promise((resolve) => {
    chrome.storage.sync.get(keys, (result) => {
      resolve(result);
    });
  });
};

// This function is a utility used to set a new value and add it to sync storage
// Simply pass in the key you want to be used and the value
export const setSyncStorage = (key, value) => {
  const test = { [key]: value };
  console.log(test);
  return new Promise((resolve) => {
    chrome.storage.sync.set({ [key]: value }, () => {
      resolve();
    });
  });
};

// This function is a utility to remove items from sync storage
// It takes in an array of keys you want to remove from storage
export const removeSyncStorage = (keys = []) => {
  return new Promise((resolve) => {
    chrome.storage.sync.remove(keys, function () {
      resolve();
    });
  });
};
