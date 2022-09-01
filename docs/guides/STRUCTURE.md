# Project Structure

The `sonr` repo follows the Go project structure outlined in https://github.com/golang-standards/project-layout.

The core packages (`/pkg`) is structured as follows:

```text
/app                ->          Main blockchain executable
/bind               ->          Binded motor executable for ios, android, and wasm
/cmd                ->          Executable binaries
/docker             ->          Docker container files
/docs               ->          Documentation with Docusaurus
/internal           ->          Identity management models and interfaces
/pkg                ->          Core packages for all executables
  └─ client         ->          +   Blockchain Client utilities
  └─ config         ->          +   Configuration settings for Motor and Highway nodes
  └─ crypto         ->          +   Cryptographic primitives and Wallet implementation
  └─ did            ->          +   DID management utilities
  └─ fs             ->          +   File System utilities for Motor
  └─ host           ->          +   Libp2p host for Motor & Highway nodes
/proto              ->          Protobuf Definition Files
/scripts            ->          Project specific scripts
/testutil           ->          Testing utilities for simulations
/vue                ->          Vue based wallet UI
/x                  ->          Cosmos Blockchain Implementation
  └─ bucket         ->          +   See /docs/articles/reference/ADR-003.md
  └─ channel        ->          +   See /docs/articles/reference/ADR-004.md
  └─ registry       ->          +   See /docs/articles/reference/ADR-001.md
  └─ schema         ->          +   See /docs/articles/reference/ADR-002.md
```
