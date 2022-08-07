---
title: Identifiers
id: identifiers
displayed_sidebar: basicsSidebar
---
# Identifiers
Sonr uses the handshake protocol to register subdomains on the .snr/ TLD -- Top-Level-Domain. Using a .snr/ instead of a traditional username was an idea that came to fruition when we realized how much we disliked passwords and not having a genuine identity on the internet. Right now 90% of all social logins on the internet are proxied through Facebook and Google. We believe that ownership of data should reside in the people, not the services that facilitate it.

### Decentralized Identifiers (DID)


A DID is a new kind of ID that anyone can use to prove they own their domain. It's similar to a website address, but instead of pointing to a website, it points to you. So if you want to prove that you own a name or an email address, you can get a DID and use it to prove that you own that name or email address.

Notable data on each service of the Sonr Protocol is referenced using DID's, making them accessible through your .snr/.

### Base DID

Our base DID (or decentralized identifier) follows a syntactic structure of the **root** (your DID), followed by a **method** (or in our case, in every case with this SDK, Sonr), then followed by a **public key**

![](https://archbee-image-uploads.s3.amazonaws.com/YigsjtwFFq_eX7dhChoeN/ze9buUbapxPP7S5ROVXn__6e60b2d-screenshot2022-03-10at25108pm.png)

### Controllers

**Structure**
Controllers are defined by a **base DID** and a fragment

**Account Devices**
Linking new devices adds new entries to the users' **base DID**

**IPFS Vault**
Users' secure IPFS Vault and Mailbox are also listed as controllers on the account's **DID Document**

**Key Management**
Individual and additional Keys can also be referenced using this fragmented structure.

### Verification

| **Key**          | **Description**                                                                                                                             |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| Authentication   | Service used Key.                                                                                                                           |
| Assertion Method | Verification key by the user.                                                                                                               |
| id               | DID of the verification method                                                                                                              |
| controller       | Pointer to the device holding the key.                                                                                                      |
| type             | The Verification method type from W3:*   `JsonWebkey2020`&#xA;`EcdsaSecp256k1RecoveryMethod2020`&#xA;`EcdsaSecp256k1VerificationKey2019`    |
| PublicKey        | Key format value. The following fields are present:*   `publicKeyJwk`

*   `publicKeyBase58`

*   `publicKeyHex`

*   `blockchainAccountId` |

### Service DID

Apps powered by Sonr will be added to the **Blockchain Registry** with their corresponding **Configuration Data**.

![](https://archbee-image-uploads.s3.amazonaws.com/YigsjtwFFq_eX7dhChoeN/ZW_uX07qd7Er8odd-Dtkh_1a3639d-screenshot2022-03-10at31119pm.png)

### Service Endpoints - (Buckets, Channels, Objects, Functions)

Modules are resolvable to authorized accounts & services through a subpath to the service extension.

![](https://archbee-image-uploads.s3.amazonaws.com/YigsjtwFFq_eX7dhChoeN/9nerqZTJR7h2HT0y2uVUR_712ad7f-screenshot2022-03-10at31530pm-1.png)

### Full Example

Here is a constructed example of a `WebAuthn`** **token being stored in an accounts `DIDDocument`

```json
{
    "@context": "https://www.w3.org/ns/did/v1",
    "assertionMethod": ["did:snr:123#key-1"],
    "alsoKnownAs": ["prad.snr"],
    "controller": "did:snr:123",
    "id": "did:snr:123",
    "verificationMethod": [
    {
      "controller": "did:snr:123",
      "id": "did:snr:123#key-1",
      "publicKeyJwk": {
        "crv": "P-256",
        "kty": "EC",
        "x": "UANQ8pgvJT33JbrnwMiu1L1JCGQFOEm1ThaNAJcFrWA=",
        "y": "UWm6q5n1iXyeCJLMGDInN40bkkKr8KkoTWDqJBZQXRo="
      },
      "type": "JsonWebKey2020"
    }
  ]
}

```


