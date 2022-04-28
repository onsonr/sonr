/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "ibc.applications.transfer.v1";

/**
 * FungibleTokenPacketData defines a struct for the packet payload
 * See FungibleTokenPacketData spec:
 * https://github.com/cosmos/ics/tree/master/spec/ics-020-fungible-token-transfer#data-structures
 */
export interface FungibleTokenPacketData {
  /** the token denomination to be transferred */
  denom: string;
  /** the token amount to be transferred */
  amount: number;
  /** the sender address */
  sender: string;
  /** the recipient address on the destination chain */
  receiver: string;
}

/**
 * DenomTrace contains the base denomination for ICS20 fungible tokens and the
 * source tracing information path.
 */
export interface DenomTrace {
  /**
   * path defines the chain of port/channel identifiers used for tracing the
   * source of the fungible token.
   */
  path: string;
  /** base denomination of the relayed fungible token. */
  base_denom: string;
}

/**
 * Params defines the set of IBC transfer parameters.
 * NOTE: To prevent a single token from being transferred, set the
 * TransfersEnabled parameter to true and then set the bank module's SendEnabled
 * parameter for the denomination to false.
 */
export interface Params {
  /**
   * send_enabled enables or disables all cross-chain token transfers from this
   * chain.
   */
  send_enabled: boolean;
  /**
   * receive_enabled enables or disables all cross-chain token transfers to this
   * chain.
   */
  receive_enabled: boolean;
}

const baseFungibleTokenPacketData: object = {
  denom: "",
  amount: 0,
  sender: "",
  receiver: "",
};

export const FungibleTokenPacketData = {
  encode(
    message: FungibleTokenPacketData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }
    if (message.amount !== 0) {
      writer.uint32(16).uint64(message.amount);
    }
    if (message.sender !== "") {
      writer.uint32(26).string(message.sender);
    }
    if (message.receiver !== "") {
      writer.uint32(34).string(message.receiver);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): FungibleTokenPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseFungibleTokenPacketData,
    } as FungibleTokenPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;
        case 2:
          message.amount = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.sender = reader.string();
          break;
        case 4:
          message.receiver = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FungibleTokenPacketData {
    const message = {
      ...baseFungibleTokenPacketData,
    } as FungibleTokenPacketData;
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom);
    } else {
      message.denom = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Number(object.amount);
    } else {
      message.amount = 0;
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
    return message;
  },

  toJSON(message: FungibleTokenPacketData): unknown {
    const obj: any = {};
    message.denom !== undefined && (obj.denom = message.denom);
    message.amount !== undefined && (obj.amount = message.amount);
    message.sender !== undefined && (obj.sender = message.sender);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    return obj;
  },

  fromPartial(
    object: DeepPartial<FungibleTokenPacketData>
  ): FungibleTokenPacketData {
    const message = {
      ...baseFungibleTokenPacketData,
    } as FungibleTokenPacketData;
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom;
    } else {
      message.denom = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = 0;
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
    return message;
  },
};

const baseDenomTrace: object = { path: "", base_denom: "" };

export const DenomTrace = {
  encode(message: DenomTrace, writer: Writer = Writer.create()): Writer {
    if (message.path !== "") {
      writer.uint32(10).string(message.path);
    }
    if (message.base_denom !== "") {
      writer.uint32(18).string(message.base_denom);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DenomTrace {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDenomTrace } as DenomTrace;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.path = reader.string();
          break;
        case 2:
          message.base_denom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DenomTrace {
    const message = { ...baseDenomTrace } as DenomTrace;
    if (object.path !== undefined && object.path !== null) {
      message.path = String(object.path);
    } else {
      message.path = "";
    }
    if (object.base_denom !== undefined && object.base_denom !== null) {
      message.base_denom = String(object.base_denom);
    } else {
      message.base_denom = "";
    }
    return message;
  },

  toJSON(message: DenomTrace): unknown {
    const obj: any = {};
    message.path !== undefined && (obj.path = message.path);
    message.base_denom !== undefined && (obj.base_denom = message.base_denom);
    return obj;
  },

  fromPartial(object: DeepPartial<DenomTrace>): DenomTrace {
    const message = { ...baseDenomTrace } as DenomTrace;
    if (object.path !== undefined && object.path !== null) {
      message.path = object.path;
    } else {
      message.path = "";
    }
    if (object.base_denom !== undefined && object.base_denom !== null) {
      message.base_denom = object.base_denom;
    } else {
      message.base_denom = "";
    }
    return message;
  },
};

const baseParams: object = { send_enabled: false, receive_enabled: false };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.send_enabled === true) {
      writer.uint32(8).bool(message.send_enabled);
    }
    if (message.receive_enabled === true) {
      writer.uint32(16).bool(message.receive_enabled);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.send_enabled = reader.bool();
          break;
        case 2:
          message.receive_enabled = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    if (object.send_enabled !== undefined && object.send_enabled !== null) {
      message.send_enabled = Boolean(object.send_enabled);
    } else {
      message.send_enabled = false;
    }
    if (
      object.receive_enabled !== undefined &&
      object.receive_enabled !== null
    ) {
      message.receive_enabled = Boolean(object.receive_enabled);
    } else {
      message.receive_enabled = false;
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.send_enabled !== undefined &&
      (obj.send_enabled = message.send_enabled);
    message.receive_enabled !== undefined &&
      (obj.receive_enabled = message.receive_enabled);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.send_enabled !== undefined && object.send_enabled !== null) {
      message.send_enabled = object.send_enabled;
    } else {
      message.send_enabled = false;
    }
    if (
      object.receive_enabled !== undefined &&
      object.receive_enabled !== null
    ) {
      message.receive_enabled = object.receive_enabled;
    } else {
      message.receive_enabled = false;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
