
<p align="center">
<img width="500" src="https://uploads-ssl.webflow.com/60e4b57e5960f8d0456720e7/60fbc0e3fcdf204c7ed9946b_Github%20-%20Core.png">
</p>

*By [Sonr](https://www.sonr.io), creators of [The Sonr App](https://www.twitter.com/TheSonrApp)*

---

**Core Framework** that manages the Sonr `Libp2p` node in Go, Handles *File Management, Connection to Peer, Pub-Sub for Lobby, and Graph Data Structure*.

> Also Manages Sonr RPC Server

## 🔷 Build
Use `make` with `ios` or `android` or `all` command in root directory, then `flutter` run in [plugin]("https://github.com/sonr-io/plugin") `$HOME/Sonr/plugin/example`


## 🔷 Usage
This project contains a `makefile` with the following commands:
```bash
# Binds Android and iOS for Plugin Path
make bind

# Binds iOS Framework ONLY
make bind.ios

# Binds AAR for Android ONLY
make bind.android

# Compiles Protobuf models for Core Library and Plugin
make proto

# Binds Binary, Creates Protobufs, and Updates App
make upgrade

# Reinitializes Gomobile and Removes Framworks from Plugin
make clean
```

## 🔷 Compatible Types

- Signed integer and floating point types.
- `int` `uint` `int16` `int32` `uint32` `int64` `uint64` `uintptr`

- String and boolean types. `bool` `string`

- Byte slice types. Note that byte slices are passed by reference,
  and support mutation.
  - `byte // alias for uint8`  `rune // alias for int32, represents a Unicode code point`
  - `complex64` `complex128`

- **Any function type all of whose parameters and results have
  supported types.**
  - Functions must return either no results,
  one result, or two results where the type of the second is
  the built-in 'error' type.

  - Any interface type, all of whose exported methods have
  supported function types.

  - Any struct type, all of whose exported methods have
  supported function types and all of whose exported fields
  have supported types.
