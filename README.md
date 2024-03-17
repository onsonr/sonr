# Generated With [Spawn](https://github.com/rollchains/spawn)

## Module Scaffolding

- `spawn module new <name>` *Generates a Cosmos module template*

## Content Generation

- `make proto-gen` *Generates golang code from proto files, stubs interfaces*

## Testnet

- `make testnet` *IBC testnet from chain <-> local cosmos-hub*
- `make sh-testnet` *Single node, no IBC. quick iteration*

## Local Images

- `make install`      *Builds the chain's binary*
- `make local-image`  *Builds the chain's docker image*

## Testing

- `go test ./... -v` *Unit test*
- `make ictest-*`  *E2E testing*