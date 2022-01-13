/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Peer, Metadata } from "../../../common/core";

export const protobufPackage = "protocols.discover.v1";

/** LobbyMessage is message passed from Peer in Lobby */
export interface LobbyMessage {
  /** Users Peer Data */
  peer?: Peer;
  /** Message to be published */
  message?: string | undefined;
  /** Metadata */
  metadata?: Metadata;
}

/** VisibilityRequest is Message for updating Peer Visibility in Exchange */
export interface VisibilityRequest {
  /** SName combined with Device ID and Hashed */
  sName: string;
  /** Buffer of Public Key */
  publicKey: Buffer;
  /** Visibility */
  visibility: VisibilityRequest_Visibility;
}

export enum VisibilityRequest_Visibility {
  VISIBILITY_UNSPECIFIED = "VISIBILITY_UNSPECIFIED",
  /** VISIBILITY_AVAILABLE - Everyone can see this peer */
  VISIBILITY_AVAILABLE = "VISIBILITY_AVAILABLE",
  /** VISIBILITY_HIDDEN - Only Linked Devices can see this peer */
  VISIBILITY_HIDDEN = "VISIBILITY_HIDDEN",
  /** VISIBILITY_FRIENDS - Only Friends can see this peer */
  VISIBILITY_FRIENDS = "VISIBILITY_FRIENDS",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function visibilityRequest_VisibilityFromJSON(
  object: any
): VisibilityRequest_Visibility {
  switch (object) {
    case 0:
    case "VISIBILITY_UNSPECIFIED":
      return VisibilityRequest_Visibility.VISIBILITY_UNSPECIFIED;
    case 1:
    case "VISIBILITY_AVAILABLE":
      return VisibilityRequest_Visibility.VISIBILITY_AVAILABLE;
    case 2:
    case "VISIBILITY_HIDDEN":
      return VisibilityRequest_Visibility.VISIBILITY_HIDDEN;
    case 3:
    case "VISIBILITY_FRIENDS":
      return VisibilityRequest_Visibility.VISIBILITY_FRIENDS;
    case -1:
    case "UNRECOGNIZED":
    default:
      return VisibilityRequest_Visibility.UNRECOGNIZED;
  }
}

export function visibilityRequest_VisibilityToJSON(
  object: VisibilityRequest_Visibility
): string {
  switch (object) {
    case VisibilityRequest_Visibility.VISIBILITY_UNSPECIFIED:
      return "VISIBILITY_UNSPECIFIED";
    case VisibilityRequest_Visibility.VISIBILITY_AVAILABLE:
      return "VISIBILITY_AVAILABLE";
    case VisibilityRequest_Visibility.VISIBILITY_HIDDEN:
      return "VISIBILITY_HIDDEN";
    case VisibilityRequest_Visibility.VISIBILITY_FRIENDS:
      return "VISIBILITY_FRIENDS";
    default:
      return "UNKNOWN";
  }
}

export function visibilityRequest_VisibilityToNumber(
  object: VisibilityRequest_Visibility
): number {
  switch (object) {
    case VisibilityRequest_Visibility.VISIBILITY_UNSPECIFIED:
      return 0;
    case VisibilityRequest_Visibility.VISIBILITY_AVAILABLE:
      return 1;
    case VisibilityRequest_Visibility.VISIBILITY_HIDDEN:
      return 2;
    case VisibilityRequest_Visibility.VISIBILITY_FRIENDS:
      return 3;
    default:
      return 0;
  }
}

/** VisibilityResponse is response for VisibilityRequest */
export interface VisibilityResponse {
  /** If Request was Successful */
  success: boolean;
  /** Error Message if Request was not successful */
  error: string;
  /** Visibility */
  visibility: VisibilityResponse_Visibility;
}

export enum VisibilityResponse_Visibility {
  VISIBILITY_UNSPECIFIED = "VISIBILITY_UNSPECIFIED",
  /** VISIBILITY_AVAILABLE - Everyone can see this peer */
  VISIBILITY_AVAILABLE = "VISIBILITY_AVAILABLE",
  /** VISIBILITY_HIDDEN - Only Linked Devices can see this peer */
  VISIBILITY_HIDDEN = "VISIBILITY_HIDDEN",
  /** VISIBILITY_FRIENDS - Only Friends can see this peer */
  VISIBILITY_FRIENDS = "VISIBILITY_FRIENDS",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function visibilityResponse_VisibilityFromJSON(
  object: any
): VisibilityResponse_Visibility {
  switch (object) {
    case 0:
    case "VISIBILITY_UNSPECIFIED":
      return VisibilityResponse_Visibility.VISIBILITY_UNSPECIFIED;
    case 1:
    case "VISIBILITY_AVAILABLE":
      return VisibilityResponse_Visibility.VISIBILITY_AVAILABLE;
    case 2:
    case "VISIBILITY_HIDDEN":
      return VisibilityResponse_Visibility.VISIBILITY_HIDDEN;
    case 3:
    case "VISIBILITY_FRIENDS":
      return VisibilityResponse_Visibility.VISIBILITY_FRIENDS;
    case -1:
    case "UNRECOGNIZED":
    default:
      return VisibilityResponse_Visibility.UNRECOGNIZED;
  }
}

export function visibilityResponse_VisibilityToJSON(
  object: VisibilityResponse_Visibility
): string {
  switch (object) {
    case VisibilityResponse_Visibility.VISIBILITY_UNSPECIFIED:
      return "VISIBILITY_UNSPECIFIED";
    case VisibilityResponse_Visibility.VISIBILITY_AVAILABLE:
      return "VISIBILITY_AVAILABLE";
    case VisibilityResponse_Visibility.VISIBILITY_HIDDEN:
      return "VISIBILITY_HIDDEN";
    case VisibilityResponse_Visibility.VISIBILITY_FRIENDS:
      return "VISIBILITY_FRIENDS";
    default:
      return "UNKNOWN";
  }
}

export function visibilityResponse_VisibilityToNumber(
  object: VisibilityResponse_Visibility
): number {
  switch (object) {
    case VisibilityResponse_Visibility.VISIBILITY_UNSPECIFIED:
      return 0;
    case VisibilityResponse_Visibility.VISIBILITY_AVAILABLE:
      return 1;
    case VisibilityResponse_Visibility.VISIBILITY_HIDDEN:
      return 2;
    case VisibilityResponse_Visibility.VISIBILITY_FRIENDS:
      return 3;
    default:
      return 0;
  }
}

function createBaseLobbyMessage(): LobbyMessage {
  return { peer: undefined, message: undefined, metadata: undefined };
}

export const LobbyMessage = {
  encode(
    message: LobbyMessage,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.peer !== undefined) {
      Peer.encode(message.peer, writer.uint32(10).fork()).ldelim();
    }
    if (message.message !== undefined) {
      writer.uint32(18).string(message.message);
    }
    if (message.metadata !== undefined) {
      Metadata.encode(message.metadata, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LobbyMessage {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLobbyMessage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.peer = Peer.decode(reader, reader.uint32());
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.metadata = Metadata.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LobbyMessage {
    return {
      peer: isSet(object.peer) ? Peer.fromJSON(object.peer) : undefined,
      message: isSet(object.message) ? String(object.message) : undefined,
      metadata: isSet(object.metadata)
        ? Metadata.fromJSON(object.metadata)
        : undefined,
    };
  },

  toJSON(message: LobbyMessage): unknown {
    const obj: any = {};
    message.peer !== undefined &&
      (obj.peer = message.peer ? Peer.toJSON(message.peer) : undefined);
    message.message !== undefined && (obj.message = message.message);
    message.metadata !== undefined &&
      (obj.metadata = message.metadata
        ? Metadata.toJSON(message.metadata)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LobbyMessage>, I>>(
    object: I
  ): LobbyMessage {
    const message = createBaseLobbyMessage();
    message.peer =
      object.peer !== undefined && object.peer !== null
        ? Peer.fromPartial(object.peer)
        : undefined;
    message.message = object.message ?? undefined;
    message.metadata =
      object.metadata !== undefined && object.metadata !== null
        ? Metadata.fromPartial(object.metadata)
        : undefined;
    return message;
  },
};

function createBaseVisibilityRequest(): VisibilityRequest {
  return {
    sName: "",
    publicKey: Buffer.alloc(0),
    visibility: VisibilityRequest_Visibility.VISIBILITY_UNSPECIFIED,
  };
}

export const VisibilityRequest = {
  encode(
    message: VisibilityRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.sName !== "") {
      writer.uint32(10).string(message.sName);
    }
    if (message.publicKey.length !== 0) {
      writer.uint32(18).bytes(message.publicKey);
    }
    if (
      message.visibility !== VisibilityRequest_Visibility.VISIBILITY_UNSPECIFIED
    ) {
      writer
        .uint32(24)
        .int32(visibilityRequest_VisibilityToNumber(message.visibility));
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VisibilityRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVisibilityRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sName = reader.string();
          break;
        case 2:
          message.publicKey = reader.bytes() as Buffer;
          break;
        case 3:
          message.visibility = visibilityRequest_VisibilityFromJSON(
            reader.int32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VisibilityRequest {
    return {
      sName: isSet(object.sName) ? String(object.sName) : "",
      publicKey: isSet(object.publicKey)
        ? Buffer.from(bytesFromBase64(object.publicKey))
        : Buffer.alloc(0),
      visibility: isSet(object.visibility)
        ? visibilityRequest_VisibilityFromJSON(object.visibility)
        : VisibilityRequest_Visibility.VISIBILITY_UNSPECIFIED,
    };
  },

  toJSON(message: VisibilityRequest): unknown {
    const obj: any = {};
    message.sName !== undefined && (obj.sName = message.sName);
    message.publicKey !== undefined &&
      (obj.publicKey = base64FromBytes(
        message.publicKey !== undefined ? message.publicKey : Buffer.alloc(0)
      ));
    message.visibility !== undefined &&
      (obj.visibility = visibilityRequest_VisibilityToJSON(message.visibility));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VisibilityRequest>, I>>(
    object: I
  ): VisibilityRequest {
    const message = createBaseVisibilityRequest();
    message.sName = object.sName ?? "";
    message.publicKey = object.publicKey ?? Buffer.alloc(0);
    message.visibility =
      object.visibility ?? VisibilityRequest_Visibility.VISIBILITY_UNSPECIFIED;
    return message;
  },
};

function createBaseVisibilityResponse(): VisibilityResponse {
  return {
    success: false,
    error: "",
    visibility: VisibilityResponse_Visibility.VISIBILITY_UNSPECIFIED,
  };
}

export const VisibilityResponse = {
  encode(
    message: VisibilityResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    if (message.error !== "") {
      writer.uint32(18).string(message.error);
    }
    if (
      message.visibility !==
      VisibilityResponse_Visibility.VISIBILITY_UNSPECIFIED
    ) {
      writer
        .uint32(24)
        .int32(visibilityResponse_VisibilityToNumber(message.visibility));
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VisibilityResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVisibilityResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        case 2:
          message.error = reader.string();
          break;
        case 3:
          message.visibility = visibilityResponse_VisibilityFromJSON(
            reader.int32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VisibilityResponse {
    return {
      success: isSet(object.success) ? Boolean(object.success) : false,
      error: isSet(object.error) ? String(object.error) : "",
      visibility: isSet(object.visibility)
        ? visibilityResponse_VisibilityFromJSON(object.visibility)
        : VisibilityResponse_Visibility.VISIBILITY_UNSPECIFIED,
    };
  },

  toJSON(message: VisibilityResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.error !== undefined && (obj.error = message.error);
    message.visibility !== undefined &&
      (obj.visibility = visibilityResponse_VisibilityToJSON(
        message.visibility
      ));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VisibilityResponse>, I>>(
    object: I
  ): VisibilityResponse {
    const message = createBaseVisibilityResponse();
    message.success = object.success ?? false;
    message.error = object.error ?? "";
    message.visibility =
      object.visibility ?? VisibilityResponse_Visibility.VISIBILITY_UNSPECIFIED;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (const byte of arr) {
    bin.push(String.fromCharCode(byte));
  }
  return btoa(bin.join(""));
}

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
