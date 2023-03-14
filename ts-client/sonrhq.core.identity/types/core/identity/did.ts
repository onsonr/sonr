/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "sonrhq.core.identity";

/** DIDMethod is the DID method for each supported resolver. */
export enum DIDMethod {
  /** DIDMethod_BLOCKCHAIN - DID method for the Sonr network */
  DIDMethod_BLOCKCHAIN = 0,
  /** DIDMethod_WEB - DID method for the Ethereum network */
  DIDMethod_WEB = 1,
  /** DIDMethod_KEY - DID method for the Cosmos network */
  DIDMethod_KEY = 2,
  /** DIDMethod_IPFS - DID method for the Filecoin network */
  DIDMethod_IPFS = 3,
  /** DIDMethod_PEER - DID method for the Handshake network */
  DIDMethod_PEER = 4,
  UNRECOGNIZED = -1,
}

export function dIDMethodFromJSON(object: any): DIDMethod {
  switch (object) {
    case 0:
    case "DIDMethod_BLOCKCHAIN":
      return DIDMethod.DIDMethod_BLOCKCHAIN;
    case 1:
    case "DIDMethod_WEB":
      return DIDMethod.DIDMethod_WEB;
    case 2:
    case "DIDMethod_KEY":
      return DIDMethod.DIDMethod_KEY;
    case 3:
    case "DIDMethod_IPFS":
      return DIDMethod.DIDMethod_IPFS;
    case 4:
    case "DIDMethod_PEER":
      return DIDMethod.DIDMethod_PEER;
    case -1:
    case "UNRECOGNIZED":
    default:
      return DIDMethod.UNRECOGNIZED;
  }
}

export function dIDMethodToJSON(object: DIDMethod): string {
  switch (object) {
    case DIDMethod.DIDMethod_BLOCKCHAIN:
      return "DIDMethod_BLOCKCHAIN";
    case DIDMethod.DIDMethod_WEB:
      return "DIDMethod_WEB";
    case DIDMethod.DIDMethod_KEY:
      return "DIDMethod_KEY";
    case DIDMethod.DIDMethod_IPFS:
      return "DIDMethod_IPFS";
    case DIDMethod.DIDMethod_PEER:
      return "DIDMethod_PEER";
    case DIDMethod.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface DidDocument {
  context: string[];
  id: string;
  /** optional */
  controller: string[];
  /** optional */
  verificationMethod: VerificationMethod[];
  /** optional */
  authentication: VerificationRelationship[];
  /** optional */
  assertionMethod: VerificationRelationship[];
  /** optional */
  capabilityInvocation: VerificationRelationship[];
  /** optional */
  capabilityDelegation: VerificationRelationship[];
  /** optional */
  keyAgreement: VerificationRelationship[];
  /** optional */
  service: Service[];
  /** optional */
  alsoKnownAs: string[];
  /** optional */
  metadata: KeyValuePair[];
}

export interface VerificationMethod {
  id: string;
  type: string;
  controller: string;
  /** optional */
  publicKeyJwk: KeyValuePair[];
  /** optional */
  publicKeyMultibase: string;
  /** optional */
  blockchainAccountId: string;
  metadata: KeyValuePair[];
}

export interface VerificationRelationship {
  verificationMethod: VerificationMethod | undefined;
  reference: string;
}

export interface Service {
  id: string;
  controller: string;
  type: string;
  origin: string;
  name: string;
  /** optional */
  serviceEndpoints: KeyValuePair[];
  /** optional */
  metadata: KeyValuePair[];
}

export interface KeyValuePair {
  key: string;
  value: string;
}

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
    metadata: [],
  };
}

export const DidDocument = {
  encode(message: DidDocument, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
      VerificationRelationship.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.assertionMethod) {
      VerificationRelationship.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.capabilityInvocation) {
      VerificationRelationship.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    for (const v of message.capabilityDelegation) {
      VerificationRelationship.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    for (const v of message.keyAgreement) {
      VerificationRelationship.encode(v!, writer.uint32(74).fork()).ldelim();
    }
    for (const v of message.service) {
      Service.encode(v!, writer.uint32(82).fork()).ldelim();
    }
    for (const v of message.alsoKnownAs) {
      writer.uint32(90).string(v!);
    }
    for (const v of message.metadata) {
      KeyValuePair.encode(v!, writer.uint32(98).fork()).ldelim();
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
          message.verificationMethod.push(VerificationMethod.decode(reader, reader.uint32()));
          break;
        case 5:
          message.authentication.push(VerificationRelationship.decode(reader, reader.uint32()));
          break;
        case 6:
          message.assertionMethod.push(VerificationRelationship.decode(reader, reader.uint32()));
          break;
        case 7:
          message.capabilityInvocation.push(VerificationRelationship.decode(reader, reader.uint32()));
          break;
        case 8:
          message.capabilityDelegation.push(VerificationRelationship.decode(reader, reader.uint32()));
          break;
        case 9:
          message.keyAgreement.push(VerificationRelationship.decode(reader, reader.uint32()));
          break;
        case 10:
          message.service.push(Service.decode(reader, reader.uint32()));
          break;
        case 11:
          message.alsoKnownAs.push(reader.string());
          break;
        case 12:
          message.metadata.push(KeyValuePair.decode(reader, reader.uint32()));
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
      context: Array.isArray(object?.context) ? object.context.map((e: any) => String(e)) : [],
      id: isSet(object.id) ? String(object.id) : "",
      controller: Array.isArray(object?.controller) ? object.controller.map((e: any) => String(e)) : [],
      verificationMethod: Array.isArray(object?.verificationMethod)
        ? object.verificationMethod.map((e: any) => VerificationMethod.fromJSON(e))
        : [],
      authentication: Array.isArray(object?.authentication)
        ? object.authentication.map((e: any) => VerificationRelationship.fromJSON(e))
        : [],
      assertionMethod: Array.isArray(object?.assertionMethod)
        ? object.assertionMethod.map((e: any) => VerificationRelationship.fromJSON(e))
        : [],
      capabilityInvocation: Array.isArray(object?.capabilityInvocation)
        ? object.capabilityInvocation.map((e: any) => VerificationRelationship.fromJSON(e))
        : [],
      capabilityDelegation: Array.isArray(object?.capabilityDelegation)
        ? object.capabilityDelegation.map((e: any) => VerificationRelationship.fromJSON(e))
        : [],
      keyAgreement: Array.isArray(object?.keyAgreement)
        ? object.keyAgreement.map((e: any) => VerificationRelationship.fromJSON(e))
        : [],
      service: Array.isArray(object?.service) ? object.service.map((e: any) => Service.fromJSON(e)) : [],
      alsoKnownAs: Array.isArray(object?.alsoKnownAs) ? object.alsoKnownAs.map((e: any) => String(e)) : [],
      metadata: Array.isArray(object?.metadata) ? object.metadata.map((e: any) => KeyValuePair.fromJSON(e)) : [],
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
      obj.verificationMethod = message.verificationMethod.map((e) => e ? VerificationMethod.toJSON(e) : undefined);
    } else {
      obj.verificationMethod = [];
    }
    if (message.authentication) {
      obj.authentication = message.authentication.map((e) => e ? VerificationRelationship.toJSON(e) : undefined);
    } else {
      obj.authentication = [];
    }
    if (message.assertionMethod) {
      obj.assertionMethod = message.assertionMethod.map((e) => e ? VerificationRelationship.toJSON(e) : undefined);
    } else {
      obj.assertionMethod = [];
    }
    if (message.capabilityInvocation) {
      obj.capabilityInvocation = message.capabilityInvocation.map((e) =>
        e ? VerificationRelationship.toJSON(e) : undefined
      );
    } else {
      obj.capabilityInvocation = [];
    }
    if (message.capabilityDelegation) {
      obj.capabilityDelegation = message.capabilityDelegation.map((e) =>
        e ? VerificationRelationship.toJSON(e) : undefined
      );
    } else {
      obj.capabilityDelegation = [];
    }
    if (message.keyAgreement) {
      obj.keyAgreement = message.keyAgreement.map((e) => e ? VerificationRelationship.toJSON(e) : undefined);
    } else {
      obj.keyAgreement = [];
    }
    if (message.service) {
      obj.service = message.service.map((e) => e ? Service.toJSON(e) : undefined);
    } else {
      obj.service = [];
    }
    if (message.alsoKnownAs) {
      obj.alsoKnownAs = message.alsoKnownAs.map((e) => e);
    } else {
      obj.alsoKnownAs = [];
    }
    if (message.metadata) {
      obj.metadata = message.metadata.map((e) => e ? KeyValuePair.toJSON(e) : undefined);
    } else {
      obj.metadata = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DidDocument>, I>>(object: I): DidDocument {
    const message = createBaseDidDocument();
    message.context = object.context?.map((e) => e) || [];
    message.id = object.id ?? "";
    message.controller = object.controller?.map((e) => e) || [];
    message.verificationMethod = object.verificationMethod?.map((e) => VerificationMethod.fromPartial(e)) || [];
    message.authentication = object.authentication?.map((e) => VerificationRelationship.fromPartial(e)) || [];
    message.assertionMethod = object.assertionMethod?.map((e) => VerificationRelationship.fromPartial(e)) || [];
    message.capabilityInvocation = object.capabilityInvocation?.map((e) => VerificationRelationship.fromPartial(e))
      || [];
    message.capabilityDelegation = object.capabilityDelegation?.map((e) => VerificationRelationship.fromPartial(e))
      || [];
    message.keyAgreement = object.keyAgreement?.map((e) => VerificationRelationship.fromPartial(e)) || [];
    message.service = object.service?.map((e) => Service.fromPartial(e)) || [];
    message.alsoKnownAs = object.alsoKnownAs?.map((e) => e) || [];
    message.metadata = object.metadata?.map((e) => KeyValuePair.fromPartial(e)) || [];
    return message;
  },
};

function createBaseVerificationMethod(): VerificationMethod {
  return {
    id: "",
    type: "",
    controller: "",
    publicKeyJwk: [],
    publicKeyMultibase: "",
    blockchainAccountId: "",
    metadata: [],
  };
}

export const VerificationMethod = {
  encode(message: VerificationMethod, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.type !== "") {
      writer.uint32(18).string(message.type);
    }
    if (message.controller !== "") {
      writer.uint32(26).string(message.controller);
    }
    for (const v of message.publicKeyJwk) {
      KeyValuePair.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.publicKeyMultibase !== "") {
      writer.uint32(42).string(message.publicKeyMultibase);
    }
    if (message.blockchainAccountId !== "") {
      writer.uint32(50).string(message.blockchainAccountId);
    }
    for (const v of message.metadata) {
      KeyValuePair.encode(v!, writer.uint32(58).fork()).ldelim();
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
          message.publicKeyJwk.push(KeyValuePair.decode(reader, reader.uint32()));
          break;
        case 5:
          message.publicKeyMultibase = reader.string();
          break;
        case 6:
          message.blockchainAccountId = reader.string();
          break;
        case 7:
          message.metadata.push(KeyValuePair.decode(reader, reader.uint32()));
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
      publicKeyJwk: Array.isArray(object?.publicKeyJwk)
        ? object.publicKeyJwk.map((e: any) => KeyValuePair.fromJSON(e))
        : [],
      publicKeyMultibase: isSet(object.publicKeyMultibase) ? String(object.publicKeyMultibase) : "",
      blockchainAccountId: isSet(object.blockchainAccountId) ? String(object.blockchainAccountId) : "",
      metadata: Array.isArray(object?.metadata) ? object.metadata.map((e: any) => KeyValuePair.fromJSON(e)) : [],
    };
  },

  toJSON(message: VerificationMethod): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.type !== undefined && (obj.type = message.type);
    message.controller !== undefined && (obj.controller = message.controller);
    if (message.publicKeyJwk) {
      obj.publicKeyJwk = message.publicKeyJwk.map((e) => e ? KeyValuePair.toJSON(e) : undefined);
    } else {
      obj.publicKeyJwk = [];
    }
    message.publicKeyMultibase !== undefined && (obj.publicKeyMultibase = message.publicKeyMultibase);
    message.blockchainAccountId !== undefined && (obj.blockchainAccountId = message.blockchainAccountId);
    if (message.metadata) {
      obj.metadata = message.metadata.map((e) => e ? KeyValuePair.toJSON(e) : undefined);
    } else {
      obj.metadata = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VerificationMethod>, I>>(object: I): VerificationMethod {
    const message = createBaseVerificationMethod();
    message.id = object.id ?? "";
    message.type = object.type ?? "";
    message.controller = object.controller ?? "";
    message.publicKeyJwk = object.publicKeyJwk?.map((e) => KeyValuePair.fromPartial(e)) || [];
    message.publicKeyMultibase = object.publicKeyMultibase ?? "";
    message.blockchainAccountId = object.blockchainAccountId ?? "";
    message.metadata = object.metadata?.map((e) => KeyValuePair.fromPartial(e)) || [];
    return message;
  },
};

function createBaseVerificationRelationship(): VerificationRelationship {
  return { verificationMethod: undefined, reference: "" };
}

export const VerificationRelationship = {
  encode(message: VerificationRelationship, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.verificationMethod !== undefined) {
      VerificationMethod.encode(message.verificationMethod, writer.uint32(10).fork()).ldelim();
    }
    if (message.reference !== "") {
      writer.uint32(18).string(message.reference);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VerificationRelationship {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVerificationRelationship();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.verificationMethod = VerificationMethod.decode(reader, reader.uint32());
          break;
        case 2:
          message.reference = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VerificationRelationship {
    return {
      verificationMethod: isSet(object.verificationMethod)
        ? VerificationMethod.fromJSON(object.verificationMethod)
        : undefined,
      reference: isSet(object.reference) ? String(object.reference) : "",
    };
  },

  toJSON(message: VerificationRelationship): unknown {
    const obj: any = {};
    message.verificationMethod !== undefined && (obj.verificationMethod = message.verificationMethod
      ? VerificationMethod.toJSON(message.verificationMethod)
      : undefined);
    message.reference !== undefined && (obj.reference = message.reference);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VerificationRelationship>, I>>(object: I): VerificationRelationship {
    const message = createBaseVerificationRelationship();
    message.verificationMethod = (object.verificationMethod !== undefined && object.verificationMethod !== null)
      ? VerificationMethod.fromPartial(object.verificationMethod)
      : undefined;
    message.reference = object.reference ?? "";
    return message;
  },
};

function createBaseService(): Service {
  return { id: "", controller: "", type: "", origin: "", name: "", serviceEndpoints: [], metadata: [] };
}

export const Service = {
  encode(message: Service, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.controller !== "") {
      writer.uint32(18).string(message.controller);
    }
    if (message.type !== "") {
      writer.uint32(26).string(message.type);
    }
    if (message.origin !== "") {
      writer.uint32(34).string(message.origin);
    }
    if (message.name !== "") {
      writer.uint32(42).string(message.name);
    }
    for (const v of message.serviceEndpoints) {
      KeyValuePair.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.metadata) {
      KeyValuePair.encode(v!, writer.uint32(58).fork()).ldelim();
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
          message.controller = reader.string();
          break;
        case 3:
          message.type = reader.string();
          break;
        case 4:
          message.origin = reader.string();
          break;
        case 5:
          message.name = reader.string();
          break;
        case 6:
          message.serviceEndpoints.push(KeyValuePair.decode(reader, reader.uint32()));
          break;
        case 7:
          message.metadata.push(KeyValuePair.decode(reader, reader.uint32()));
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
      controller: isSet(object.controller) ? String(object.controller) : "",
      type: isSet(object.type) ? String(object.type) : "",
      origin: isSet(object.origin) ? String(object.origin) : "",
      name: isSet(object.name) ? String(object.name) : "",
      serviceEndpoints: Array.isArray(object?.serviceEndpoints)
        ? object.serviceEndpoints.map((e: any) => KeyValuePair.fromJSON(e))
        : [],
      metadata: Array.isArray(object?.metadata) ? object.metadata.map((e: any) => KeyValuePair.fromJSON(e)) : [],
    };
  },

  toJSON(message: Service): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.controller !== undefined && (obj.controller = message.controller);
    message.type !== undefined && (obj.type = message.type);
    message.origin !== undefined && (obj.origin = message.origin);
    message.name !== undefined && (obj.name = message.name);
    if (message.serviceEndpoints) {
      obj.serviceEndpoints = message.serviceEndpoints.map((e) => e ? KeyValuePair.toJSON(e) : undefined);
    } else {
      obj.serviceEndpoints = [];
    }
    if (message.metadata) {
      obj.metadata = message.metadata.map((e) => e ? KeyValuePair.toJSON(e) : undefined);
    } else {
      obj.metadata = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Service>, I>>(object: I): Service {
    const message = createBaseService();
    message.id = object.id ?? "";
    message.controller = object.controller ?? "";
    message.type = object.type ?? "";
    message.origin = object.origin ?? "";
    message.name = object.name ?? "";
    message.serviceEndpoints = object.serviceEndpoints?.map((e) => KeyValuePair.fromPartial(e)) || [];
    message.metadata = object.metadata?.map((e) => KeyValuePair.fromPartial(e)) || [];
    return message;
  },
};

function createBaseKeyValuePair(): KeyValuePair {
  return { key: "", value: "" };
}

export const KeyValuePair = {
  encode(message: KeyValuePair, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): KeyValuePair {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseKeyValuePair();
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

  fromJSON(object: any): KeyValuePair {
    return { key: isSet(object.key) ? String(object.key) : "", value: isSet(object.value) ? String(object.value) : "" };
  },

  toJSON(message: KeyValuePair): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<KeyValuePair>, I>>(object: I): KeyValuePair {
    const message = createBaseKeyValuePair();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
