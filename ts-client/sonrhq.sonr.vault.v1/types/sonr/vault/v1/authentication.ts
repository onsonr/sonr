/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { DidDocument } from "../../../core/identity/did";
import { AccountInfo } from "../../common/info";

export const protobufPackage = "sonrhq.sonr.vault.v1";

/** RegisterStartRequest is the request to register a new account. */
export interface RegisterStartRequest {
  /** The origin of the request. This is used to query the Blockchain for the Service DID. */
  origin: string;
  /** The user defined label for the device. */
  deviceLabel: string;
  /** The security threshold for the wallet account. */
  securityThreshold: number;
  /** The recovery passcode for the wallet account. */
  passcode: string;
  /** The Unique Identifier for the client device. Typically in a cookie. */
  uuid: string;
}

/** RegisterStartResponse is the response to a Register request. */
export interface RegisterStartResponse {
  /** Credential options for the user to sign with WebAuthn. */
  creationOptions: string;
  /** Relaying party id for the request. */
  rpId: string;
  /** Relaying party name for the request. */
  rpName: string;
}

/** RegisterFinishRequest is the request to CreateAccount a new key from the private key. */
export interface RegisterFinishRequest {
  /** The previously generated session id. */
  uuid: string;
  /** The signed credential response from the user. */
  credentialResponse: string;
  /** The origin of the request. This is used to query the Blockchain for the Service DID. */
  origin: string;
}

/** RegisterFinishResponse is the response to a CreateAccount request. */
export interface RegisterFinishResponse {
  /** The id of the account. */
  id: Uint8Array;
  /** The address of the account. */
  address: string;
  /** Relaying party id for the request. */
  rpId: string;
  /** Relaying party name for the request. */
  rpName: string;
  /** The DID Document for the wallet. */
  didDocument:
    | DidDocument
    | undefined;
  /** The account info for the wallet. */
  accountInfo:
    | AccountInfo
    | undefined;
  /** The UCAN token authorization header for subsequent requests. */
  ucanTokenHeader: Uint8Array;
}

/** LoginStartRequest is the request to login to an account. */
export interface LoginStartRequest {
  /** The origin of the request. This is used to query the Blockchain for the Service DID. */
  origin: string;
  /** The Sonr account address for the user. */
  accountAddress: string;
}

/** LoginStartResponse is the response to a Login request. */
export interface LoginStartResponse {
  /** Success is true if the account exists. */
  success: boolean;
  /** The account address for the user. */
  accountAddress: string;
  /** Json encoded WebAuthn credential options for the user to sign with. */
  credentialOptions: string;
  /** Relaying party id for the request. */
  rpId: string;
  /** Relaying party name for the request. */
  rpName: string;
}

/** LoginFinishRequest is the request to login to an account. */
export interface LoginFinishRequest {
  /** Address of the account to login to. */
  accountAddress: string;
  /** The signed credential response from the user. */
  credentialResponse: string;
  /** The origin of the request. This is used to query the Blockchain for the Service DID. */
  origin: string;
}

/** LoginFinishResponse is the response to a Login request. */
export interface LoginFinishResponse {
  /** Success is true if the account exists. */
  success: boolean;
  /** The account address for the user. */
  accountAddress: string;
  /** Relaying party id for the request. */
  rpId: string;
  /** Relaying party name for the request. */
  rpName: string;
  /** The DID Document for the wallet. */
  didDocument:
    | DidDocument
    | undefined;
  /** The account info for the wallet. */
  accountInfo:
    | AccountInfo
    | undefined;
  /** The UCAN token authorization header for subsequent requests. */
  ucanTokenHeader: Uint8Array;
}

function createBaseRegisterStartRequest(): RegisterStartRequest {
  return { origin: "", deviceLabel: "", securityThreshold: 0, passcode: "", uuid: "" };
}

export const RegisterStartRequest = {
  encode(message: RegisterStartRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.origin !== "") {
      writer.uint32(10).string(message.origin);
    }
    if (message.deviceLabel !== "") {
      writer.uint32(18).string(message.deviceLabel);
    }
    if (message.securityThreshold !== 0) {
      writer.uint32(24).int32(message.securityThreshold);
    }
    if (message.passcode !== "") {
      writer.uint32(34).string(message.passcode);
    }
    if (message.uuid !== "") {
      writer.uint32(42).string(message.uuid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterStartRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterStartRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.origin = reader.string();
          break;
        case 2:
          message.deviceLabel = reader.string();
          break;
        case 3:
          message.securityThreshold = reader.int32();
          break;
        case 4:
          message.passcode = reader.string();
          break;
        case 5:
          message.uuid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegisterStartRequest {
    return {
      origin: isSet(object.origin) ? String(object.origin) : "",
      deviceLabel: isSet(object.deviceLabel) ? String(object.deviceLabel) : "",
      securityThreshold: isSet(object.securityThreshold) ? Number(object.securityThreshold) : 0,
      passcode: isSet(object.passcode) ? String(object.passcode) : "",
      uuid: isSet(object.uuid) ? String(object.uuid) : "",
    };
  },

  toJSON(message: RegisterStartRequest): unknown {
    const obj: any = {};
    message.origin !== undefined && (obj.origin = message.origin);
    message.deviceLabel !== undefined && (obj.deviceLabel = message.deviceLabel);
    message.securityThreshold !== undefined && (obj.securityThreshold = Math.round(message.securityThreshold));
    message.passcode !== undefined && (obj.passcode = message.passcode);
    message.uuid !== undefined && (obj.uuid = message.uuid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterStartRequest>, I>>(object: I): RegisterStartRequest {
    const message = createBaseRegisterStartRequest();
    message.origin = object.origin ?? "";
    message.deviceLabel = object.deviceLabel ?? "";
    message.securityThreshold = object.securityThreshold ?? 0;
    message.passcode = object.passcode ?? "";
    message.uuid = object.uuid ?? "";
    return message;
  },
};

function createBaseRegisterStartResponse(): RegisterStartResponse {
  return { creationOptions: "", rpId: "", rpName: "" };
}

export const RegisterStartResponse = {
  encode(message: RegisterStartResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creationOptions !== "") {
      writer.uint32(10).string(message.creationOptions);
    }
    if (message.rpId !== "") {
      writer.uint32(18).string(message.rpId);
    }
    if (message.rpName !== "") {
      writer.uint32(26).string(message.rpName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterStartResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterStartResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creationOptions = reader.string();
          break;
        case 2:
          message.rpId = reader.string();
          break;
        case 3:
          message.rpName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegisterStartResponse {
    return {
      creationOptions: isSet(object.creationOptions) ? String(object.creationOptions) : "",
      rpId: isSet(object.rpId) ? String(object.rpId) : "",
      rpName: isSet(object.rpName) ? String(object.rpName) : "",
    };
  },

  toJSON(message: RegisterStartResponse): unknown {
    const obj: any = {};
    message.creationOptions !== undefined && (obj.creationOptions = message.creationOptions);
    message.rpId !== undefined && (obj.rpId = message.rpId);
    message.rpName !== undefined && (obj.rpName = message.rpName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterStartResponse>, I>>(object: I): RegisterStartResponse {
    const message = createBaseRegisterStartResponse();
    message.creationOptions = object.creationOptions ?? "";
    message.rpId = object.rpId ?? "";
    message.rpName = object.rpName ?? "";
    return message;
  },
};

function createBaseRegisterFinishRequest(): RegisterFinishRequest {
  return { uuid: "", credentialResponse: "", origin: "" };
}

export const RegisterFinishRequest = {
  encode(message: RegisterFinishRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uuid !== "") {
      writer.uint32(10).string(message.uuid);
    }
    if (message.credentialResponse !== "") {
      writer.uint32(18).string(message.credentialResponse);
    }
    if (message.origin !== "") {
      writer.uint32(26).string(message.origin);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterFinishRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterFinishRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uuid = reader.string();
          break;
        case 2:
          message.credentialResponse = reader.string();
          break;
        case 3:
          message.origin = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegisterFinishRequest {
    return {
      uuid: isSet(object.uuid) ? String(object.uuid) : "",
      credentialResponse: isSet(object.credentialResponse) ? String(object.credentialResponse) : "",
      origin: isSet(object.origin) ? String(object.origin) : "",
    };
  },

  toJSON(message: RegisterFinishRequest): unknown {
    const obj: any = {};
    message.uuid !== undefined && (obj.uuid = message.uuid);
    message.credentialResponse !== undefined && (obj.credentialResponse = message.credentialResponse);
    message.origin !== undefined && (obj.origin = message.origin);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterFinishRequest>, I>>(object: I): RegisterFinishRequest {
    const message = createBaseRegisterFinishRequest();
    message.uuid = object.uuid ?? "";
    message.credentialResponse = object.credentialResponse ?? "";
    message.origin = object.origin ?? "";
    return message;
  },
};

function createBaseRegisterFinishResponse(): RegisterFinishResponse {
  return {
    id: new Uint8Array(),
    address: "",
    rpId: "",
    rpName: "",
    didDocument: undefined,
    accountInfo: undefined,
    ucanTokenHeader: new Uint8Array(),
  };
}

export const RegisterFinishResponse = {
  encode(message: RegisterFinishResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id.length !== 0) {
      writer.uint32(10).bytes(message.id);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.rpId !== "") {
      writer.uint32(26).string(message.rpId);
    }
    if (message.rpName !== "") {
      writer.uint32(34).string(message.rpName);
    }
    if (message.didDocument !== undefined) {
      DidDocument.encode(message.didDocument, writer.uint32(42).fork()).ldelim();
    }
    if (message.accountInfo !== undefined) {
      AccountInfo.encode(message.accountInfo, writer.uint32(50).fork()).ldelim();
    }
    if (message.ucanTokenHeader.length !== 0) {
      writer.uint32(58).bytes(message.ucanTokenHeader);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterFinishResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterFinishResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.bytes();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.rpId = reader.string();
          break;
        case 4:
          message.rpName = reader.string();
          break;
        case 5:
          message.didDocument = DidDocument.decode(reader, reader.uint32());
          break;
        case 6:
          message.accountInfo = AccountInfo.decode(reader, reader.uint32());
          break;
        case 7:
          message.ucanTokenHeader = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegisterFinishResponse {
    return {
      id: isSet(object.id) ? bytesFromBase64(object.id) : new Uint8Array(),
      address: isSet(object.address) ? String(object.address) : "",
      rpId: isSet(object.rpId) ? String(object.rpId) : "",
      rpName: isSet(object.rpName) ? String(object.rpName) : "",
      didDocument: isSet(object.didDocument) ? DidDocument.fromJSON(object.didDocument) : undefined,
      accountInfo: isSet(object.accountInfo) ? AccountInfo.fromJSON(object.accountInfo) : undefined,
      ucanTokenHeader: isSet(object.ucanTokenHeader) ? bytesFromBase64(object.ucanTokenHeader) : new Uint8Array(),
    };
  },

  toJSON(message: RegisterFinishResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = base64FromBytes(message.id !== undefined ? message.id : new Uint8Array()));
    message.address !== undefined && (obj.address = message.address);
    message.rpId !== undefined && (obj.rpId = message.rpId);
    message.rpName !== undefined && (obj.rpName = message.rpName);
    message.didDocument !== undefined
      && (obj.didDocument = message.didDocument ? DidDocument.toJSON(message.didDocument) : undefined);
    message.accountInfo !== undefined
      && (obj.accountInfo = message.accountInfo ? AccountInfo.toJSON(message.accountInfo) : undefined);
    message.ucanTokenHeader !== undefined
      && (obj.ucanTokenHeader = base64FromBytes(
        message.ucanTokenHeader !== undefined ? message.ucanTokenHeader : new Uint8Array(),
      ));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterFinishResponse>, I>>(object: I): RegisterFinishResponse {
    const message = createBaseRegisterFinishResponse();
    message.id = object.id ?? new Uint8Array();
    message.address = object.address ?? "";
    message.rpId = object.rpId ?? "";
    message.rpName = object.rpName ?? "";
    message.didDocument = (object.didDocument !== undefined && object.didDocument !== null)
      ? DidDocument.fromPartial(object.didDocument)
      : undefined;
    message.accountInfo = (object.accountInfo !== undefined && object.accountInfo !== null)
      ? AccountInfo.fromPartial(object.accountInfo)
      : undefined;
    message.ucanTokenHeader = object.ucanTokenHeader ?? new Uint8Array();
    return message;
  },
};

function createBaseLoginStartRequest(): LoginStartRequest {
  return { origin: "", accountAddress: "" };
}

export const LoginStartRequest = {
  encode(message: LoginStartRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.origin !== "") {
      writer.uint32(10).string(message.origin);
    }
    if (message.accountAddress !== "") {
      writer.uint32(18).string(message.accountAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginStartRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginStartRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.origin = reader.string();
          break;
        case 2:
          message.accountAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LoginStartRequest {
    return {
      origin: isSet(object.origin) ? String(object.origin) : "",
      accountAddress: isSet(object.accountAddress) ? String(object.accountAddress) : "",
    };
  },

  toJSON(message: LoginStartRequest): unknown {
    const obj: any = {};
    message.origin !== undefined && (obj.origin = message.origin);
    message.accountAddress !== undefined && (obj.accountAddress = message.accountAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LoginStartRequest>, I>>(object: I): LoginStartRequest {
    const message = createBaseLoginStartRequest();
    message.origin = object.origin ?? "";
    message.accountAddress = object.accountAddress ?? "";
    return message;
  },
};

function createBaseLoginStartResponse(): LoginStartResponse {
  return { success: false, accountAddress: "", credentialOptions: "", rpId: "", rpName: "" };
}

export const LoginStartResponse = {
  encode(message: LoginStartResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    if (message.accountAddress !== "") {
      writer.uint32(18).string(message.accountAddress);
    }
    if (message.credentialOptions !== "") {
      writer.uint32(26).string(message.credentialOptions);
    }
    if (message.rpId !== "") {
      writer.uint32(34).string(message.rpId);
    }
    if (message.rpName !== "") {
      writer.uint32(42).string(message.rpName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginStartResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginStartResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        case 2:
          message.accountAddress = reader.string();
          break;
        case 3:
          message.credentialOptions = reader.string();
          break;
        case 4:
          message.rpId = reader.string();
          break;
        case 5:
          message.rpName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LoginStartResponse {
    return {
      success: isSet(object.success) ? Boolean(object.success) : false,
      accountAddress: isSet(object.accountAddress) ? String(object.accountAddress) : "",
      credentialOptions: isSet(object.credentialOptions) ? String(object.credentialOptions) : "",
      rpId: isSet(object.rpId) ? String(object.rpId) : "",
      rpName: isSet(object.rpName) ? String(object.rpName) : "",
    };
  },

  toJSON(message: LoginStartResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.accountAddress !== undefined && (obj.accountAddress = message.accountAddress);
    message.credentialOptions !== undefined && (obj.credentialOptions = message.credentialOptions);
    message.rpId !== undefined && (obj.rpId = message.rpId);
    message.rpName !== undefined && (obj.rpName = message.rpName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LoginStartResponse>, I>>(object: I): LoginStartResponse {
    const message = createBaseLoginStartResponse();
    message.success = object.success ?? false;
    message.accountAddress = object.accountAddress ?? "";
    message.credentialOptions = object.credentialOptions ?? "";
    message.rpId = object.rpId ?? "";
    message.rpName = object.rpName ?? "";
    return message;
  },
};

function createBaseLoginFinishRequest(): LoginFinishRequest {
  return { accountAddress: "", credentialResponse: "", origin: "" };
}

export const LoginFinishRequest = {
  encode(message: LoginFinishRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accountAddress !== "") {
      writer.uint32(10).string(message.accountAddress);
    }
    if (message.credentialResponse !== "") {
      writer.uint32(18).string(message.credentialResponse);
    }
    if (message.origin !== "") {
      writer.uint32(26).string(message.origin);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginFinishRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginFinishRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accountAddress = reader.string();
          break;
        case 2:
          message.credentialResponse = reader.string();
          break;
        case 3:
          message.origin = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LoginFinishRequest {
    return {
      accountAddress: isSet(object.accountAddress) ? String(object.accountAddress) : "",
      credentialResponse: isSet(object.credentialResponse) ? String(object.credentialResponse) : "",
      origin: isSet(object.origin) ? String(object.origin) : "",
    };
  },

  toJSON(message: LoginFinishRequest): unknown {
    const obj: any = {};
    message.accountAddress !== undefined && (obj.accountAddress = message.accountAddress);
    message.credentialResponse !== undefined && (obj.credentialResponse = message.credentialResponse);
    message.origin !== undefined && (obj.origin = message.origin);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LoginFinishRequest>, I>>(object: I): LoginFinishRequest {
    const message = createBaseLoginFinishRequest();
    message.accountAddress = object.accountAddress ?? "";
    message.credentialResponse = object.credentialResponse ?? "";
    message.origin = object.origin ?? "";
    return message;
  },
};

function createBaseLoginFinishResponse(): LoginFinishResponse {
  return {
    success: false,
    accountAddress: "",
    rpId: "",
    rpName: "",
    didDocument: undefined,
    accountInfo: undefined,
    ucanTokenHeader: new Uint8Array(),
  };
}

export const LoginFinishResponse = {
  encode(message: LoginFinishResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    if (message.accountAddress !== "") {
      writer.uint32(18).string(message.accountAddress);
    }
    if (message.rpId !== "") {
      writer.uint32(26).string(message.rpId);
    }
    if (message.rpName !== "") {
      writer.uint32(34).string(message.rpName);
    }
    if (message.didDocument !== undefined) {
      DidDocument.encode(message.didDocument, writer.uint32(42).fork()).ldelim();
    }
    if (message.accountInfo !== undefined) {
      AccountInfo.encode(message.accountInfo, writer.uint32(50).fork()).ldelim();
    }
    if (message.ucanTokenHeader.length !== 0) {
      writer.uint32(58).bytes(message.ucanTokenHeader);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginFinishResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginFinishResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        case 2:
          message.accountAddress = reader.string();
          break;
        case 3:
          message.rpId = reader.string();
          break;
        case 4:
          message.rpName = reader.string();
          break;
        case 5:
          message.didDocument = DidDocument.decode(reader, reader.uint32());
          break;
        case 6:
          message.accountInfo = AccountInfo.decode(reader, reader.uint32());
          break;
        case 7:
          message.ucanTokenHeader = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LoginFinishResponse {
    return {
      success: isSet(object.success) ? Boolean(object.success) : false,
      accountAddress: isSet(object.accountAddress) ? String(object.accountAddress) : "",
      rpId: isSet(object.rpId) ? String(object.rpId) : "",
      rpName: isSet(object.rpName) ? String(object.rpName) : "",
      didDocument: isSet(object.didDocument) ? DidDocument.fromJSON(object.didDocument) : undefined,
      accountInfo: isSet(object.accountInfo) ? AccountInfo.fromJSON(object.accountInfo) : undefined,
      ucanTokenHeader: isSet(object.ucanTokenHeader) ? bytesFromBase64(object.ucanTokenHeader) : new Uint8Array(),
    };
  },

  toJSON(message: LoginFinishResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.accountAddress !== undefined && (obj.accountAddress = message.accountAddress);
    message.rpId !== undefined && (obj.rpId = message.rpId);
    message.rpName !== undefined && (obj.rpName = message.rpName);
    message.didDocument !== undefined
      && (obj.didDocument = message.didDocument ? DidDocument.toJSON(message.didDocument) : undefined);
    message.accountInfo !== undefined
      && (obj.accountInfo = message.accountInfo ? AccountInfo.toJSON(message.accountInfo) : undefined);
    message.ucanTokenHeader !== undefined
      && (obj.ucanTokenHeader = base64FromBytes(
        message.ucanTokenHeader !== undefined ? message.ucanTokenHeader : new Uint8Array(),
      ));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LoginFinishResponse>, I>>(object: I): LoginFinishResponse {
    const message = createBaseLoginFinishResponse();
    message.success = object.success ?? false;
    message.accountAddress = object.accountAddress ?? "";
    message.rpId = object.rpId ?? "";
    message.rpName = object.rpName ?? "";
    message.didDocument = (object.didDocument !== undefined && object.didDocument !== null)
      ? DidDocument.fromPartial(object.didDocument)
      : undefined;
    message.accountInfo = (object.accountInfo !== undefined && object.accountInfo !== null)
      ? AccountInfo.fromPartial(object.accountInfo)
      : undefined;
    message.ucanTokenHeader = object.ucanTokenHeader ?? new Uint8Array();
    return message;
  },
};

/** Vault is the service used for managing a node's keypair. */
export interface VaultAuthentication {
  /**
   * Login Start
   *
   * {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
   * It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
   *
   * #### {{.RequestType.Name}}
   * | Name | Type | Description |
   * | ---- | ---- | ----------- |{{range .RequestType.Fields}}
   * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * #### {{.ResponseType.Name}}
   * | Name | Type | Description |
   * | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
   * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   */
  LoginStart(request: LoginStartRequest): Promise<LoginStartResponse>;
  /**
   * Login Finish
   *
   * {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
   * It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
   *
   * #### {{.RequestType.Name}}
   * | Name | Type | Description |
   * | ---- | ---- | ----------- |{{range .RequestType.Fields}}
   * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * #### {{.ResponseType.Name}}
   * | Name | Type | Description |
   * | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
   * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   */
  LoginFinish(request: LoginFinishRequest): Promise<LoginFinishResponse>;
  /**
   * Register Start
   *
   * {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
   * It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
   *
   * #### {{.RequestType.Name}}
   * | Name | Type | Description |
   * | ---- | ---- | ----------- |{{range .RequestType.Fields}}
   * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * #### {{.ResponseType.Name}}
   * | Name | Type | Description |
   * | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
   * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   */
  RegisterStart(request: RegisterStartRequest): Promise<RegisterStartResponse>;
  /**
   * Register Finish
   *
   * {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
   * It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
   *
   * #### {{.RequestType.Name}}
   * | Name | Type | Description |
   * | ---- | ---- | ----------- |{{range .RequestType.Fields}}
   * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   *
   * #### {{.ResponseType.Name}}
   * | Name | Type | Description |
   * | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
   * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
   */
  RegisterFinish(request: RegisterFinishRequest): Promise<RegisterFinishResponse>;
}

export class VaultAuthenticationClientImpl implements VaultAuthentication {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.LoginStart = this.LoginStart.bind(this);
    this.LoginFinish = this.LoginFinish.bind(this);
    this.RegisterStart = this.RegisterStart.bind(this);
    this.RegisterFinish = this.RegisterFinish.bind(this);
  }
  LoginStart(request: LoginStartRequest): Promise<LoginStartResponse> {
    const data = LoginStartRequest.encode(request).finish();
    const promise = this.rpc.request("sonrhq.sonr.vault.v1.VaultAuthentication", "LoginStart", data);
    return promise.then((data) => LoginStartResponse.decode(new _m0.Reader(data)));
  }

  LoginFinish(request: LoginFinishRequest): Promise<LoginFinishResponse> {
    const data = LoginFinishRequest.encode(request).finish();
    const promise = this.rpc.request("sonrhq.sonr.vault.v1.VaultAuthentication", "LoginFinish", data);
    return promise.then((data) => LoginFinishResponse.decode(new _m0.Reader(data)));
  }

  RegisterStart(request: RegisterStartRequest): Promise<RegisterStartResponse> {
    const data = RegisterStartRequest.encode(request).finish();
    const promise = this.rpc.request("sonrhq.sonr.vault.v1.VaultAuthentication", "RegisterStart", data);
    return promise.then((data) => RegisterStartResponse.decode(new _m0.Reader(data)));
  }

  RegisterFinish(request: RegisterFinishRequest): Promise<RegisterFinishResponse> {
    const data = RegisterFinishRequest.encode(request).finish();
    const promise = this.rpc.request("sonrhq.sonr.vault.v1.VaultAuthentication", "RegisterFinish", data);
    return promise.then((data) => RegisterFinishResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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
