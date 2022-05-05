/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Credential } from "../registry/credential";

export const protobufPackage = "sonrio.sonr.registry";

/** WhoIs is the entry pointing a registered name to a user account address, Did Url string, and a DIDDocument. */
export interface WhoIs {
  /** Name is the registered name of the User or Application */
  name: string;
  /** DID is the DID of the account */
  did: string;
  /** Document is the DID Document of the registered name and account encoded as JSON */
  document: Uint8Array;
  /** Creator is the Account Address of the creator of the DID Document */
  creator: string;
  /** Credentials are the biometric info of the registered name and account encoded with public key */
  credentials: Credential[];
  /** Type is the type of the registered name */
  type: WhoIs_Type;
  /** Additional Metadata for associated WhoIs */
  metadata: { [key: string]: string };
  /** Timestamp is the time of the last update of the DID Document */
  timestamp: number;
  /** IsActive is the status of the DID Document */
  is_active: boolean;
}

/** Type is the type of the registered name */
export enum WhoIs_Type {
  /** User - User is the type of the registered name */
  User = 0,
  /** Application - Application is the type of the registered name */
  Application = 1,
  UNRECOGNIZED = -1,
}

export function whoIs_TypeFromJSON(object: any): WhoIs_Type {
  switch (object) {
    case 0:
    case "User":
      return WhoIs_Type.User;
    case 1:
    case "Application":
      return WhoIs_Type.Application;
    case -1:
    case "UNRECOGNIZED":
    default:
      return WhoIs_Type.UNRECOGNIZED;
  }
}

export function whoIs_TypeToJSON(object: WhoIs_Type): string {
  switch (object) {
    case WhoIs_Type.User:
      return "User";
    case WhoIs_Type.Application:
      return "Application";
    default:
      return "UNKNOWN";
  }
}

export interface WhoIs_MetadataEntry {
  key: string;
  value: string;
}

/** Session is the metadata for current user or application */
export interface Session {
  /** Base DID is the current Account or Application whois DID url */
  base_did: string;
  /** WhoIs is the current Document for the DID */
  whois: WhoIs | undefined;
  /** Credential is the current Credential for the DID */
  credential: Credential | undefined;
}

const baseWhoIs: object = {
  name: "",
  did: "",
  creator: "",
  type: 0,
  timestamp: 0,
  is_active: false,
};

export const WhoIs = {
  encode(message: WhoIs, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.document.length !== 0) {
      writer.uint32(26).bytes(message.document);
    }
    if (message.creator !== "") {
      writer.uint32(34).string(message.creator);
    }
    for (const v of message.credentials) {
      Credential.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.type !== 0) {
      writer.uint32(48).int32(message.type);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      WhoIs_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(58).fork()
      ).ldelim();
    });
    if (message.timestamp !== 0) {
      writer.uint32(64).int64(message.timestamp);
    }
    if (message.is_active === true) {
      writer.uint32(72).bool(message.is_active);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): WhoIs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseWhoIs } as WhoIs;
    message.credentials = [];
    message.metadata = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.document = reader.bytes();
          break;
        case 4:
          message.creator = reader.string();
          break;
        case 5:
          message.credentials.push(Credential.decode(reader, reader.uint32()));
          break;
        case 6:
          message.type = reader.int32() as any;
          break;
        case 7:
          const entry7 = WhoIs_MetadataEntry.decode(reader, reader.uint32());
          if (entry7.value !== undefined) {
            message.metadata[entry7.key] = entry7.value;
          }
          break;
        case 8:
          message.timestamp = longToNumber(reader.int64() as Long);
          break;
        case 9:
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
    message.credentials = [];
    message.metadata = {};
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.document !== undefined && object.document !== null) {
      message.document = bytesFromBase64(object.document);
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.credentials !== undefined && object.credentials !== null) {
      for (const e of object.credentials) {
        message.credentials.push(Credential.fromJSON(e));
      }
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = whoIs_TypeFromJSON(object.type);
    } else {
      message.type = 0;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
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
    message.name !== undefined && (obj.name = message.name);
    message.did !== undefined && (obj.did = message.did);
    message.document !== undefined &&
      (obj.document = base64FromBytes(
        message.document !== undefined ? message.document : new Uint8Array()
      ));
    message.creator !== undefined && (obj.creator = message.creator);
    if (message.credentials) {
      obj.credentials = message.credentials.map((e) =>
        e ? Credential.toJSON(e) : undefined
      );
    } else {
      obj.credentials = [];
    }
    message.type !== undefined && (obj.type = whoIs_TypeToJSON(message.type));
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    message.is_active !== undefined && (obj.is_active = message.is_active);
    return obj;
  },

  fromPartial(object: DeepPartial<WhoIs>): WhoIs {
    const message = { ...baseWhoIs } as WhoIs;
    message.credentials = [];
    message.metadata = {};
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.document !== undefined && object.document !== null) {
      message.document = object.document;
    } else {
      message.document = new Uint8Array();
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.credentials !== undefined && object.credentials !== null) {
      for (const e of object.credentials) {
        message.credentials.push(Credential.fromPartial(e));
      }
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = 0;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
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

const baseWhoIs_MetadataEntry: object = { key: "", value: "" };

export const WhoIs_MetadataEntry = {
  encode(
    message: WhoIs_MetadataEntry,
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

  decode(input: Reader | Uint8Array, length?: number): WhoIs_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseWhoIs_MetadataEntry } as WhoIs_MetadataEntry;
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

  fromJSON(object: any): WhoIs_MetadataEntry {
    const message = { ...baseWhoIs_MetadataEntry } as WhoIs_MetadataEntry;
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

  toJSON(message: WhoIs_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(object: DeepPartial<WhoIs_MetadataEntry>): WhoIs_MetadataEntry {
    const message = { ...baseWhoIs_MetadataEntry } as WhoIs_MetadataEntry;
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

const baseSession: object = { base_did: "" };

export const Session = {
  encode(message: Session, writer: Writer = Writer.create()): Writer {
    if (message.base_did !== "") {
      writer.uint32(10).string(message.base_did);
    }
    if (message.whois !== undefined) {
      WhoIs.encode(message.whois, writer.uint32(18).fork()).ldelim();
    }
    if (message.credential !== undefined) {
      Credential.encode(message.credential, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Session {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSession } as Session;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.base_did = reader.string();
          break;
        case 2:
          message.whois = WhoIs.decode(reader, reader.uint32());
          break;
        case 3:
          message.credential = Credential.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Session {
    const message = { ...baseSession } as Session;
    if (object.base_did !== undefined && object.base_did !== null) {
      message.base_did = String(object.base_did);
    } else {
      message.base_did = "";
    }
    if (object.whois !== undefined && object.whois !== null) {
      message.whois = WhoIs.fromJSON(object.whois);
    } else {
      message.whois = undefined;
    }
    if (object.credential !== undefined && object.credential !== null) {
      message.credential = Credential.fromJSON(object.credential);
    } else {
      message.credential = undefined;
    }
    return message;
  },

  toJSON(message: Session): unknown {
    const obj: any = {};
    message.base_did !== undefined && (obj.base_did = message.base_did);
    message.whois !== undefined &&
      (obj.whois = message.whois ? WhoIs.toJSON(message.whois) : undefined);
    message.credential !== undefined &&
      (obj.credential = message.credential
        ? Credential.toJSON(message.credential)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Session>): Session {
    const message = { ...baseSession } as Session;
    if (object.base_did !== undefined && object.base_did !== null) {
      message.base_did = object.base_did;
    } else {
      message.base_did = "";
    }
    if (object.whois !== undefined && object.whois !== null) {
      message.whois = WhoIs.fromPartial(object.whois);
    } else {
      message.whois = undefined;
    }
    if (object.credential !== undefined && object.credential !== null) {
      message.credential = Credential.fromPartial(object.credential);
    } else {
      message.credential = undefined;
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
