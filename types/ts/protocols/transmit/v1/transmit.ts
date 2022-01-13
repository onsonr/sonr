/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import {
  Direction,
  Payload,
  FileItem,
  directionFromJSON,
  directionToJSON,
} from "../../../common/v1/data";
import { Peer } from "../../../common/v1/core";

export const protobufPackage = "protocols.transmit.v1";

export interface Session {
  direction: Direction;
  from: Peer | undefined;
  to: Peer | undefined;
  payload: Payload | undefined;
  lastUpdated: number;
  items: SessionItem[];
  currentIndex: number;
  results: { [key: number]: boolean };
}

export interface Session_ResultsEntry {
  key: number;
  value: boolean;
}

export interface SessionItem {
  index: number;
  count: number;
  item: FileItem | undefined;
  written: number;
  size: number;
  totalSize: number;
  direction: Direction;
  path: string;
}

export interface SessionPayload {
  payload: Payload | undefined;
  direction: Direction;
}

function createBaseSession(): Session {
  return {
    direction: 0,
    from: undefined,
    to: undefined,
    payload: undefined,
    lastUpdated: 0,
    items: [],
    currentIndex: 0,
    results: {},
  };
}

export const Session = {
  encode(
    message: Session,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.direction !== 0) {
      writer.uint32(8).int32(message.direction);
    }
    if (message.from !== undefined) {
      Peer.encode(message.from, writer.uint32(18).fork()).ldelim();
    }
    if (message.to !== undefined) {
      Peer.encode(message.to, writer.uint32(26).fork()).ldelim();
    }
    if (message.payload !== undefined) {
      Payload.encode(message.payload, writer.uint32(34).fork()).ldelim();
    }
    if (message.lastUpdated !== 0) {
      writer.uint32(40).int64(message.lastUpdated);
    }
    for (const v of message.items) {
      SessionItem.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    if (message.currentIndex !== 0) {
      writer.uint32(56).int32(message.currentIndex);
    }
    Object.entries(message.results).forEach(([key, value]) => {
      Session_ResultsEntry.encode(
        { key: key as any, value },
        writer.uint32(66).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Session {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSession();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.direction = reader.int32() as any;
          break;
        case 2:
          message.from = Peer.decode(reader, reader.uint32());
          break;
        case 3:
          message.to = Peer.decode(reader, reader.uint32());
          break;
        case 4:
          message.payload = Payload.decode(reader, reader.uint32());
          break;
        case 5:
          message.lastUpdated = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.items.push(SessionItem.decode(reader, reader.uint32()));
          break;
        case 7:
          message.currentIndex = reader.int32();
          break;
        case 8:
          const entry8 = Session_ResultsEntry.decode(reader, reader.uint32());
          if (entry8.value !== undefined) {
            message.results[entry8.key] = entry8.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Session {
    return {
      direction: isSet(object.direction)
        ? directionFromJSON(object.direction)
        : 0,
      from: isSet(object.from) ? Peer.fromJSON(object.from) : undefined,
      to: isSet(object.to) ? Peer.fromJSON(object.to) : undefined,
      payload: isSet(object.payload)
        ? Payload.fromJSON(object.payload)
        : undefined,
      lastUpdated: isSet(object.lastUpdated) ? Number(object.lastUpdated) : 0,
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => SessionItem.fromJSON(e))
        : [],
      currentIndex: isSet(object.currentIndex)
        ? Number(object.currentIndex)
        : 0,
      results: isObject(object.results)
        ? Object.entries(object.results).reduce<{ [key: number]: boolean }>(
            (acc, [key, value]) => {
              acc[Number(key)] = Boolean(value);
              return acc;
            },
            {}
          )
        : {},
    };
  },

  toJSON(message: Session): unknown {
    const obj: any = {};
    message.direction !== undefined &&
      (obj.direction = directionToJSON(message.direction));
    message.from !== undefined &&
      (obj.from = message.from ? Peer.toJSON(message.from) : undefined);
    message.to !== undefined &&
      (obj.to = message.to ? Peer.toJSON(message.to) : undefined);
    message.payload !== undefined &&
      (obj.payload = message.payload
        ? Payload.toJSON(message.payload)
        : undefined);
    message.lastUpdated !== undefined &&
      (obj.lastUpdated = Math.round(message.lastUpdated));
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? SessionItem.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.currentIndex !== undefined &&
      (obj.currentIndex = Math.round(message.currentIndex));
    obj.results = {};
    if (message.results) {
      Object.entries(message.results).forEach(([k, v]) => {
        obj.results[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Session>, I>>(object: I): Session {
    const message = createBaseSession();
    message.direction = object.direction ?? 0;
    message.from =
      object.from !== undefined && object.from !== null
        ? Peer.fromPartial(object.from)
        : undefined;
    message.to =
      object.to !== undefined && object.to !== null
        ? Peer.fromPartial(object.to)
        : undefined;
    message.payload =
      object.payload !== undefined && object.payload !== null
        ? Payload.fromPartial(object.payload)
        : undefined;
    message.lastUpdated = object.lastUpdated ?? 0;
    message.items = object.items?.map((e) => SessionItem.fromPartial(e)) || [];
    message.currentIndex = object.currentIndex ?? 0;
    message.results = Object.entries(object.results ?? {}).reduce<{
      [key: number]: boolean;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[Number(key)] = Boolean(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseSession_ResultsEntry(): Session_ResultsEntry {
  return { key: 0, value: false };
}

export const Session_ResultsEntry = {
  encode(
    message: Session_ResultsEntry,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.key !== 0) {
      writer.uint32(8).int32(message.key);
    }
    if (message.value === true) {
      writer.uint32(16).bool(message.value);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): Session_ResultsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSession_ResultsEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.int32();
          break;
        case 2:
          message.value = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Session_ResultsEntry {
    return {
      key: isSet(object.key) ? Number(object.key) : 0,
      value: isSet(object.value) ? Boolean(object.value) : false,
    };
  },

  toJSON(message: Session_ResultsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = Math.round(message.key));
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Session_ResultsEntry>, I>>(
    object: I
  ): Session_ResultsEntry {
    const message = createBaseSession_ResultsEntry();
    message.key = object.key ?? 0;
    message.value = object.value ?? false;
    return message;
  },
};

function createBaseSessionItem(): SessionItem {
  return {
    index: 0,
    count: 0,
    item: undefined,
    written: 0,
    size: 0,
    totalSize: 0,
    direction: 0,
    path: "",
  };
}

export const SessionItem = {
  encode(
    message: SessionItem,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.index !== 0) {
      writer.uint32(8).int32(message.index);
    }
    if (message.count !== 0) {
      writer.uint32(16).int32(message.count);
    }
    if (message.item !== undefined) {
      FileItem.encode(message.item, writer.uint32(26).fork()).ldelim();
    }
    if (message.written !== 0) {
      writer.uint32(32).int64(message.written);
    }
    if (message.size !== 0) {
      writer.uint32(40).int64(message.size);
    }
    if (message.totalSize !== 0) {
      writer.uint32(48).int64(message.totalSize);
    }
    if (message.direction !== 0) {
      writer.uint32(56).int32(message.direction);
    }
    if (message.path !== "") {
      writer.uint32(66).string(message.path);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SessionItem {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSessionItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.int32();
          break;
        case 2:
          message.count = reader.int32();
          break;
        case 3:
          message.item = FileItem.decode(reader, reader.uint32());
          break;
        case 4:
          message.written = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.size = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.totalSize = longToNumber(reader.int64() as Long);
          break;
        case 7:
          message.direction = reader.int32() as any;
          break;
        case 8:
          message.path = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SessionItem {
    return {
      index: isSet(object.index) ? Number(object.index) : 0,
      count: isSet(object.count) ? Number(object.count) : 0,
      item: isSet(object.item) ? FileItem.fromJSON(object.item) : undefined,
      written: isSet(object.written) ? Number(object.written) : 0,
      size: isSet(object.size) ? Number(object.size) : 0,
      totalSize: isSet(object.totalSize) ? Number(object.totalSize) : 0,
      direction: isSet(object.direction)
        ? directionFromJSON(object.direction)
        : 0,
      path: isSet(object.path) ? String(object.path) : "",
    };
  },

  toJSON(message: SessionItem): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = Math.round(message.index));
    message.count !== undefined && (obj.count = Math.round(message.count));
    message.item !== undefined &&
      (obj.item = message.item ? FileItem.toJSON(message.item) : undefined);
    message.written !== undefined &&
      (obj.written = Math.round(message.written));
    message.size !== undefined && (obj.size = Math.round(message.size));
    message.totalSize !== undefined &&
      (obj.totalSize = Math.round(message.totalSize));
    message.direction !== undefined &&
      (obj.direction = directionToJSON(message.direction));
    message.path !== undefined && (obj.path = message.path);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SessionItem>, I>>(
    object: I
  ): SessionItem {
    const message = createBaseSessionItem();
    message.index = object.index ?? 0;
    message.count = object.count ?? 0;
    message.item =
      object.item !== undefined && object.item !== null
        ? FileItem.fromPartial(object.item)
        : undefined;
    message.written = object.written ?? 0;
    message.size = object.size ?? 0;
    message.totalSize = object.totalSize ?? 0;
    message.direction = object.direction ?? 0;
    message.path = object.path ?? "";
    return message;
  },
};

function createBaseSessionPayload(): SessionPayload {
  return { payload: undefined, direction: 0 };
}

export const SessionPayload = {
  encode(
    message: SessionPayload,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.payload !== undefined) {
      Payload.encode(message.payload, writer.uint32(10).fork()).ldelim();
    }
    if (message.direction !== 0) {
      writer.uint32(16).int32(message.direction);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SessionPayload {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSessionPayload();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.payload = Payload.decode(reader, reader.uint32());
          break;
        case 2:
          message.direction = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SessionPayload {
    return {
      payload: isSet(object.payload)
        ? Payload.fromJSON(object.payload)
        : undefined,
      direction: isSet(object.direction)
        ? directionFromJSON(object.direction)
        : 0,
    };
  },

  toJSON(message: SessionPayload): unknown {
    const obj: any = {};
    message.payload !== undefined &&
      (obj.payload = message.payload
        ? Payload.toJSON(message.payload)
        : undefined);
    message.direction !== undefined &&
      (obj.direction = directionToJSON(message.direction));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SessionPayload>, I>>(
    object: I
  ): SessionPayload {
    const message = createBaseSessionPayload();
    message.payload =
      object.payload !== undefined && object.payload !== null
        ? Payload.fromPartial(object.payload)
        : undefined;
    message.direction = object.direction ?? 0;
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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
