# RFC 004: Account System Refactor

## Status
- Draft v2 (May 2023)

## Current Limitations

1. **Account Representation**: Limited by `google.Protobuf.Any` encapsulation and basic authentication methods
2. **Interface Constraints**: Lacks support for advanced functionalities like vesting and complex auth systems
3. **Implementation Rigidity**: Poor differentiation between account types (e.g., `ModuleAccount`)
4. **Authorization System**: Basic `x/auth` module with limited scope beyond `x/bank` functionality
5. **Dependency Issues**: Cyclic dependencies between modules (e.g., `x/auth` ↔ `x/bank` for vesting)

## Proposal

This proposal aims to transform the way accounts are managed within the Cosmos SDK by introducing significant changes to
their structure and functionality.

### Rethinking Account Representation and Business Logic

Instead of representing accounts as simple `google.Protobuf.Any` structures stored in state with no business logic
attached, this proposal suggests a more sophisticated account representation that is closer to module entities.
In fact, accounts should be able to receive messages and process them in the same way modules do, and be capable of storing
state in a isolated (prefixed) portion of state belonging only to them, in the same way as modules do.

### Account Message Reception

We propose that accounts should be able to receive messages in the same way modules can, allowing them to manage their
own state modifications without relying on other modules. This change would enable more advanced account functionality, such as the
`VestingAccount` example, where the x/bank module previously needed to change the vestingState by casting the abstracted
account to `VestingAccount` and triggering the `TrackDelegation` call. Accounts are already capable of sending messages when
a state transition, originating from a transaction, is executed.

When accounts receive messages, they will be able to identify the sender of the message and decide how to process the
state transition, if at all.

### Consequences

These changes would have significant implications for the Cosmos SDK, resulting in a system of actors that are equal from
the runtime perspective. The runtime would only be responsible for propagating messages between actors and would not
manage the authorization system. Instead, actors would manage their own authorizations. For instance, there would be no
need for the `x/auth` module to manage minting or burning of coins permissions, as it would fall within the scope of the
`x/bank` module.

The key difference between accounts and modules would lie in the origin of the message (state transition). Accounts
(ExternallyOwnedAccount), which have credentials (e.g., a public/private key pairing), originate state transitions from
transactions. In contrast, module state transitions do not have authentication credentials backing them and can be
caused by two factors: either as a consequence of a state transition coming from a transaction or triggered by a scheduler
(e.g., the runtime's Begin/EndBlock).

By implementing these proposed changes, the Cosmos SDK will benefit from a more extensible, versatile, and efficient account
management system that is better suited to address the requirements of the Cosmos ecosystem.

#### Standardization

With `x/accounts` allowing a modular api there becomes a need for standardization of accounts or the interfaces wallets and other clients should expect to use. For this reason we will be using the [`CIP` repo](https://github.com/cosmos/cips) in order to standardize interfaces in order for wallets to know what to expect when interacting with accounts.

## Implementation

### Account Definition

We define the new `Account` type, which is what an account needs to implement to be treated as such.
An `Account` type is defined at APP level, so it cannot be dynamically loaded as the chain is running without upgrading the
node code, unless we create something like a `CosmWasmAccount` which is an account backed by an `x/wasm` contract.

```go
// Account is what the developer implements to define an account.
type Account[InitMsg proto.Message] interface {
	// Init is the function that initialises an account instance of a given kind.
	// InitMsg is used to initialise the initial state of an account.
	Init(ctx *Context, msg InitMsg) error
	// RegisterExecuteHandlers registers an account's execution messages.
	RegisterExecuteHandlers(executeRouter *ExecuteRouter)
	// RegisterQueryHandlers registers an account's query messages.
	RegisterQueryHandlers(queryRouter *QueryRouter)
	// RegisterMigrationHandlers registers an account's migration messages.
	RegisterMigrationHandlers(migrationRouter *MigrationRouter)
}
```

### The InternalAccount definition

The public `Account` interface implementation is then converted by the runtime into an `InternalAccount` implementation,
which contains all the information and business logic needed to operate the account.

```go
type Schema struct {
	state StateSchema // represents the state of an account
	init InitSchema // represents the init msg schema
	exec ExecSchema // represents the multiple execution msg schemas, containing also responses
	query QuerySchema // represents the multiple query msg schemas, containing also responses
	migrate *MigrateSchema // represents the multiple migrate msg schemas, containing also responses, it's optional
}

type InternalAccount struct {
	init    func(ctx *Context, msg proto.Message) (*InitResponse, error)
	execute func(ctx *Context, msg proto.Message) (*ExecuteResponse, error)
	query   func(ctx *Context, msg proto.Message) (proto.Message, error)
    schema  func() *Schema
    migrate func(ctx *Context, msg proto.Message) (*MigrateResponse, error)
}
```

This is an internal view of the account as intended by the system. It is not meant to be what developers implement. An
example implementation of the `InternalAccount` type can be found in [this](https://github.com/testinginprod/accounts-poc/blob/main/examples/recover/recover.go)
example of account whose credentials can be recovered. In fact, even if the `Internal` implementation is untyped (with
respect to `proto.Message`), the concrete implementation is fully typed.

During any of the execution methods of `InternalAccount`, `schema` excluded, the account is given a `Context` which provides:

- A namespaced `KVStore` for the account, which isolates the account state from others (NOTE: no `store keys` needed,
  the account address serves as `store key`).
- Information regarding itself (its address)
- Information regarding the sender.
- ...

#### Init

Init defines the entrypoint that allows for a new account instance of a given kind to be initialised.
The account is passed some opaque protobuf message which is then interpreted and contains the instructions that
constitute the initial state of an account once it is deployed.

An `Account` code can be deployed multiple times through the `Init` function, similar to how a `CosmWasm` contract code
can be deployed (Instantiated) multiple times.

#### Execute

Execute defines the entrypoint that allows an `Account` to process a state transition, the account can decide then how to
process the state transition based on the message provided and the sender of the transition.

#### Query

Query defines a read-only entrypoint that provides a stable interface that links an account with its state. The reason for
which `Query` is still being preferred as an addition to raw state reflection is to:

- Provide a stable interface for querying (state can be optimised and change more frequently than a query)
- Provide a way to define an account `Interface` with respect to its `Read/Write` paths.
- Provide a way to query information that cannot be processed from raw state reflection, ex: compute information from lazy
  state that has not been yet concretely processed (eg: balances with respect to lazy inputs/outputs)

#### Schema

Schema provides the definition of an account from `API` perspective, and it's the only thing that should be taken into account
when interacting with an account from another account or module, for example: an account is an `authz-interface` account if
it has the following message in its execution messages `MsgProxyStateTransition{ state_transition: google.Protobuf.Any }`.

### Migrate

Migrate defines the entrypoint that allows an `Account` to migrate its state from a previous version to a new one. Migrations
can be initiated only by the account itself, concretely this means that the migrate action sender can only be the account address
itself, if the account wants to allow another address to migrate it on its behalf then it could create an execution message
that makes the account migrate itself.

### x/accounts module

In order to create accounts we define a new module `x/accounts`, note that `x/accounts` deploys account with no authentication
credentials attached to it which means no action of an account can be incepted from a TX, we will later explore how the
`x/authn` module uses `x/accounts` to deploy authenticated accounts.

This also has another important implication for which account addresses are now fully decoupled from the authentication mechanism
which makes in turn off-chain operations a little more complex, as the chain becomes the real link between account identifier
and credentials.

We could also introduce a way to deterministically compute the account address.

Note, from the transaction point of view, the `init_message` and `execute_message` are opaque `google.Protobuf.Any`.

The module protobuf definition for `x/accounts` are the following:

```protobuf
// Msg defines the Msg service.
service Msg {
  rpc Deploy(MsgDeploy) returns (MsgDeployResponse);
  rpc Execute(MsgExecute) returns (MsgExecuteResponse);
  rpc Migrate(MsgMigrate) returns (MsgMigrateResponse);
}

message MsgDeploy {
  string sender = 1;
  string kind = 2;
  google.Protobuf.Any init_message = 3;
  repeated google.Protobuf.Any authorize_messages = 4 [(gogoproto.nullable) = false];
}

message MsgDeployResponse {
  string address = 1;
  uint64 id = 2;
  google.Protobuf.Any data = 3;
}

message MsgExecute {
  string sender = 1;
  string address = 2;
  google.Protobuf.Any message = 3;
  repeated google.Protobuf.Any authorize_messages = 4 [(gogoproto.nullable) = false];
}

message MsgExecuteResponse {
  google.Protobuf.Any data = 1;
}

message MsgMigrate {
  string sender = 1;
  string new_account_kind = 2;
  google.Protobuf.Any migrate_message = 3;
}

message MsgMigrateResponse {
  google.Protobuf.Any data = 1;
}

```

#### MsgDeploy

Deploys a new instance of the given account `kind` with initial settings represented by the `init_message` which is a `google.Protobuf.Any`.
Of course the `init_message` can be empty. A response is returned containing the account ID and humanised address, alongside some response
that the account instantiation might produce.

#### Address derivation

In order to decouple public keys from account addresses, we introduce a new address derivation mechanism which is

#### MsgExecute

Sends a `StateTransition` execution request, where the state transition is represented by the `message` which is a `google.Protobuf.Any`.
The account can then decide if to process it or not based on the `sender`.

### MsgMigrate

Migrates an account to a new version of itself, the new version is represented by the `new_account_kind`. The state transition
can only be incepted by the account itself, which means that the `sender` must be the account address itself. During the migration
the account current state is given to the new version of the account, which then executes the migration logic using the `migrate_message`,
it might change state or not, it's up to the account to decide. The response contains possible data that the account might produce
after the migration.

#### Authorize Messages

The `Deploy` and `Execute` messages have a field in common called `authorize_messages`, these messages are messages that the account
can execute on behalf of the sender. For example, in case an account is expecting some funds to be sent from the sender,
the sender can attach a `MsgSend` that the account can execute on the sender's behalf. These authorizations are short-lived,
they live only for the duration of the `Deploy` or `Execute` message execution, or until they are consumed.

An alternative would have been to add a `funds` field, like it happens in cosmwasm, which guarantees the called contract that
the funds are available and sent in the context of the message execution. This would have been a simpler approach, but it would
have been limited to the context of `MsgSend` only, where the asset is `sdk.Coins`. The proposed generic way, instead, allows
the account to execute any message on behalf of the sender, which is more flexible, it could include NFT send execution, or
more complex things like `MsgMultiSend` or `MsgDelegate`, etc.

### Further discussion

#### Sub-accounts

We could provide a way to link accounts to other accounts. Maybe during deployment the sender could decide to link the
newly created to its own account, although there might be use-cases for which the deployer is different from the account
that needs to be linked, in this case a handshake protocol on linking would need to be defined.

#### Predictable address creation

We need to provide a way to create an account with a predictable address, this might serve a lot of purposes, like accounts
wanting to generate an address that:

- nobody else can claim besides the account used to generate the new account
- is predictable

For example:

```protobuf

message MsgDeployPredictable {
  string sender = 1;
  uint32 nonce = 2;
  ...
}
```

And then the address becomes `bechify(concat(sender, nonce))`

`x/accounts` would still use the monotonically increasing sequence as account number.

#### Joining Multiple Accounts

As developers are building new kinds of accounts, it becomes necessary to provide a default way to combine the
functionalities of different account types. This allows developers to avoid duplicating code and enables end-users to
create or migrate to accounts with multiple functionalities without requiring custom development.

To address this need, we propose the inclusion of a default account type called "MultiAccount". The MultiAccount type is
designed to merge the functionalities of other accounts by combining their execution, query, and migration APIs.
The account joining process would only fail in the case of API (intended as non-state Schema APIs) conflicts, ensuring
compatibility and consistency.

With the introduction of the MultiAccount type, users would have the option to either migrate their existing accounts to
a MultiAccount type or extend an existing MultiAccount with newer APIs. This flexibility empowers users to leverage
various account functionalities without compromising compatibility or resorting to manual code duplication.

The MultiAccount type serves as a standardized solution for combining different account functionalities within the
cosmos-sdk ecosystem. By adopting this approach, developers can streamline the development process and users can benefit
from a modular and extensible account system.

# ADR 071: Cryptography v2- Multi-curve support

## Change log

- May 7th 2024: Initial Draft (Zondax AG: @raynaudoe @juliantoledano @jleni @educlerici-zondax @lucaslopezf)
- June 13th 2024: Add CometBFT implementation proposal (Zondax AG: @raynaudoe @juliantoledano @jleni @educlerici-zondax @lucaslopezf)
- July 2nd 2024: Split ADR proposal, add link to ADR in cosmos/crypto (Zondax AG: @raynaudoe @juliantoledano @jleni @educlerici-zondax @lucaslopezf)

## Status

DRAFT

## Abstract

This ADR proposes the refactoring of the existing `Keyring` and `cosmos-sdk/crypto` code to implement [ADR-001-CryptoProviders](https://github.com/cosmos/crypto/blob/main/docs/architecture/adr-001-crypto-provider.md).

For in-depth details of the `CryptoProviders` and their design please refer to ADR mentioned above.

## Introduction

The introduction of multi-curve support in the cosmos-sdk cryptographic package offers significant advantages. By not being restricted to a single cryptographic curve, developers can choose the most appropriate curve based on security, performance, and compatibility requirements. This flexibility enhances the application's ability to adapt to evolving security standards and optimizes performance for specific use cases, helping to future-proofing the sdk's cryptographic capabilities.

The enhancements in this proposal not only render the ["Keyring ADR"](https://github.com/cosmos/cosmos-sdk/issues/14940) obsolete, but also encompass its key aspects, replacing it with a more flexible and comprehensive approach. Furthermore, the gRPC service proposed in the mentioned ADR can be easily implemented as a specialized `CryptoProvider`.

### Glossary

1. **Interface**: In the context of this document, "interface" refers to Go's interface.

2. **Module**: In this document, "module" refers to a Go module.

3. **Package**: In the context of Go, a "package" refers to a unit of code organization.

## Context

In order to fully understand the need for changes and the proposed improvements, it's crucial to consider the current state of affairs:

- The Cosmos SDK currently lacks a comprehensive ADR for the cryptographic package.

- If a blockchain project requires a cryptographic curve that is not supported by the current SDK, the most likely scenario is that they will need to fork the SDK repository and make modifications. These modifications could potentially make the fork incompatible with future updates from the upstream SDK, complicating maintenance and integration.

- Type leakage of specific crypto data types expose backward compatibility and extensibility challenges.

- The demand for a more flexible and extensible approach to cryptography and address management is high.

- Architectural changes are necessary to resolve many of the currently open issues related to new curves support.

- There is a current trend towards modularity in the Interchain stack (e.g., runtime modules).

- Security implications are a critical consideration during the redesign work.

## Objectives

The key objectives for this proposal are:

- Leverage `CryptoProviders`: Utilize them as APIs for cryptographic tools, ensuring modularity, flexibility, and ease of integration.

Developer-Centric Approach

- Prioritize clear, intuitive interfaces and best-practice design principles.

Quality Assurance

- Enhanced Test Coverage: Improve testing methodologies to ensure the robustness and reliability of the module.

## Technical Goals

New Keyring:

- Design a new `Keyring` interface with modular backends injection system to support hardware devices and cloud-based HSMs. This feature is optional and tied to complexity; if it proves too complex, it will be deferred to a future release as an enhancement.

## Proposed architecture

### Components

The main components to be used will be the same as those found in the [ADR-001](https://github.com/cosmos/crypto/blob/main/docs/architecture/adr-001-crypto-provider.md#components).

#### Storage and persistence

The storage and persistence layer is tasked with storing a `CryptoProvider`s. Specifically, this layer must:

- Securely store the crypto provider's associated private key (only if stored locally, otherwise a reference to the private key will be stored instead).
- Store the [`ProviderMetadata`](https://github.com/cosmos/crypto/blob/main/docs/architecture/adr-001-crypto-provider.md#metadata) struct which contains the data that distinguishes that provider.

The purpose of this layer is to ensure that upon retrieval of the persisted data, we can access the provider's type, version, and specific configuration (which varies based on the provider type). This information will subsequently be utilized to initialize the appropriate factory, as detailed in the following section on the factory pattern.

The storage proposal involves using a modified version of the [Record](https://github.com/cosmos/cosmos-sdk/blob/main/proto/cosmos/crypto/keyring/v1/record.proto) struct, which is already defined in **Keyring/v1**. Additionally, we propose utilizing the existing keyring backends (keychain, filesystem, memory, etc.) to store these `Record`s in the same manner as the current **Keyring/v1**.

_Note: This approach will facilitate a smoother migration path from the current Keyring/v1 to the proposed architecture._

Below is the proposed protobuf message to be included in the modified `Record.proto` file

##### Protobuf message structure

The [record.proto](https://github.com/cosmos/cosmos-sdk/blob/main/proto/cosmos/crypto/keyring/v1/record.proto) file will be modified to include the `CryptoProvider` message as an optional field as follows.

```protobuf

// record.proto

message Record {
  string name = 1;
  google.protobuf.Any pub_key = 2;

  oneof item {
    Local local = 3;
    Ledger ledger = 4;
    Multi multi = 5;
    Offline offline = 6;
    CryptoProvider crypto_provider = 7; // <- New
  }

  message Local {
    google.protobuf.Any priv_key = 1;
  }

  message Ledger {
    hd.v1.BIP44Params path = 1;
  }

  message Multi {}

  message Offline {}
}
```

##### Creating and loading a `CryptoProvider`

For creating providers, we propose a _factory pattern_ and a _registry_ for these builders. Examples of these
patterns can be found [here](https://github.com/cosmos/crypto/blob/main/docs/architecture/adr-001-crypto-provider.md#illustrative-code-snippets)

##### Keyring

The new `Keyring` interface will serve as a central hub for managing and fetching `CryptoProviders`. To ensure a smoother migration path, the new Keyring will be backward compatible with the previous version. Since this will be the main API from which applications will obtain their `CryptoProvider` instances, the proposal is to extend the Keyring interface to include the methods:

```go
type KeyringV2 interface {
  // methods from Keyring/v1

  // ListCryptoProviders returns a list of all the stored CryptoProvider metadata.
  ListCryptoProviders() ([]ProviderMetadata, error)

  // GetCryptoProvider retrieves a specific CryptoProvider by its id.
  GetCryptoProvider(id string) (CryptoProvider, error)
}
```

_Note_: Methods to obtain a provider from a public key or other means that make it easier to load the desired provider can be added.

##### Especial use case: remote signers

It's important to note that the `CryptoProvider` interface is versatile enough to be implemented as a remote signer. This capability allows for the integration of remote cryptographic operations, which can be particularly useful in distributed or cloud-based environments where local cryptographic resources are limited or need to be managed centrally.

## Alternatives

It is important to note that all the code presented in this document is not in its final form and could be subject to changes at the time of implementation. The examples and implementations discussed should be interpreted as alternatives, providing a conceptual framework rather than definitive solutions. This flexibility allows for adjustments based on further insights, technical evaluations, or changing requirements as development progresses.

## Decision

We will:

- Leverage crypto providers
- Refactor the module structure as described above.
- Define types and interfaces as the code attached.
- Refactor existing code into new structure and interfaces.
- Implement Unit Tests to ensure no backward compatibility issues.

## Consequences

### Impact on the SDK codebase

We can divide the impact of this ADR into two main categories: state machine code and client related code.

#### Client

The major impact will be on the client side, where the current `Keyring` interface will be replaced by the new `KeyringV2` interface. At first, the impact will be low since `CryptoProvider` is an optional field in the `Record` message, so there's no mandatory requirement for migrating to this new concept right away. This allows a progressive transition where the risks of breaking changes or regressions are minimized.

#### State Machine

The impact on the state machine code will be minimal, the modules affected (at the time of writing this ADR)
are the `x/accounts` module, specifically the `Authenticate` function and the `x/auth/ante` module. This function will need to be adapted to use a `CryptoProvider` service to make use of the `Verifier` instance.

Worth mentioning that there's also the alternative of using `Verifier` instances in a standalone fashion (see note below).

The specific way to adapt these modules will be deeply analyzed and decided at implementation time of this ADR.

_Note_: All cryptographic tools (hashers, verifiers, signers, etc.) will continue to be available as standalone packages that can be imported and utilized directly without the need for a `CryptoProvider` instance. However, the `CryptoProvider` is the recommended method for using these tools as it offers a more secure way to handle sensitive data, enhanced modularity, and the ability to store configurations and metadata within the `CryptoProvider` definition.

### Backwards Compatibility

The proposed migration path is similar to what the cosmos-sdk has done in the past. To ensure a smooth transition, the following steps will be taken:

Once ADR-001 is implemented with a stable release:

- Deprecate the old crypto package. The old crypto package will still be usable, but it will be marked as deprecated and users can opt to use the new package.
- Migrate the codebase to use the new cosmos/crypto package and remove the old crypto one.

### Positive

- Single place of truth
- Easier to use interfaces
- Easier to extend
- Unit test for each crypto package
- Greater maintainability
- Incentivize addition of implementations instead of forks
- Decoupling behavior from implementation
- Sanitization of code

### Negative

- It will involve an effort to adapt existing code.
- It will require attention to detail and audition.

### Neutral

- It will involve extensive testing.

## Test Cases

- The code will be unit tested to ensure a high code coverage
- There should be integration tests around Keyring and CryptoProviders.

> While an ADR is in the DRAFT or PROPOSED stage, this section should contain a
> summary of issues to be solved in future iterations (usually referencing comments
> from a pull-request discussion).
>
> Later, this section can optionally list ideas or improvements the author or
> reviewers found during the analysis of this ADR.

# ADR-71 Bank V2

## Status

DRAFT

## Changelog

- 2024-05-08: Initial Draft (@samricotta, @julienrbrt)

## Abstract

The primary objective of refactoring the bank module is to simplify and enhance the functionality of the Cosmos SDK. Over time the bank module has been burdened with numerous responsibilities including transaction handling, account restrictions, delegation counting, and the minting and burning of coins.

In addition to the above, the bank module is currently too rigid and handles too many tasks, so this proposal aims to streamline the module by focusing on core functions `Send`, `Mint`, and `Burn`.

Currently, the module is split across different keepers with scattered and duplicates functionalities (with 4 send functions for instance).

Additionally, the integration of the token factory into the bank module allows for standardization, and better integration within the core modules.

This rewrite will reduce complexity and enhance the efficiency and UX of the bank module.

## Context

The current implementation of the bank module is characterised by its handling of a broad array of functions, leading to significant complexity in using and extending the bank module.

These issues have underscored the need for a refactoring strategy that simplifies the module’s architecture and focuses on its most essential operations.

Additionally, there is an overlap in functionality with a Token Factory module, which could be integrated to streamline oper.

## Decision

**Permission Tightening**: Access to the module can be restricted to selected denominations only, ensuring that it operates within designated boundaries and does not exceed its intended scope. Currently, the permissions allow all denoms, so this should be changed. Send restrictions functionality will be maintained.

**Simplification of Logic**: The bank module will focus on core functionalities `Send`, `Mint`, and `Burn`. This refinement aims to streamline the architecture, enhancing both maintainability and performance.

**Integration of Token Factory**: The Token Factory will be merged into the bank module. This consolidation of related functionalities aims to reduce redundancy and enhance coherence within the system. Migrations functions will be provided for migrating from Osmosis' Token Factory module to bank/v2.

**Legacy Support**: A legacy wrapper will be implemented to ensure compatibility with about 90% of existing functions. This measure will facilitate a smooth transition while keeping older systems functional.

**Denom Implementation**: A asset interface will be added to standardise interactions such as transfers, balance inquiries, minting, and burning across different tokens. This will allow the bank module to support arbitrary asset types, enabling developers to implement custom, ERC20-like denominations.

For example, currently if a team would like to extend the transfer method the changes would apply universally, affecting all denom’s. With the proposed Asset Interface, it allows teams to customise or extend the transfer method specifically for their own tokens without impacting others.

These improvements are expected to enhance the flexibility of the bank module, allowing for the creation of custom tokens similar to ERC20 standards and assets backed by CosmWasm (CW) contracts. The integration efforts will also aim to unify CW20 with bank coins across the Cosmos chains.

Example of denom interface:

```go
type AssetInterface interface {
    Transfer(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, amount sdk.Coin) error
    Mint(ctx sdk.Context, to sdk.AccAddress, amount sdk.Coin) error
    Burn(ctx sdk.Context, from sdk.AccAddress, amount sdk.Coin) error
    QueryBalance(ctx sdk.Context, account sdk.AccAddress) (sdk.Coin, error)
}
```

Overview of flow:

1. Alice initiates a transfer by entering Bob's address and the amount (100 ATOM)
2. The Bank module verifies that the ATOM token implements the `AssetInterface` by querying the `ATOM_Denom_Account`, which is an `x/account` denom account.
3. The Bank module executes the transfer by subtracting 100 ATOM from Alice’s balance and adding 100 ATOM to Bob’s balance.
4. The Bank module calls the Transfer method on the `ATOM_Denom_Account`. The Transfer method, defined in the `AssetInterface`, handles the logic to subtract 100 ATOM from Alice’s balance and add 100 ATOM to Bob’s balance.
5. The Bank module updates the chain and returns the new balances.
6. Both Alice and Bob successfully receive the updated balances.

## Migration Plans

Bank is a widely used module, so getting a v2 needs to be thought thoroughly. In order to not force all dependencies to immediately migrate to bank/v2, the same _upgrading_ path will be taken as for the `gov` module.

This means `cosmossdk.io/bank` will stay one module and there won't be a new `cosmossdk.io/bank/v2` go module. Instead the bank protos will be versioned from `v1beta1` (current bank) to `v2`.

Bank `v1beta1` endpoints will use the new bank v2 implementation for maximum backward compatibility.

The bank `v1beta1` keepers will be deprecated and potentially eventually removed, but its proto and messages definitions will remain.

Additionally, as bank plans to integrate token factory, migrations functions will be provided to migrate from Osmosis token factory implementation (most widely used implementation) to the new bank/v2 token factory.

## Consequences

### Positive

- Simplified interaction with bank APIs
- Backward compatible changes (no contracts or apis broken)
- Optional migration (note: bank `v1beta1` won't get any new feature after bank `v2` release)

### Neutral

- Asset implementation not available cross-chain (IBC-ed custom asset should possibly fallback to the default implementation)
- Many assets may slow down bank balances requests

### Negative

- Temporarily duplicate functionalities as bank `v1beta1` are `v2` are living alongside
- Difficultity to ever completely remove bank `v1beta1`

### References

- Current bank module implementation: https://github.com/cosmos/cosmos-sdk/blob/v0.50.6/x/bank/keeper/keeper.go#L22-L53
- Osmosis token factory: https://github.com/osmosis-labs/osmosis/tree/v25.0.0/x/tokenfactory/keeper
