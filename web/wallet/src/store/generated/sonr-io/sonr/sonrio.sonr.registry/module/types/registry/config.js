/* eslint-disable */
import { Did } from "../registry/did";
import { ObjectDoc } from "../object/object";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "sonrio.sonr.registry";
const baseServiceConfig = {
    name: "",
    description: "",
    endpoints: "",
    version: "",
};
export const ServiceConfig = {
    encode(message, writer = Writer.create()) {
        if (message.name !== "") {
            writer.uint32(10).string(message.name);
        }
        if (message.description !== "") {
            writer.uint32(18).string(message.description);
        }
        if (message.did !== undefined) {
            Did.encode(message.did, writer.uint32(26).fork()).ldelim();
        }
        for (const v of message.maintainers) {
            Did.encode(v, writer.uint32(34).fork()).ldelim();
        }
        for (const v of message.channels) {
            Did.encode(v, writer.uint32(42).fork()).ldelim();
        }
        for (const v of message.buckets) {
            Did.encode(v, writer.uint32(50).fork()).ldelim();
        }
        Object.entries(message.objects).forEach(([key, value]) => {
            ServiceConfig_ObjectsEntry.encode({ key: key, value }, writer.uint32(58).fork()).ldelim();
        });
        for (const v of message.endpoints) {
            writer.uint32(66).string(v);
        }
        Object.entries(message.metadata).forEach(([key, value]) => {
            ServiceConfig_MetadataEntry.encode({ key: key, value }, writer.uint32(74).fork()).ldelim();
        });
        if (message.version !== "") {
            writer.uint32(82).string(message.version);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseServiceConfig };
        message.maintainers = [];
        message.channels = [];
        message.buckets = [];
        message.objects = {};
        message.endpoints = [];
        message.metadata = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.name = reader.string();
                    break;
                case 2:
                    message.description = reader.string();
                    break;
                case 3:
                    message.did = Did.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.maintainers.push(Did.decode(reader, reader.uint32()));
                    break;
                case 5:
                    message.channels.push(Did.decode(reader, reader.uint32()));
                    break;
                case 6:
                    message.buckets.push(Did.decode(reader, reader.uint32()));
                    break;
                case 7:
                    const entry7 = ServiceConfig_ObjectsEntry.decode(reader, reader.uint32());
                    if (entry7.value !== undefined) {
                        message.objects[entry7.key] = entry7.value;
                    }
                    break;
                case 8:
                    message.endpoints.push(reader.string());
                    break;
                case 9:
                    const entry9 = ServiceConfig_MetadataEntry.decode(reader, reader.uint32());
                    if (entry9.value !== undefined) {
                        message.metadata[entry9.key] = entry9.value;
                    }
                    break;
                case 10:
                    message.version = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseServiceConfig };
        message.maintainers = [];
        message.channels = [];
        message.buckets = [];
        message.objects = {};
        message.endpoints = [];
        message.metadata = {};
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
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
        if (object.maintainers !== undefined && object.maintainers !== null) {
            for (const e of object.maintainers) {
                message.maintainers.push(Did.fromJSON(e));
            }
        }
        if (object.channels !== undefined && object.channels !== null) {
            for (const e of object.channels) {
                message.channels.push(Did.fromJSON(e));
            }
        }
        if (object.buckets !== undefined && object.buckets !== null) {
            for (const e of object.buckets) {
                message.buckets.push(Did.fromJSON(e));
            }
        }
        if (object.objects !== undefined && object.objects !== null) {
            Object.entries(object.objects).forEach(([key, value]) => {
                message.objects[key] = ObjectDoc.fromJSON(value);
            });
        }
        if (object.endpoints !== undefined && object.endpoints !== null) {
            for (const e of object.endpoints) {
                message.endpoints.push(String(e));
            }
        }
        if (object.metadata !== undefined && object.metadata !== null) {
            Object.entries(object.metadata).forEach(([key, value]) => {
                message.metadata[key] = String(value);
            });
        }
        if (object.version !== undefined && object.version !== null) {
            message.version = String(object.version);
        }
        else {
            message.version = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.name !== undefined && (obj.name = message.name);
        message.description !== undefined &&
            (obj.description = message.description);
        message.did !== undefined &&
            (obj.did = message.did ? Did.toJSON(message.did) : undefined);
        if (message.maintainers) {
            obj.maintainers = message.maintainers.map((e) => e ? Did.toJSON(e) : undefined);
        }
        else {
            obj.maintainers = [];
        }
        if (message.channels) {
            obj.channels = message.channels.map((e) => e ? Did.toJSON(e) : undefined);
        }
        else {
            obj.channels = [];
        }
        if (message.buckets) {
            obj.buckets = message.buckets.map((e) => (e ? Did.toJSON(e) : undefined));
        }
        else {
            obj.buckets = [];
        }
        obj.objects = {};
        if (message.objects) {
            Object.entries(message.objects).forEach(([k, v]) => {
                obj.objects[k] = ObjectDoc.toJSON(v);
            });
        }
        if (message.endpoints) {
            obj.endpoints = message.endpoints.map((e) => e);
        }
        else {
            obj.endpoints = [];
        }
        obj.metadata = {};
        if (message.metadata) {
            Object.entries(message.metadata).forEach(([k, v]) => {
                obj.metadata[k] = v;
            });
        }
        message.version !== undefined && (obj.version = message.version);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseServiceConfig };
        message.maintainers = [];
        message.channels = [];
        message.buckets = [];
        message.objects = {};
        message.endpoints = [];
        message.metadata = {};
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
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
        if (object.maintainers !== undefined && object.maintainers !== null) {
            for (const e of object.maintainers) {
                message.maintainers.push(Did.fromPartial(e));
            }
        }
        if (object.channels !== undefined && object.channels !== null) {
            for (const e of object.channels) {
                message.channels.push(Did.fromPartial(e));
            }
        }
        if (object.buckets !== undefined && object.buckets !== null) {
            for (const e of object.buckets) {
                message.buckets.push(Did.fromPartial(e));
            }
        }
        if (object.objects !== undefined && object.objects !== null) {
            Object.entries(object.objects).forEach(([key, value]) => {
                if (value !== undefined) {
                    message.objects[key] = ObjectDoc.fromPartial(value);
                }
            });
        }
        if (object.endpoints !== undefined && object.endpoints !== null) {
            for (const e of object.endpoints) {
                message.endpoints.push(e);
            }
        }
        if (object.metadata !== undefined && object.metadata !== null) {
            Object.entries(object.metadata).forEach(([key, value]) => {
                if (value !== undefined) {
                    message.metadata[key] = String(value);
                }
            });
        }
        if (object.version !== undefined && object.version !== null) {
            message.version = object.version;
        }
        else {
            message.version = "";
        }
        return message;
    },
};
const baseServiceConfig_ObjectsEntry = { key: "" };
export const ServiceConfig_ObjectsEntry = {
    encode(message, writer = Writer.create()) {
        if (message.key !== "") {
            writer.uint32(10).string(message.key);
        }
        if (message.value !== undefined) {
            ObjectDoc.encode(message.value, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseServiceConfig_ObjectsEntry,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.key = reader.string();
                    break;
                case 2:
                    message.value = ObjectDoc.decode(reader, reader.uint32());
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
            ...baseServiceConfig_ObjectsEntry,
        };
        if (object.key !== undefined && object.key !== null) {
            message.key = String(object.key);
        }
        else {
            message.key = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = ObjectDoc.fromJSON(object.value);
        }
        else {
            message.value = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.key !== undefined && (obj.key = message.key);
        message.value !== undefined &&
            (obj.value = message.value ? ObjectDoc.toJSON(message.value) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseServiceConfig_ObjectsEntry,
        };
        if (object.key !== undefined && object.key !== null) {
            message.key = object.key;
        }
        else {
            message.key = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = ObjectDoc.fromPartial(object.value);
        }
        else {
            message.value = undefined;
        }
        return message;
    },
};
const baseServiceConfig_MetadataEntry = { key: "", value: "" };
export const ServiceConfig_MetadataEntry = {
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
            ...baseServiceConfig_MetadataEntry,
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
            ...baseServiceConfig_MetadataEntry,
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
            ...baseServiceConfig_MetadataEntry,
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
