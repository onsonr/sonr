/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
export const protobufPackage = "sonrio.sonr.blob";
const baseMsgUploadBlob = {
    creator: "",
    label: "",
    path: "",
    refDid: "",
    size: 0,
    lastModified: 0,
};
export const MsgUploadBlob = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUploadBlob };
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
    fromJSON(object) {
        const message = { ...baseMsgUploadBlob };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.label !== undefined && object.label !== null) {
            message.label = String(object.label);
        }
        else {
            message.label = "";
        }
        if (object.path !== undefined && object.path !== null) {
            message.path = String(object.path);
        }
        else {
            message.path = "";
        }
        if (object.refDid !== undefined && object.refDid !== null) {
            message.refDid = String(object.refDid);
        }
        else {
            message.refDid = "";
        }
        if (object.size !== undefined && object.size !== null) {
            message.size = Number(object.size);
        }
        else {
            message.size = 0;
        }
        if (object.lastModified !== undefined && object.lastModified !== null) {
            message.lastModified = Number(object.lastModified);
        }
        else {
            message.lastModified = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.label !== undefined && (obj.label = message.label);
        message.path !== undefined && (obj.path = message.path);
        message.refDid !== undefined && (obj.refDid = message.refDid);
        message.size !== undefined && (obj.size = message.size);
        message.lastModified !== undefined &&
            (obj.lastModified = message.lastModified);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUploadBlob };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.label !== undefined && object.label !== null) {
            message.label = object.label;
        }
        else {
            message.label = "";
        }
        if (object.path !== undefined && object.path !== null) {
            message.path = object.path;
        }
        else {
            message.path = "";
        }
        if (object.refDid !== undefined && object.refDid !== null) {
            message.refDid = object.refDid;
        }
        else {
            message.refDid = "";
        }
        if (object.size !== undefined && object.size !== null) {
            message.size = object.size;
        }
        else {
            message.size = 0;
        }
        if (object.lastModified !== undefined && object.lastModified !== null) {
            message.lastModified = object.lastModified;
        }
        else {
            message.lastModified = 0;
        }
        return message;
    },
};
const baseMsgUploadBlobResponse = {};
export const MsgUploadBlobResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUploadBlobResponse };
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
        const message = { ...baseMsgUploadBlobResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgUploadBlobResponse };
        return message;
    },
};
const baseMsgDownloadBlob = {
    creator: "",
    did: "",
    path: "",
    timeout: 0,
};
export const MsgDownloadBlob = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDownloadBlob };
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
    fromJSON(object) {
        const message = { ...baseMsgDownloadBlob };
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
        if (object.path !== undefined && object.path !== null) {
            message.path = String(object.path);
        }
        else {
            message.path = "";
        }
        if (object.timeout !== undefined && object.timeout !== null) {
            message.timeout = Number(object.timeout);
        }
        else {
            message.timeout = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.did !== undefined && (obj.did = message.did);
        message.path !== undefined && (obj.path = message.path);
        message.timeout !== undefined && (obj.timeout = message.timeout);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgDownloadBlob };
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
        if (object.path !== undefined && object.path !== null) {
            message.path = object.path;
        }
        else {
            message.path = "";
        }
        if (object.timeout !== undefined && object.timeout !== null) {
            message.timeout = object.timeout;
        }
        else {
            message.timeout = 0;
        }
        return message;
    },
};
const baseMsgDownloadBlobResponse = {};
export const MsgDownloadBlobResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgDownloadBlobResponse,
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
            ...baseMsgDownloadBlobResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgDownloadBlobResponse,
        };
        return message;
    },
};
const baseMsgSyncBlob = { creator: "", did: "", path: "", timeout: 0 };
export const MsgSyncBlob = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSyncBlob };
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
    fromJSON(object) {
        const message = { ...baseMsgSyncBlob };
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
        if (object.path !== undefined && object.path !== null) {
            message.path = String(object.path);
        }
        else {
            message.path = "";
        }
        if (object.timeout !== undefined && object.timeout !== null) {
            message.timeout = Number(object.timeout);
        }
        else {
            message.timeout = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.did !== undefined && (obj.did = message.did);
        message.path !== undefined && (obj.path = message.path);
        message.timeout !== undefined && (obj.timeout = message.timeout);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSyncBlob };
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
        if (object.path !== undefined && object.path !== null) {
            message.path = object.path;
        }
        else {
            message.path = "";
        }
        if (object.timeout !== undefined && object.timeout !== null) {
            message.timeout = object.timeout;
        }
        else {
            message.timeout = 0;
        }
        return message;
    },
};
const baseMsgSyncBlobResponse = {};
export const MsgSyncBlobResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSyncBlobResponse };
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
        const message = { ...baseMsgSyncBlobResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgSyncBlobResponse };
        return message;
    },
};
const baseMsgDeleteBlob = { creator: "", did: "", publicKey: "" };
export const MsgDeleteBlob = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDeleteBlob };
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
    fromJSON(object) {
        const message = { ...baseMsgDeleteBlob };
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
        message.did !== undefined && (obj.did = message.did);
        message.publicKey !== undefined && (obj.publicKey = message.publicKey);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgDeleteBlob };
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
        if (object.publicKey !== undefined && object.publicKey !== null) {
            message.publicKey = object.publicKey;
        }
        else {
            message.publicKey = "";
        }
        return message;
    },
};
const baseMsgDeleteBlobResponse = {};
export const MsgDeleteBlobResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDeleteBlobResponse };
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
        const message = { ...baseMsgDeleteBlobResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgDeleteBlobResponse };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    UploadBlob(request) {
        const data = MsgUploadBlob.encode(request).finish();
        const promise = this.rpc.request("sonrio.sonr.blob.Msg", "UploadBlob", data);
        return promise.then((data) => MsgUploadBlobResponse.decode(new Reader(data)));
    }
    DownloadBlob(request) {
        const data = MsgDownloadBlob.encode(request).finish();
        const promise = this.rpc.request("sonrio.sonr.blob.Msg", "DownloadBlob", data);
        return promise.then((data) => MsgDownloadBlobResponse.decode(new Reader(data)));
    }
    SyncBlob(request) {
        const data = MsgSyncBlob.encode(request).finish();
        const promise = this.rpc.request("sonrio.sonr.blob.Msg", "SyncBlob", data);
        return promise.then((data) => MsgSyncBlobResponse.decode(new Reader(data)));
    }
    DeleteBlob(request) {
        const data = MsgDeleteBlob.encode(request).finish();
        const promise = this.rpc.request("sonrio.sonr.blob.Msg", "DeleteBlob", data);
        return promise.then((data) => MsgDeleteBlobResponse.decode(new Reader(data)));
    }
}
