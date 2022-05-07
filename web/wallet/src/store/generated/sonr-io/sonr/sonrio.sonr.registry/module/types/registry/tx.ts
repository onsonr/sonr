/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Credential } from "../registry/credential";
import { WhoIs, Session } from "../registry/who_is";

export const protobufPackage = "sonrio.sonr.registry";

export interface MsgRegisterApplication {
  /** Creator is the account address of the creator of the Application. */
  creator: string;
  /** Client side JSON Web Token for AssertionMethod */
  credential: Credential | undefined;
  /** Application Name is the endpoint of the Application. */
  application_name: string;
  /** Application Description is the description of the Application. */
  application_description: string;
  /** Application URL is the URL of the Application. */
  application_url: string;
  /** Application Category is the category of the Application. */
  application_category: string;
}

export interface MsgRegisterApplicationResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** WhoIs for the registered name */
  who_is: WhoIs | undefined;
  /** Session returns the session for the name */
  session: Session | undefined;
}

/** MsgRegisterName is a request to register a name with the ".snr" name of a peer */
export interface MsgRegisterName {
  /** Account address of the name owner */
  creator: string;
  /** Selected Name to register */
  name_to_register: string;
  /** Client side JSON Web Token for AssertionMethod */
  credential: Credential | undefined;
  /** The Updated Metadata */
  metadata: { [key: string]: string };
}

export interface MsgRegisterName_MetadataEntry {
  key: string;
  value: string;
}

export interface MsgRegisterNameResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** WhoIs for the registered name */
  who_is: WhoIs | undefined;
  /** Session returns the session for the name */
  session: Session | undefined;
}

/** MsgAccessName defines the MsgAccessName transaction. */
export interface MsgAccessName {
  /** The account that is accessing the name */
  creator: string;
  /** The name to access */
  name: string;
  /** Client side JSON Web Token for AssertionMethod */
  credential: Credential | undefined;
}

export interface MsgAccessNameResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** WhoIs for the registered name */
  who_is: WhoIs | undefined;
  /** Session returns the session for the name */
  session: Session | undefined;
}

export interface MsgUpdateName {
  /** The account that owns the name. */
  creator: string;
  /** The did of the peer to update the name of */
  did: string;
  /** Client side JSON Web Token for AssertionMethod. For additional devices being linked. */
  credential: Credential | undefined;
  /** The Updated Metadata */
  metadata: { [key: string]: string };
  /** Session returns the session for the name */
  session: Session | undefined;
}

export interface MsgUpdateName_MetadataEntry {
  key: string;
  value: string;
}

export interface MsgUpdateNameResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** WhoIs for the registered name */
  who_is: WhoIs | undefined;
}

export interface MsgAccessApplication {
  /** The account that is accessing the Application */
  creator: string;
  /** The name of the Application to access */
  app_name: string;
  /** Client side JSON Web Token for AssertionMethod */
  credential: Credential | undefined;
}

export interface MsgAccessApplicationResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** Data of the response */
  metadata: { [key: string]: string };
  /** WhoIs for the registered name */
  who_is: WhoIs | undefined;
  /** Session returns the session for the name */
  session: Session | undefined;
}

export interface MsgAccessApplicationResponse_MetadataEntry {
  key: string;
  value: string;
}

export interface MsgUpdateApplication {
  /** The account that owns the name. */
  creator: string;
  /** The name of the peer to update the Application details of */
  did: string;
  /** The updated configuration for the Application */
  metadata: { [key: string]: string };
  /** Session returns the session for the name */
  session: Session | undefined;
}

export interface MsgUpdateApplication_MetadataEntry {
  key: string;
  value: string;
}

export interface MsgUpdateApplicationResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** Data of the response */
  metadata: { [key: string]: string };
  /** WhoIs for the registered name */
  who_is: WhoIs | undefined;
}

export interface MsgUpdateApplicationResponse_MetadataEntry {
  key: string;
  value: string;
}

export interface MsgCreateWhoIs {
  creator: string;
  did: string;
  document: Uint8Array;
  credentials: Credential[];
  name: string;
}

export interface MsgCreateWhoIsResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  who_is: WhoIs | undefined;
}

export interface MsgUpdateWhoIs {
  creator: string;
  did: string;
  document: Uint8Array;
  credentials: Credential[];
}

export interface MsgUpdateWhoIsResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  who_is: WhoIs | undefined;
}

export interface MsgDeleteWhoIs {
  creator: string;
  did: string;
}

export interface MsgDeleteWhoIsResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
}

const baseMsgRegisterApplication: object = {
  creator: "",
  application_name: "",
  application_description: "",
  application_url: "",
  application_category: "",
};

export const MsgRegisterApplication = {
  encode(
    message: MsgRegisterApplication,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.credential !== undefined) {
      Credential.encode(message.credential, writer.uint32(18).fork()).ldelim();
    }
    if (message.application_name !== "") {
      writer.uint32(26).string(message.application_name);
    }
    if (message.application_description !== "") {
      writer.uint32(34).string(message.application_description);
    }
    if (message.application_url !== "") {
      writer.uint32(42).string(message.application_url);
    }
    if (message.application_category !== "") {
      writer.uint32(50).string(message.application_category);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRegisterApplication {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRegisterApplication } as MsgRegisterApplication;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.credential = Credential.decode(reader, reader.uint32());
          break;
        case 3:
          message.application_name = reader.string();
          break;
        case 4:
          message.application_description = reader.string();
          break;
        case 5:
          message.application_url = reader.string();
          break;
        case 6:
          message.application_category = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterApplication {
    const message = { ...baseMsgRegisterApplication } as MsgRegisterApplication;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.credential !== undefined && object.credential !== null) {
      message.credential = Credential.fromJSON(object.credential);
    } else {
      message.credential = undefined;
    }
    if (
      object.application_name !== undefined &&
      object.application_name !== null
    ) {
      message.application_name = String(object.application_name);
    } else {
      message.application_name = "";
    }
    if (
      object.application_description !== undefined &&
      object.application_description !== null
    ) {
      message.application_description = String(object.application_description);
    } else {
      message.application_description = "";
    }
    if (
      object.application_url !== undefined &&
      object.application_url !== null
    ) {
      message.application_url = String(object.application_url);
    } else {
      message.application_url = "";
    }
    if (
      object.application_category !== undefined &&
      object.application_category !== null
    ) {
      message.application_category = String(object.application_category);
    } else {
      message.application_category = "";
    }
    return message;
  },

  toJSON(message: MsgRegisterApplication): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.credential !== undefined &&
      (obj.credential = message.credential
        ? Credential.toJSON(message.credential)
        : undefined);
    message.application_name !== undefined &&
      (obj.application_name = message.application_name);
    message.application_description !== undefined &&
      (obj.application_description = message.application_description);
    message.application_url !== undefined &&
      (obj.application_url = message.application_url);
    message.application_category !== undefined &&
      (obj.application_category = message.application_category);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRegisterApplication>
  ): MsgRegisterApplication {
    const message = { ...baseMsgRegisterApplication } as MsgRegisterApplication;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.credential !== undefined && object.credential !== null) {
      message.credential = Credential.fromPartial(object.credential);
    } else {
      message.credential = undefined;
    }
    if (
      object.application_name !== undefined &&
      object.application_name !== null
    ) {
      message.application_name = object.application_name;
    } else {
      message.application_name = "";
    }
    if (
      object.application_description !== undefined &&
      object.application_description !== null
    ) {
      message.application_description = object.application_description;
    } else {
      message.application_description = "";
    }
    if (
      object.application_url !== undefined &&
      object.application_url !== null
    ) {
      message.application_url = object.application_url;
    } else {
      message.application_url = "";
    }
    if (
      object.application_category !== undefined &&
      object.application_category !== null
    ) {
      message.application_category = object.application_category;
    } else {
      message.application_category = "";
    }
    return message;
  },
};

const baseMsgRegisterApplicationResponse: object = { code: 0, message: "" };

export const MsgRegisterApplicationResponse = {
  encode(
    message: MsgRegisterApplicationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.who_is !== undefined) {
      WhoIs.encode(message.who_is, writer.uint32(26).fork()).ldelim();
    }
    if (message.session !== undefined) {
      Session.encode(message.session, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRegisterApplicationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRegisterApplicationResponse,
    } as MsgRegisterApplicationResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.who_is = WhoIs.decode(reader, reader.uint32());
          break;
        case 4:
          message.session = Session.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterApplicationResponse {
    const message = {
      ...baseMsgRegisterApplicationResponse,
    } as MsgRegisterApplicationResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = Number(object.code);
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromJSON(object.who_is);
    } else {
      message.who_is = undefined;
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromJSON(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },

  toJSON(message: MsgRegisterApplicationResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.message !== undefined && (obj.message = message.message);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    message.session !== undefined &&
      (obj.session = message.session
        ? Session.toJSON(message.session)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRegisterApplicationResponse>
  ): MsgRegisterApplicationResponse {
    const message = {
      ...baseMsgRegisterApplicationResponse,
    } as MsgRegisterApplicationResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromPartial(object.who_is);
    } else {
      message.who_is = undefined;
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromPartial(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },
};

const baseMsgRegisterName: object = { creator: "", name_to_register: "" };

export const MsgRegisterName = {
  encode(message: MsgRegisterName, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name_to_register !== "") {
      writer.uint32(18).string(message.name_to_register);
    }
    if (message.credential !== undefined) {
      Credential.encode(message.credential, writer.uint32(26).fork()).ldelim();
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      MsgRegisterName_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRegisterName {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRegisterName } as MsgRegisterName;
    message.metadata = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name_to_register = reader.string();
          break;
        case 3:
          message.credential = Credential.decode(reader, reader.uint32());
          break;
        case 4:
          const entry4 = MsgRegisterName_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry4.value !== undefined) {
            message.metadata[entry4.key] = entry4.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterName {
    const message = { ...baseMsgRegisterName } as MsgRegisterName;
    message.metadata = {};
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (
      object.name_to_register !== undefined &&
      object.name_to_register !== null
    ) {
      message.name_to_register = String(object.name_to_register);
    } else {
      message.name_to_register = "";
    }
    if (object.credential !== undefined && object.credential !== null) {
      message.credential = Credential.fromJSON(object.credential);
    } else {
      message.credential = undefined;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    return message;
  },

  toJSON(message: MsgRegisterName): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name_to_register !== undefined &&
      (obj.name_to_register = message.name_to_register);
    message.credential !== undefined &&
      (obj.credential = message.credential
        ? Credential.toJSON(message.credential)
        : undefined);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRegisterName>): MsgRegisterName {
    const message = { ...baseMsgRegisterName } as MsgRegisterName;
    message.metadata = {};
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (
      object.name_to_register !== undefined &&
      object.name_to_register !== null
    ) {
      message.name_to_register = object.name_to_register;
    } else {
      message.name_to_register = "";
    }
    if (object.credential !== undefined && object.credential !== null) {
      message.credential = Credential.fromPartial(object.credential);
    } else {
      message.credential = undefined;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
    }
    return message;
  },
};

const baseMsgRegisterName_MetadataEntry: object = { key: "", value: "" };

export const MsgRegisterName_MetadataEntry = {
  encode(
    message: MsgRegisterName_MetadataEntry,
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

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRegisterName_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRegisterName_MetadataEntry,
    } as MsgRegisterName_MetadataEntry;
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

  fromJSON(object: any): MsgRegisterName_MetadataEntry {
    const message = {
      ...baseMsgRegisterName_MetadataEntry,
    } as MsgRegisterName_MetadataEntry;
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

  toJSON(message: MsgRegisterName_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRegisterName_MetadataEntry>
  ): MsgRegisterName_MetadataEntry {
    const message = {
      ...baseMsgRegisterName_MetadataEntry,
    } as MsgRegisterName_MetadataEntry;
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

const baseMsgRegisterNameResponse: object = { code: 0, message: "" };

export const MsgRegisterNameResponse = {
  encode(
    message: MsgRegisterNameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.who_is !== undefined) {
      WhoIs.encode(message.who_is, writer.uint32(26).fork()).ldelim();
    }
    if (message.session !== undefined) {
      Session.encode(message.session, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRegisterNameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRegisterNameResponse,
    } as MsgRegisterNameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.who_is = WhoIs.decode(reader, reader.uint32());
          break;
        case 4:
          message.session = Session.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterNameResponse {
    const message = {
      ...baseMsgRegisterNameResponse,
    } as MsgRegisterNameResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = Number(object.code);
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromJSON(object.who_is);
    } else {
      message.who_is = undefined;
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromJSON(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },

  toJSON(message: MsgRegisterNameResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.message !== undefined && (obj.message = message.message);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    message.session !== undefined &&
      (obj.session = message.session
        ? Session.toJSON(message.session)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRegisterNameResponse>
  ): MsgRegisterNameResponse {
    const message = {
      ...baseMsgRegisterNameResponse,
    } as MsgRegisterNameResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromPartial(object.who_is);
    } else {
      message.who_is = undefined;
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromPartial(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },
};

const baseMsgAccessName: object = { creator: "", name: "" };

export const MsgAccessName = {
  encode(message: MsgAccessName, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.credential !== undefined) {
      Credential.encode(message.credential, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAccessName {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAccessName } as MsgAccessName;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
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

  fromJSON(object: any): MsgAccessName {
    const message = { ...baseMsgAccessName } as MsgAccessName;
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
    if (object.credential !== undefined && object.credential !== null) {
      message.credential = Credential.fromJSON(object.credential);
    } else {
      message.credential = undefined;
    }
    return message;
  },

  toJSON(message: MsgAccessName): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.credential !== undefined &&
      (obj.credential = message.credential
        ? Credential.toJSON(message.credential)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAccessName>): MsgAccessName {
    const message = { ...baseMsgAccessName } as MsgAccessName;
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
    if (object.credential !== undefined && object.credential !== null) {
      message.credential = Credential.fromPartial(object.credential);
    } else {
      message.credential = undefined;
    }
    return message;
  },
};

const baseMsgAccessNameResponse: object = { code: 0, message: "" };

export const MsgAccessNameResponse = {
  encode(
    message: MsgAccessNameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.who_is !== undefined) {
      WhoIs.encode(message.who_is, writer.uint32(26).fork()).ldelim();
    }
    if (message.session !== undefined) {
      Session.encode(message.session, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAccessNameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAccessNameResponse } as MsgAccessNameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.who_is = WhoIs.decode(reader, reader.uint32());
          break;
        case 4:
          message.session = Session.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAccessNameResponse {
    const message = { ...baseMsgAccessNameResponse } as MsgAccessNameResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = Number(object.code);
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromJSON(object.who_is);
    } else {
      message.who_is = undefined;
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromJSON(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },

  toJSON(message: MsgAccessNameResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.message !== undefined && (obj.message = message.message);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    message.session !== undefined &&
      (obj.session = message.session
        ? Session.toJSON(message.session)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgAccessNameResponse>
  ): MsgAccessNameResponse {
    const message = { ...baseMsgAccessNameResponse } as MsgAccessNameResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromPartial(object.who_is);
    } else {
      message.who_is = undefined;
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromPartial(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },
};

const baseMsgUpdateName: object = { creator: "", did: "" };

export const MsgUpdateName = {
  encode(message: MsgUpdateName, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.credential !== undefined) {
      Credential.encode(message.credential, writer.uint32(26).fork()).ldelim();
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      MsgUpdateName_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    if (message.session !== undefined) {
      Session.encode(message.session, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateName {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateName } as MsgUpdateName;
    message.metadata = {};
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
          message.credential = Credential.decode(reader, reader.uint32());
          break;
        case 4:
          const entry4 = MsgUpdateName_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry4.value !== undefined) {
            message.metadata[entry4.key] = entry4.value;
          }
          break;
        case 5:
          message.session = Session.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateName {
    const message = { ...baseMsgUpdateName } as MsgUpdateName;
    message.metadata = {};
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
    if (object.credential !== undefined && object.credential !== null) {
      message.credential = Credential.fromJSON(object.credential);
    } else {
      message.credential = undefined;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromJSON(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },

  toJSON(message: MsgUpdateName): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.credential !== undefined &&
      (obj.credential = message.credential
        ? Credential.toJSON(message.credential)
        : undefined);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    message.session !== undefined &&
      (obj.session = message.session
        ? Session.toJSON(message.session)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateName>): MsgUpdateName {
    const message = { ...baseMsgUpdateName } as MsgUpdateName;
    message.metadata = {};
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
    if (object.credential !== undefined && object.credential !== null) {
      message.credential = Credential.fromPartial(object.credential);
    } else {
      message.credential = undefined;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromPartial(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },
};

const baseMsgUpdateName_MetadataEntry: object = { key: "", value: "" };

export const MsgUpdateName_MetadataEntry = {
  encode(
    message: MsgUpdateName_MetadataEntry,
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

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateName_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateName_MetadataEntry,
    } as MsgUpdateName_MetadataEntry;
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

  fromJSON(object: any): MsgUpdateName_MetadataEntry {
    const message = {
      ...baseMsgUpdateName_MetadataEntry,
    } as MsgUpdateName_MetadataEntry;
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

  toJSON(message: MsgUpdateName_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateName_MetadataEntry>
  ): MsgUpdateName_MetadataEntry {
    const message = {
      ...baseMsgUpdateName_MetadataEntry,
    } as MsgUpdateName_MetadataEntry;
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

const baseMsgUpdateNameResponse: object = { code: 0, message: "" };

export const MsgUpdateNameResponse = {
  encode(
    message: MsgUpdateNameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.who_is !== undefined) {
      WhoIs.encode(message.who_is, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateNameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateNameResponse } as MsgUpdateNameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.who_is = WhoIs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateNameResponse {
    const message = { ...baseMsgUpdateNameResponse } as MsgUpdateNameResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = Number(object.code);
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromJSON(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },

  toJSON(message: MsgUpdateNameResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.message !== undefined && (obj.message = message.message);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateNameResponse>
  ): MsgUpdateNameResponse {
    const message = { ...baseMsgUpdateNameResponse } as MsgUpdateNameResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromPartial(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },
};

const baseMsgAccessApplication: object = { creator: "", app_name: "" };

export const MsgAccessApplication = {
  encode(
    message: MsgAccessApplication,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.app_name !== "") {
      writer.uint32(18).string(message.app_name);
    }
    if (message.credential !== undefined) {
      Credential.encode(message.credential, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAccessApplication {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAccessApplication } as MsgAccessApplication;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.app_name = reader.string();
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

  fromJSON(object: any): MsgAccessApplication {
    const message = { ...baseMsgAccessApplication } as MsgAccessApplication;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.app_name !== undefined && object.app_name !== null) {
      message.app_name = String(object.app_name);
    } else {
      message.app_name = "";
    }
    if (object.credential !== undefined && object.credential !== null) {
      message.credential = Credential.fromJSON(object.credential);
    } else {
      message.credential = undefined;
    }
    return message;
  },

  toJSON(message: MsgAccessApplication): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.app_name !== undefined && (obj.app_name = message.app_name);
    message.credential !== undefined &&
      (obj.credential = message.credential
        ? Credential.toJSON(message.credential)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAccessApplication>): MsgAccessApplication {
    const message = { ...baseMsgAccessApplication } as MsgAccessApplication;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.app_name !== undefined && object.app_name !== null) {
      message.app_name = object.app_name;
    } else {
      message.app_name = "";
    }
    if (object.credential !== undefined && object.credential !== null) {
      message.credential = Credential.fromPartial(object.credential);
    } else {
      message.credential = undefined;
    }
    return message;
  },
};

const baseMsgAccessApplicationResponse: object = { code: 0, message: "" };

export const MsgAccessApplicationResponse = {
  encode(
    message: MsgAccessApplicationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      MsgAccessApplicationResponse_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(26).fork()
      ).ldelim();
    });
    if (message.who_is !== undefined) {
      WhoIs.encode(message.who_is, writer.uint32(34).fork()).ldelim();
    }
    if (message.session !== undefined) {
      Session.encode(message.session, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAccessApplicationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAccessApplicationResponse,
    } as MsgAccessApplicationResponse;
    message.metadata = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          const entry3 = MsgAccessApplicationResponse_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry3.value !== undefined) {
            message.metadata[entry3.key] = entry3.value;
          }
          break;
        case 4:
          message.who_is = WhoIs.decode(reader, reader.uint32());
          break;
        case 5:
          message.session = Session.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAccessApplicationResponse {
    const message = {
      ...baseMsgAccessApplicationResponse,
    } as MsgAccessApplicationResponse;
    message.metadata = {};
    if (object.code !== undefined && object.code !== null) {
      message.code = Number(object.code);
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromJSON(object.who_is);
    } else {
      message.who_is = undefined;
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromJSON(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },

  toJSON(message: MsgAccessApplicationResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.message !== undefined && (obj.message = message.message);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    message.session !== undefined &&
      (obj.session = message.session
        ? Session.toJSON(message.session)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgAccessApplicationResponse>
  ): MsgAccessApplicationResponse {
    const message = {
      ...baseMsgAccessApplicationResponse,
    } as MsgAccessApplicationResponse;
    message.metadata = {};
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromPartial(object.who_is);
    } else {
      message.who_is = undefined;
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromPartial(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },
};

const baseMsgAccessApplicationResponse_MetadataEntry: object = {
  key: "",
  value: "",
};

export const MsgAccessApplicationResponse_MetadataEntry = {
  encode(
    message: MsgAccessApplicationResponse_MetadataEntry,
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

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAccessApplicationResponse_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAccessApplicationResponse_MetadataEntry,
    } as MsgAccessApplicationResponse_MetadataEntry;
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

  fromJSON(object: any): MsgAccessApplicationResponse_MetadataEntry {
    const message = {
      ...baseMsgAccessApplicationResponse_MetadataEntry,
    } as MsgAccessApplicationResponse_MetadataEntry;
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

  toJSON(message: MsgAccessApplicationResponse_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgAccessApplicationResponse_MetadataEntry>
  ): MsgAccessApplicationResponse_MetadataEntry {
    const message = {
      ...baseMsgAccessApplicationResponse_MetadataEntry,
    } as MsgAccessApplicationResponse_MetadataEntry;
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

const baseMsgUpdateApplication: object = { creator: "", did: "" };

export const MsgUpdateApplication = {
  encode(
    message: MsgUpdateApplication,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      MsgUpdateApplication_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(26).fork()
      ).ldelim();
    });
    if (message.session !== undefined) {
      Session.encode(message.session, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateApplication {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateApplication } as MsgUpdateApplication;
    message.metadata = {};
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
          const entry3 = MsgUpdateApplication_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry3.value !== undefined) {
            message.metadata[entry3.key] = entry3.value;
          }
          break;
        case 4:
          message.session = Session.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateApplication {
    const message = { ...baseMsgUpdateApplication } as MsgUpdateApplication;
    message.metadata = {};
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
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromJSON(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },

  toJSON(message: MsgUpdateApplication): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    message.session !== undefined &&
      (obj.session = message.session
        ? Session.toJSON(message.session)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateApplication>): MsgUpdateApplication {
    const message = { ...baseMsgUpdateApplication } as MsgUpdateApplication;
    message.metadata = {};
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
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
    }
    if (object.session !== undefined && object.session !== null) {
      message.session = Session.fromPartial(object.session);
    } else {
      message.session = undefined;
    }
    return message;
  },
};

const baseMsgUpdateApplication_MetadataEntry: object = { key: "", value: "" };

export const MsgUpdateApplication_MetadataEntry = {
  encode(
    message: MsgUpdateApplication_MetadataEntry,
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

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateApplication_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateApplication_MetadataEntry,
    } as MsgUpdateApplication_MetadataEntry;
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

  fromJSON(object: any): MsgUpdateApplication_MetadataEntry {
    const message = {
      ...baseMsgUpdateApplication_MetadataEntry,
    } as MsgUpdateApplication_MetadataEntry;
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

  toJSON(message: MsgUpdateApplication_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateApplication_MetadataEntry>
  ): MsgUpdateApplication_MetadataEntry {
    const message = {
      ...baseMsgUpdateApplication_MetadataEntry,
    } as MsgUpdateApplication_MetadataEntry;
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

const baseMsgUpdateApplicationResponse: object = { code: 0, message: "" };

export const MsgUpdateApplicationResponse = {
  encode(
    message: MsgUpdateApplicationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      MsgUpdateApplicationResponse_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(26).fork()
      ).ldelim();
    });
    if (message.who_is !== undefined) {
      WhoIs.encode(message.who_is, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateApplicationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateApplicationResponse,
    } as MsgUpdateApplicationResponse;
    message.metadata = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          const entry3 = MsgUpdateApplicationResponse_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry3.value !== undefined) {
            message.metadata[entry3.key] = entry3.value;
          }
          break;
        case 4:
          message.who_is = WhoIs.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateApplicationResponse {
    const message = {
      ...baseMsgUpdateApplicationResponse,
    } as MsgUpdateApplicationResponse;
    message.metadata = {};
    if (object.code !== undefined && object.code !== null) {
      message.code = Number(object.code);
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromJSON(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },

  toJSON(message: MsgUpdateApplicationResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.message !== undefined && (obj.message = message.message);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateApplicationResponse>
  ): MsgUpdateApplicationResponse {
    const message = {
      ...baseMsgUpdateApplicationResponse,
    } as MsgUpdateApplicationResponse;
    message.metadata = {};
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromPartial(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },
};

const baseMsgUpdateApplicationResponse_MetadataEntry: object = {
  key: "",
  value: "",
};

export const MsgUpdateApplicationResponse_MetadataEntry = {
  encode(
    message: MsgUpdateApplicationResponse_MetadataEntry,
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

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateApplicationResponse_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateApplicationResponse_MetadataEntry,
    } as MsgUpdateApplicationResponse_MetadataEntry;
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

  fromJSON(object: any): MsgUpdateApplicationResponse_MetadataEntry {
    const message = {
      ...baseMsgUpdateApplicationResponse_MetadataEntry,
    } as MsgUpdateApplicationResponse_MetadataEntry;
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

  toJSON(message: MsgUpdateApplicationResponse_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateApplicationResponse_MetadataEntry>
  ): MsgUpdateApplicationResponse_MetadataEntry {
    const message = {
      ...baseMsgUpdateApplicationResponse_MetadataEntry,
    } as MsgUpdateApplicationResponse_MetadataEntry;
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

const baseMsgCreateWhoIs: object = { creator: "", did: "", name: "" };

export const MsgCreateWhoIs = {
  encode(message: MsgCreateWhoIs, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.document.length !== 0) {
      writer.uint32(26).bytes(message.document);
    }
    for (const v of message.credentials) {
      Credential.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.name !== "") {
      writer.uint32(42).string(message.name);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateWhoIs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateWhoIs } as MsgCreateWhoIs;
    message.credentials = [];
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
          message.document = reader.bytes();
          break;
        case 4:
          message.credentials.push(Credential.decode(reader, reader.uint32()));
          break;
        case 5:
          message.name = reader.string();
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
    message.credentials = [];
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
    if (object.document !== undefined && object.document !== null) {
      message.document = bytesFromBase64(object.document);
    }
    if (object.credentials !== undefined && object.credentials !== null) {
      for (const e of object.credentials) {
        message.credentials.push(Credential.fromJSON(e));
      }
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    return message;
  },

  toJSON(message: MsgCreateWhoIs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.document !== undefined &&
      (obj.document = base64FromBytes(
        message.document !== undefined ? message.document : new Uint8Array()
      ));
    if (message.credentials) {
      obj.credentials = message.credentials.map((e) =>
        e ? Credential.toJSON(e) : undefined
      );
    } else {
      obj.credentials = [];
    }
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateWhoIs>): MsgCreateWhoIs {
    const message = { ...baseMsgCreateWhoIs } as MsgCreateWhoIs;
    message.credentials = [];
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
    if (object.document !== undefined && object.document !== null) {
      message.document = object.document;
    } else {
      message.document = new Uint8Array();
    }
    if (object.credentials !== undefined && object.credentials !== null) {
      for (const e of object.credentials) {
        message.credentials.push(Credential.fromPartial(e));
      }
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    return message;
  },
};

const baseMsgCreateWhoIsResponse: object = { code: 0, message: "" };

export const MsgCreateWhoIsResponse = {
  encode(
    message: MsgCreateWhoIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.who_is !== undefined) {
      WhoIs.encode(message.who_is, writer.uint32(26).fork()).ldelim();
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
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
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
    if (object.code !== undefined && object.code !== null) {
      message.code = Number(object.code);
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
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
    message.code !== undefined && (obj.code = message.code);
    message.message !== undefined && (obj.message = message.message);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateWhoIsResponse>
  ): MsgCreateWhoIsResponse {
    const message = { ...baseMsgCreateWhoIsResponse } as MsgCreateWhoIsResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
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
    if (message.document.length !== 0) {
      writer.uint32(26).bytes(message.document);
    }
    for (const v of message.credentials) {
      Credential.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateWhoIs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateWhoIs } as MsgUpdateWhoIs;
    message.credentials = [];
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
          message.document = reader.bytes();
          break;
        case 4:
          message.credentials.push(Credential.decode(reader, reader.uint32()));
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
    message.credentials = [];
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
    if (object.document !== undefined && object.document !== null) {
      message.document = bytesFromBase64(object.document);
    }
    if (object.credentials !== undefined && object.credentials !== null) {
      for (const e of object.credentials) {
        message.credentials.push(Credential.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: MsgUpdateWhoIs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.document !== undefined &&
      (obj.document = base64FromBytes(
        message.document !== undefined ? message.document : new Uint8Array()
      ));
    if (message.credentials) {
      obj.credentials = message.credentials.map((e) =>
        e ? Credential.toJSON(e) : undefined
      );
    } else {
      obj.credentials = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateWhoIs>): MsgUpdateWhoIs {
    const message = { ...baseMsgUpdateWhoIs } as MsgUpdateWhoIs;
    message.credentials = [];
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
    if (object.document !== undefined && object.document !== null) {
      message.document = object.document;
    } else {
      message.document = new Uint8Array();
    }
    if (object.credentials !== undefined && object.credentials !== null) {
      for (const e of object.credentials) {
        message.credentials.push(Credential.fromPartial(e));
      }
    }
    return message;
  },
};

const baseMsgUpdateWhoIsResponse: object = { code: 0, message: "" };

export const MsgUpdateWhoIsResponse = {
  encode(
    message: MsgUpdateWhoIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.who_is !== undefined) {
      WhoIs.encode(message.who_is, writer.uint32(26).fork()).ldelim();
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
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
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
    if (object.code !== undefined && object.code !== null) {
      message.code = Number(object.code);
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
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
    message.code !== undefined && (obj.code = message.code);
    message.message !== undefined && (obj.message = message.message);
    message.who_is !== undefined &&
      (obj.who_is = message.who_is ? WhoIs.toJSON(message.who_is) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateWhoIsResponse>
  ): MsgUpdateWhoIsResponse {
    const message = { ...baseMsgUpdateWhoIsResponse } as MsgUpdateWhoIsResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.who_is !== undefined && object.who_is !== null) {
      message.who_is = WhoIs.fromPartial(object.who_is);
    } else {
      message.who_is = undefined;
    }
    return message;
  },
};

const baseMsgDeleteWhoIs: object = { creator: "", did: "" };

export const MsgDeleteWhoIs = {
  encode(message: MsgDeleteWhoIs, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteWhoIs {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteWhoIs } as MsgDeleteWhoIs;
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

  fromJSON(object: any): MsgDeleteWhoIs {
    const message = { ...baseMsgDeleteWhoIs } as MsgDeleteWhoIs;
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

  toJSON(message: MsgDeleteWhoIs): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteWhoIs>): MsgDeleteWhoIs {
    const message = { ...baseMsgDeleteWhoIs } as MsgDeleteWhoIs;
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

const baseMsgDeleteWhoIsResponse: object = { code: 0, message: "" };

export const MsgDeleteWhoIsResponse = {
  encode(
    message: MsgDeleteWhoIsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteWhoIsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteWhoIsResponse } as MsgDeleteWhoIsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteWhoIsResponse {
    const message = { ...baseMsgDeleteWhoIsResponse } as MsgDeleteWhoIsResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = Number(object.code);
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteWhoIsResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.message !== undefined && (obj.message = message.message);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteWhoIsResponse>
  ): MsgDeleteWhoIsResponse {
    const message = { ...baseMsgDeleteWhoIsResponse } as MsgDeleteWhoIsResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = 0;
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    return message;
  },
};

/** Msg defines the Msg Application. */
export interface Msg {
  /**
   * Register Application
   *
   * RegisterApplication registers a new application on the Registry module.
   */
  RegisterApplication(
    request: MsgRegisterApplication
  ): Promise<MsgRegisterApplicationResponse>;
  /**
   * Register Name
   *
   * RegisterName registers a .snr/ name for an account on the Registry module.
   */
  RegisterName(request: MsgRegisterName): Promise<MsgRegisterNameResponse>;
  /**
   * Access Name
   *
   * AccessName allows an account to access a .snr/ name on the Registry module. The equivalent of
   * of a traditional Login method.
   */
  AccessName(request: MsgAccessName): Promise<MsgAccessNameResponse>;
  /**
   * Update Name
   *
   * UpdateName allows an account to update a .snr/ name on the Registry module. Or,
   * in other words, link a new device to an existing .snr/ name.
   */
  UpdateName(request: MsgUpdateName): Promise<MsgUpdateNameResponse>;
  /**
   * Access Application
   *
   * AccessApplication allows an account to access an application on the Registry module.
   */
  AccessApplication(
    request: MsgAccessApplication
  ): Promise<MsgAccessApplicationResponse>;
  /**
   * Update Application
   *
   * UpdateApplication allows an account to update an application's config on the Registry module.
   */
  UpdateApplication(
    request: MsgUpdateApplication
  ): Promise<MsgUpdateApplicationResponse>;
  /**
   * Create WhoIs
   *
   * CreateWhoIs allows an account to create a WhoIs on the Registry module.
   */
  CreateWhoIs(request: MsgCreateWhoIs): Promise<MsgCreateWhoIsResponse>;
  /**
   * Update WhoIs
   *
   * UpdateWhoIs allows an account to update a WhoIs on the Registry module.
   */
  UpdateWhoIs(request: MsgUpdateWhoIs): Promise<MsgUpdateWhoIsResponse>;
  /**
   * Delete WhoIs
   *
   * DeleteWhoIs allows an account to delete a WhoIs on the Registry module.
   */
  DeleteWhoIs(request: MsgDeleteWhoIs): Promise<MsgDeleteWhoIsResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  RegisterApplication(
    request: MsgRegisterApplication
  ): Promise<MsgRegisterApplicationResponse> {
    const data = MsgRegisterApplication.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "RegisterApplication",
      data
    );
    return promise.then((data) =>
      MsgRegisterApplicationResponse.decode(new Reader(data))
    );
  }

  RegisterName(request: MsgRegisterName): Promise<MsgRegisterNameResponse> {
    const data = MsgRegisterName.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "RegisterName",
      data
    );
    return promise.then((data) =>
      MsgRegisterNameResponse.decode(new Reader(data))
    );
  }

  AccessName(request: MsgAccessName): Promise<MsgAccessNameResponse> {
    const data = MsgAccessName.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "AccessName",
      data
    );
    return promise.then((data) =>
      MsgAccessNameResponse.decode(new Reader(data))
    );
  }

  UpdateName(request: MsgUpdateName): Promise<MsgUpdateNameResponse> {
    const data = MsgUpdateName.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "UpdateName",
      data
    );
    return promise.then((data) =>
      MsgUpdateNameResponse.decode(new Reader(data))
    );
  }

  AccessApplication(
    request: MsgAccessApplication
  ): Promise<MsgAccessApplicationResponse> {
    const data = MsgAccessApplication.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "AccessApplication",
      data
    );
    return promise.then((data) =>
      MsgAccessApplicationResponse.decode(new Reader(data))
    );
  }

  UpdateApplication(
    request: MsgUpdateApplication
  ): Promise<MsgUpdateApplicationResponse> {
    const data = MsgUpdateApplication.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "UpdateApplication",
      data
    );
    return promise.then((data) =>
      MsgUpdateApplicationResponse.decode(new Reader(data))
    );
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

  DeleteWhoIs(request: MsgDeleteWhoIs): Promise<MsgDeleteWhoIsResponse> {
    const data = MsgDeleteWhoIs.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "DeleteWhoIs",
      data
    );
    return promise.then((data) =>
      MsgDeleteWhoIsResponse.decode(new Reader(data))
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
