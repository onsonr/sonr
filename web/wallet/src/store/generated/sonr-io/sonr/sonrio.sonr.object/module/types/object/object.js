/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "sonrio.sonr.object";
/** ObjectFieldType is the type of the field */
export var ObjectFieldType;
(function (ObjectFieldType) {
    /** OBJECT_FIELD_TYPE_UNSPECIFIED - ObjectFieldTypeUnspecified is the default value */
    ObjectFieldType[ObjectFieldType["OBJECT_FIELD_TYPE_UNSPECIFIED"] = 0] = "OBJECT_FIELD_TYPE_UNSPECIFIED";
    /** OBJECT_FIELD_TYPE_STRING - ObjectFieldTypeString is a string or text field */
    ObjectFieldType[ObjectFieldType["OBJECT_FIELD_TYPE_STRING"] = 1] = "OBJECT_FIELD_TYPE_STRING";
    /** OBJECT_FIELD_TYPE_NUMBER - ObjectFieldTypeInt is an integer */
    ObjectFieldType[ObjectFieldType["OBJECT_FIELD_TYPE_NUMBER"] = 2] = "OBJECT_FIELD_TYPE_NUMBER";
    /** OBJECT_FIELD_TYPE_BOOL - ObjectFieldTypeBool is a boolean */
    ObjectFieldType[ObjectFieldType["OBJECT_FIELD_TYPE_BOOL"] = 3] = "OBJECT_FIELD_TYPE_BOOL";
    /** OBJECT_FIELD_TYPE_ARRAY - ObjectFieldTypeArray is a list of values */
    ObjectFieldType[ObjectFieldType["OBJECT_FIELD_TYPE_ARRAY"] = 4] = "OBJECT_FIELD_TYPE_ARRAY";
    /** OBJECT_FIELD_TYPE_TIMESTAMP - ObjectFieldTypeDateTime is a datetime */
    ObjectFieldType[ObjectFieldType["OBJECT_FIELD_TYPE_TIMESTAMP"] = 5] = "OBJECT_FIELD_TYPE_TIMESTAMP";
    /** OBJECT_FIELD_TYPE_GEOPOINT - ObjectFieldTypeGeopoint is a geopoint */
    ObjectFieldType[ObjectFieldType["OBJECT_FIELD_TYPE_GEOPOINT"] = 6] = "OBJECT_FIELD_TYPE_GEOPOINT";
    /** OBJECT_FIELD_TYPE_BLOB - ObjectFieldTypeBlob is a blob of data */
    ObjectFieldType[ObjectFieldType["OBJECT_FIELD_TYPE_BLOB"] = 7] = "OBJECT_FIELD_TYPE_BLOB";
    /** OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS - ObjectFieldTypeETU is a pointer to an Ethereum account address. */
    ObjectFieldType[ObjectFieldType["OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS"] = 8] = "OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS";
    ObjectFieldType[ObjectFieldType["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(ObjectFieldType || (ObjectFieldType = {}));
export function objectFieldTypeFromJSON(object) {
    switch (object) {
        case 0:
        case "OBJECT_FIELD_TYPE_UNSPECIFIED":
            return ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED;
        case 1:
        case "OBJECT_FIELD_TYPE_STRING":
            return ObjectFieldType.OBJECT_FIELD_TYPE_STRING;
        case 2:
        case "OBJECT_FIELD_TYPE_NUMBER":
            return ObjectFieldType.OBJECT_FIELD_TYPE_NUMBER;
        case 3:
        case "OBJECT_FIELD_TYPE_BOOL":
            return ObjectFieldType.OBJECT_FIELD_TYPE_BOOL;
        case 4:
        case "OBJECT_FIELD_TYPE_ARRAY":
            return ObjectFieldType.OBJECT_FIELD_TYPE_ARRAY;
        case 5:
        case "OBJECT_FIELD_TYPE_TIMESTAMP":
            return ObjectFieldType.OBJECT_FIELD_TYPE_TIMESTAMP;
        case 6:
        case "OBJECT_FIELD_TYPE_GEOPOINT":
            return ObjectFieldType.OBJECT_FIELD_TYPE_GEOPOINT;
        case 7:
        case "OBJECT_FIELD_TYPE_BLOB":
            return ObjectFieldType.OBJECT_FIELD_TYPE_BLOB;
        case 8:
        case "OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS":
            return ObjectFieldType.OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS;
        case -1:
        case "UNRECOGNIZED":
        default:
            return ObjectFieldType.UNRECOGNIZED;
    }
}
export function objectFieldTypeToJSON(object) {
    switch (object) {
        case ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED:
            return "OBJECT_FIELD_TYPE_UNSPECIFIED";
        case ObjectFieldType.OBJECT_FIELD_TYPE_STRING:
            return "OBJECT_FIELD_TYPE_STRING";
        case ObjectFieldType.OBJECT_FIELD_TYPE_NUMBER:
            return "OBJECT_FIELD_TYPE_NUMBER";
        case ObjectFieldType.OBJECT_FIELD_TYPE_BOOL:
            return "OBJECT_FIELD_TYPE_BOOL";
        case ObjectFieldType.OBJECT_FIELD_TYPE_ARRAY:
            return "OBJECT_FIELD_TYPE_ARRAY";
        case ObjectFieldType.OBJECT_FIELD_TYPE_TIMESTAMP:
            return "OBJECT_FIELD_TYPE_TIMESTAMP";
        case ObjectFieldType.OBJECT_FIELD_TYPE_GEOPOINT:
            return "OBJECT_FIELD_TYPE_GEOPOINT";
        case ObjectFieldType.OBJECT_FIELD_TYPE_BLOB:
            return "OBJECT_FIELD_TYPE_BLOB";
        case ObjectFieldType.OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS:
            return "OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS";
        default:
            return "UNKNOWN";
    }
}
const baseObjectDoc = {
    label: "",
    description: "",
    did: "",
    bucketDid: "",
};
export const ObjectDoc = {
    encode(message, writer = Writer.create()) {
        if (message.label !== "") {
            writer.uint32(10).string(message.label);
        }
        if (message.description !== "") {
            writer.uint32(18).string(message.description);
        }
        if (message.did !== "") {
            writer.uint32(26).string(message.did);
        }
        if (message.bucketDid !== "") {
            writer.uint32(34).string(message.bucketDid);
        }
        Object.entries(message.fields).forEach(([key, value]) => {
            ObjectDoc_FieldsEntry.encode({ key: key, value }, writer.uint32(42).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseObjectDoc };
        message.fields = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.label = reader.string();
                    break;
                case 2:
                    message.description = reader.string();
                    break;
                case 3:
                    message.did = reader.string();
                    break;
                case 4:
                    message.bucketDid = reader.string();
                    break;
                case 5:
                    const entry5 = ObjectDoc_FieldsEntry.decode(reader, reader.uint32());
                    if (entry5.value !== undefined) {
                        message.fields[entry5.key] = entry5.value;
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
        const message = { ...baseObjectDoc };
        message.fields = {};
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
            message.did = String(object.did);
        }
        else {
            message.did = "";
        }
        if (object.bucketDid !== undefined && object.bucketDid !== null) {
            message.bucketDid = String(object.bucketDid);
        }
        else {
            message.bucketDid = "";
        }
        if (object.fields !== undefined && object.fields !== null) {
            Object.entries(object.fields).forEach(([key, value]) => {
                message.fields[key] = ObjectField.fromJSON(value);
            });
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.label !== undefined && (obj.label = message.label);
        message.description !== undefined &&
            (obj.description = message.description);
        message.did !== undefined && (obj.did = message.did);
        message.bucketDid !== undefined && (obj.bucketDid = message.bucketDid);
        obj.fields = {};
        if (message.fields) {
            Object.entries(message.fields).forEach(([k, v]) => {
                obj.fields[k] = ObjectField.toJSON(v);
            });
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseObjectDoc };
        message.fields = {};
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
            message.did = object.did;
        }
        else {
            message.did = "";
        }
        if (object.bucketDid !== undefined && object.bucketDid !== null) {
            message.bucketDid = object.bucketDid;
        }
        else {
            message.bucketDid = "";
        }
        if (object.fields !== undefined && object.fields !== null) {
            Object.entries(object.fields).forEach(([key, value]) => {
                if (value !== undefined) {
                    message.fields[key] = ObjectField.fromPartial(value);
                }
            });
        }
        return message;
    },
};
const baseObjectDoc_FieldsEntry = { key: "" };
export const ObjectDoc_FieldsEntry = {
    encode(message, writer = Writer.create()) {
        if (message.key !== "") {
            writer.uint32(10).string(message.key);
        }
        if (message.value !== undefined) {
            ObjectField.encode(message.value, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseObjectDoc_FieldsEntry };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.key = reader.string();
                    break;
                case 2:
                    message.value = ObjectField.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseObjectDoc_FieldsEntry };
        if (object.key !== undefined && object.key !== null) {
            message.key = String(object.key);
        }
        else {
            message.key = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = ObjectField.fromJSON(object.value);
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
            (obj.value = message.value
                ? ObjectField.toJSON(message.value)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseObjectDoc_FieldsEntry };
        if (object.key !== undefined && object.key !== null) {
            message.key = object.key;
        }
        else {
            message.key = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = ObjectField.fromPartial(object.value);
        }
        else {
            message.value = undefined;
        }
        return message;
    },
};
const baseObjectField = { label: "", type: 0, did: "" };
export const ObjectField = {
    encode(message, writer = Writer.create()) {
        if (message.label !== "") {
            writer.uint32(10).string(message.label);
        }
        if (message.type !== 0) {
            writer.uint32(16).int32(message.type);
        }
        if (message.did !== "") {
            writer.uint32(26).string(message.did);
        }
        if (message.stringValue !== undefined) {
            ObjectFieldText.encode(message.stringValue, writer.uint32(34).fork()).ldelim();
        }
        if (message.numberValue !== undefined) {
            ObjectFieldNumber.encode(message.numberValue, writer.uint32(42).fork()).ldelim();
        }
        if (message.boolValue !== undefined) {
            ObjectFieldBool.encode(message.boolValue, writer.uint32(50).fork()).ldelim();
        }
        if (message.arrayValue !== undefined) {
            ObjectFieldArray.encode(message.arrayValue, writer.uint32(58).fork()).ldelim();
        }
        if (message.timeStampValue !== undefined) {
            ObjectFieldTime.encode(message.timeStampValue, writer.uint32(66).fork()).ldelim();
        }
        if (message.geopointValue !== undefined) {
            ObjectFieldGeopoint.encode(message.geopointValue, writer.uint32(74).fork()).ldelim();
        }
        if (message.blobValue !== undefined) {
            ObjectFieldBlob.encode(message.blobValue, writer.uint32(82).fork()).ldelim();
        }
        if (message.blockchainAddrValue !== undefined) {
            ObjectFieldBlockchainAddress.encode(message.blockchainAddrValue, writer.uint32(98).fork()).ldelim();
        }
        Object.entries(message.metadata).forEach(([key, value]) => {
            ObjectField_MetadataEntry.encode({ key: key, value }, writer.uint32(106).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseObjectField };
        message.metadata = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.label = reader.string();
                    break;
                case 2:
                    message.type = reader.int32();
                    break;
                case 3:
                    message.did = reader.string();
                    break;
                case 4:
                    message.stringValue = ObjectFieldText.decode(reader, reader.uint32());
                    break;
                case 5:
                    message.numberValue = ObjectFieldNumber.decode(reader, reader.uint32());
                    break;
                case 6:
                    message.boolValue = ObjectFieldBool.decode(reader, reader.uint32());
                    break;
                case 7:
                    message.arrayValue = ObjectFieldArray.decode(reader, reader.uint32());
                    break;
                case 8:
                    message.timeStampValue = ObjectFieldTime.decode(reader, reader.uint32());
                    break;
                case 9:
                    message.geopointValue = ObjectFieldGeopoint.decode(reader, reader.uint32());
                    break;
                case 10:
                    message.blobValue = ObjectFieldBlob.decode(reader, reader.uint32());
                    break;
                case 12:
                    message.blockchainAddrValue = ObjectFieldBlockchainAddress.decode(reader, reader.uint32());
                    break;
                case 13:
                    const entry13 = ObjectField_MetadataEntry.decode(reader, reader.uint32());
                    if (entry13.value !== undefined) {
                        message.metadata[entry13.key] = entry13.value;
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
        const message = { ...baseObjectField };
        message.metadata = {};
        if (object.label !== undefined && object.label !== null) {
            message.label = String(object.label);
        }
        else {
            message.label = "";
        }
        if (object.type !== undefined && object.type !== null) {
            message.type = objectFieldTypeFromJSON(object.type);
        }
        else {
            message.type = 0;
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = String(object.did);
        }
        else {
            message.did = "";
        }
        if (object.stringValue !== undefined && object.stringValue !== null) {
            message.stringValue = ObjectFieldText.fromJSON(object.stringValue);
        }
        else {
            message.stringValue = undefined;
        }
        if (object.numberValue !== undefined && object.numberValue !== null) {
            message.numberValue = ObjectFieldNumber.fromJSON(object.numberValue);
        }
        else {
            message.numberValue = undefined;
        }
        if (object.boolValue !== undefined && object.boolValue !== null) {
            message.boolValue = ObjectFieldBool.fromJSON(object.boolValue);
        }
        else {
            message.boolValue = undefined;
        }
        if (object.arrayValue !== undefined && object.arrayValue !== null) {
            message.arrayValue = ObjectFieldArray.fromJSON(object.arrayValue);
        }
        else {
            message.arrayValue = undefined;
        }
        if (object.timeStampValue !== undefined && object.timeStampValue !== null) {
            message.timeStampValue = ObjectFieldTime.fromJSON(object.timeStampValue);
        }
        else {
            message.timeStampValue = undefined;
        }
        if (object.geopointValue !== undefined && object.geopointValue !== null) {
            message.geopointValue = ObjectFieldGeopoint.fromJSON(object.geopointValue);
        }
        else {
            message.geopointValue = undefined;
        }
        if (object.blobValue !== undefined && object.blobValue !== null) {
            message.blobValue = ObjectFieldBlob.fromJSON(object.blobValue);
        }
        else {
            message.blobValue = undefined;
        }
        if (object.blockchainAddrValue !== undefined &&
            object.blockchainAddrValue !== null) {
            message.blockchainAddrValue = ObjectFieldBlockchainAddress.fromJSON(object.blockchainAddrValue);
        }
        else {
            message.blockchainAddrValue = undefined;
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
        message.label !== undefined && (obj.label = message.label);
        message.type !== undefined &&
            (obj.type = objectFieldTypeToJSON(message.type));
        message.did !== undefined && (obj.did = message.did);
        message.stringValue !== undefined &&
            (obj.stringValue = message.stringValue
                ? ObjectFieldText.toJSON(message.stringValue)
                : undefined);
        message.numberValue !== undefined &&
            (obj.numberValue = message.numberValue
                ? ObjectFieldNumber.toJSON(message.numberValue)
                : undefined);
        message.boolValue !== undefined &&
            (obj.boolValue = message.boolValue
                ? ObjectFieldBool.toJSON(message.boolValue)
                : undefined);
        message.arrayValue !== undefined &&
            (obj.arrayValue = message.arrayValue
                ? ObjectFieldArray.toJSON(message.arrayValue)
                : undefined);
        message.timeStampValue !== undefined &&
            (obj.timeStampValue = message.timeStampValue
                ? ObjectFieldTime.toJSON(message.timeStampValue)
                : undefined);
        message.geopointValue !== undefined &&
            (obj.geopointValue = message.geopointValue
                ? ObjectFieldGeopoint.toJSON(message.geopointValue)
                : undefined);
        message.blobValue !== undefined &&
            (obj.blobValue = message.blobValue
                ? ObjectFieldBlob.toJSON(message.blobValue)
                : undefined);
        message.blockchainAddrValue !== undefined &&
            (obj.blockchainAddrValue = message.blockchainAddrValue
                ? ObjectFieldBlockchainAddress.toJSON(message.blockchainAddrValue)
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
        const message = { ...baseObjectField };
        message.metadata = {};
        if (object.label !== undefined && object.label !== null) {
            message.label = object.label;
        }
        else {
            message.label = "";
        }
        if (object.type !== undefined && object.type !== null) {
            message.type = object.type;
        }
        else {
            message.type = 0;
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = object.did;
        }
        else {
            message.did = "";
        }
        if (object.stringValue !== undefined && object.stringValue !== null) {
            message.stringValue = ObjectFieldText.fromPartial(object.stringValue);
        }
        else {
            message.stringValue = undefined;
        }
        if (object.numberValue !== undefined && object.numberValue !== null) {
            message.numberValue = ObjectFieldNumber.fromPartial(object.numberValue);
        }
        else {
            message.numberValue = undefined;
        }
        if (object.boolValue !== undefined && object.boolValue !== null) {
            message.boolValue = ObjectFieldBool.fromPartial(object.boolValue);
        }
        else {
            message.boolValue = undefined;
        }
        if (object.arrayValue !== undefined && object.arrayValue !== null) {
            message.arrayValue = ObjectFieldArray.fromPartial(object.arrayValue);
        }
        else {
            message.arrayValue = undefined;
        }
        if (object.timeStampValue !== undefined && object.timeStampValue !== null) {
            message.timeStampValue = ObjectFieldTime.fromPartial(object.timeStampValue);
        }
        else {
            message.timeStampValue = undefined;
        }
        if (object.geopointValue !== undefined && object.geopointValue !== null) {
            message.geopointValue = ObjectFieldGeopoint.fromPartial(object.geopointValue);
        }
        else {
            message.geopointValue = undefined;
        }
        if (object.blobValue !== undefined && object.blobValue !== null) {
            message.blobValue = ObjectFieldBlob.fromPartial(object.blobValue);
        }
        else {
            message.blobValue = undefined;
        }
        if (object.blockchainAddrValue !== undefined &&
            object.blockchainAddrValue !== null) {
            message.blockchainAddrValue = ObjectFieldBlockchainAddress.fromPartial(object.blockchainAddrValue);
        }
        else {
            message.blockchainAddrValue = undefined;
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
const baseObjectField_MetadataEntry = { key: "", value: "" };
export const ObjectField_MetadataEntry = {
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
            ...baseObjectField_MetadataEntry,
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
            ...baseObjectField_MetadataEntry,
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
            ...baseObjectField_MetadataEntry,
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
const baseObjectFieldArray = { label: "", type: 0, did: "" };
export const ObjectFieldArray = {
    encode(message, writer = Writer.create()) {
        if (message.label !== "") {
            writer.uint32(10).string(message.label);
        }
        if (message.type !== 0) {
            writer.uint32(16).int32(message.type);
        }
        if (message.did !== "") {
            writer.uint32(26).string(message.did);
        }
        for (const v of message.items) {
            ObjectField.encode(v, writer.uint32(34).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseObjectFieldArray };
        message.items = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.label = reader.string();
                    break;
                case 2:
                    message.type = reader.int32();
                    break;
                case 3:
                    message.did = reader.string();
                    break;
                case 4:
                    message.items.push(ObjectField.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseObjectFieldArray };
        message.items = [];
        if (object.label !== undefined && object.label !== null) {
            message.label = String(object.label);
        }
        else {
            message.label = "";
        }
        if (object.type !== undefined && object.type !== null) {
            message.type = objectFieldTypeFromJSON(object.type);
        }
        else {
            message.type = 0;
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = String(object.did);
        }
        else {
            message.did = "";
        }
        if (object.items !== undefined && object.items !== null) {
            for (const e of object.items) {
                message.items.push(ObjectField.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.label !== undefined && (obj.label = message.label);
        message.type !== undefined &&
            (obj.type = objectFieldTypeToJSON(message.type));
        message.did !== undefined && (obj.did = message.did);
        if (message.items) {
            obj.items = message.items.map((e) => e ? ObjectField.toJSON(e) : undefined);
        }
        else {
            obj.items = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseObjectFieldArray };
        message.items = [];
        if (object.label !== undefined && object.label !== null) {
            message.label = object.label;
        }
        else {
            message.label = "";
        }
        if (object.type !== undefined && object.type !== null) {
            message.type = object.type;
        }
        else {
            message.type = 0;
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = object.did;
        }
        else {
            message.did = "";
        }
        if (object.items !== undefined && object.items !== null) {
            for (const e of object.items) {
                message.items.push(ObjectField.fromPartial(e));
            }
        }
        return message;
    },
};
const baseObjectFieldText = { label: "", did: "", value: "" };
export const ObjectFieldText = {
    encode(message, writer = Writer.create()) {
        if (message.label !== "") {
            writer.uint32(10).string(message.label);
        }
        if (message.did !== "") {
            writer.uint32(18).string(message.did);
        }
        if (message.value !== "") {
            writer.uint32(26).string(message.value);
        }
        Object.entries(message.metadata).forEach(([key, value]) => {
            ObjectFieldText_MetadataEntry.encode({ key: key, value }, writer.uint32(34).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseObjectFieldText };
        message.metadata = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.label = reader.string();
                    break;
                case 2:
                    message.did = reader.string();
                    break;
                case 3:
                    message.value = reader.string();
                    break;
                case 4:
                    const entry4 = ObjectFieldText_MetadataEntry.decode(reader, reader.uint32());
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
        const message = { ...baseObjectFieldText };
        message.metadata = {};
        if (object.label !== undefined && object.label !== null) {
            message.label = String(object.label);
        }
        else {
            message.label = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = String(object.did);
        }
        else {
            message.did = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = String(object.value);
        }
        else {
            message.value = "";
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
        message.label !== undefined && (obj.label = message.label);
        message.did !== undefined && (obj.did = message.did);
        message.value !== undefined && (obj.value = message.value);
        obj.metadata = {};
        if (message.metadata) {
            Object.entries(message.metadata).forEach(([k, v]) => {
                obj.metadata[k] = v;
            });
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseObjectFieldText };
        message.metadata = {};
        if (object.label !== undefined && object.label !== null) {
            message.label = object.label;
        }
        else {
            message.label = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = object.did;
        }
        else {
            message.did = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = object.value;
        }
        else {
            message.value = "";
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
const baseObjectFieldText_MetadataEntry = { key: "", value: "" };
export const ObjectFieldText_MetadataEntry = {
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
            ...baseObjectFieldText_MetadataEntry,
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
            ...baseObjectFieldText_MetadataEntry,
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
            ...baseObjectFieldText_MetadataEntry,
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
const baseObjectFieldNumber = { label: "", did: "", value: 0 };
export const ObjectFieldNumber = {
    encode(message, writer = Writer.create()) {
        if (message.label !== "") {
            writer.uint32(10).string(message.label);
        }
        if (message.did !== "") {
            writer.uint32(18).string(message.did);
        }
        if (message.value !== 0) {
            writer.uint32(25).double(message.value);
        }
        Object.entries(message.metadata).forEach(([key, value]) => {
            ObjectFieldNumber_MetadataEntry.encode({ key: key, value }, writer.uint32(34).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseObjectFieldNumber };
        message.metadata = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.label = reader.string();
                    break;
                case 2:
                    message.did = reader.string();
                    break;
                case 3:
                    message.value = reader.double();
                    break;
                case 4:
                    const entry4 = ObjectFieldNumber_MetadataEntry.decode(reader, reader.uint32());
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
        const message = { ...baseObjectFieldNumber };
        message.metadata = {};
        if (object.label !== undefined && object.label !== null) {
            message.label = String(object.label);
        }
        else {
            message.label = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = String(object.did);
        }
        else {
            message.did = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = Number(object.value);
        }
        else {
            message.value = 0;
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
        message.label !== undefined && (obj.label = message.label);
        message.did !== undefined && (obj.did = message.did);
        message.value !== undefined && (obj.value = message.value);
        obj.metadata = {};
        if (message.metadata) {
            Object.entries(message.metadata).forEach(([k, v]) => {
                obj.metadata[k] = v;
            });
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseObjectFieldNumber };
        message.metadata = {};
        if (object.label !== undefined && object.label !== null) {
            message.label = object.label;
        }
        else {
            message.label = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = object.did;
        }
        else {
            message.did = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = object.value;
        }
        else {
            message.value = 0;
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
const baseObjectFieldNumber_MetadataEntry = { key: "", value: "" };
export const ObjectFieldNumber_MetadataEntry = {
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
            ...baseObjectFieldNumber_MetadataEntry,
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
            ...baseObjectFieldNumber_MetadataEntry,
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
            ...baseObjectFieldNumber_MetadataEntry,
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
const baseObjectFieldBool = { label: "", did: "", value: false };
export const ObjectFieldBool = {
    encode(message, writer = Writer.create()) {
        if (message.label !== "") {
            writer.uint32(10).string(message.label);
        }
        if (message.did !== "") {
            writer.uint32(18).string(message.did);
        }
        if (message.value === true) {
            writer.uint32(24).bool(message.value);
        }
        Object.entries(message.metadata).forEach(([key, value]) => {
            ObjectFieldBool_MetadataEntry.encode({ key: key, value }, writer.uint32(34).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseObjectFieldBool };
        message.metadata = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.label = reader.string();
                    break;
                case 2:
                    message.did = reader.string();
                    break;
                case 3:
                    message.value = reader.bool();
                    break;
                case 4:
                    const entry4 = ObjectFieldBool_MetadataEntry.decode(reader, reader.uint32());
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
        const message = { ...baseObjectFieldBool };
        message.metadata = {};
        if (object.label !== undefined && object.label !== null) {
            message.label = String(object.label);
        }
        else {
            message.label = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = String(object.did);
        }
        else {
            message.did = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = Boolean(object.value);
        }
        else {
            message.value = false;
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
        message.label !== undefined && (obj.label = message.label);
        message.did !== undefined && (obj.did = message.did);
        message.value !== undefined && (obj.value = message.value);
        obj.metadata = {};
        if (message.metadata) {
            Object.entries(message.metadata).forEach(([k, v]) => {
                obj.metadata[k] = v;
            });
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseObjectFieldBool };
        message.metadata = {};
        if (object.label !== undefined && object.label !== null) {
            message.label = object.label;
        }
        else {
            message.label = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = object.did;
        }
        else {
            message.did = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = object.value;
        }
        else {
            message.value = false;
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
const baseObjectFieldBool_MetadataEntry = { key: "", value: "" };
export const ObjectFieldBool_MetadataEntry = {
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
            ...baseObjectFieldBool_MetadataEntry,
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
            ...baseObjectFieldBool_MetadataEntry,
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
            ...baseObjectFieldBool_MetadataEntry,
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
const baseObjectFieldTime = { label: "", did: "", value: 0 };
export const ObjectFieldTime = {
    encode(message, writer = Writer.create()) {
        if (message.label !== "") {
            writer.uint32(10).string(message.label);
        }
        if (message.did !== "") {
            writer.uint32(18).string(message.did);
        }
        if (message.value !== 0) {
            writer.uint32(24).int64(message.value);
        }
        Object.entries(message.metadata).forEach(([key, value]) => {
            ObjectFieldTime_MetadataEntry.encode({ key: key, value }, writer.uint32(34).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseObjectFieldTime };
        message.metadata = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.label = reader.string();
                    break;
                case 2:
                    message.did = reader.string();
                    break;
                case 3:
                    message.value = longToNumber(reader.int64());
                    break;
                case 4:
                    const entry4 = ObjectFieldTime_MetadataEntry.decode(reader, reader.uint32());
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
        const message = { ...baseObjectFieldTime };
        message.metadata = {};
        if (object.label !== undefined && object.label !== null) {
            message.label = String(object.label);
        }
        else {
            message.label = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = String(object.did);
        }
        else {
            message.did = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = Number(object.value);
        }
        else {
            message.value = 0;
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
        message.label !== undefined && (obj.label = message.label);
        message.did !== undefined && (obj.did = message.did);
        message.value !== undefined && (obj.value = message.value);
        obj.metadata = {};
        if (message.metadata) {
            Object.entries(message.metadata).forEach(([k, v]) => {
                obj.metadata[k] = v;
            });
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseObjectFieldTime };
        message.metadata = {};
        if (object.label !== undefined && object.label !== null) {
            message.label = object.label;
        }
        else {
            message.label = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = object.did;
        }
        else {
            message.did = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = object.value;
        }
        else {
            message.value = 0;
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
const baseObjectFieldTime_MetadataEntry = { key: "", value: "" };
export const ObjectFieldTime_MetadataEntry = {
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
            ...baseObjectFieldTime_MetadataEntry,
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
            ...baseObjectFieldTime_MetadataEntry,
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
            ...baseObjectFieldTime_MetadataEntry,
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
const baseObjectFieldGeopoint = {
    label: "",
    type: 0,
    did: "",
    latitude: 0,
    longitude: 0,
};
export const ObjectFieldGeopoint = {
    encode(message, writer = Writer.create()) {
        if (message.label !== "") {
            writer.uint32(10).string(message.label);
        }
        if (message.type !== 0) {
            writer.uint32(16).int32(message.type);
        }
        if (message.did !== "") {
            writer.uint32(26).string(message.did);
        }
        if (message.latitude !== 0) {
            writer.uint32(33).double(message.latitude);
        }
        if (message.longitude !== 0) {
            writer.uint32(41).double(message.longitude);
        }
        Object.entries(message.metadata).forEach(([key, value]) => {
            ObjectFieldGeopoint_MetadataEntry.encode({ key: key, value }, writer.uint32(50).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseObjectFieldGeopoint };
        message.metadata = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.label = reader.string();
                    break;
                case 2:
                    message.type = reader.int32();
                    break;
                case 3:
                    message.did = reader.string();
                    break;
                case 4:
                    message.latitude = reader.double();
                    break;
                case 5:
                    message.longitude = reader.double();
                    break;
                case 6:
                    const entry6 = ObjectFieldGeopoint_MetadataEntry.decode(reader, reader.uint32());
                    if (entry6.value !== undefined) {
                        message.metadata[entry6.key] = entry6.value;
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
        const message = { ...baseObjectFieldGeopoint };
        message.metadata = {};
        if (object.label !== undefined && object.label !== null) {
            message.label = String(object.label);
        }
        else {
            message.label = "";
        }
        if (object.type !== undefined && object.type !== null) {
            message.type = objectFieldTypeFromJSON(object.type);
        }
        else {
            message.type = 0;
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = String(object.did);
        }
        else {
            message.did = "";
        }
        if (object.latitude !== undefined && object.latitude !== null) {
            message.latitude = Number(object.latitude);
        }
        else {
            message.latitude = 0;
        }
        if (object.longitude !== undefined && object.longitude !== null) {
            message.longitude = Number(object.longitude);
        }
        else {
            message.longitude = 0;
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
        message.label !== undefined && (obj.label = message.label);
        message.type !== undefined &&
            (obj.type = objectFieldTypeToJSON(message.type));
        message.did !== undefined && (obj.did = message.did);
        message.latitude !== undefined && (obj.latitude = message.latitude);
        message.longitude !== undefined && (obj.longitude = message.longitude);
        obj.metadata = {};
        if (message.metadata) {
            Object.entries(message.metadata).forEach(([k, v]) => {
                obj.metadata[k] = v;
            });
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseObjectFieldGeopoint };
        message.metadata = {};
        if (object.label !== undefined && object.label !== null) {
            message.label = object.label;
        }
        else {
            message.label = "";
        }
        if (object.type !== undefined && object.type !== null) {
            message.type = object.type;
        }
        else {
            message.type = 0;
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = object.did;
        }
        else {
            message.did = "";
        }
        if (object.latitude !== undefined && object.latitude !== null) {
            message.latitude = object.latitude;
        }
        else {
            message.latitude = 0;
        }
        if (object.longitude !== undefined && object.longitude !== null) {
            message.longitude = object.longitude;
        }
        else {
            message.longitude = 0;
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
const baseObjectFieldGeopoint_MetadataEntry = { key: "", value: "" };
export const ObjectFieldGeopoint_MetadataEntry = {
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
            ...baseObjectFieldGeopoint_MetadataEntry,
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
            ...baseObjectFieldGeopoint_MetadataEntry,
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
            ...baseObjectFieldGeopoint_MetadataEntry,
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
const baseObjectFieldBlob = { label: "", did: "" };
export const ObjectFieldBlob = {
    encode(message, writer = Writer.create()) {
        if (message.label !== "") {
            writer.uint32(10).string(message.label);
        }
        if (message.did !== "") {
            writer.uint32(18).string(message.did);
        }
        if (message.value.length !== 0) {
            writer.uint32(26).bytes(message.value);
        }
        Object.entries(message.metadata).forEach(([key, value]) => {
            ObjectFieldBlob_MetadataEntry.encode({ key: key, value }, writer.uint32(34).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseObjectFieldBlob };
        message.metadata = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.label = reader.string();
                    break;
                case 2:
                    message.did = reader.string();
                    break;
                case 3:
                    message.value = reader.bytes();
                    break;
                case 4:
                    const entry4 = ObjectFieldBlob_MetadataEntry.decode(reader, reader.uint32());
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
        const message = { ...baseObjectFieldBlob };
        message.metadata = {};
        if (object.label !== undefined && object.label !== null) {
            message.label = String(object.label);
        }
        else {
            message.label = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = String(object.did);
        }
        else {
            message.did = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = bytesFromBase64(object.value);
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
        message.label !== undefined && (obj.label = message.label);
        message.did !== undefined && (obj.did = message.did);
        message.value !== undefined &&
            (obj.value = base64FromBytes(message.value !== undefined ? message.value : new Uint8Array()));
        obj.metadata = {};
        if (message.metadata) {
            Object.entries(message.metadata).forEach(([k, v]) => {
                obj.metadata[k] = v;
            });
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseObjectFieldBlob };
        message.metadata = {};
        if (object.label !== undefined && object.label !== null) {
            message.label = object.label;
        }
        else {
            message.label = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = object.did;
        }
        else {
            message.did = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = object.value;
        }
        else {
            message.value = new Uint8Array();
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
const baseObjectFieldBlob_MetadataEntry = { key: "", value: "" };
export const ObjectFieldBlob_MetadataEntry = {
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
            ...baseObjectFieldBlob_MetadataEntry,
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
            ...baseObjectFieldBlob_MetadataEntry,
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
            ...baseObjectFieldBlob_MetadataEntry,
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
const baseObjectFieldBlockchainAddress = {
    label: "",
    did: "",
    value: "",
};
export const ObjectFieldBlockchainAddress = {
    encode(message, writer = Writer.create()) {
        if (message.label !== "") {
            writer.uint32(10).string(message.label);
        }
        if (message.did !== "") {
            writer.uint32(18).string(message.did);
        }
        if (message.value !== "") {
            writer.uint32(26).string(message.value);
        }
        Object.entries(message.metadata).forEach(([key, value]) => {
            ObjectFieldBlockchainAddress_MetadataEntry.encode({ key: key, value }, writer.uint32(34).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseObjectFieldBlockchainAddress,
        };
        message.metadata = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.label = reader.string();
                    break;
                case 2:
                    message.did = reader.string();
                    break;
                case 3:
                    message.value = reader.string();
                    break;
                case 4:
                    const entry4 = ObjectFieldBlockchainAddress_MetadataEntry.decode(reader, reader.uint32());
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
        const message = {
            ...baseObjectFieldBlockchainAddress,
        };
        message.metadata = {};
        if (object.label !== undefined && object.label !== null) {
            message.label = String(object.label);
        }
        else {
            message.label = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = String(object.did);
        }
        else {
            message.did = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = String(object.value);
        }
        else {
            message.value = "";
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
        message.label !== undefined && (obj.label = message.label);
        message.did !== undefined && (obj.did = message.did);
        message.value !== undefined && (obj.value = message.value);
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
            ...baseObjectFieldBlockchainAddress,
        };
        message.metadata = {};
        if (object.label !== undefined && object.label !== null) {
            message.label = object.label;
        }
        else {
            message.label = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = object.did;
        }
        else {
            message.did = "";
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = object.value;
        }
        else {
            message.value = "";
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
const baseObjectFieldBlockchainAddress_MetadataEntry = {
    key: "",
    value: "",
};
export const ObjectFieldBlockchainAddress_MetadataEntry = {
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
            ...baseObjectFieldBlockchainAddress_MetadataEntry,
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
            ...baseObjectFieldBlockchainAddress_MetadataEntry,
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
            ...baseObjectFieldBlockchainAddress_MetadataEntry,
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
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
const atob = globalThis.atob ||
    ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64) {
    const bin = atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
        arr[i] = bin.charCodeAt(i);
    }
    return arr;
}
const btoa = globalThis.btoa ||
    ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr) {
    const bin = [];
    for (let i = 0; i < arr.byteLength; ++i) {
        bin.push(String.fromCharCode(arr[i]));
    }
    return btoa(bin.join(""));
}
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
