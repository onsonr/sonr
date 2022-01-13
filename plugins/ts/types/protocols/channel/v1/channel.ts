/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "protocols.channel.v1";

/** EventType is the type of event being performed on a channel. */
export enum EventType {
  /** EVENT_TYPE_UNSPECIFIED - EventTypeUnspecified is the default value. */
  EVENT_TYPE_UNSPECIFIED = "EVENT_TYPE_UNSPECIFIED",
  /** EVENT_TYPE_GET - EventTypeGet is a get event. */
  EVENT_TYPE_GET = "EVENT_TYPE_GET",
  /** EVENT_TYPE_SET - EventTypeSet is a set event. */
  EVENT_TYPE_SET = "EVENT_TYPE_SET",
  /** EVENT_TYPE_DELETE - EventTypeDelete is a delete event. */
  EVENT_TYPE_DELETE = "EVENT_TYPE_DELETE",
  /** EVENT_TYPE_PUT - EventTypePut is a put event. */
  EVENT_TYPE_PUT = "EVENT_TYPE_PUT",
  /** EVENT_TYPE_SYNC - EventTypeSync is a sync event. */
  EVENT_TYPE_SYNC = "EVENT_TYPE_SYNC",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function eventTypeFromJSON(object: any): EventType {
  switch (object) {
    case 0:
    case "EVENT_TYPE_UNSPECIFIED":
      return EventType.EVENT_TYPE_UNSPECIFIED;
    case 1:
    case "EVENT_TYPE_GET":
      return EventType.EVENT_TYPE_GET;
    case 2:
    case "EVENT_TYPE_SET":
      return EventType.EVENT_TYPE_SET;
    case 3:
    case "EVENT_TYPE_DELETE":
      return EventType.EVENT_TYPE_DELETE;
    case 4:
    case "EVENT_TYPE_PUT":
      return EventType.EVENT_TYPE_PUT;
    case 5:
    case "EVENT_TYPE_SYNC":
      return EventType.EVENT_TYPE_SYNC;
    case -1:
    case "UNRECOGNIZED":
    default:
      return EventType.UNRECOGNIZED;
  }
}

export function eventTypeToJSON(object: EventType): string {
  switch (object) {
    case EventType.EVENT_TYPE_UNSPECIFIED:
      return "EVENT_TYPE_UNSPECIFIED";
    case EventType.EVENT_TYPE_GET:
      return "EVENT_TYPE_GET";
    case EventType.EVENT_TYPE_SET:
      return "EVENT_TYPE_SET";
    case EventType.EVENT_TYPE_DELETE:
      return "EVENT_TYPE_DELETE";
    case EventType.EVENT_TYPE_PUT:
      return "EVENT_TYPE_PUT";
    case EventType.EVENT_TYPE_SYNC:
      return "EVENT_TYPE_SYNC";
    default:
      return "UNKNOWN";
  }
}

export function eventTypeToNumber(object: EventType): number {
  switch (object) {
    case EventType.EVENT_TYPE_UNSPECIFIED:
      return 0;
    case EventType.EVENT_TYPE_GET:
      return 1;
    case EventType.EVENT_TYPE_SET:
      return 2;
    case EventType.EVENT_TYPE_DELETE:
      return 3;
    case EventType.EVENT_TYPE_PUT:
      return 4;
    case EventType.EVENT_TYPE_SYNC:
      return 5;
    default:
      return 0;
  }
}

/** Event is the base event type for all channel events. */
export interface Event {
  /** Peer is the peer that originated the event. */
  peer: string;
  /** Type is the type of event being performed on a channel. */
  type: EventType;
  /** Entry is the entry being modified in the Store. */
  entry?: StoreEntry;
  /** Store is the store being operated on. */
  store?: Store;
}

/** Store is a disk based key-value store for channel data. */
export interface Store {
  /** Data is the data being stored. */
  data: { [key: string]: StoreEntry };
  /** Capacity is the maximum number of entries that can be stored. */
  capacity: number;
  /** Modified is the last time the store was modified. */
  modified: number;
  /** TTL is the time to live for entries in the store. */
  ttl: number;
}

export interface Store_DataEntry {
  key: string;
  value?: StoreEntry;
}

/** StoreEntry is the data being stored in the Store. */
export interface StoreEntry {
  /** Peer is the peer that originated the event. */
  peer: string;
  /** Key is the key being modified in the Store. */
  key: string;
  /** Value is the value being modified in the Store. */
  value: Buffer;
  /** Expiration is the expiration time for the entry. */
  expiration: number;
  /** Created is the time the entry was created. */
  created: number;
  /** Modified is the time the entry was last modified. */
  modified: number;
}

function createBaseEvent(): Event {
  return {
    peer: "",
    type: EventType.EVENT_TYPE_UNSPECIFIED,
    entry: undefined,
    store: undefined,
  };
}

export const Event = {
  encode(message: Event, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.peer !== "") {
      writer.uint32(10).string(message.peer);
    }
    if (message.type !== EventType.EVENT_TYPE_UNSPECIFIED) {
      writer.uint32(16).int32(eventTypeToNumber(message.type));
    }
    if (message.entry !== undefined) {
      StoreEntry.encode(message.entry, writer.uint32(26).fork()).ldelim();
    }
    if (message.store !== undefined) {
      Store.encode(message.store, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Event {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.peer = reader.string();
          break;
        case 2:
          message.type = eventTypeFromJSON(reader.int32());
          break;
        case 3:
          message.entry = StoreEntry.decode(reader, reader.uint32());
          break;
        case 4:
          message.store = Store.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Event {
    return {
      peer: isSet(object.peer) ? String(object.peer) : "",
      type: isSet(object.type)
        ? eventTypeFromJSON(object.type)
        : EventType.EVENT_TYPE_UNSPECIFIED,
      entry: isSet(object.entry)
        ? StoreEntry.fromJSON(object.entry)
        : undefined,
      store: isSet(object.store) ? Store.fromJSON(object.store) : undefined,
    };
  },

  toJSON(message: Event): unknown {
    const obj: any = {};
    message.peer !== undefined && (obj.peer = message.peer);
    message.type !== undefined && (obj.type = eventTypeToJSON(message.type));
    message.entry !== undefined &&
      (obj.entry = message.entry
        ? StoreEntry.toJSON(message.entry)
        : undefined);
    message.store !== undefined &&
      (obj.store = message.store ? Store.toJSON(message.store) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Event>, I>>(object: I): Event {
    const message = createBaseEvent();
    message.peer = object.peer ?? "";
    message.type = object.type ?? EventType.EVENT_TYPE_UNSPECIFIED;
    message.entry =
      object.entry !== undefined && object.entry !== null
        ? StoreEntry.fromPartial(object.entry)
        : undefined;
    message.store =
      object.store !== undefined && object.store !== null
        ? Store.fromPartial(object.store)
        : undefined;
    return message;
  },
};

function createBaseStore(): Store {
  return { data: {}, capacity: 0, modified: 0, ttl: 0 };
}

export const Store = {
  encode(message: Store, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    Object.entries(message.data).forEach(([key, value]) => {
      Store_DataEntry.encode(
        { key: key as any, value },
        writer.uint32(10).fork()
      ).ldelim();
    });
    if (message.capacity !== 0) {
      writer.uint32(16).int32(message.capacity);
    }
    if (message.modified !== 0) {
      writer.uint32(24).int64(message.modified);
    }
    if (message.ttl !== 0) {
      writer.uint32(32).int64(message.ttl);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Store {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStore();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          const entry1 = Store_DataEntry.decode(reader, reader.uint32());
          if (entry1.value !== undefined) {
            message.data[entry1.key] = entry1.value;
          }
          break;
        case 2:
          message.capacity = reader.int32();
          break;
        case 3:
          message.modified = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.ttl = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Store {
    return {
      data: isObject(object.data)
        ? Object.entries(object.data).reduce<{ [key: string]: StoreEntry }>(
            (acc, [key, value]) => {
              acc[key] = StoreEntry.fromJSON(value);
              return acc;
            },
            {}
          )
        : {},
      capacity: isSet(object.capacity) ? Number(object.capacity) : 0,
      modified: isSet(object.modified) ? Number(object.modified) : 0,
      ttl: isSet(object.ttl) ? Number(object.ttl) : 0,
    };
  },

  toJSON(message: Store): unknown {
    const obj: any = {};
    obj.data = {};
    if (message.data) {
      Object.entries(message.data).forEach(([k, v]) => {
        obj.data[k] = StoreEntry.toJSON(v);
      });
    }
    message.capacity !== undefined &&
      (obj.capacity = Math.round(message.capacity));
    message.modified !== undefined &&
      (obj.modified = Math.round(message.modified));
    message.ttl !== undefined && (obj.ttl = Math.round(message.ttl));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Store>, I>>(object: I): Store {
    const message = createBaseStore();
    message.data = Object.entries(object.data ?? {}).reduce<{
      [key: string]: StoreEntry;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = StoreEntry.fromPartial(value);
      }
      return acc;
    }, {});
    message.capacity = object.capacity ?? 0;
    message.modified = object.modified ?? 0;
    message.ttl = object.ttl ?? 0;
    return message;
  },
};

function createBaseStore_DataEntry(): Store_DataEntry {
  return { key: "", value: undefined };
}

export const Store_DataEntry = {
  encode(
    message: Store_DataEntry,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      StoreEntry.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Store_DataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStore_DataEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = StoreEntry.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Store_DataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value)
        ? StoreEntry.fromJSON(object.value)
        : undefined,
    };
  },

  toJSON(message: Store_DataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value
        ? StoreEntry.toJSON(message.value)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Store_DataEntry>, I>>(
    object: I
  ): Store_DataEntry {
    const message = createBaseStore_DataEntry();
    message.key = object.key ?? "";
    message.value =
      object.value !== undefined && object.value !== null
        ? StoreEntry.fromPartial(object.value)
        : undefined;
    return message;
  },
};

function createBaseStoreEntry(): StoreEntry {
  return {
    peer: "",
    key: "",
    value: Buffer.alloc(0),
    expiration: 0,
    created: 0,
    modified: 0,
  };
}

export const StoreEntry = {
  encode(
    message: StoreEntry,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.peer !== "") {
      writer.uint32(10).string(message.peer);
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    if (message.value.length !== 0) {
      writer.uint32(26).bytes(message.value);
    }
    if (message.expiration !== 0) {
      writer.uint32(32).int64(message.expiration);
    }
    if (message.created !== 0) {
      writer.uint32(40).int64(message.created);
    }
    if (message.modified !== 0) {
      writer.uint32(48).int64(message.modified);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StoreEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStoreEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.peer = reader.string();
          break;
        case 2:
          message.key = reader.string();
          break;
        case 3:
          message.value = reader.bytes() as Buffer;
          break;
        case 4:
          message.expiration = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.modified = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoreEntry {
    return {
      peer: isSet(object.peer) ? String(object.peer) : "",
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value)
        ? Buffer.from(bytesFromBase64(object.value))
        : Buffer.alloc(0),
      expiration: isSet(object.expiration) ? Number(object.expiration) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      modified: isSet(object.modified) ? Number(object.modified) : 0,
    };
  },

  toJSON(message: StoreEntry): unknown {
    const obj: any = {};
    message.peer !== undefined && (obj.peer = message.peer);
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = base64FromBytes(
        message.value !== undefined ? message.value : Buffer.alloc(0)
      ));
    message.expiration !== undefined &&
      (obj.expiration = Math.round(message.expiration));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.modified !== undefined &&
      (obj.modified = Math.round(message.modified));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<StoreEntry>, I>>(
    object: I
  ): StoreEntry {
    const message = createBaseStoreEntry();
    message.peer = object.peer ?? "";
    message.key = object.key ?? "";
    message.value = object.value ?? Buffer.alloc(0);
    message.expiration = object.expiration ?? 0;
    message.created = object.created ?? 0;
    message.modified = object.modified ?? 0;
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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
