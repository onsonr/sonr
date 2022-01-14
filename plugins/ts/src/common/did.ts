/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "common";

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
}

/** Service is a Application that runs on the Sonr network. */
export interface Service {
  /** ID is the DID of the service. */
  id: string;
  /** Type is the type of the service. */
  type: string;
  /** ServiceEndpoint is the endpoint of the service. */
  serviceEndpoint: string;
}

/** VerificationMethod is a method that can be used to verify the DID. */
export interface VerificationMethod {
  /** ID is the DID of the verification method. */
  id: string;
  /** Type is the type of the verification method. */
  type: string;
  /** Controller is the DID of the controller of the verification method. */
  controller: string;
  /** PublicKeyHex is the public key of the verification method in hexidecimal. */
  publicKeyHex: string;
  /** PublicKeyBase58 is the public key of the verification method in base58. */
  publicKeyBase58: string;
  /** BlockchainAccountId is the blockchain account id of the verification method. */
  blockchainAccountId: string;
}

function createBaseDid(): Did {
  return {
    method: "",
    network: "",
    id: "",
    paths: [],
    query: "",
    fragment: "",
  };
}

export const Did = {
  encode(message: Did, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): Did {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDid();
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
    return {
      method: isSet(object.method) ? String(object.method) : "",
      network: isSet(object.network) ? String(object.network) : "",
      id: isSet(object.id) ? String(object.id) : "",
      paths: Array.isArray(object?.paths)
        ? object.paths.map((e: any) => String(e))
        : [],
      query: isSet(object.query) ? String(object.query) : "",
      fragment: isSet(object.fragment) ? String(object.fragment) : "",
    };
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

  fromPartial<I extends Exact<DeepPartial<Did>, I>>(object: I): Did {
    const message = createBaseDid();
    message.method = object.method ?? "";
    message.network = object.network ?? "";
    message.id = object.id ?? "";
    message.paths = object.paths?.map((e) => e) || [];
    message.query = object.query ?? "";
    message.fragment = object.fragment ?? "";
    return message;
  },
};

function createBaseDidDocument(): DidDocument {
  return {
    context: [],
    id: "",
    controller: [],
    verificationMethod: [],
    authentication: [],
    assertionMethod: [],
    capabilityInvocation: [],
    capabilityDelegation: [],
    keyAgreement: [],
    service: [],
    alsoKnownAs: [],
  };
}

export const DidDocument = {
  encode(
    message: DidDocument,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
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
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DidDocument {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDidDocument();
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
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DidDocument {
    return {
      context: Array.isArray(object?.context)
        ? object.context.map((e: any) => String(e))
        : [],
      id: isSet(object.id) ? String(object.id) : "",
      controller: Array.isArray(object?.controller)
        ? object.controller.map((e: any) => String(e))
        : [],
      verificationMethod: Array.isArray(object?.verificationMethod)
        ? object.verificationMethod.map((e: any) =>
            VerificationMethod.fromJSON(e)
          )
        : [],
      authentication: Array.isArray(object?.authentication)
        ? object.authentication.map((e: any) => String(e))
        : [],
      assertionMethod: Array.isArray(object?.assertionMethod)
        ? object.assertionMethod.map((e: any) => String(e))
        : [],
      capabilityInvocation: Array.isArray(object?.capabilityInvocation)
        ? object.capabilityInvocation.map((e: any) => String(e))
        : [],
      capabilityDelegation: Array.isArray(object?.capabilityDelegation)
        ? object.capabilityDelegation.map((e: any) => String(e))
        : [],
      keyAgreement: Array.isArray(object?.keyAgreement)
        ? object.keyAgreement.map((e: any) => String(e))
        : [],
      service: Array.isArray(object?.service)
        ? object.service.map((e: any) => Service.fromJSON(e))
        : [],
      alsoKnownAs: Array.isArray(object?.alsoKnownAs)
        ? object.alsoKnownAs.map((e: any) => String(e))
        : [],
    };
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
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DidDocument>, I>>(
    object: I
  ): DidDocument {
    const message = createBaseDidDocument();
    message.context = object.context?.map((e) => e) || [];
    message.id = object.id ?? "";
    message.controller = object.controller?.map((e) => e) || [];
    message.verificationMethod =
      object.verificationMethod?.map((e) =>
        VerificationMethod.fromPartial(e)
      ) || [];
    message.authentication = object.authentication?.map((e) => e) || [];
    message.assertionMethod = object.assertionMethod?.map((e) => e) || [];
    message.capabilityInvocation =
      object.capabilityInvocation?.map((e) => e) || [];
    message.capabilityDelegation =
      object.capabilityDelegation?.map((e) => e) || [];
    message.keyAgreement = object.keyAgreement?.map((e) => e) || [];
    message.service = object.service?.map((e) => Service.fromPartial(e)) || [];
    message.alsoKnownAs = object.alsoKnownAs?.map((e) => e) || [];
    return message;
  },
};

function createBaseService(): Service {
  return { id: "", type: "", serviceEndpoint: "" };
}

export const Service = {
  encode(
    message: Service,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.type !== "") {
      writer.uint32(18).string(message.type);
    }
    if (message.serviceEndpoint !== "") {
      writer.uint32(26).string(message.serviceEndpoint);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Service {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseService();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.type = reader.string();
          break;
        case 3:
          message.serviceEndpoint = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Service {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      type: isSet(object.type) ? String(object.type) : "",
      serviceEndpoint: isSet(object.serviceEndpoint)
        ? String(object.serviceEndpoint)
        : "",
    };
  },

  toJSON(message: Service): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.type !== undefined && (obj.type = message.type);
    message.serviceEndpoint !== undefined &&
      (obj.serviceEndpoint = message.serviceEndpoint);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Service>, I>>(object: I): Service {
    const message = createBaseService();
    message.id = object.id ?? "";
    message.type = object.type ?? "";
    message.serviceEndpoint = object.serviceEndpoint ?? "";
    return message;
  },
};

function createBaseVerificationMethod(): VerificationMethod {
  return {
    id: "",
    type: "",
    controller: "",
    publicKeyHex: "",
    publicKeyBase58: "",
    blockchainAccountId: "",
  };
}

export const VerificationMethod = {
  encode(
    message: VerificationMethod,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.type !== "") {
      writer.uint32(18).string(message.type);
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

  decode(input: _m0.Reader | Uint8Array, length?: number): VerificationMethod {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVerificationMethod();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.type = reader.string();
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
    return {
      id: isSet(object.id) ? String(object.id) : "",
      type: isSet(object.type) ? String(object.type) : "",
      controller: isSet(object.controller) ? String(object.controller) : "",
      publicKeyHex: isSet(object.publicKeyHex)
        ? String(object.publicKeyHex)
        : "",
      publicKeyBase58: isSet(object.publicKeyBase58)
        ? String(object.publicKeyBase58)
        : "",
      blockchainAccountId: isSet(object.blockchainAccountId)
        ? String(object.blockchainAccountId)
        : "",
    };
  },

  toJSON(message: VerificationMethod): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.type !== undefined && (obj.type = message.type);
    message.controller !== undefined && (obj.controller = message.controller);
    message.publicKeyHex !== undefined &&
      (obj.publicKeyHex = message.publicKeyHex);
    message.publicKeyBase58 !== undefined &&
      (obj.publicKeyBase58 = message.publicKeyBase58);
    message.blockchainAccountId !== undefined &&
      (obj.blockchainAccountId = message.blockchainAccountId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VerificationMethod>, I>>(
    object: I
  ): VerificationMethod {
    const message = createBaseVerificationMethod();
    message.id = object.id ?? "";
    message.type = object.type ?? "";
    message.controller = object.controller ?? "";
    message.publicKeyHex = object.publicKeyHex ?? "";
    message.publicKeyBase58 = object.publicKeyBase58 ?? "";
    message.blockchainAccountId = object.blockchainAccountId ?? "";
    return message;
  },
};

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<
        Exclude<keyof I, KeysOfUnion<P>>,
        never
      >;

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
