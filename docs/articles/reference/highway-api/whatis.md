---
title: WhatIs
slug: N4iK-whatis
createdAt: 2022-04-28T19:47:16.000Z
updatedAt: 2022-04-29T18:17:41.000Z
---
#WhatIs
##Overview


##Usage
### Create WhatIs

```
MsgCreateWhatIs {
  string creator;
  string did;
  ObjectDoc object_doc;
}
```

### Definition Response from CreateWhatIs

```
message MsgCreateWhatIsResponse {
  string did;
}
```

### Update WhatIs

```
MsgUpdateWhatIs {
  string creator;
  string did;
  ObjectDoc object_doc;
}
```

### Definition Response from UpdateWhatIs

```
MsgUpdateWhatIsResponse {
  string did;
}
```



### Delete WhatIs

```
MsgDeleteWhatIs {
  string creator;
  string did;
}
```



### Definition Response from DeleteWhatIs

```
MsgDeleteWhatIsResponse {
  string did;
}
```
