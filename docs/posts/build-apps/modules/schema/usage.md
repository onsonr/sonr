---
title: Usage
id: usage
displayed_sidebar: modulesSidebar
---

# Usage

# Introduction

Both `message` endpoints and `Query` endpoints are accessible through `grpc` and the `cli`

## Keeper Message Enpoints
### `CreateSchema(SchemaDefinition)` 
Register's a new type definition for a given application. this type defention is then asserted against when uploading content

```go
{
 creator string 
 label string
 fields map<string, SchemaKind>
}
```

- `Creator` The identifier of the application the schema is registering for
- `Label` The human readable description of the schema
- `Fields` The data `Schema` being persisted

Returns a `WhatIs`

---
### `DeprecateSchema(MsgDeprecateSchema)`
Allows for Schemas to be depricated. Depricated schemas are still accessible but will allow schemas developers to indicate it is no longer supported.

```go
{
  Creator string 
  Did string 
}
```

- `Creator` The Account Address Singing this message
- `DID`     The identifier of the Schema

Returns a `status code` and `message` detailing the operation.

## Usage

### `CreateSchema`

`GRPC`

```bash
$ grpcurl -d '{"creator": <address>, "Label": "Message schema v1" "fields": {"message": 0, "icon": 2}}'  \ 
sonrio.sonr.schema.Msg/MsgCreateSchema
```

### `DepicateSchema`

`GRPC`

```bash
$ grpcurl -d '{"creator": <address>, "did": "did:snr:123"}}'  \ 
sonrio.sonr.schema.Msg/MsgCreateSchema
```

## Query Methods
Query methods are used for accessing state kept within the `Keeper`
### `QuerySchema`
For cases it is dersired to lookup a Schema Definition for verifying on an uploaded object.

```go
{
    Creator string
    Did     string
}
```
- `Creator` identifier of the schema owner which will be a `Application`
- `DID` identifier of the schema being queried for


Returns a `SchemaDefinition`

---
### `QueryWhatIs`
For cases where applications need to verify existing data, or verify a schema belongs to a certain `Creator`. `QueryWhatIs` should be used over `QuerySchema` when only verification of the data is needed.


```go
{
    Creator string
    Did     string
}
```
`
- `Creator` identifier of the schema owner which will be a `Application`
- `DID` identifier of the schema being queried for


Returns a `WhatIs`

---

