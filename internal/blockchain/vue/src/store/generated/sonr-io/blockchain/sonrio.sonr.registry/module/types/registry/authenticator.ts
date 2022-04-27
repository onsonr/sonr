/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sonrio.sonr.registry";

export interface Authenticator {
  /**
   * The AAGUID of the authenticator. An AAGUID is defined as an array containing the globally unique
   * identifier of the authenticator model being sought.
   */
  aaguid: Uint8Array;
  /**
   * SignCount -Upon a new login operation, the Relying Party compares the stored signature counter value
   * with the new sign_count value returned in the assertion’s authenticator data. If this new
   * signCount value is less than or equal to the stored value, a cloned authenticator may
   * exist, or the authenticator may be malfunctioning.
   */
  sign_count: number;
  /**
   * CloneWarning - This is a signal that the authenticator may be cloned, i.e. at least two copies of the
   * credential private key may exist and are being used in parallel. Relying Parties should incorporate
   * this information into their risk scoring. Whether the Relying Party updates the stored signature
   * counter value in this case, or not, or fails the authentication ceremony or not, is Relying Party-specific.
   */
  clone_warning: boolean;
}

const baseAuthenticator: object = { sign_count: 0, clone_warning: false };

export const Authenticator = {
  encode(message: Authenticator, writer: Writer = Writer.create()): Writer {
    if (message.aaguid.length !== 0) {
      writer.uint32(10).bytes(message.aaguid);
    }
    if (message.sign_count !== 0) {
      writer.uint32(16).uint32(message.sign_count);
    }
    if (message.clone_warning === true) {
      writer.uint32(24).bool(message.clone_warning);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Authenticator {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAuthenticator } as Authenticator;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.aaguid = reader.bytes();
          break;
        case 2:
          message.sign_count = reader.uint32();
          break;
        case 3:
          message.clone_warning = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Authenticator {
    const message = { ...baseAuthenticator } as Authenticator;
    if (object.aaguid !== undefined && object.aaguid !== null) {
      message.aaguid = bytesFromBase64(object.aaguid);
    }
    if (object.sign_count !== undefined && object.sign_count !== null) {
      message.sign_count = Number(object.sign_count);
    } else {
      message.sign_count = 0;
    }
    if (object.clone_warning !== undefined && object.clone_warning !== null) {
      message.clone_warning = Boolean(object.clone_warning);
    } else {
      message.clone_warning = false;
    }
    return message;
  },

  toJSON(message: Authenticator): unknown {
    const obj: any = {};
    message.aaguid !== undefined &&
      (obj.aaguid = base64FromBytes(
        message.aaguid !== undefined ? message.aaguid : new Uint8Array()
      ));
    message.sign_count !== undefined && (obj.sign_count = message.sign_count);
    message.clone_warning !== undefined &&
      (obj.clone_warning = message.clone_warning);
    return obj;
  },

  fromPartial(object: DeepPartial<Authenticator>): Authenticator {
    const message = { ...baseAuthenticator } as Authenticator;
    if (object.aaguid !== undefined && object.aaguid !== null) {
      message.aaguid = object.aaguid;
    } else {
      message.aaguid = new Uint8Array();
    }
    if (object.sign_count !== undefined && object.sign_count !== null) {
      message.sign_count = object.sign_count;
    } else {
      message.sign_count = 0;
    }
    if (object.clone_warning !== undefined && object.clone_warning !== null) {
      message.clone_warning = object.clone_warning;
    } else {
      message.clone_warning = false;
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
