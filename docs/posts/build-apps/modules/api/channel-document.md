---
title: Channel Document
slug: ZBT1-channel-document
createdAt: 2022-04-26T14:49:08.000Z
updatedAt: 2022-05-05T15:33:40.000Z
---
#Object Documents
##Overview
Object Documents are used to relate channels to a given object and its correlative type. Once a channel is Created with an associated Object Document it cannot be modified. An object can be deactivated, making it no longer availble to be updated.
##Usage
### Object Types

The following are Types related to a given object.

```azcli
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
