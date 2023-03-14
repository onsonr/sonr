/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { DidDocument, Service } from "./did";

export const protobufPackage = "sonrhq.core.identity";

/** GenesisState defines the identity module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  didDocumentList: DidDocument[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  serviceList: Service[];
}

/** Params defines the parameters for the module. */
export interface Params {
  didBaseContext: string;
  didMethodContext: string;
  didMethodName: string;
  didMethodVersion: string;
  didNetwork: string;
  ipfsGateway: string;
  ipfsApi: string;
  webauthnAttestionPreference: string;
  webauthnAuthenticatorAttachment: string;
  webauthnTimeout: number;
}

function createBaseGenesisState(): GenesisState {
  return { params: undefined, didDocumentList: [], serviceList: [] };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.didDocumentList) {
      DidDocument.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.serviceList) {
      Service.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.didDocumentList.push(DidDocument.decode(reader, reader.uint32()));
          break;
        case 3:
          message.serviceList.push(Service.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      didDocumentList: Array.isArray(object?.didDocumentList)
        ? object.didDocumentList.map((e: any) => DidDocument.fromJSON(e))
        : [],
      serviceList: Array.isArray(object?.serviceList) ? object.serviceList.map((e: any) => Service.fromJSON(e)) : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.didDocumentList) {
      obj.didDocumentList = message.didDocumentList.map((e) => e ? DidDocument.toJSON(e) : undefined);
    } else {
      obj.didDocumentList = [];
    }
    if (message.serviceList) {
      obj.serviceList = message.serviceList.map((e) => e ? Service.toJSON(e) : undefined);
    } else {
      obj.serviceList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.didDocumentList = object.didDocumentList?.map((e) => DidDocument.fromPartial(e)) || [];
    message.serviceList = object.serviceList?.map((e) => Service.fromPartial(e)) || [];
    return message;
  },
};

function createBaseParams(): Params {
  return {
    didBaseContext: "",
    didMethodContext: "",
    didMethodName: "",
    didMethodVersion: "",
    didNetwork: "",
    ipfsGateway: "",
    ipfsApi: "",
    webauthnAttestionPreference: "",
    webauthnAuthenticatorAttachment: "",
    webauthnTimeout: 0,
  };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.didBaseContext !== "") {
      writer.uint32(10).string(message.didBaseContext);
    }
    if (message.didMethodContext !== "") {
      writer.uint32(18).string(message.didMethodContext);
    }
    if (message.didMethodName !== "") {
      writer.uint32(26).string(message.didMethodName);
    }
    if (message.didMethodVersion !== "") {
      writer.uint32(34).string(message.didMethodVersion);
    }
    if (message.didNetwork !== "") {
      writer.uint32(42).string(message.didNetwork);
    }
    if (message.ipfsGateway !== "") {
      writer.uint32(50).string(message.ipfsGateway);
    }
    if (message.ipfsApi !== "") {
      writer.uint32(58).string(message.ipfsApi);
    }
    if (message.webauthnAttestionPreference !== "") {
      writer.uint32(66).string(message.webauthnAttestionPreference);
    }
    if (message.webauthnAuthenticatorAttachment !== "") {
      writer.uint32(74).string(message.webauthnAuthenticatorAttachment);
    }
    if (message.webauthnTimeout !== 0) {
      writer.uint32(80).int32(message.webauthnTimeout);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.didBaseContext = reader.string();
          break;
        case 2:
          message.didMethodContext = reader.string();
          break;
        case 3:
          message.didMethodName = reader.string();
          break;
        case 4:
          message.didMethodVersion = reader.string();
          break;
        case 5:
          message.didNetwork = reader.string();
          break;
        case 6:
          message.ipfsGateway = reader.string();
          break;
        case 7:
          message.ipfsApi = reader.string();
          break;
        case 8:
          message.webauthnAttestionPreference = reader.string();
          break;
        case 9:
          message.webauthnAuthenticatorAttachment = reader.string();
          break;
        case 10:
          message.webauthnTimeout = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    return {
      didBaseContext: isSet(object.didBaseContext) ? String(object.didBaseContext) : "",
      didMethodContext: isSet(object.didMethodContext) ? String(object.didMethodContext) : "",
      didMethodName: isSet(object.didMethodName) ? String(object.didMethodName) : "",
      didMethodVersion: isSet(object.didMethodVersion) ? String(object.didMethodVersion) : "",
      didNetwork: isSet(object.didNetwork) ? String(object.didNetwork) : "",
      ipfsGateway: isSet(object.ipfsGateway) ? String(object.ipfsGateway) : "",
      ipfsApi: isSet(object.ipfsApi) ? String(object.ipfsApi) : "",
      webauthnAttestionPreference: isSet(object.webauthnAttestionPreference)
        ? String(object.webauthnAttestionPreference)
        : "",
      webauthnAuthenticatorAttachment: isSet(object.webauthnAuthenticatorAttachment)
        ? String(object.webauthnAuthenticatorAttachment)
        : "",
      webauthnTimeout: isSet(object.webauthnTimeout) ? Number(object.webauthnTimeout) : 0,
    };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.didBaseContext !== undefined && (obj.didBaseContext = message.didBaseContext);
    message.didMethodContext !== undefined && (obj.didMethodContext = message.didMethodContext);
    message.didMethodName !== undefined && (obj.didMethodName = message.didMethodName);
    message.didMethodVersion !== undefined && (obj.didMethodVersion = message.didMethodVersion);
    message.didNetwork !== undefined && (obj.didNetwork = message.didNetwork);
    message.ipfsGateway !== undefined && (obj.ipfsGateway = message.ipfsGateway);
    message.ipfsApi !== undefined && (obj.ipfsApi = message.ipfsApi);
    message.webauthnAttestionPreference !== undefined
      && (obj.webauthnAttestionPreference = message.webauthnAttestionPreference);
    message.webauthnAuthenticatorAttachment !== undefined
      && (obj.webauthnAuthenticatorAttachment = message.webauthnAuthenticatorAttachment);
    message.webauthnTimeout !== undefined && (obj.webauthnTimeout = Math.round(message.webauthnTimeout));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.didBaseContext = object.didBaseContext ?? "";
    message.didMethodContext = object.didMethodContext ?? "";
    message.didMethodName = object.didMethodName ?? "";
    message.didMethodVersion = object.didMethodVersion ?? "";
    message.didNetwork = object.didNetwork ?? "";
    message.ipfsGateway = object.ipfsGateway ?? "";
    message.ipfsApi = object.ipfsApi ?? "";
    message.webauthnAttestionPreference = object.webauthnAttestionPreference ?? "";
    message.webauthnAuthenticatorAttachment = object.webauthnAuthenticatorAttachment ?? "";
    message.webauthnTimeout = object.webauthnTimeout ?? 0;
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
