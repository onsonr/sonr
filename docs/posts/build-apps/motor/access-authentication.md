---
title: Access & Authentication
id: access-authentication
displayed_sidebar: modulesSidebar
---
# Access and Authentication
We think authentication should be simple, yet secure. The Sonr network uses Webauthn a key-based authentication system utilizing credential systems found on the user operating system to perform key exchanges with our Highway Nodes which then grants you access to the network. The following is a diagram outlining our authentication and registration flows.




<!-- [t]("https://www.figma.com/file/4BeBs2QYmytTN0RII1i4d8/Webauthn-flow?node-id=0%3A1") -->





# Platform Credentials

Currently, our implementations of Webauthn use the 'platform-specific' credential options meaning our servers will request your operating system to use whichever authentication method is most native to it. For information on what authentication mechanisims are supported





# DID Registration

When a user registers their domain will be prompted to supply your user credentials for relating to our generated `DID` which is then paired with provided `PublicKeyCredentials`

The folowing is an example of a generated `WhoIs` which repersented a user in our registry



```javascript
{
  "@context": [
    "https://www.w3.org/ns/did/v1"
  ],
  "id": "did:snr:04cf1e20-378a-4e38-ab1b-401a5018c9ff",
  "controller": [
    "did:snr:04cf1e20-378a-4e38-ab1b-401a5018c9ff",
    "did:snr:f03a00f1-9615-4060-bd00-bd282e150c46"
  ],
  "verificationMethod": [
    {
      "id": "did:snr:04cf1e20-378a-4e38-ab1b-401a5018c9ff#key-1",
      "type": "JsonWebKey2020",
      "controller": "did:snr:04cf1e20-378a-4e38-ab1b-401a5018c9ff",
      "publicKeyJwk": {
        "kty": "EC",
        "crv": "P-256",
        "x": "SVqB4JcUD6lsfvqMr-OKUNUphdNn64Eay60978ZlL74",
        "y": "lf0u0pMj4lGAzZix5u4Cm5CMQIgMNpkwy163wtKYVKI"
      }
    },
    {
      "id": "did:snr:04cf1e20-378a-4e38-ab1b-401a5018c9ff#key-2",
      "type": "JsonWebKey2020",
      "controller": "did:snr:04cf1e20-378a-4e38-ab1b-401a5018c9ff",
      "publicKeyJwk": {
        "kty": "EC",
        "crv": "P-256",
        "x": "SVqB4JcUD6lsfvqMr-OKUNUphdNn64Eay60978ZlL74",
        "y": "lf0u0pMj4lGAzZix5u4Cm5CMQIgMNpkwy163wtKYVKI"
      }
    },
    {
      "id": "did:snr:04cf1e20-378a-4e38-ab1b-401a5018c9ff#added-assertion-method-1",
      "controller": "did:snr:04cf1e20-378a-4e38-ab1b-401a5018c9ff",
      "publicKeyBase58": "GGRj8PAR5tRgD5xqAhPna1bLa3UoYuxNEEhRmcYCPBm5",
      "type": "Ed25519VerificationKey2018"
    }
  ],
  "authentication": [
    "did:snr:04cf1e20-378a-4e38-ab1b-401a5018c9ff#key-1",
    "did:snr:04cf1e20-378a-4e38-ab1b-401a5018c9ff#key-2"
  ],
  "assertionMethod": [
    "did:snr:04cf1e20-378a-4e38-ab1b-401a5018c9ff#key-1"
  ],
  "service": [
    {
      "id": "did:snr:04cf1e20-378a-4e38-ab1b-401a5018c9ff#service-1",
      "type": "DIDCommMessaging",
      "serviceEndpoint": "did:snr:<vendor>#service-76"
    },
    {
      "id": "did:snr:04cf1e20-378a-4e38-ab1b-401a5018c9ff#service-2",
      "type": "EncryptedDataVault",
      "serviceEndpoint": "did:snr:<vendor>#service-2"
    }
  ]
}
```



``
