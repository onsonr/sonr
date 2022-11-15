# Usage

## Working the Makefile

This repository utilizes a Makefile to automate common tasks. To see a list of available commands, run `make` in the root of the project.

```bash
$ make

 Makefile
 > The following Makefile is used for various actions for the Sonr project.

 bind        :   Binds Android, iOS and Web for Plugin Path
 └─ android       - Android AAR
 └─ ios           - iOS Framework
 └─ web           - iOS Framework
 proto       :   Compiles Go Proto Files and pushes to Buf.Build
 └─ go            - Generate to x/*/types and thirdparty/types/*
 └─ buf           - Build and push to buf.build/sonr-io/blockchain
 clean       :   Clean all artifacts and tidy
```

## Interacting with the Motor

For more information on how to interact with the Motor, see [Motor SDK](https://docs.sonr.io/motor-sdk/overview.html) on the Sonr documentation.
