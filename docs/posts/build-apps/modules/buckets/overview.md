---
title: Buckets Overview
id: overview
displayed_sidebar: modulesSidebar
---

# `x/buckets`

Similar to Amazon S3 or DigitalOcean Spaces, developers can leverage our decentralized storage module for uploading either application specific assets or user specific assets. While we encourage developers to use our SDK for best results, this storage is S3-compliant.

The Sonr bucket module is used to record the defined collections of Objects utilized by an Application on the Sonr Network. A bucket can be either public access, private access, or restricted access based on Developer configuration. A bucket is used to help organize similar objects for a given application.

## Overview

The record type utilized in the **Bucket module** is the `WhichIs` type. This type provides both an interface to utilize VerifiableCredentials and modify the BucketDoc type definition

### Bucket Types

*   **App-specific bucket** -- a bucket created by the developer containing visual assets and other media germane to the App.

*   **User-specific bucket** -- a bucket contrived through user-created data, facilitated by an **Object** or **Channel**.

## Usage

> Blockchain Methods supplied by Channel Module. Full implementation is still a work in progress.

### `CreateBucket()` - Creates a new bucket implementation for a given application

```Text
- (`string`) Creator                : The Account Address signing this message
- (`Session`) Session               : The Session for the authenticated user
- (`string`) Label                  : Name of the bucket defined by developer
- (`string`) Description            : Description of the bucket defined by developer
- (`string`) Kind                   : Functionality of the bucket i.e. ('public', 'private', 'restricted') *WIP*
- (`List`) InitialObjects           : The initial list of objects to add to the bucket
```

### `UpdateBucket()` - Modifies the bucket configuration and/or updates the bucket objects

```Text
- (`string`) Creator                : The Account Address signing this message
- (`Session`) Session               : The Session for the authenticated user
- (`string`) Label                  : Name of the bucket defined by developer
- (`string`) Description            : Description of the bucket defined by developer
- (`List`) AddedObjects             : The list of objects to add to the bucket
- (`List`) RemovedObjects           : The list of objects to remove from the bucket
```

## Status Codes



```
200 - SUCCESS
300 - MULTIPLE CHOICE
304 - NOT MODIFIED
400 - BAD REQUEST
401 - NOT AUTHORIZED

```
