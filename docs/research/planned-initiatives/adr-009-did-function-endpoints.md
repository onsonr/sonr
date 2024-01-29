---
description: >-
  Functions on Sonr are user-defined functions that can be accessed via HTTP.
  These functions provide a flexible and customizable way to interact with the
  platform.
---

# ADR-009: DID Function Endpoints

## Context

Functions on Sonr are user-defined functions that can be accessed via `HTTP`. These functions provide a flexible and customizable way to interact with the platform. Users can invoke these functions by sending a request to a static endpoint, which includes the function name and a parameter. The parameter is obtained from `ipfs` using the specified `CID`, ensuring the reliability and security of the data. By leveraging this mechanism, users can easily integrate their own logic and algorithms into the Sonr platform, expanding its capabilities and possibilities.

***

## O**bjective**

* Enable the deployment and invocation of user-defined functions
* Allow for custom-defined behavior and interaction with the chain
* Facilitate the implementation of custom logic and algorithms on the Sonr platform

***

## Solution

The proposed solution aims to enable the deployment and invocation of user-defined functions, allowing for custom-defined behavior and interaction with the chain, ultimately enhancing the capabilities and possibilities of the Sonr platform.

***

## Definitions

*   **`URL Parameters`**

    Parameters to functions will be nameless and of type `string`. the provided data will be deserialized to the described types and given to functions as what they are specified to be in the delceration upon creation.
*   `**User Function**`

    A User `Defined Function` is represented by a single `binary` file which is assumed to be executable, and corresponding `callback Urls` associated with said executable file. Said functions are not permitted to return data to the caller directly, but rather the outer managment of the defined function will provide data to the given `urls` as `base64 encoded` representations as to not permit users to modify the system state of any `highway` node directly.

***

## Sequence Methods

### 1. Calling a Function

Static and Variable functions are useful when the function needs to make its own interaction with the chain or fund calls within. A function that’s invoked should have the following defined within the `body` of the `HTTP` request. The Response will simply be an empty body with an `HTTP` status.

| Parameters   | Description                                                                                                                                                                                                                                                                                                                                |
| ------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| DID          | A DID which uniquely identifies the function                                                                                                                                                                                                                                                                                               |
| Label        | A label for conveniently referencing the function.                                                                                                                                                                                                                                                                                         |
| CID          | Reflects whether or not the bucket is active. Inactive functions cannot be modified.                                                                                                                                                                                                                                                       |
| Metadata     | A map that stores the encrypted keys for each party who should have access to the content referenced by a CID. Each CID entry contains a map of public keys (those of the members being granted access) to the JWK encrypted with that public key (PK). This allows the user to find the JWK they need to decrypt the data found with CID. |
| CallbackURLs | Maps developer-defined keys to the CID for content. Content is posted via the APIs described below and can be found again via this map.                                                                                                                                                                                                    |

### 2. **Returning data from a user defined function**

User defined functions cannot modify the state of `highways`, meaning data stored within the network. Instead one or many `callback URLs` are specifiable, allowing for data to be sent to said urls through `url encoding` data return by any function must be `encodable` to `Base64` standard or data will not be sent.

***

## Economic Impact

### Static Cost Functions

Cost is statically defined at function creation time. The run time cost is added in addition to the cost paid by the caller

\$$ Price Formula: total\_{cost} = k + R(t) \$$

| Caller Fees     | Validator Rewards | Service Rewards |
| --------------- | ----------------- | --------------- |
| $total\_{cost}$ | $R(t)$            | $k$             |

***

### Variable Cost Functions

Cost is statically defined at function creation time. The run time cost is added in addition to the cost paid by the caller

\$$ Price Formula: total\_{cost} = k + R(t) + F(c) \$$

| Caller Fees     | Validator Rewards | Service Rewards |
| --------------- | ----------------- | --------------- |
| $total\_{cost}$ | $R(t)$            | $F(c)$          |

***

### Execution-based or “Free” Functions

Cost is statically defined at function creation time. The run time cost is added in addition to the cost paid by the caller.

\$$ Price Formula: total\_{cost} = k \$$

| Caller Fees     | Validator Rewards | Service Rewards |
| --------------- | ----------------- | --------------- |
| $total\_{cost}$ | $k$               | $-$             |

***

## Implementation

### Status

This proposal is **under development** by the core Sonr Team.

| Development Phase | Devnet  |
| ----------------- | ------- |
| Target Completion | Q4 2023 |

***

