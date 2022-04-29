/* eslint-disable */
import { Did } from "../registry/did";
import { ObjectDoc } from "../object/object";
import { Peer } from "../registry/peer";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sonrio.sonr.bucket";

/** BucketType is the type of a bucket. */
export enum BucketType {
  /** BUCKET_TYPE_UNSPECIFIED - BucketTypeUnspecified is the default value. */
  BUCKET_TYPE_UNSPECIFIED = 0,
  /** BUCKET_TYPE_APP - BucketTypeApp is an App specific bucket. For Assets regarding the service. */
  BUCKET_TYPE_APP = 1,
  /**
   * BUCKET_TYPE_USER - BucketTypeUser is a User specific bucket. For any remote user data that is required
   * to be stored in the Network.
   */
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

/** EventType is the type of event being performed on a Bucket. */
export enum BucketEventType {
  /** Bucket_EVENT_TYPE_UNSPECIFIED - EventTypeUnspecified is the default value. */
  Bucket_EVENT_TYPE_UNSPECIFIED = 0,
  /** Bucket_EVENT_TYPE_GET - EventTypeGet is a get event being performed on a Bucket record. */
  Bucket_EVENT_TYPE_GET = 1,
  /** Bucket_EVENT_TYPE_SET - EventTypeSet is a set event on the Bucket store. */
  Bucket_EVENT_TYPE_SET = 2,
  /** Bucket_EVENT_TYPE_DELETE - EventTypeDelete is a delete event on the Bucket store. */
  Bucket_EVENT_TYPE_DELETE = 3,
  UNRECOGNIZED = -1,
}

export function bucketEventTypeFromJSON(object: any): BucketEventType {
  switch (object) {
    case 0:
    case "Bucket_EVENT_TYPE_UNSPECIFIED":
      return BucketEventType.Bucket_EVENT_TYPE_UNSPECIFIED;
    case 1:
    case "Bucket_EVENT_TYPE_GET":
      return BucketEventType.Bucket_EVENT_TYPE_GET;
    case 2:
    case "Bucket_EVENT_TYPE_SET":
      return BucketEventType.Bucket_EVENT_TYPE_SET;
    case 3:
    case "Bucket_EVENT_TYPE_DELETE":
      return BucketEventType.Bucket_EVENT_TYPE_DELETE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return BucketEventType.UNRECOGNIZED;
  }
}

export function bucketEventTypeToJSON(object: BucketEventType): string {
  switch (object) {
    case BucketEventType.Bucket_EVENT_TYPE_UNSPECIFIED:
      return "Bucket_EVENT_TYPE_UNSPECIFIED";
    case BucketEventType.Bucket_EVENT_TYPE_GET:
      return "Bucket_EVENT_TYPE_GET";
    case BucketEventType.Bucket_EVENT_TYPE_SET:
      return "Bucket_EVENT_TYPE_SET";
    case BucketEventType.Bucket_EVENT_TYPE_DELETE:
      return "Bucket_EVENT_TYPE_DELETE";
    default:
      return "UNKNOWN";
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
  did: Did | undefined;
  /** Objects are stored in a tree structure. */
  objects: ObjectDoc[];
}

/** BucketEvent is the base event type for all Bucket events. */
export interface BucketEvent {
  /** Owner is the peer that originated the event. */
  peer: Peer | undefined;
  /** Type is the type of event being performed on a Bucket. */
  type: BucketEventType;
  /** Object is the entry being modified in the Bucket. */
  object: ObjectDoc | undefined;
  /** Metadata is the metadata associated with the event. */
  metadata: { [key: string]: string };
}

export interface BucketEvent_MetadataEntry {
  key: string;
  value: string;
}

const baseBucket: object = { label: "", description: "", type: 0 };

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
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.objects) {
      ObjectDoc.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Bucket {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBucket } as Bucket;
    message.objects = [];
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
          message.did = Did.decode(reader, reader.uint32());
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
    const message = { ...baseBucket } as Bucket;
    message.objects = [];
    if (object.label !== undefined && object.label !== null) {
      message.label = String(object.label);
    } else {
      message.label = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = bucketTypeFromJSON(object.type);
    } else {
      message.type = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = Did.fromJSON(object.did);
    } else {
      message.did = undefined;
    }
    if (object.objects !== undefined && object.objects !== null) {
      for (const e of object.objects) {
        message.objects.push(ObjectDoc.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Bucket): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.description !== undefined &&
      (obj.description = message.description);
    message.type !== undefined && (obj.type = bucketTypeToJSON(message.type));
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    if (message.objects) {
      obj.objects = message.objects.map((e) =>
        e ? ObjectDoc.toJSON(e) : undefined
      );
    } else {
      obj.objects = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Bucket>): Bucket {
    const message = { ...baseBucket } as Bucket;
    message.objects = [];
    if (object.label !== undefined && object.label !== null) {
      message.label = object.label;
    } else {
      message.label = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = Did.fromPartial(object.did);
    } else {
      message.did = undefined;
    }
    if (object.objects !== undefined && object.objects !== null) {
      for (const e of object.objects) {
        message.objects.push(ObjectDoc.fromPartial(e));
      }
    }
    return message;
  },
};

const baseBucketEvent: object = { type: 0 };

export const BucketEvent = {
  encode(message: BucketEvent, writer: Writer = Writer.create()): Writer {
    if (message.peer !== undefined) {
      Peer.encode(message.peer, writer.uint32(10).fork()).ldelim();
    }
    if (message.type !== 0) {
      writer.uint32(16).int32(message.type);
    }
    if (message.object !== undefined) {
      ObjectDoc.encode(message.object, writer.uint32(26).fork()).ldelim();
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      BucketEvent_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): BucketEvent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBucketEvent } as BucketEvent;
    message.metadata = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.peer = Peer.decode(reader, reader.uint32());
          break;
        case 2:
          message.type = reader.int32() as any;
          break;
        case 3:
          message.object = ObjectDoc.decode(reader, reader.uint32());
          break;
        case 4:
          const entry4 = BucketEvent_MetadataEntry.decode(
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

  fromJSON(object: any): BucketEvent {
    const message = { ...baseBucketEvent } as BucketEvent;
    message.metadata = {};
    if (object.peer !== undefined && object.peer !== null) {
      message.peer = Peer.fromJSON(object.peer);
    } else {
      message.peer = undefined;
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = bucketEventTypeFromJSON(object.type);
    } else {
      message.type = 0;
    }
    if (object.object !== undefined && object.object !== null) {
      message.object = ObjectDoc.fromJSON(object.object);
    } else {
      message.object = undefined;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    return message;
  },

  toJSON(message: BucketEvent): unknown {
    const obj: any = {};
    message.peer !== undefined &&
      (obj.peer = message.peer ? Peer.toJSON(message.peer) : undefined);
    message.type !== undefined &&
      (obj.type = bucketEventTypeToJSON(message.type));
    message.object !== undefined &&
      (obj.object = message.object
        ? ObjectDoc.toJSON(message.object)
        : undefined);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<BucketEvent>): BucketEvent {
    const message = { ...baseBucketEvent } as BucketEvent;
    message.metadata = {};
    if (object.peer !== undefined && object.peer !== null) {
      message.peer = Peer.fromPartial(object.peer);
    } else {
      message.peer = undefined;
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = 0;
    }
    if (object.object !== undefined && object.object !== null) {
      message.object = ObjectDoc.fromPartial(object.object);
    } else {
      message.object = undefined;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
    }
    return message;
  },
};

const baseBucketEvent_MetadataEntry: object = { key: "", value: "" };

export const BucketEvent_MetadataEntry = {
  encode(
    message: BucketEvent_MetadataEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): BucketEvent_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseBucketEvent_MetadataEntry,
    } as BucketEvent_MetadataEntry;
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

  fromJSON(object: any): BucketEvent_MetadataEntry {
    const message = {
      ...baseBucketEvent_MetadataEntry,
    } as BucketEvent_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: BucketEvent_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<BucketEvent_MetadataEntry>
  ): BucketEvent_MetadataEntry {
    const message = {
      ...baseBucketEvent_MetadataEntry,
    } as BucketEvent_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
