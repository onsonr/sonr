---
title: Usage and Examples
id: usage
displayed_sidebar: modulesSidebar
---

# Usage

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
---

