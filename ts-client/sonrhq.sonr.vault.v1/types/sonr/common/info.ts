/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "sonrhq.sonr.common";

/** Account is used for storing all credentials and their locations to be encrypted. */
export interface AccountInfo {
  /** Address is the associated Sonr address. */
  address: string;
  /** Credentials is a list of all credentials associated with the account. */
  network: string;
  /** Label is the label of the account. */
  label: string;
  /** Index is the index of the account. */
  index: number;
  /** Balance is the balance of the account. */
  balance: number;
}

/** Basic Info Sent to Peers to Establish Connections */
export interface PeerInfo {
  /** User Sonr Account Decentralized ID */
  id: string;
  /** User Defined Label for Peer, also known as PartyID */
  name: string;
  /** Peer ID */
  peerId: string;
  /** Peer Multiaddress */
  multiaddr: string;
}

export interface WalletInfo {
  /** Controller is the associated Sonr address. */
  controller: string;
  /** DiscoverPaths is a list of all known hardened coin type paths. */
  discoveredPaths: number[];
  /** Algorithm is the algorithm of the wallet. CMP is the default. */
  algorithm: string;
  /** CreatedAt is the time the wallet was created. */
  createdAt: number;
  /** LastUpdated is the last time the wallet was updated. */
  lastUpdated: number;
}

function createBaseAccountInfo(): AccountInfo {
  return { address: "", network: "", label: "", index: 0, balance: 0 };
}

export const AccountInfo = {
  encode(message: AccountInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.network !== "") {
      writer.uint32(18).string(message.network);
    }
    if (message.label !== "") {
      writer.uint32(26).string(message.label);
    }
    if (message.index !== 0) {
      writer.uint32(32).uint32(message.index);
    }
    if (message.balance !== 0) {
      writer.uint32(40).int32(message.balance);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AccountInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAccountInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.network = reader.string();
          break;
        case 3:
          message.label = reader.string();
          break;
        case 4:
          message.index = reader.uint32();
          break;
        case 5:
          message.balance = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AccountInfo {
    return {
      address: isSet(object.address) ? String(object.address) : "",
      network: isSet(object.network) ? String(object.network) : "",
      label: isSet(object.label) ? String(object.label) : "",
      index: isSet(object.index) ? Number(object.index) : 0,
      balance: isSet(object.balance) ? Number(object.balance) : 0,
    };
  },

  toJSON(message: AccountInfo): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.network !== undefined && (obj.network = message.network);
    message.label !== undefined && (obj.label = message.label);
    message.index !== undefined && (obj.index = Math.round(message.index));
    message.balance !== undefined && (obj.balance = Math.round(message.balance));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AccountInfo>, I>>(object: I): AccountInfo {
    const message = createBaseAccountInfo();
    message.address = object.address ?? "";
    message.network = object.network ?? "";
    message.label = object.label ?? "";
    message.index = object.index ?? 0;
    message.balance = object.balance ?? 0;
    return message;
  },
};

function createBasePeerInfo(): PeerInfo {
  return { id: "", name: "", peerId: "", multiaddr: "" };
}

export const PeerInfo = {
  encode(message: PeerInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.peerId !== "") {
      writer.uint32(26).string(message.peerId);
    }
    if (message.multiaddr !== "") {
      writer.uint32(34).string(message.multiaddr);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PeerInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePeerInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.peerId = reader.string();
          break;
        case 4:
          message.multiaddr = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PeerInfo {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      name: isSet(object.name) ? String(object.name) : "",
      peerId: isSet(object.peerId) ? String(object.peerId) : "",
      multiaddr: isSet(object.multiaddr) ? String(object.multiaddr) : "",
    };
  },

  toJSON(message: PeerInfo): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.peerId !== undefined && (obj.peerId = message.peerId);
    message.multiaddr !== undefined && (obj.multiaddr = message.multiaddr);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PeerInfo>, I>>(object: I): PeerInfo {
    const message = createBasePeerInfo();
    message.id = object.id ?? "";
    message.name = object.name ?? "";
    message.peerId = object.peerId ?? "";
    message.multiaddr = object.multiaddr ?? "";
    return message;
  },
};

function createBaseWalletInfo(): WalletInfo {
  return { controller: "", discoveredPaths: [], algorithm: "", createdAt: 0, lastUpdated: 0 };
}

export const WalletInfo = {
  encode(message: WalletInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.controller !== "") {
      writer.uint32(10).string(message.controller);
    }
    writer.uint32(18).fork();
    for (const v of message.discoveredPaths) {
      writer.int32(v);
    }
    writer.ldelim();
    if (message.algorithm !== "") {
      writer.uint32(26).string(message.algorithm);
    }
    if (message.createdAt !== 0) {
      writer.uint32(32).int64(message.createdAt);
    }
    if (message.lastUpdated !== 0) {
      writer.uint32(40).int64(message.lastUpdated);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WalletInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWalletInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.controller = reader.string();
          break;
        case 2:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.discoveredPaths.push(reader.int32());
            }
          } else {
            message.discoveredPaths.push(reader.int32());
          }
          break;
        case 3:
          message.algorithm = reader.string();
          break;
        case 4:
          message.createdAt = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.lastUpdated = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WalletInfo {
    return {
      controller: isSet(object.controller) ? String(object.controller) : "",
      discoveredPaths: Array.isArray(object?.discoveredPaths) ? object.discoveredPaths.map((e: any) => Number(e)) : [],
      algorithm: isSet(object.algorithm) ? String(object.algorithm) : "",
      createdAt: isSet(object.createdAt) ? Number(object.createdAt) : 0,
      lastUpdated: isSet(object.lastUpdated) ? Number(object.lastUpdated) : 0,
    };
  },

  toJSON(message: WalletInfo): unknown {
    const obj: any = {};
    message.controller !== undefined && (obj.controller = message.controller);
    if (message.discoveredPaths) {
      obj.discoveredPaths = message.discoveredPaths.map((e) => Math.round(e));
    } else {
      obj.discoveredPaths = [];
    }
    message.algorithm !== undefined && (obj.algorithm = message.algorithm);
    message.createdAt !== undefined && (obj.createdAt = Math.round(message.createdAt));
    message.lastUpdated !== undefined && (obj.lastUpdated = Math.round(message.lastUpdated));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<WalletInfo>, I>>(object: I): WalletInfo {
    const message = createBaseWalletInfo();
    message.controller = object.controller ?? "";
    message.discoveredPaths = object.discoveredPaths?.map((e) => e) || [];
    message.algorithm = object.algorithm ?? "";
    message.createdAt = object.createdAt ?? 0;
    message.lastUpdated = object.lastUpdated ?? 0;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
