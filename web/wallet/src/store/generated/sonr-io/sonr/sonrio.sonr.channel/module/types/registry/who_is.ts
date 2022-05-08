/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sonrio.sonr.registry";

export enum WhoIsType {
  /** USER - User is the type of the registered name */
  USER = 0,
  /** APPLICATION - Application is the type of the registered name */
  APPLICATION = 1,
  UNRECOGNIZED = -1,
}

export function whoIsTypeFromJSON(object: any): WhoIsType {
  switch (object) {
    case 0:
    case "USER":
      return WhoIsType.USER;
    case 1:
    case "APPLICATION":
      return WhoIsType.APPLICATION;
    case -1:
    case "UNRECOGNIZED":
    default:
      return WhoIsType.UNRECOGNIZED;
  }
}

export function whoIsTypeToJSON(object: WhoIsType): string {
  switch (object) {
    case WhoIsType.USER:
      return "USER";
    case WhoIsType.APPLICATION:
      return "APPLICATION";
    default:
      return "UNKNOWN";
  }
}

export interface WhoIs {
  /** Alias is the list of registered `alsoKnownAs` identifiers of the User or Application */
  alias: string[];
  /** Owner is the top level DID of the User or Application derived from the multisignature wallet. */
  owner: string;
  /** DIDDocument is the bytes representation of DIDDocument within the WhoIs */
  did_document: Uint8Array;
  /** Credentials are the biometric info of the registered name and account encoded with public key */
  controllers: string[];
  /** Type is the type of the registered name */
  type: WhoIsType;
  /** Timestamp is the time of the last update of the DID Document */
  timestamp: number;
  /** IsActive is the status of the DID Document */
  is_active: boolean;
}

const baseWhoIs: object = {
  alias: "",
  owner: "",
  controllers: "",
  type: 0,
  timestamp: 0,
  is_active: false,
};

export const WhoIs = {
  encode(message: WhoIs, writer: Writer = Writer.create()): Writer {
    for (const v of message.alias) {
      writer.uint32(10).string(v!);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.did_document.length !== 0) {
      writer.uint32(26).bytes(message.did_document);
    }
    for (const v of message.controllers) {
      writer.uint32(34).string(v!);
    }
    if (message.type !== 0) {
      writer.uint32(40).int32(message.type);
    }
    if (message.timestamp !== 0) {
      writer.uint32(48).int64(message.timestamp);
    }
    if (message.is_active === true) {
      writer.uint32(56).bool(message.is_active);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): WhoIs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseWhoIs } as WhoIs;
    message.alias = [];
    message.controllers = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.alias.push(reader.string());
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 3:
          message.did_document = reader.bytes();
          break;
        case 4:
          message.controllers.push(reader.string());
          break;
        case 5:
          message.type = reader.int32() as any;
          break;
        case 6:
          message.timestamp = longToNumber(reader.int64() as Long);
          break;
        case 7:
          message.is_active = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WhoIs {
    const message = { ...baseWhoIs } as WhoIs;
    message.alias = [];
    message.controllers = [];
    if (object.alias !== undefined && object.alias !== null) {
      for (const e of object.alias) {
        message.alias.push(String(e));
      }
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.did_document !== undefined && object.did_document !== null) {
      message.did_document = bytesFromBase64(object.did_document);
    }
    if (object.controllers !== undefined && object.controllers !== null) {
      for (const e of object.controllers) {
        message.controllers.push(String(e));
      }
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = whoIsTypeFromJSON(object.type);
    } else {
      message.type = 0;
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = Number(object.timestamp);
    } else {
      message.timestamp = 0;
    }
    if (object.is_active !== undefined && object.is_active !== null) {
      message.is_active = Boolean(object.is_active);
    } else {
      message.is_active = false;
    }
    return message;
  },

  toJSON(message: WhoIs): unknown {
    const obj: any = {};
    if (message.alias) {
      obj.alias = message.alias.map((e) => e);
    } else {
      obj.alias = [];
    }
    message.owner !== undefined && (obj.owner = message.owner);
    message.did_document !== undefined &&
      (obj.did_document = base64FromBytes(
        message.did_document !== undefined
          ? message.did_document
          : new Uint8Array()
      ));
    if (message.controllers) {
      obj.controllers = message.controllers.map((e) => e);
    } else {
      obj.controllers = [];
    }
    message.type !== undefined && (obj.type = whoIsTypeToJSON(message.type));
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    message.is_active !== undefined && (obj.is_active = message.is_active);
    return obj;
  },

  fromPartial(object: DeepPartial<WhoIs>): WhoIs {
    const message = { ...baseWhoIs } as WhoIs;
    message.alias = [];
    message.controllers = [];
    if (object.alias !== undefined && object.alias !== null) {
      for (const e of object.alias) {
        message.alias.push(e);
      }
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.did_document !== undefined && object.did_document !== null) {
      message.did_document = object.did_document;
    } else {
      message.did_document = new Uint8Array();
    }
    if (object.controllers !== undefined && object.controllers !== null) {
      for (const e of object.controllers) {
        message.controllers.push(e);
      }
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = 0;
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = object.timestamp;
    } else {
      message.timestamp = 0;
    }
    if (object.is_active !== undefined && object.is_active !== null) {
      message.is_active = object.is_active;
    } else {
      message.is_active = false;
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
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
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
