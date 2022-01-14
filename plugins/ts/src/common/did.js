"use strict";
exports.__esModule = true;
exports.VerificationMethod = exports.Service = exports.DidDocument = exports.Did = exports.protobufPackage = void 0;
/* eslint-disable */
var long_1 = require("long");
var minimal_1 = require("protobufjs/minimal");
exports.protobufPackage = "common";
function createBaseDid() {
    return {
        method: "",
        network: "",
        id: "",
        paths: [],
        query: "",
        fragment: ""
    };
}
exports.Did = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        if (message.method !== "") {
            writer.uint32(10).string(message.method);
        }
        if (message.network !== "") {
            writer.uint32(18).string(message.network);
        }
        if (message.id !== "") {
            writer.uint32(26).string(message.id);
        }
        for (var _i = 0, _a = message.paths; _i < _a.length; _i++) {
            var v = _a[_i];
            writer.uint32(34).string(v);
        }
        if (message.query !== "") {
            writer.uint32(42).string(message.query);
        }
        if (message.fragment !== "") {
            writer.uint32(50).string(message.fragment);
        }
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBaseDid();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.method = reader.string();
                    break;
                case 2:
                    message.network = reader.string();
                    break;
                case 3:
                    message.id = reader.string();
                    break;
                case 4:
                    message.paths.push(reader.string());
                    break;
                case 5:
                    message.query = reader.string();
                    break;
                case 6:
                    message.fragment = reader.string();
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
            method: isSet(object.method) ? String(object.method) : "",
            network: isSet(object.network) ? String(object.network) : "",
            id: isSet(object.id) ? String(object.id) : "",
            paths: Array.isArray(object === null || object === void 0 ? void 0 : object.paths)
                ? object.paths.map(function (e) { return String(e); })
                : [],
            query: isSet(object.query) ? String(object.query) : "",
            fragment: isSet(object.fragment) ? String(object.fragment) : ""
        };
    },
    toJSON: function (message) {
        var obj = {};
        message.method !== undefined && (obj.method = message.method);
        message.network !== undefined && (obj.network = message.network);
        message.id !== undefined && (obj.id = message.id);
        if (message.paths) {
            obj.paths = message.paths.map(function (e) { return e; });
        }
        else {
            obj.paths = [];
        }
        message.query !== undefined && (obj.query = message.query);
        message.fragment !== undefined && (obj.fragment = message.fragment);
        return obj;
    },
    fromPartial: function (object) {
        var _a, _b, _c, _d, _e, _f;
        var message = createBaseDid();
        message.method = (_a = object.method) !== null && _a !== void 0 ? _a : "";
        message.network = (_b = object.network) !== null && _b !== void 0 ? _b : "";
        message.id = (_c = object.id) !== null && _c !== void 0 ? _c : "";
        message.paths = ((_d = object.paths) === null || _d === void 0 ? void 0 : _d.map(function (e) { return e; })) || [];
        message.query = (_e = object.query) !== null && _e !== void 0 ? _e : "";
        message.fragment = (_f = object.fragment) !== null && _f !== void 0 ? _f : "";
        return message;
    }
};
function createBaseDidDocument() {
    return {
        context: [],
        id: "",
        controller: [],
        verificationMethod: [],
        authentication: [],
        assertionMethod: [],
        capabilityInvocation: [],
        capabilityDelegation: [],
        keyAgreement: [],
        service: [],
        alsoKnownAs: []
    };
}
exports.DidDocument = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        for (var _i = 0, _a = message.context; _i < _a.length; _i++) {
            var v = _a[_i];
            writer.uint32(10).string(v);
        }
        if (message.id !== "") {
            writer.uint32(18).string(message.id);
        }
        for (var _b = 0, _c = message.controller; _b < _c.length; _b++) {
            var v = _c[_b];
            writer.uint32(26).string(v);
        }
        for (var _d = 0, _e = message.verificationMethod; _d < _e.length; _d++) {
            var v = _e[_d];
            exports.VerificationMethod.encode(v, writer.uint32(34).fork()).ldelim();
        }
        for (var _f = 0, _g = message.authentication; _f < _g.length; _f++) {
            var v = _g[_f];
            writer.uint32(42).string(v);
        }
        for (var _h = 0, _j = message.assertionMethod; _h < _j.length; _h++) {
            var v = _j[_h];
            writer.uint32(50).string(v);
        }
        for (var _k = 0, _l = message.capabilityInvocation; _k < _l.length; _k++) {
            var v = _l[_k];
            writer.uint32(58).string(v);
        }
        for (var _m = 0, _o = message.capabilityDelegation; _m < _o.length; _m++) {
            var v = _o[_m];
            writer.uint32(66).string(v);
        }
        for (var _p = 0, _q = message.keyAgreement; _p < _q.length; _p++) {
            var v = _q[_p];
            writer.uint32(74).string(v);
        }
        for (var _r = 0, _s = message.service; _r < _s.length; _r++) {
            var v = _s[_r];
            exports.Service.encode(v, writer.uint32(82).fork()).ldelim();
        }
        for (var _t = 0, _u = message.alsoKnownAs; _t < _u.length; _t++) {
            var v = _u[_t];
            writer.uint32(90).string(v);
        }
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBaseDidDocument();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.context.push(reader.string());
                    break;
                case 2:
                    message.id = reader.string();
                    break;
                case 3:
                    message.controller.push(reader.string());
                    break;
                case 4:
                    message.verificationMethod.push(exports.VerificationMethod.decode(reader, reader.uint32()));
                    break;
                case 5:
                    message.authentication.push(reader.string());
                    break;
                case 6:
                    message.assertionMethod.push(reader.string());
                    break;
                case 7:
                    message.capabilityInvocation.push(reader.string());
                    break;
                case 8:
                    message.capabilityDelegation.push(reader.string());
                    break;
                case 9:
                    message.keyAgreement.push(reader.string());
                    break;
                case 10:
                    message.service.push(exports.Service.decode(reader, reader.uint32()));
                    break;
                case 11:
                    message.alsoKnownAs.push(reader.string());
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
            context: Array.isArray(object === null || object === void 0 ? void 0 : object.context)
                ? object.context.map(function (e) { return String(e); })
                : [],
            id: isSet(object.id) ? String(object.id) : "",
            controller: Array.isArray(object === null || object === void 0 ? void 0 : object.controller)
                ? object.controller.map(function (e) { return String(e); })
                : [],
            verificationMethod: Array.isArray(object === null || object === void 0 ? void 0 : object.verificationMethod)
                ? object.verificationMethod.map(function (e) {
                    return exports.VerificationMethod.fromJSON(e);
                })
                : [],
            authentication: Array.isArray(object === null || object === void 0 ? void 0 : object.authentication)
                ? object.authentication.map(function (e) { return String(e); })
                : [],
            assertionMethod: Array.isArray(object === null || object === void 0 ? void 0 : object.assertionMethod)
                ? object.assertionMethod.map(function (e) { return String(e); })
                : [],
            capabilityInvocation: Array.isArray(object === null || object === void 0 ? void 0 : object.capabilityInvocation)
                ? object.capabilityInvocation.map(function (e) { return String(e); })
                : [],
            capabilityDelegation: Array.isArray(object === null || object === void 0 ? void 0 : object.capabilityDelegation)
                ? object.capabilityDelegation.map(function (e) { return String(e); })
                : [],
            keyAgreement: Array.isArray(object === null || object === void 0 ? void 0 : object.keyAgreement)
                ? object.keyAgreement.map(function (e) { return String(e); })
                : [],
            service: Array.isArray(object === null || object === void 0 ? void 0 : object.service)
                ? object.service.map(function (e) { return exports.Service.fromJSON(e); })
                : [],
            alsoKnownAs: Array.isArray(object === null || object === void 0 ? void 0 : object.alsoKnownAs)
                ? object.alsoKnownAs.map(function (e) { return String(e); })
                : []
        };
    },
    toJSON: function (message) {
        var obj = {};
        if (message.context) {
            obj.context = message.context.map(function (e) { return e; });
        }
        else {
            obj.context = [];
        }
        message.id !== undefined && (obj.id = message.id);
        if (message.controller) {
            obj.controller = message.controller.map(function (e) { return e; });
        }
        else {
            obj.controller = [];
        }
        if (message.verificationMethod) {
            obj.verificationMethod = message.verificationMethod.map(function (e) {
                return e ? exports.VerificationMethod.toJSON(e) : undefined;
            });
        }
        else {
            obj.verificationMethod = [];
        }
        if (message.authentication) {
            obj.authentication = message.authentication.map(function (e) { return e; });
        }
        else {
            obj.authentication = [];
        }
        if (message.assertionMethod) {
            obj.assertionMethod = message.assertionMethod.map(function (e) { return e; });
        }
        else {
            obj.assertionMethod = [];
        }
        if (message.capabilityInvocation) {
            obj.capabilityInvocation = message.capabilityInvocation.map(function (e) { return e; });
        }
        else {
            obj.capabilityInvocation = [];
        }
        if (message.capabilityDelegation) {
            obj.capabilityDelegation = message.capabilityDelegation.map(function (e) { return e; });
        }
        else {
            obj.capabilityDelegation = [];
        }
        if (message.keyAgreement) {
            obj.keyAgreement = message.keyAgreement.map(function (e) { return e; });
        }
        else {
            obj.keyAgreement = [];
        }
        if (message.service) {
            obj.service = message.service.map(function (e) {
                return e ? exports.Service.toJSON(e) : undefined;
            });
        }
        else {
            obj.service = [];
        }
        if (message.alsoKnownAs) {
            obj.alsoKnownAs = message.alsoKnownAs.map(function (e) { return e; });
        }
        else {
            obj.alsoKnownAs = [];
        }
        return obj;
    },
    fromPartial: function (object) {
        var _a, _b, _c, _d, _e, _f, _g, _h, _j, _k, _l;
        var message = createBaseDidDocument();
        message.context = ((_a = object.context) === null || _a === void 0 ? void 0 : _a.map(function (e) { return e; })) || [];
        message.id = (_b = object.id) !== null && _b !== void 0 ? _b : "";
        message.controller = ((_c = object.controller) === null || _c === void 0 ? void 0 : _c.map(function (e) { return e; })) || [];
        message.verificationMethod =
            ((_d = object.verificationMethod) === null || _d === void 0 ? void 0 : _d.map(function (e) {
                return exports.VerificationMethod.fromPartial(e);
            })) || [];
        message.authentication = ((_e = object.authentication) === null || _e === void 0 ? void 0 : _e.map(function (e) { return e; })) || [];
        message.assertionMethod = ((_f = object.assertionMethod) === null || _f === void 0 ? void 0 : _f.map(function (e) { return e; })) || [];
        message.capabilityInvocation =
            ((_g = object.capabilityInvocation) === null || _g === void 0 ? void 0 : _g.map(function (e) { return e; })) || [];
        message.capabilityDelegation =
            ((_h = object.capabilityDelegation) === null || _h === void 0 ? void 0 : _h.map(function (e) { return e; })) || [];
        message.keyAgreement = ((_j = object.keyAgreement) === null || _j === void 0 ? void 0 : _j.map(function (e) { return e; })) || [];
        message.service = ((_k = object.service) === null || _k === void 0 ? void 0 : _k.map(function (e) { return exports.Service.fromPartial(e); })) || [];
        message.alsoKnownAs = ((_l = object.alsoKnownAs) === null || _l === void 0 ? void 0 : _l.map(function (e) { return e; })) || [];
        return message;
    }
};
function createBaseService() {
    return { id: "", type: "", serviceEndpoint: "" };
}
exports.Service = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.type !== "") {
            writer.uint32(18).string(message.type);
        }
        if (message.serviceEndpoint !== "") {
            writer.uint32(26).string(message.serviceEndpoint);
        }
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBaseService();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.type = reader.string();
                    break;
                case 3:
                    message.serviceEndpoint = reader.string();
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
            type: isSet(object.type) ? String(object.type) : "",
            serviceEndpoint: isSet(object.serviceEndpoint)
                ? String(object.serviceEndpoint)
                : ""
        };
    },
    toJSON: function (message) {
        var obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.type !== undefined && (obj.type = message.type);
        message.serviceEndpoint !== undefined &&
            (obj.serviceEndpoint = message.serviceEndpoint);
        return obj;
    },
    fromPartial: function (object) {
        var _a, _b, _c;
        var message = createBaseService();
        message.id = (_a = object.id) !== null && _a !== void 0 ? _a : "";
        message.type = (_b = object.type) !== null && _b !== void 0 ? _b : "";
        message.serviceEndpoint = (_c = object.serviceEndpoint) !== null && _c !== void 0 ? _c : "";
        return message;
    }
};
function createBaseVerificationMethod() {
    return {
        id: "",
        type: "",
        controller: "",
        publicKeyHex: "",
        publicKeyBase58: "",
        blockchainAccountId: ""
    };
}
exports.VerificationMethod = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_1["default"].Writer.create(); }
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.type !== "") {
            writer.uint32(18).string(message.type);
        }
        if (message.controller !== "") {
            writer.uint32(26).string(message.controller);
        }
        if (message.publicKeyHex !== "") {
            writer.uint32(34).string(message.publicKeyHex);
        }
        if (message.publicKeyBase58 !== "") {
            writer.uint32(42).string(message.publicKeyBase58);
        }
        if (message.blockchainAccountId !== "") {
            writer.uint32(50).string(message.blockchainAccountId);
        }
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_1["default"].Reader ? input : new minimal_1["default"].Reader(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBaseVerificationMethod();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.type = reader.string();
                    break;
                case 3:
                    message.controller = reader.string();
                    break;
                case 4:
                    message.publicKeyHex = reader.string();
                    break;
                case 5:
                    message.publicKeyBase58 = reader.string();
                    break;
                case 6:
                    message.blockchainAccountId = reader.string();
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
            type: isSet(object.type) ? String(object.type) : "",
            controller: isSet(object.controller) ? String(object.controller) : "",
            publicKeyHex: isSet(object.publicKeyHex)
                ? String(object.publicKeyHex)
                : "",
            publicKeyBase58: isSet(object.publicKeyBase58)
                ? String(object.publicKeyBase58)
                : "",
            blockchainAccountId: isSet(object.blockchainAccountId)
                ? String(object.blockchainAccountId)
                : ""
        };
    },
    toJSON: function (message) {
        var obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.type !== undefined && (obj.type = message.type);
        message.controller !== undefined && (obj.controller = message.controller);
        message.publicKeyHex !== undefined &&
            (obj.publicKeyHex = message.publicKeyHex);
        message.publicKeyBase58 !== undefined &&
            (obj.publicKeyBase58 = message.publicKeyBase58);
        message.blockchainAccountId !== undefined &&
            (obj.blockchainAccountId = message.blockchainAccountId);
        return obj;
    },
    fromPartial: function (object) {
        var _a, _b, _c, _d, _e, _f;
        var message = createBaseVerificationMethod();
        message.id = (_a = object.id) !== null && _a !== void 0 ? _a : "";
        message.type = (_b = object.type) !== null && _b !== void 0 ? _b : "";
        message.controller = (_c = object.controller) !== null && _c !== void 0 ? _c : "";
        message.publicKeyHex = (_d = object.publicKeyHex) !== null && _d !== void 0 ? _d : "";
        message.publicKeyBase58 = (_e = object.publicKeyBase58) !== null && _e !== void 0 ? _e : "";
        message.blockchainAccountId = (_f = object.blockchainAccountId) !== null && _f !== void 0 ? _f : "";
        return message;
    }
};
if (minimal_1["default"].util.Long !== long_1["default"]) {
    minimal_1["default"].util.Long = long_1["default"];
    minimal_1["default"].configure();
}
function isSet(value) {
    return value !== null && value !== undefined;
}
