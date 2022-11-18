/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "sonrio.sonr.identity";

export interface DidDocument {
  /** optional */
  context: string[];
  /** optional */
  creator: string;
  did: string;
  /** optional */
  controller: string[];
  /** optional */
  verificationMethod: VerificationMethod[];
  /** optional */
  authentication: string[];
  /** optional */
  assertionMethod: string[];
  /** optional */
  capabilityInvocation: string[];
  /** optional */
  capabilityDelegation: string[];
  /** optional */
  keyAgreement: string[];
  /** optional */
  service: Service[];
  /** optional */
  alsoKnownAs: string[];
}

export interface VerificationMethod {
  id: string;
  type: string;
  controller: string;
  /** optional */
  publicKeyJwk: KeyValuePair[];
  /** optional */
  publicKeyBase58: string;
  /** optional */
  credentialJson: Uint8Array;
}

export interface Service {
  id: string;
  type: string;
  serviceEndpoint: string;
  /** optional */
  serviceEndpointMap: KeyValuePair[];
}

export interface KeyValuePair {
  key: string;
  value: string;
}

function createBaseDidDocument(): DidDocument {
  return {
    context: [],
    creator: "",
    did: "",
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
  encode(message: DidDocument, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.context) {
      writer.uint32(10).string(v!);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(26).string(message.did);
    }
    for (const v of message.controller) {
      writer.uint32(34).string(v!);
    }
    for (const v of message.verificationMethod) {
      VerificationMethod.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.authentication) {
      writer.uint32(50).string(v!);
    }
    for (const v of message.assertionMethod) {
      writer.uint32(58).string(v!);
    }
    for (const v of message.capabilityInvocation) {
      writer.uint32(66).string(v!);
    }
    for (const v of message.capabilityDelegation) {
      writer.uint32(74).string(v!);
    }
    for (const v of message.keyAgreement) {
      writer.uint32(82).string(v!);
    }
    for (const v of message.service) {
      Service.encode(v!, writer.uint32(90).fork()).ldelim();
    }
    for (const v of message.alsoKnownAs) {
      writer.uint32(98).string(v!);
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
          message.creator = reader.string();
          break;
        case 3:
          message.did = reader.string();
          break;
        case 4:
          message.controller.push(reader.string());
          break;
        case 5:
          message.verificationMethod.push(VerificationMethod.decode(reader, reader.uint32()));
          break;
        case 6:
          message.authentication.push(reader.string());
          break;
        case 7:
          message.assertionMethod.push(reader.string());
          break;
        case 8:
          message.capabilityInvocation.push(reader.string());
          break;
        case 9:
          message.capabilityDelegation.push(reader.string());
          break;
        case 10:
          message.keyAgreement.push(reader.string());
          break;
        case 11:
          message.service.push(Service.decode(reader, reader.uint32()));
          break;
        case 12:
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
      context: Array.isArray(object?.context) ? object.context.map((e: any) => String(e)) : [],
      creator: isSet(object.creator) ? String(object.creator) : "",
      did: isSet(object.did) ? String(object.did) : "",
      controller: Array.isArray(object?.controller) ? object.controller.map((e: any) => String(e)) : [],
      verificationMethod: Array.isArray(object?.verificationMethod)
        ? object.verificationMethod.map((e: any) => VerificationMethod.fromJSON(e))
        : [],
      authentication: Array.isArray(object?.authentication) ? object.authentication.map((e: any) => String(e)) : [],
      assertionMethod: Array.isArray(object?.assertionMethod) ? object.assertionMethod.map((e: any) => String(e)) : [],
      capabilityInvocation: Array.isArray(object?.capabilityInvocation)
        ? object.capabilityInvocation.map((e: any) => String(e))
        : [],
      capabilityDelegation: Array.isArray(object?.capabilityDelegation)
        ? object.capabilityDelegation.map((e: any) => String(e))
        : [],
      keyAgreement: Array.isArray(object?.keyAgreement) ? object.keyAgreement.map((e: any) => String(e)) : [],
      service: Array.isArray(object?.service) ? object.service.map((e: any) => Service.fromJSON(e)) : [],
      alsoKnownAs: Array.isArray(object?.alsoKnownAs) ? object.alsoKnownAs.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: DidDocument): unknown {
    const obj: any = {};
    if (message.context) {
      obj.context = message.context.map((e) => e);
    } else {
      obj.context = [];
    }
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
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
      obj.service = message.service.map((e) => e ? Service.toJSON(e) : undefined);
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

  fromPartial<I extends Exact<DeepPartial<DidDocument>, I>>(object: I): DidDocument {
    const message = createBaseDidDocument();
    message.context = object.context?.map((e) => e) || [];
    message.creator = object.creator ?? "";
    message.did = object.did ?? "";
    message.controller = object.controller?.map((e) => e) || [];
    message.verificationMethod = object.verificationMethod?.map((e) => VerificationMethod.fromPartial(e)) || [];
    message.authentication = object.authentication?.map((e) => e) || [];
    message.assertionMethod = object.assertionMethod?.map((e) => e) || [];
    message.capabilityInvocation = object.capabilityInvocation?.map((e) => e) || [];
    message.capabilityDelegation = object.capabilityDelegation?.map((e) => e) || [];
    message.keyAgreement = object.keyAgreement?.map((e) => e) || [];
    message.service = object.service?.map((e) => Service.fromPartial(e)) || [];
    message.alsoKnownAs = object.alsoKnownAs?.map((e) => e) || [];
    return message;
  },
};

function createBaseVerificationMethod(): VerificationMethod {
  return { id: "", type: "", controller: "", publicKeyJwk: [], publicKeyBase58: "", credentialJson: new Uint8Array() };
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
    if (message.publicKeyBase58 !== "") {
      writer.uint32(42).string(message.publicKeyBase58);
    }
    if (message.credentialJson.length !== 0) {
      writer.uint32(50).bytes(message.credentialJson);
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
          message.publicKeyBase58 = reader.string();
          break;
        case 6:
          message.credentialJson = reader.bytes();
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
      publicKeyBase58: isSet(object.publicKeyBase58) ? String(object.publicKeyBase58) : "",
      credentialJson: isSet(object.credentialJson) ? bytesFromBase64(object.credentialJson) : new Uint8Array(),
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
    message.publicKeyBase58 !== undefined && (obj.publicKeyBase58 = message.publicKeyBase58);
    message.credentialJson !== undefined
      && (obj.credentialJson = base64FromBytes(
        message.credentialJson !== undefined ? message.credentialJson : new Uint8Array(),
      ));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VerificationMethod>, I>>(object: I): VerificationMethod {
    const message = createBaseVerificationMethod();
    message.id = object.id ?? "";
    message.type = object.type ?? "";
    message.controller = object.controller ?? "";
    message.publicKeyJwk = object.publicKeyJwk?.map((e) => KeyValuePair.fromPartial(e)) || [];
    message.publicKeyBase58 = object.publicKeyBase58 ?? "";
    message.credentialJson = object.credentialJson ?? new Uint8Array();
    return message;
  },
};

function createBaseService(): Service {
  return { id: "", type: "", serviceEndpoint: "", serviceEndpointMap: [] };
}

export const Service = {
  encode(message: Service, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.type !== "") {
      writer.uint32(18).string(message.type);
    }
    if (message.serviceEndpoint !== "") {
      writer.uint32(26).string(message.serviceEndpoint);
    }
    for (const v of message.serviceEndpointMap) {
      KeyValuePair.encode(v!, writer.uint32(34).fork()).ldelim();
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
        case 4:
          message.serviceEndpointMap.push(KeyValuePair.decode(reader, reader.uint32()));
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
      serviceEndpoint: isSet(object.serviceEndpoint) ? String(object.serviceEndpoint) : "",
      serviceEndpointMap: Array.isArray(object?.serviceEndpointMap)
        ? object.serviceEndpointMap.map((e: any) => KeyValuePair.fromJSON(e))
        : [],
    };
  },

  toJSON(message: Service): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.type !== undefined && (obj.type = message.type);
    message.serviceEndpoint !== undefined && (obj.serviceEndpoint = message.serviceEndpoint);
    if (message.serviceEndpointMap) {
      obj.serviceEndpointMap = message.serviceEndpointMap.map((e) => e ? KeyValuePair.toJSON(e) : undefined);
    } else {
      obj.serviceEndpointMap = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Service>, I>>(object: I): Service {
    const message = createBaseService();
    message.id = object.id ?? "";
    message.type = object.type ?? "";
    message.serviceEndpoint = object.serviceEndpoint ?? "";
    message.serviceEndpointMap = object.serviceEndpointMap?.map((e) => KeyValuePair.fromPartial(e)) || [];
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

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

function bytesFromBase64(b64: string): Uint8Array {
  if (globalThis.Buffer) {
    return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = globalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if (globalThis.Buffer) {
    return globalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(String.fromCharCode(byte));
    });
    return globalThis.btoa(bin.join(""));
  }
}

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
