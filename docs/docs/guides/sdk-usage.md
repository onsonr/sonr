## `x/did` - Auth & AuthZ

> The DID module is responsible for managing the creation and management of DIDs.
> Controllers represent on-chain accounts backed by a MPC keypair. Controllers
> provide methods for Wallet Account Abstraction (WAA) and are responsible for
> managing the creation and management of DIDs for an individual user.

### Features

- DID Controllers leverage the Cosmos SDK's `x/accounts` std interface for WAA.
- DIDs are represented by a `x/did` controller and are required to state the
  controller's public key, and which map to the controller's capabilities.
- General Sign/Verify methods are provides from the QueryServer for HTTP requests.
- The Execute method is used to broadcast transactions across the network. (TODO)
- Biscuits are used to authenticate and authorize requests between services. (TODO)

### References

- [State](https://github.com/onsonr/sonr/tree/develop/x/did#state)
- [State Transitions](https://github.com/onsonr/sonr/tree/develop/x/did#state-transitions)
- [Messages](https://github.com/onsonr/sonr/tree/develop/x/did#messages)
- [Queries](https://github.com/onsonr/sonr/tree/develop/x/did#query)
- [Params](https://github.com/onsonr/sonr/tree/develop/x/did#params)
- [Client](https://github.com/onsonr/sonr/tree/develop/x/did#client)
- [Future Improvements](https://github.com/onsonr/sonr/tree/develop/x/did#future-improvements)
- [Tests](https://github.com/onsonr/sonr/tree/develop/x/did#tests)
- [Appendix](https://github.com/onsonr/sonr/tree/develop/x/did#appendix)

---

## `x/macaroon`

> The macaroon module is responsible for issuing and verifying macaroons. Macaroons
> are used to authenticate and authorize requests between services.
> Macaroons are requested by NFT Records from [`x/service`](2-‐-Modules-Overview.md#x-service) and granted by controllers from [`x/did`](2-‐-Modules-Overview.md#x/did)

### Features

- On Controller creation, a macaroon is created with an admin scope and a default expiry of _315,569,520 blocks (or ~10 years)_.
- On Service registration, a macaroon is created with a service scope and a default expiry of _31,556,952 blocks (or ~1 year)_.
- Macaroons contain the scope of access for a service and the expiry of the permissions in `blockHeight`.

### References

- [State](https://github.com/onsonr/sonr/tree/develop/x/macaroon#state)
- [State Transitions](https://github.com/onsonr/sonr/tree/develop/x/macaroon#state-transitions)
- [Messages](https://github.com/onsonr/sonr/tree/develop/x/macaroon#messages)
- [Queries](https://github.com/onsonr/sonr/tree/develop/x/macaroon#query)
- [Params](https://github.com/onsonr/sonr/tree/develop/x/macaroon#params)
- [Client](https://github.com/onsonr/sonr/tree/develop/x/macaroon#client)
- [Future Improvements](https://github.com/onsonr/sonr/tree/develop/x/macaroon#future-improvements)
- [Tests](https://github.com/onsonr/sonr/tree/develop/x/macaroon#tests)
- [Appendix](https://github.com/onsonr/sonr/tree/develop/x/macaroon#appendix)

---

## `x/service`

> The service module is responsible for managing decentralized services. Services
> on the Sonr network are essentially on-chain MultiSig wallets that are
> represented by a NFT. Service admins are represented by
> a [`x/did`](2-‐-Modules-Overview.md#x-did) controller and are required to state
> the service's scope of access, and which map to the services' capabilities.

### Features

- Needs a Valid Domain with .htaccess file to be whitelisted.

### References

- [State](https://github.com/onsonr/sonr/tree/develop/x/service#state)
- [State Transitions](https://github.com/onsonr/sonr/tree/develop/x/service#state-transitions)
- [Messages](https://github.com/onsonr/sonr/tree/develop/x/service#messages)
- [Queries](https://github.com/onsonr/sonr/tree/develop/x/service#query)
- [Params](https://github.com/onsonr/sonr/tree/develop/x/service#params)
- [Client](https://github.com/onsonr/sonr/tree/develop/x/service#client)
- [Future Improvements](https://github.com/onsonr/sonr/tree/develop/x/service#future-improvements)
- [Tests](https://github.com/onsonr/sonr/tree/develop/x/service#tests)
- [Appendix](https://github.com/onsonr/sonr/tree/develop/x/service#appendix)

---

## `x/vault`

> The vault module is responsible for managing the storage and acccess-control of
> Decentralized Web Nodes (DWNs) from IPFS. Vaults contain user-facing keys and
> are represented by a [`x/did`](2-‐-Modules-Overview.md#x-did) controller.

### Features

- Vaults can be created by anyone, but efforts are made to restrict 1 per user.
- Vaults are stored in IPFS and when claimed, the bech32 Sonr Address is pinned to IPFS.

### References

- [State](https://github.com/onsonr/sonr/tree/develop/x/vault#state)
- [State Transitions](https://github.com/onsonr/sonr/tree/develop/x/vault#state-transitions)
- [Messages](https://github.com/onsonr/sonr/tree/develop/x/vault#messages)
- [Queries](https://github.com/onsonr/sonr/tree/develop/x/vault#query)
- [Params](https://github.com/onsonr/sonr/tree/develop/x/vault#params)
- [Client](https://github.com/onsonr/sonr/tree/develop/x/vault#client)
- [Future Improvements](https://github.com/onsonr/sonr/tree/develop/x/vault#future-improvements)
- [Tests](https://github.com/onsonr/sonr/tree/develop/x/vault#tests)
- [Appendix](https://github.com/onsonr/sonr/tree/develop/x/vault#appendix)
