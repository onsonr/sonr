---
title: Objects
slug: SqYX-objects
createdAt: 2022-04-20T19:04:08.000Z
updatedAt: 2022-05-03T17:24:26.000Z
---
#Objects
##Overview
An object repersents unstructured data that is persisted for later use by a Sonr user or application. Objjects can be associated with a given Channel, which will then publish new state of said object when an update to an object occures. See Channels for information on Listening and recieving updates from Objects.
##Usage
### Object Types

```
TypeKind_Invalid
TypeKind_Map
TypeKind_List
TypeKind_Unit
TypeKind_Bool
TypeKind_Int
TypeKind_Float
TypeKind_String
TypeKind_Bytes
TypeKind_Link
TypeKind_Struct
TypeKind_Union
TypeKind_Enum
TypeKind_Any
```

### Creating a new Object

```
MsgCreateObject {
  string creator;
  string label;
  string description;
  repeated TypeField initial_fields;
  sonrio.sonr.registry.Session session;
}
```

### Response Definition from CreateObject

```
MsgCreateObjectResponse {
  int32 code;      // Code of the response
  string message;  // Message of the response
  WhatIs what_is;  // WhatIs of the Channel
}
```

### Upadating an Object

```
MsgUpdateObject {
  string creator;
  string label;                          // Label of the Object
  sonrio.sonr.registry.Session session;  // Authenticated session data
  repeated TypeField added_fields;       // Added fields to the object
  repeated TypeField removed_fields;     // Removed fields from the object
  string cid;                            // Contend Identifier of the object
}
```

# Definition Response from UpdateObject

```azcli
MsgUpdateObjectResponse {
  int32 code;      // Code of the response
  string message;  // Message of the response
  WhatIs what_is;  // WhatIs of the Channel
}
```



### Deactivating an Object

```
MsgDeactivateObject {
  string creator;
  string did;
  sonrio.sonr.registry.Session session;
}
```

### Response Definition from DeactivateObject



```
MsgDeactivateObjectResponse {
  int32 code;      // Code of the response 
  string message;  // Message of the response
}
```














