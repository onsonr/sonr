/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Did, DidDocument } from "../registry/did";
export const protobufPackage = "sonrio.sonr.registry";
const baseMsgRegisterService = {
    creator: "",
    serviceName: "",
    publicKey: "",
};
export const MsgRegisterService = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.serviceName !== "") {
            writer.uint32(18).string(message.serviceName);
        }
        if (message.publicKey !== "") {
            writer.uint32(26).string(message.publicKey);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRegisterService };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.serviceName = reader.string();
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
    fromJSON(object) {
        const message = { ...baseMsgRegisterService };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.serviceName !== undefined && object.serviceName !== null) {
            message.serviceName = String(object.serviceName);
        }
        else {
            message.serviceName = "";
        }
        if (object.publicKey !== undefined && object.publicKey !== null) {
            message.publicKey = String(object.publicKey);
        }
        else {
            message.publicKey = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.serviceName !== undefined &&
            (obj.serviceName = message.serviceName);
        message.publicKey !== undefined && (obj.publicKey = message.publicKey);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgRegisterService };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.serviceName !== undefined && object.serviceName !== null) {
            message.serviceName = object.serviceName;
        }
        else {
            message.serviceName = "";
        }
        if (object.publicKey !== undefined && object.publicKey !== null) {
            message.publicKey = object.publicKey;
        }
        else {
            message.publicKey = "";
        }
        return message;
    },
};
const baseMsgRegisterServiceResponse = {};
export const MsgRegisterServiceResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgRegisterServiceResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgRegisterServiceResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgRegisterServiceResponse,
        };
        return message;
    },
};
const baseMsgRegisterName = {
    creator: "",
    deviceId: "",
    os: "",
    model: "",
    arch: "",
    publicKey: "",
    nameToRegister: "",
};
export const MsgRegisterName = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.deviceId !== "") {
            writer.uint32(18).string(message.deviceId);
        }
        if (message.os !== "") {
            writer.uint32(26).string(message.os);
        }
        if (message.model !== "") {
            writer.uint32(34).string(message.model);
        }
        if (message.arch !== "") {
            writer.uint32(42).string(message.arch);
        }
        if (message.publicKey !== "") {
            writer.uint32(50).string(message.publicKey);
        }
        if (message.nameToRegister !== "") {
            writer.uint32(58).string(message.nameToRegister);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRegisterName };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.deviceId = reader.string();
                    break;
                case 3:
                    message.os = reader.string();
                    break;
                case 4:
                    message.model = reader.string();
                    break;
                case 5:
                    message.arch = reader.string();
                    break;
                case 6:
                    message.publicKey = reader.string();
                    break;
                case 7:
                    message.nameToRegister = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgRegisterName };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.deviceId !== undefined && object.deviceId !== null) {
            message.deviceId = String(object.deviceId);
        }
        else {
            message.deviceId = "";
        }
        if (object.os !== undefined && object.os !== null) {
            message.os = String(object.os);
        }
        else {
            message.os = "";
        }
        if (object.model !== undefined && object.model !== null) {
            message.model = String(object.model);
        }
        else {
            message.model = "";
        }
        if (object.arch !== undefined && object.arch !== null) {
            message.arch = String(object.arch);
        }
        else {
            message.arch = "";
        }
        if (object.publicKey !== undefined && object.publicKey !== null) {
            message.publicKey = String(object.publicKey);
        }
        else {
            message.publicKey = "";
        }
        if (object.nameToRegister !== undefined && object.nameToRegister !== null) {
            message.nameToRegister = String(object.nameToRegister);
        }
        else {
            message.nameToRegister = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.deviceId !== undefined && (obj.deviceId = message.deviceId);
        message.os !== undefined && (obj.os = message.os);
        message.model !== undefined && (obj.model = message.model);
        message.arch !== undefined && (obj.arch = message.arch);
        message.publicKey !== undefined && (obj.publicKey = message.publicKey);
        message.nameToRegister !== undefined &&
            (obj.nameToRegister = message.nameToRegister);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgRegisterName };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.deviceId !== undefined && object.deviceId !== null) {
            message.deviceId = object.deviceId;
        }
        else {
            message.deviceId = "";
        }
        if (object.os !== undefined && object.os !== null) {
            message.os = object.os;
        }
        else {
            message.os = "";
        }
        if (object.model !== undefined && object.model !== null) {
            message.model = object.model;
        }
        else {
            message.model = "";
        }
        if (object.arch !== undefined && object.arch !== null) {
            message.arch = object.arch;
        }
        else {
            message.arch = "";
        }
        if (object.publicKey !== undefined && object.publicKey !== null) {
            message.publicKey = object.publicKey;
        }
        else {
            message.publicKey = "";
        }
        if (object.nameToRegister !== undefined && object.nameToRegister !== null) {
            message.nameToRegister = object.nameToRegister;
        }
        else {
            message.nameToRegister = "";
        }
        return message;
    },
};
const baseMsgRegisterNameResponse = { isSuccess: false };
export const MsgRegisterNameResponse = {
    encode(message, writer = Writer.create()) {
        if (message.isSuccess === true) {
            writer.uint32(8).bool(message.isSuccess);
        }
        if (message.did !== undefined) {
            Did.encode(message.did, writer.uint32(18).fork()).ldelim();
        }
        if (message.didDocument !== undefined) {
            DidDocument.encode(message.didDocument, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgRegisterNameResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.isSuccess = reader.bool();
                    break;
                case 2:
                    message.did = Did.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.didDocument = DidDocument.decode(reader, reader.uint32());
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
            ...baseMsgRegisterNameResponse,
        };
        if (object.isSuccess !== undefined && object.isSuccess !== null) {
            message.isSuccess = Boolean(object.isSuccess);
        }
        else {
            message.isSuccess = false;
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = Did.fromJSON(object.did);
        }
        else {
            message.did = undefined;
        }
        if (object.didDocument !== undefined && object.didDocument !== null) {
            message.didDocument = DidDocument.fromJSON(object.didDocument);
        }
        else {
            message.didDocument = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.isSuccess !== undefined && (obj.isSuccess = message.isSuccess);
        message.did !== undefined &&
            (obj.did = message.did ? Did.toJSON(message.did) : undefined);
        message.didDocument !== undefined &&
            (obj.didDocument = message.didDocument
                ? DidDocument.toJSON(message.didDocument)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMsgRegisterNameResponse,
        };
        if (object.isSuccess !== undefined && object.isSuccess !== null) {
            message.isSuccess = object.isSuccess;
        }
        else {
            message.isSuccess = false;
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = Did.fromPartial(object.did);
        }
        else {
            message.did = undefined;
        }
        if (object.didDocument !== undefined && object.didDocument !== null) {
            message.didDocument = DidDocument.fromPartial(object.didDocument);
        }
        else {
            message.didDocument = undefined;
        }
        return message;
    },
};
const baseMsgAccessName = {
    creator: "",
    name: "",
    publicKey: "",
    peerId: "",
};
export const MsgAccessName = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.name !== "") {
            writer.uint32(18).string(message.name);
        }
        if (message.publicKey !== "") {
            writer.uint32(26).string(message.publicKey);
        }
        if (message.peerId !== "") {
            writer.uint32(34).string(message.peerId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAccessName };
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
                    message.publicKey = reader.string();
                    break;
                case 4:
                    message.peerId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgAccessName };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
        }
        if (object.publicKey !== undefined && object.publicKey !== null) {
            message.publicKey = String(object.publicKey);
        }
        else {
            message.publicKey = "";
        }
        if (object.peerId !== undefined && object.peerId !== null) {
            message.peerId = String(object.peerId);
        }
        else {
            message.peerId = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.name !== undefined && (obj.name = message.name);
        message.publicKey !== undefined && (obj.publicKey = message.publicKey);
        message.peerId !== undefined && (obj.peerId = message.peerId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgAccessName };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
        }
        if (object.publicKey !== undefined && object.publicKey !== null) {
            message.publicKey = object.publicKey;
        }
        else {
            message.publicKey = "";
        }
        if (object.peerId !== undefined && object.peerId !== null) {
            message.peerId = object.peerId;
        }
        else {
            message.peerId = "";
        }
        return message;
    },
};
const baseMsgAccessNameResponse = {
    name: "",
    publicKey: "",
    peerId: "",
};
export const MsgAccessNameResponse = {
    encode(message, writer = Writer.create()) {
        if (message.name !== "") {
            writer.uint32(10).string(message.name);
        }
        if (message.publicKey !== "") {
            writer.uint32(18).string(message.publicKey);
        }
        if (message.peerId !== "") {
            writer.uint32(26).string(message.peerId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAccessNameResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.name = reader.string();
                    break;
                case 2:
                    message.publicKey = reader.string();
                    break;
                case 3:
                    message.peerId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgAccessNameResponse };
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
        }
        if (object.publicKey !== undefined && object.publicKey !== null) {
            message.publicKey = String(object.publicKey);
        }
        else {
            message.publicKey = "";
        }
        if (object.peerId !== undefined && object.peerId !== null) {
            message.peerId = String(object.peerId);
        }
        else {
            message.peerId = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.name !== undefined && (obj.name = message.name);
        message.publicKey !== undefined && (obj.publicKey = message.publicKey);
        message.peerId !== undefined && (obj.peerId = message.peerId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgAccessNameResponse };
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
        }
        if (object.publicKey !== undefined && object.publicKey !== null) {
            message.publicKey = object.publicKey;
        }
        else {
            message.publicKey = "";
        }
        if (object.peerId !== undefined && object.peerId !== null) {
            message.peerId = object.peerId;
        }
        else {
            message.peerId = "";
        }
        return message;
    },
};
const baseMsgUpdateName = { creator: "", name: "" };
export const MsgUpdateName = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.name !== "") {
            writer.uint32(18).string(message.name);
        }
        Object.entries(message.metadata).forEach(([key, value]) => {
            MsgUpdateName_MetadataEntry.encode({ key: key, value }, writer.uint32(26).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateName };
        message.metadata = {};
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
                    const entry3 = MsgUpdateName_MetadataEntry.decode(reader, reader.uint32());
                    if (entry3.value !== undefined) {
                        message.metadata[entry3.key] = entry3.value;
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
        const message = { ...baseMsgUpdateName };
        message.metadata = {};
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
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
        message.creator !== undefined && (obj.creator = message.creator);
        message.name !== undefined && (obj.name = message.name);
        obj.metadata = {};
        if (message.metadata) {
            Object.entries(message.metadata).forEach(([k, v]) => {
                obj.metadata[k] = v;
            });
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateName };
        message.metadata = {};
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
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
const baseMsgUpdateName_MetadataEntry = { key: "", value: "" };
export const MsgUpdateName_MetadataEntry = {
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
            ...baseMsgUpdateName_MetadataEntry,
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
            ...baseMsgUpdateName_MetadataEntry,
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
            ...baseMsgUpdateName_MetadataEntry,
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
const baseMsgUpdateNameResponse = {};
export const MsgUpdateNameResponse = {
    encode(message, writer = Writer.create()) {
        if (message.didDocument !== undefined) {
            DidDocument.encode(message.didDocument, writer.uint32(10).fork()).ldelim();
        }
        Object.entries(message.metadata).forEach(([key, value]) => {
            MsgUpdateNameResponse_MetadataEntry.encode({ key: key, value }, writer.uint32(18).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateNameResponse };
        message.metadata = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.didDocument = DidDocument.decode(reader, reader.uint32());
                    break;
                case 2:
                    const entry2 = MsgUpdateNameResponse_MetadataEntry.decode(reader, reader.uint32());
                    if (entry2.value !== undefined) {
                        message.metadata[entry2.key] = entry2.value;
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
        const message = { ...baseMsgUpdateNameResponse };
        message.metadata = {};
        if (object.didDocument !== undefined && object.didDocument !== null) {
            message.didDocument = DidDocument.fromJSON(object.didDocument);
        }
        else {
            message.didDocument = undefined;
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
        message.didDocument !== undefined &&
            (obj.didDocument = message.didDocument
                ? DidDocument.toJSON(message.didDocument)
                : undefined);
        obj.metadata = {};
        if (message.metadata) {
            Object.entries(message.metadata).forEach(([k, v]) => {
                obj.metadata[k] = v;
            });
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateNameResponse };
        message.metadata = {};
        if (object.didDocument !== undefined && object.didDocument !== null) {
            message.didDocument = DidDocument.fromPartial(object.didDocument);
        }
        else {
            message.didDocument = undefined;
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
const baseMsgUpdateNameResponse_MetadataEntry = { key: "", value: "" };
export const MsgUpdateNameResponse_MetadataEntry = {
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
            ...baseMsgUpdateNameResponse_MetadataEntry,
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
            ...baseMsgUpdateNameResponse_MetadataEntry,
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
            ...baseMsgUpdateNameResponse_MetadataEntry,
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
const baseMsgAccessService = { creator: "", did: "" };
export const MsgAccessService = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.did !== "") {
            writer.uint32(18).string(message.did);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAccessService };
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
    fromJSON(object) {
        const message = { ...baseMsgAccessService };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = String(object.did);
        }
        else {
            message.did = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.did !== undefined && (obj.did = message.did);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgAccessService };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = object.did;
        }
        else {
            message.did = "";
        }
        return message;
    },
};
const baseMsgAccessServiceResponse = { code: 0, message: "" };
export const MsgAccessServiceResponse = {
    encode(message, writer = Writer.create()) {
        if (message.code !== 0) {
            writer.uint32(8).int32(message.code);
        }
        if (message.message !== "") {
            writer.uint32(18).string(message.message);
        }
        Object.entries(message.metadata).forEach(([key, value]) => {
            MsgAccessServiceResponse_MetadataEntry.encode({ key: key, value }, writer.uint32(26).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgAccessServiceResponse,
        };
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
                    const entry3 = MsgAccessServiceResponse_MetadataEntry.decode(reader, reader.uint32());
                    if (entry3.value !== undefined) {
                        message.metadata[entry3.key] = entry3.value;
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
        const message = {
            ...baseMsgAccessServiceResponse,
        };
        message.metadata = {};
        if (object.code !== undefined && object.code !== null) {
            message.code = Number(object.code);
        }
        else {
            message.code = 0;
        }
        if (object.message !== undefined && object.message !== null) {
            message.message = String(object.message);
        }
        else {
            message.message = "";
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
        message.code !== undefined && (obj.code = message.code);
        message.message !== undefined && (obj.message = message.message);
        obj.metadata = {};
        if (message.metadata) {
            Object.entries(message.metadata).forEach(([k, v]) => {
                obj.metadata[k] = v;
            });
        }
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMsgAccessServiceResponse,
        };
        message.metadata = {};
        if (object.code !== undefined && object.code !== null) {
            message.code = object.code;
        }
        else {
            message.code = 0;
        }
        if (object.message !== undefined && object.message !== null) {
            message.message = object.message;
        }
        else {
            message.message = "";
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
const baseMsgAccessServiceResponse_MetadataEntry = {
    key: "",
    value: "",
};
export const MsgAccessServiceResponse_MetadataEntry = {
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
            ...baseMsgAccessServiceResponse_MetadataEntry,
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
            ...baseMsgAccessServiceResponse_MetadataEntry,
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
            ...baseMsgAccessServiceResponse_MetadataEntry,
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
const baseMsgUpdateService = { creator: "", did: "" };
export const MsgUpdateService = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.did !== "") {
            writer.uint32(18).string(message.did);
        }
        Object.entries(message.configuration).forEach(([key, value]) => {
            MsgUpdateService_ConfigurationEntry.encode({ key: key, value }, writer.uint32(26).fork()).ldelim();
        });
        Object.entries(message.metadata).forEach(([key, value]) => {
            MsgUpdateService_MetadataEntry.encode({ key: key, value }, writer.uint32(34).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateService };
        message.configuration = {};
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
                    const entry3 = MsgUpdateService_ConfigurationEntry.decode(reader, reader.uint32());
                    if (entry3.value !== undefined) {
                        message.configuration[entry3.key] = entry3.value;
                    }
                    break;
                case 4:
                    const entry4 = MsgUpdateService_MetadataEntry.decode(reader, reader.uint32());
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
        const message = { ...baseMsgUpdateService };
        message.configuration = {};
        message.metadata = {};
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = String(object.did);
        }
        else {
            message.did = "";
        }
        if (object.configuration !== undefined && object.configuration !== null) {
            Object.entries(object.configuration).forEach(([key, value]) => {
                message.configuration[key] = String(value);
            });
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
        message.creator !== undefined && (obj.creator = message.creator);
        message.did !== undefined && (obj.did = message.did);
        obj.configuration = {};
        if (message.configuration) {
            Object.entries(message.configuration).forEach(([k, v]) => {
                obj.configuration[k] = v;
            });
        }
        obj.metadata = {};
        if (message.metadata) {
            Object.entries(message.metadata).forEach(([k, v]) => {
                obj.metadata[k] = v;
            });
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateService };
        message.configuration = {};
        message.metadata = {};
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = object.did;
        }
        else {
            message.did = "";
        }
        if (object.configuration !== undefined && object.configuration !== null) {
            Object.entries(object.configuration).forEach(([key, value]) => {
                if (value !== undefined) {
                    message.configuration[key] = String(value);
                }
            });
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
const baseMsgUpdateService_ConfigurationEntry = { key: "", value: "" };
export const MsgUpdateService_ConfigurationEntry = {
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
            ...baseMsgUpdateService_ConfigurationEntry,
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
            ...baseMsgUpdateService_ConfigurationEntry,
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
            ...baseMsgUpdateService_ConfigurationEntry,
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
const baseMsgUpdateService_MetadataEntry = { key: "", value: "" };
export const MsgUpdateService_MetadataEntry = {
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
            ...baseMsgUpdateService_MetadataEntry,
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
            ...baseMsgUpdateService_MetadataEntry,
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
            ...baseMsgUpdateService_MetadataEntry,
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
const baseMsgUpdateServiceResponse = {};
export const MsgUpdateServiceResponse = {
    encode(message, writer = Writer.create()) {
        if (message.didDocument !== undefined) {
            DidDocument.encode(message.didDocument, writer.uint32(10).fork()).ldelim();
        }
        Object.entries(message.configuration).forEach(([key, value]) => {
            MsgUpdateServiceResponse_ConfigurationEntry.encode({ key: key, value }, writer.uint32(18).fork()).ldelim();
        });
        Object.entries(message.metadata).forEach(([key, value]) => {
            MsgUpdateServiceResponse_MetadataEntry.encode({ key: key, value }, writer.uint32(26).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgUpdateServiceResponse,
        };
        message.configuration = {};
        message.metadata = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.didDocument = DidDocument.decode(reader, reader.uint32());
                    break;
                case 2:
                    const entry2 = MsgUpdateServiceResponse_ConfigurationEntry.decode(reader, reader.uint32());
                    if (entry2.value !== undefined) {
                        message.configuration[entry2.key] = entry2.value;
                    }
                    break;
                case 3:
                    const entry3 = MsgUpdateServiceResponse_MetadataEntry.decode(reader, reader.uint32());
                    if (entry3.value !== undefined) {
                        message.metadata[entry3.key] = entry3.value;
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
        const message = {
            ...baseMsgUpdateServiceResponse,
        };
        message.configuration = {};
        message.metadata = {};
        if (object.didDocument !== undefined && object.didDocument !== null) {
            message.didDocument = DidDocument.fromJSON(object.didDocument);
        }
        else {
            message.didDocument = undefined;
        }
        if (object.configuration !== undefined && object.configuration !== null) {
            Object.entries(object.configuration).forEach(([key, value]) => {
                message.configuration[key] = String(value);
            });
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
        message.didDocument !== undefined &&
            (obj.didDocument = message.didDocument
                ? DidDocument.toJSON(message.didDocument)
                : undefined);
        obj.configuration = {};
        if (message.configuration) {
            Object.entries(message.configuration).forEach(([k, v]) => {
                obj.configuration[k] = v;
            });
        }
        obj.metadata = {};
        if (message.metadata) {
            Object.entries(message.metadata).forEach(([k, v]) => {
                obj.metadata[k] = v;
            });
        }
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMsgUpdateServiceResponse,
        };
        message.configuration = {};
        message.metadata = {};
        if (object.didDocument !== undefined && object.didDocument !== null) {
            message.didDocument = DidDocument.fromPartial(object.didDocument);
        }
        else {
            message.didDocument = undefined;
        }
        if (object.configuration !== undefined && object.configuration !== null) {
            Object.entries(object.configuration).forEach(([key, value]) => {
                if (value !== undefined) {
                    message.configuration[key] = String(value);
                }
            });
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
const baseMsgUpdateServiceResponse_ConfigurationEntry = {
    key: "",
    value: "",
};
export const MsgUpdateServiceResponse_ConfigurationEntry = {
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
            ...baseMsgUpdateServiceResponse_ConfigurationEntry,
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
            ...baseMsgUpdateServiceResponse_ConfigurationEntry,
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
            ...baseMsgUpdateServiceResponse_ConfigurationEntry,
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
const baseMsgUpdateServiceResponse_MetadataEntry = {
    key: "",
    value: "",
};
export const MsgUpdateServiceResponse_MetadataEntry = {
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
            ...baseMsgUpdateServiceResponse_MetadataEntry,
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
            ...baseMsgUpdateServiceResponse_MetadataEntry,
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
            ...baseMsgUpdateServiceResponse_MetadataEntry,
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
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    RegisterService(request) {
        const data = MsgRegisterService.encode(request).finish();
        const promise = this.rpc.request("sonrio.sonr.registry.Msg", "RegisterService", data);
        return promise.then((data) => MsgRegisterServiceResponse.decode(new Reader(data)));
    }
    RegisterName(request) {
        const data = MsgRegisterName.encode(request).finish();
        const promise = this.rpc.request("sonrio.sonr.registry.Msg", "RegisterName", data);
        return promise.then((data) => MsgRegisterNameResponse.decode(new Reader(data)));
    }
    AccessName(request) {
        const data = MsgAccessName.encode(request).finish();
        const promise = this.rpc.request("sonrio.sonr.registry.Msg", "AccessName", data);
        return promise.then((data) => MsgAccessNameResponse.decode(new Reader(data)));
    }
    UpdateName(request) {
        const data = MsgUpdateName.encode(request).finish();
        const promise = this.rpc.request("sonrio.sonr.registry.Msg", "UpdateName", data);
        return promise.then((data) => MsgUpdateNameResponse.decode(new Reader(data)));
    }
    AccessService(request) {
        const data = MsgAccessService.encode(request).finish();
        const promise = this.rpc.request("sonrio.sonr.registry.Msg", "AccessService", data);
        return promise.then((data) => MsgAccessServiceResponse.decode(new Reader(data)));
    }
    UpdateService(request) {
        const data = MsgUpdateService.encode(request).finish();
        const promise = this.rpc.request("sonrio.sonr.registry.Msg", "UpdateService", data);
        return promise.then((data) => MsgUpdateServiceResponse.decode(new Reader(data)));
    }
}
