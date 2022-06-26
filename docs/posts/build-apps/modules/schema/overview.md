---
title: Objects
id: objects
displayed_sidebar: highwaySidebar
---

# Schemas
## Introduction
The Sonr Schema module is used to store the records of verifiable objects for a specific application powered by the Sonr Network. Schemas are used to create custom application protcols which can be asserted on in order to verify your application data.

## Overview
Schemas are implemented on the `IPLD Object Model` which allows developers to register specific application data schemas.

## Concepts

### Schema Definition
A `Schema Definition` is used to describe an application Schema that will be stored for later assertion against. the provided `Schema Definition` is then used to Derive both the `WhatIs` and `Schema Reference` that will be recorded on chain. Schemas comply to the `IPLD Object` specification. 


```Text
- (`string`) Creator    : The Account Address signing this message
- (`string`) Label      : Identifier or description of the schema
- (`Map`) Fields        : Map of the initial property names to `SchemaKinds`
```
---
Fields contained within the `SchemaDefinition` are described below:
Each field reperesents an `ipld` see [here](https://ipld.io/docs/schemas/features/typekinds/)
```go
// Represents the types of fields a schema can have
enum SchemaKind {
  INVALID = 0;
  MAP = 1;
  LIST = 2;
  UNIT = 3;
  BOOL = 4;
  INT = 5;
  FLOAT = 6;
  STRING = 7;
  BYTES = 8;
  LINK = 9;
  STRUCT = 10;
  UNION = 11;
  ENUM = 12;
  ANY = 13;
}

```

### What Is records
A `ScehamReference` is used to store information about a `ScehmaDefinition` on our blockchain. This is stored within what is called a `WhatIs` record. Which holds infromation describing the registered Schema.

```go
message WhatIs {
  // DID is the DID of the object
  string did = 1;

  // Object_doc is the object document
  SchemaReference schema = 2;

  // Creator is the DID of the creator
  string creator = 3;

  // Timestamp is the time of the last update of the DID Document
  int64 timestamp = 4;

  // IsActive is the status of the DID Document
  bool is_active = 5;
}

```

### Schema Reference
A `Schema Reference` is used to repersent off chain information related to the `Schema` being registered. This is held within the `WhatIs` record that is written to the chain. A `Schema Reference` helps in retrieving a `Schema` Which is held within other storage.

```go
{
    // the DID for this schema
    did string

    // an alternative reference point
    label string

    // a reference to information stored within an IPFS node.
    cid string
}
```
