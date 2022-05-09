---
title: WhoIs
slug: 5AIN-whois
createdAt: 2022-05-02T17:52:10.000Z
updatedAt: 2022-05-02T22:00:54.000Z
---
#WhoIs
##Overview
##Usage




### Creating a New WhoIs

The following is an example of a request to create a new WhoIs

```
MsgCreateWhoIs {
  string creator;
  string did;
  bytes document;
  repeated Credential credentials;
  string name;
}
```



### Response Definition from CreateWhoIs

The following is an example of a response after creating a new WhoIs

```
MsgCreateWhoIsResponse {
    int32 code;        // Code of the response
    string message;    // Message of the response
    WhoIs who_is;
}
```



### Updating a WhoIs

The following is an example of a request to update a WhoIs

```
MsgUpdateWhoIs {
  string creator;
  string did;
  bytes document;
  repeated Credential credentials;
}
```



### Response Definition from UpdateWhoIs

The following is an example of a response after Updating a WhoIs

```
MsgUpdateWhoIsResponse {
    int32 code = 1;// Code of the response
    string message = 2;// Message of the response
    WhoIs who_is = 3;
}
```





### Deleting a WhoIs

The following is an example of a request to Delet a WhoIs

```
MsgDeleteWhoIs {
  string creator;
  string did;
}
```





### Response Definition from DeleteWhoIs

The following is an example of a response after Deleting a WhoIs

```
MsgDeleteWhoIsResponse {
    int32 code;        // Code of the response
    string message;    // Message of the response
}
```
