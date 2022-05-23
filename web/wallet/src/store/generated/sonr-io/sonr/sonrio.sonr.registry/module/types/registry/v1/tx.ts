/* eslint-disable */
import {
  WhoIsType,
  WhoIs,
  whoIsTypeFromJSON,
  whoIsTypeToJSON,
} from "../../registry/v1/who_is";
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "sonrio.sonr.registry";

/** swagger:model MsgCreateWhoIs */
export interface MsgCreateWhoIs {
  /** Creator is the wallet address of the creator of the transaction. */
  creator: string;
  /** DidDocument is the DID document to be stored, in JSON format (see https://w3c-ccg.github.io/did-spec/#did-json-ld). */
  did_document: Uint8Array;
  /** WhoIsType is the type of the WhoIs to be created. Possible values are: "USER", "APPLICATION". */
  whois_type: WhoIsType;
}

export interface MsgCreateWhoIsResponse {
  /** Did is the top level DID of the created WhoIs. */
  success: boolean;
  /** WhoIs is the created WhoIs, contains the DID document and associated metadata. */
  who_is: WhoIs | undefined;
}

export interface MsgUpdateWhoIs {
  /** Creator is the wallet address of the creator of the transaction. */
  creator: string;
  /** DidDocument is the DID document to be stored, in JSON format (see https://w3c-ccg.github.io/did-spec/#did-json-ld). */
  did_document: Uint8Array;
}

export interface MsgUpdateWhoIsResponse {
  /** Did is the top level DID of the WhoIs. */
  success: boolean;
  /** WhoIs is the created WhoIs, contains the DID document and associated metadata. */
  who_is: WhoIs | undefined;
}

export interface MsgDeactivateWhoIs {
  /** Creator is the wallet address of the creator of the transaction. */
  creator: string;
}

export interface MsgDeactivateWhoIsResponse {
  /** Success is true if the WhoIs was successfully deactivated. */
  success: boolean;
  /** Did is the top level DID of the WhoIs. */
  did: string;
}

export interface MsgBuyAlias {
  /** Creator is the wallet address of the creator of the transaction. */
  creator: string;
  /** Name is the name of the alias app extension to be bought. i.e. example.snr/{name} */
  name: string;
}

export interface MsgBuyAliasResponse {
  /** Did is the top level DID of the WhoIs. */
  success: boolean;
  /** WhoIs is the updated WhoIs, contains the DID document and associated metadata. */
  who_is: WhoIs | undefined;
}

export interface MsgTransferAlias {
  /** Creator is the wallet address of the creator of the transaction. */
  creator: string;
  /** Alias is the name of the user domain alias to be transferred to the recipient. i.e. {alias}.snr */
  alias: string;
  /** Recipient is the wallet address of the recipient of the alias. */
  recipient: string;
  /** Amount is the amount of the alias to be transferred. */
  amount: number;
}

export interface MsgTransferAliasResponse {
  /** Success is true if the Alias was successfully transferred. */
  success: boolean;
  /** WhoIs is the updated WhoIs, contains the DID document and associated metadata. */
  who_is: WhoIs | undefined;
}

export interface MsgSellAlias {
  /** Creator is the wallet address of the creator of the transaction. */
  creator: string;
  /** Alias is the name of the app alias to be transferred to the recipient.  i.e. example.snr/{name} */
  alias: string;
  /** Amount is the amount of the alias to be transferred. */
  amount: number;
}

export interface MsgSellAliasResponse {
  /** Success is true if the Alias was successfully transferred. */
  success: boolean;
  /** WhoIs is the updated WhoIs, contains the DID document and associated metadata. */
  who_is: WhoIs | undefined;
}

const baseMsgCreateWhoIs: object = { creator: "", whois_type: 0 };

export const MsgCreateWhoIs = {
  encode(message: MsgCreateWhoIs, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did_document.length !== 0) {
      writer.uint32(18).bytes(message.did_document);
    }
    if (message.whois_type !== 0) {
      writer.uint32(24).int32(message.whois_type);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateWhoIs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateWhoIs } as MsgCreateWhoIs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did_document = reader.bytes();
          break;
        case 3:
          message.whois_type = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateWhoIs {
    const message = { ...baseMsgCreateWhoIs } as MsgCreateWhoIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.did_document !== undefined && object.did_document !== null) {
      message.did_document = bytesFromBase64(object.did_document);
    }
    if (object.whois_type !== undefined && object.whois_type !== null) {
      message.whois_type = whoIsTypeFromJSON(object.whois_type);
    } else {
      message.whois_type = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateWhoIs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did_document !== undefined &&
      (obj.did_document = base64FromBytes(
        message.did_document !== undefined
          ? message.did_document
          : new Uint8Array()
      ));
    message.whois_type !== undefined &&
      (obj.whois_type = whoIsTypeToJSON(message.whois_type));
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateWhoIs>): MsgCreateWhoIs {
    const message = { ...baseMsgCreateWhoIs } as MsgCreateWhoIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.did_document !== undefined && object.did_document !== null) {
      message.did_document = object.did_document;
    } else {
      message.did_document = new Uint8Array();
    }
    if (object.whois_type !== undefined && object.whois_type !== null) {
      message.whois_type = object.whois_type;
    } else {
      message.whois_type = 0;
    }
    return message;
  },
};

const baseMsgCreateWhoIsResponse: object = { success: false };

export const MsgCreateWhoIsResponse = {
  encode(
    message: MsgCreateWhoIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    if (message.who_is !== undefined) {
      WhoIs.encode(message.who_is, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateWhoIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateWhoIsResponse } as MsgCreateWhoIsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        case 2:
          message.who_is = WhoIs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateWhoIsResponse {
    const message = { ...baseMsgCreateWhoIsResponse } as MsgCreateWhoIsResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromJSON(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },

  toJSON(message: MsgCreateWhoIsResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateWhoIsResponse>
  ): MsgCreateWhoIsResponse {
    const message = { ...baseMsgCreateWhoIsResponse } as MsgCreateWhoIsResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromPartial(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },
};

const baseMsgUpdateWhoIs: object = { creator: "" };

export const MsgUpdateWhoIs = {
  encode(message: MsgUpdateWhoIs, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did_document.length !== 0) {
      writer.uint32(18).bytes(message.did_document);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateWhoIs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateWhoIs } as MsgUpdateWhoIs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did_document = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateWhoIs {
    const message = { ...baseMsgUpdateWhoIs } as MsgUpdateWhoIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.did_document !== undefined && object.did_document !== null) {
      message.did_document = bytesFromBase64(object.did_document);
    }
    return message;
  },

  toJSON(message: MsgUpdateWhoIs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did_document !== undefined &&
      (obj.did_document = base64FromBytes(
        message.did_document !== undefined
          ? message.did_document
          : new Uint8Array()
      ));
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateWhoIs>): MsgUpdateWhoIs {
    const message = { ...baseMsgUpdateWhoIs } as MsgUpdateWhoIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.did_document !== undefined && object.did_document !== null) {
      message.did_document = object.did_document;
    } else {
      message.did_document = new Uint8Array();
    }
    return message;
  },
};

const baseMsgUpdateWhoIsResponse: object = { success: false };

export const MsgUpdateWhoIsResponse = {
  encode(
    message: MsgUpdateWhoIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    if (message.who_is !== undefined) {
      WhoIs.encode(message.who_is, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateWhoIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateWhoIsResponse } as MsgUpdateWhoIsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        case 2:
          message.who_is = WhoIs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateWhoIsResponse {
    const message = { ...baseMsgUpdateWhoIsResponse } as MsgUpdateWhoIsResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromJSON(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },

  toJSON(message: MsgUpdateWhoIsResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateWhoIsResponse>
  ): MsgUpdateWhoIsResponse {
    const message = { ...baseMsgUpdateWhoIsResponse } as MsgUpdateWhoIsResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromPartial(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },
};

const baseMsgDeactivateWhoIs: object = { creator: "" };

export const MsgDeactivateWhoIs = {
  encode(
    message: MsgDeactivateWhoIs,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeactivateWhoIs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeactivateWhoIs } as MsgDeactivateWhoIs;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeactivateWhoIs {
    const message = { ...baseMsgDeactivateWhoIs } as MsgDeactivateWhoIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgDeactivateWhoIs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeactivateWhoIs>): MsgDeactivateWhoIs {
    const message = { ...baseMsgDeactivateWhoIs } as MsgDeactivateWhoIs;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgDeactivateWhoIsResponse: object = { success: false, did: "" };

export const MsgDeactivateWhoIsResponse = {
  encode(
    message: MsgDeactivateWhoIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeactivateWhoIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeactivateWhoIsResponse,
    } as MsgDeactivateWhoIsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
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

  fromJSON(object: any): MsgDeactivateWhoIsResponse {
    const message = {
      ...baseMsgDeactivateWhoIsResponse,
    } as MsgDeactivateWhoIsResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    return message;
  },

  toJSON(message: MsgDeactivateWhoIsResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeactivateWhoIsResponse>
  ): MsgDeactivateWhoIsResponse {
    const message = {
      ...baseMsgDeactivateWhoIsResponse,
    } as MsgDeactivateWhoIsResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    return message;
  },
};

const baseMsgBuyAlias: object = { creator: "", name: "" };

export const MsgBuyAlias = {
  encode(message: MsgBuyAlias, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBuyAlias {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgBuyAlias } as MsgBuyAlias;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBuyAlias {
    const message = { ...baseMsgBuyAlias } as MsgBuyAlias;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },

  toJSON(message: MsgBuyAlias): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgBuyAlias>): MsgBuyAlias {
    const message = { ...baseMsgBuyAlias } as MsgBuyAlias;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
};

const baseMsgBuyAliasResponse: object = { success: false };

export const MsgBuyAliasResponse = {
  encode(
    message: MsgBuyAliasResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    if (message.who_is !== undefined) {
      WhoIs.encode(message.who_is, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBuyAliasResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgBuyAliasResponse } as MsgBuyAliasResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        case 2:
          message.who_is = WhoIs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBuyAliasResponse {
    const message = { ...baseMsgBuyAliasResponse } as MsgBuyAliasResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromJSON(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },

  toJSON(message: MsgBuyAliasResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgBuyAliasResponse>): MsgBuyAliasResponse {
    const message = { ...baseMsgBuyAliasResponse } as MsgBuyAliasResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromPartial(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },
};

const baseMsgTransferAlias: object = {
  creator: "",
  alias: "",
  recipient: "",
  amount: 0,
};

export const MsgTransferAlias = {
  encode(message: MsgTransferAlias, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.alias !== "") {
      writer.uint32(18).string(message.alias);
    }
    if (message.recipient !== "") {
      writer.uint32(26).string(message.recipient);
    }
    if (message.amount !== 0) {
      writer.uint32(32).int32(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgTransferAlias {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgTransferAlias } as MsgTransferAlias;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.alias = reader.string();
          break;
        case 3:
          message.recipient = reader.string();
          break;
        case 4:
          message.amount = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgTransferAlias {
    const message = { ...baseMsgTransferAlias } as MsgTransferAlias;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = String(object.alias);
    } else {
      message.alias = "";
    }
    if (object.recipient !== undefined && object.recipient !== null) {
      message.recipient = String(object.recipient);
    } else {
      message.recipient = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Number(object.amount);
    } else {
      message.amount = 0;
    }
    return message;
  },

  toJSON(message: MsgTransferAlias): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.alias !== undefined && (obj.alias = message.alias);
    message.recipient !== undefined && (obj.recipient = message.recipient);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgTransferAlias>): MsgTransferAlias {
    const message = { ...baseMsgTransferAlias } as MsgTransferAlias;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = object.alias;
    } else {
      message.alias = "";
    }
    if (object.recipient !== undefined && object.recipient !== null) {
      message.recipient = object.recipient;
    } else {
      message.recipient = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = 0;
    }
    return message;
  },
};

const baseMsgTransferAliasResponse: object = { success: false };

export const MsgTransferAliasResponse = {
  encode(
    message: MsgTransferAliasResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    if (message.who_is !== undefined) {
      WhoIs.encode(message.who_is, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgTransferAliasResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgTransferAliasResponse,
    } as MsgTransferAliasResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        case 2:
          message.who_is = WhoIs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgTransferAliasResponse {
    const message = {
      ...baseMsgTransferAliasResponse,
    } as MsgTransferAliasResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromJSON(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },

  toJSON(message: MsgTransferAliasResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgTransferAliasResponse>
  ): MsgTransferAliasResponse {
    const message = {
      ...baseMsgTransferAliasResponse,
    } as MsgTransferAliasResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromPartial(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },
};

const baseMsgSellAlias: object = { creator: "", alias: "", amount: 0 };

export const MsgSellAlias = {
  encode(message: MsgSellAlias, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.alias !== "") {
      writer.uint32(18).string(message.alias);
    }
    if (message.amount !== 0) {
      writer.uint32(24).int32(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSellAlias {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSellAlias } as MsgSellAlias;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.alias = reader.string();
          break;
        case 3:
          message.amount = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSellAlias {
    const message = { ...baseMsgSellAlias } as MsgSellAlias;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = String(object.alias);
    } else {
      message.alias = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Number(object.amount);
    } else {
      message.amount = 0;
    }
    return message;
  },

  toJSON(message: MsgSellAlias): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.alias !== undefined && (obj.alias = message.alias);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSellAlias>): MsgSellAlias {
    const message = { ...baseMsgSellAlias } as MsgSellAlias;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.alias !== undefined && object.alias !== null) {
      message.alias = object.alias;
    } else {
      message.alias = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = 0;
    }
    return message;
  },
};

const baseMsgSellAliasResponse: object = { success: false };

export const MsgSellAliasResponse = {
  encode(
    message: MsgSellAliasResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    if (message.who_is !== undefined) {
      WhoIs.encode(message.who_is, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSellAliasResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSellAliasResponse } as MsgSellAliasResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        case 2:
          message.who_is = WhoIs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSellAliasResponse {
    const message = { ...baseMsgSellAliasResponse } as MsgSellAliasResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromJSON(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },

  toJSON(message: MsgSellAliasResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSellAliasResponse>): MsgSellAliasResponse {
    const message = { ...baseMsgSellAliasResponse } as MsgSellAliasResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromPartial(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateWhoIs(request: MsgCreateWhoIs): Promise<MsgCreateWhoIsResponse>;
  UpdateWhoIs(request: MsgUpdateWhoIs): Promise<MsgUpdateWhoIsResponse>;
  DeactivateWhoIs(
    request: MsgDeactivateWhoIs
  ): Promise<MsgDeactivateWhoIsResponse>;
  BuyAlias(request: MsgBuyAlias): Promise<MsgBuyAliasResponse>;
  SellAlias(request: MsgSellAlias): Promise<MsgSellAliasResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  TransferAlias(request: MsgTransferAlias): Promise<MsgTransferAliasResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateWhoIs(request: MsgCreateWhoIs): Promise<MsgCreateWhoIsResponse> {
    const data = MsgCreateWhoIs.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "CreateWhoIs",
      data
    );
    return promise.then((data) =>
      MsgCreateWhoIsResponse.decode(new Reader(data))
    );
  }

  UpdateWhoIs(request: MsgUpdateWhoIs): Promise<MsgUpdateWhoIsResponse> {
    const data = MsgUpdateWhoIs.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "UpdateWhoIs",
      data
    );
    return promise.then((data) =>
      MsgUpdateWhoIsResponse.decode(new Reader(data))
    );
  }

  DeactivateWhoIs(
    request: MsgDeactivateWhoIs
  ): Promise<MsgDeactivateWhoIsResponse> {
    const data = MsgDeactivateWhoIs.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "DeactivateWhoIs",
      data
    );
    return promise.then((data) =>
      MsgDeactivateWhoIsResponse.decode(new Reader(data))
    );
  }

  BuyAlias(request: MsgBuyAlias): Promise<MsgBuyAliasResponse> {
    const data = MsgBuyAlias.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "BuyAlias",
      data
    );
    return promise.then((data) => MsgBuyAliasResponse.decode(new Reader(data)));
  }

  SellAlias(request: MsgSellAlias): Promise<MsgSellAliasResponse> {
    const data = MsgSellAlias.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "SellAlias",
      data
    );
    return promise.then((data) =>
      MsgSellAliasResponse.decode(new Reader(data))
    );
  }

  TransferAlias(request: MsgTransferAlias): Promise<MsgTransferAliasResponse> {
    const data = MsgTransferAlias.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "TransferAlias",
      data
    );
    return promise.then((data) =>
      MsgTransferAliasResponse.decode(new Reader(data))
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
