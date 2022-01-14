"use strict";
exports.__esModule = true;
exports.ObjectField_MetadataEntry = exports.ObjectField = exports.ObjectDoc_FieldsEntry = exports.ObjectDoc = exports.Bucket = exports.objectFieldTypeToNumber = exports.objectFieldTypeToJSON = exports.objectFieldTypeFromJSON = exports.ObjectFieldType = exports.bucketTypeToNumber = exports.bucketTypeToJSON = exports.bucketTypeFromJSON = exports.BucketType = exports.protobufPackage = void 0;
/* eslint-disable */
var long_1 = require("long");
var minimal_1 = require("protobufjs/minimal");
exports.protobufPackage = "common";
/** BucketType is the type of a bucket. */
var BucketType;
(function (BucketType) {
    /** BUCKET_TYPE_UNSPECIFIED - BucketTypeUnspecified is the default value. */
    BucketType["BUCKET_TYPE_UNSPECIFIED"] = "BUCKET_TYPE_UNSPECIFIED";
    /** BUCKET_TYPE_APP - BucketTypeApp is an App specific bucket. For Assets regarding the service. */
    BucketType["BUCKET_TYPE_APP"] = "BUCKET_TYPE_APP";
    /**
     * BUCKET_TYPE_USER - BucketTypeUser is a User specific bucket. For any remote user data that is required
     * to be stored in the Network.
     */
    BucketType["BUCKET_TYPE_USER"] = "BUCKET_TYPE_USER";
    BucketType["UNRECOGNIZED"] = "UNRECOGNIZED";
})(BucketType = exports.BucketType || (exports.BucketType = {}));
function bucketTypeFromJSON(object) {
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
exports.bucketTypeFromJSON = bucketTypeFromJSON;
function bucketTypeToJSON(object) {
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
exports.bucketTypeToJSON = bucketTypeToJSON;
function bucketTypeToNumber(object) {
    switch (object) {
        case BucketType.BUCKET_TYPE_UNSPECIFIED:
            return 0;
        case BucketType.BUCKET_TYPE_APP:
            return 1;
        case BucketType.BUCKET_TYPE_USER:
            return 2;
        default:
            return 0;
    }
}
exports.bucketTypeToNumber = bucketTypeToNumber;
/** ObjectFieldType is the type of the field */
var ObjectFieldType;
(function (ObjectFieldType) {
    /** OBJECT_FIELD_TYPE_UNSPECIFIED - ObjectFieldTypeUnspecified is the default value */
    ObjectFieldType["OBJECT_FIELD_TYPE_UNSPECIFIED"] = "OBJECT_FIELD_TYPE_UNSPECIFIED";
    /** OBJECT_FIELD_TYPE_STRING - ObjectFieldTypeString is a string or text field */
    ObjectFieldType["OBJECT_FIELD_TYPE_STRING"] = "OBJECT_FIELD_TYPE_STRING";
    /** OBJECT_FIELD_TYPE_INT - ObjectFieldTypeInt is an integer */
    ObjectFieldType["OBJECT_FIELD_TYPE_INT"] = "OBJECT_FIELD_TYPE_INT";
    /** OBJECT_FIELD_TYPE_FLOAT - ObjectFieldTypeFloat is a floating point number */
    ObjectFieldType["OBJECT_FIELD_TYPE_FLOAT"] = "OBJECT_FIELD_TYPE_FLOAT";
    /** OBJECT_FIELD_TYPE_BOOL - ObjectFieldTypeBool is a boolean */
    ObjectFieldType["OBJECT_FIELD_TYPE_BOOL"] = "OBJECT_FIELD_TYPE_BOOL";
    /** OBJECT_FIELD_TYPE_DATETIME - ObjectFieldTypeDateTime is a datetime */
    ObjectFieldType["OBJECT_FIELD_TYPE_DATETIME"] = "OBJECT_FIELD_TYPE_DATETIME";
    /** OBJECT_FIELD_TYPE_BLOB - ObjectFieldTypeBlob is a blob which is a byte array */
    ObjectFieldType["OBJECT_FIELD_TYPE_BLOB"] = "OBJECT_FIELD_TYPE_BLOB";
    /** OBJECT_FIELD_TYPE_REFERENCE - ObjectFieldTypeReference is a reference to another object */
    ObjectFieldType["OBJECT_FIELD_TYPE_REFERENCE"] = "OBJECT_FIELD_TYPE_REFERENCE";
    ObjectFieldType["UNRECOGNIZED"] = "UNRECOGNIZED";
})(ObjectFieldType = exports.ObjectFieldType || (exports.ObjectFieldType = {}));
function objectFieldTypeFromJSON(object) {
    switch (object) {
        case 0:
        case "OBJECT_FIELD_TYPE_UNSPECIFIED":
            return ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED;
        case 1:
        case "OBJECT_FIELD_TYPE_STRING":
            return ObjectFieldType.OBJECT_FIELD_TYPE_STRING;
        case 2:
        case "OBJECT_FIELD_TYPE_INT":
            return ObjectFieldType.OBJECT_FIELD_TYPE_INT;
        case 3:
        case "OBJECT_FIELD_TYPE_FLOAT":
            return ObjectFieldType.OBJECT_FIELD_TYPE_FLOAT;
        case 4:
        case "OBJECT_FIELD_TYPE_BOOL":
            return ObjectFieldType.OBJECT_FIELD_TYPE_BOOL;
        case 5:
        case "OBJECT_FIELD_TYPE_DATETIME":
            return ObjectFieldType.OBJECT_FIELD_TYPE_DATETIME;
        case 6:
        case "OBJECT_FIELD_TYPE_BLOB":
            return ObjectFieldType.OBJECT_FIELD_TYPE_BLOB;
        case 7:
        case "OBJECT_FIELD_TYPE_REFERENCE":
            return ObjectFieldType.OBJECT_FIELD_TYPE_REFERENCE;
        case -1:
        case "UNRECOGNIZED":
        default:
            return ObjectFieldType.UNRECOGNIZED;
    }
}
exports.objectFieldTypeFromJSON = objectFieldTypeFromJSON;
function objectFieldTypeToJSON(object) {
    switch (object) {
        case ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED:
            return "OBJECT_FIELD_TYPE_UNSPECIFIED";
        case ObjectFieldType.OBJECT_FIELD_TYPE_STRING:
            return "OBJECT_FIELD_TYPE_STRING";
        case ObjectFieldType.OBJECT_FIELD_TYPE_INT:
            return "OBJECT_FIELD_TYPE_INT";
        case ObjectFieldType.OBJECT_FIELD_TYPE_FLOAT:
            return "OBJECT_FIELD_TYPE_FLOAT";
        case ObjectFieldType.OBJECT_FIELD_TYPE_BOOL:
            return "OBJECT_FIELD_TYPE_BOOL";
        case ObjectFieldType.OBJECT_FIELD_TYPE_DATETIME:
            return "OBJECT_FIELD_TYPE_DATETIME";
        case ObjectFieldType.OBJECT_FIELD_TYPE_BLOB:
            return "OBJECT_FIELD_TYPE_BLOB";
        case ObjectFieldType.OBJECT_FIELD_TYPE_REFERENCE:
            return "OBJECT_FIELD_TYPE_REFERENCE";
        default:
            return "UNKNOWN";
    }
}
exports.objectFieldTypeToJSON = objectFieldTypeToJSON;
function objectFieldTypeToNumber(object) {
    switch (object) {
        case ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED:
            return 0;
        case ObjectFieldType.OBJECT_FIELD_TYPE_STRING:
            return 1;
        case ObjectFieldType.OBJECT_FIELD_TYPE_INT:
            return 2;
        case ObjectFieldType.OBJECT_FIELD_TYPE_FLOAT:
            return 3;
        case ObjectFieldType.OBJECT_FIELD_TYPE_BOOL:
            return 4;
        case ObjectFieldType.OBJECT_FIELD_TYPE_DATETIME:
            return 5;
        case ObjectFieldType.OBJECT_FIELD_TYPE_BLOB:
            return 6;
        case ObjectFieldType.OBJECT_FIELD_TYPE_REFERENCE:
            return 7;
        default:
            return 0;
    }
}
exports.objectFieldTypeToNumber = objectFieldTypeToNumber;
function createBaseBucket() {
    return {
        label: "",
        description: "",
        type: BucketType.BUCKET_TYPE_UNSPECIFIED,
        did: "",
        objects: []
    };
}
exports.Bucket = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        if (message.label !== "") {
            writer.uint32(10).string(message.label);
        }
        if (message.description !== "") {
            writer.uint32(18).string(message.description);
        }
        if (message.type !== BucketType.BUCKET_TYPE_UNSPECIFIED) {
            writer.uint32(24).int32(bucketTypeToNumber(message.type));
        }
        if (message.did !== "") {
            writer.uint32(34).string(message.did);
        }
        for (var _i = 0, _a = message.objects; _i < _a.length; _i++) {
            var v = _a[_i];
            exports.ObjectDoc.encode(v, writer.uint32(42).fork()).ldelim();
        }
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBaseBucket();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.label = reader.string();
                    break;
                case 2:
                    message.description = reader.string();
                    break;
                case 3:
                    message.type = bucketTypeFromJSON(reader.int32());
                    break;
                case 4:
                    message.did = reader.string();
                    break;
                case 5:
                    message.objects.push(exports.ObjectDoc.decode(reader, reader.uint32()));
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
            label: isSet(object.label) ? String(object.label) : "",
            description: isSet(object.description) ? String(object.description) : "",
            type: isSet(object.type)
                ? bucketTypeFromJSON(object.type)
                : BucketType.BUCKET_TYPE_UNSPECIFIED,
            did: isSet(object.did) ? String(object.did) : "",
            objects: Array.isArray(object === null || object === void 0 ? void 0 : object.objects)
                ? object.objects.map(function (e) { return exports.ObjectDoc.fromJSON(e); })
                : []
        };
    },
    toJSON: function (message) {
        var obj = {};
        message.label !== undefined && (obj.label = message.label);
        message.description !== undefined &&
            (obj.description = message.description);
        message.type !== undefined && (obj.type = bucketTypeToJSON(message.type));
        message.did !== undefined && (obj.did = message.did);
        if (message.objects) {
            obj.objects = message.objects.map(function (e) {
                return e ? exports.ObjectDoc.toJSON(e) : undefined;
            });
        }
        else {
            obj.objects = [];
        }
        return obj;
    },
    fromPartial: function (object) {
        var _a, _b, _c, _d, _e;
        var message = createBaseBucket();
        message.label = (_a = object.label) !== null && _a !== void 0 ? _a : "";
        message.description = (_b = object.description) !== null && _b !== void 0 ? _b : "";
        message.type = (_c = object.type) !== null && _c !== void 0 ? _c : BucketType.BUCKET_TYPE_UNSPECIFIED;
        message.did = (_d = object.did) !== null && _d !== void 0 ? _d : "";
        message.objects =
            ((_e = object.objects) === null || _e === void 0 ? void 0 : _e.map(function (e) { return exports.ObjectDoc.fromPartial(e); })) || [];
        return message;
    }
};
function createBaseObjectDoc() {
    return { did: "", service: "", tags: [], fields: {} };
}
exports.ObjectDoc = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        if (message.did !== "") {
            writer.uint32(10).string(message.did);
        }
        if (message.service !== "") {
            writer.uint32(18).string(message.service);
        }
        for (var _i = 0, _a = message.tags; _i < _a.length; _i++) {
            var v = _a[_i];
            writer.uint32(26).string(v);
        }
        Object.entries(message.fields).forEach(function (_a) {
            var key = _a[0], value = _a[1];
            exports.ObjectDoc_FieldsEntry.encode({ key: key, value: value }, writer.uint32(34).fork()).ldelim();
        });
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBaseObjectDoc();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.did = reader.string();
                    break;
                case 2:
                    message.service = reader.string();
                    break;
                case 3:
                    message.tags.push(reader.string());
                    break;
                case 4:
                    var entry4 = exports.ObjectDoc_FieldsEntry.decode(reader, reader.uint32());
                    if (entry4.value !== undefined) {
                        message.fields[entry4.key] = entry4.value;
                    }
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
            did: isSet(object.did) ? String(object.did) : "",
            service: isSet(object.service) ? String(object.service) : "",
            tags: Array.isArray(object === null || object === void 0 ? void 0 : object.tags)
                ? object.tags.map(function (e) { return String(e); })
                : [],
            fields: isObject(object.fields)
                ? Object.entries(object.fields).reduce(function (acc, _a) {
                    var key = _a[0], value = _a[1];
                    acc[key] = exports.ObjectField.fromJSON(value);
                    return acc;
                }, {})
                : {}
        };
    },
    toJSON: function (message) {
        var obj = {};
        message.did !== undefined && (obj.did = message.did);
        message.service !== undefined && (obj.service = message.service);
        if (message.tags) {
            obj.tags = message.tags.map(function (e) { return e; });
        }
        else {
            obj.tags = [];
        }
        obj.fields = {};
        if (message.fields) {
            Object.entries(message.fields).forEach(function (_a) {
                var k = _a[0], v = _a[1];
                obj.fields[k] = exports.ObjectField.toJSON(v);
            });
        }
        return obj;
    },
    fromPartial: function (object) {
        var _a, _b, _c, _d;
        var message = createBaseObjectDoc();
        message.did = (_a = object.did) !== null && _a !== void 0 ? _a : "";
        message.service = (_b = object.service) !== null && _b !== void 0 ? _b : "";
        message.tags = ((_c = object.tags) === null || _c === void 0 ? void 0 : _c.map(function (e) { return e; })) || [];
        message.fields = Object.entries((_d = object.fields) !== null && _d !== void 0 ? _d : {}).reduce(function (acc, _a) {
            var key = _a[0], value = _a[1];
            if (value !== undefined) {
                acc[key] = exports.ObjectField.fromPartial(value);
            }
            return acc;
        }, {});
        return message;
    }
};
function createBaseObjectDoc_FieldsEntry() {
    return { key: "", value: undefined };
}
exports.ObjectDoc_FieldsEntry = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        if (message.key !== "") {
            writer.uint32(10).string(message.key);
        }
        if (message.value !== undefined) {
            exports.ObjectField.encode(message.value, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBaseObjectDoc_FieldsEntry();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.key = reader.string();
                    break;
                case 2:
                    message.value = exports.ObjectField.decode(reader, reader.uint32());
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
            key: isSet(object.key) ? String(object.key) : "",
            value: isSet(object.value)
                ? exports.ObjectField.fromJSON(object.value)
                : undefined
        };
    },
    toJSON: function (message) {
        var obj = {};
        message.key !== undefined && (obj.key = message.key);
        message.value !== undefined &&
            (obj.value = message.value
                ? exports.ObjectField.toJSON(message.value)
                : undefined);
        return obj;
    },
    fromPartial: function (object) {
        var _a;
        var message = createBaseObjectDoc_FieldsEntry();
        message.key = (_a = object.key) !== null && _a !== void 0 ? _a : "";
        message.value =
            object.value !== undefined && object.value !== null
                ? exports.ObjectField.fromPartial(object.value)
                : undefined;
        return message;
    }
};
function createBaseObjectField() {
    return {
        name: "",
        type: ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED,
        stringValue: undefined,
        intValue: undefined,
        floatValue: undefined,
        boolValue: undefined,
        dateValue: undefined,
        blobValue: undefined,
        referenceValue: undefined,
        metadata: {}
    };
}
exports.ObjectField = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        if (message.name !== "") {
            writer.uint32(10).string(message.name);
        }
        if (message.type !== ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED) {
            writer.uint32(16).int32(objectFieldTypeToNumber(message.type));
        }
        if (message.stringValue !== undefined) {
            writer.uint32(26).string(message.stringValue);
        }
        if (message.intValue !== undefined) {
            writer.uint32(32).int32(message.intValue);
        }
        if (message.floatValue !== undefined) {
            writer.uint32(41).double(message.floatValue);
        }
        if (message.boolValue !== undefined) {
            writer.uint32(48).bool(message.boolValue);
        }
        if (message.dateValue !== undefined) {
            writer.uint32(56).int64(message.dateValue);
        }
        if (message.blobValue !== undefined) {
            writer.uint32(66).bytes(message.blobValue);
        }
        if (message.referenceValue !== undefined) {
            writer.uint32(74).string(message.referenceValue);
        }
        Object.entries(message.metadata).forEach(function (_a) {
            var key = _a[0], value = _a[1];
            exports.ObjectField_MetadataEntry.encode({ key: key, value: value }, writer.uint32(82).fork()).ldelim();
        });
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBaseObjectField();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.name = reader.string();
                    break;
                case 2:
                    message.type = objectFieldTypeFromJSON(reader.int32());
                    break;
                case 3:
                    message.stringValue = reader.string();
                    break;
                case 4:
                    message.intValue = reader.int32();
                    break;
                case 5:
                    message.floatValue = reader.double();
                    break;
                case 6:
                    message.boolValue = reader.bool();
                    break;
                case 7:
                    message.dateValue = longToNumber(reader.int64());
                    break;
                case 8:
                    message.blobValue = reader.bytes();
                    break;
                case 9:
                    message.referenceValue = reader.string();
                    break;
                case 10:
                    var entry10 = exports.ObjectField_MetadataEntry.decode(reader, reader.uint32());
                    if (entry10.value !== undefined) {
                        message.metadata[entry10.key] = entry10.value;
                    }
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
            type: isSet(object.type)
                ? objectFieldTypeFromJSON(object.type)
                : ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED,
            stringValue: isSet(object.stringValue)
                ? String(object.stringValue)
                : undefined,
            intValue: isSet(object.intValue) ? Number(object.intValue) : undefined,
            floatValue: isSet(object.floatValue)
                ? Number(object.floatValue)
                : undefined,
            boolValue: isSet(object.boolValue)
                ? Boolean(object.boolValue)
                : undefined,
            dateValue: isSet(object.dateValue) ? Number(object.dateValue) : undefined,
            blobValue: isSet(object.blobValue)
                ? Buffer.from(bytesFromBase64(object.blobValue))
                : undefined,
            referenceValue: isSet(object.referenceValue)
                ? String(object.referenceValue)
                : undefined,
            metadata: isObject(object.metadata)
                ? Object.entries(object.metadata).reduce(function (acc, _a) {
                    var key = _a[0], value = _a[1];
                    acc[key] = String(value);
                    return acc;
                }, {})
                : {}
        };
    },
    toJSON: function (message) {
        var obj = {};
        message.name !== undefined && (obj.name = message.name);
        message.type !== undefined &&
            (obj.type = objectFieldTypeToJSON(message.type));
        message.stringValue !== undefined &&
            (obj.stringValue = message.stringValue);
        message.intValue !== undefined &&
            (obj.intValue = Math.round(message.intValue));
        message.floatValue !== undefined && (obj.floatValue = message.floatValue);
        message.boolValue !== undefined && (obj.boolValue = message.boolValue);
        message.dateValue !== undefined &&
            (obj.dateValue = Math.round(message.dateValue));
        message.blobValue !== undefined &&
            (obj.blobValue =
                message.blobValue !== undefined
                    ? base64FromBytes(message.blobValue)
                    : undefined);
        message.referenceValue !== undefined &&
            (obj.referenceValue = message.referenceValue);
        obj.metadata = {};
        if (message.metadata) {
            Object.entries(message.metadata).forEach(function (_a) {
                var k = _a[0], v = _a[1];
                obj.metadata[k] = v;
            });
        }
        return obj;
    },
    fromPartial: function (object) {
        var _a, _b, _c, _d, _e, _f, _g, _h, _j, _k;
        var message = createBaseObjectField();
        message.name = (_a = object.name) !== null && _a !== void 0 ? _a : "";
        message.type = (_b = object.type) !== null && _b !== void 0 ? _b : ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED;
        message.stringValue = (_c = object.stringValue) !== null && _c !== void 0 ? _c : undefined;
        message.intValue = (_d = object.intValue) !== null && _d !== void 0 ? _d : undefined;
        message.floatValue = (_e = object.floatValue) !== null && _e !== void 0 ? _e : undefined;
        message.boolValue = (_f = object.boolValue) !== null && _f !== void 0 ? _f : undefined;
        message.dateValue = (_g = object.dateValue) !== null && _g !== void 0 ? _g : undefined;
        message.blobValue = (_h = object.blobValue) !== null && _h !== void 0 ? _h : undefined;
        message.referenceValue = (_j = object.referenceValue) !== null && _j !== void 0 ? _j : undefined;
        message.metadata = Object.entries((_k = object.metadata) !== null && _k !== void 0 ? _k : {}).reduce(function (acc, _a) {
            var key = _a[0], value = _a[1];
            if (value !== undefined) {
                acc[key] = String(value);
            }
            return acc;
        }, {});
        return message;
    }
};
function createBaseObjectField_MetadataEntry() {
    return { key: "", value: "" };
}
exports.ObjectField_MetadataEntry = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        if (message.key !== "") {
            writer.uint32(10).string(message.key);
        }
        if (message.value !== "") {
            writer.uint32(18).string(message.value);
        }
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBaseObjectField_MetadataEntry();
        while (reader.pos < end) {
            var tag = reader.uint32();
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
    fromJSON: function (object) {
        return {
            key: isSet(object.key) ? String(object.key) : "",
            value: isSet(object.value) ? String(object.value) : ""
        };
    },
    toJSON: function (message) {
        var obj = {};
        message.key !== undefined && (obj.key = message.key);
        message.value !== undefined && (obj.value = message.value);
        return obj;
    },
    fromPartial: function (object) {
        var _a, _b;
        var message = createBaseObjectField_MetadataEntry();
        message.key = (_a = object.key) !== null && _a !== void 0 ? _a : "";
        message.value = (_b = object.value) !== null && _b !== void 0 ? _b : "";
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
function isObject(value) {
    return typeof value === "object" && value !== null;
}
function isSet(value) {
    return value !== null && value !== undefined;
}
