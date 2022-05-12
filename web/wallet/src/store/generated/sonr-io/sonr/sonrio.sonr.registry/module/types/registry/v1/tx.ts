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
  did: string;
  /** WhoIs is the created WhoIs, contains the DID document and associated metadata. */
  who_is: WhoIs | undefined;
}

export interface MsgUpdateWhoIs {
  /** Creator is the wallet address of the creator of the transaction. */
  creator: string;
  /** Did is the top level DID of the WhoIs. */
  did: string;
  /** DidDocument is the DID document to be stored, in JSON format (see https://w3c-ccg.github.io/did-spec/#did-json-ld). */
  did_document: Uint8Array;
}

export interface MsgUpdateWhoIsResponse {
  /** Did is the top level DID of the WhoIs. */
  did: string;
  /** WhoIs is the created WhoIs, contains the DID document and associated metadata. */
  who_is: WhoIs | undefined;
}

export interface MsgDeactivateWhoIs {
  /** Creator is the wallet address of the creator of the transaction. */
  creator: string;
  /** Did is the top level DID of the WhoIs. */
  did: string;
}

export interface MsgDeactivateWhoIsResponse {
  /** Success is true if the WhoIs was successfully deactivated. */
  success: boolean;
  /** Did is the top level DID of the WhoIs. */
  did: string;
}

export interface MsgBuyNameAlias {
  /** Creator is the wallet address of the creator of the transaction. */
  creator: string;
  /** Did is the top level DID of the WhoIs. */
  did: string;
  /** Name is the name of the alias to be bought. i.e. {alias}.snr */
  name: string;
}

export interface MsgBuyNameAliasResponse {
  /** Did is the top level DID of the WhoIs. */
  did: string;
  /** WhoIs is the created WhoIs, contains the DID document and associated metadata. */
  who_is: WhoIs | undefined;
}

export interface MsgBuyAppAlias {
  /** Creator is the wallet address of the creator of the transaction. */
  creator: string;
  /** Did is the top level DID of the WhoIs. */
  did: string;
  /** Name is the name of the alias app extension to be bought. i.e. example.snr/{name} */
  name: string;
}

export interface MsgBuyAppAliasResponse {
  /** Did is the top level DID of the WhoIs. */
  did: string;
  /** WhoIs is the updated WhoIs, contains the DID document and associated metadata. */
  who_is: WhoIs | undefined;
}

export interface MsgTransferNameAlias {
  /** Creator is the wallet address of the creator of the transaction. */
  creator: string;
  /** Did is the top level DID of the WhoIs. */
  did: string;
  /** Alias is the name of the user domain alias to be transferred to the recipient. i.e. {alias}.snr */
  alias: string;
  /** Recipient is the wallet address of the recipient of the alias. */
  recipient: string;
  /** Amount is the amount of the alias to be transferred. */
  amount: number;
}

export interface MsgTransferNameAliasResponse {
  /** Success is true if the Alias was successfully transferred. */
  success: boolean;
  /** WhoIs is the updated WhoIs, contains the DID document and associated metadata. */
  who_is: WhoIs | undefined;
}

export interface MsgTransferAppAlias {
  /** Creator is the wallet address of the creator of the transaction. */
  creator: string;
  /** Did is the top level DID of the WhoIs. */
  did: string;
  /** Alias is the name of the app alias to be transferred to the recipient.  i.e. example.snr/{name} */
  alias: string;
  /** Recipient is the wallet address of the recipient of the alias. */
  recipient: string;
  /** Amount is the amount of the alias to be transferred. */
  amount: number;
}

export interface MsgTransferAppAliasResponse {
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

const baseMsgCreateWhoIsResponse: object = { did: "" };

export const MsgCreateWhoIsResponse = {
  encode(
    message: MsgCreateWhoIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
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
          message.did = reader.string();
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
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
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
    message.did !== undefined && (obj.did = message.did);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateWhoIsResponse>
  ): MsgCreateWhoIsResponse {
    const message = { ...baseMsgCreateWhoIsResponse } as MsgCreateWhoIsResponse;
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromPartial(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },
};

const baseMsgUpdateWhoIs: object = { creator: "", did: "" };

export const MsgUpdateWhoIs = {
  encode(message: MsgUpdateWhoIs, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.did_document.length !== 0) {
      writer.uint32(26).bytes(message.did_document);
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
          message.did = reader.string();
          break;
        case 3:
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
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.did_document !== undefined && object.did_document !== null) {
      message.did_document = bytesFromBase64(object.did_document);
    }
    return message;
  },

  toJSON(message: MsgUpdateWhoIs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
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
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.did_document !== undefined && object.did_document !== null) {
      message.did_document = object.did_document;
    } else {
      message.did_document = new Uint8Array();
    }
    return message;
  },
};

const baseMsgUpdateWhoIsResponse: object = { did: "" };

export const MsgUpdateWhoIsResponse = {
  encode(
    message: MsgUpdateWhoIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
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
          message.did = reader.string();
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
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
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
    message.did !== undefined && (obj.did = message.did);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateWhoIsResponse>
  ): MsgUpdateWhoIsResponse {
    const message = { ...baseMsgUpdateWhoIsResponse } as MsgUpdateWhoIsResponse;
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromPartial(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },
};

const baseMsgDeactivateWhoIs: object = { creator: "", did: "" };

export const MsgDeactivateWhoIs = {
  encode(
    message: MsgDeactivateWhoIs,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
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

  fromJSON(object: any): MsgDeactivateWhoIs {
    const message = { ...baseMsgDeactivateWhoIs } as MsgDeactivateWhoIs;
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

  toJSON(message: MsgDeactivateWhoIs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeactivateWhoIs>): MsgDeactivateWhoIs {
    const message = { ...baseMsgDeactivateWhoIs } as MsgDeactivateWhoIs;
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

const baseMsgBuyNameAlias: object = { creator: "", did: "", name: "" };

export const MsgBuyNameAlias = {
  encode(message: MsgBuyNameAlias, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBuyNameAlias {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgBuyNameAlias } as MsgBuyNameAlias;
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
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBuyNameAlias {
    const message = { ...baseMsgBuyNameAlias } as MsgBuyNameAlias;
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
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },

  toJSON(message: MsgBuyNameAlias): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgBuyNameAlias>): MsgBuyNameAlias {
    const message = { ...baseMsgBuyNameAlias } as MsgBuyNameAlias;
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
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
};

const baseMsgBuyNameAliasResponse: object = { did: "" };

export const MsgBuyNameAliasResponse = {
  encode(
    message: MsgBuyNameAliasResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    if (message.who_is !== undefined) {
      WhoIs.encode(message.who_is, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBuyNameAliasResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgBuyNameAliasResponse,
    } as MsgBuyNameAliasResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
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

  fromJSON(object: any): MsgBuyNameAliasResponse {
    const message = {
      ...baseMsgBuyNameAliasResponse,
    } as MsgBuyNameAliasResponse;
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromJSON(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },

  toJSON(message: MsgBuyNameAliasResponse): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgBuyNameAliasResponse>
  ): MsgBuyNameAliasResponse {
    const message = {
      ...baseMsgBuyNameAliasResponse,
    } as MsgBuyNameAliasResponse;
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromPartial(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },
};

const baseMsgBuyAppAlias: object = { creator: "", did: "", name: "" };

export const MsgBuyAppAlias = {
  encode(message: MsgBuyAppAlias, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBuyAppAlias {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgBuyAppAlias } as MsgBuyAppAlias;
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
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBuyAppAlias {
    const message = { ...baseMsgBuyAppAlias } as MsgBuyAppAlias;
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
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },

  toJSON(message: MsgBuyAppAlias): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgBuyAppAlias>): MsgBuyAppAlias {
    const message = { ...baseMsgBuyAppAlias } as MsgBuyAppAlias;
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
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
};

const baseMsgBuyAppAliasResponse: object = { did: "" };

export const MsgBuyAppAliasResponse = {
  encode(
    message: MsgBuyAppAliasResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    if (message.who_is !== undefined) {
      WhoIs.encode(message.who_is, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBuyAppAliasResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgBuyAppAliasResponse } as MsgBuyAppAliasResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
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

  fromJSON(object: any): MsgBuyAppAliasResponse {
    const message = { ...baseMsgBuyAppAliasResponse } as MsgBuyAppAliasResponse;
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromJSON(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },

  toJSON(message: MsgBuyAppAliasResponse): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgBuyAppAliasResponse>
  ): MsgBuyAppAliasResponse {
    const message = { ...baseMsgBuyAppAliasResponse } as MsgBuyAppAliasResponse;
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromPartial(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },
};

const baseMsgTransferNameAlias: object = {
  creator: "",
  did: "",
  alias: "",
  recipient: "",
  amount: 0,
};

export const MsgTransferNameAlias = {
  encode(
    message: MsgTransferNameAlias,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.alias !== "") {
      writer.uint32(26).string(message.alias);
    }
    if (message.recipient !== "") {
      writer.uint32(34).string(message.recipient);
    }
    if (message.amount !== 0) {
      writer.uint32(40).int32(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgTransferNameAlias {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgTransferNameAlias } as MsgTransferNameAlias;
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
          message.alias = reader.string();
          break;
        case 4:
          message.recipient = reader.string();
          break;
        case 5:
          message.amount = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgTransferNameAlias {
    const message = { ...baseMsgTransferNameAlias } as MsgTransferNameAlias;
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

  toJSON(message: MsgTransferNameAlias): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.alias !== undefined && (obj.alias = message.alias);
    message.recipient !== undefined && (obj.recipient = message.recipient);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgTransferNameAlias>): MsgTransferNameAlias {
    const message = { ...baseMsgTransferNameAlias } as MsgTransferNameAlias;
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

const baseMsgTransferNameAliasResponse: object = { success: false };

export const MsgTransferNameAliasResponse = {
  encode(
    message: MsgTransferNameAliasResponse,
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
  ): MsgTransferNameAliasResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgTransferNameAliasResponse,
    } as MsgTransferNameAliasResponse;
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

  fromJSON(object: any): MsgTransferNameAliasResponse {
    const message = {
      ...baseMsgTransferNameAliasResponse,
    } as MsgTransferNameAliasResponse;
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

  toJSON(message: MsgTransferNameAliasResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgTransferNameAliasResponse>
  ): MsgTransferNameAliasResponse {
    const message = {
      ...baseMsgTransferNameAliasResponse,
    } as MsgTransferNameAliasResponse;
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

const baseMsgTransferAppAlias: object = {
  creator: "",
  did: "",
  alias: "",
  recipient: "",
  amount: 0,
};

export const MsgTransferAppAlias = {
  encode(
    message: MsgTransferAppAlias,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.alias !== "") {
      writer.uint32(26).string(message.alias);
    }
    if (message.recipient !== "") {
      writer.uint32(34).string(message.recipient);
    }
    if (message.amount !== 0) {
      writer.uint32(40).int32(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgTransferAppAlias {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgTransferAppAlias } as MsgTransferAppAlias;
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
          message.alias = reader.string();
          break;
        case 4:
          message.recipient = reader.string();
          break;
        case 5:
          message.amount = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgTransferAppAlias {
    const message = { ...baseMsgTransferAppAlias } as MsgTransferAppAlias;
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

  toJSON(message: MsgTransferAppAlias): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.alias !== undefined && (obj.alias = message.alias);
    message.recipient !== undefined && (obj.recipient = message.recipient);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgTransferAppAlias>): MsgTransferAppAlias {
    const message = { ...baseMsgTransferAppAlias } as MsgTransferAppAlias;
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

const baseMsgTransferAppAliasResponse: object = { success: false };

export const MsgTransferAppAliasResponse = {
  encode(
    message: MsgTransferAppAliasResponse,
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
  ): MsgTransferAppAliasResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgTransferAppAliasResponse,
    } as MsgTransferAppAliasResponse;
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

  fromJSON(object: any): MsgTransferAppAliasResponse {
    const message = {
      ...baseMsgTransferAppAliasResponse,
    } as MsgTransferAppAliasResponse;
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

  toJSON(message: MsgTransferAppAliasResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgTransferAppAliasResponse>
  ): MsgTransferAppAliasResponse {
    const message = {
      ...baseMsgTransferAppAliasResponse,
    } as MsgTransferAppAliasResponse;
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
  DeleteWhoIs(request: MsgDeactivateWhoIs): Promise<MsgDeactivateWhoIsResponse>;
  BuyNameAlias(request: MsgBuyNameAlias): Promise<MsgBuyNameAliasResponse>;
  BuyAppAlias(request: MsgBuyAppAlias): Promise<MsgBuyAppAliasResponse>;
  TransferNameAlias(
    request: MsgTransferNameAlias
  ): Promise<MsgTransferNameAliasResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  TransferAppAlias(
    request: MsgTransferAppAlias
  ): Promise<MsgTransferAppAliasResponse>;
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

  DeleteWhoIs(
    request: MsgDeactivateWhoIs
  ): Promise<MsgDeactivateWhoIsResponse> {
    const data = MsgDeactivateWhoIs.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "DeleteWhoIs",
      data
    );
    return promise.then((data) =>
      MsgDeactivateWhoIsResponse.decode(new Reader(data))
    );
  }

  BuyNameAlias(request: MsgBuyNameAlias): Promise<MsgBuyNameAliasResponse> {
    const data = MsgBuyNameAlias.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "BuyNameAlias",
      data
    );
    return promise.then((data) =>
      MsgBuyNameAliasResponse.decode(new Reader(data))
    );
  }

  BuyAppAlias(request: MsgBuyAppAlias): Promise<MsgBuyAppAliasResponse> {
    const data = MsgBuyAppAlias.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "BuyAppAlias",
      data
    );
    return promise.then((data) =>
      MsgBuyAppAliasResponse.decode(new Reader(data))
    );
  }

  TransferNameAlias(
    request: MsgTransferNameAlias
  ): Promise<MsgTransferNameAliasResponse> {
    const data = MsgTransferNameAlias.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "TransferNameAlias",
      data
    );
    return promise.then((data) =>
      MsgTransferNameAliasResponse.decode(new Reader(data))
    );
  }

  TransferAppAlias(
    request: MsgTransferAppAlias
  ): Promise<MsgTransferAppAliasResponse> {
    const data = MsgTransferAppAlias.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "TransferAppAlias",
      data
    );
    return promise.then((data) =>
      MsgTransferAppAliasResponse.decode(new Reader(data))
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
