---
title: Which Is
slug: qJce-which-is
createdAt: 2022-04-28T20:03:37.000Z
updatedAt: 2022-05-02T16:25:32.000Z
---
#WhichIs
##Overview
##Usage
### Create WhichIs

```
MsgCreateWhichIs {
  string creator;
  string did;
  BucketDoc bucket;
}
```



### Definition Response from CreateWhichIs

```
MsgCreateWhichIsResponse {

}
```



### Update WhichIs

```
MsgUpdateWhichIs {
  string creator;
  string did;
  BucketDoc bucket;
}
```



### Definition Response from UpdateWhichIs

```
MsgUpdateWhichIsResponse {}
```



### Deactivate WhichIs

```
message MsgDeleteWhichIs {
  string creator;
  string did;
}
```



### Definition Response from DeactivateWhichIs

```
message MsgDeleteWhichIsResponse {}
```


