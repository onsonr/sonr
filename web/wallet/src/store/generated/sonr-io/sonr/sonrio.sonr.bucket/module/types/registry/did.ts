/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sonrio.sonr.registry";

/** NetworkType is the type of network the DID is on. */
export enum NetworkType {
  /** NETWORK_TYPE_UNSPECIFIED - Unspecified is the default value. Gets converted to "did:sonr:". */
  NETWORK_TYPE_UNSPECIFIED = 0,
  /** NETWORK_TYPE_MAINNET - Mainnet is the main network. It prefix is "did:sonr:" or "did:sonr:mainnet:". */
  NETWORK_TYPE_MAINNET = 1,
  /** NETWORK_TYPE_TESTNET - Testnet is the deployed test network. It's prefix is "did:sonr:testnet:". */
  NETWORK_TYPE_TESTNET = 2,
  /** NETWORK_TYPE_DEVNET - Devnet is the localhost test network. It's prefix is "did:sonr:devnet:". */
  NETWORK_TYPE_DEVNET = 3,
  UNRECOGNIZED = -1,
}

export function networkTypeFromJSON(object: any): NetworkType {
  switch (object) {
    case 0:
    case "NETWORK_TYPE_UNSPECIFIED":
      return NetworkType.NETWORK_TYPE_UNSPECIFIED;
    case 1:
    case "NETWORK_TYPE_MAINNET":
      return NetworkType.NETWORK_TYPE_MAINNET;
    case 2:
    case "NETWORK_TYPE_TESTNET":
      return NetworkType.NETWORK_TYPE_TESTNET;
    case 3:
    case "NETWORK_TYPE_DEVNET":
      return NetworkType.NETWORK_TYPE_DEVNET;
    case -1:
    case "UNRECOGNIZED":
    default:
      return NetworkType.UNRECOGNIZED;
  }
}

export function networkTypeToJSON(object: NetworkType): string {
  switch (object) {
    case NetworkType.NETWORK_TYPE_UNSPECIFIED:
      return "NETWORK_TYPE_UNSPECIFIED";
    case NetworkType.NETWORK_TYPE_MAINNET:
      return "NETWORK_TYPE_MAINNET";
    case NetworkType.NETWORK_TYPE_TESTNET:
      return "NETWORK_TYPE_TESTNET";
    case NetworkType.NETWORK_TYPE_DEVNET:
      return "NETWORK_TYPE_DEVNET";
    default:
      return "UNKNOWN";
  }
}

/** ServiceProtocol are core modules that can be installed on custom services on the Sonr network. */
export enum ServiceProtocol {
  /** SERVICE_PROTOCOL_UNSPECIFIED - SERVICE_PROTOCOL_UNSPECIFIED is the default value. */
  SERVICE_PROTOCOL_UNSPECIFIED = 0,
  /** SERVICE_PROTOCOL_BUCKETS - SERVICE_PROTOCOL_BUCKETS is the module that provides the ability to store and retrieve data. */
  SERVICE_PROTOCOL_BUCKETS = 1,
  /** SERVICE_PROTOCOL_CHANNEL - SERVICE_PROTOCOL_CHANNEL is the module that provides the ability to communicate with other services. */
  SERVICE_PROTOCOL_CHANNEL = 2,
  /** SERVICE_PROTOCOL_OBJECTS - SERVICE_PROTOCOL_OBJECTS is the module that provides the ability to create new schemas for data on the network. */
  SERVICE_PROTOCOL_OBJECTS = 3,
  /** SERVICE_PROTOCOL_FUNCTIONS - SERVICE_PROTOCOL_FUNCTIONS is the module that provides the ability to create new functions for data on the network. */
  SERVICE_PROTOCOL_FUNCTIONS = 4,
  UNRECOGNIZED = -1,
}

export function serviceProtocolFromJSON(object: any): ServiceProtocol {
  switch (object) {
    case 0:
    case "SERVICE_PROTOCOL_UNSPECIFIED":
      return ServiceProtocol.SERVICE_PROTOCOL_UNSPECIFIED;
    case 1:
    case "SERVICE_PROTOCOL_BUCKETS":
      return ServiceProtocol.SERVICE_PROTOCOL_BUCKETS;
    case 2:
    case "SERVICE_PROTOCOL_CHANNEL":
      return ServiceProtocol.SERVICE_PROTOCOL_CHANNEL;
    case 3:
    case "SERVICE_PROTOCOL_OBJECTS":
      return ServiceProtocol.SERVICE_PROTOCOL_OBJECTS;
    case 4:
    case "SERVICE_PROTOCOL_FUNCTIONS":
      return ServiceProtocol.SERVICE_PROTOCOL_FUNCTIONS;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ServiceProtocol.UNRECOGNIZED;
  }
}

export function serviceProtocolToJSON(object: ServiceProtocol): string {
  switch (object) {
    case ServiceProtocol.SERVICE_PROTOCOL_UNSPECIFIED:
      return "SERVICE_PROTOCOL_UNSPECIFIED";
    case ServiceProtocol.SERVICE_PROTOCOL_BUCKETS:
      return "SERVICE_PROTOCOL_BUCKETS";
    case ServiceProtocol.SERVICE_PROTOCOL_CHANNEL:
      return "SERVICE_PROTOCOL_CHANNEL";
    case ServiceProtocol.SERVICE_PROTOCOL_OBJECTS:
      return "SERVICE_PROTOCOL_OBJECTS";
    case ServiceProtocol.SERVICE_PROTOCOL_FUNCTIONS:
      return "SERVICE_PROTOCOL_FUNCTIONS";
    default:
      return "UNKNOWN";
  }
}

/** ServiceType is the type of service that is being registered. */
export enum ServiceType {
  /** SERVICE_TYPE_UNSPECIFIED - SERVICE_TYPE_UNSPECIFIED is the default value. */
  SERVICE_TYPE_UNSPECIFIED = 0,
  /** SERVICE_TYPE_DID_COMM_MESSAGING - SERVICE_TYPE_APPLICATION is the type of service that is a DApp. */
  SERVICE_TYPE_DID_COMM_MESSAGING = 1,
  /** SERVICE_TYPE_LINKED_DOMAINS - SERVICE_TYPE_SERVICE is the type of service that is a service. */
  SERVICE_TYPE_LINKED_DOMAINS = 2,
  /** SERVICE_TYPE_SONR - SERVICE_TYPE_SONR is the type of service that is a DApp. */
  SERVICE_TYPE_SONR = 3,
  UNRECOGNIZED = -1,
}

export function serviceTypeFromJSON(object: any): ServiceType {
  switch (object) {
    case 0:
    case "SERVICE_TYPE_UNSPECIFIED":
      return ServiceType.SERVICE_TYPE_UNSPECIFIED;
    case 1:
    case "SERVICE_TYPE_DID_COMM_MESSAGING":
      return ServiceType.SERVICE_TYPE_DID_COMM_MESSAGING;
    case 2:
    case "SERVICE_TYPE_LINKED_DOMAINS":
      return ServiceType.SERVICE_TYPE_LINKED_DOMAINS;
    case 3:
    case "SERVICE_TYPE_SONR":
      return ServiceType.SERVICE_TYPE_SONR;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ServiceType.UNRECOGNIZED;
  }
}

export function serviceTypeToJSON(object: ServiceType): string {
  switch (object) {
    case ServiceType.SERVICE_TYPE_UNSPECIFIED:
      return "SERVICE_TYPE_UNSPECIFIED";
    case ServiceType.SERVICE_TYPE_DID_COMM_MESSAGING:
      return "SERVICE_TYPE_DID_COMM_MESSAGING";
    case ServiceType.SERVICE_TYPE_LINKED_DOMAINS:
      return "SERVICE_TYPE_LINKED_DOMAINS";
    case ServiceType.SERVICE_TYPE_SONR:
      return "SERVICE_TYPE_SONR";
    default:
      return "UNKNOWN";
  }
}

/**
 * Did represents a string that has been parsed and validated as a DID. The parts are stored
 * in the individual fields.
 */
export interface Did {
  /** Method is the method used to create the DID. For the Sonr network it is "sonr". */
  method: string;
  /** Network is the network the DID is on. For testnet it is "testnet". i.e "did:sonr:testnet:". */
  network: string;
  /** id is the trailing identifier after the network. i.e. "did:sonr:testnet:abc123" */
  id: string;
  /** Paths is a list of paths that the DID is valid for. This is used to identify the Service. */
  paths: string[];
  /** Query is the query string that was used to create the DID. This is followed by a '?'. */
  query: string;
  /** Fragment is the fragment string that was used to create the DID. This is followed by a '#'. */
  fragment: string;
}

/** DidDocument is the document that describes a DID. This document is stored on the blockchain. */
export interface DidDocument {
  /** Context is the context of the DID document. This is used to identify the Service. */
  context: string[];
  /** Id is the DID of the document. */
  id: string;
  /** Controller is the DID of the controller of the document. This will be the individual user devices and mailboxes. */
  controller: string[];
  /** VerificationMethod is the list of verification methods for the user. */
  verificationMethod: VerificationMethod[];
  /** Authentication is the list of authentication methods for the user. */
  authentication: string[];
  /** AssertionMethod is the list of assertion methods for the user. */
  assertionMethod: string[];
  /** CapabilityInvocation is the list of capability invocation methods for the user. */
  capabilityInvocation: string[];
  /** CapabilityDelegation is the list of capability delegation methods for the user. */
  capabilityDelegation: string[];
  /** KeyAgreement is the list of key agreement methods for the user. */
  keyAgreement: string[];
  /** Service is the list of services or DApps that the user has access to. */
  service: Service[];
  /** AlsoKnownAs is the list of ".snr" aliases for the user. */
  alsoKnownAs: string[];
  /** Metadata is the metadata of the service. */
  metadata: { [key: string]: string };
}

export interface DidDocument_MetadataEntry {
  key: string;
  value: string;
}

/** Service is a Application that runs on the Sonr network. */
export interface Service {
  /** ID is the DID of the service. */
  id: string;
  /** Type is the type of the service. */
  type: ServiceType;
  /** ServiceEndpoint is the endpoint of the service. */
  serviceEndpoint: ServiceEndpoint | undefined;
  /** Metadata is the metadata of the service. */
  metadata: { [key: string]: string };
}

export interface Service_MetadataEntry {
  key: string;
  value: string;
}

/** ServiceEndpoint is the endpoint of the service. */
export interface ServiceEndpoint {
  /** TransportType is the type of transport used to connect to the service. */
  transportType: string;
  /** Network is the network the service is on. */
  network: string;
  /**
   * SupportedProtocols is the list of protocols supported by the service.
   * (e.g. "channels", "buckets", "objects", "storage")
   */
  supportedProtocols: ServiceProtocol[];
}

/** VerificationMethod is a method that can be used to verify the DID. */
export interface VerificationMethod {
  /** ID is the DID of the verification method. */
  id: string;
  /** Type is the type of the verification method. */
  type: VerificationMethod_Type;
  /** Controller is the DID of the controller of the verification method. */
  controller: string;
  /** PublicKeyHex is the public key of the verification method in hexidecimal. */
  publicKeyHex: string;
  /** PublicKeyBase58 is the public key of the verification method in base58. */
  publicKeyBase58: string;
  /** BlockchainAccountId is the blockchain account id of the verification method. */
  blockchainAccountId: string;
}

export enum VerificationMethod_Type {
  /** TYPE_UNSPECIFIED - TYPE_UNSPECIFIED is the default value. */
  TYPE_UNSPECIFIED = 0,
  /** TYPE_ECDSA_SECP256K1 - TYPE_ECDSA_SECP256K1 represents the Ed25519VerificationKey2018 key type. */
  TYPE_ECDSA_SECP256K1 = 1,
  /** TYPE_X25519 - TYPE_X25519 represents the X25519KeyAgreementKey2019 key type. */
  TYPE_X25519 = 2,
  /** TYPE_ED25519 - TYPE_ED25519 represents the Ed25519VerificationKey2018 key type. */
  TYPE_ED25519 = 3,
  /** TYPE_BLS_12381_G1 - TYPE_BLS_12381_G1 represents the Bls12381G1Key2020 key type */
  TYPE_BLS_12381_G1 = 4,
  /** TYPE_BLS_12381_G2 - TYPE_BLS_12381_G2 represents the Bls12381G2Key2020 key type */
  TYPE_BLS_12381_G2 = 5,
  /** TYPE_RSA - TYPE_RSA represents the RsaVerificationKey2018 key type. */
  TYPE_RSA = 6,
  /** TYPE_VERIFIABLE_CONDITION - TYPE_VERIFIABLE_CONDITION represents the VerifiableCondition2021 key type. */
  TYPE_VERIFIABLE_CONDITION = 7,
  UNRECOGNIZED = -1,
}

export function verificationMethod_TypeFromJSON(
  object: any
): VerificationMethod_Type {
  switch (object) {
    case 0:
    case "TYPE_UNSPECIFIED":
      return VerificationMethod_Type.TYPE_UNSPECIFIED;
    case 1:
    case "TYPE_ECDSA_SECP256K1":
      return VerificationMethod_Type.TYPE_ECDSA_SECP256K1;
    case 2:
    case "TYPE_X25519":
      return VerificationMethod_Type.TYPE_X25519;
    case 3:
    case "TYPE_ED25519":
      return VerificationMethod_Type.TYPE_ED25519;
    case 4:
    case "TYPE_BLS_12381_G1":
      return VerificationMethod_Type.TYPE_BLS_12381_G1;
    case 5:
    case "TYPE_BLS_12381_G2":
      return VerificationMethod_Type.TYPE_BLS_12381_G2;
    case 6:
    case "TYPE_RSA":
      return VerificationMethod_Type.TYPE_RSA;
    case 7:
    case "TYPE_VERIFIABLE_CONDITION":
      return VerificationMethod_Type.TYPE_VERIFIABLE_CONDITION;
    case -1:
    case "UNRECOGNIZED":
    default:
      return VerificationMethod_Type.UNRECOGNIZED;
  }
}

export function verificationMethod_TypeToJSON(
  object: VerificationMethod_Type
): string {
  switch (object) {
    case VerificationMethod_Type.TYPE_UNSPECIFIED:
      return "TYPE_UNSPECIFIED";
    case VerificationMethod_Type.TYPE_ECDSA_SECP256K1:
      return "TYPE_ECDSA_SECP256K1";
    case VerificationMethod_Type.TYPE_X25519:
      return "TYPE_X25519";
    case VerificationMethod_Type.TYPE_ED25519:
      return "TYPE_ED25519";
    case VerificationMethod_Type.TYPE_BLS_12381_G1:
      return "TYPE_BLS_12381_G1";
    case VerificationMethod_Type.TYPE_BLS_12381_G2:
      return "TYPE_BLS_12381_G2";
    case VerificationMethod_Type.TYPE_RSA:
      return "TYPE_RSA";
    case VerificationMethod_Type.TYPE_VERIFIABLE_CONDITION:
      return "TYPE_VERIFIABLE_CONDITION";
    default:
      return "UNKNOWN";
  }
}

const baseDid: object = {
  method: "",
  network: "",
  id: "",
  paths: "",
  query: "",
  fragment: "",
};

export const Did = {
  encode(message: Did, writer: Writer = Writer.create()): Writer {
    if (message.method !== "") {
      writer.uint32(10).string(message.method);
    }
    if (message.network !== "") {
      writer.uint32(18).string(message.network);
    }
    if (message.id !== "") {
      writer.uint32(26).string(message.id);
    }
    for (const v of message.paths) {
      writer.uint32(34).string(v!);
    }
    if (message.query !== "") {
      writer.uint32(42).string(message.query);
    }
    if (message.fragment !== "") {
      writer.uint32(50).string(message.fragment);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Did {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDid } as Did;
    message.paths = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.method = reader.string();
          break;
        case 2:
          message.network = reader.string();
          break;
        case 3:
          message.id = reader.string();
          break;
        case 4:
          message.paths.push(reader.string());
          break;
        case 5:
          message.query = reader.string();
          break;
        case 6:
          message.fragment = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Did {
    const message = { ...baseDid } as Did;
    message.paths = [];
    if (object.method !== undefined && object.method !== null) {
      message.method = String(object.method);
    } else {
      message.method = "";
    }
    if (object.network !== undefined && object.network !== null) {
      message.network = String(object.network);
    } else {
      message.network = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.paths !== undefined && object.paths !== null) {
      for (const e of object.paths) {
        message.paths.push(String(e));
      }
    }
    if (object.query !== undefined && object.query !== null) {
      message.query = String(object.query);
    } else {
      message.query = "";
    }
    if (object.fragment !== undefined && object.fragment !== null) {
      message.fragment = String(object.fragment);
    } else {
      message.fragment = "";
    }
    return message;
  },

  toJSON(message: Did): unknown {
    const obj: any = {};
    message.method !== undefined && (obj.method = message.method);
    message.network !== undefined && (obj.network = message.network);
    message.id !== undefined && (obj.id = message.id);
    if (message.paths) {
      obj.paths = message.paths.map((e) => e);
    } else {
      obj.paths = [];
    }
    message.query !== undefined && (obj.query = message.query);
    message.fragment !== undefined && (obj.fragment = message.fragment);
    return obj;
  },

  fromPartial(object: DeepPartial<Did>): Did {
    const message = { ...baseDid } as Did;
    message.paths = [];
    if (object.method !== undefined && object.method !== null) {
      message.method = object.method;
    } else {
      message.method = "";
    }
    if (object.network !== undefined && object.network !== null) {
      message.network = object.network;
    } else {
      message.network = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.paths !== undefined && object.paths !== null) {
      for (const e of object.paths) {
        message.paths.push(e);
      }
    }
    if (object.query !== undefined && object.query !== null) {
      message.query = object.query;
    } else {
      message.query = "";
    }
    if (object.fragment !== undefined && object.fragment !== null) {
      message.fragment = object.fragment;
    } else {
      message.fragment = "";
    }
    return message;
  },
};

const baseDidDocument: object = {
  context: "",
  id: "",
  controller: "",
  authentication: "",
  assertionMethod: "",
  capabilityInvocation: "",
  capabilityDelegation: "",
  keyAgreement: "",
  alsoKnownAs: "",
};

export const DidDocument = {
  encode(message: DidDocument, writer: Writer = Writer.create()): Writer {
    for (const v of message.context) {
      writer.uint32(10).string(v!);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    for (const v of message.controller) {
      writer.uint32(26).string(v!);
    }
    for (const v of message.verificationMethod) {
      VerificationMethod.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.authentication) {
      writer.uint32(42).string(v!);
    }
    for (const v of message.assertionMethod) {
      writer.uint32(50).string(v!);
    }
    for (const v of message.capabilityInvocation) {
      writer.uint32(58).string(v!);
    }
    for (const v of message.capabilityDelegation) {
      writer.uint32(66).string(v!);
    }
    for (const v of message.keyAgreement) {
      writer.uint32(74).string(v!);
    }
    for (const v of message.service) {
      Service.encode(v!, writer.uint32(82).fork()).ldelim();
    }
    for (const v of message.alsoKnownAs) {
      writer.uint32(90).string(v!);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      DidDocument_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(98).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DidDocument {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDidDocument } as DidDocument;
    message.context = [];
    message.controller = [];
    message.verificationMethod = [];
    message.authentication = [];
    message.assertionMethod = [];
    message.capabilityInvocation = [];
    message.capabilityDelegation = [];
    message.keyAgreement = [];
    message.service = [];
    message.alsoKnownAs = [];
    message.metadata = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.context.push(reader.string());
          break;
        case 2:
          message.id = reader.string();
          break;
        case 3:
          message.controller.push(reader.string());
          break;
        case 4:
          message.verificationMethod.push(
            VerificationMethod.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.authentication.push(reader.string());
          break;
        case 6:
          message.assertionMethod.push(reader.string());
          break;
        case 7:
          message.capabilityInvocation.push(reader.string());
          break;
        case 8:
          message.capabilityDelegation.push(reader.string());
          break;
        case 9:
          message.keyAgreement.push(reader.string());
          break;
        case 10:
          message.service.push(Service.decode(reader, reader.uint32()));
          break;
        case 11:
          message.alsoKnownAs.push(reader.string());
          break;
        case 12:
          const entry12 = DidDocument_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry12.value !== undefined) {
            message.metadata[entry12.key] = entry12.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DidDocument {
    const message = { ...baseDidDocument } as DidDocument;
    message.context = [];
    message.controller = [];
    message.verificationMethod = [];
    message.authentication = [];
    message.assertionMethod = [];
    message.capabilityInvocation = [];
    message.capabilityDelegation = [];
    message.keyAgreement = [];
    message.service = [];
    message.alsoKnownAs = [];
    message.metadata = {};
    if (object.context !== undefined && object.context !== null) {
      for (const e of object.context) {
        message.context.push(String(e));
      }
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.controller !== undefined && object.controller !== null) {
      for (const e of object.controller) {
        message.controller.push(String(e));
      }
    }
    if (
      object.verificationMethod !== undefined &&
      object.verificationMethod !== null
    ) {
      for (const e of object.verificationMethod) {
        message.verificationMethod.push(VerificationMethod.fromJSON(e));
      }
    }
    if (object.authentication !== undefined && object.authentication !== null) {
      for (const e of object.authentication) {
        message.authentication.push(String(e));
      }
    }
    if (
      object.assertionMethod !== undefined &&
      object.assertionMethod !== null
    ) {
      for (const e of object.assertionMethod) {
        message.assertionMethod.push(String(e));
      }
    }
    if (
      object.capabilityInvocation !== undefined &&
      object.capabilityInvocation !== null
    ) {
      for (const e of object.capabilityInvocation) {
        message.capabilityInvocation.push(String(e));
      }
    }
    if (
      object.capabilityDelegation !== undefined &&
      object.capabilityDelegation !== null
    ) {
      for (const e of object.capabilityDelegation) {
        message.capabilityDelegation.push(String(e));
      }
    }
    if (object.keyAgreement !== undefined && object.keyAgreement !== null) {
      for (const e of object.keyAgreement) {
        message.keyAgreement.push(String(e));
      }
    }
    if (object.service !== undefined && object.service !== null) {
      for (const e of object.service) {
        message.service.push(Service.fromJSON(e));
      }
    }
    if (object.alsoKnownAs !== undefined && object.alsoKnownAs !== null) {
      for (const e of object.alsoKnownAs) {
        message.alsoKnownAs.push(String(e));
      }
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    return message;
  },

  toJSON(message: DidDocument): unknown {
    const obj: any = {};
    if (message.context) {
      obj.context = message.context.map((e) => e);
    } else {
      obj.context = [];
    }
    message.id !== undefined && (obj.id = message.id);
    if (message.controller) {
      obj.controller = message.controller.map((e) => e);
    } else {
      obj.controller = [];
    }
    if (message.verificationMethod) {
      obj.verificationMethod = message.verificationMethod.map((e) =>
        e ? VerificationMethod.toJSON(e) : undefined
      );
    } else {
      obj.verificationMethod = [];
    }
    if (message.authentication) {
      obj.authentication = message.authentication.map((e) => e);
    } else {
      obj.authentication = [];
    }
    if (message.assertionMethod) {
      obj.assertionMethod = message.assertionMethod.map((e) => e);
    } else {
      obj.assertionMethod = [];
    }
    if (message.capabilityInvocation) {
      obj.capabilityInvocation = message.capabilityInvocation.map((e) => e);
    } else {
      obj.capabilityInvocation = [];
    }
    if (message.capabilityDelegation) {
      obj.capabilityDelegation = message.capabilityDelegation.map((e) => e);
    } else {
      obj.capabilityDelegation = [];
    }
    if (message.keyAgreement) {
      obj.keyAgreement = message.keyAgreement.map((e) => e);
    } else {
      obj.keyAgreement = [];
    }
    if (message.service) {
      obj.service = message.service.map((e) =>
        e ? Service.toJSON(e) : undefined
      );
    } else {
      obj.service = [];
    }
    if (message.alsoKnownAs) {
      obj.alsoKnownAs = message.alsoKnownAs.map((e) => e);
    } else {
      obj.alsoKnownAs = [];
    }
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<DidDocument>): DidDocument {
    const message = { ...baseDidDocument } as DidDocument;
    message.context = [];
    message.controller = [];
    message.verificationMethod = [];
    message.authentication = [];
    message.assertionMethod = [];
    message.capabilityInvocation = [];
    message.capabilityDelegation = [];
    message.keyAgreement = [];
    message.service = [];
    message.alsoKnownAs = [];
    message.metadata = {};
    if (object.context !== undefined && object.context !== null) {
      for (const e of object.context) {
        message.context.push(e);
      }
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.controller !== undefined && object.controller !== null) {
      for (const e of object.controller) {
        message.controller.push(e);
      }
    }
    if (
      object.verificationMethod !== undefined &&
      object.verificationMethod !== null
    ) {
      for (const e of object.verificationMethod) {
        message.verificationMethod.push(VerificationMethod.fromPartial(e));
      }
    }
    if (object.authentication !== undefined && object.authentication !== null) {
      for (const e of object.authentication) {
        message.authentication.push(e);
      }
    }
    if (
      object.assertionMethod !== undefined &&
      object.assertionMethod !== null
    ) {
      for (const e of object.assertionMethod) {
        message.assertionMethod.push(e);
      }
    }
    if (
      object.capabilityInvocation !== undefined &&
      object.capabilityInvocation !== null
    ) {
      for (const e of object.capabilityInvocation) {
        message.capabilityInvocation.push(e);
      }
    }
    if (
      object.capabilityDelegation !== undefined &&
      object.capabilityDelegation !== null
    ) {
      for (const e of object.capabilityDelegation) {
        message.capabilityDelegation.push(e);
      }
    }
    if (object.keyAgreement !== undefined && object.keyAgreement !== null) {
      for (const e of object.keyAgreement) {
        message.keyAgreement.push(e);
      }
    }
    if (object.service !== undefined && object.service !== null) {
      for (const e of object.service) {
        message.service.push(Service.fromPartial(e));
      }
    }
    if (object.alsoKnownAs !== undefined && object.alsoKnownAs !== null) {
      for (const e of object.alsoKnownAs) {
        message.alsoKnownAs.push(e);
      }
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
    }
    return message;
  },
};

const baseDidDocument_MetadataEntry: object = { key: "", value: "" };

export const DidDocument_MetadataEntry = {
  encode(
    message: DidDocument_MetadataEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): DidDocument_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseDidDocument_MetadataEntry,
    } as DidDocument_MetadataEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DidDocument_MetadataEntry {
    const message = {
      ...baseDidDocument_MetadataEntry,
    } as DidDocument_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: DidDocument_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<DidDocument_MetadataEntry>
  ): DidDocument_MetadataEntry {
    const message = {
      ...baseDidDocument_MetadataEntry,
    } as DidDocument_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    return message;
  },
};

const baseService: object = { id: "", type: 0 };

export const Service = {
  encode(message: Service, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.type !== 0) {
      writer.uint32(16).int32(message.type);
    }
    if (message.serviceEndpoint !== undefined) {
      ServiceEndpoint.encode(
        message.serviceEndpoint,
        writer.uint32(26).fork()
      ).ldelim();
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      Service_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Service {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseService } as Service;
    message.metadata = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.type = reader.int32() as any;
          break;
        case 3:
          message.serviceEndpoint = ServiceEndpoint.decode(
            reader,
            reader.uint32()
          );
          break;
        case 4:
          const entry4 = Service_MetadataEntry.decode(reader, reader.uint32());
          if (entry4.value !== undefined) {
            message.metadata[entry4.key] = entry4.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Service {
    const message = { ...baseService } as Service;
    message.metadata = {};
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = serviceTypeFromJSON(object.type);
    } else {
      message.type = 0;
    }
    if (
      object.serviceEndpoint !== undefined &&
      object.serviceEndpoint !== null
    ) {
      message.serviceEndpoint = ServiceEndpoint.fromJSON(
        object.serviceEndpoint
      );
    } else {
      message.serviceEndpoint = undefined;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    return message;
  },

  toJSON(message: Service): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.type !== undefined && (obj.type = serviceTypeToJSON(message.type));
    message.serviceEndpoint !== undefined &&
      (obj.serviceEndpoint = message.serviceEndpoint
        ? ServiceEndpoint.toJSON(message.serviceEndpoint)
        : undefined);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Service>): Service {
    const message = { ...baseService } as Service;
    message.metadata = {};
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = 0;
    }
    if (
      object.serviceEndpoint !== undefined &&
      object.serviceEndpoint !== null
    ) {
      message.serviceEndpoint = ServiceEndpoint.fromPartial(
        object.serviceEndpoint
      );
    } else {
      message.serviceEndpoint = undefined;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
    }
    return message;
  },
};

const baseService_MetadataEntry: object = { key: "", value: "" };

export const Service_MetadataEntry = {
  encode(
    message: Service_MetadataEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Service_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseService_MetadataEntry } as Service_MetadataEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Service_MetadataEntry {
    const message = { ...baseService_MetadataEntry } as Service_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: Service_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<Service_MetadataEntry>
  ): Service_MetadataEntry {
    const message = { ...baseService_MetadataEntry } as Service_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    return message;
  },
};

const baseServiceEndpoint: object = {
  transportType: "",
  network: "",
  supportedProtocols: 0,
};

export const ServiceEndpoint = {
  encode(message: ServiceEndpoint, writer: Writer = Writer.create()): Writer {
    if (message.transportType !== "") {
      writer.uint32(10).string(message.transportType);
    }
    if (message.network !== "") {
      writer.uint32(18).string(message.network);
    }
    writer.uint32(26).fork();
    for (const v of message.supportedProtocols) {
      writer.int32(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ServiceEndpoint {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseServiceEndpoint } as ServiceEndpoint;
    message.supportedProtocols = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.transportType = reader.string();
          break;
        case 2:
          message.network = reader.string();
          break;
        case 3:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.supportedProtocols.push(reader.int32() as any);
            }
          } else {
            message.supportedProtocols.push(reader.int32() as any);
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ServiceEndpoint {
    const message = { ...baseServiceEndpoint } as ServiceEndpoint;
    message.supportedProtocols = [];
    if (object.transportType !== undefined && object.transportType !== null) {
      message.transportType = String(object.transportType);
    } else {
      message.transportType = "";
    }
    if (object.network !== undefined && object.network !== null) {
      message.network = String(object.network);
    } else {
      message.network = "";
    }
    if (
      object.supportedProtocols !== undefined &&
      object.supportedProtocols !== null
    ) {
      for (const e of object.supportedProtocols) {
        message.supportedProtocols.push(serviceProtocolFromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: ServiceEndpoint): unknown {
    const obj: any = {};
    message.transportType !== undefined &&
      (obj.transportType = message.transportType);
    message.network !== undefined && (obj.network = message.network);
    if (message.supportedProtocols) {
      obj.supportedProtocols = message.supportedProtocols.map((e) =>
        serviceProtocolToJSON(e)
      );
    } else {
      obj.supportedProtocols = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ServiceEndpoint>): ServiceEndpoint {
    const message = { ...baseServiceEndpoint } as ServiceEndpoint;
    message.supportedProtocols = [];
    if (object.transportType !== undefined && object.transportType !== null) {
      message.transportType = object.transportType;
    } else {
      message.transportType = "";
    }
    if (object.network !== undefined && object.network !== null) {
      message.network = object.network;
    } else {
      message.network = "";
    }
    if (
      object.supportedProtocols !== undefined &&
      object.supportedProtocols !== null
    ) {
      for (const e of object.supportedProtocols) {
        message.supportedProtocols.push(e);
      }
    }
    return message;
  },
};

const baseVerificationMethod: object = {
  id: "",
  type: 0,
  controller: "",
  publicKeyHex: "",
  publicKeyBase58: "",
  blockchainAccountId: "",
};

export const VerificationMethod = {
  encode(
    message: VerificationMethod,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.type !== 0) {
      writer.uint32(16).int32(message.type);
    }
    if (message.controller !== "") {
      writer.uint32(26).string(message.controller);
    }
    if (message.publicKeyHex !== "") {
      writer.uint32(34).string(message.publicKeyHex);
    }
    if (message.publicKeyBase58 !== "") {
      writer.uint32(42).string(message.publicKeyBase58);
    }
    if (message.blockchainAccountId !== "") {
      writer.uint32(50).string(message.blockchainAccountId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VerificationMethod {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVerificationMethod } as VerificationMethod;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.type = reader.int32() as any;
          break;
        case 3:
          message.controller = reader.string();
          break;
        case 4:
          message.publicKeyHex = reader.string();
          break;
        case 5:
          message.publicKeyBase58 = reader.string();
          break;
        case 6:
          message.blockchainAccountId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VerificationMethod {
    const message = { ...baseVerificationMethod } as VerificationMethod;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = verificationMethod_TypeFromJSON(object.type);
    } else {
      message.type = 0;
    }
    if (object.controller !== undefined && object.controller !== null) {
      message.controller = String(object.controller);
    } else {
      message.controller = "";
    }
    if (object.publicKeyHex !== undefined && object.publicKeyHex !== null) {
      message.publicKeyHex = String(object.publicKeyHex);
    } else {
      message.publicKeyHex = "";
    }
    if (
      object.publicKeyBase58 !== undefined &&
      object.publicKeyBase58 !== null
    ) {
      message.publicKeyBase58 = String(object.publicKeyBase58);
    } else {
      message.publicKeyBase58 = "";
    }
    if (
      object.blockchainAccountId !== undefined &&
      object.blockchainAccountId !== null
    ) {
      message.blockchainAccountId = String(object.blockchainAccountId);
    } else {
      message.blockchainAccountId = "";
    }
    return message;
  },

  toJSON(message: VerificationMethod): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.type !== undefined &&
      (obj.type = verificationMethod_TypeToJSON(message.type));
    message.controller !== undefined && (obj.controller = message.controller);
    message.publicKeyHex !== undefined &&
      (obj.publicKeyHex = message.publicKeyHex);
    message.publicKeyBase58 !== undefined &&
      (obj.publicKeyBase58 = message.publicKeyBase58);
    message.blockchainAccountId !== undefined &&
      (obj.blockchainAccountId = message.blockchainAccountId);
    return obj;
  },

  fromPartial(object: DeepPartial<VerificationMethod>): VerificationMethod {
    const message = { ...baseVerificationMethod } as VerificationMethod;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = 0;
    }
    if (object.controller !== undefined && object.controller !== null) {
      message.controller = object.controller;
    } else {
      message.controller = "";
    }
    if (object.publicKeyHex !== undefined && object.publicKeyHex !== null) {
      message.publicKeyHex = object.publicKeyHex;
    } else {
      message.publicKeyHex = "";
    }
    if (
      object.publicKeyBase58 !== undefined &&
      object.publicKeyBase58 !== null
    ) {
      message.publicKeyBase58 = object.publicKeyBase58;
    } else {
      message.publicKeyBase58 = "";
    }
    if (
      object.blockchainAccountId !== undefined &&
      object.blockchainAccountId !== null
    ) {
      message.blockchainAccountId = object.blockchainAccountId;
    } else {
      message.blockchainAccountId = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
