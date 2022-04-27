/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { ObjectField } from "../object/object";

export const protobufPackage = "sonrio.sonr.object";

export interface MsgCreateObject {
  creator: string;
  label: string;
  description: string;
  fields: ObjectField[];
}

export interface MsgCreateObjectResponse {}

export interface MsgReadObject {
  creator: string;
  did: string;
}

export interface MsgReadObjectResponse {}

export interface MsgUpdateObject {
  creator: string;
  did: string;
}

export interface MsgUpdateObjectResponse {}

export interface MsgDeactivateObject {
  creator: string;
  did: string;
  publicKey: string;
}

export interface MsgDeactivateObjectResponse {}

const baseMsgCreateObject: object = { creator: "", label: "", description: "" };

export const MsgCreateObject = {
  encode(message: MsgCreateObject, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.label !== "") {
      writer.uint32(18).string(message.label);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    for (const v of message.fields) {
      ObjectField.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateObject {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateObject } as MsgCreateObject;
    message.fields = [];
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
          message.fields.push(ObjectField.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateObject {
    const message = { ...baseMsgCreateObject } as MsgCreateObject;
    message.fields = [];
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
    if (object.fields !== undefined && object.fields !== null) {
      for (const e of object.fields) {
        message.fields.push(ObjectField.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: MsgCreateObject): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.label !== undefined && (obj.label = message.label);
    message.description !== undefined &&
      (obj.description = message.description);
    if (message.fields) {
      obj.fields = message.fields.map((e) =>
        e ? ObjectField.toJSON(e) : undefined
      );
    } else {
      obj.fields = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateObject>): MsgCreateObject {
    const message = { ...baseMsgCreateObject } as MsgCreateObject;
    message.fields = [];
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
    if (object.fields !== undefined && object.fields !== null) {
      for (const e of object.fields) {
        message.fields.push(ObjectField.fromPartial(e));
      }
    }
    return message;
  },
};

const baseMsgCreateObjectResponse: object = {};

export const MsgCreateObjectResponse = {
  encode(_: MsgCreateObjectResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateObjectResponse,
    } as MsgCreateObjectResponse;
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

  fromJSON(_: any): MsgCreateObjectResponse {
    const message = {
      ...baseMsgCreateObjectResponse,
    } as MsgCreateObjectResponse;
    return message;
  },

  toJSON(_: MsgCreateObjectResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateObjectResponse>
  ): MsgCreateObjectResponse {
    const message = {
      ...baseMsgCreateObjectResponse,
    } as MsgCreateObjectResponse;
    return message;
  },
};

const baseMsgReadObject: object = { creator: "", did: "" };

export const MsgReadObject = {
  encode(message: MsgReadObject, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgReadObject {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgReadObject } as MsgReadObject;
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

  fromJSON(object: any): MsgReadObject {
    const message = { ...baseMsgReadObject } as MsgReadObject;
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

  toJSON(message: MsgReadObject): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgReadObject>): MsgReadObject {
    const message = { ...baseMsgReadObject } as MsgReadObject;
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

const baseMsgReadObjectResponse: object = {};

export const MsgReadObjectResponse = {
  encode(_: MsgReadObjectResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgReadObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgReadObjectResponse } as MsgReadObjectResponse;
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

  fromJSON(_: any): MsgReadObjectResponse {
    const message = { ...baseMsgReadObjectResponse } as MsgReadObjectResponse;
    return message;
  },

  toJSON(_: MsgReadObjectResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgReadObjectResponse>): MsgReadObjectResponse {
    const message = { ...baseMsgReadObjectResponse } as MsgReadObjectResponse;
    return message;
  },
};

const baseMsgUpdateObject: object = { creator: "", did: "" };

export const MsgUpdateObject = {
  encode(message: MsgUpdateObject, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateObject {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateObject } as MsgUpdateObject;
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

  fromJSON(object: any): MsgUpdateObject {
    const message = { ...baseMsgUpdateObject } as MsgUpdateObject;
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

  toJSON(message: MsgUpdateObject): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateObject>): MsgUpdateObject {
    const message = { ...baseMsgUpdateObject } as MsgUpdateObject;
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

const baseMsgUpdateObjectResponse: object = {};

export const MsgUpdateObjectResponse = {
  encode(_: MsgUpdateObjectResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateObjectResponse,
    } as MsgUpdateObjectResponse;
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

  fromJSON(_: any): MsgUpdateObjectResponse {
    const message = {
      ...baseMsgUpdateObjectResponse,
    } as MsgUpdateObjectResponse;
    return message;
  },

  toJSON(_: MsgUpdateObjectResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateObjectResponse>
  ): MsgUpdateObjectResponse {
    const message = {
      ...baseMsgUpdateObjectResponse,
    } as MsgUpdateObjectResponse;
    return message;
  },
};

const baseMsgDeactivateObject: object = { creator: "", did: "", publicKey: "" };

export const MsgDeactivateObject = {
  encode(message: MsgDeactivateObject, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): MsgDeactivateObject {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeactivateObject } as MsgDeactivateObject;
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

  fromJSON(object: any): MsgDeactivateObject {
    const message = { ...baseMsgDeactivateObject } as MsgDeactivateObject;
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

  toJSON(message: MsgDeactivateObject): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.publicKey !== undefined && (obj.publicKey = message.publicKey);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeactivateObject>): MsgDeactivateObject {
    const message = { ...baseMsgDeactivateObject } as MsgDeactivateObject;
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

const baseMsgDeactivateObjectResponse: object = {};

export const MsgDeactivateObjectResponse = {
  encode(_: MsgDeactivateObjectResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeactivateObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeactivateObjectResponse,
    } as MsgDeactivateObjectResponse;
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

  fromJSON(_: any): MsgDeactivateObjectResponse {
    const message = {
      ...baseMsgDeactivateObjectResponse,
    } as MsgDeactivateObjectResponse;
    return message;
  },

  toJSON(_: MsgDeactivateObjectResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeactivateObjectResponse>
  ): MsgDeactivateObjectResponse {
    const message = {
      ...baseMsgDeactivateObjectResponse,
    } as MsgDeactivateObjectResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateObject(request: MsgCreateObject): Promise<MsgCreateObjectResponse>;
  ReadObject(request: MsgReadObject): Promise<MsgReadObjectResponse>;
  UpdateObject(request: MsgUpdateObject): Promise<MsgUpdateObjectResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteObject(request: MsgDeactivateObject): Promise<MsgDeactivateObjectResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateObject(request: MsgCreateObject): Promise<MsgCreateObjectResponse> {
    const data = MsgCreateObject.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.object.Msg",
      "CreateObject",
      data
    );
    return promise.then((data) =>
      MsgCreateObjectResponse.decode(new Reader(data))
    );
  }

  ReadObject(request: MsgReadObject): Promise<MsgReadObjectResponse> {
    const data = MsgReadObject.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.object.Msg",
      "ReadObject",
      data
    );
    return promise.then((data) =>
      MsgReadObjectResponse.decode(new Reader(data))
    );
  }

  UpdateObject(request: MsgUpdateObject): Promise<MsgUpdateObjectResponse> {
    const data = MsgUpdateObject.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.object.Msg",
      "UpdateObject",
      data
    );
    return promise.then((data) =>
      MsgUpdateObjectResponse.decode(new Reader(data))
    );
  }

  DeleteObject(request: MsgDeactivateObject): Promise<MsgDeactivateObjectResponse> {
    const data = MsgDeactivateObject.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.object.Msg",
      "DeleteObject",
      data
    );
    return promise.then((data) =>
      MsgDeactivateObjectResponse.decode(new Reader(data))
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
