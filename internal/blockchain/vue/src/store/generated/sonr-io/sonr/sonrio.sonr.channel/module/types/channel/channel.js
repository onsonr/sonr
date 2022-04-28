/* eslint-disable */
import { Did } from "../registry/did";
import { ObjectDoc } from "../object/object";
import { Peer } from "../registry/peer";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "sonrio.sonr.channel";
const baseChannel = { label: "", description: "" };
export const Channel = {
    encode(message, writer = Writer.create()) {
        if (message.label !== "") {
            writer.uint32(10).string(message.label);
        }
        if (message.description !== "") {
            writer.uint32(18).string(message.description);
        }
        if (message.did !== undefined) {
            Did.encode(message.did, writer.uint32(34).fork()).ldelim();
        }
        if (message.registeredObject !== undefined) {
            ObjectDoc.encode(message.registeredObject, writer.uint32(42).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseChannel };
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
                    message.did = Did.decode(reader, reader.uint32());
                    break;
                case 5:
                    message.registeredObject = ObjectDoc.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseChannel };
        if (object.label !== undefined && object.label !== null) {
            message.label = String(object.label);
        }
        else {
            message.label = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = String(object.description);
        }
        else {
            message.description = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = Did.fromJSON(object.did);
        }
        else {
            message.did = undefined;
        }
        if (object.registeredObject !== undefined &&
            object.registeredObject !== null) {
            message.registeredObject = ObjectDoc.fromJSON(object.registeredObject);
        }
        else {
            message.registeredObject = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.label !== undefined && (obj.label = message.label);
        message.description !== undefined &&
            (obj.description = message.description);
        message.did !== undefined &&
            (obj.did = message.did ? Did.toJSON(message.did) : undefined);
        message.registeredObject !== undefined &&
            (obj.registeredObject = message.registeredObject
                ? ObjectDoc.toJSON(message.registeredObject)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseChannel };
        if (object.label !== undefined && object.label !== null) {
            message.label = object.label;
        }
        else {
            message.label = "";
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = object.description;
        }
        else {
            message.description = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = Did.fromPartial(object.did);
        }
        else {
            message.did = undefined;
        }
        if (object.registeredObject !== undefined &&
            object.registeredObject !== null) {
            message.registeredObject = ObjectDoc.fromPartial(object.registeredObject);
        }
        else {
            message.registeredObject = undefined;
        }
        return message;
    },
};
const baseChannelMessage = {};
export const ChannelMessage = {
    encode(message, writer = Writer.create()) {
        if (message.peer !== undefined) {
            Peer.encode(message.peer, writer.uint32(10).fork()).ldelim();
        }
        if (message.did !== undefined) {
            Did.encode(message.did, writer.uint32(18).fork()).ldelim();
        }
        if (message.data !== undefined) {
            ObjectDoc.encode(message.data, writer.uint32(26).fork()).ldelim();
        }
        Object.entries(message.metadata).forEach(([key, value]) => {
            ChannelMessage_MetadataEntry.encode({ key: key, value }, writer.uint32(34).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseChannelMessage };
        message.metadata = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.peer = Peer.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.did = Did.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.data = ObjectDoc.decode(reader, reader.uint32());
                    break;
                case 4:
                    const entry4 = ChannelMessage_MetadataEntry.decode(reader, reader.uint32());
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
    fromJSON(object) {
        const message = { ...baseChannelMessage };
        message.metadata = {};
        if (object.peer !== undefined && object.peer !== null) {
            message.peer = Peer.fromJSON(object.peer);
        }
        else {
            message.peer = undefined;
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = Did.fromJSON(object.did);
        }
        else {
            message.did = undefined;
        }
        if (object.data !== undefined && object.data !== null) {
            message.data = ObjectDoc.fromJSON(object.data);
        }
        else {
            message.data = undefined;
        }
        if (object.metadata !== undefined && object.metadata !== null) {
            Object.entries(object.metadata).forEach(([key, value]) => {
                message.metadata[key] = String(value);
            });
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.peer !== undefined &&
            (obj.peer = message.peer ? Peer.toJSON(message.peer) : undefined);
        message.did !== undefined &&
            (obj.did = message.did ? Did.toJSON(message.did) : undefined);
        message.data !== undefined &&
            (obj.data = message.data ? ObjectDoc.toJSON(message.data) : undefined);
        obj.metadata = {};
        if (message.metadata) {
            Object.entries(message.metadata).forEach(([k, v]) => {
                obj.metadata[k] = v;
            });
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseChannelMessage };
        message.metadata = {};
        if (object.peer !== undefined && object.peer !== null) {
            message.peer = Peer.fromPartial(object.peer);
        }
        else {
            message.peer = undefined;
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = Did.fromPartial(object.did);
        }
        else {
            message.did = undefined;
        }
        if (object.data !== undefined && object.data !== null) {
            message.data = ObjectDoc.fromPartial(object.data);
        }
        else {
            message.data = undefined;
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
const baseChannelMessage_MetadataEntry = { key: "", value: "" };
export const ChannelMessage_MetadataEntry = {
    encode(message, writer = Writer.create()) {
        if (message.key !== "") {
            writer.uint32(10).string(message.key);
        }
        if (message.value !== "") {
            writer.uint32(18).string(message.value);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseChannelMessage_MetadataEntry,
        };
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
    fromJSON(object) {
        const message = {
            ...baseChannelMessage_MetadataEntry,
        };
        if (object.key !== undefined && object.key !== null) {
            message.key = String(object.key);
        }
        else {
            message.key = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = String(object.value);
        }
        else {
            message.value = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.key !== undefined && (obj.key = message.key);
        message.value !== undefined && (obj.value = message.value);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseChannelMessage_MetadataEntry,
        };
        if (object.key !== undefined && object.key !== null) {
            message.key = object.key;
        }
        else {
            message.key = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = object.value;
        }
        else {
            message.value = "";
        }
        return message;
    },
};
