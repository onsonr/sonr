---
title: Overview
---

# Overview

:::note Synopsis
Learn about what the Fee Middleware module is, and how to build custom modules that utilize the Fee Middleware functionality
:::

## What is the Fee Middleware module?

IBC does not depend on relayer operators for transaction verification. However, the relayer infrastructure ensures liveness of the Interchain network — operators listen for packets sent through channels opened between chains, and perform the vital service of ferrying these packets (and proof of the transaction on the sending chain/receipt on the receiving chain) to the clients on each side of the channel.

Though relaying is permissionless and completely decentralized and accessible, it does come with operational costs. Running full nodes to query transaction proofs and paying for transaction fees associated with IBC packets are two of the primary cost burdens which have driven the overall discussion on **a general, in-protocol incentivization mechanism for relayers**.

Initially, a [simple proposal](https://github.com/cosmos/ibc/pull/577/files) was created to incentivize relaying on ICS20 token transfers on the destination chain. However, the proposal was specific to ICS20 token transfers and would have to be reimplemented in this format on every other IBC application module.

After much discussion, the proposal was expanded to a [general incentivisation design](https://github.com/cosmos/ibc/tree/master/spec/app/ics-029-fee-payment) that can be adopted by any ICS application protocol as [middleware](../../01-ibc/04-middleware/02-develop.md).

## Concepts

ICS29 fee payments in this middleware design are built on the assumption that sender chains are the source of incentives — the chain on which packets are incentivized is the chain that distributes fees to relayer operators. However, as part of the IBC packet flow, messages have to be submitted on both sender and destination chains. This introduces the requirement of a mapping of relayer operator's addresses on both chains.

To achieve the stated requirements, the **fee middleware module has two main groups of functionality**:

- Registering of relayer addresses associated with each party involved in relaying the packet on the source chain. This registration process can be automated on start up of relayer infrastructure and happens only once, not every packet flow.

  This is described in the [Fee distribution section](04-fee-distribution.md).

- Escrowing fees by any party which will be paid out to each rightful party on completion of the packet lifecycle.

  This is described in the [Fee messages section](03-msgs.md).

We complete the introduction by giving a list of definitions of relevant terminology.

`Forward relayer`: The relayer that submits the `MsgRecvPacket` message for a given packet (on the destination chain).

`Reverse relayer`: The relayer that submits the `MsgAcknowledgement` message for a given packet (on the source chain).

`Timeout relayer`: The relayer that submits the `MsgTimeout` or `MsgTimeoutOnClose` messages for a given packet (on the source chain).

`Payee`: The account address on the source chain to be paid on completion of the packet lifecycle. The packet lifecycle on the source chain completes with the receipt of a `MsgTimeout`/`MsgTimeoutOnClose` or a `MsgAcknowledgement`.

`Counterparty payee`: The account address to be paid on completion of the packet lifecycle on the destination chain. The package lifecycle on the destination chain completes with a successful `MsgRecvPacket`.

`Refund address`: The address of the account paying for the incentivization of packet relaying. The account is refunded timeout fees upon successful acknowledgement. In the event of a packet timeout, both acknowledgement and receive fees are refunded.

## Known Limitations

- At the time of the release of the feature (ibc-go v4) fee payments middleware only supported incentivisation of new channels; however, with the release of channel upgradeability (ibc-go v8.1) it is possible to enable incentivisation of all existing channels.
- Even though unlikely, there exists a DoS attack vector on a fee-enabled channel if 1) there exists a relayer software implementation that is incentivised to timeout packets if the timeout fee is greater than the sum of the fees to receive and acknowledge the packet, and 2) only this type of implementation is used by operators relaying on the channel. In this situation, an attacker could continuously incentivise the relayers to never deliver the packets by incrementing the timeout fee of the packets above the sum of the receive and acknowledge fees. However, this situation is unlikely to occur because 1) another relayer behaving honestly could relay the packets before they timeout, and 2) the attack would be costly because the attacker would need to incentivise the timeout fee of the packets with their own funds. Given the low impact and unlikelihood of the attack we have decided to accept this risk and not implement any mitigation mesaures.


## Module Integration

The Fee Middleware module, as the name suggests, plays the role of an IBC middleware and as such must be configured by chain developers to route and handle IBC messages correctly.
For Cosmos SDK chains this setup is done via the `app/app.go` file, where modules are constructed and configured in order to bootstrap the blockchain application.

## Example integration of the Fee Middleware module

```go
// app.go

// Register the AppModule for the fee middleware module
ModuleBasics = module.NewBasicManager(
  ...
  ibcfee.AppModuleBasic{},
  ...
)

...

// Add module account permissions for the fee middleware module
maccPerms = map[string][]string{
  ...
  ibcfeetypes.ModuleName:            nil,
}

...

// Add fee middleware Keeper
type App struct {
  ...

  IBCFeeKeeper ibcfeekeeper.Keeper

  ...
}

...

// Create store keys
keys := sdk.NewKVStoreKeys(
  ...
  ibcfeetypes.StoreKey,
  ...
)

...

app.IBCFeeKeeper = ibcfeekeeper.NewKeeper(
  appCodec, keys[ibcfeetypes.StoreKey],
  app.IBCKeeper.ChannelKeeper, // may be replaced with IBC middleware
  app.IBCKeeper.ChannelKeeper,
  &app.IBCKeeper.PortKeeper, app.AccountKeeper, app.BankKeeper,
)


// See the section below for configuring an application stack with the fee middleware module

...

// Register fee middleware AppModule
app.moduleManager = module.NewManager(
  ...
  ibcfee.NewAppModule(app.IBCFeeKeeper),
)

...

// Add fee middleware to begin blocker logic
app.moduleManager.SetOrderBeginBlockers(
  ...
  ibcfeetypes.ModuleName,
  ...
)

// Add fee middleware to end blocker logic
app.moduleManager.SetOrderEndBlockers(
  ...
  ibcfeetypes.ModuleName,
  ...
)

// Add fee middleware to init genesis logic
app.moduleManager.SetOrderInitGenesis(
  ...
  ibcfeetypes.ModuleName,
  ...
)
```

## Configuring an application stack with Fee Middleware

As mentioned in [IBC middleware development](../../01-ibc/04-middleware/02-develop.md) an application stack may be composed of many or no middlewares that nest a base application.
These layers form the complete set of application logic that enable developers to build composable and flexible IBC application stacks.
For example, an application stack may be just a single base application like `transfer`, however, the same application stack composed with `29-fee` will nest the `transfer` base application
by wrapping it with the Fee Middleware module.

### Transfer

See below for an example of how to create an application stack using `transfer` and `29-fee`.
The following `transferStack` is configured in `app/app.go` and added to the IBC `Router`.
The in-line comments describe the execution flow of packets between the application stack and IBC core.

```go
// Create Transfer Stack
// SendPacket, since it is originating from the application to core IBC:
// transferKeeper.SendPacket -> fee.SendPacket -> channel.SendPacket

// RecvPacket, message that originates from core IBC and goes down to app, the flow is the other way
// channel.RecvPacket -> fee.OnRecvPacket -> transfer.OnRecvPacket

// transfer stack contains (from top to bottom):
// - IBC Fee Middleware
// - Transfer

// create IBC module from bottom to top of stack
var transferStack porttypes.IBCModule
transferStack = transfer.NewIBCModule(app.TransferKeeper)
transferStack = ibcfee.NewIBCMiddleware(transferStack, app.IBCFeeKeeper)

// Add transfer stack to IBC Router
ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferStack)
```

### Interchain Accounts

See below for an example of how to create an application stack using `27-interchain-accounts` and `29-fee`.
The following `icaControllerStack` and `icaHostStack` are configured in `app/app.go` and added to the IBC `Router` with the associated authentication module.
The in-line comments describe the execution flow of packets between the application stack and IBC core.

```go
// Create Interchain Accounts Stack
// SendPacket, since it is originating from the application to core IBC:
// icaAuthModuleKeeper.SendTx -> icaController.SendPacket -> fee.SendPacket -> channel.SendPacket

// initialize ICA module with mock module as the authentication module on the controller side
var icaControllerStack porttypes.IBCModule
icaControllerStack = ibcmock.NewIBCModule(&mockModule, ibcmock.NewMockIBCApp("", scopedICAMockKeeper))
app.ICAAuthModule = icaControllerStack.(ibcmock.IBCModule)
icaControllerStack = icacontroller.NewIBCMiddleware(icaControllerStack, app.ICAControllerKeeper)
icaControllerStack = ibcfee.NewIBCMiddleware(icaControllerStack, app.IBCFeeKeeper)

// RecvPacket, message that originates from core IBC and goes down to app, the flow is:
// channel.RecvPacket -> fee.OnRecvPacket -> icaHost.OnRecvPacket

var icaHostStack porttypes.IBCModule
icaHostStack = icahost.NewIBCModule(app.ICAHostKeeper)
icaHostStack = ibcfee.NewIBCMiddleware(icaHostStack, app.IBCFeeKeeper)

// Add authentication module, controller and host to IBC router
ibcRouter.
  // the ICA Controller middleware needs to be explicitly added to the IBC Router because the
  // ICA controller module owns the port capability for ICA. The ICA authentication module
  // owns the channel capability.
  AddRoute(ibcmock.ModuleName+icacontrollertypes.SubModuleName, icaControllerStack) // ica with mock auth module stack route to ica (top level of middleware stack)
  AddRoute(icacontrollertypes.SubModuleName, icaControllerStack).
  AddRoute(icahosttypes.SubModuleName, icaHostStack).
```

## Fee Distribution

Packet fees are divided into 3 distinct amounts in order to compensate relayer operators for packet relaying on fee enabled IBC channels.

- `RecvFee`: The sum of all packet receive fees distributed to a payee for successful execution of `MsgRecvPacket`.
- `AckFee`: The sum of all packet acknowledgement fees distributed to a payee for successful execution of `MsgAcknowledgement`.
- `TimeoutFee`: The sum of all packet timeout fees distributed to a payee for successful execution of `MsgTimeout`.

## Register a counterparty payee address for forward relaying

As mentioned in [ICS29 Concepts](01-overview.md#concepts), the forward relayer describes the actor who performs the submission of `MsgRecvPacket` on the destination chain.
Fee distribution for incentivized packet relays takes place on the packet source chain.

> Relayer operators are expected to register a counterparty payee address, in order to be compensated accordingly with `RecvFee`s upon completion of a packet lifecycle.

The counterparty payee address registered on the destination chain is encoded into the packet acknowledgement and communicated as such to the source chain for fee distribution.
**If a counterparty payee is not registered for the forward relayer on the destination chain, the escrowed fees will be refunded upon fee distribution.**

### Relayer operator actions

A transaction must be submitted **to the destination chain** including a `CounterpartyPayee` address of an account on the source chain.
The transaction must be signed by the `Relayer`.

Note: If a module account address is used as the `CounterpartyPayee` but the module has been set as a blocked address in the `BankKeeper`, the refunding to the module account will fail. This is because many modules use invariants to compare internal tracking of module account balances against the actual balance of the account stored in the `BankKeeper`. If a token transfer to the module account occurs without going through this module and updating the account balance of the module on the `BankKeeper`, then invariants may break and unknown behaviour could occur depending on the module implementation. Therefore, if it is desirable to use a module account that is currently blocked, the module developers should be consulted to gauge to possibility of removing the module account from the blocked list.

```go
type MsgRegisterCounterpartyPayee struct {
  // unique port identifier
  PortId string
  // unique channel identifier
  ChannelId string
  // the relayer address
  Relayer string
  // the counterparty payee address
  CounterpartyPayee string
}
```

> This message is expected to fail if:
>
> - `PortId` is invalid (see [24-host naming requirements](https://github.com/cosmos/ibc/blob/master/spec/core/ics-024-host-requirements/README.md#paths-identifiers-separators).
> - `ChannelId` is invalid (see [24-host naming requirements](https://github.com/cosmos/ibc/blob/master/spec/core/ics-024-host-requirements/README.md#paths-identifiers-separators)).
> - `Relayer` is an invalid address (see [Cosmos SDK Addresses](https://github.com/cosmos/cosmos-sdk/blob/main/docs/learn/beginner/03-accounts.md#addresses)).
> - `CounterpartyPayee` is empty or contains more than 2048 bytes.

See below for an example CLI command:

```bash
simd tx ibc-fee register-counterparty-payee transfer channel-0 \
  cosmos1rsp837a4kvtgp2m4uqzdge0zzu6efqgucm0qdh \
  osmo1v5y0tz01llxzf4c2afml8s3awue0ymju22wxx2 \
  --from cosmos1rsp837a4kvtgp2m4uqzdge0zzu6efqgucm0qdh
```

## Register an alternative payee address for reverse and timeout relaying

As mentioned in [ICS29 Concepts](01-overview.md#concepts), the reverse relayer describes the actor who performs the submission of `MsgAcknowledgement` on the source chain.
Similarly the timeout relayer describes the actor who performs the submission of `MsgTimeout` (or `MsgTimeoutOnClose`) on the source chain.

> Relayer operators **may choose** to register an optional payee address, in order to be compensated accordingly with `AckFee`s and `TimeoutFee`s upon completion of a packet life cycle.

If a payee is not registered for the reverse or timeout relayer on the source chain, then fee distribution assumes the default behaviour, where fees are paid out to the relayer account which delivers `MsgAcknowledgement` or `MsgTimeout`/`MsgTimeoutOnClose`.

### Relayer operator actions

A transaction must be submitted **to the source chain** including a `Payee` address of an account on the source chain.
The transaction must be signed by the `Relayer`.

Note: If a module account address is used as the `Payee` it is recommended to [turn off invariant checks](https://github.com/cosmos/ibc-go/blob/v7.0.0/testing/simapp/app.go#L727) for that module.

```go
type MsgRegisterPayee struct {
  // unique port identifier
  PortId string
  // unique channel identifier
  ChannelId string
  // the relayer address
  Relayer string
  // the payee address
  Payee string
}
```

> This message is expected to fail if:
>
> - `PortId` is invalid (see [24-host naming requirements](https://github.com/cosmos/ibc/blob/master/spec/core/ics-024-host-requirements/README.md#paths-identifiers-separators).
> - `ChannelId` is invalid (see [24-host naming requirements](https://github.com/cosmos/ibc/blob/master/spec/core/ics-024-host-requirements/README.md#paths-identifiers-separators)).
> - `Relayer` is an invalid address (see [Cosmos SDK Addresses](https://github.com/cosmos/cosmos-sdk/blob/main/docs/learn/beginner/03-accounts.md#addresses)).
> - `Payee` is an invalid address (see [Cosmos SDK Addresses](https://github.com/cosmos/cosmos-sdk/blob/main/docs/learn/beginner/03-accounts.md#addresses)).

See below for an example CLI command:

```bash
simd tx ibc-fee register-payee transfer channel-0 \
  cosmos1rsp837a4kvtgp2m4uqzdge0zzu6efqgucm0qdh \
  cosmos153lf4zntqt33a4v0sm5cytrxyqn78q7kz8j8x5 \
  --from cosmos1rsp837a4kvtgp2m4uqzdge0zzu6efqgucm0qdh
```
