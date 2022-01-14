/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { ObjectDoc } from "../../object/v1/object";

export const protobufPackage = "protocols.bucket.v1";

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
