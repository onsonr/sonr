"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.derivePrivateKeyFromWebAuthnCredentialAndPin = void 0;
async function derivePrivateKeyFromWebAuthnCredentialAndPin(pin) {
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
            pubKeyCredParams: [{ type: 'public-key', alg: -7 }],
            authenticatorSelection: {
                authenticatorAttachment: 'platform',
            },
        },
    });
    if (!credential || credential.type !== 'public-key') {
        throw new Error('Failed to generate WebAuthn credential');
    }
    const publicKeyCredential = credential;
    // Derive a symmetric key from the PIN using PBKDF2
    const pinBuffer = new TextEncoder().encode(pin);
    const pinHash = await crypto.subtle.digest('SHA-256', pinBuffer);
    const baseKey = await crypto.subtle.importKey('raw', pinHash, { name: 'PBKDF2' }, false, ['deriveKey']);
    const derivedKey = await crypto.subtle.deriveKey({
        name: 'PBKDF2',
        salt: new Uint8Array(16),
        iterations: 100000,
        hash: 'SHA-256',
    }, baseKey, { name: 'AES-GCM', length: 256 }, true, ['encrypt', 'decrypt']);
    // Extract the private key from the WebAuthn credential
    // const cosePublicKey = publicKeyCredential.response.getPublicKey();
    // if (!cosePublicKey) {
    //     throw new Error('Failed to get COSE public key from WebAuthn credential');
    // }
    // TODO: Extract the private key from the COSE public key and WebAuthn credential
    // Encrypt the private key with the derived symmetric key
    const encryptedPrivateKey = await crypto.subtle.encrypt({
        name: 'AES-GCM',
        iv: new Uint8Array(12),
    }, derivedKey, new Uint8Array( /* Private key bytes */));
    // Import the encrypted private key as a CryptoKey object
    const privateKey = await crypto.subtle.importKey('pkcs8', encryptedPrivateKey, { name: 'ECDSA', namedCurve: 'P-256' }, true, ['sign']);
    return privateKey;
}
exports.derivePrivateKeyFromWebAuthnCredentialAndPin = derivePrivateKeyFromWebAuthnCredentialAndPin;
