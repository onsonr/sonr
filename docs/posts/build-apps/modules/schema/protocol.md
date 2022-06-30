---
title: Messages
id: protocol
displayed_sidebar: modulesSidebar
---

# Messages


### `CreateSchema(SchemaDefinition)` 
Register's a new type definition for a given application. this type defention is then asserted against when uploading content

```go
{
    Creator string
    Label   string
    fields  map<string, SchemaKind>
}
```

- `Creator` The identifier of the application the schema is registering for
- `Label` The human readable description of the schema
<!-- Here is where the dev server error was. The link "here" has no url in the parenthesis. -->
<!-- - `fIelds` The data `Schema` being persisted see [here]() for schema data types -->

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
See `Keepers` section for examples of message endpoint usage.