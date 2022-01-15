/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Did } from "../../common/did";
import { ObjectDoc } from "../../common/object";

export const protobufPackage = "service.v1";

/** ServiceConfig is the configuration for a service. */
export interface ServiceConfig {
  /** Name is the name of the service. */
  name: string;
  /** Description is a human readable description of the service. */
  description: string;
  /** Owner is the DID of the service owner. */
  owner?: Did;
  /** Tags is a list of tags the service is registered with. */
  tags: string[];
  /** Channels is a list of channels the service is registered on. */
  channels: Did[];
  /** Buckets is a list of buckets the service is registered on. */
  buckets: Did[];
  /** Objects is a map of objects the service is registered on. */
  objects: { [key: string]: ObjectDoc };
  /** Endpoints is a list of endpoints the service is registered on. */
  endpoints: string[];
  /** Metadata is the metadata associated with the event. */
  metadata: { [key: string]: string };
  /** Version is the version of the service. Version must be a semantic version. */
  version: string;
}

export interface ServiceConfig_ObjectsEntry {
  key: string;
  value?: ObjectDoc;
}

export interface ServiceConfig_MetadataEntry {
  key: string;
  value: string;
}

function createBaseServiceConfig(): ServiceConfig {
  return {
    name: "",
    description: "",
    owner: undefined,
    tags: [],
    channels: [],
    buckets: [],
    objects: {},
    endpoints: [],
    metadata: {},
    version: "",
  };
}

export const ServiceConfig = {
  encode(
    message: ServiceConfig,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.owner !== undefined) {
      Did.encode(message.owner, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.tags) {
      writer.uint32(26).string(v!);
    }
    for (const v of message.channels) {
      Did.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.buckets) {
      Did.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    Object.entries(message.objects).forEach(([key, value]) => {
      ServiceConfig_ObjectsEntry.encode(
        { key: key as any, value },
        writer.uint32(58).fork()
      ).ldelim();
    });
    for (const v of message.endpoints) {
      writer.uint32(66).string(v!);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ServiceConfig_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(74).fork()
      ).ldelim();
    });
    if (message.version !== "") {
      writer.uint32(82).string(message.version);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ServiceConfig {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseServiceConfig();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 4:
          message.owner = Did.decode(reader, reader.uint32());
          break;
        case 3:
          message.tags.push(reader.string());
          break;
        case 5:
          message.channels.push(Did.decode(reader, reader.uint32()));
          break;
        case 6:
          message.buckets.push(Did.decode(reader, reader.uint32()));
          break;
        case 7:
          const entry7 = ServiceConfig_ObjectsEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry7.value !== undefined) {
            message.objects[entry7.key] = entry7.value;
          }
          break;
        case 8:
          message.endpoints.push(reader.string());
          break;
        case 9:
          const entry9 = ServiceConfig_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry9.value !== undefined) {
            message.metadata[entry9.key] = entry9.value;
          }
          break;
        case 10:
          message.version = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ServiceConfig {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      owner: isSet(object.owner) ? Did.fromJSON(object.owner) : undefined,
      tags: Array.isArray(object?.tags)
        ? object.tags.map((e: any) => String(e))
        : [],
      channels: Array.isArray(object?.channels)
        ? object.channels.map((e: any) => Did.fromJSON(e))
        : [],
      buckets: Array.isArray(object?.buckets)
        ? object.buckets.map((e: any) => Did.fromJSON(e))
        : [],
      objects: isObject(object.objects)
        ? Object.entries(object.objects).reduce<{ [key: string]: ObjectDoc }>(
            (acc, [key, value]) => {
              acc[key] = ObjectDoc.fromJSON(value);
              return acc;
            },
            {}
          )
        : {},
      endpoints: Array.isArray(object?.endpoints)
        ? object.endpoints.map((e: any) => String(e))
        : [],
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
      version: isSet(object.version) ? String(object.version) : "",
    };
  },

  toJSON(message: ServiceConfig): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined &&
      (obj.description = message.description);
    message.owner !== undefined &&
      (obj.owner = message.owner ? Did.toJSON(message.owner) : undefined);
    if (message.tags) {
      obj.tags = message.tags.map((e) => e);
    } else {
      obj.tags = [];
    }
    if (message.channels) {
      obj.channels = message.channels.map((e) =>
        e ? Did.toJSON(e) : undefined
      );
    } else {
      obj.channels = [];
    }
    if (message.buckets) {
      obj.buckets = message.buckets.map((e) => (e ? Did.toJSON(e) : undefined));
    } else {
      obj.buckets = [];
    }
    obj.objects = {};
    if (message.objects) {
      Object.entries(message.objects).forEach(([k, v]) => {
        obj.objects[k] = ObjectDoc.toJSON(v);
      });
    }
    if (message.endpoints) {
      obj.endpoints = message.endpoints.map((e) => e);
    } else {
      obj.endpoints = [];
    }
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    message.version !== undefined && (obj.version = message.version);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ServiceConfig>, I>>(
    object: I
  ): ServiceConfig {
    const message = createBaseServiceConfig();
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.owner =
      object.owner !== undefined && object.owner !== null
        ? Did.fromPartial(object.owner)
        : undefined;
    message.tags = object.tags?.map((e) => e) || [];
    message.channels = object.channels?.map((e) => Did.fromPartial(e)) || [];
    message.buckets = object.buckets?.map((e) => Did.fromPartial(e)) || [];
    message.objects = Object.entries(object.objects ?? {}).reduce<{
      [key: string]: ObjectDoc;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = ObjectDoc.fromPartial(value);
      }
      return acc;
    }, {});
    message.endpoints = object.endpoints?.map((e) => e) || [];
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    message.version = object.version ?? "";
    return message;
  },
};

function createBaseServiceConfig_ObjectsEntry(): ServiceConfig_ObjectsEntry {
  return { key: "", value: undefined };
}

export const ServiceConfig_ObjectsEntry = {
  encode(
    message: ServiceConfig_ObjectsEntry,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      ObjectDoc.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ServiceConfig_ObjectsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseServiceConfig_ObjectsEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = ObjectDoc.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ServiceConfig_ObjectsEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? ObjectDoc.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: ServiceConfig_ObjectsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value ? ObjectDoc.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ServiceConfig_ObjectsEntry>, I>>(
    object: I
  ): ServiceConfig_ObjectsEntry {
    const message = createBaseServiceConfig_ObjectsEntry();
    message.key = object.key ?? "";
    message.value =
      object.value !== undefined && object.value !== null
        ? ObjectDoc.fromPartial(object.value)
        : undefined;
    return message;
  },
};

function createBaseServiceConfig_MetadataEntry(): ServiceConfig_MetadataEntry {
  return { key: "", value: "" };
}

export const ServiceConfig_MetadataEntry = {
  encode(
    message: ServiceConfig_MetadataEntry,
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
  ): ServiceConfig_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseServiceConfig_MetadataEntry();
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

  fromJSON(object: any): ServiceConfig_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: ServiceConfig_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ServiceConfig_MetadataEntry>, I>>(
    object: I
  ): ServiceConfig_MetadataEntry {
    const message = createBaseServiceConfig_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
