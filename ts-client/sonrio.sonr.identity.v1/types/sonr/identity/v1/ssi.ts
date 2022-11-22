/* eslint-disable */
export const protobufPackage = "sonrio.sonr.identity.v1";

/** KeyType is the type of key used to sign a DID document. */
export enum KeyType {
  /** KeyType_UNSPECIFIED - No key type specified */
  KeyType_UNSPECIFIED = 0,
  /** KeyType_JSON_WEB_KEY_2020 - JsonWebKey2020 is a VerificationMethod type. https://w3c-ccg.github.io/lds-jws2020/ */
  KeyType_JSON_WEB_KEY_2020 = 1,
  /** KeyType_ED25519_VERIFICATION_KEY_2018 - ED25519VerificationKey2018 is the Ed25519VerificationKey2018 verification key type as specified here: https://w3c-ccg.github.io/lds-ed25519-2018/ */
  KeyType_ED25519_VERIFICATION_KEY_2018 = 2,
  /** KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019 - ECDSASECP256K1VerificationKey2019 is the EcdsaSecp256k1VerificationKey2019 verification key type as specified here: https://w3c-ccg.github.io/lds-ecdsa-secp256k1-2019/ */
  KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019 = 3,
  /** KeyType_RSA_VERIFICATION_KEY_2018 - RSAVerificationKey2018 is the RsaVerificationKey2018 verification key type as specified here: https://w3c-ccg.github.io/lds-rsa2018/ */
  KeyType_RSA_VERIFICATION_KEY_2018 = 4,
  UNRECOGNIZED = -1,
}

export function keyTypeFromJSON(object: any): KeyType {
  switch (object) {
    case 0:
    case "KeyType_UNSPECIFIED":
      return KeyType.KeyType_UNSPECIFIED;
    case 1:
    case "KeyType_JSON_WEB_KEY_2020":
      return KeyType.KeyType_JSON_WEB_KEY_2020;
    case 2:
    case "KeyType_ED25519_VERIFICATION_KEY_2018":
      return KeyType.KeyType_ED25519_VERIFICATION_KEY_2018;
    case 3:
    case "KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019":
      return KeyType.KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019;
    case 4:
    case "KeyType_RSA_VERIFICATION_KEY_2018":
      return KeyType.KeyType_RSA_VERIFICATION_KEY_2018;
    case -1:
    case "UNRECOGNIZED":
    default:
      return KeyType.UNRECOGNIZED;
  }
}

export function keyTypeToJSON(object: KeyType): string {
  switch (object) {
    case KeyType.KeyType_UNSPECIFIED:
      return "KeyType_UNSPECIFIED";
    case KeyType.KeyType_JSON_WEB_KEY_2020:
      return "KeyType_JSON_WEB_KEY_2020";
    case KeyType.KeyType_ED25519_VERIFICATION_KEY_2018:
      return "KeyType_ED25519_VERIFICATION_KEY_2018";
    case KeyType.KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019:
      return "KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019";
    case KeyType.KeyType_RSA_VERIFICATION_KEY_2018:
      return "KeyType_RSA_VERIFICATION_KEY_2018";
    case KeyType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

/** ProofType is the type of proof used to present claims over a DID document. */
export enum ProofType {
  /** ProofType_UNSPECIFIED - No proof type specified */
  ProofType_UNSPECIFIED = 0,
  /** ProofType_JSON_WEB_SIGNATURE_2020 - JsonWebSignature2020 is a proof type. https://w3c-ccg.github.io/lds-jws2020/ */
  ProofType_JSON_WEB_SIGNATURE_2020 = 1,
  /** ProofType_ED25519_SIGNATURE_2018 - ED25519Signature2018 is the Ed25519Signature2018 proof type as specified here: https://w3c-ccg.github.io/lds-ed25519-2018/ */
  ProofType_ED25519_SIGNATURE_2018 = 2,
  /** ProofType_ECDSA_SECP256K1_SIGNATURE_2019 - EcdsaSecp256k1Signature2019 is the EcdsaSecp256k1Signature2019 proof type as specified here: https://w3c-ccg.github.io/lds-ecdsa-secp256k1-2019/ */
  ProofType_ECDSA_SECP256K1_SIGNATURE_2019 = 3,
  /** ProofType_RSA_SIGNATURE_2018 - RsaSignature2018 is the RsaSignature2018 proof type as specified here: https://w3c-ccg.github.io/lds-rsa2018/ */
  ProofType_RSA_SIGNATURE_2018 = 4,
  UNRECOGNIZED = -1,
}

export function proofTypeFromJSON(object: any): ProofType {
  switch (object) {
    case 0:
    case "ProofType_UNSPECIFIED":
      return ProofType.ProofType_UNSPECIFIED;
    case 1:
    case "ProofType_JSON_WEB_SIGNATURE_2020":
      return ProofType.ProofType_JSON_WEB_SIGNATURE_2020;
    case 2:
    case "ProofType_ED25519_SIGNATURE_2018":
      return ProofType.ProofType_ED25519_SIGNATURE_2018;
    case 3:
    case "ProofType_ECDSA_SECP256K1_SIGNATURE_2019":
      return ProofType.ProofType_ECDSA_SECP256K1_SIGNATURE_2019;
    case 4:
    case "ProofType_RSA_SIGNATURE_2018":
      return ProofType.ProofType_RSA_SIGNATURE_2018;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ProofType.UNRECOGNIZED;
  }
}

export function proofTypeToJSON(object: ProofType): string {
  switch (object) {
    case ProofType.ProofType_UNSPECIFIED:
      return "ProofType_UNSPECIFIED";
    case ProofType.ProofType_JSON_WEB_SIGNATURE_2020:
      return "ProofType_JSON_WEB_SIGNATURE_2020";
    case ProofType.ProofType_ED25519_SIGNATURE_2018:
      return "ProofType_ED25519_SIGNATURE_2018";
    case ProofType.ProofType_ECDSA_SECP256K1_SIGNATURE_2019:
      return "ProofType_ECDSA_SECP256K1_SIGNATURE_2019";
    case ProofType.ProofType_RSA_SIGNATURE_2018:
      return "ProofType_RSA_SIGNATURE_2018";
    case ProofType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}
