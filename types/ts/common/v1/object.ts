/* eslint-disable */
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "common.v1";

/** BucketType is the type of a bucket. */
export enum BucketType {
  BUCKET_TYPE_UNSPECIFIED = 0,
  BUCKET_TYPE_APP = 1,
  BUCKET_TYPE_USER = 2,
  UNRECOGNIZED = -1,
}

export function bucketTypeFromJSON(object: any): BucketType {
  switch (object) {
    case 0:
    case "BUCKET_TYPE_UNSPECIFIED":
      return BucketType.BUCKET_TYPE_UNSPECIFIED;
    case 1:
    case "BUCKET_TYPE_APP":
      return BucketType.BUCKET_TYPE_APP;
    case 2:
    case "BUCKET_TYPE_USER":
      return BucketType.BUCKET_TYPE_USER;
    case -1:
    case "UNRECOGNIZED":
    default:
      return BucketType.UNRECOGNIZED;
  }
}

export function bucketTypeToJSON(object: BucketType): string {
  switch (object) {
    case BucketType.BUCKET_TYPE_UNSPECIFIED:
      return "BUCKET_TYPE_UNSPECIFIED";
    case BucketType.BUCKET_TYPE_APP:
      return "BUCKET_TYPE_APP";
    case BucketType.BUCKET_TYPE_USER:
      return "BUCKET_TYPE_USER";
    default:
      return "UNKNOWN";
  }
}

/** ObjectFieldType is the type of the field */
export enum ObjectFieldType {
  OBJECT_FIELD_TYPE_UNSPECIFIED = 0,
  OBJECT_FIELD_TYPE_STRING = 1,
  OBJECT_FIELD_TYPE_INT = 2,
  OBJECT_FIELD_TYPE_FLOAT = 3,
  OBJECT_FIELD_TYPE_BOOL = 4,
  OBJECT_FIELD_TYPE_DATETIME = 5,
  OBJECT_FIELD_TYPE_BLOB = 6,
  OBJECT_FIELD_TYPE_REFERENCE = 7,
  UNRECOGNIZED = -1,
}

export function objectFieldTypeFromJSON(object: any): ObjectFieldType {
  switch (object) {
    case 0:
    case "OBJECT_FIELD_TYPE_UNSPECIFIED":
      return ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED;
    case 1:
    case "OBJECT_FIELD_TYPE_STRING":
      return ObjectFieldType.OBJECT_FIELD_TYPE_STRING;
    case 2:
    case "OBJECT_FIELD_TYPE_INT":
      return ObjectFieldType.OBJECT_FIELD_TYPE_INT;
    case 3:
    case "OBJECT_FIELD_TYPE_FLOAT":
      return ObjectFieldType.OBJECT_FIELD_TYPE_FLOAT;
    case 4:
    case "OBJECT_FIELD_TYPE_BOOL":
      return ObjectFieldType.OBJECT_FIELD_TYPE_BOOL;
    case 5:
    case "OBJECT_FIELD_TYPE_DATETIME":
      return ObjectFieldType.OBJECT_FIELD_TYPE_DATETIME;
    case 6:
    case "OBJECT_FIELD_TYPE_BLOB":
      return ObjectFieldType.OBJECT_FIELD_TYPE_BLOB;
    case 7:
    case "OBJECT_FIELD_TYPE_REFERENCE":
      return ObjectFieldType.OBJECT_FIELD_TYPE_REFERENCE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ObjectFieldType.UNRECOGNIZED;
  }
}

export function objectFieldTypeToJSON(object: ObjectFieldType): string {
  switch (object) {
    case ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED:
      return "OBJECT_FIELD_TYPE_UNSPECIFIED";
    case ObjectFieldType.OBJECT_FIELD_TYPE_STRING:
      return "OBJECT_FIELD_TYPE_STRING";
    case ObjectFieldType.OBJECT_FIELD_TYPE_INT:
      return "OBJECT_FIELD_TYPE_INT";
    case ObjectFieldType.OBJECT_FIELD_TYPE_FLOAT:
      return "OBJECT_FIELD_TYPE_FLOAT";
    case ObjectFieldType.OBJECT_FIELD_TYPE_BOOL:
      return "OBJECT_FIELD_TYPE_BOOL";
    case ObjectFieldType.OBJECT_FIELD_TYPE_DATETIME:
      return "OBJECT_FIELD_TYPE_DATETIME";
    case ObjectFieldType.OBJECT_FIELD_TYPE_BLOB:
      return "OBJECT_FIELD_TYPE_BLOB";
    case ObjectFieldType.OBJECT_FIELD_TYPE_REFERENCE:
      return "OBJECT_FIELD_TYPE_REFERENCE";
    default:
      return "UNKNOWN";
  }
}

export interface Bucket {
  /** Label is human-readable name of the bucket. */
  label: string;
  /** Description is a human-readable description of the bucket. */
  description: string;
  /** Type is the kind of bucket for either App specific or User specific data. */
  type: BucketType;
  /** Did is the identifier of the bucket. */
  did: string;
  /** Objects are stored in a tree structure. */
  objects: Object[];
}

export interface Object {
  did: string;
  service: string;
  tags: string[];
  fields: { [key: string]: ObjectField };
}

export interface Object_FieldsEntry {
  key: string;
  value: ObjectField | undefined;
}

export interface ObjectField {
  /** Name is the name of the field */
  name: string;
  /** Type is the type of the field */
  type: ObjectFieldType;
  /** String is the value of the field */
  stringValue: string | undefined;
  /** Int is the value of the field */
  intValue: number | undefined;
  /** Bool is the value of the field */
  boolValue: boolean | undefined;
  /** Blob is the value of the field */
  blobValue: Uint8Array | undefined;
  /** Reference is the value of the field */
  referenceValue: string | undefined;
}

function createBaseBucket(): Bucket {
  return { label: "", description: "", type: 0, did: "", objects: [] };
}

export const Bucket = {
  encode(message: Bucket, writer: Writer = Writer.create()): Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.type !== 0) {
      writer.uint32(24).int32(message.type);
    }
    if (message.did !== "") {
      writer.uint32(34).string(message.did);
    }
    for (const v of message.objects) {
      Object.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Bucket {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBucket();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.type = reader.int32() as any;
          break;
        case 4:
          message.did = reader.string();
          break;
        case 5:
          message.objects.push(Object.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Bucket {
    return {
      label: isSet(object.label) ? String(object.label) : "",
      description: isSet(object.description) ? String(object.description) : "",
      type: isSet(object.type) ? bucketTypeFromJSON(object.type) : 0,
      did: isSet(object.did) ? String(object.did) : "",
      objects: Array.isArray(object?.objects)
        ? object.objects.map((e: any) => Object.fromJSON(e))
        : [],
    };
  },

  toJSON(message: Bucket): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.description !== undefined &&
      (obj.description = message.description);
    message.type !== undefined && (obj.type = bucketTypeToJSON(message.type));
    message.did !== undefined && (obj.did = message.did);
    if (message.objects) {
      obj.objects = message.objects.map((e) =>
        e ? Object.toJSON(e) : undefined
      );
    } else {
      obj.objects = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Bucket>, I>>(object: I): Bucket {
    const message = createBaseBucket();
    message.label = object.label ?? "";
    message.description = object.description ?? "";
    message.type = object.type ?? 0;
    message.did = object.did ?? "";
    message.objects = object.objects?.map((e) => Object.fromPartial(e)) || [];
    return message;
  },
};

function createBaseObject(): Object {
  return { did: "", service: "", tags: [], fields: {} };
}

export const Object = {
  encode(message: Object, writer: Writer = Writer.create()): Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    if (message.service !== "") {
      writer.uint32(18).string(message.service);
    }
    for (const v of message.tags) {
      writer.uint32(26).string(v!);
    }
    Object.entries(message.fields).forEach(([key, value]) => {
      Object_FieldsEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Object {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseObject();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          message.service = reader.string();
          break;
        case 3:
          message.tags.push(reader.string());
          break;
        case 4:
          const entry4 = Object_FieldsEntry.decode(reader, reader.uint32());
          if (entry4.value !== undefined) {
            message.fields[entry4.key] = entry4.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Object {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      service: isSet(object.service) ? String(object.service) : "",
      tags: Array.isArray(object?.tags)
        ? object.tags.map((e: any) => String(e))
        : [],
      fields: isObject(object.fields)
        ? Object.entries(object.fields).reduce<{ [key: string]: ObjectField }>(
            (acc, [key, value]) => {
              acc[key] = ObjectField.fromJSON(value);
              return acc;
            },
            {}
          )
        : {},
    };
  },

  toJSON(message: Object): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    message.service !== undefined && (obj.service = message.service);
    if (message.tags) {
      obj.tags = message.tags.map((e) => e);
    } else {
      obj.tags = [];
    }
    obj.fields = {};
    if (message.fields) {
      Object.entries(message.fields).forEach(([k, v]) => {
        obj.fields[k] = ObjectField.toJSON(v);
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Object>, I>>(object: I): Object {
    const message = createBaseObject();
    message.did = object.did ?? "";
    message.service = object.service ?? "";
    message.tags = object.tags?.map((e) => e) || [];
    message.fields = Object.entries(object.fields ?? {}).reduce<{
      [key: string]: ObjectField;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = ObjectField.fromPartial(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseObject_FieldsEntry(): Object_FieldsEntry {
  return { key: "", value: undefined };
}

export const Object_FieldsEntry = {
  encode(
    message: Object_FieldsEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      ObjectField.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Object_FieldsEntry {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseObject_FieldsEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = ObjectField.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Object_FieldsEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value)
        ? ObjectField.fromJSON(object.value)
        : undefined,
    };
  },

  toJSON(message: Object_FieldsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value
        ? ObjectField.toJSON(message.value)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Object_FieldsEntry>, I>>(
    object: I
  ): Object_FieldsEntry {
    const message = createBaseObject_FieldsEntry();
    message.key = object.key ?? "";
    message.value =
      object.value !== undefined && object.value !== null
        ? ObjectField.fromPartial(object.value)
        : undefined;
    return message;
  },
};

function createBaseObjectField(): ObjectField {
  return {
    name: "",
    type: 0,
    stringValue: undefined,
    intValue: undefined,
    boolValue: undefined,
    blobValue: undefined,
    referenceValue: undefined,
  };
}

export const ObjectField = {
  encode(message: ObjectField, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.type !== 0) {
      writer.uint32(16).int32(message.type);
    }
    if (message.stringValue !== undefined) {
      writer.uint32(26).string(message.stringValue);
    }
    if (message.intValue !== undefined) {
      writer.uint32(32).int32(message.intValue);
    }
    if (message.boolValue !== undefined) {
      writer.uint32(40).bool(message.boolValue);
    }
    if (message.blobValue !== undefined) {
      writer.uint32(50).bytes(message.blobValue);
    }
    if (message.referenceValue !== undefined) {
      writer.uint32(58).string(message.referenceValue);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ObjectField {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseObjectField();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.type = reader.int32() as any;
          break;
        case 3:
          message.stringValue = reader.string();
          break;
        case 4:
          message.intValue = reader.int32();
          break;
        case 5:
          message.boolValue = reader.bool();
          break;
        case 6:
          message.blobValue = reader.bytes();
          break;
        case 7:
          message.referenceValue = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectField {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      type: isSet(object.type) ? objectFieldTypeFromJSON(object.type) : 0,
      stringValue: isSet(object.stringValue)
        ? String(object.stringValue)
        : undefined,
      intValue: isSet(object.intValue) ? Number(object.intValue) : undefined,
      boolValue: isSet(object.boolValue)
        ? Boolean(object.boolValue)
        : undefined,
      blobValue: isSet(object.blobValue)
        ? bytesFromBase64(object.blobValue)
        : undefined,
      referenceValue: isSet(object.referenceValue)
        ? String(object.referenceValue)
        : undefined,
    };
  },

  toJSON(message: ObjectField): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.type !== undefined &&
      (obj.type = objectFieldTypeToJSON(message.type));
    message.stringValue !== undefined &&
      (obj.stringValue = message.stringValue);
    message.intValue !== undefined &&
      (obj.intValue = Math.round(message.intValue));
    message.boolValue !== undefined && (obj.boolValue = message.boolValue);
    message.blobValue !== undefined &&
      (obj.blobValue =
        message.blobValue !== undefined
          ? base64FromBytes(message.blobValue)
          : undefined);
    message.referenceValue !== undefined &&
      (obj.referenceValue = message.referenceValue);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ObjectField>, I>>(
    object: I
  ): ObjectField {
    const message = createBaseObjectField();
    message.name = object.name ?? "";
    message.type = object.type ?? 0;
    message.stringValue = object.stringValue ?? undefined;
    message.intValue = object.intValue ?? undefined;
    message.boolValue = object.boolValue ?? undefined;
    message.blobValue = object.blobValue ?? undefined;
    message.referenceValue = object.referenceValue ?? undefined;
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

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
