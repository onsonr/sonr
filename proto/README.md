# Sonr Buf Build Definitions

Contains all `protobuf` definitions for the sonr protocol. Uses `Buf Build` to generate definitions various targets.

## Installation
- Install [buf.build cli](https://buf.build/)
- Install [protoc](https://grpc.io/docs/protoc-installation/)

these installations will be enough local generation of `Go` definitions and publishing to the Sonr `buf registry`
## Building JS/TS definitions
For target TypeScript and JavaScript we will need `protoc` installed and the [](https://www.npmjs.com/package/protoc-gen-ts)

### building proto defs
- in the root of `proto` directory run `buf generate`