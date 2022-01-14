/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "protocols.channel.v1";

/** EventType is the type of event being performed on a channel. */
export enum ChannelEventType {
  /** CHANNEL_EVENT_TYPE_UNSPECIFIED - EventTypeUnspecified is the default value. */
  CHANNEL_EVENT_TYPE_UNSPECIFIED = "CHANNEL_EVENT_TYPE_UNSPECIFIED",
  /** CHANNEL_EVENT_TYPE_GET - EventTypeGet is a get event being performed on a channel record in the store. */
  CHANNEL_EVENT_TYPE_GET = "CHANNEL_EVENT_TYPE_GET",
  /** CHANNEL_EVENT_TYPE_SET - EventTypeSet is a set event on the record store. */
  CHANNEL_EVENT_TYPE_SET = "CHANNEL_EVENT_TYPE_SET",
  /** CHANNEL_EVENT_TYPE_DELETE - EventTypeDelete is a delete event on the record store. */
  CHANNEL_EVENT_TYPE_DELETE = "CHANNEL_EVENT_TYPE_DELETE",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function channelEventTypeFromJSON(object: any): ChannelEventType {
  switch (object) {
    case 0:
    case "CHANNEL_EVENT_TYPE_UNSPECIFIED":
      return ChannelEventType.CHANNEL_EVENT_TYPE_UNSPECIFIED;
    case 1:
    case "CHANNEL_EVENT_TYPE_GET":
      return ChannelEventType.CHANNEL_EVENT_TYPE_GET;
    case 2:
    case "CHANNEL_EVENT_TYPE_SET":
      return ChannelEventType.CHANNEL_EVENT_TYPE_SET;
    case 3:
    case "CHANNEL_EVENT_TYPE_DELETE":
      return ChannelEventType.CHANNEL_EVENT_TYPE_DELETE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ChannelEventType.UNRECOGNIZED;
  }
}

export function channelEventTypeToJSON(object: ChannelEventType): string {
  switch (object) {
    case ChannelEventType.CHANNEL_EVENT_TYPE_UNSPECIFIED:
      return "CHANNEL_EVENT_TYPE_UNSPECIFIED";
    case ChannelEventType.CHANNEL_EVENT_TYPE_GET:
      return "CHANNEL_EVENT_TYPE_GET";
    case ChannelEventType.CHANNEL_EVENT_TYPE_SET:
      return "CHANNEL_EVENT_TYPE_SET";
    case ChannelEventType.CHANNEL_EVENT_TYPE_DELETE:
      return "CHANNEL_EVENT_TYPE_DELETE";
    default:
      return "UNKNOWN";
  }
}

export function channelEventTypeToNumber(object: ChannelEventType): number {
  switch (object) {
    case ChannelEventType.CHANNEL_EVENT_TYPE_UNSPECIFIED:
      return 0;
    case ChannelEventType.CHANNEL_EVENT_TYPE_GET:
      return 1;
    case ChannelEventType.CHANNEL_EVENT_TYPE_SET:
      return 2;
    case ChannelEventType.CHANNEL_EVENT_TYPE_DELETE:
      return 3;
    default:
      return 0;
  }
}

/** ChannelEvent is the base event type for all channel events. */
export interface ChannelEvent {
  /** Owner is the peer that originated the event. */
  owner: string;
  /** Type is the type of event being performed on a channel. */
  type: ChannelEventType;
  /** Record is the entry being modified in the Store. */
  record?: ChannelStoreRecord;
  /** Metadata is the metadata associated with the event. */
  metadata: { [key: string]: string };
}

export interface ChannelEvent_MetadataEntry {
  key: string;
  value: string;
}

/** ChannelMessage is a message sent to a channel. */
export interface ChannelMessage {
  /** Owner is the peer that originated the message. */
  owner: string;
  /** Text is the message text. */
  text: string;
  /** Data is the data being sent. */
  data: Buffer;
  /** Metadata is the metadata associated with the message. */
  metadata: { [key: string]: string };
}

export interface ChannelMessage_MetadataEntry {
  key: string;
  value: string;
}

/** Store is a disk based key-value store for channel data. */
export interface ChannelStore {
  /** Entries is the data being stored. */
  entries: { [key: string]: ChannelStoreRecord };
  /** Capacity is the maximum number of entries that can be stored. */
  capacity: number;
  /** Modified is the last time the store was modified. */
  modified: number;
  /** TTL is the time to live for entries in the store. */
  ttl: number;
}

export interface ChannelStore_EntriesEntry {
  key: string;
  value?: ChannelStoreRecord;
}

/** ChannelStoreRecord is the data being stored in the ChannelStore. */
export interface ChannelStoreRecord {
  /** Owner is the peer that originated the event. */
  owner: string;
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

function createBaseChannelEvent(): ChannelEvent {
  return {
    owner: "",
    type: ChannelEventType.CHANNEL_EVENT_TYPE_UNSPECIFIED,
    record: undefined,
    metadata: {},
  };
}

export const ChannelEvent = {
  encode(
    message: ChannelEvent,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.type !== ChannelEventType.CHANNEL_EVENT_TYPE_UNSPECIFIED) {
      writer.uint32(16).int32(channelEventTypeToNumber(message.type));
    }
    if (message.record !== undefined) {
      ChannelStoreRecord.encode(
        message.record,
        writer.uint32(26).fork()
      ).ldelim();
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ChannelEvent_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ChannelEvent {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChannelEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.type = channelEventTypeFromJSON(reader.int32());
          break;
        case 3:
          message.record = ChannelStoreRecord.decode(reader, reader.uint32());
          break;
        case 4:
          const entry4 = ChannelEvent_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry4.value !== undefined) {
            message.metadata[entry4.key] = entry4.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ChannelEvent {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      type: isSet(object.type)
        ? channelEventTypeFromJSON(object.type)
        : ChannelEventType.CHANNEL_EVENT_TYPE_UNSPECIFIED,
      record: isSet(object.record)
        ? ChannelStoreRecord.fromJSON(object.record)
        : undefined,
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
    };
  },

  toJSON(message: ChannelEvent): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.type !== undefined &&
      (obj.type = channelEventTypeToJSON(message.type));
    message.record !== undefined &&
      (obj.record = message.record
        ? ChannelStoreRecord.toJSON(message.record)
        : undefined);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ChannelEvent>, I>>(
    object: I
  ): ChannelEvent {
    const message = createBaseChannelEvent();
    message.owner = object.owner ?? "";
    message.type =
      object.type ?? ChannelEventType.CHANNEL_EVENT_TYPE_UNSPECIFIED;
    message.record =
      object.record !== undefined && object.record !== null
        ? ChannelStoreRecord.fromPartial(object.record)
        : undefined;
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseChannelEvent_MetadataEntry(): ChannelEvent_MetadataEntry {
  return { key: "", value: "" };
}

export const ChannelEvent_MetadataEntry = {
  encode(
    message: ChannelEvent_MetadataEntry,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ChannelEvent_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChannelEvent_MetadataEntry();
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

  fromJSON(object: any): ChannelEvent_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: ChannelEvent_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ChannelEvent_MetadataEntry>, I>>(
    object: I
  ): ChannelEvent_MetadataEntry {
    const message = createBaseChannelEvent_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseChannelMessage(): ChannelMessage {
  return { owner: "", text: "", data: Buffer.alloc(0), metadata: {} };
}

export const ChannelMessage = {
  encode(
    message: ChannelMessage,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.text !== "") {
      writer.uint32(18).string(message.text);
    }
    if (message.data.length !== 0) {
      writer.uint32(26).bytes(message.data);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ChannelMessage_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ChannelMessage {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChannelMessage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.text = reader.string();
          break;
        case 3:
          message.data = reader.bytes() as Buffer;
          break;
        case 4:
          const entry4 = ChannelMessage_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry4.value !== undefined) {
            message.metadata[entry4.key] = entry4.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ChannelMessage {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      text: isSet(object.text) ? String(object.text) : "",
      data: isSet(object.data)
        ? Buffer.from(bytesFromBase64(object.data))
        : Buffer.alloc(0),
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
    };
  },

  toJSON(message: ChannelMessage): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.text !== undefined && (obj.text = message.text);
    message.data !== undefined &&
      (obj.data = base64FromBytes(
        message.data !== undefined ? message.data : Buffer.alloc(0)
      ));
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ChannelMessage>, I>>(
    object: I
  ): ChannelMessage {
    const message = createBaseChannelMessage();
    message.owner = object.owner ?? "";
    message.text = object.text ?? "";
    message.data = object.data ?? Buffer.alloc(0);
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseChannelMessage_MetadataEntry(): ChannelMessage_MetadataEntry {
  return { key: "", value: "" };
}

export const ChannelMessage_MetadataEntry = {
  encode(
    message: ChannelMessage_MetadataEntry,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ChannelMessage_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChannelMessage_MetadataEntry();
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

  fromJSON(object: any): ChannelMessage_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: ChannelMessage_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ChannelMessage_MetadataEntry>, I>>(
    object: I
  ): ChannelMessage_MetadataEntry {
    const message = createBaseChannelMessage_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseChannelStore(): ChannelStore {
  return { entries: {}, capacity: 0, modified: 0, ttl: 0 };
}

export const ChannelStore = {
  encode(
    message: ChannelStore,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    Object.entries(message.entries).forEach(([key, value]) => {
      ChannelStore_EntriesEntry.encode(
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

  decode(input: _m0.Reader | Uint8Array, length?: number): ChannelStore {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChannelStore();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          const entry1 = ChannelStore_EntriesEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry1.value !== undefined) {
            message.entries[entry1.key] = entry1.value;
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

  fromJSON(object: any): ChannelStore {
    return {
      entries: isObject(object.entries)
        ? Object.entries(object.entries).reduce<{
            [key: string]: ChannelStoreRecord;
          }>((acc, [key, value]) => {
            acc[key] = ChannelStoreRecord.fromJSON(value);
            return acc;
          }, {})
        : {},
      capacity: isSet(object.capacity) ? Number(object.capacity) : 0,
      modified: isSet(object.modified) ? Number(object.modified) : 0,
      ttl: isSet(object.ttl) ? Number(object.ttl) : 0,
    };
  },

  toJSON(message: ChannelStore): unknown {
    const obj: any = {};
    obj.entries = {};
    if (message.entries) {
      Object.entries(message.entries).forEach(([k, v]) => {
        obj.entries[k] = ChannelStoreRecord.toJSON(v);
      });
    }
    message.capacity !== undefined &&
      (obj.capacity = Math.round(message.capacity));
    message.modified !== undefined &&
      (obj.modified = Math.round(message.modified));
    message.ttl !== undefined && (obj.ttl = Math.round(message.ttl));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ChannelStore>, I>>(
    object: I
  ): ChannelStore {
    const message = createBaseChannelStore();
    message.entries = Object.entries(object.entries ?? {}).reduce<{
      [key: string]: ChannelStoreRecord;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = ChannelStoreRecord.fromPartial(value);
      }
      return acc;
    }, {});
    message.capacity = object.capacity ?? 0;
    message.modified = object.modified ?? 0;
    message.ttl = object.ttl ?? 0;
    return message;
  },
};

function createBaseChannelStore_EntriesEntry(): ChannelStore_EntriesEntry {
  return { key: "", value: undefined };
}

export const ChannelStore_EntriesEntry = {
  encode(
    message: ChannelStore_EntriesEntry,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      ChannelStoreRecord.encode(
        message.value,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ChannelStore_EntriesEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChannelStore_EntriesEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = ChannelStoreRecord.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ChannelStore_EntriesEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value)
        ? ChannelStoreRecord.fromJSON(object.value)
        : undefined,
    };
  },

  toJSON(message: ChannelStore_EntriesEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value
        ? ChannelStoreRecord.toJSON(message.value)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ChannelStore_EntriesEntry>, I>>(
    object: I
  ): ChannelStore_EntriesEntry {
    const message = createBaseChannelStore_EntriesEntry();
    message.key = object.key ?? "";
    message.value =
      object.value !== undefined && object.value !== null
        ? ChannelStoreRecord.fromPartial(object.value)
        : undefined;
    return message;
  },
};

function createBaseChannelStoreRecord(): ChannelStoreRecord {
  return {
    owner: "",
    key: "",
    value: Buffer.alloc(0),
    expiration: 0,
    created: 0,
    modified: 0,
  };
}

export const ChannelStoreRecord = {
  encode(
    message: ChannelStoreRecord,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
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

  decode(input: _m0.Reader | Uint8Array, length?: number): ChannelStoreRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChannelStoreRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
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

  fromJSON(object: any): ChannelStoreRecord {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value)
        ? Buffer.from(bytesFromBase64(object.value))
        : Buffer.alloc(0),
      expiration: isSet(object.expiration) ? Number(object.expiration) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      modified: isSet(object.modified) ? Number(object.modified) : 0,
    };
  },

  toJSON(message: ChannelStoreRecord): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
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

  fromPartial<I extends Exact<DeepPartial<ChannelStoreRecord>, I>>(
    object: I
  ): ChannelStoreRecord {
    const message = createBaseChannelStoreRecord();
    message.owner = object.owner ?? "";
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
