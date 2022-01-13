/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "common";

/** BucketType is the type of a bucket. */
export enum BucketType {
  /** BUCKET_TYPE_UNSPECIFIED - BucketTypeUnspecified is the default value. */
  BUCKET_TYPE_UNSPECIFIED = "BUCKET_TYPE_UNSPECIFIED",
  /** BUCKET_TYPE_APP - BucketTypeApp is an App specific bucket. For Assets regarding the service. */
  BUCKET_TYPE_APP = "BUCKET_TYPE_APP",
  /**
   * BUCKET_TYPE_USER - BucketTypeUser is a User specific bucket. For any remote user data that is required
   * to be stored in the Network.
   */
  BUCKET_TYPE_USER = "BUCKET_TYPE_USER",
  UNRECOGNIZED = "UNRECOGNIZED",
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

export function bucketTypeToNumber(object: BucketType): number {
  switch (object) {
    case BucketType.BUCKET_TYPE_UNSPECIFIED:
      return 0;
    case BucketType.BUCKET_TYPE_APP:
      return 1;
    case BucketType.BUCKET_TYPE_USER:
      return 2;
    default:
      return 0;
  }
}

/** ObjectFieldType is the type of the field */
export enum ObjectFieldType {
  /** OBJECT_FIELD_TYPE_UNSPECIFIED - ObjectFieldTypeUnspecified is the default value */
  OBJECT_FIELD_TYPE_UNSPECIFIED = "OBJECT_FIELD_TYPE_UNSPECIFIED",
  /** OBJECT_FIELD_TYPE_STRING - ObjectFieldTypeString is a string or text field */
  OBJECT_FIELD_TYPE_STRING = "OBJECT_FIELD_TYPE_STRING",
  /** OBJECT_FIELD_TYPE_INT - ObjectFieldTypeInt is an integer */
  OBJECT_FIELD_TYPE_INT = "OBJECT_FIELD_TYPE_INT",
  /** OBJECT_FIELD_TYPE_FLOAT - ObjectFieldTypeFloat is a floating point number */
  OBJECT_FIELD_TYPE_FLOAT = "OBJECT_FIELD_TYPE_FLOAT",
  /** OBJECT_FIELD_TYPE_BOOL - ObjectFieldTypeBool is a boolean */
  OBJECT_FIELD_TYPE_BOOL = "OBJECT_FIELD_TYPE_BOOL",
  /** OBJECT_FIELD_TYPE_DATETIME - ObjectFieldTypeDateTime is a datetime */
  OBJECT_FIELD_TYPE_DATETIME = "OBJECT_FIELD_TYPE_DATETIME",
  /** OBJECT_FIELD_TYPE_BLOB - ObjectFieldTypeBlob is a blob which is a byte array */
  OBJECT_FIELD_TYPE_BLOB = "OBJECT_FIELD_TYPE_BLOB",
  /** OBJECT_FIELD_TYPE_REFERENCE - ObjectFieldTypeReference is a reference to another object */
  OBJECT_FIELD_TYPE_REFERENCE = "OBJECT_FIELD_TYPE_REFERENCE",
  UNRECOGNIZED = "UNRECOGNIZED",
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

export function objectFieldTypeToNumber(object: ObjectFieldType): number {
  switch (object) {
    case ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED:
      return 0;
    case ObjectFieldType.OBJECT_FIELD_TYPE_STRING:
      return 1;
    case ObjectFieldType.OBJECT_FIELD_TYPE_INT:
      return 2;
    case ObjectFieldType.OBJECT_FIELD_TYPE_FLOAT:
      return 3;
    case ObjectFieldType.OBJECT_FIELD_TYPE_BOOL:
      return 4;
    case ObjectFieldType.OBJECT_FIELD_TYPE_DATETIME:
      return 5;
    case ObjectFieldType.OBJECT_FIELD_TYPE_BLOB:
      return 6;
    case ObjectFieldType.OBJECT_FIELD_TYPE_REFERENCE:
      return 7;
    default:
      return 0;
  }
}

/** Bucket is a collection of objects. */
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
  objects: ObjectDoc[];
}

/** ObjectDoc is a document for an Object stored in the graph. */
export interface ObjectDoc {
  /** Did is the identifier of the object. */
  did: string;
  /** Service is the service that created the object. */
  service: string;
  /** Tags are the tags associated with the object. */
  tags: string[];
  /** Fields are the fields associated with the object. */
  fields: { [key: string]: ObjectField };
}

export interface ObjectDoc_FieldsEntry {
  key: string;
  value?: ObjectField;
}

/** ObjectField is a field of an Object. */
export interface ObjectField {
  /** Name is the name of the field */
  name: string;
  /** Type is the type of the field */
  type: ObjectFieldType;
  /** String is the value of the field */
  stringValue: string | undefined;
  /** Int is the value of the field */
  intValue: number | undefined;
  /** Float is the value of the field */
  floatValue: number | undefined;
  /** Bool is the value of the field */
  boolValue: boolean | undefined;
  /** Date is defined by milliseconds since epoch. */
  dateValue: number | undefined;
  /** Blob is the value of the field */
  blobValue: Buffer | undefined;
  /** Reference is the value of the field */
  referenceValue: string | undefined;
  /** Metadata is additional info about the field */
  metadata: { [key: string]: string };
}

export interface ObjectField_MetadataEntry {
  key: string;
  value: string;
}

function createBaseBucket(): Bucket {
  return {
    label: "",
    description: "",
    type: BucketType.BUCKET_TYPE_UNSPECIFIED,
    did: "",
    objects: [],
  };
}

export const Bucket = {
  encode(
    message: Bucket,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.type !== BucketType.BUCKET_TYPE_UNSPECIFIED) {
      writer.uint32(24).int32(bucketTypeToNumber(message.type));
    }
    if (message.did !== "") {
      writer.uint32(34).string(message.did);
    }
    for (const v of message.objects) {
      ObjectDoc.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Bucket {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
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
          message.type = bucketTypeFromJSON(reader.int32());
          break;
        case 4:
          message.did = reader.string();
          break;
        case 5:
          message.objects.push(ObjectDoc.decode(reader, reader.uint32()));
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
      type: isSet(object.type)
        ? bucketTypeFromJSON(object.type)
        : BucketType.BUCKET_TYPE_UNSPECIFIED,
      did: isSet(object.did) ? String(object.did) : "",
      objects: Array.isArray(object?.objects)
        ? object.objects.map((e: any) => ObjectDoc.fromJSON(e))
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
        e ? ObjectDoc.toJSON(e) : undefined
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
    message.type = object.type ?? BucketType.BUCKET_TYPE_UNSPECIFIED;
    message.did = object.did ?? "";
    message.objects =
      object.objects?.map((e) => ObjectDoc.fromPartial(e)) || [];
    return message;
  },
};

function createBaseObjectDoc(): ObjectDoc {
  return { did: "", service: "", tags: [], fields: {} };
}

export const ObjectDoc = {
  encode(
    message: ObjectDoc,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
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
      ObjectDoc_FieldsEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ObjectDoc {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseObjectDoc();
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
          const entry4 = ObjectDoc_FieldsEntry.decode(reader, reader.uint32());
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

  fromJSON(object: any): ObjectDoc {
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

  toJSON(message: ObjectDoc): unknown {
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

  fromPartial<I extends Exact<DeepPartial<ObjectDoc>, I>>(
    object: I
  ): ObjectDoc {
    const message = createBaseObjectDoc();
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

function createBaseObjectDoc_FieldsEntry(): ObjectDoc_FieldsEntry {
  return { key: "", value: undefined };
}

export const ObjectDoc_FieldsEntry = {
  encode(
    message: ObjectDoc_FieldsEntry,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      ObjectField.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ObjectDoc_FieldsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseObjectDoc_FieldsEntry();
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

  fromJSON(object: any): ObjectDoc_FieldsEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value)
        ? ObjectField.fromJSON(object.value)
        : undefined,
    };
  },

  toJSON(message: ObjectDoc_FieldsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value
        ? ObjectField.toJSON(message.value)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ObjectDoc_FieldsEntry>, I>>(
    object: I
  ): ObjectDoc_FieldsEntry {
    const message = createBaseObjectDoc_FieldsEntry();
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
    type: ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED,
    stringValue: undefined,
    intValue: undefined,
    floatValue: undefined,
    boolValue: undefined,
    dateValue: undefined,
    blobValue: undefined,
    referenceValue: undefined,
    metadata: {},
  };
}

export const ObjectField = {
  encode(
    message: ObjectField,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.type !== ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED) {
      writer.uint32(16).int32(objectFieldTypeToNumber(message.type));
    }
    if (message.stringValue !== undefined) {
      writer.uint32(26).string(message.stringValue);
    }
    if (message.intValue !== undefined) {
      writer.uint32(32).int32(message.intValue);
    }
    if (message.floatValue !== undefined) {
      writer.uint32(41).double(message.floatValue);
    }
    if (message.boolValue !== undefined) {
      writer.uint32(48).bool(message.boolValue);
    }
    if (message.dateValue !== undefined) {
      writer.uint32(56).int64(message.dateValue);
    }
    if (message.blobValue !== undefined) {
      writer.uint32(66).bytes(message.blobValue);
    }
    if (message.referenceValue !== undefined) {
      writer.uint32(74).string(message.referenceValue);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ObjectField_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(82).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ObjectField {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseObjectField();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.type = objectFieldTypeFromJSON(reader.int32());
          break;
        case 3:
          message.stringValue = reader.string();
          break;
        case 4:
          message.intValue = reader.int32();
          break;
        case 5:
          message.floatValue = reader.double();
          break;
        case 6:
          message.boolValue = reader.bool();
          break;
        case 7:
          message.dateValue = longToNumber(reader.int64() as Long);
          break;
        case 8:
          message.blobValue = reader.bytes() as Buffer;
          break;
        case 9:
          message.referenceValue = reader.string();
          break;
        case 10:
          const entry10 = ObjectField_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry10.value !== undefined) {
            message.metadata[entry10.key] = entry10.value;
          }
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
      type: isSet(object.type)
        ? objectFieldTypeFromJSON(object.type)
        : ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED,
      stringValue: isSet(object.stringValue)
        ? String(object.stringValue)
        : undefined,
      intValue: isSet(object.intValue) ? Number(object.intValue) : undefined,
      floatValue: isSet(object.floatValue)
        ? Number(object.floatValue)
        : undefined,
      boolValue: isSet(object.boolValue)
        ? Boolean(object.boolValue)
        : undefined,
      dateValue: isSet(object.dateValue) ? Number(object.dateValue) : undefined,
      blobValue: isSet(object.blobValue)
        ? Buffer.from(bytesFromBase64(object.blobValue))
        : undefined,
      referenceValue: isSet(object.referenceValue)
        ? String(object.referenceValue)
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

  toJSON(message: ObjectField): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.type !== undefined &&
      (obj.type = objectFieldTypeToJSON(message.type));
    message.stringValue !== undefined &&
      (obj.stringValue = message.stringValue);
    message.intValue !== undefined &&
      (obj.intValue = Math.round(message.intValue));
    message.floatValue !== undefined && (obj.floatValue = message.floatValue);
    message.boolValue !== undefined && (obj.boolValue = message.boolValue);
    message.dateValue !== undefined &&
      (obj.dateValue = Math.round(message.dateValue));
    message.blobValue !== undefined &&
      (obj.blobValue =
        message.blobValue !== undefined
          ? base64FromBytes(message.blobValue)
          : undefined);
    message.referenceValue !== undefined &&
      (obj.referenceValue = message.referenceValue);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ObjectField>, I>>(
    object: I
  ): ObjectField {
    const message = createBaseObjectField();
    message.name = object.name ?? "";
    message.type = object.type ?? ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED;
    message.stringValue = object.stringValue ?? undefined;
    message.intValue = object.intValue ?? undefined;
    message.floatValue = object.floatValue ?? undefined;
    message.boolValue = object.boolValue ?? undefined;
    message.dateValue = object.dateValue ?? undefined;
    message.blobValue = object.blobValue ?? undefined;
    message.referenceValue = object.referenceValue ?? undefined;
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

function createBaseObjectField_MetadataEntry(): ObjectField_MetadataEntry {
  return { key: "", value: "" };
}

export const ObjectField_MetadataEntry = {
  encode(
    message: ObjectField_MetadataEntry,
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
  ): ObjectField_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseObjectField_MetadataEntry();
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

  fromJSON(object: any): ObjectField_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: ObjectField_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ObjectField_MetadataEntry>, I>>(
    object: I
  ): ObjectField_MetadataEntry {
    const message = createBaseObjectField_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
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
