<div align="center">
    <img src=".meta/header.png" alt="Sonr-Core-Header"/>
  <br>
</div>

# Description

> Core Framework that manages the Sonr Libp2p node in Go, Handles File Management, Connection to Peer, Pub-Sub for Lobby, and Graph Data Structure.

# Build
> Use `make` with `ios` or `android` or `all` command in root directory, then `flutter` run in [plugin]("https://github.com/sonr-io/plugin") `$HOME/Sonr/plugin/example`

## Compatible types

- Signed integer and floating point types.
- `int` `uint` `int16` `int32` `uint32` `int64` `uint64` `uintptr`

- String and boolean types. `bool` `string`

- Byte slice types. Note that byte slices are passed by reference,
  and support mutation.
  - `byte // alias for uint8`  `rune // alias for int32, represents a Unicode code point`
  - `complex64` `complex128`

**Any function type all of whose parameters and results have
  supported types.**
  Functions must return either no results,
  one result, or two results where the type of the second is
  the built-in 'error' type.

Any interface type, all of whose exported methods have
  supported function types.

Any struct type, all of whose exported methods have
  supported function types and all of whose exported fields
  have supported types.

## ProtoBuf Types

| .proto Type | Notes                                                                                                                                           | Java Type  | Go Type | Dart Type |
|-------------|-------------------------------------------------------------------------------------------------------------------------------------------------|------------|---------|-----------|
| double      |                                                                                                                                                 | double     | float64 | double    |
| float       |                                                                                                                                                 | float      | float32 | double    |
| int32       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int        | int32   | int       |
| int64       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | long       | int64   | Int64     |
| uint32      | Uses variable-length encoding.                                                                                                                  | int[1]     | uint32  | int       |
| uint64      | Uses variable-length encoding.                                                                                                                  | long[1]    | uint64  | Int64     |
| sint32      | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s.                            | int        | int32   | int       |
| sint64      | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s.                            | long       | int64   | Int64     |
| fixed32     | Always four bytes. More efficient than uint32 if values are often greater than 228.                                                             | int[1]     | uint32  | int       |
| fixed64     | Always eight bytes. More efficient than uint64 if values are often greater than 256.                                                            | long[1]    | uint64  | Int64     |
| sfixed32    | Always four bytes.                                                                                                                              | int        | int32   | int       |
| sfixed64    | Always eight bytes.                                                                                                                             | long       | int64   | Int64     |
| bool        |                                                                                                                                                 | boolean    | bool    | bool      |
| string      | A string must always contain UTF-8 encoded or 7-bit ASCII text, and cannot be longer than 232.                                                  | String     | string  | String    |
| bytes       | May contain any arbitrary sequence of bytes no longer than 232.                                                                                 | ByteString | []byte  | List      |
