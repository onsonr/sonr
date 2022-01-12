/* eslint-disable */
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import * as Long from "long";
import { Payload } from "../../../common/v1/data";
import { Profile, Metadata, Peer } from "../../../common/v1/core";

export const protobufPackage = "protocols.exchange.v1";

export interface MailboxMessage {
  /** ID is the Message ID */
  id: string;
  /** Payload is the message data */
  payload: Payload | undefined;
  /** Users Peer Data */
  from: Profile | undefined;
  /** Receivers Peer Data */
  to: Profile | undefined;
  /** Metadata */
  metadata: Metadata | undefined;
  /** Timestamp */
  createdAt: number;
}

/** Invitation Message sent on RPC */
export interface InviteRequest {
  /** Attached Data */
  payload: Payload | undefined;
  /** Users Peer Data */
  from: Peer | undefined;
  /** Receivers Peer Data */
  to: Peer | undefined;
  /** Metadata */
  metadata: Metadata | undefined;
}

/** Reply Message sent on RPC */
export interface InviteResponse {
  /** Success */
  decision: boolean;
  /** Users Peer Data */
  from: Peer | undefined;
  /** Receivers Peer Data */
  to: Peer | undefined;
  /** Metadata */
  metadata: Metadata | undefined;
}

function createBaseMailboxMessage(): MailboxMessage {
  return {
    id: "",
    payload: undefined,
    from: undefined,
    to: undefined,
    metadata: undefined,
    createdAt: 0,
  };
}

export const MailboxMessage = {
  encode(message: MailboxMessage, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.payload !== undefined) {
      Payload.encode(message.payload, writer.uint32(18).fork()).ldelim();
    }
    if (message.from !== undefined) {
      Profile.encode(message.from, writer.uint32(26).fork()).ldelim();
    }
    if (message.to !== undefined) {
      Profile.encode(message.to, writer.uint32(34).fork()).ldelim();
    }
    if (message.metadata !== undefined) {
      Metadata.encode(message.metadata, writer.uint32(42).fork()).ldelim();
    }
    if (message.createdAt !== 0) {
      writer.uint32(48).int64(message.createdAt);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MailboxMessage {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMailboxMessage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.payload = Payload.decode(reader, reader.uint32());
          break;
        case 3:
          message.from = Profile.decode(reader, reader.uint32());
          break;
        case 4:
          message.to = Profile.decode(reader, reader.uint32());
          break;
        case 5:
          message.metadata = Metadata.decode(reader, reader.uint32());
          break;
        case 6:
          message.createdAt = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MailboxMessage {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      payload: isSet(object.payload)
        ? Payload.fromJSON(object.payload)
        : undefined,
      from: isSet(object.from) ? Profile.fromJSON(object.from) : undefined,
      to: isSet(object.to) ? Profile.fromJSON(object.to) : undefined,
      metadata: isSet(object.metadata)
        ? Metadata.fromJSON(object.metadata)
        : undefined,
      createdAt: isSet(object.createdAt) ? Number(object.createdAt) : 0,
    };
  },

  toJSON(message: MailboxMessage): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.payload !== undefined &&
      (obj.payload = message.payload
        ? Payload.toJSON(message.payload)
        : undefined);
    message.from !== undefined &&
      (obj.from = message.from ? Profile.toJSON(message.from) : undefined);
    message.to !== undefined &&
      (obj.to = message.to ? Profile.toJSON(message.to) : undefined);
    message.metadata !== undefined &&
      (obj.metadata = message.metadata
        ? Metadata.toJSON(message.metadata)
        : undefined);
    message.createdAt !== undefined &&
      (obj.createdAt = Math.round(message.createdAt));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MailboxMessage>, I>>(
    object: I
  ): MailboxMessage {
    const message = createBaseMailboxMessage();
    message.id = object.id ?? "";
    message.payload =
      object.payload !== undefined && object.payload !== null
        ? Payload.fromPartial(object.payload)
        : undefined;
    message.from =
      object.from !== undefined && object.from !== null
        ? Profile.fromPartial(object.from)
        : undefined;
    message.to =
      object.to !== undefined && object.to !== null
        ? Profile.fromPartial(object.to)
        : undefined;
    message.metadata =
      object.metadata !== undefined && object.metadata !== null
        ? Metadata.fromPartial(object.metadata)
        : undefined;
    message.createdAt = object.createdAt ?? 0;
    return message;
  },
};

function createBaseInviteRequest(): InviteRequest {
  return {
    payload: undefined,
    from: undefined,
    to: undefined,
    metadata: undefined,
  };
}

export const InviteRequest = {
  encode(message: InviteRequest, writer: Writer = Writer.create()): Writer {
    if (message.payload !== undefined) {
      Payload.encode(message.payload, writer.uint32(10).fork()).ldelim();
    }
    if (message.from !== undefined) {
      Peer.encode(message.from, writer.uint32(26).fork()).ldelim();
    }
    if (message.to !== undefined) {
      Peer.encode(message.to, writer.uint32(34).fork()).ldelim();
    }
    if (message.metadata !== undefined) {
      Metadata.encode(message.metadata, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): InviteRequest {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInviteRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.payload = Payload.decode(reader, reader.uint32());
          break;
        case 3:
          message.from = Peer.decode(reader, reader.uint32());
          break;
        case 4:
          message.to = Peer.decode(reader, reader.uint32());
          break;
        case 5:
          message.metadata = Metadata.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InviteRequest {
    return {
      payload: isSet(object.payload)
        ? Payload.fromJSON(object.payload)
        : undefined,
      from: isSet(object.from) ? Peer.fromJSON(object.from) : undefined,
      to: isSet(object.to) ? Peer.fromJSON(object.to) : undefined,
      metadata: isSet(object.metadata)
        ? Metadata.fromJSON(object.metadata)
        : undefined,
    };
  },

  toJSON(message: InviteRequest): unknown {
    const obj: any = {};
    message.payload !== undefined &&
      (obj.payload = message.payload
        ? Payload.toJSON(message.payload)
        : undefined);
    message.from !== undefined &&
      (obj.from = message.from ? Peer.toJSON(message.from) : undefined);
    message.to !== undefined &&
      (obj.to = message.to ? Peer.toJSON(message.to) : undefined);
    message.metadata !== undefined &&
      (obj.metadata = message.metadata
        ? Metadata.toJSON(message.metadata)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<InviteRequest>, I>>(
    object: I
  ): InviteRequest {
    const message = createBaseInviteRequest();
    message.payload =
      object.payload !== undefined && object.payload !== null
        ? Payload.fromPartial(object.payload)
        : undefined;
    message.from =
      object.from !== undefined && object.from !== null
        ? Peer.fromPartial(object.from)
        : undefined;
    message.to =
      object.to !== undefined && object.to !== null
        ? Peer.fromPartial(object.to)
        : undefined;
    message.metadata =
      object.metadata !== undefined && object.metadata !== null
        ? Metadata.fromPartial(object.metadata)
        : undefined;
    return message;
  },
};

function createBaseInviteResponse(): InviteResponse {
  return {
    decision: false,
    from: undefined,
    to: undefined,
    metadata: undefined,
  };
}

export const InviteResponse = {
  encode(message: InviteResponse, writer: Writer = Writer.create()): Writer {
    if (message.decision === true) {
      writer.uint32(8).bool(message.decision);
    }
    if (message.from !== undefined) {
      Peer.encode(message.from, writer.uint32(26).fork()).ldelim();
    }
    if (message.to !== undefined) {
      Peer.encode(message.to, writer.uint32(34).fork()).ldelim();
    }
    if (message.metadata !== undefined) {
      Metadata.encode(message.metadata, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): InviteResponse {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInviteResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.decision = reader.bool();
          break;
        case 3:
          message.from = Peer.decode(reader, reader.uint32());
          break;
        case 4:
          message.to = Peer.decode(reader, reader.uint32());
          break;
        case 5:
          message.metadata = Metadata.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InviteResponse {
    return {
      decision: isSet(object.decision) ? Boolean(object.decision) : false,
      from: isSet(object.from) ? Peer.fromJSON(object.from) : undefined,
      to: isSet(object.to) ? Peer.fromJSON(object.to) : undefined,
      metadata: isSet(object.metadata)
        ? Metadata.fromJSON(object.metadata)
        : undefined,
    };
  },

  toJSON(message: InviteResponse): unknown {
    const obj: any = {};
    message.decision !== undefined && (obj.decision = message.decision);
    message.from !== undefined &&
      (obj.from = message.from ? Peer.toJSON(message.from) : undefined);
    message.to !== undefined &&
      (obj.to = message.to ? Peer.toJSON(message.to) : undefined);
    message.metadata !== undefined &&
      (obj.metadata = message.metadata
        ? Metadata.toJSON(message.metadata)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<InviteResponse>, I>>(
    object: I
  ): InviteResponse {
    const message = createBaseInviteResponse();
    message.decision = object.decision ?? false;
    message.from =
      object.from !== undefined && object.from !== null
        ? Peer.fromPartial(object.from)
        : undefined;
    message.to =
      object.to !== undefined && object.to !== null
        ? Peer.fromPartial(object.to)
        : undefined;
    message.metadata =
      object.metadata !== undefined && object.metadata !== null
        ? Metadata.fromPartial(object.metadata)
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

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
