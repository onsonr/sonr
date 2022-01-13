/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Peer, Profile, Metadata } from "../../../common/core";
import {
  Payload,
  Direction,
  directionToNumber,
  directionFromJSON,
  directionToJSON,
} from "../../../common/data";

export const protobufPackage = "node.motor.v1";

/** (Client) ShareResponse is response to ShareRequest */
export interface ShareResponse {
  /** True if Supply is Active */
  success: boolean;
  /** Error Message if Supply is not Active */
  error: string;
}

/** (Client) RespondResponse is response to RespondRequest */
export interface DecideResponse {
  /** True if Supply is Active */
  success: boolean;
  /** Error Message if Supply is not Active */
  error: string;
}

/** (Client) SearchResponse is Message for Searching for Peer */
export interface SearchResponse {
  /** Success */
  success: boolean;
  /** Error Message */
  error: string;
  /** Peer Data */
  peer?: Peer;
  /** Peer ID */
  peerId: string;
  /** SName */
  sName: string;
}

/** DecisionEvent is emitted when a decision is made by Peer. */
export interface OnTransmitDecisionResponse {
  /** true = accept, false = reject */
  decision: boolean;
  /** Peer that made decision */
  from?: Peer;
  /** Timestamp */
  received: number;
}

/** Message Sent when peer messages Lobby */
export interface OnLobbyRefreshResponse {
  /** OLC Code of Topic */
  olc: string;
  /** User Information */
  peers: Peer[];
  /** Invite received Timestamp */
  received: number;
}

/** InviteEvent notifies Peer that an Invite has been received */
export interface OnTransmitInviteResponse {
  /** Invite received Timestamp */
  received: number;
  /** Peer that sent the Invite */
  from?: Peer;
  /** Attached Data */
  payload?: Payload;
}

/** Received Mail Event */
export interface OnMailboxMessageResponse {
  /** ID is the Message ID */
  id: string;
  /** Buffer is the message data */
  buffer: Buffer;
  /** Users Peer Data */
  from?: Profile;
  /** Receivers Peer Data */
  to?: Profile;
  /** Metadata */
  metadata?: Metadata;
}

/** Transfer Progress Event */
export interface OnTransmitProgressResponse {
  /** Current Transfer Progress */
  progress: number;
  /** Timestamp */
  received: number;
  /** Current position of item in list */
  current: number;
  /** Total number of items in list */
  total: number;
  /** Direction of Transfer */
  direction: Direction;
}

/** Message Sent after Completed Transfer */
export interface OnTransmitCompleteResponse {
  /** Direction of Transfer */
  direction: Direction;
  /** Transfer Data */
  payload?: Payload;
  /** Peer that sent the Complete Event */
  from?: Peer;
  /** Peer that received the Complete Event */
  to?: Peer;
  /** Transfer Created Timestamp */
  createdAt: number;
  /** Transfer Received Timestamp */
  receivedAt: number;
  /** Transfer Success */
  results: { [key: number]: boolean };
}

export interface OnTransmitCompleteResponse_ResultsEntry {
  key: number;
  value: boolean;
}

function createBaseShareResponse(): ShareResponse {
  return { success: false, error: "" };
}

export const ShareResponse = {
  encode(
    message: ShareResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    if (message.error !== "") {
      writer.uint32(18).string(message.error);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ShareResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseShareResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        case 2:
          message.error = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ShareResponse {
    return {
      success: isSet(object.success) ? Boolean(object.success) : false,
      error: isSet(object.error) ? String(object.error) : "",
    };
  },

  toJSON(message: ShareResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.error !== undefined && (obj.error = message.error);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ShareResponse>, I>>(
    object: I
  ): ShareResponse {
    const message = createBaseShareResponse();
    message.success = object.success ?? false;
    message.error = object.error ?? "";
    return message;
  },
};

function createBaseDecideResponse(): DecideResponse {
  return { success: false, error: "" };
}

export const DecideResponse = {
  encode(
    message: DecideResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    if (message.error !== "") {
      writer.uint32(18).string(message.error);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DecideResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDecideResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        case 2:
          message.error = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DecideResponse {
    return {
      success: isSet(object.success) ? Boolean(object.success) : false,
      error: isSet(object.error) ? String(object.error) : "",
    };
  },

  toJSON(message: DecideResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.error !== undefined && (obj.error = message.error);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DecideResponse>, I>>(
    object: I
  ): DecideResponse {
    const message = createBaseDecideResponse();
    message.success = object.success ?? false;
    message.error = object.error ?? "";
    return message;
  },
};

function createBaseSearchResponse(): SearchResponse {
  return { success: false, error: "", peer: undefined, peerId: "", sName: "" };
}

export const SearchResponse = {
  encode(
    message: SearchResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    if (message.error !== "") {
      writer.uint32(18).string(message.error);
    }
    if (message.peer !== undefined) {
      Peer.encode(message.peer, writer.uint32(26).fork()).ldelim();
    }
    if (message.peerId !== "") {
      writer.uint32(34).string(message.peerId);
    }
    if (message.sName !== "") {
      writer.uint32(42).string(message.sName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SearchResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSearchResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        case 2:
          message.error = reader.string();
          break;
        case 3:
          message.peer = Peer.decode(reader, reader.uint32());
          break;
        case 4:
          message.peerId = reader.string();
          break;
        case 5:
          message.sName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SearchResponse {
    return {
      success: isSet(object.success) ? Boolean(object.success) : false,
      error: isSet(object.error) ? String(object.error) : "",
      peer: isSet(object.peer) ? Peer.fromJSON(object.peer) : undefined,
      peerId: isSet(object.peerId) ? String(object.peerId) : "",
      sName: isSet(object.sName) ? String(object.sName) : "",
    };
  },

  toJSON(message: SearchResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.error !== undefined && (obj.error = message.error);
    message.peer !== undefined &&
      (obj.peer = message.peer ? Peer.toJSON(message.peer) : undefined);
    message.peerId !== undefined && (obj.peerId = message.peerId);
    message.sName !== undefined && (obj.sName = message.sName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SearchResponse>, I>>(
    object: I
  ): SearchResponse {
    const message = createBaseSearchResponse();
    message.success = object.success ?? false;
    message.error = object.error ?? "";
    message.peer =
      object.peer !== undefined && object.peer !== null
        ? Peer.fromPartial(object.peer)
        : undefined;
    message.peerId = object.peerId ?? "";
    message.sName = object.sName ?? "";
    return message;
  },
};

function createBaseOnTransmitDecisionResponse(): OnTransmitDecisionResponse {
  return { decision: false, from: undefined, received: 0 };
}

export const OnTransmitDecisionResponse = {
  encode(
    message: OnTransmitDecisionResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.decision === true) {
      writer.uint32(8).bool(message.decision);
    }
    if (message.from !== undefined) {
      Peer.encode(message.from, writer.uint32(18).fork()).ldelim();
    }
    if (message.received !== 0) {
      writer.uint32(24).int64(message.received);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): OnTransmitDecisionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOnTransmitDecisionResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.decision = reader.bool();
          break;
        case 2:
          message.from = Peer.decode(reader, reader.uint32());
          break;
        case 3:
          message.received = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OnTransmitDecisionResponse {
    return {
      decision: isSet(object.decision) ? Boolean(object.decision) : false,
      from: isSet(object.from) ? Peer.fromJSON(object.from) : undefined,
      received: isSet(object.received) ? Number(object.received) : 0,
    };
  },

  toJSON(message: OnTransmitDecisionResponse): unknown {
    const obj: any = {};
    message.decision !== undefined && (obj.decision = message.decision);
    message.from !== undefined &&
      (obj.from = message.from ? Peer.toJSON(message.from) : undefined);
    message.received !== undefined &&
      (obj.received = Math.round(message.received));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OnTransmitDecisionResponse>, I>>(
    object: I
  ): OnTransmitDecisionResponse {
    const message = createBaseOnTransmitDecisionResponse();
    message.decision = object.decision ?? false;
    message.from =
      object.from !== undefined && object.from !== null
        ? Peer.fromPartial(object.from)
        : undefined;
    message.received = object.received ?? 0;
    return message;
  },
};

function createBaseOnLobbyRefreshResponse(): OnLobbyRefreshResponse {
  return { olc: "", peers: [], received: 0 };
}

export const OnLobbyRefreshResponse = {
  encode(
    message: OnLobbyRefreshResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.olc !== "") {
      writer.uint32(10).string(message.olc);
    }
    for (const v of message.peers) {
      Peer.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.received !== 0) {
      writer.uint32(24).int64(message.received);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): OnLobbyRefreshResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOnLobbyRefreshResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.olc = reader.string();
          break;
        case 2:
          message.peers.push(Peer.decode(reader, reader.uint32()));
          break;
        case 3:
          message.received = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OnLobbyRefreshResponse {
    return {
      olc: isSet(object.olc) ? String(object.olc) : "",
      peers: Array.isArray(object?.peers)
        ? object.peers.map((e: any) => Peer.fromJSON(e))
        : [],
      received: isSet(object.received) ? Number(object.received) : 0,
    };
  },

  toJSON(message: OnLobbyRefreshResponse): unknown {
    const obj: any = {};
    message.olc !== undefined && (obj.olc = message.olc);
    if (message.peers) {
      obj.peers = message.peers.map((e) => (e ? Peer.toJSON(e) : undefined));
    } else {
      obj.peers = [];
    }
    message.received !== undefined &&
      (obj.received = Math.round(message.received));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OnLobbyRefreshResponse>, I>>(
    object: I
  ): OnLobbyRefreshResponse {
    const message = createBaseOnLobbyRefreshResponse();
    message.olc = object.olc ?? "";
    message.peers = object.peers?.map((e) => Peer.fromPartial(e)) || [];
    message.received = object.received ?? 0;
    return message;
  },
};

function createBaseOnTransmitInviteResponse(): OnTransmitInviteResponse {
  return { received: 0, from: undefined, payload: undefined };
}

export const OnTransmitInviteResponse = {
  encode(
    message: OnTransmitInviteResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.received !== 0) {
      writer.uint32(8).int64(message.received);
    }
    if (message.from !== undefined) {
      Peer.encode(message.from, writer.uint32(18).fork()).ldelim();
    }
    if (message.payload !== undefined) {
      Payload.encode(message.payload, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): OnTransmitInviteResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOnTransmitInviteResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.received = longToNumber(reader.int64() as Long);
          break;
        case 2:
          message.from = Peer.decode(reader, reader.uint32());
          break;
        case 3:
          message.payload = Payload.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OnTransmitInviteResponse {
    return {
      received: isSet(object.received) ? Number(object.received) : 0,
      from: isSet(object.from) ? Peer.fromJSON(object.from) : undefined,
      payload: isSet(object.payload)
        ? Payload.fromJSON(object.payload)
        : undefined,
    };
  },

  toJSON(message: OnTransmitInviteResponse): unknown {
    const obj: any = {};
    message.received !== undefined &&
      (obj.received = Math.round(message.received));
    message.from !== undefined &&
      (obj.from = message.from ? Peer.toJSON(message.from) : undefined);
    message.payload !== undefined &&
      (obj.payload = message.payload
        ? Payload.toJSON(message.payload)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OnTransmitInviteResponse>, I>>(
    object: I
  ): OnTransmitInviteResponse {
    const message = createBaseOnTransmitInviteResponse();
    message.received = object.received ?? 0;
    message.from =
      object.from !== undefined && object.from !== null
        ? Peer.fromPartial(object.from)
        : undefined;
    message.payload =
      object.payload !== undefined && object.payload !== null
        ? Payload.fromPartial(object.payload)
        : undefined;
    return message;
  },
};

function createBaseOnMailboxMessageResponse(): OnMailboxMessageResponse {
  return {
    id: "",
    buffer: Buffer.alloc(0),
    from: undefined,
    to: undefined,
    metadata: undefined,
  };
}

export const OnMailboxMessageResponse = {
  encode(
    message: OnMailboxMessageResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.buffer.length !== 0) {
      writer.uint32(18).bytes(message.buffer);
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
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): OnMailboxMessageResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOnMailboxMessageResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.buffer = reader.bytes() as Buffer;
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
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OnMailboxMessageResponse {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      buffer: isSet(object.buffer)
        ? Buffer.from(bytesFromBase64(object.buffer))
        : Buffer.alloc(0),
      from: isSet(object.from) ? Profile.fromJSON(object.from) : undefined,
      to: isSet(object.to) ? Profile.fromJSON(object.to) : undefined,
      metadata: isSet(object.metadata)
        ? Metadata.fromJSON(object.metadata)
        : undefined,
    };
  },

  toJSON(message: OnMailboxMessageResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.buffer !== undefined &&
      (obj.buffer = base64FromBytes(
        message.buffer !== undefined ? message.buffer : Buffer.alloc(0)
      ));
    message.from !== undefined &&
      (obj.from = message.from ? Profile.toJSON(message.from) : undefined);
    message.to !== undefined &&
      (obj.to = message.to ? Profile.toJSON(message.to) : undefined);
    message.metadata !== undefined &&
      (obj.metadata = message.metadata
        ? Metadata.toJSON(message.metadata)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OnMailboxMessageResponse>, I>>(
    object: I
  ): OnMailboxMessageResponse {
    const message = createBaseOnMailboxMessageResponse();
    message.id = object.id ?? "";
    message.buffer = object.buffer ?? Buffer.alloc(0);
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
    return message;
  },
};

function createBaseOnTransmitProgressResponse(): OnTransmitProgressResponse {
  return {
    progress: 0,
    received: 0,
    current: 0,
    total: 0,
    direction: Direction.DIRECTION_UNSPECIFIED,
  };
}

export const OnTransmitProgressResponse = {
  encode(
    message: OnTransmitProgressResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.progress !== 0) {
      writer.uint32(9).double(message.progress);
    }
    if (message.received !== 0) {
      writer.uint32(16).int64(message.received);
    }
    if (message.current !== 0) {
      writer.uint32(24).int32(message.current);
    }
    if (message.total !== 0) {
      writer.uint32(32).int32(message.total);
    }
    if (message.direction !== Direction.DIRECTION_UNSPECIFIED) {
      writer.uint32(40).int32(directionToNumber(message.direction));
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): OnTransmitProgressResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOnTransmitProgressResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.progress = reader.double();
          break;
        case 2:
          message.received = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.current = reader.int32();
          break;
        case 4:
          message.total = reader.int32();
          break;
        case 5:
          message.direction = directionFromJSON(reader.int32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OnTransmitProgressResponse {
    return {
      progress: isSet(object.progress) ? Number(object.progress) : 0,
      received: isSet(object.received) ? Number(object.received) : 0,
      current: isSet(object.current) ? Number(object.current) : 0,
      total: isSet(object.total) ? Number(object.total) : 0,
      direction: isSet(object.direction)
        ? directionFromJSON(object.direction)
        : Direction.DIRECTION_UNSPECIFIED,
    };
  },

  toJSON(message: OnTransmitProgressResponse): unknown {
    const obj: any = {};
    message.progress !== undefined && (obj.progress = message.progress);
    message.received !== undefined &&
      (obj.received = Math.round(message.received));
    message.current !== undefined &&
      (obj.current = Math.round(message.current));
    message.total !== undefined && (obj.total = Math.round(message.total));
    message.direction !== undefined &&
      (obj.direction = directionToJSON(message.direction));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OnTransmitProgressResponse>, I>>(
    object: I
  ): OnTransmitProgressResponse {
    const message = createBaseOnTransmitProgressResponse();
    message.progress = object.progress ?? 0;
    message.received = object.received ?? 0;
    message.current = object.current ?? 0;
    message.total = object.total ?? 0;
    message.direction = object.direction ?? Direction.DIRECTION_UNSPECIFIED;
    return message;
  },
};

function createBaseOnTransmitCompleteResponse(): OnTransmitCompleteResponse {
  return {
    direction: Direction.DIRECTION_UNSPECIFIED,
    payload: undefined,
    from: undefined,
    to: undefined,
    createdAt: 0,
    receivedAt: 0,
    results: {},
  };
}

export const OnTransmitCompleteResponse = {
  encode(
    message: OnTransmitCompleteResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.direction !== Direction.DIRECTION_UNSPECIFIED) {
      writer.uint32(8).int32(directionToNumber(message.direction));
    }
    if (message.payload !== undefined) {
      Payload.encode(message.payload, writer.uint32(18).fork()).ldelim();
    }
    if (message.from !== undefined) {
      Peer.encode(message.from, writer.uint32(26).fork()).ldelim();
    }
    if (message.to !== undefined) {
      Peer.encode(message.to, writer.uint32(34).fork()).ldelim();
    }
    if (message.createdAt !== 0) {
      writer.uint32(40).int64(message.createdAt);
    }
    if (message.receivedAt !== 0) {
      writer.uint32(48).int64(message.receivedAt);
    }
    Object.entries(message.results).forEach(([key, value]) => {
      OnTransmitCompleteResponse_ResultsEntry.encode(
        { key: key as any, value },
        writer.uint32(58).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): OnTransmitCompleteResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOnTransmitCompleteResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.direction = directionFromJSON(reader.int32());
          break;
        case 2:
          message.payload = Payload.decode(reader, reader.uint32());
          break;
        case 3:
          message.from = Peer.decode(reader, reader.uint32());
          break;
        case 4:
          message.to = Peer.decode(reader, reader.uint32());
          break;
        case 5:
          message.createdAt = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.receivedAt = longToNumber(reader.int64() as Long);
          break;
        case 7:
          const entry7 = OnTransmitCompleteResponse_ResultsEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry7.value !== undefined) {
            message.results[entry7.key] = entry7.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OnTransmitCompleteResponse {
    return {
      direction: isSet(object.direction)
        ? directionFromJSON(object.direction)
        : Direction.DIRECTION_UNSPECIFIED,
      payload: isSet(object.payload)
        ? Payload.fromJSON(object.payload)
        : undefined,
      from: isSet(object.from) ? Peer.fromJSON(object.from) : undefined,
      to: isSet(object.to) ? Peer.fromJSON(object.to) : undefined,
      createdAt: isSet(object.createdAt) ? Number(object.createdAt) : 0,
      receivedAt: isSet(object.receivedAt) ? Number(object.receivedAt) : 0,
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

  toJSON(message: OnTransmitCompleteResponse): unknown {
    const obj: any = {};
    message.direction !== undefined &&
      (obj.direction = directionToJSON(message.direction));
    message.payload !== undefined &&
      (obj.payload = message.payload
        ? Payload.toJSON(message.payload)
        : undefined);
    message.from !== undefined &&
      (obj.from = message.from ? Peer.toJSON(message.from) : undefined);
    message.to !== undefined &&
      (obj.to = message.to ? Peer.toJSON(message.to) : undefined);
    message.createdAt !== undefined &&
      (obj.createdAt = Math.round(message.createdAt));
    message.receivedAt !== undefined &&
      (obj.receivedAt = Math.round(message.receivedAt));
    obj.results = {};
    if (message.results) {
      Object.entries(message.results).forEach(([k, v]) => {
        obj.results[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OnTransmitCompleteResponse>, I>>(
    object: I
  ): OnTransmitCompleteResponse {
    const message = createBaseOnTransmitCompleteResponse();
    message.direction = object.direction ?? Direction.DIRECTION_UNSPECIFIED;
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
    message.createdAt = object.createdAt ?? 0;
    message.receivedAt = object.receivedAt ?? 0;
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

function createBaseOnTransmitCompleteResponse_ResultsEntry(): OnTransmitCompleteResponse_ResultsEntry {
  return { key: 0, value: false };
}

export const OnTransmitCompleteResponse_ResultsEntry = {
  encode(
    message: OnTransmitCompleteResponse_ResultsEntry,
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
  ): OnTransmitCompleteResponse_ResultsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOnTransmitCompleteResponse_ResultsEntry();
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

  fromJSON(object: any): OnTransmitCompleteResponse_ResultsEntry {
    return {
      key: isSet(object.key) ? Number(object.key) : 0,
      value: isSet(object.value) ? Boolean(object.value) : false,
    };
  },

  toJSON(message: OnTransmitCompleteResponse_ResultsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = Math.round(message.key));
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<OnTransmitCompleteResponse_ResultsEntry>, I>
  >(object: I): OnTransmitCompleteResponse_ResultsEntry {
    const message = createBaseOnTransmitCompleteResponse_ResultsEntry();
    message.key = object.key ?? 0;
    message.value = object.value ?? false;
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
