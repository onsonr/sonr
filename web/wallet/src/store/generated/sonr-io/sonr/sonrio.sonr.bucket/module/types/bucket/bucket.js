/* eslint-disable */
import { Did } from "../registry/did";
import { ObjectDoc } from "../object/object";
import { Peer } from "../registry/peer";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "sonrio.sonr.bucket";
/** BucketType is the type of a bucket. */
export var BucketType;
(function (BucketType) {
    /** BUCKET_TYPE_UNSPECIFIED - BucketTypeUnspecified is the default value. */
    BucketType[BucketType["BUCKET_TYPE_UNSPECIFIED"] = 0] = "BUCKET_TYPE_UNSPECIFIED";
    /** BUCKET_TYPE_APP - BucketTypeApp is an App specific bucket. For Assets regarding the service. */
    BucketType[BucketType["BUCKET_TYPE_APP"] = 1] = "BUCKET_TYPE_APP";
    /**
     * BUCKET_TYPE_USER - BucketTypeUser is a User specific bucket. For any remote user data that is required
     * to be stored in the Network.
     */
    BucketType[BucketType["BUCKET_TYPE_USER"] = 2] = "BUCKET_TYPE_USER";
    BucketType[BucketType["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(BucketType || (BucketType = {}));
export function bucketTypeFromJSON(object) {
    switch (object) {
        case 0:
        case "BUCKET_TYPE_UNSPECIFIED":
            return BucketType.BUCKET_TYPE_UNSPECIFIED;
        case 1:
        case "BUCKET_TYPE_APP":
            return BucketType.BUCKET_TYPE_APP;
        case 2:
        case "BUCKET_TYPE_USER":
            return BucketType.BUCKET_TYPE_USER;
        case -1:
        case "UNRECOGNIZED":
        default:
            return BucketType.UNRECOGNIZED;
    }
}
export function bucketTypeToJSON(object) {
    switch (object) {
        case BucketType.BUCKET_TYPE_UNSPECIFIED:
            return "BUCKET_TYPE_UNSPECIFIED";
        case BucketType.BUCKET_TYPE_APP:
            return "BUCKET_TYPE_APP";
        case BucketType.BUCKET_TYPE_USER:
            return "BUCKET_TYPE_USER";
        default:
            return "UNKNOWN";
    }
}
/** EventType is the type of event being performed on a Bucket. */
export var BucketEventType;
(function (BucketEventType) {
    /** Bucket_EVENT_TYPE_UNSPECIFIED - EventTypeUnspecified is the default value. */
    BucketEventType[BucketEventType["Bucket_EVENT_TYPE_UNSPECIFIED"] = 0] = "Bucket_EVENT_TYPE_UNSPECIFIED";
    /** Bucket_EVENT_TYPE_GET - EventTypeGet is a get event being performed on a Bucket record. */
    BucketEventType[BucketEventType["Bucket_EVENT_TYPE_GET"] = 1] = "Bucket_EVENT_TYPE_GET";
    /** Bucket_EVENT_TYPE_SET - EventTypeSet is a set event on the Bucket store. */
    BucketEventType[BucketEventType["Bucket_EVENT_TYPE_SET"] = 2] = "Bucket_EVENT_TYPE_SET";
    /** Bucket_EVENT_TYPE_DELETE - EventTypeDelete is a delete event on the Bucket store. */
    BucketEventType[BucketEventType["Bucket_EVENT_TYPE_DELETE"] = 3] = "Bucket_EVENT_TYPE_DELETE";
    BucketEventType[BucketEventType["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(BucketEventType || (BucketEventType = {}));
export function bucketEventTypeFromJSON(object) {
    switch (object) {
        case 0:
        case "Bucket_EVENT_TYPE_UNSPECIFIED":
            return BucketEventType.Bucket_EVENT_TYPE_UNSPECIFIED;
        case 1:
        case "Bucket_EVENT_TYPE_GET":
            return BucketEventType.Bucket_EVENT_TYPE_GET;
        case 2:
        case "Bucket_EVENT_TYPE_SET":
            return BucketEventType.Bucket_EVENT_TYPE_SET;
        case 3:
        case "Bucket_EVENT_TYPE_DELETE":
            return BucketEventType.Bucket_EVENT_TYPE_DELETE;
        case -1:
        case "UNRECOGNIZED":
        default:
            return BucketEventType.UNRECOGNIZED;
    }
}
export function bucketEventTypeToJSON(object) {
    switch (object) {
        case BucketEventType.Bucket_EVENT_TYPE_UNSPECIFIED:
            return "Bucket_EVENT_TYPE_UNSPECIFIED";
        case BucketEventType.Bucket_EVENT_TYPE_GET:
            return "Bucket_EVENT_TYPE_GET";
        case BucketEventType.Bucket_EVENT_TYPE_SET:
            return "Bucket_EVENT_TYPE_SET";
        case BucketEventType.Bucket_EVENT_TYPE_DELETE:
            return "Bucket_EVENT_TYPE_DELETE";
        default:
            return "UNKNOWN";
    }
}
const baseBucket = { label: "", description: "", type: 0 };
export const Bucket = {
    encode(message, writer = Writer.create()) {
        if (message.label !== "") {
            writer.uint32(10).string(message.label);
        }
        if (message.description !== "") {
            writer.uint32(18).string(message.description);
        }
        if (message.type !== 0) {
            writer.uint32(24).int32(message.type);
        }
        if (message.did !== undefined) {
            Did.encode(message.did, writer.uint32(34).fork()).ldelim();
        }
        for (const v of message.objects) {
            ObjectDoc.encode(v, writer.uint32(42).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseBucket };
        message.objects = [];
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
                    message.type = reader.int32();
                    break;
                case 4:
                    message.did = Did.decode(reader, reader.uint32());
                    break;
                case 5:
                    message.objects.push(ObjectDoc.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseBucket };
        message.objects = [];
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
        if (object.type !== undefined && object.type !== null) {
            message.type = bucketTypeFromJSON(object.type);
        }
        else {
            message.type = 0;
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = Did.fromJSON(object.did);
        }
        else {
            message.did = undefined;
        }
        if (object.objects !== undefined && object.objects !== null) {
            for (const e of object.objects) {
                message.objects.push(ObjectDoc.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.label !== undefined && (obj.label = message.label);
        message.description !== undefined &&
            (obj.description = message.description);
        message.type !== undefined && (obj.type = bucketTypeToJSON(message.type));
        message.did !== undefined &&
            (obj.did = message.did ? Did.toJSON(message.did) : undefined);
        if (message.objects) {
            obj.objects = message.objects.map((e) => e ? ObjectDoc.toJSON(e) : undefined);
        }
        else {
            obj.objects = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseBucket };
        message.objects = [];
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
        if (object.type !== undefined && object.type !== null) {
            message.type = object.type;
        }
        else {
            message.type = 0;
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = Did.fromPartial(object.did);
        }
        else {
            message.did = undefined;
        }
        if (object.objects !== undefined && object.objects !== null) {
            for (const e of object.objects) {
                message.objects.push(ObjectDoc.fromPartial(e));
            }
        }
        return message;
    },
};
const baseBucketEvent = { type: 0 };
export const BucketEvent = {
    encode(message, writer = Writer.create()) {
        if (message.peer !== undefined) {
            Peer.encode(message.peer, writer.uint32(10).fork()).ldelim();
        }
        if (message.type !== 0) {
            writer.uint32(16).int32(message.type);
        }
        if (message.object !== undefined) {
            ObjectDoc.encode(message.object, writer.uint32(26).fork()).ldelim();
        }
        Object.entries(message.metadata).forEach(([key, value]) => {
            BucketEvent_MetadataEntry.encode({ key: key, value }, writer.uint32(34).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseBucketEvent };
        message.metadata = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.peer = Peer.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.type = reader.int32();
                    break;
                case 3:
                    message.object = ObjectDoc.decode(reader, reader.uint32());
                    break;
                case 4:
                    const entry4 = BucketEvent_MetadataEntry.decode(reader, reader.uint32());
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
        const message = { ...baseBucketEvent };
        message.metadata = {};
        if (object.peer !== undefined && object.peer !== null) {
            message.peer = Peer.fromJSON(object.peer);
        }
        else {
            message.peer = undefined;
        }
        if (object.type !== undefined && object.type !== null) {
            message.type = bucketEventTypeFromJSON(object.type);
        }
        else {
            message.type = 0;
        }
        if (object.object !== undefined && object.object !== null) {
            message.object = ObjectDoc.fromJSON(object.object);
        }
        else {
            message.object = undefined;
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
        message.type !== undefined &&
            (obj.type = bucketEventTypeToJSON(message.type));
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
    fromPartial(object) {
        const message = { ...baseBucketEvent };
        message.metadata = {};
        if (object.peer !== undefined && object.peer !== null) {
            message.peer = Peer.fromPartial(object.peer);
        }
        else {
            message.peer = undefined;
        }
        if (object.type !== undefined && object.type !== null) {
            message.type = object.type;
        }
        else {
            message.type = 0;
        }
        if (object.object !== undefined && object.object !== null) {
            message.object = ObjectDoc.fromPartial(object.object);
        }
        else {
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
const baseBucketEvent_MetadataEntry = { key: "", value: "" };
export const BucketEvent_MetadataEntry = {
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
            ...baseBucketEvent_MetadataEntry,
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
            ...baseBucketEvent_MetadataEntry,
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
            ...baseBucketEvent_MetadataEntry,
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
