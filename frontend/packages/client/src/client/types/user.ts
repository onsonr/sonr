import { DidDocument } from "./did";

type User = {
    did: string;
    didDocument: DidDocument;
    username: string;
    address: string;
}

type Account = {
    // This field represents the unique account address associated with the user. It is typically a hash or an encoded public key, depending on the underlying blockchain or network.
    address: string;

    // This field contains the human-readable name associated with the account. It is used for easier identification and management of the account by the user.
    name: string;

    // This field stores the Decentralized Identifier (DID) of the user, which is a unique, resolvable, and cryptographically verifiable identifier. DIDs are used to enable secure and decentralized identity management.
    did: string;

    // This field specifies the type of the cryptocurrency or token associated with the account. It is used to differentiate between various cryptocurrencies or tokens that the user may hold in their wallet.
    coin_type: string;

    // This field represents the identifier of the blockchain or network that the account is associated with. Chain IDs are used to distinguish between different blockchains or networks, such as Ethereum, Cosmos, or Filecoin.
    chain_id: string;

    // This field stores the base64 encoded public key of the account, which is used for cryptographic operations such as signing and verifying transactions. The public key is derived from the user's private key and is an essential part of the account's security.
    public_key: string;

    // This field stores the type of the public key. It is used to differentiate between various public key types, such as secp256k1, ed25519, and sr25519.
    type: string;
};

export async function derivePrivateKeyFromWebAuthnCredentialAndPin(
    pin: string
): Promise<CryptoKey> {
    // Generate a new WebAuthn credential
    const credential = await navigator.credentials.create({
        publicKey: {
            challenge: new Uint8Array(16),
            rp: { name: 'Example RP' },
            user: {
                id: new Uint8Array(16),
                name: 'username',
                displayName: 'User Name',
            },
            pubKeyCredParams: [{ type: 'public-key', alg: -7 }], // ES256
            authenticatorSelection: {
                authenticatorAttachment: 'platform',
            },
        },
    });

    if (!credential || credential.type !== 'public-key') {
        throw new Error('Failed to generate WebAuthn credential');
    }

    const publicKeyCredential = credential as PublicKeyCredential;

    // Derive a symmetric key from the PIN using PBKDF2
    const pinBuffer = new TextEncoder().encode(pin);
    const pinHash = await crypto.subtle.digest('SHA-256', pinBuffer);
    const baseKey = await crypto.subtle.importKey(
        'raw',
        pinHash,
        { name: 'PBKDF2' },
        false,
        ['deriveKey']
    );
    const derivedKey = await crypto.subtle.deriveKey(
        {
            name: 'PBKDF2',
            salt: new Uint8Array(16),
            iterations: 100000,
            hash: 'SHA-256',
        },
        baseKey,
        { name: 'AES-GCM', length: 256 },
        true,
        ['encrypt', 'decrypt']
    );

    // Extract the private key from the WebAuthn credential
    // const cosePublicKey = publicKeyCredential.response.getPublicKey();
    // if (!cosePublicKey) {
    //     throw new Error('Failed to get COSE public key from WebAuthn credential');
    // }

    // TODO: Extract the private key from the COSE public key and WebAuthn credential

    // Encrypt the private key with the derived symmetric key
    const encryptedPrivateKey = await crypto.subtle.encrypt(
        {
            name: 'AES-GCM',
            iv: new Uint8Array(12),
        },
        derivedKey,
        new Uint8Array(/* Private key bytes */)
    );

    // Import the encrypted private key as a CryptoKey object
    const privateKey = await crypto.subtle.importKey(
        'pkcs8',
        encryptedPrivateKey,
        { name: 'ECDSA', namedCurve: 'P-256' },
        true,
        ['sign']
    );

    return privateKey;
}


export type { User, Account };
