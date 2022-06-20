---
title: Blobs
slug: 4iZp-blobs
createdAt: 2022-04-20T19:10:14.000Z
updatedAt: 2022-04-28T18:32:53.000Z
---

# Blobs
## Overview
`Blobs` are instances of data related to files stored on chain. These pieces of data are used to retrieve files from our persistence layer. The following is our data models for blobs being stored on our blockchain (distributed persistence layer). Which is then used to retreive your data from our network file systems.
##Usage
###Uploading a Blob to Sonr

```
MsgUploadBlob {
	state         MessageState
	sizeCache     SizeCache
	unknownFields UnknownFields
	Label string
	Path string
}
```

### Downloading a blob

```
MsgDownloadBlob {
	state
	sizeCache
	unknownFields
	Cid string
	OutPath string
}
```




