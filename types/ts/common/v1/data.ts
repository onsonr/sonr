/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { MIME, Profile } from "../../common/v1/core";

export const protobufPackage = "common.v1";

export enum Direction {
  DIRECTION_UNSPECIFIED = 0,
  DIRECTION_INCOMING = 1,
  DIRECTION_OUTGOING = 2,
  UNRECOGNIZED = -1,
}

export function directionFromJSON(object: any): Direction {
  switch (object) {
    case 0:
    case "DIRECTION_UNSPECIFIED":
      return Direction.DIRECTION_UNSPECIFIED;
    case 1:
    case "DIRECTION_INCOMING":
      return Direction.DIRECTION_INCOMING;
    case 2:
    case "DIRECTION_OUTGOING":
      return Direction.DIRECTION_OUTGOING;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Direction.UNRECOGNIZED;
  }
}

export function directionToJSON(object: Direction): string {
  switch (object) {
    case Direction.DIRECTION_UNSPECIFIED:
      return "DIRECTION_UNSPECIFIED";
    case Direction.DIRECTION_INCOMING:
      return "DIRECTION_INCOMING";
    case Direction.DIRECTION_OUTGOING:
      return "DIRECTION_OUTGOING";
    default:
      return "UNKNOWN";
  }
}

/** For Transfer File Payload */
export interface FileItem {
  /** Standard Mime Type */
  mime: MIME | undefined;
  /** File Name without Path */
  name: string;
  /** File Location */
  path: string;
  /** File Size in Bytes */
  size: number;
  /** Thumbnail of File */
  thumbnail: Thumbnail | undefined;
  /** Last Modified Time in Seconds */
  lastModified: number;
}

/** Payload is Data thats being Passed */
export interface Payload {
  /** Payload Items */
  items: Payload_Item[];
  /** PROFILE: General Sender Info */
  owner: Profile | undefined;
  /** Payload Size in Bytes */
  size: number;
  /** Payload Creation Time in Seconds */
  createdAt: number;
}

/** Item in Payload */
export interface Payload_Item {
  /** MIME of the Item */
  mime: MIME | undefined;
  /** Size of the Item in Bytes */
  size: number;
  /** FILE: File Item */
  file: FileItem | undefined;
  /** URL: Url Item */
  url: string | undefined;
  /** MESSAGE: Message Item */
  message: string | undefined;
  /** Thumbnail of the Item */
  thumbnail: Thumbnail | undefined;
}

/** PayloadList is a list of Payload.Item's for Persistent Store */
export interface PayloadList {
  /** Payload List */
  payloads: Payload[];
  /** Key of the Payload List */
  key: string;
  /** Last Modified Time in Seconds */
  lastModified: number;
}

/** SupplyItem is an item supplied to be a payload */
export interface SupplyItem {
  /** Supply Path */
  path: string;
  /** Supply Path of the Thumbnail */
  thumbnail?: Uint8Array | undefined;
}

/** Thumbnail of File */
export interface Thumbnail {
  /** Thumbnail Buffer */
  buffer: Uint8Array;
  /** Mime Type */
  mime: MIME | undefined;
}

function createBaseFileItem(): FileItem {
  return {
    mime: undefined,
    name: "",
    path: "",
    size: 0,
    thumbnail: undefined,
    lastModified: 0,
  };
}

export const FileItem = {
  encode(
    message: FileItem,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.mime !== undefined) {
      MIME.encode(message.mime, writer.uint32(10).fork()).ldelim();
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.path !== "") {
      writer.uint32(26).string(message.path);
    }
    if (message.size !== 0) {
      writer.uint32(32).int64(message.size);
    }
    if (message.thumbnail !== undefined) {
      Thumbnail.encode(message.thumbnail, writer.uint32(42).fork()).ldelim();
    }
    if (message.lastModified !== 0) {
      writer.uint32(48).int64(message.lastModified);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FileItem {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFileItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mime = MIME.decode(reader, reader.uint32());
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.path = reader.string();
          break;
        case 4:
          message.size = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.thumbnail = Thumbnail.decode(reader, reader.uint32());
          break;
        case 6:
          message.lastModified = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FileItem {
    return {
      mime: isSet(object.mime) ? MIME.fromJSON(object.mime) : undefined,
      name: isSet(object.name) ? String(object.name) : "",
      path: isSet(object.path) ? String(object.path) : "",
      size: isSet(object.size) ? Number(object.size) : 0,
      thumbnail: isSet(object.thumbnail)
        ? Thumbnail.fromJSON(object.thumbnail)
        : undefined,
      lastModified: isSet(object.lastModified)
        ? Number(object.lastModified)
        : 0,
    };
  },

  toJSON(message: FileItem): unknown {
    const obj: any = {};
    message.mime !== undefined &&
      (obj.mime = message.mime ? MIME.toJSON(message.mime) : undefined);
    message.name !== undefined && (obj.name = message.name);
    message.path !== undefined && (obj.path = message.path);
    message.size !== undefined && (obj.size = Math.round(message.size));
    message.thumbnail !== undefined &&
      (obj.thumbnail = message.thumbnail
        ? Thumbnail.toJSON(message.thumbnail)
        : undefined);
    message.lastModified !== undefined &&
      (obj.lastModified = Math.round(message.lastModified));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FileItem>, I>>(object: I): FileItem {
    const message = createBaseFileItem();
    message.mime =
      object.mime !== undefined && object.mime !== null
        ? MIME.fromPartial(object.mime)
        : undefined;
    message.name = object.name ?? "";
    message.path = object.path ?? "";
    message.size = object.size ?? 0;
    message.thumbnail =
      object.thumbnail !== undefined && object.thumbnail !== null
        ? Thumbnail.fromPartial(object.thumbnail)
        : undefined;
    message.lastModified = object.lastModified ?? 0;
    return message;
  },
};

function createBasePayload(): Payload {
  return { items: [], owner: undefined, size: 0, createdAt: 0 };
}

export const Payload = {
  encode(
    message: Payload,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      Payload_Item.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.owner !== undefined) {
      Profile.encode(message.owner, writer.uint32(18).fork()).ldelim();
    }
    if (message.size !== 0) {
      writer.uint32(24).int64(message.size);
    }
    if (message.createdAt !== 0) {
      writer.uint32(32).int64(message.createdAt);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Payload {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePayload();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(Payload_Item.decode(reader, reader.uint32()));
          break;
        case 2:
          message.owner = Profile.decode(reader, reader.uint32());
          break;
        case 3:
          message.size = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.createdAt = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Payload {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => Payload_Item.fromJSON(e))
        : [],
      owner: isSet(object.owner) ? Profile.fromJSON(object.owner) : undefined,
      size: isSet(object.size) ? Number(object.size) : 0,
      createdAt: isSet(object.createdAt) ? Number(object.createdAt) : 0,
    };
  },

  toJSON(message: Payload): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? Payload_Item.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.owner !== undefined &&
      (obj.owner = message.owner ? Profile.toJSON(message.owner) : undefined);
    message.size !== undefined && (obj.size = Math.round(message.size));
    message.createdAt !== undefined &&
      (obj.createdAt = Math.round(message.createdAt));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Payload>, I>>(object: I): Payload {
    const message = createBasePayload();
    message.items = object.items?.map((e) => Payload_Item.fromPartial(e)) || [];
    message.owner =
      object.owner !== undefined && object.owner !== null
        ? Profile.fromPartial(object.owner)
        : undefined;
    message.size = object.size ?? 0;
    message.createdAt = object.createdAt ?? 0;
    return message;
  },
};

function createBasePayload_Item(): Payload_Item {
  return {
    mime: undefined,
    size: 0,
    file: undefined,
    url: undefined,
    message: undefined,
    thumbnail: undefined,
  };
}

export const Payload_Item = {
  encode(
    message: Payload_Item,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.mime !== undefined) {
      MIME.encode(message.mime, writer.uint32(10).fork()).ldelim();
    }
    if (message.size !== 0) {
      writer.uint32(16).int64(message.size);
    }
    if (message.file !== undefined) {
      FileItem.encode(message.file, writer.uint32(26).fork()).ldelim();
    }
    if (message.url !== undefined) {
      writer.uint32(34).string(message.url);
    }
    if (message.message !== undefined) {
      writer.uint32(42).string(message.message);
    }
    if (message.thumbnail !== undefined) {
      Thumbnail.encode(message.thumbnail, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Payload_Item {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePayload_Item();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mime = MIME.decode(reader, reader.uint32());
          break;
        case 2:
          message.size = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.file = FileItem.decode(reader, reader.uint32());
          break;
        case 4:
          message.url = reader.string();
          break;
        case 5:
          message.message = reader.string();
          break;
        case 6:
          message.thumbnail = Thumbnail.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Payload_Item {
    return {
      mime: isSet(object.mime) ? MIME.fromJSON(object.mime) : undefined,
      size: isSet(object.size) ? Number(object.size) : 0,
      file: isSet(object.file) ? FileItem.fromJSON(object.file) : undefined,
      url: isSet(object.url) ? String(object.url) : undefined,
      message: isSet(object.message) ? String(object.message) : undefined,
      thumbnail: isSet(object.thumbnail)
        ? Thumbnail.fromJSON(object.thumbnail)
        : undefined,
    };
  },

  toJSON(message: Payload_Item): unknown {
    const obj: any = {};
    message.mime !== undefined &&
      (obj.mime = message.mime ? MIME.toJSON(message.mime) : undefined);
    message.size !== undefined && (obj.size = Math.round(message.size));
    message.file !== undefined &&
      (obj.file = message.file ? FileItem.toJSON(message.file) : undefined);
    message.url !== undefined && (obj.url = message.url);
    message.message !== undefined && (obj.message = message.message);
    message.thumbnail !== undefined &&
      (obj.thumbnail = message.thumbnail
        ? Thumbnail.toJSON(message.thumbnail)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Payload_Item>, I>>(
    object: I
  ): Payload_Item {
    const message = createBasePayload_Item();
    message.mime =
      object.mime !== undefined && object.mime !== null
        ? MIME.fromPartial(object.mime)
        : undefined;
    message.size = object.size ?? 0;
    message.file =
      object.file !== undefined && object.file !== null
        ? FileItem.fromPartial(object.file)
        : undefined;
    message.url = object.url ?? undefined;
    message.message = object.message ?? undefined;
    message.thumbnail =
      object.thumbnail !== undefined && object.thumbnail !== null
        ? Thumbnail.fromPartial(object.thumbnail)
        : undefined;
    return message;
  },
};

function createBasePayloadList(): PayloadList {
  return { payloads: [], key: "", lastModified: 0 };
}

export const PayloadList = {
  encode(
    message: PayloadList,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.payloads) {
      Payload.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    if (message.lastModified !== 0) {
      writer.uint32(24).int64(message.lastModified);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PayloadList {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePayloadList();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.payloads.push(Payload.decode(reader, reader.uint32()));
          break;
        case 2:
          message.key = reader.string();
          break;
        case 3:
          message.lastModified = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PayloadList {
    return {
      payloads: Array.isArray(object?.payloads)
        ? object.payloads.map((e: any) => Payload.fromJSON(e))
        : [],
      key: isSet(object.key) ? String(object.key) : "",
      lastModified: isSet(object.lastModified)
        ? Number(object.lastModified)
        : 0,
    };
  },

  toJSON(message: PayloadList): unknown {
    const obj: any = {};
    if (message.payloads) {
      obj.payloads = message.payloads.map((e) =>
        e ? Payload.toJSON(e) : undefined
      );
    } else {
      obj.payloads = [];
    }
    message.key !== undefined && (obj.key = message.key);
    message.lastModified !== undefined &&
      (obj.lastModified = Math.round(message.lastModified));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PayloadList>, I>>(
    object: I
  ): PayloadList {
    const message = createBasePayloadList();
    message.payloads =
      object.payloads?.map((e) => Payload.fromPartial(e)) || [];
    message.key = object.key ?? "";
    message.lastModified = object.lastModified ?? 0;
    return message;
  },
};

function createBaseSupplyItem(): SupplyItem {
  return { path: "", thumbnail: undefined };
}

export const SupplyItem = {
  encode(
    message: SupplyItem,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.path !== "") {
      writer.uint32(10).string(message.path);
    }
    if (message.thumbnail !== undefined) {
      writer.uint32(18).bytes(message.thumbnail);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SupplyItem {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSupplyItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.path = reader.string();
          break;
        case 2:
          message.thumbnail = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SupplyItem {
    return {
      path: isSet(object.path) ? String(object.path) : "",
      thumbnail: isSet(object.thumbnail)
        ? bytesFromBase64(object.thumbnail)
        : undefined,
    };
  },

  toJSON(message: SupplyItem): unknown {
    const obj: any = {};
    message.path !== undefined && (obj.path = message.path);
    message.thumbnail !== undefined &&
      (obj.thumbnail =
        message.thumbnail !== undefined
          ? base64FromBytes(message.thumbnail)
          : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SupplyItem>, I>>(
    object: I
  ): SupplyItem {
    const message = createBaseSupplyItem();
    message.path = object.path ?? "";
    message.thumbnail = object.thumbnail ?? undefined;
    return message;
  },
};

function createBaseThumbnail(): Thumbnail {
  return { buffer: new Uint8Array(), mime: undefined };
}

export const Thumbnail = {
  encode(
    message: Thumbnail,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.buffer.length !== 0) {
      writer.uint32(10).bytes(message.buffer);
    }
    if (message.mime !== undefined) {
      MIME.encode(message.mime, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Thumbnail {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseThumbnail();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.buffer = reader.bytes();
          break;
        case 2:
          message.mime = MIME.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Thumbnail {
    return {
      buffer: isSet(object.buffer)
        ? bytesFromBase64(object.buffer)
        : new Uint8Array(),
      mime: isSet(object.mime) ? MIME.fromJSON(object.mime) : undefined,
    };
  },

  toJSON(message: Thumbnail): unknown {
    const obj: any = {};
    message.buffer !== undefined &&
      (obj.buffer = base64FromBytes(
        message.buffer !== undefined ? message.buffer : new Uint8Array()
      ));
    message.mime !== undefined &&
      (obj.mime = message.mime ? MIME.toJSON(message.mime) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Thumbnail>, I>>(
    object: I
  ): Thumbnail {
    const message = createBaseThumbnail();
    message.buffer = object.buffer ?? new Uint8Array();
    message.mime =
      object.mime !== undefined && object.mime !== null
        ? MIME.fromPartial(object.mime)
        : undefined;
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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
