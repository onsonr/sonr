/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { DidDocument } from "./did";

export const protobufPackage = "sonrhq.core.identity";

export interface MsgCreateDidDocument {
  creator: string;
  document: DidDocument | undefined;
}

export interface MsgCreateDidDocumentResponse {
}

export interface MsgUpdateDidDocument {
  creator: string;
  document: DidDocument | undefined;
}

export interface MsgUpdateDidDocumentResponse {
  creator: string;
}

export interface MsgDeleteDidDocument {
  creator: string;
  did: string;
}

export interface MsgDeleteDidDocumentResponse {
}

export interface MsgRegisterService {
  creator: string;
  index: string;
  domain: string;
}

export interface MsgRegisterServiceResponse {
}

export interface MsgUpdateService {
  creator: string;
  index: string;
  domain: string;
}

export interface MsgUpdateServiceResponse {
}

export interface MsgDeactivateService {
  creator: string;
  index: string;
  domain: string;
}

export interface MsgDeactivateServiceResponse {
}

function createBaseMsgCreateDidDocument(): MsgCreateDidDocument {
  return { creator: "", document: undefined };
}

export const MsgCreateDidDocument = {
  encode(message: MsgCreateDidDocument, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.document !== undefined) {
      DidDocument.encode(message.document, writer.uint32(18).fork()).ldelim();
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
          message.document = DidDocument.decode(reader, reader.uint32());
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
      document: isSet(object.document) ? DidDocument.fromJSON(object.document) : undefined,
    };
  },

  toJSON(message: MsgCreateDidDocument): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.document !== undefined
      && (obj.document = message.document ? DidDocument.toJSON(message.document) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateDidDocument>, I>>(object: I): MsgCreateDidDocument {
    const message = createBaseMsgCreateDidDocument();
    message.creator = object.creator ?? "";
    message.document = (object.document !== undefined && object.document !== null)
      ? DidDocument.fromPartial(object.document)
      : undefined;
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
  return { creator: "", document: undefined };
}

export const MsgUpdateDidDocument = {
  encode(message: MsgUpdateDidDocument, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.document !== undefined) {
      DidDocument.encode(message.document, writer.uint32(18).fork()).ldelim();
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
          message.document = DidDocument.decode(reader, reader.uint32());
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
      document: isSet(object.document) ? DidDocument.fromJSON(object.document) : undefined,
    };
  },

  toJSON(message: MsgUpdateDidDocument): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.document !== undefined
      && (obj.document = message.document ? DidDocument.toJSON(message.document) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateDidDocument>, I>>(object: I): MsgUpdateDidDocument {
    const message = createBaseMsgUpdateDidDocument();
    message.creator = object.creator ?? "";
    message.document = (object.document !== undefined && object.document !== null)
      ? DidDocument.fromPartial(object.document)
      : undefined;
    return message;
  },
};

function createBaseMsgUpdateDidDocumentResponse(): MsgUpdateDidDocumentResponse {
  return { creator: "" };
}

export const MsgUpdateDidDocumentResponse = {
  encode(message: MsgUpdateDidDocumentResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateDidDocumentResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateDidDocumentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateDidDocumentResponse {
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgUpdateDidDocumentResponse): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateDidDocumentResponse>, I>>(object: I): MsgUpdateDidDocumentResponse {
    const message = createBaseMsgUpdateDidDocumentResponse();
    message.creator = object.creator ?? "";
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

function createBaseMsgRegisterService(): MsgRegisterService {
  return { creator: "", index: "", domain: "" };
}

export const MsgRegisterService = {
  encode(message: MsgRegisterService, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.domain !== "") {
      writer.uint32(26).string(message.domain);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegisterService {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegisterService();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.domain = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterService {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      domain: isSet(object.domain) ? String(object.domain) : "",
    };
  },

  toJSON(message: MsgRegisterService): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.domain !== undefined && (obj.domain = message.domain);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegisterService>, I>>(object: I): MsgRegisterService {
    const message = createBaseMsgRegisterService();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.domain = object.domain ?? "";
    return message;
  },
};

function createBaseMsgRegisterServiceResponse(): MsgRegisterServiceResponse {
  return {};
}

export const MsgRegisterServiceResponse = {
  encode(_: MsgRegisterServiceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegisterServiceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegisterServiceResponse();
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

  fromJSON(_: any): MsgRegisterServiceResponse {
    return {};
  },

  toJSON(_: MsgRegisterServiceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegisterServiceResponse>, I>>(_: I): MsgRegisterServiceResponse {
    const message = createBaseMsgRegisterServiceResponse();
    return message;
  },
};

function createBaseMsgUpdateService(): MsgUpdateService {
  return { creator: "", index: "", domain: "" };
}

export const MsgUpdateService = {
  encode(message: MsgUpdateService, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.domain !== "") {
      writer.uint32(26).string(message.domain);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateService {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateService();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.domain = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateService {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      domain: isSet(object.domain) ? String(object.domain) : "",
    };
  },

  toJSON(message: MsgUpdateService): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.domain !== undefined && (obj.domain = message.domain);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateService>, I>>(object: I): MsgUpdateService {
    const message = createBaseMsgUpdateService();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.domain = object.domain ?? "";
    return message;
  },
};

function createBaseMsgUpdateServiceResponse(): MsgUpdateServiceResponse {
  return {};
}

export const MsgUpdateServiceResponse = {
  encode(_: MsgUpdateServiceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateServiceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateServiceResponse();
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

  fromJSON(_: any): MsgUpdateServiceResponse {
    return {};
  },

  toJSON(_: MsgUpdateServiceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateServiceResponse>, I>>(_: I): MsgUpdateServiceResponse {
    const message = createBaseMsgUpdateServiceResponse();
    return message;
  },
};

function createBaseMsgDeactivateService(): MsgDeactivateService {
  return { creator: "", index: "", domain: "" };
}

export const MsgDeactivateService = {
  encode(message: MsgDeactivateService, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.domain !== "") {
      writer.uint32(26).string(message.domain);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeactivateService {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeactivateService();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.domain = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeactivateService {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      domain: isSet(object.domain) ? String(object.domain) : "",
    };
  },

  toJSON(message: MsgDeactivateService): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.domain !== undefined && (obj.domain = message.domain);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeactivateService>, I>>(object: I): MsgDeactivateService {
    const message = createBaseMsgDeactivateService();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.domain = object.domain ?? "";
    return message;
  },
};

function createBaseMsgDeactivateServiceResponse(): MsgDeactivateServiceResponse {
  return {};
}

export const MsgDeactivateServiceResponse = {
  encode(_: MsgDeactivateServiceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeactivateServiceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeactivateServiceResponse();
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

  fromJSON(_: any): MsgDeactivateServiceResponse {
    return {};
  },

  toJSON(_: MsgDeactivateServiceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeactivateServiceResponse>, I>>(_: I): MsgDeactivateServiceResponse {
    const message = createBaseMsgDeactivateServiceResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateDidDocument(request: MsgCreateDidDocument): Promise<MsgCreateDidDocumentResponse>;
  UpdateDidDocument(request: MsgUpdateDidDocument): Promise<MsgUpdateDidDocumentResponse>;
  DeleteDidDocument(request: MsgDeleteDidDocument): Promise<MsgDeleteDidDocumentResponse>;
  RegisterService(request: MsgRegisterService): Promise<MsgRegisterServiceResponse>;
  UpdateService(request: MsgUpdateService): Promise<MsgUpdateServiceResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeactivateService(request: MsgDeactivateService): Promise<MsgDeactivateServiceResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateDidDocument = this.CreateDidDocument.bind(this);
    this.UpdateDidDocument = this.UpdateDidDocument.bind(this);
    this.DeleteDidDocument = this.DeleteDidDocument.bind(this);
    this.RegisterService = this.RegisterService.bind(this);
    this.UpdateService = this.UpdateService.bind(this);
    this.DeactivateService = this.DeactivateService.bind(this);
  }
  CreateDidDocument(request: MsgCreateDidDocument): Promise<MsgCreateDidDocumentResponse> {
    const data = MsgCreateDidDocument.encode(request).finish();
    const promise = this.rpc.request("sonrhq.core.identity.Msg", "CreateDidDocument", data);
    return promise.then((data) => MsgCreateDidDocumentResponse.decode(new _m0.Reader(data)));
  }

  UpdateDidDocument(request: MsgUpdateDidDocument): Promise<MsgUpdateDidDocumentResponse> {
    const data = MsgUpdateDidDocument.encode(request).finish();
    const promise = this.rpc.request("sonrhq.core.identity.Msg", "UpdateDidDocument", data);
    return promise.then((data) => MsgUpdateDidDocumentResponse.decode(new _m0.Reader(data)));
  }

  DeleteDidDocument(request: MsgDeleteDidDocument): Promise<MsgDeleteDidDocumentResponse> {
    const data = MsgDeleteDidDocument.encode(request).finish();
    const promise = this.rpc.request("sonrhq.core.identity.Msg", "DeleteDidDocument", data);
    return promise.then((data) => MsgDeleteDidDocumentResponse.decode(new _m0.Reader(data)));
  }

  RegisterService(request: MsgRegisterService): Promise<MsgRegisterServiceResponse> {
    const data = MsgRegisterService.encode(request).finish();
    const promise = this.rpc.request("sonrhq.core.identity.Msg", "RegisterService", data);
    return promise.then((data) => MsgRegisterServiceResponse.decode(new _m0.Reader(data)));
  }

  UpdateService(request: MsgUpdateService): Promise<MsgUpdateServiceResponse> {
    const data = MsgUpdateService.encode(request).finish();
    const promise = this.rpc.request("sonrhq.core.identity.Msg", "UpdateService", data);
    return promise.then((data) => MsgUpdateServiceResponse.decode(new _m0.Reader(data)));
  }

  DeactivateService(request: MsgDeactivateService): Promise<MsgDeactivateServiceResponse> {
    const data = MsgDeactivateService.encode(request).finish();
    const promise = this.rpc.request("sonrhq.core.identity.Msg", "DeactivateService", data);
    return promise.then((data) => MsgDeactivateServiceResponse.decode(new _m0.Reader(data)));
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
