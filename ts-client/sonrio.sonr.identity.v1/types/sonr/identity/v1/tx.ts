/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "sonrio.sonr.identity.v1";

export interface MsgCreateDidDocument {
  creator: string;
  did: string;
  context: string;
  controller: string;
  verificationMethod: string;
  authentication: string;
  assertionMethod: string;
  capibilityInvocation: string;
  capabilityDelegation: string;
  keyAgreement: string;
  service: string;
  alsoKnownAs: string;
}

export interface MsgCreateDidDocumentResponse {
}

export interface MsgUpdateDidDocument {
  creator: string;
  did: string;
  context: string;
  controller: string;
  verificationMethod: string;
  authentication: string;
  assertionMethod: string;
  capibilityInvocation: string;
  capabilityDelegation: string;
  keyAgreement: string;
  service: string;
  alsoKnownAs: string;
}

export interface MsgUpdateDidDocumentResponse {
}

export interface MsgDeleteDidDocument {
  creator: string;
  did: string;
}

export interface MsgDeleteDidDocumentResponse {
}

function createBaseMsgCreateDidDocument(): MsgCreateDidDocument {
  return {
    creator: "",
    did: "",
    context: "",
    controller: "",
    verificationMethod: "",
    authentication: "",
    assertionMethod: "",
    capibilityInvocation: "",
    capabilityDelegation: "",
    keyAgreement: "",
    service: "",
    alsoKnownAs: "",
  };
}

export const MsgCreateDidDocument = {
  encode(message: MsgCreateDidDocument, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.context !== "") {
      writer.uint32(26).string(message.context);
    }
    if (message.controller !== "") {
      writer.uint32(34).string(message.controller);
    }
    if (message.verificationMethod !== "") {
      writer.uint32(42).string(message.verificationMethod);
    }
    if (message.authentication !== "") {
      writer.uint32(50).string(message.authentication);
    }
    if (message.assertionMethod !== "") {
      writer.uint32(58).string(message.assertionMethod);
    }
    if (message.capibilityInvocation !== "") {
      writer.uint32(66).string(message.capibilityInvocation);
    }
    if (message.capabilityDelegation !== "") {
      writer.uint32(74).string(message.capabilityDelegation);
    }
    if (message.keyAgreement !== "") {
      writer.uint32(82).string(message.keyAgreement);
    }
    if (message.service !== "") {
      writer.uint32(90).string(message.service);
    }
    if (message.alsoKnownAs !== "") {
      writer.uint32(98).string(message.alsoKnownAs);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateDidDocument {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateDidDocument();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.context = reader.string();
          break;
        case 4:
          message.controller = reader.string();
          break;
        case 5:
          message.verificationMethod = reader.string();
          break;
        case 6:
          message.authentication = reader.string();
          break;
        case 7:
          message.assertionMethod = reader.string();
          break;
        case 8:
          message.capibilityInvocation = reader.string();
          break;
        case 9:
          message.capabilityDelegation = reader.string();
          break;
        case 10:
          message.keyAgreement = reader.string();
          break;
        case 11:
          message.service = reader.string();
          break;
        case 12:
          message.alsoKnownAs = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateDidDocument {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      did: isSet(object.did) ? String(object.did) : "",
      context: isSet(object.context) ? String(object.context) : "",
      controller: isSet(object.controller) ? String(object.controller) : "",
      verificationMethod: isSet(object.verificationMethod) ? String(object.verificationMethod) : "",
      authentication: isSet(object.authentication) ? String(object.authentication) : "",
      assertionMethod: isSet(object.assertionMethod) ? String(object.assertionMethod) : "",
      capibilityInvocation: isSet(object.capibilityInvocation) ? String(object.capibilityInvocation) : "",
      capabilityDelegation: isSet(object.capabilityDelegation) ? String(object.capabilityDelegation) : "",
      keyAgreement: isSet(object.keyAgreement) ? String(object.keyAgreement) : "",
      service: isSet(object.service) ? String(object.service) : "",
      alsoKnownAs: isSet(object.alsoKnownAs) ? String(object.alsoKnownAs) : "",
    };
  },

  toJSON(message: MsgCreateDidDocument): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.context !== undefined && (obj.context = message.context);
    message.controller !== undefined && (obj.controller = message.controller);
    message.verificationMethod !== undefined && (obj.verificationMethod = message.verificationMethod);
    message.authentication !== undefined && (obj.authentication = message.authentication);
    message.assertionMethod !== undefined && (obj.assertionMethod = message.assertionMethod);
    message.capibilityInvocation !== undefined && (obj.capibilityInvocation = message.capibilityInvocation);
    message.capabilityDelegation !== undefined && (obj.capabilityDelegation = message.capabilityDelegation);
    message.keyAgreement !== undefined && (obj.keyAgreement = message.keyAgreement);
    message.service !== undefined && (obj.service = message.service);
    message.alsoKnownAs !== undefined && (obj.alsoKnownAs = message.alsoKnownAs);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateDidDocument>, I>>(object: I): MsgCreateDidDocument {
    const message = createBaseMsgCreateDidDocument();
    message.creator = object.creator ?? "";
    message.did = object.did ?? "";
    message.context = object.context ?? "";
    message.controller = object.controller ?? "";
    message.verificationMethod = object.verificationMethod ?? "";
    message.authentication = object.authentication ?? "";
    message.assertionMethod = object.assertionMethod ?? "";
    message.capibilityInvocation = object.capibilityInvocation ?? "";
    message.capabilityDelegation = object.capabilityDelegation ?? "";
    message.keyAgreement = object.keyAgreement ?? "";
    message.service = object.service ?? "";
    message.alsoKnownAs = object.alsoKnownAs ?? "";
    return message;
  },
};

function createBaseMsgCreateDidDocumentResponse(): MsgCreateDidDocumentResponse {
  return {};
}

export const MsgCreateDidDocumentResponse = {
  encode(_: MsgCreateDidDocumentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateDidDocumentResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateDidDocumentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateDidDocumentResponse {
    return {};
  },

  toJSON(_: MsgCreateDidDocumentResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateDidDocumentResponse>, I>>(_: I): MsgCreateDidDocumentResponse {
    const message = createBaseMsgCreateDidDocumentResponse();
    return message;
  },
};

function createBaseMsgUpdateDidDocument(): MsgUpdateDidDocument {
  return {
    creator: "",
    did: "",
    context: "",
    controller: "",
    verificationMethod: "",
    authentication: "",
    assertionMethod: "",
    capibilityInvocation: "",
    capabilityDelegation: "",
    keyAgreement: "",
    service: "",
    alsoKnownAs: "",
  };
}

export const MsgUpdateDidDocument = {
  encode(message: MsgUpdateDidDocument, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.context !== "") {
      writer.uint32(26).string(message.context);
    }
    if (message.controller !== "") {
      writer.uint32(34).string(message.controller);
    }
    if (message.verificationMethod !== "") {
      writer.uint32(42).string(message.verificationMethod);
    }
    if (message.authentication !== "") {
      writer.uint32(50).string(message.authentication);
    }
    if (message.assertionMethod !== "") {
      writer.uint32(58).string(message.assertionMethod);
    }
    if (message.capibilityInvocation !== "") {
      writer.uint32(66).string(message.capibilityInvocation);
    }
    if (message.capabilityDelegation !== "") {
      writer.uint32(74).string(message.capabilityDelegation);
    }
    if (message.keyAgreement !== "") {
      writer.uint32(82).string(message.keyAgreement);
    }
    if (message.service !== "") {
      writer.uint32(90).string(message.service);
    }
    if (message.alsoKnownAs !== "") {
      writer.uint32(98).string(message.alsoKnownAs);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateDidDocument {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateDidDocument();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.context = reader.string();
          break;
        case 4:
          message.controller = reader.string();
          break;
        case 5:
          message.verificationMethod = reader.string();
          break;
        case 6:
          message.authentication = reader.string();
          break;
        case 7:
          message.assertionMethod = reader.string();
          break;
        case 8:
          message.capibilityInvocation = reader.string();
          break;
        case 9:
          message.capabilityDelegation = reader.string();
          break;
        case 10:
          message.keyAgreement = reader.string();
          break;
        case 11:
          message.service = reader.string();
          break;
        case 12:
          message.alsoKnownAs = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateDidDocument {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      did: isSet(object.did) ? String(object.did) : "",
      context: isSet(object.context) ? String(object.context) : "",
      controller: isSet(object.controller) ? String(object.controller) : "",
      verificationMethod: isSet(object.verificationMethod) ? String(object.verificationMethod) : "",
      authentication: isSet(object.authentication) ? String(object.authentication) : "",
      assertionMethod: isSet(object.assertionMethod) ? String(object.assertionMethod) : "",
      capibilityInvocation: isSet(object.capibilityInvocation) ? String(object.capibilityInvocation) : "",
      capabilityDelegation: isSet(object.capabilityDelegation) ? String(object.capabilityDelegation) : "",
      keyAgreement: isSet(object.keyAgreement) ? String(object.keyAgreement) : "",
      service: isSet(object.service) ? String(object.service) : "",
      alsoKnownAs: isSet(object.alsoKnownAs) ? String(object.alsoKnownAs) : "",
    };
  },

  toJSON(message: MsgUpdateDidDocument): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.context !== undefined && (obj.context = message.context);
    message.controller !== undefined && (obj.controller = message.controller);
    message.verificationMethod !== undefined && (obj.verificationMethod = message.verificationMethod);
    message.authentication !== undefined && (obj.authentication = message.authentication);
    message.assertionMethod !== undefined && (obj.assertionMethod = message.assertionMethod);
    message.capibilityInvocation !== undefined && (obj.capibilityInvocation = message.capibilityInvocation);
    message.capabilityDelegation !== undefined && (obj.capabilityDelegation = message.capabilityDelegation);
    message.keyAgreement !== undefined && (obj.keyAgreement = message.keyAgreement);
    message.service !== undefined && (obj.service = message.service);
    message.alsoKnownAs !== undefined && (obj.alsoKnownAs = message.alsoKnownAs);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateDidDocument>, I>>(object: I): MsgUpdateDidDocument {
    const message = createBaseMsgUpdateDidDocument();
    message.creator = object.creator ?? "";
    message.did = object.did ?? "";
    message.context = object.context ?? "";
    message.controller = object.controller ?? "";
    message.verificationMethod = object.verificationMethod ?? "";
    message.authentication = object.authentication ?? "";
    message.assertionMethod = object.assertionMethod ?? "";
    message.capibilityInvocation = object.capibilityInvocation ?? "";
    message.capabilityDelegation = object.capabilityDelegation ?? "";
    message.keyAgreement = object.keyAgreement ?? "";
    message.service = object.service ?? "";
    message.alsoKnownAs = object.alsoKnownAs ?? "";
    return message;
  },
};

function createBaseMsgUpdateDidDocumentResponse(): MsgUpdateDidDocumentResponse {
  return {};
}

export const MsgUpdateDidDocumentResponse = {
  encode(_: MsgUpdateDidDocumentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateDidDocumentResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateDidDocumentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateDidDocumentResponse {
    return {};
  },

  toJSON(_: MsgUpdateDidDocumentResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateDidDocumentResponse>, I>>(_: I): MsgUpdateDidDocumentResponse {
    const message = createBaseMsgUpdateDidDocumentResponse();
    return message;
  },
};

function createBaseMsgDeleteDidDocument(): MsgDeleteDidDocument {
  return { creator: "", did: "" };
}

export const MsgDeleteDidDocument = {
  encode(message: MsgDeleteDidDocument, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteDidDocument {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteDidDocument();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteDidDocument {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      did: isSet(object.did) ? String(object.did) : "",
    };
  },

  toJSON(message: MsgDeleteDidDocument): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteDidDocument>, I>>(object: I): MsgDeleteDidDocument {
    const message = createBaseMsgDeleteDidDocument();
    message.creator = object.creator ?? "";
    message.did = object.did ?? "";
    return message;
  },
};

function createBaseMsgDeleteDidDocumentResponse(): MsgDeleteDidDocumentResponse {
  return {};
}

export const MsgDeleteDidDocumentResponse = {
  encode(_: MsgDeleteDidDocumentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteDidDocumentResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteDidDocumentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteDidDocumentResponse {
    return {};
  },

  toJSON(_: MsgDeleteDidDocumentResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteDidDocumentResponse>, I>>(_: I): MsgDeleteDidDocumentResponse {
    const message = createBaseMsgDeleteDidDocumentResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateDidDocument(request: MsgCreateDidDocument): Promise<MsgCreateDidDocumentResponse>;
  UpdateDidDocument(request: MsgUpdateDidDocument): Promise<MsgUpdateDidDocumentResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteDidDocument(request: MsgDeleteDidDocument): Promise<MsgDeleteDidDocumentResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateDidDocument = this.CreateDidDocument.bind(this);
    this.UpdateDidDocument = this.UpdateDidDocument.bind(this);
    this.DeleteDidDocument = this.DeleteDidDocument.bind(this);
  }
  CreateDidDocument(request: MsgCreateDidDocument): Promise<MsgCreateDidDocumentResponse> {
    const data = MsgCreateDidDocument.encode(request).finish();
    const promise = this.rpc.request("sonrio.sonr.identity.v1.Msg", "CreateDidDocument", data);
    return promise.then((data) => MsgCreateDidDocumentResponse.decode(new _m0.Reader(data)));
  }

  UpdateDidDocument(request: MsgUpdateDidDocument): Promise<MsgUpdateDidDocumentResponse> {
    const data = MsgUpdateDidDocument.encode(request).finish();
    const promise = this.rpc.request("sonrio.sonr.identity.v1.Msg", "UpdateDidDocument", data);
    return promise.then((data) => MsgUpdateDidDocumentResponse.decode(new _m0.Reader(data)));
  }

  DeleteDidDocument(request: MsgDeleteDidDocument): Promise<MsgDeleteDidDocumentResponse> {
    const data = MsgDeleteDidDocument.encode(request).finish();
    const promise = this.rpc.request("sonrio.sonr.identity.v1.Msg", "DeleteDidDocument", data);
    return promise.then((data) => MsgDeleteDidDocumentResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
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
