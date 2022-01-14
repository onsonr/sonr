"use strict";
exports.__esModule = true;
exports.ProfileList = exports.Profile = exports.Peer_Device = exports.Peer = exports.MIME = exports.Metadata = exports.Location_Placemark = exports.Location = exports.peer_StatusToNumber = exports.peer_StatusToJSON = exports.peer_StatusFromJSON = exports.Peer_Status = exports.mIME_TypeToNumber = exports.mIME_TypeToJSON = exports.mIME_TypeFromJSON = exports.MIME_Type = exports.connectionToNumber = exports.connectionToJSON = exports.connectionFromJSON = exports.Connection = exports.protobufPackage = void 0;
/* eslint-disable */
var long_1 = require("long");
var minimal_1 = require("protobufjs/minimal");
exports.protobufPackage = "common";
/** Internet Connection Type */
var Connection;
(function (Connection) {
    Connection["CONNECTION_UNSPECIFIED"] = "CONNECTION_UNSPECIFIED";
    /** CONNECTION_WIFI - ConnectionWifi is used for WiFi connections. */
    Connection["CONNECTION_WIFI"] = "CONNECTION_WIFI";
    /** CONNECTION_ETHERNET - ConnectionEthernet is used for Ethernet connections. */
    Connection["CONNECTION_ETHERNET"] = "CONNECTION_ETHERNET";
    /** CONNECTION_MOBILE - ConnectionMobile is used for mobile connections. */
    Connection["CONNECTION_MOBILE"] = "CONNECTION_MOBILE";
    /** CONNECTION_OFFLINE - CONNECTION_OFFLINE */
    Connection["CONNECTION_OFFLINE"] = "CONNECTION_OFFLINE";
    Connection["UNRECOGNIZED"] = "UNRECOGNIZED";
})(Connection = exports.Connection || (exports.Connection = {}));
function connectionFromJSON(object) {
    switch (object) {
        case 0:
        case "CONNECTION_UNSPECIFIED":
            return Connection.CONNECTION_UNSPECIFIED;
        case 1:
        case "CONNECTION_WIFI":
            return Connection.CONNECTION_WIFI;
        case 2:
        case "CONNECTION_ETHERNET":
            return Connection.CONNECTION_ETHERNET;
        case 3:
        case "CONNECTION_MOBILE":
            return Connection.CONNECTION_MOBILE;
        case 4:
        case "CONNECTION_OFFLINE":
            return Connection.CONNECTION_OFFLINE;
        case -1:
        case "UNRECOGNIZED":
        default:
            return Connection.UNRECOGNIZED;
    }
}
exports.connectionFromJSON = connectionFromJSON;
function connectionToJSON(object) {
    switch (object) {
        case Connection.CONNECTION_UNSPECIFIED:
            return "CONNECTION_UNSPECIFIED";
        case Connection.CONNECTION_WIFI:
            return "CONNECTION_WIFI";
        case Connection.CONNECTION_ETHERNET:
            return "CONNECTION_ETHERNET";
        case Connection.CONNECTION_MOBILE:
            return "CONNECTION_MOBILE";
        case Connection.CONNECTION_OFFLINE:
            return "CONNECTION_OFFLINE";
        default:
            return "UNKNOWN";
    }
}
exports.connectionToJSON = connectionToJSON;
function connectionToNumber(object) {
    switch (object) {
        case Connection.CONNECTION_UNSPECIFIED:
            return 0;
        case Connection.CONNECTION_WIFI:
            return 1;
        case Connection.CONNECTION_ETHERNET:
            return 2;
        case Connection.CONNECTION_MOBILE:
            return 3;
        case Connection.CONNECTION_OFFLINE:
            return 4;
        default:
            return 0;
    }
}
exports.connectionToNumber = connectionToNumber;
/** File Content Type */
var MIME_Type;
(function (MIME_Type) {
    /** TYPE_UNSPECIFIED - Other File Type - If cannot derive from Subtype */
    MIME_Type["TYPE_UNSPECIFIED"] = "TYPE_UNSPECIFIED";
    /** TYPE_AUDIO - Sound, Audio Files */
    MIME_Type["TYPE_AUDIO"] = "TYPE_AUDIO";
    /** TYPE_DOCUMENT - Document Files - PDF, Word, Excel, etc. */
    MIME_Type["TYPE_DOCUMENT"] = "TYPE_DOCUMENT";
    /** TYPE_IMAGE - Image Files */
    MIME_Type["TYPE_IMAGE"] = "TYPE_IMAGE";
    /** TYPE_TEXT - Text Based Files */
    MIME_Type["TYPE_TEXT"] = "TYPE_TEXT";
    /** TYPE_VIDEO - Video Files */
    MIME_Type["TYPE_VIDEO"] = "TYPE_VIDEO";
    /** TYPE_URL - URL Links */
    MIME_Type["TYPE_URL"] = "TYPE_URL";
    MIME_Type["UNRECOGNIZED"] = "UNRECOGNIZED";
})(MIME_Type = exports.MIME_Type || (exports.MIME_Type = {}));
function mIME_TypeFromJSON(object) {
    switch (object) {
        case 0:
        case "TYPE_UNSPECIFIED":
            return MIME_Type.TYPE_UNSPECIFIED;
        case 1:
        case "TYPE_AUDIO":
            return MIME_Type.TYPE_AUDIO;
        case 2:
        case "TYPE_DOCUMENT":
            return MIME_Type.TYPE_DOCUMENT;
        case 3:
        case "TYPE_IMAGE":
            return MIME_Type.TYPE_IMAGE;
        case 4:
        case "TYPE_TEXT":
            return MIME_Type.TYPE_TEXT;
        case 5:
        case "TYPE_VIDEO":
            return MIME_Type.TYPE_VIDEO;
        case 6:
        case "TYPE_URL":
            return MIME_Type.TYPE_URL;
        case -1:
        case "UNRECOGNIZED":
        default:
            return MIME_Type.UNRECOGNIZED;
    }
}
exports.mIME_TypeFromJSON = mIME_TypeFromJSON;
function mIME_TypeToJSON(object) {
    switch (object) {
        case MIME_Type.TYPE_UNSPECIFIED:
            return "TYPE_UNSPECIFIED";
        case MIME_Type.TYPE_AUDIO:
            return "TYPE_AUDIO";
        case MIME_Type.TYPE_DOCUMENT:
            return "TYPE_DOCUMENT";
        case MIME_Type.TYPE_IMAGE:
            return "TYPE_IMAGE";
        case MIME_Type.TYPE_TEXT:
            return "TYPE_TEXT";
        case MIME_Type.TYPE_VIDEO:
            return "TYPE_VIDEO";
        case MIME_Type.TYPE_URL:
            return "TYPE_URL";
        default:
            return "UNKNOWN";
    }
}
exports.mIME_TypeToJSON = mIME_TypeToJSON;
function mIME_TypeToNumber(object) {
    switch (object) {
        case MIME_Type.TYPE_UNSPECIFIED:
            return 0;
        case MIME_Type.TYPE_AUDIO:
            return 1;
        case MIME_Type.TYPE_DOCUMENT:
            return 2;
        case MIME_Type.TYPE_IMAGE:
            return 3;
        case MIME_Type.TYPE_TEXT:
            return 4;
        case MIME_Type.TYPE_VIDEO:
            return 5;
        case MIME_Type.TYPE_URL:
            return 6;
        default:
            return 0;
    }
}
exports.mIME_TypeToNumber = mIME_TypeToNumber;
/** Peers Active Status */
var Peer_Status;
(function (Peer_Status) {
    /** STATUS_UNSPECIFIED - Offline - Not Online or Not a Full Node */
    Peer_Status["STATUS_UNSPECIFIED"] = "STATUS_UNSPECIFIED";
    /** STATUS_ONLINE - Online - Full Node Available */
    Peer_Status["STATUS_ONLINE"] = "STATUS_ONLINE";
    /** STATUS_AWAY - Away - Not Online, but has a full node */
    Peer_Status["STATUS_AWAY"] = "STATUS_AWAY";
    /** STATUS_BUSY - Busy - Online, but busy with Transfer */
    Peer_Status["STATUS_BUSY"] = "STATUS_BUSY";
    Peer_Status["UNRECOGNIZED"] = "UNRECOGNIZED";
})(Peer_Status = exports.Peer_Status || (exports.Peer_Status = {}));
function peer_StatusFromJSON(object) {
    switch (object) {
        case 0:
        case "STATUS_UNSPECIFIED":
            return Peer_Status.STATUS_UNSPECIFIED;
        case 1:
        case "STATUS_ONLINE":
            return Peer_Status.STATUS_ONLINE;
        case 2:
        case "STATUS_AWAY":
            return Peer_Status.STATUS_AWAY;
        case 3:
        case "STATUS_BUSY":
            return Peer_Status.STATUS_BUSY;
        case -1:
        case "UNRECOGNIZED":
        default:
            return Peer_Status.UNRECOGNIZED;
    }
}
exports.peer_StatusFromJSON = peer_StatusFromJSON;
function peer_StatusToJSON(object) {
    switch (object) {
        case Peer_Status.STATUS_UNSPECIFIED:
            return "STATUS_UNSPECIFIED";
        case Peer_Status.STATUS_ONLINE:
            return "STATUS_ONLINE";
        case Peer_Status.STATUS_AWAY:
            return "STATUS_AWAY";
        case Peer_Status.STATUS_BUSY:
            return "STATUS_BUSY";
        default:
            return "UNKNOWN";
    }
}
exports.peer_StatusToJSON = peer_StatusToJSON;
function peer_StatusToNumber(object) {
    switch (object) {
        case Peer_Status.STATUS_UNSPECIFIED:
            return 0;
        case Peer_Status.STATUS_ONLINE:
            return 1;
        case Peer_Status.STATUS_AWAY:
            return 2;
        case Peer_Status.STATUS_BUSY:
            return 3;
        default:
            return 0;
    }
}
exports.peer_StatusToNumber = peer_StatusToNumber;
function createBaseLocation() {
    return { latitude: 0, longitude: 0, placemark: undefined, lastModified: 0 };
}
exports.Location = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        if (message.latitude !== 0) {
            writer.uint32(9).double(message.latitude);
        }
        if (message.longitude !== 0) {
            writer.uint32(17).double(message.longitude);
        }
        if (message.placemark !== undefined) {
            exports.Location_Placemark.encode(message.placemark, writer.uint32(26).fork()).ldelim();
        }
        if (message.lastModified !== 0) {
            writer.uint32(32).int64(message.lastModified);
        }
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBaseLocation();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.latitude = reader.double();
                    break;
                case 2:
                    message.longitude = reader.double();
                    break;
                case 3:
                    message.placemark = exports.Location_Placemark.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.lastModified = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON: function (object) {
        return {
            latitude: isSet(object.latitude) ? Number(object.latitude) : 0,
            longitude: isSet(object.longitude) ? Number(object.longitude) : 0,
            placemark: isSet(object.placemark)
                ? exports.Location_Placemark.fromJSON(object.placemark)
                : undefined,
            lastModified: isSet(object.lastModified)
                ? Number(object.lastModified)
                : 0
        };
    },
    toJSON: function (message) {
        var obj = {};
        message.latitude !== undefined && (obj.latitude = message.latitude);
        message.longitude !== undefined && (obj.longitude = message.longitude);
        message.placemark !== undefined &&
            (obj.placemark = message.placemark
                ? exports.Location_Placemark.toJSON(message.placemark)
                : undefined);
        message.lastModified !== undefined &&
            (obj.lastModified = Math.round(message.lastModified));
        return obj;
    },
    fromPartial: function (object) {
        var _a, _b, _c;
        var message = createBaseLocation();
        message.latitude = (_a = object.latitude) !== null && _a !== void 0 ? _a : 0;
        message.longitude = (_b = object.longitude) !== null && _b !== void 0 ? _b : 0;
        message.placemark =
            object.placemark !== undefined && object.placemark !== null
                ? exports.Location_Placemark.fromPartial(object.placemark)
                : undefined;
        message.lastModified = (_c = object.lastModified) !== null && _c !== void 0 ? _c : 0;
        return message;
    }
};
function createBaseLocation_Placemark() {
    return {
        name: "",
        street: "",
        isoCountryCode: "",
        country: "",
        postalCode: "",
        administrativeArea: "",
        subAdministrativeArea: "",
        locality: "",
        subLocality: "",
        thoroughfare: "",
        subThoroughfare: ""
    };
}
exports.Location_Placemark = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        if (message.name !== "") {
            writer.uint32(10).string(message.name);
        }
        if (message.street !== "") {
            writer.uint32(18).string(message.street);
        }
        if (message.isoCountryCode !== "") {
            writer.uint32(26).string(message.isoCountryCode);
        }
        if (message.country !== "") {
            writer.uint32(34).string(message.country);
        }
        if (message.postalCode !== "") {
            writer.uint32(42).string(message.postalCode);
        }
        if (message.administrativeArea !== "") {
            writer.uint32(50).string(message.administrativeArea);
        }
        if (message.subAdministrativeArea !== "") {
            writer.uint32(58).string(message.subAdministrativeArea);
        }
        if (message.locality !== "") {
            writer.uint32(66).string(message.locality);
        }
        if (message.subLocality !== "") {
            writer.uint32(74).string(message.subLocality);
        }
        if (message.thoroughfare !== "") {
            writer.uint32(82).string(message.thoroughfare);
        }
        if (message.subThoroughfare !== "") {
            writer.uint32(90).string(message.subThoroughfare);
        }
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBaseLocation_Placemark();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.name = reader.string();
                    break;
                case 2:
                    message.street = reader.string();
                    break;
                case 3:
                    message.isoCountryCode = reader.string();
                    break;
                case 4:
                    message.country = reader.string();
                    break;
                case 5:
                    message.postalCode = reader.string();
                    break;
                case 6:
                    message.administrativeArea = reader.string();
                    break;
                case 7:
                    message.subAdministrativeArea = reader.string();
                    break;
                case 8:
                    message.locality = reader.string();
                    break;
                case 9:
                    message.subLocality = reader.string();
                    break;
                case 10:
                    message.thoroughfare = reader.string();
                    break;
                case 11:
                    message.subThoroughfare = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON: function (object) {
        return {
            name: isSet(object.name) ? String(object.name) : "",
            street: isSet(object.street) ? String(object.street) : "",
            isoCountryCode: isSet(object.isoCountryCode)
                ? String(object.isoCountryCode)
                : "",
            country: isSet(object.country) ? String(object.country) : "",
            postalCode: isSet(object.postalCode) ? String(object.postalCode) : "",
            administrativeArea: isSet(object.administrativeArea)
                ? String(object.administrativeArea)
                : "",
            subAdministrativeArea: isSet(object.subAdministrativeArea)
                ? String(object.subAdministrativeArea)
                : "",
            locality: isSet(object.locality) ? String(object.locality) : "",
            subLocality: isSet(object.subLocality) ? String(object.subLocality) : "",
            thoroughfare: isSet(object.thoroughfare)
                ? String(object.thoroughfare)
                : "",
            subThoroughfare: isSet(object.subThoroughfare)
                ? String(object.subThoroughfare)
                : ""
        };
    },
    toJSON: function (message) {
        var obj = {};
        message.name !== undefined && (obj.name = message.name);
        message.street !== undefined && (obj.street = message.street);
        message.isoCountryCode !== undefined &&
            (obj.isoCountryCode = message.isoCountryCode);
        message.country !== undefined && (obj.country = message.country);
        message.postalCode !== undefined && (obj.postalCode = message.postalCode);
        message.administrativeArea !== undefined &&
            (obj.administrativeArea = message.administrativeArea);
        message.subAdministrativeArea !== undefined &&
            (obj.subAdministrativeArea = message.subAdministrativeArea);
        message.locality !== undefined && (obj.locality = message.locality);
        message.subLocality !== undefined &&
            (obj.subLocality = message.subLocality);
        message.thoroughfare !== undefined &&
            (obj.thoroughfare = message.thoroughfare);
        message.subThoroughfare !== undefined &&
            (obj.subThoroughfare = message.subThoroughfare);
        return obj;
    },
    fromPartial: function (object) {
        var _a, _b, _c, _d, _e, _f, _g, _h, _j, _k, _l;
        var message = createBaseLocation_Placemark();
        message.name = (_a = object.name) !== null && _a !== void 0 ? _a : "";
        message.street = (_b = object.street) !== null && _b !== void 0 ? _b : "";
        message.isoCountryCode = (_c = object.isoCountryCode) !== null && _c !== void 0 ? _c : "";
        message.country = (_d = object.country) !== null && _d !== void 0 ? _d : "";
        message.postalCode = (_e = object.postalCode) !== null && _e !== void 0 ? _e : "";
        message.administrativeArea = (_f = object.administrativeArea) !== null && _f !== void 0 ? _f : "";
        message.subAdministrativeArea = (_g = object.subAdministrativeArea) !== null && _g !== void 0 ? _g : "";
        message.locality = (_h = object.locality) !== null && _h !== void 0 ? _h : "";
        message.subLocality = (_j = object.subLocality) !== null && _j !== void 0 ? _j : "";
        message.thoroughfare = (_k = object.thoroughfare) !== null && _k !== void 0 ? _k : "";
        message.subThoroughfare = (_l = object.subThoroughfare) !== null && _l !== void 0 ? _l : "";
        return message;
    }
};
function createBaseMetadata() {
    return {
        timestamp: 0,
        nodeId: "",
        signature: Buffer.alloc(0),
        publicKey: Buffer.alloc(0)
    };
}
exports.Metadata = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        if (message.timestamp !== 0) {
            writer.uint32(8).int64(message.timestamp);
        }
        if (message.nodeId !== "") {
            writer.uint32(18).string(message.nodeId);
        }
        if (message.signature.length !== 0) {
            writer.uint32(26).bytes(message.signature);
        }
        if (message.publicKey.length !== 0) {
            writer.uint32(34).bytes(message.publicKey);
        }
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBaseMetadata();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.timestamp = longToNumber(reader.int64());
                    break;
                case 2:
                    message.nodeId = reader.string();
                    break;
                case 3:
                    message.signature = reader.bytes();
                    break;
                case 4:
                    message.publicKey = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON: function (object) {
        return {
            timestamp: isSet(object.timestamp) ? Number(object.timestamp) : 0,
            nodeId: isSet(object.nodeId) ? String(object.nodeId) : "",
            signature: isSet(object.signature)
                ? Buffer.from(bytesFromBase64(object.signature))
                : Buffer.alloc(0),
            publicKey: isSet(object.publicKey)
                ? Buffer.from(bytesFromBase64(object.publicKey))
                : Buffer.alloc(0)
        };
    },
    toJSON: function (message) {
        var obj = {};
        message.timestamp !== undefined &&
            (obj.timestamp = Math.round(message.timestamp));
        message.nodeId !== undefined && (obj.nodeId = message.nodeId);
        message.signature !== undefined &&
            (obj.signature = base64FromBytes(message.signature !== undefined ? message.signature : Buffer.alloc(0)));
        message.publicKey !== undefined &&
            (obj.publicKey = base64FromBytes(message.publicKey !== undefined ? message.publicKey : Buffer.alloc(0)));
        return obj;
    },
    fromPartial: function (object) {
        var _a, _b, _c, _d;
        var message = createBaseMetadata();
        message.timestamp = (_a = object.timestamp) !== null && _a !== void 0 ? _a : 0;
        message.nodeId = (_b = object.nodeId) !== null && _b !== void 0 ? _b : "";
        message.signature = (_c = object.signature) !== null && _c !== void 0 ? _c : Buffer.alloc(0);
        message.publicKey = (_d = object.publicKey) !== null && _d !== void 0 ? _d : Buffer.alloc(0);
        return message;
    }
};
function createBaseMIME() {
    return { type: MIME_Type.TYPE_UNSPECIFIED, subtype: "", value: "" };
}
exports.MIME = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        if (message.type !== MIME_Type.TYPE_UNSPECIFIED) {
            writer.uint32(8).int32(mIME_TypeToNumber(message.type));
        }
        if (message.subtype !== "") {
            writer.uint32(18).string(message.subtype);
        }
        if (message.value !== "") {
            writer.uint32(26).string(message.value);
        }
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBaseMIME();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.type = mIME_TypeFromJSON(reader.int32());
                    break;
                case 2:
                    message.subtype = reader.string();
                    break;
                case 3:
                    message.value = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON: function (object) {
        return {
            type: isSet(object.type)
                ? mIME_TypeFromJSON(object.type)
                : MIME_Type.TYPE_UNSPECIFIED,
            subtype: isSet(object.subtype) ? String(object.subtype) : "",
            value: isSet(object.value) ? String(object.value) : ""
        };
    },
    toJSON: function (message) {
        var obj = {};
        message.type !== undefined && (obj.type = mIME_TypeToJSON(message.type));
        message.subtype !== undefined && (obj.subtype = message.subtype);
        message.value !== undefined && (obj.value = message.value);
        return obj;
    },
    fromPartial: function (object) {
        var _a, _b, _c;
        var message = createBaseMIME();
        message.type = (_a = object.type) !== null && _a !== void 0 ? _a : MIME_Type.TYPE_UNSPECIFIED;
        message.subtype = (_b = object.subtype) !== null && _b !== void 0 ? _b : "";
        message.value = (_c = object.value) !== null && _c !== void 0 ? _c : "";
        return message;
    }
};
function createBasePeer() {
    return {
        sName: "",
        status: Peer_Status.STATUS_UNSPECIFIED,
        device: undefined,
        profile: undefined,
        publicKey: Buffer.alloc(0),
        peerId: "",
        lastModified: 0
    };
}
exports.Peer = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        if (message.sName !== "") {
            writer.uint32(10).string(message.sName);
        }
        if (message.status !== Peer_Status.STATUS_UNSPECIFIED) {
            writer.uint32(16).int32(peer_StatusToNumber(message.status));
        }
        if (message.device !== undefined) {
            exports.Peer_Device.encode(message.device, writer.uint32(26).fork()).ldelim();
        }
        if (message.profile !== undefined) {
            exports.Profile.encode(message.profile, writer.uint32(34).fork()).ldelim();
        }
        if (message.publicKey.length !== 0) {
            writer.uint32(42).bytes(message.publicKey);
        }
        if (message.peerId !== "") {
            writer.uint32(50).string(message.peerId);
        }
        if (message.lastModified !== 0) {
            writer.uint32(56).int64(message.lastModified);
        }
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBasePeer();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.sName = reader.string();
                    break;
                case 2:
                    message.status = peer_StatusFromJSON(reader.int32());
                    break;
                case 3:
                    message.device = exports.Peer_Device.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.profile = exports.Profile.decode(reader, reader.uint32());
                    break;
                case 5:
                    message.publicKey = reader.bytes();
                    break;
                case 6:
                    message.peerId = reader.string();
                    break;
                case 7:
                    message.lastModified = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON: function (object) {
        return {
            sName: isSet(object.sName) ? String(object.sName) : "",
            status: isSet(object.status)
                ? peer_StatusFromJSON(object.status)
                : Peer_Status.STATUS_UNSPECIFIED,
            device: isSet(object.device)
                ? exports.Peer_Device.fromJSON(object.device)
                : undefined,
            profile: isSet(object.profile)
                ? exports.Profile.fromJSON(object.profile)
                : undefined,
            publicKey: isSet(object.publicKey)
                ? Buffer.from(bytesFromBase64(object.publicKey))
                : Buffer.alloc(0),
            peerId: isSet(object.peerId) ? String(object.peerId) : "",
            lastModified: isSet(object.lastModified)
                ? Number(object.lastModified)
                : 0
        };
    },
    toJSON: function (message) {
        var obj = {};
        message.sName !== undefined && (obj.sName = message.sName);
        message.status !== undefined &&
            (obj.status = peer_StatusToJSON(message.status));
        message.device !== undefined &&
            (obj.device = message.device
                ? exports.Peer_Device.toJSON(message.device)
                : undefined);
        message.profile !== undefined &&
            (obj.profile = message.profile
                ? exports.Profile.toJSON(message.profile)
                : undefined);
        message.publicKey !== undefined &&
            (obj.publicKey = base64FromBytes(message.publicKey !== undefined ? message.publicKey : Buffer.alloc(0)));
        message.peerId !== undefined && (obj.peerId = message.peerId);
        message.lastModified !== undefined &&
            (obj.lastModified = Math.round(message.lastModified));
        return obj;
    },
    fromPartial: function (object) {
        var _a, _b, _c, _d, _e;
        var message = createBasePeer();
        message.sName = (_a = object.sName) !== null && _a !== void 0 ? _a : "";
        message.status = (_b = object.status) !== null && _b !== void 0 ? _b : Peer_Status.STATUS_UNSPECIFIED;
        message.device =
            object.device !== undefined && object.device !== null
                ? exports.Peer_Device.fromPartial(object.device)
                : undefined;
        message.profile =
            object.profile !== undefined && object.profile !== null
                ? exports.Profile.fromPartial(object.profile)
                : undefined;
        message.publicKey = (_c = object.publicKey) !== null && _c !== void 0 ? _c : Buffer.alloc(0);
        message.peerId = (_d = object.peerId) !== null && _d !== void 0 ? _d : "";
        message.lastModified = (_e = object.lastModified) !== null && _e !== void 0 ? _e : 0;
        return message;
    }
};
function createBasePeer_Device() {
    return { id: "", hostName: "", os: "", arch: "", model: "" };
}
exports.Peer_Device = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.hostName !== "") {
            writer.uint32(18).string(message.hostName);
        }
        if (message.os !== "") {
            writer.uint32(26).string(message.os);
        }
        if (message.arch !== "") {
            writer.uint32(34).string(message.arch);
        }
        if (message.model !== "") {
            writer.uint32(42).string(message.model);
        }
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBasePeer_Device();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.hostName = reader.string();
                    break;
                case 3:
                    message.os = reader.string();
                    break;
                case 4:
                    message.arch = reader.string();
                    break;
                case 5:
                    message.model = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON: function (object) {
        return {
            id: isSet(object.id) ? String(object.id) : "",
            hostName: isSet(object.hostName) ? String(object.hostName) : "",
            os: isSet(object.os) ? String(object.os) : "",
            arch: isSet(object.arch) ? String(object.arch) : "",
            model: isSet(object.model) ? String(object.model) : ""
        };
    },
    toJSON: function (message) {
        var obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.hostName !== undefined && (obj.hostName = message.hostName);
        message.os !== undefined && (obj.os = message.os);
        message.arch !== undefined && (obj.arch = message.arch);
        message.model !== undefined && (obj.model = message.model);
        return obj;
    },
    fromPartial: function (object) {
        var _a, _b, _c, _d, _e;
        var message = createBasePeer_Device();
        message.id = (_a = object.id) !== null && _a !== void 0 ? _a : "";
        message.hostName = (_b = object.hostName) !== null && _b !== void 0 ? _b : "";
        message.os = (_c = object.os) !== null && _c !== void 0 ? _c : "";
        message.arch = (_d = object.arch) !== null && _d !== void 0 ? _d : "";
        message.model = (_e = object.model) !== null && _e !== void 0 ? _e : "";
        return message;
    }
};
function createBaseProfile() {
    return {
        sName: "",
        firstName: "",
        lastName: "",
        picture: Buffer.alloc(0),
        bio: "",
        lastModified: 0
    };
}
exports.Profile = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        if (message.sName !== "") {
            writer.uint32(10).string(message.sName);
        }
        if (message.firstName !== "") {
            writer.uint32(18).string(message.firstName);
        }
        if (message.lastName !== "") {
            writer.uint32(26).string(message.lastName);
        }
        if (message.picture.length !== 0) {
            writer.uint32(34).bytes(message.picture);
        }
        if (message.bio !== "") {
            writer.uint32(50).string(message.bio);
        }
        if (message.lastModified !== 0) {
            writer.uint32(56).int64(message.lastModified);
        }
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBaseProfile();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.sName = reader.string();
                    break;
                case 2:
                    message.firstName = reader.string();
                    break;
                case 3:
                    message.lastName = reader.string();
                    break;
                case 4:
                    message.picture = reader.bytes();
                    break;
                case 6:
                    message.bio = reader.string();
                    break;
                case 7:
                    message.lastModified = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON: function (object) {
        return {
            sName: isSet(object.sName) ? String(object.sName) : "",
            firstName: isSet(object.firstName) ? String(object.firstName) : "",
            lastName: isSet(object.lastName) ? String(object.lastName) : "",
            picture: isSet(object.picture)
                ? Buffer.from(bytesFromBase64(object.picture))
                : Buffer.alloc(0),
            bio: isSet(object.bio) ? String(object.bio) : "",
            lastModified: isSet(object.lastModified)
                ? Number(object.lastModified)
                : 0
        };
    },
    toJSON: function (message) {
        var obj = {};
        message.sName !== undefined && (obj.sName = message.sName);
        message.firstName !== undefined && (obj.firstName = message.firstName);
        message.lastName !== undefined && (obj.lastName = message.lastName);
        message.picture !== undefined &&
            (obj.picture = base64FromBytes(message.picture !== undefined ? message.picture : Buffer.alloc(0)));
        message.bio !== undefined && (obj.bio = message.bio);
        message.lastModified !== undefined &&
            (obj.lastModified = Math.round(message.lastModified));
        return obj;
    },
    fromPartial: function (object) {
        var _a, _b, _c, _d, _e, _f;
        var message = createBaseProfile();
        message.sName = (_a = object.sName) !== null && _a !== void 0 ? _a : "";
        message.firstName = (_b = object.firstName) !== null && _b !== void 0 ? _b : "";
        message.lastName = (_c = object.lastName) !== null && _c !== void 0 ? _c : "";
        message.picture = (_d = object.picture) !== null && _d !== void 0 ? _d : Buffer.alloc(0);
        message.bio = (_e = object.bio) !== null && _e !== void 0 ? _e : "";
        message.lastModified = (_f = object.lastModified) !== null && _f !== void 0 ? _f : 0;
        return message;
    }
};
function createBaseProfileList() {
    return { profiles: [], createdAt: 0, key: "", lastModified: 0 };
}
exports.ProfileList = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        for (var _i = 0, _a = message.profiles; _i < _a.length; _i++) {
            var v = _a[_i];
            exports.Profile.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.createdAt !== 0) {
            writer.uint32(16).int64(message.createdAt);
        }
        if (message.key !== "") {
            writer.uint32(26).string(message.key);
        }
        if (message.lastModified !== 0) {
            writer.uint32(32).int64(message.lastModified);
        }
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBaseProfileList();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.profiles.push(exports.Profile.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.createdAt = longToNumber(reader.int64());
                    break;
                case 3:
                    message.key = reader.string();
                    break;
                case 4:
                    message.lastModified = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON: function (object) {
        return {
            profiles: Array.isArray(object === null || object === void 0 ? void 0 : object.profiles)
                ? object.profiles.map(function (e) { return exports.Profile.fromJSON(e); })
                : [],
            createdAt: isSet(object.createdAt) ? Number(object.createdAt) : 0,
            key: isSet(object.key) ? String(object.key) : "",
            lastModified: isSet(object.lastModified)
                ? Number(object.lastModified)
                : 0
        };
    },
    toJSON: function (message) {
        var obj = {};
        if (message.profiles) {
            obj.profiles = message.profiles.map(function (e) {
                return e ? exports.Profile.toJSON(e) : undefined;
            });
        }
        else {
            obj.profiles = [];
        }
        message.createdAt !== undefined &&
            (obj.createdAt = Math.round(message.createdAt));
        message.key !== undefined && (obj.key = message.key);
        message.lastModified !== undefined &&
            (obj.lastModified = Math.round(message.lastModified));
        return obj;
    },
    fromPartial: function (object) {
        var _a, _b, _c, _d;
        var message = createBaseProfileList();
        message.profiles =
            ((_a = object.profiles) === null || _a === void 0 ? void 0 : _a.map(function (e) { return exports.Profile.fromPartial(e); })) || [];
        message.createdAt = (_b = object.createdAt) !== null && _b !== void 0 ? _b : 0;
        message.key = (_c = object.key) !== null && _c !== void 0 ? _c : "";
        message.lastModified = (_d = object.lastModified) !== null && _d !== void 0 ? _d : 0;
        return message;
    }
};
var globalThis = (function () {
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
var atob = globalThis.atob ||
    (function (b64) { return globalThis.Buffer.from(b64, "base64").toString("binary"); });
function bytesFromBase64(b64) {
    var bin = atob(b64);
    var arr = new Uint8Array(bin.length);
    for (var i = 0; i < bin.length; ++i) {
        arr[i] = bin.charCodeAt(i);
    }
    return arr;
}
var btoa = globalThis.btoa ||
    (function (bin) { return globalThis.Buffer.from(bin, "binary").toString("base64"); });
function base64FromBytes(arr) {
    var bin = [];
    for (var _i = 0, arr_1 = arr; _i < arr_1.length; _i++) {
        var byte = arr_1[_i];
        bin.push(String.fromCharCode(byte));
    }
    return btoa(bin.join(""));
}
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (minimal_1["default"].util.Long !== long_1["default"]) {
    minimal_1["default"].util.Long = long_1["default"];
    minimal_1["default"].configure();
}
function isSet(value) {
    return value !== null && value !== undefined;
}
