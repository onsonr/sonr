# Sonr Typescript Client

Sonr is a peer-to-peer identity and asset management system that leverages DID documents, Webauthn, and IPFS — providing users with a secure, user-friendly way to manage their digital identity and assets. This client is written in TypeScript, allowing developers to interact with Sonr's Blockchain effectively.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API](#api)
- [Contributing](#contributing)
- [License](#license)

## Installation

Before you begin, make sure you have Node.js and npm installed on your machine. Then, you can install the Sonr client via npm:

```bash
npm install @sonrhq/client
```

## Usage

Here's a basic example of how to use the Sonr client:

```typescript
import SonrClient from '@sonrhq/client';

// Initialize the client with your service origin
const client = new SonrClient('localhost');

// Check if the user is authenticated
if (client.isAuthenticated()) {
    console.log('User is authenticated.');
} else {
    console.log('User is not authenticated.');
}

// Get the user's address
const address = client.getAddress();
console.log(`Address: ${address}`);

// Get block response
client.getBlock().then(blockResponse => {
    console.log(blockResponse);
}).catch(error => {
    console.error(error);
});

// Get primary document
const primaryDoc = client.getPrimaryDoc();
console.log(`Primary Document: ${JSON.stringify(primaryDoc)}`);
```

For detailed usage, please refer to the [API documentation](#api).

## API

Here is a brief overview of the API provided by the SonrClient class:

### `SonrClient(origin: string)`

Constructor function that initializes a DID and Services object with a given origin.

### `isAuthenticated(): boolean`

Checks if the user is authenticated by verifying if their account information is defined.

### `getAddress(): string`

Returns the address as a string.

### `getBlock(): Promise<BlockResponse>`

Retrieves a block response from a specified URL using axios in TypeScript.

### `getPrimaryDoc(): DidDocument`

Returns the primary document of type DidDocument.

### `register({ alias, onCredentialSet, onRegisterComplete }: SonrRegisterProps): Promise<RegistrationResponse>`

Registers a user by generating a web authentication credential and sending it to the server for verification.

### `login({ alias, onCredentialSet, onLoginComplete }: SonrLoginProps): Promise<LoginResponse>`

Logs in a user by starting and finishing a web authentication process.

## Contributing

Contributions are welcome! Please see our [contributing guide](CONTRIBUTING.md) for more details.

## License

This project is licensed under the terms of the [MIT License](../../../LICENSE).
