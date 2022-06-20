---
title: Buckets
slug: VHNs-buckets
createdAt: 2022-04-20T18:59:08.000Z
updatedAt: 2022-04-28T22:11:21.000Z
---
# Buckets
##Overview
<!--
  asdasdasdasdasd
-->
Similar to Amazon S3 or DigitalOcean Spaces, developers can leverage our decentralized storage module for uploading either application specific assets or user specific assets. While we encourage developers to use our SDK for best results, this storage is S3-compliant.

The Sonr bucket module is used to record the defined collections of Objects utilized by an Application on the Sonr Network. A bucket can be either public access, private access, or restricted access based on Developer configuration. A bucket is used to help organize similar objects for a given application.
##Usage
### Creating a new Bucket

```
MsgCreateBucket {
  string creator;
  string label;
  string description;
  string kind;
  sonrio.sonr.registry.Session session; // Authenticated user session data
  repeated string initial_object_dids;  // Provided initial objects for the bucket
}
```



### Response Definition from CreateBucket

```
MsgCreateBucketResponse {
    int32 code;       // Code of the response
    string message;   // Message of the response
    WhichIs which_is; // Whichis response of the ObjectDoc
}
```



### Updating a Bucket

```azcli
MsgUpdateBucket {
  string creator;
  string label;                          // The Bucket label
  string description;                    // New bucket description
  sonrio.sonr.registry.Session session;  // Session data of authenticated user
  repeated string added_object_dids;     // Added Objects
  repeated string removed_object_dids;   // Removed Objects
}
```



### Response Definition from UpdateBucket

```
MsgUpdateBucketResponse { 
    int32 code;        // Code of the response
    string message;    // Message of the response
    WhichIs which_is;  // Whichis response of the ObjectDoc
}
```



### Deactivating a Bucket

```
MsgDeactivateBucket {
  string creator;
  string did;
  sonrio.sonr.registry.Session session;
}
```





### Response Definition from DeactivateBucket

```
MsgDeactivateBucketResponse {
    int32 code;     // Code of the response
    string message; // Message of the response
}
```


