/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Session } from "../registry/who_is";
import { WhichIs } from "../bucket/which_is";
import { BucketDoc } from "../bucket/bucket";

export const protobufPackage = "sonrio.sonr.bucket";

export interface MsgCreateBucket {
  creator: string;
  label: string;
  description: string;
  kind: string;
  /** Authenticated user session data */
  session: Session | undefined;
  /** Provided initial objects for the bucket */
  initial_object_dids: string[];
}

export interface MsgCreateBucketResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** Whichis response of the ObjectDoc */
  which_is: WhichIs | undefined;
}

export interface MsgUpdateBucket {
  creator: string;
  /** The Bucket label */
  label: string;
  /** New bucket description */
  description: string;
  /** Session data of authenticated user */
  session: Session | undefined;
  /** Added Objects */
  added_object_dids: string[];
  /** Removed Objects */
  removed_object_dids: string[];
}

export interface MsgUpdateBucketResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** Whichis response of the ObjectDoc */
  which_is: WhichIs | undefined;
}

export interface MsgDeactivateBucket {
  creator: string;
  did: string;
  session: Session | undefined;
}

export interface MsgDeactivateBucketResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
}

export interface MsgCreateWhichIs {
  creator: string;
  did: string;
  bucket: BucketDoc | undefined;
}

export interface MsgCreateWhichIsResponse {}

export interface MsgUpdateWhichIs {
  creator: string;
  did: string;
  bucket: BucketDoc | undefined;
}

export interface MsgUpdateWhichIsResponse {}

export interface MsgDeleteWhichIs {
  creator: string;
  did: string;
}

export interface MsgDeleteWhichIsResponse {}

const baseMsgCreateBucket: object = {
  creator: "",
  label: "",
  description: "",
  kind: "",
  initial_object_dids: "",
};

export const MsgCreateBucket = {
  encode(message: MsgCreateBucket, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.label !== "") {
      writer.uint32(18).string(message.label);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.kind !== "") {
      writer.uint32(34).string(message.kind);
    }
    if (message.session !== undefined) {
      Session.encode(message.session, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.initial_object_dids) {
      writer.uint32(50).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateBucket {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateBucket } as MsgCreateBucket;
    message.initial_object_dids = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.label = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.kind = reader.string();
          break;
        case 5:
          message.session = Session.decode(reader, reader.uint32());
          break;
        case 6:
          message.initial_object_dids.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateBucket {
    const message = { ...baseMsgCreateBucket } as MsgCreateBucket;
    message.initial_object_dids = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
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
    if (object.kind !== undefined && object.kind !== null) {
      message.kind = String(object.kind);
    } else {
      message.kind = "";
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromJSON(object.session);
    } else {
      message.session = undefined;
    }
    if (
      object.initial_object_dids !== undefined &&
      object.initial_object_dids !== null
    ) {
      for (const e of object.initial_object_dids) {
        message.initial_object_dids.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: MsgCreateBucket): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.label !== undefined && (obj.label = message.label);
    message.description !== undefined &&
      (obj.description = message.description);
    message.kind !== undefined && (obj.kind = message.kind);
    message.session !== undefined &&
      (obj.session = message.session
        ? Session.toJSON(message.session)
        : undefined);
    if (message.initial_object_dids) {
      obj.initial_object_dids = message.initial_object_dids.map((e) => e);
    } else {
      obj.initial_object_dids = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateBucket>): MsgCreateBucket {
    const message = { ...baseMsgCreateBucket } as MsgCreateBucket;
    message.initial_object_dids = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
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
    if (object.kind !== undefined && object.kind !== null) {
      message.kind = object.kind;
    } else {
      message.kind = "";
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromPartial(object.session);
    } else {
      message.session = undefined;
    }
    if (
      object.initial_object_dids !== undefined &&
      object.initial_object_dids !== null
    ) {
      for (const e of object.initial_object_dids) {
        message.initial_object_dids.push(e);
      }
    }
    return message;
  },
};

const baseMsgCreateBucketResponse: object = { code: 0, message: "" };

export const MsgCreateBucketResponse = {
  encode(
    message: MsgCreateBucketResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.which_is !== undefined) {
      WhichIs.encode(message.which_is, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateBucketResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateBucketResponse,
    } as MsgCreateBucketResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.which_is = WhichIs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateBucketResponse {
    const message = {
      ...baseMsgCreateBucketResponse,
    } as MsgCreateBucketResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = Number(object.code);
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (object.which_is !== undefined && object.which_is !== null) {
      message.which_is = WhichIs.fromJSON(object.which_is);
    } else {
      message.which_is = undefined;
    }
    return message;
  },

  toJSON(message: MsgCreateBucketResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.message !== undefined && (obj.message = message.message);
    message.which_is !== undefined &&
      (obj.which_is = message.which_is
        ? WhichIs.toJSON(message.which_is)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateBucketResponse>
  ): MsgCreateBucketResponse {
    const message = {
      ...baseMsgCreateBucketResponse,
    } as MsgCreateBucketResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.which_is !== undefined && object.which_is !== null) {
      message.which_is = WhichIs.fromPartial(object.which_is);
    } else {
      message.which_is = undefined;
    }
    return message;
  },
};

const baseMsgUpdateBucket: object = {
  creator: "",
  label: "",
  description: "",
  added_object_dids: "",
  removed_object_dids: "",
};

export const MsgUpdateBucket = {
  encode(message: MsgUpdateBucket, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.label !== "") {
      writer.uint32(18).string(message.label);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.session !== undefined) {
      Session.encode(message.session, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.added_object_dids) {
      writer.uint32(42).string(v!);
    }
    for (const v of message.removed_object_dids) {
      writer.uint32(50).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateBucket {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateBucket } as MsgUpdateBucket;
    message.added_object_dids = [];
    message.removed_object_dids = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.label = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.session = Session.decode(reader, reader.uint32());
          break;
        case 5:
          message.added_object_dids.push(reader.string());
          break;
        case 6:
          message.removed_object_dids.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateBucket {
    const message = { ...baseMsgUpdateBucket } as MsgUpdateBucket;
    message.added_object_dids = [];
    message.removed_object_dids = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
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
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromJSON(object.session);
    } else {
      message.session = undefined;
    }
    if (
      object.added_object_dids !== undefined &&
      object.added_object_dids !== null
    ) {
      for (const e of object.added_object_dids) {
        message.added_object_dids.push(String(e));
      }
    }
    if (
      object.removed_object_dids !== undefined &&
      object.removed_object_dids !== null
    ) {
      for (const e of object.removed_object_dids) {
        message.removed_object_dids.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: MsgUpdateBucket): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.label !== undefined && (obj.label = message.label);
    message.description !== undefined &&
      (obj.description = message.description);
    message.session !== undefined &&
      (obj.session = message.session
        ? Session.toJSON(message.session)
        : undefined);
    if (message.added_object_dids) {
      obj.added_object_dids = message.added_object_dids.map((e) => e);
    } else {
      obj.added_object_dids = [];
    }
    if (message.removed_object_dids) {
      obj.removed_object_dids = message.removed_object_dids.map((e) => e);
    } else {
      obj.removed_object_dids = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateBucket>): MsgUpdateBucket {
    const message = { ...baseMsgUpdateBucket } as MsgUpdateBucket;
    message.added_object_dids = [];
    message.removed_object_dids = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
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
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromPartial(object.session);
    } else {
      message.session = undefined;
    }
    if (
      object.added_object_dids !== undefined &&
      object.added_object_dids !== null
    ) {
      for (const e of object.added_object_dids) {
        message.added_object_dids.push(e);
      }
    }
    if (
      object.removed_object_dids !== undefined &&
      object.removed_object_dids !== null
    ) {
      for (const e of object.removed_object_dids) {
        message.removed_object_dids.push(e);
      }
    }
    return message;
  },
};

const baseMsgUpdateBucketResponse: object = { code: 0, message: "" };

export const MsgUpdateBucketResponse = {
  encode(
    message: MsgUpdateBucketResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.which_is !== undefined) {
      WhichIs.encode(message.which_is, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateBucketResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateBucketResponse,
    } as MsgUpdateBucketResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.which_is = WhichIs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateBucketResponse {
    const message = {
      ...baseMsgUpdateBucketResponse,
    } as MsgUpdateBucketResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = Number(object.code);
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (object.which_is !== undefined && object.which_is !== null) {
      message.which_is = WhichIs.fromJSON(object.which_is);
    } else {
      message.which_is = undefined;
    }
    return message;
  },

  toJSON(message: MsgUpdateBucketResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.message !== undefined && (obj.message = message.message);
    message.which_is !== undefined &&
      (obj.which_is = message.which_is
        ? WhichIs.toJSON(message.which_is)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateBucketResponse>
  ): MsgUpdateBucketResponse {
    const message = {
      ...baseMsgUpdateBucketResponse,
    } as MsgUpdateBucketResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.which_is !== undefined && object.which_is !== null) {
      message.which_is = WhichIs.fromPartial(object.which_is);
    } else {
      message.which_is = undefined;
    }
    return message;
  },
};

const baseMsgDeactivateBucket: object = { creator: "", did: "" };

export const MsgDeactivateBucket = {
  encode(
    message: MsgDeactivateBucket,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.session !== undefined) {
      Session.encode(message.session, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeactivateBucket {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeactivateBucket } as MsgDeactivateBucket;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.session = Session.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeactivateBucket {
    const message = { ...baseMsgDeactivateBucket } as MsgDeactivateBucket;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromJSON(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },

  toJSON(message: MsgDeactivateBucket): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.session !== undefined &&
      (obj.session = message.session
        ? Session.toJSON(message.session)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeactivateBucket>): MsgDeactivateBucket {
    const message = { ...baseMsgDeactivateBucket } as MsgDeactivateBucket;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromPartial(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },
};

const baseMsgDeactivateBucketResponse: object = { code: 0, message: "" };

export const MsgDeactivateBucketResponse = {
  encode(
    message: MsgDeactivateBucketResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeactivateBucketResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeactivateBucketResponse,
    } as MsgDeactivateBucketResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeactivateBucketResponse {
    const message = {
      ...baseMsgDeactivateBucketResponse,
    } as MsgDeactivateBucketResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = Number(object.code);
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    return message;
  },

  toJSON(message: MsgDeactivateBucketResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.message !== undefined && (obj.message = message.message);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeactivateBucketResponse>
  ): MsgDeactivateBucketResponse {
    const message = {
      ...baseMsgDeactivateBucketResponse,
    } as MsgDeactivateBucketResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    return message;
  },
};

const baseMsgCreateWhichIs: object = { creator: "", did: "" };

export const MsgCreateWhichIs = {
  encode(message: MsgCreateWhichIs, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.bucket !== undefined) {
      BucketDoc.encode(message.bucket, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateWhichIs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateWhichIs } as MsgCreateWhichIs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.bucket = BucketDoc.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateWhichIs {
    const message = { ...baseMsgCreateWhichIs } as MsgCreateWhichIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.bucket !== undefined && object.bucket !== null) {
      message.bucket = BucketDoc.fromJSON(object.bucket);
    } else {
      message.bucket = undefined;
    }
    return message;
  },

  toJSON(message: MsgCreateWhichIs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.bucket !== undefined &&
      (obj.bucket = message.bucket
        ? BucketDoc.toJSON(message.bucket)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateWhichIs>): MsgCreateWhichIs {
    const message = { ...baseMsgCreateWhichIs } as MsgCreateWhichIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.bucket !== undefined && object.bucket !== null) {
      message.bucket = BucketDoc.fromPartial(object.bucket);
    } else {
      message.bucket = undefined;
    }
    return message;
  },
};

const baseMsgCreateWhichIsResponse: object = {};

export const MsgCreateWhichIsResponse = {
  encode(
    _: MsgCreateWhichIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateWhichIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateWhichIsResponse,
    } as MsgCreateWhichIsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateWhichIsResponse {
    const message = {
      ...baseMsgCreateWhichIsResponse,
    } as MsgCreateWhichIsResponse;
    return message;
  },

  toJSON(_: MsgCreateWhichIsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateWhichIsResponse>
  ): MsgCreateWhichIsResponse {
    const message = {
      ...baseMsgCreateWhichIsResponse,
    } as MsgCreateWhichIsResponse;
    return message;
  },
};

const baseMsgUpdateWhichIs: object = { creator: "", did: "" };

export const MsgUpdateWhichIs = {
  encode(message: MsgUpdateWhichIs, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.bucket !== undefined) {
      BucketDoc.encode(message.bucket, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateWhichIs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateWhichIs } as MsgUpdateWhichIs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.bucket = BucketDoc.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateWhichIs {
    const message = { ...baseMsgUpdateWhichIs } as MsgUpdateWhichIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.bucket !== undefined && object.bucket !== null) {
      message.bucket = BucketDoc.fromJSON(object.bucket);
    } else {
      message.bucket = undefined;
    }
    return message;
  },

  toJSON(message: MsgUpdateWhichIs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.bucket !== undefined &&
      (obj.bucket = message.bucket
        ? BucketDoc.toJSON(message.bucket)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateWhichIs>): MsgUpdateWhichIs {
    const message = { ...baseMsgUpdateWhichIs } as MsgUpdateWhichIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.bucket !== undefined && object.bucket !== null) {
      message.bucket = BucketDoc.fromPartial(object.bucket);
    } else {
      message.bucket = undefined;
    }
    return message;
  },
};

const baseMsgUpdateWhichIsResponse: object = {};

export const MsgUpdateWhichIsResponse = {
  encode(
    _: MsgUpdateWhichIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateWhichIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateWhichIsResponse,
    } as MsgUpdateWhichIsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateWhichIsResponse {
    const message = {
      ...baseMsgUpdateWhichIsResponse,
    } as MsgUpdateWhichIsResponse;
    return message;
  },

  toJSON(_: MsgUpdateWhichIsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateWhichIsResponse>
  ): MsgUpdateWhichIsResponse {
    const message = {
      ...baseMsgUpdateWhichIsResponse,
    } as MsgUpdateWhichIsResponse;
    return message;
  },
};

const baseMsgDeleteWhichIs: object = { creator: "", did: "" };

export const MsgDeleteWhichIs = {
  encode(message: MsgDeleteWhichIs, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteWhichIs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteWhichIs } as MsgDeleteWhichIs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteWhichIs {
    const message = { ...baseMsgDeleteWhichIs } as MsgDeleteWhichIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteWhichIs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteWhichIs>): MsgDeleteWhichIs {
    const message = { ...baseMsgDeleteWhichIs } as MsgDeleteWhichIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    return message;
  },
};

const baseMsgDeleteWhichIsResponse: object = {};

export const MsgDeleteWhichIsResponse = {
  encode(
    _: MsgDeleteWhichIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteWhichIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteWhichIsResponse,
    } as MsgDeleteWhichIsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteWhichIsResponse {
    const message = {
      ...baseMsgDeleteWhichIsResponse,
    } as MsgDeleteWhichIsResponse;
    return message;
  },

  toJSON(_: MsgDeleteWhichIsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteWhichIsResponse>
  ): MsgDeleteWhichIsResponse {
    const message = {
      ...baseMsgDeleteWhichIsResponse,
    } as MsgDeleteWhichIsResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /**
   * CreateBucket
   *
   * CreateBucket defines a new collection on the bucket module of the blockchain.
   */
  CreateBucket(request: MsgCreateBucket): Promise<MsgCreateBucketResponse>;
  /**
   * UpdateBucket
   *
   * UpdateBucket updates existing collection on the bucket module of the blockchain.
   */
  UpdateBucket(request: MsgUpdateBucket): Promise<MsgUpdateBucketResponse>;
  /**
   * DeactivateBucket
   *
   * DeactivateBucket deactivates existing collection on the bucket module of the blockchain.
   */
  DeactivateBucket(
    request: MsgDeactivateBucket
  ): Promise<MsgDeactivateBucketResponse>;
  /**
   * CreateWhichIs
   *
   * CreateWhichIs method creates a new BucketDoc record for the bucket module.
   */
  CreateWhichIs(request: MsgCreateWhichIs): Promise<MsgCreateWhichIsResponse>;
  /**
   * UpdateWhichIs
   *
   * UpdateWhichIs method updates an existing BucketDoc from the bucket store.
   */
  UpdateWhichIs(request: MsgUpdateWhichIs): Promise<MsgUpdateWhichIsResponse>;
  /**
   * DeleteWhichIs
   *
   * DeleteWhichIs method deletes an existing BucketDoc from the bucket store.
   */
  DeleteWhichIs(request: MsgDeleteWhichIs): Promise<MsgDeleteWhichIsResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateBucket(request: MsgCreateBucket): Promise<MsgCreateBucketResponse> {
    const data = MsgCreateBucket.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.bucket.Msg",
      "CreateBucket",
      data
    );
    return promise.then((data) =>
      MsgCreateBucketResponse.decode(new Reader(data))
    );
  }

  UpdateBucket(request: MsgUpdateBucket): Promise<MsgUpdateBucketResponse> {
    const data = MsgUpdateBucket.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.bucket.Msg",
      "UpdateBucket",
      data
    );
    return promise.then((data) =>
      MsgUpdateBucketResponse.decode(new Reader(data))
    );
  }

  DeactivateBucket(
    request: MsgDeactivateBucket
  ): Promise<MsgDeactivateBucketResponse> {
    const data = MsgDeactivateBucket.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.bucket.Msg",
      "DeactivateBucket",
      data
    );
    return promise.then((data) =>
      MsgDeactivateBucketResponse.decode(new Reader(data))
    );
  }

  CreateWhichIs(request: MsgCreateWhichIs): Promise<MsgCreateWhichIsResponse> {
    const data = MsgCreateWhichIs.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.bucket.Msg",
      "CreateWhichIs",
      data
    );
    return promise.then((data) =>
      MsgCreateWhichIsResponse.decode(new Reader(data))
    );
  }

  UpdateWhichIs(request: MsgUpdateWhichIs): Promise<MsgUpdateWhichIsResponse> {
    const data = MsgUpdateWhichIs.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.bucket.Msg",
      "UpdateWhichIs",
      data
    );
    return promise.then((data) =>
      MsgUpdateWhichIsResponse.decode(new Reader(data))
    );
  }

  DeleteWhichIs(request: MsgDeleteWhichIs): Promise<MsgDeleteWhichIsResponse> {
    const data = MsgDeleteWhichIs.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.bucket.Msg",
      "DeleteWhichIs",
      data
    );
    return promise.then((data) =>
      MsgDeleteWhichIsResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
