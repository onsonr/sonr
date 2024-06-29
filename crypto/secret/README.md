# Zero-Knowledge, Post-Quantum Hybrid Encryption Scheme

`Secret`` is a cutting-edge encryption scheme developed by Sonr, combining zero-knowledge proofs, post-quantum cryptography, and hybrid encryption techniques. This innovative approach provides robust security for the decentralized identity ecosystem, ensuring data privacy and integrity in the face of current and future threats.
Features

Zero-Knowledge Proofs: Utilizes cryptographic accumulators for efficient membership verification without revealing sensitive information.
Post-Quantum Security: Employs Kyber768, a lattice-based key encapsulation mechanism (KEM) resistant to quantum computing attacks.
Hybrid Encryption: Combines the strengths of asymmetric and symmetric cryptography for optimal performance and security.
IPFS Integration: Seamlessly works with IPFS (InterPlanetary File System) for decentralized data storage.
Deterministic Key Derivation: Ensures consistent key generation based on accumulator state and IPFS vault CID.

## How It Works

### Encryption 

1. Marshals the cryptographic accumulator.
2. Derives a Kyber keypair using the accumulator and IPFS vault CID.
3. Encapsulates a shared secret using Kyber768.
4. Encrypts the message using AES-GCM with the shared secret.
5. Prepends the marshaled accumulator to the encrypted data.


### Decryption

1. Extracts and unmarshals the accumulator from the encrypted data.
2. Derives the Kyber keypair using the extracted accumulator.
3. Decapsulates the shared secret.
4. Decrypts the message using AES-GCM.
5. Verifies the witness against the extracted accumulator.


## Security Considerations

- Post-quantum secure due to the use of Kyber768.
- Zero-knowledge proofs protect sensitive information during verification.
- Hybrid approach combines the security of asymmetric cryptography with the efficiency of symmetric encryption.
- Accumulator-based access control adds an extra layer of security.

## Use Cases

- Secure data sharing in decentralized identity systems.
- Privacy-preserving credential verification.
- Quantum-resistant communication for long-term data protection.
- Decentralized access control for sensitive information.
