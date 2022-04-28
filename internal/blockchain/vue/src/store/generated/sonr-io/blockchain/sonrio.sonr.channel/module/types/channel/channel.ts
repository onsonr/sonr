/* eslint-disable */
import { ObjectDoc } from "../object/object";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sonrio.sonr.channel";

export interface ChannelDoc {
  /** Label is human-readable name of the channel. */
  label: string;
  /** Description is a human-readable description of the channel. */
  description: string;
  /** Did is the identifier of the channel. */
  did: string;
  /** RegisterdObject is the object that is registered as the payload for the channel. */
  registered_object: ObjectDoc | undefined;
}

/** ChannelMessage is a message sent to a channel. */
export interface ChannelMessage {
  /** Owner is the peer that originated the message. */
  peer_did: string;
  /** Did is the identifier of the channel. */
  did: string;
  /** Data is the data being sent. */
  object: ObjectDoc | undefined;
  /** Metadata is the metadata associated with the message. */
  metadata: { [key: string]: string };
}

export interface ChannelMessage_MetadataEntry {
  key: string;
  value: string;
}

const baseChannelDoc: object = { label: "", description: "", did: "" };

export const ChannelDoc = {
  encode(message: ChannelDoc, writer: Writer = Writer.create()): Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.did !== "") {
      writer.uint32(34).string(message.did);
    }
    if (message.registered_object !== undefined) {
      ObjectDoc.encode(
        message.registered_object,
        writer.uint32(42).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ChannelDoc {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseChannelDoc } as ChannelDoc;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 4:
          message.did = reader.string();
          break;
        case 5:
          message.registered_object = ObjectDoc.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ChannelDoc {
    const message = { ...baseChannelDoc } as ChannelDoc;
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
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (
      object.registered_object !== undefined &&
      object.registered_object !== null
    ) {
      message.registered_object = ObjectDoc.fromJSON(object.registered_object);
    } else {
      message.registered_object = undefined;
    }
    return message;
  },

  toJSON(message: ChannelDoc): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.description !== undefined &&
      (obj.description = message.description);
    message.did !== undefined && (obj.did = message.did);
    message.registered_object !== undefined &&
      (obj.registered_object = message.registered_object
        ? ObjectDoc.toJSON(message.registered_object)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<ChannelDoc>): ChannelDoc {
    const message = { ...baseChannelDoc } as ChannelDoc;
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
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (
      object.registered_object !== undefined &&
      object.registered_object !== null
    ) {
      message.registered_object = ObjectDoc.fromPartial(
        object.registered_object
      );
    } else {
      message.registered_object = undefined;
    }
    return message;
  },
};

const baseChannelMessage: object = { peer_did: "", did: "" };

export const ChannelMessage = {
  encode(message: ChannelMessage, writer: Writer = Writer.create()): Writer {
    if (message.peer_did !== "") {
      writer.uint32(10).string(message.peer_did);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.object !== undefined) {
      ObjectDoc.encode(message.object, writer.uint32(26).fork()).ldelim();
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ChannelMessage_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ChannelMessage {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseChannelMessage } as ChannelMessage;
    message.metadata = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.peer_did = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.object = ObjectDoc.decode(reader, reader.uint32());
          break;
        case 4:
          const entry4 = ChannelMessage_MetadataEntry.decode(
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

  fromJSON(object: any): ChannelMessage {
    const message = { ...baseChannelMessage } as ChannelMessage;
    message.metadata = {};
    if (object.peer_did !== undefined && object.peer_did !== null) {
      message.peer_did = String(object.peer_did);
    } else {
      message.peer_did = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.object !== undefined && object.object !== null) {
      message.object = ObjectDoc.fromJSON(object.object);
    } else {
      message.object = undefined;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    return message;
  },

  toJSON(message: ChannelMessage): unknown {
    const obj: any = {};
    message.peer_did !== undefined && (obj.peer_did = message.peer_did);
    message.did !== undefined && (obj.did = message.did);
    message.object !== undefined &&
      (obj.object = message.object
        ? ObjectDoc.toJSON(message.object)
        : undefined);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ChannelMessage>): ChannelMessage {
    const message = { ...baseChannelMessage } as ChannelMessage;
    message.metadata = {};
    if (object.peer_did !== undefined && object.peer_did !== null) {
      message.peer_did = object.peer_did;
    } else {
      message.peer_did = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.object !== undefined && object.object !== null) {
      message.object = ObjectDoc.fromPartial(object.object);
    } else {
      message.object = undefined;
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

const baseChannelMessage_MetadataEntry: object = { key: "", value: "" };

export const ChannelMessage_MetadataEntry = {
  encode(
    message: ChannelMessage_MetadataEntry,
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
  ): ChannelMessage_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseChannelMessage_MetadataEntry,
    } as ChannelMessage_MetadataEntry;
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

  fromJSON(object: any): ChannelMessage_MetadataEntry {
    const message = {
      ...baseChannelMessage_MetadataEntry,
    } as ChannelMessage_MetadataEntry;
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

  toJSON(message: ChannelMessage_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ChannelMessage_MetadataEntry>
  ): ChannelMessage_MetadataEntry {
    const message = {
      ...baseChannelMessage_MetadataEntry,
    } as ChannelMessage_MetadataEntry;
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
