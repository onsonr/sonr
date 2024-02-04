---
description: >-
  With the increasing concerns around data privacy and security, it is crucial
  to develop a solution that ensures the confidentiality of user information
  while providing a seamless and user-friendly exp
---

# ADR-004: User Session Authorization

## Context

With the increasing concerns around data privacy and security, it is crucial to develop a solution that ensures the confidentiality of user information while providing a seamless and user-friendly experience.

---

## O**bjective**

- Simplify Login experience with device native biometrics
- Compute Graph based relations for Zero-Knowledge elements and DIDs to ensure privacy
- Enable a truly decentralized system at the point of device, to service, to blockchain node.

---

## Solution

1 Sentence elevator pitch for this proposal explaining the improved changes.

#### Prevent User Session Tracking without Consent

More on detail number one.

#### Facilitate Direct Authentication Flow

More on detail number two.

The proposed solution for encrypted and anonymous PassKeys involves two user journeys: registering new users and authenticating existing users. During registration, users provide necessary information and generate a secure Webauthn credential anonymously linked to their account.

Basic profile information can be stored as encrypted data. For authentication, the Webauthn credential is auto-filled based on a registered session, providing a seamless experience. The service owner covers the authentication fee, and successful authentication results in a record stored on the chain for auditing and added security.

---

## Definitions

- `JSON Web Token (JWT)`

  A JSON Web Token (JWT) is a compact, URL-safe means of representing claims between two parties. It is commonly used for authentication and authorization purposes in web applications. A JWT consists of three parts: a header, a payload, and a signature, which are encoded and digitally signed to ensure integrity and authenticity of the token.

- `PassKeys`

  **PassKeys** are secure and anonymous credentials generated during user registration for decentralized identity systems. These credentials, based on WebAuthn technology, are linked to the user's account while preserving their privacy. PassKeys enable seamless authentication by auto-filling the credentials during login, providing a user-friendly experience while ensuring the confidentiality and integrity of user information.

- `PublicKeyAttestation`

  **PublicKeyAttestation** in the context of WebAuthn refers to the process of verifying the authenticity and integrity of a public key credential during registration. It involves generating an attestation statement by the authenticator (e.g., hardware security key) to prove that it possesses a private key corresponding to the public key being registered. This attestation statement is used to ensure that the public key being registered is genuine and has not been tampered with, enhancing the security and trustworthiness of the authentication process.

- `PublicKeyAssertion`

  **PublicKeyAssertion** is a term used in WebAuthn, which refers to a cryptographic proof provided by a user during the authentication process. It is a mechanism that allows the user to prove their possession of a private key corresponding to a registered public key. This assertion is used to verify the user's identity and grant access to requested resources or services without the need for traditional passwords.

- `WebAuthn`

  **WebAuthn** is a web standard that enables strong and passwordless authentication. It allows users to authenticate to online services using public-key cryptography and biometric factors like fingerprints or facial recognition. By eliminating the need for traditional passwords, **WebAuthn** provides a more secure and convenient authentication experience while protecting user privacy.

- `Zero-Knowledge Accumulator`

  A **zero-knowledge accumulator** is a cryptographic primitive that allows for the efficient verification of membership in a set without revealing the individual elements of the set. It enables a prover to demonstrate knowledge of elements in a set without disclosing those elements. This provides a way to prove possession or knowledge of certain data without revealing any specific details about that data, ensuring privacy and confidentiality.

## Sequence Methods

### 1. Registering New Users

If a valid registration request is sent, the token reward is minted for all parties of the Account Generation exchange.

- Options and Minimum accepted identifiers queried from Services
- User generates a Webauthn credential which is anonymously points to the users account
- User can store basic profile info as encrypted data when requested by services

### 2. Authenticating Existing Users

Upon successful authentication a New record of the event is stored on chain with an encrypted fingerprint and the account which was authenticated. The Service owner is responsible for paying the _Authentication Fee_.

- Client auto-fills webauthn credential based off previously registered session.
- Manager of the Service Record requesting authentication pays standard fees.

---

## Economic Impact

### Network Fees

<table><thead><tr><th></th><th width="115">Fees</th><th width="172">Sender</th><th>Receiver</th></tr></thead><tbody><tr><td>Lookup User Identifier existence in Zero-knowledge Accumulator</td><td>SNR 0.5</td><td>Service Owner</td><td>Validator</td></tr><tr><td>Verify a Signature</td><td>SNR 0.1</td><td>Service Owner</td><td>Validator</td></tr></tbody></table>

### Staking Incentives

| Action                   | Minimum Delegation | Unlock Period |
| ------------------------ | ------------------ | ------------- |
| Persisting a Username    | USNR 200,000,000   | 30 Days       |
| Elevate Developer Access | USNR 500,000,000   | 12 Months     |

---

## Implementation

### API Methods

| Summary                     | Method | Endpoint        |
| --------------------------- | ------ | --------------- |
| Query for an Identity       | GET    | /core/identity/ |
| Send a transaction on-chain | POST   | /core/identity/ |

---

## Status

This proposal is **under development** by the core Sonr Team.

| Development Phase | Devnet  |
| ----------------- | ------- |
| Target Completion | Q4 2023 |

---

## References

- [Sonr ADR-001: Decentralized Identity System](https://www.notion.so/ADR-002-Decentralized-Identity-Specification-01102d0fa712448b8893fe1bdc689d1e?pvs=21)
- [Sonr ADR-003: Decentralized Network-Relaying Services](https://www.notion.so/ADR-003-Authoritative-Application-Records-9b579f508d14454bbe995c9dc430c345?pvs=21)
