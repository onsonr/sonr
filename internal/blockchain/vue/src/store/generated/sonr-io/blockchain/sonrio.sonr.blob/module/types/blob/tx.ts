/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "sonrio.sonr.blob";

export interface MsgUploadBlob {
  creator: string;
  label: string;
  path: string;
  refDid: string;
  size: number;
  lastModified: number;
}

export interface MsgUploadBlobResponse {}

export interface MsgDownloadBlob {
  creator: string;
  did: string;
  path: string;
  timeout: number;
}

export interface MsgDownloadBlobResponse {}

export interface MsgSyncBlob {
  creator: string;
  did: string;
  path: string;
  timeout: number;
}

export interface MsgSyncBlobResponse {}

export interface MsgDeleteBlob {
  creator: string;
  did: string;
  publicKey: string;
}

export interface MsgDeleteBlobResponse {}

export interface MsgCreateThereIs {
  creator: string;
  index: string;
  did: string;
  documentJson: string;
}

export interface MsgCreateThereIsResponse {}

export interface MsgUpdateThereIs {
  creator: string;
  index: string;
  did: string;
  documentJson: string;
}

export interface MsgUpdateThereIsResponse {}

export interface MsgDeleteThereIs {
  creator: string;
  index: string;
}

export interface MsgDeleteThereIsResponse {}

const baseMsgUploadBlob: object = {
  creator: "",
  label: "",
  path: "",
  refDid: "",
  size: 0,
  lastModified: 0,
};

export const MsgUploadBlob = {
  encode(message: MsgUploadBlob, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.label !== "") {
      writer.uint32(18).string(message.label);
    }
    if (message.path !== "") {
      writer.uint32(26).string(message.path);
    }
    if (message.refDid !== "") {
      writer.uint32(34).string(message.refDid);
    }
    if (message.size !== 0) {
      writer.uint32(40).int32(message.size);
    }
    if (message.lastModified !== 0) {
      writer.uint32(48).int32(message.lastModified);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUploadBlob {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUploadBlob } as MsgUploadBlob;
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
          message.path = reader.string();
          break;
        case 4:
          message.refDid = reader.string();
          break;
        case 5:
          message.size = reader.int32();
          break;
        case 6:
          message.lastModified = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUploadBlob {
    const message = { ...baseMsgUploadBlob } as MsgUploadBlob;
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
    if (object.path !== undefined && object.path !== null) {
      message.path = String(object.path);
    } else {
      message.path = "";
    }
    if (object.refDid !== undefined && object.refDid !== null) {
      message.refDid = String(object.refDid);
    } else {
      message.refDid = "";
    }
    if (object.size !== undefined && object.size !== null) {
      message.size = Number(object.size);
    } else {
      message.size = 0;
    }
    if (object.lastModified !== undefined && object.lastModified !== null) {
      message.lastModified = Number(object.lastModified);
    } else {
      message.lastModified = 0;
    }
    return message;
  },

  toJSON(message: MsgUploadBlob): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.label !== undefined && (obj.label = message.label);
    message.path !== undefined && (obj.path = message.path);
    message.refDid !== undefined && (obj.refDid = message.refDid);
    message.size !== undefined && (obj.size = message.size);
    message.lastModified !== undefined &&
      (obj.lastModified = message.lastModified);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUploadBlob>): MsgUploadBlob {
    const message = { ...baseMsgUploadBlob } as MsgUploadBlob;
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
    if (object.path !== undefined && object.path !== null) {
      message.path = object.path;
    } else {
      message.path = "";
    }
    if (object.refDid !== undefined && object.refDid !== null) {
      message.refDid = object.refDid;
    } else {
      message.refDid = "";
    }
    if (object.size !== undefined && object.size !== null) {
      message.size = object.size;
    } else {
      message.size = 0;
    }
    if (object.lastModified !== undefined && object.lastModified !== null) {
      message.lastModified = object.lastModified;
    } else {
      message.lastModified = 0;
    }
    return message;
  },
};

const baseMsgUploadBlobResponse: object = {};

export const MsgUploadBlobResponse = {
  encode(_: MsgUploadBlobResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUploadBlobResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUploadBlobResponse } as MsgUploadBlobResponse;
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

  fromJSON(_: any): MsgUploadBlobResponse {
    const message = { ...baseMsgUploadBlobResponse } as MsgUploadBlobResponse;
    return message;
  },

  toJSON(_: MsgUploadBlobResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgUploadBlobResponse>): MsgUploadBlobResponse {
    const message = { ...baseMsgUploadBlobResponse } as MsgUploadBlobResponse;
    return message;
  },
};

const baseMsgDownloadBlob: object = {
  creator: "",
  did: "",
  path: "",
  timeout: 0,
};

export const MsgDownloadBlob = {
  encode(message: MsgDownloadBlob, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.path !== "") {
      writer.uint32(26).string(message.path);
    }
    if (message.timeout !== 0) {
      writer.uint32(32).int32(message.timeout);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDownloadBlob {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDownloadBlob } as MsgDownloadBlob;
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
          message.path = reader.string();
          break;
        case 4:
          message.timeout = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDownloadBlob {
    const message = { ...baseMsgDownloadBlob } as MsgDownloadBlob;
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
    if (object.path !== undefined && object.path !== null) {
      message.path = String(object.path);
    } else {
      message.path = "";
    }
    if (object.timeout !== undefined && object.timeout !== null) {
      message.timeout = Number(object.timeout);
    } else {
      message.timeout = 0;
    }
    return message;
  },

  toJSON(message: MsgDownloadBlob): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.path !== undefined && (obj.path = message.path);
    message.timeout !== undefined && (obj.timeout = message.timeout);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDownloadBlob>): MsgDownloadBlob {
    const message = { ...baseMsgDownloadBlob } as MsgDownloadBlob;
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
    if (object.path !== undefined && object.path !== null) {
      message.path = object.path;
    } else {
      message.path = "";
    }
    if (object.timeout !== undefined && object.timeout !== null) {
      message.timeout = object.timeout;
    } else {
      message.timeout = 0;
    }
    return message;
  },
};

const baseMsgDownloadBlobResponse: object = {};

export const MsgDownloadBlobResponse = {
  encode(_: MsgDownloadBlobResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDownloadBlobResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDownloadBlobResponse,
    } as MsgDownloadBlobResponse;
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

  fromJSON(_: any): MsgDownloadBlobResponse {
    const message = {
      ...baseMsgDownloadBlobResponse,
    } as MsgDownloadBlobResponse;
    return message;
  },

  toJSON(_: MsgDownloadBlobResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDownloadBlobResponse>
  ): MsgDownloadBlobResponse {
    const message = {
      ...baseMsgDownloadBlobResponse,
    } as MsgDownloadBlobResponse;
    return message;
  },
};

const baseMsgSyncBlob: object = { creator: "", did: "", path: "", timeout: 0 };

export const MsgSyncBlob = {
  encode(message: MsgSyncBlob, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.path !== "") {
      writer.uint32(26).string(message.path);
    }
    if (message.timeout !== 0) {
      writer.uint32(32).int32(message.timeout);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSyncBlob {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSyncBlob } as MsgSyncBlob;
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
          message.path = reader.string();
          break;
        case 4:
          message.timeout = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSyncBlob {
    const message = { ...baseMsgSyncBlob } as MsgSyncBlob;
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
    if (object.path !== undefined && object.path !== null) {
      message.path = String(object.path);
    } else {
      message.path = "";
    }
    if (object.timeout !== undefined && object.timeout !== null) {
      message.timeout = Number(object.timeout);
    } else {
      message.timeout = 0;
    }
    return message;
  },

  toJSON(message: MsgSyncBlob): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.path !== undefined && (obj.path = message.path);
    message.timeout !== undefined && (obj.timeout = message.timeout);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSyncBlob>): MsgSyncBlob {
    const message = { ...baseMsgSyncBlob } as MsgSyncBlob;
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
    if (object.path !== undefined && object.path !== null) {
      message.path = object.path;
    } else {
      message.path = "";
    }
    if (object.timeout !== undefined && object.timeout !== null) {
      message.timeout = object.timeout;
    } else {
      message.timeout = 0;
    }
    return message;
  },
};

const baseMsgSyncBlobResponse: object = {};

export const MsgSyncBlobResponse = {
  encode(_: MsgSyncBlobResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSyncBlobResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSyncBlobResponse } as MsgSyncBlobResponse;
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

  fromJSON(_: any): MsgSyncBlobResponse {
    const message = { ...baseMsgSyncBlobResponse } as MsgSyncBlobResponse;
    return message;
  },

  toJSON(_: MsgSyncBlobResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSyncBlobResponse>): MsgSyncBlobResponse {
    const message = { ...baseMsgSyncBlobResponse } as MsgSyncBlobResponse;
    return message;
  },
};

const baseMsgDeleteBlob: object = { creator: "", did: "", publicKey: "" };

export const MsgDeleteBlob = {
  encode(message: MsgDeleteBlob, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.publicKey !== "") {
      writer.uint32(26).string(message.publicKey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteBlob {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteBlob } as MsgDeleteBlob;
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
          message.publicKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteBlob {
    const message = { ...baseMsgDeleteBlob } as MsgDeleteBlob;
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
    if (object.publicKey !== undefined && object.publicKey !== null) {
      message.publicKey = String(object.publicKey);
    } else {
      message.publicKey = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteBlob): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.publicKey !== undefined && (obj.publicKey = message.publicKey);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteBlob>): MsgDeleteBlob {
    const message = { ...baseMsgDeleteBlob } as MsgDeleteBlob;
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
    if (object.publicKey !== undefined && object.publicKey !== null) {
      message.publicKey = object.publicKey;
    } else {
      message.publicKey = "";
    }
    return message;
  },
};

const baseMsgDeleteBlobResponse: object = {};

export const MsgDeleteBlobResponse = {
  encode(_: MsgDeleteBlobResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteBlobResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteBlobResponse } as MsgDeleteBlobResponse;
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

  fromJSON(_: any): MsgDeleteBlobResponse {
    const message = { ...baseMsgDeleteBlobResponse } as MsgDeleteBlobResponse;
    return message;
  },

  toJSON(_: MsgDeleteBlobResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgDeleteBlobResponse>): MsgDeleteBlobResponse {
    const message = { ...baseMsgDeleteBlobResponse } as MsgDeleteBlobResponse;
    return message;
  },
};

const baseMsgCreateThereIs: object = {
  creator: "",
  index: "",
  did: "",
  documentJson: "",
};

export const MsgCreateThereIs = {
  encode(message: MsgCreateThereIs, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.did !== "") {
      writer.uint32(26).string(message.did);
    }
    if (message.documentJson !== "") {
      writer.uint32(34).string(message.documentJson);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateThereIs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateThereIs } as MsgCreateThereIs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.did = reader.string();
          break;
        case 4:
          message.documentJson = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateThereIs {
    const message = { ...baseMsgCreateThereIs } as MsgCreateThereIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.documentJson !== undefined && object.documentJson !== null) {
      message.documentJson = String(object.documentJson);
    } else {
      message.documentJson = "";
    }
    return message;
  },

  toJSON(message: MsgCreateThereIs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.did !== undefined && (obj.did = message.did);
    message.documentJson !== undefined &&
      (obj.documentJson = message.documentJson);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateThereIs>): MsgCreateThereIs {
    const message = { ...baseMsgCreateThereIs } as MsgCreateThereIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.documentJson !== undefined && object.documentJson !== null) {
      message.documentJson = object.documentJson;
    } else {
      message.documentJson = "";
    }
    return message;
  },
};

const baseMsgCreateThereIsResponse: object = {};

export const MsgCreateThereIsResponse = {
  encode(
    _: MsgCreateThereIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateThereIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateThereIsResponse,
    } as MsgCreateThereIsResponse;
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

  fromJSON(_: any): MsgCreateThereIsResponse {
    const message = {
      ...baseMsgCreateThereIsResponse,
    } as MsgCreateThereIsResponse;
    return message;
  },

  toJSON(_: MsgCreateThereIsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateThereIsResponse>
  ): MsgCreateThereIsResponse {
    const message = {
      ...baseMsgCreateThereIsResponse,
    } as MsgCreateThereIsResponse;
    return message;
  },
};

const baseMsgUpdateThereIs: object = {
  creator: "",
  index: "",
  did: "",
  documentJson: "",
};

export const MsgUpdateThereIs = {
  encode(message: MsgUpdateThereIs, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.did !== "") {
      writer.uint32(26).string(message.did);
    }
    if (message.documentJson !== "") {
      writer.uint32(34).string(message.documentJson);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateThereIs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateThereIs } as MsgUpdateThereIs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.did = reader.string();
          break;
        case 4:
          message.documentJson = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateThereIs {
    const message = { ...baseMsgUpdateThereIs } as MsgUpdateThereIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.documentJson !== undefined && object.documentJson !== null) {
      message.documentJson = String(object.documentJson);
    } else {
      message.documentJson = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateThereIs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.did !== undefined && (obj.did = message.did);
    message.documentJson !== undefined &&
      (obj.documentJson = message.documentJson);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateThereIs>): MsgUpdateThereIs {
    const message = { ...baseMsgUpdateThereIs } as MsgUpdateThereIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.documentJson !== undefined && object.documentJson !== null) {
      message.documentJson = object.documentJson;
    } else {
      message.documentJson = "";
    }
    return message;
  },
};

const baseMsgUpdateThereIsResponse: object = {};

export const MsgUpdateThereIsResponse = {
  encode(
    _: MsgUpdateThereIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateThereIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateThereIsResponse,
    } as MsgUpdateThereIsResponse;
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

  fromJSON(_: any): MsgUpdateThereIsResponse {
    const message = {
      ...baseMsgUpdateThereIsResponse,
    } as MsgUpdateThereIsResponse;
    return message;
  },

  toJSON(_: MsgUpdateThereIsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateThereIsResponse>
  ): MsgUpdateThereIsResponse {
    const message = {
      ...baseMsgUpdateThereIsResponse,
    } as MsgUpdateThereIsResponse;
    return message;
  },
};

const baseMsgDeleteThereIs: object = { creator: "", index: "" };

export const MsgDeleteThereIs = {
  encode(message: MsgDeleteThereIs, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteThereIs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteThereIs } as MsgDeleteThereIs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteThereIs {
    const message = { ...baseMsgDeleteThereIs } as MsgDeleteThereIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteThereIs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteThereIs>): MsgDeleteThereIs {
    const message = { ...baseMsgDeleteThereIs } as MsgDeleteThereIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgDeleteThereIsResponse: object = {};

export const MsgDeleteThereIsResponse = {
  encode(
    _: MsgDeleteThereIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteThereIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteThereIsResponse,
    } as MsgDeleteThereIsResponse;
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

  fromJSON(_: any): MsgDeleteThereIsResponse {
    const message = {
      ...baseMsgDeleteThereIsResponse,
    } as MsgDeleteThereIsResponse;
    return message;
  },

  toJSON(_: MsgDeleteThereIsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteThereIsResponse>
  ): MsgDeleteThereIsResponse {
    const message = {
      ...baseMsgDeleteThereIsResponse,
    } as MsgDeleteThereIsResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  UploadBlob(request: MsgUploadBlob): Promise<MsgUploadBlobResponse>;
  DownloadBlob(request: MsgDownloadBlob): Promise<MsgDownloadBlobResponse>;
  SyncBlob(request: MsgSyncBlob): Promise<MsgSyncBlobResponse>;
  DeleteBlob(request: MsgDeleteBlob): Promise<MsgDeleteBlobResponse>;
  CreateThereIs(request: MsgCreateThereIs): Promise<MsgCreateThereIsResponse>;
  UpdateThereIs(request: MsgUpdateThereIs): Promise<MsgUpdateThereIsResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteThereIs(request: MsgDeleteThereIs): Promise<MsgDeleteThereIsResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  UploadBlob(request: MsgUploadBlob): Promise<MsgUploadBlobResponse> {
    const data = MsgUploadBlob.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.blob.Msg",
      "UploadBlob",
      data
    );
    return promise.then((data) =>
      MsgUploadBlobResponse.decode(new Reader(data))
    );
  }

  DownloadBlob(request: MsgDownloadBlob): Promise<MsgDownloadBlobResponse> {
    const data = MsgDownloadBlob.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.blob.Msg",
      "DownloadBlob",
      data
    );
    return promise.then((data) =>
      MsgDownloadBlobResponse.decode(new Reader(data))
    );
  }

  SyncBlob(request: MsgSyncBlob): Promise<MsgSyncBlobResponse> {
    const data = MsgSyncBlob.encode(request).finish();
    const promise = this.rpc.request("sonrio.sonr.blob.Msg", "SyncBlob", data);
    return promise.then((data) => MsgSyncBlobResponse.decode(new Reader(data)));
  }

  DeleteBlob(request: MsgDeleteBlob): Promise<MsgDeleteBlobResponse> {
    const data = MsgDeleteBlob.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.blob.Msg",
      "DeleteBlob",
      data
    );
    return promise.then((data) =>
      MsgDeleteBlobResponse.decode(new Reader(data))
    );
  }

  CreateThereIs(request: MsgCreateThereIs): Promise<MsgCreateThereIsResponse> {
    const data = MsgCreateThereIs.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.blob.Msg",
      "CreateThereIs",
      data
    );
    return promise.then((data) =>
      MsgCreateThereIsResponse.decode(new Reader(data))
    );
  }

  UpdateThereIs(request: MsgUpdateThereIs): Promise<MsgUpdateThereIsResponse> {
    const data = MsgUpdateThereIs.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.blob.Msg",
      "UpdateThereIs",
      data
    );
    return promise.then((data) =>
      MsgUpdateThereIsResponse.decode(new Reader(data))
    );
  }

  DeleteThereIs(request: MsgDeleteThereIs): Promise<MsgDeleteThereIsResponse> {
    const data = MsgDeleteThereIs.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.blob.Msg",
      "DeleteThereIs",
      data
    );
    return promise.then((data) =>
      MsgDeleteThereIsResponse.decode(new Reader(data))
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
