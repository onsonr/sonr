/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "sonrio.sonr.registry";
/** NetworkType is the type of network the DID is on. */
export var NetworkType;
(function (NetworkType) {
    /** NETWORK_TYPE_UNSPECIFIED - Unspecified is the default value. Gets converted to "did:sonr:". */
    NetworkType[NetworkType["NETWORK_TYPE_UNSPECIFIED"] = 0] = "NETWORK_TYPE_UNSPECIFIED";
    /** NETWORK_TYPE_MAINNET - Mainnet is the main network. It prefix is "did:sonr:" or "did:sonr:mainnet:". */
    NetworkType[NetworkType["NETWORK_TYPE_MAINNET"] = 1] = "NETWORK_TYPE_MAINNET";
    /** NETWORK_TYPE_TESTNET - Testnet is the deployed test network. It's prefix is "did:sonr:testnet:". */
    NetworkType[NetworkType["NETWORK_TYPE_TESTNET"] = 2] = "NETWORK_TYPE_TESTNET";
    /** NETWORK_TYPE_DEVNET - Devnet is the localhost test network. It's prefix is "did:sonr:devnet:". */
    NetworkType[NetworkType["NETWORK_TYPE_DEVNET"] = 3] = "NETWORK_TYPE_DEVNET";
    NetworkType[NetworkType["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(NetworkType || (NetworkType = {}));
export function networkTypeFromJSON(object) {
    switch (object) {
        case 0:
        case "NETWORK_TYPE_UNSPECIFIED":
            return NetworkType.NETWORK_TYPE_UNSPECIFIED;
        case 1:
        case "NETWORK_TYPE_MAINNET":
            return NetworkType.NETWORK_TYPE_MAINNET;
        case 2:
        case "NETWORK_TYPE_TESTNET":
            return NetworkType.NETWORK_TYPE_TESTNET;
        case 3:
        case "NETWORK_TYPE_DEVNET":
            return NetworkType.NETWORK_TYPE_DEVNET;
        case -1:
        case "UNRECOGNIZED":
        default:
            return NetworkType.UNRECOGNIZED;
    }
}
export function networkTypeToJSON(object) {
    switch (object) {
        case NetworkType.NETWORK_TYPE_UNSPECIFIED:
            return "NETWORK_TYPE_UNSPECIFIED";
        case NetworkType.NETWORK_TYPE_MAINNET:
            return "NETWORK_TYPE_MAINNET";
        case NetworkType.NETWORK_TYPE_TESTNET:
            return "NETWORK_TYPE_TESTNET";
        case NetworkType.NETWORK_TYPE_DEVNET:
            return "NETWORK_TYPE_DEVNET";
        default:
            return "UNKNOWN";
    }
}
/** ServiceProtocol are core modules that can be installed on custom services on the Sonr network. */
export var ServiceProtocol;
(function (ServiceProtocol) {
    /** SERVICE_PROTOCOL_UNSPECIFIED - SERVICE_PROTOCOL_UNSPECIFIED is the default value. */
    ServiceProtocol[ServiceProtocol["SERVICE_PROTOCOL_UNSPECIFIED"] = 0] = "SERVICE_PROTOCOL_UNSPECIFIED";
    /** SERVICE_PROTOCOL_BUCKETS - SERVICE_PROTOCOL_BUCKETS is the module that provides the ability to store and retrieve data. */
    ServiceProtocol[ServiceProtocol["SERVICE_PROTOCOL_BUCKETS"] = 1] = "SERVICE_PROTOCOL_BUCKETS";
    /** SERVICE_PROTOCOL_CHANNEL - SERVICE_PROTOCOL_CHANNEL is the module that provides the ability to communicate with other services. */
    ServiceProtocol[ServiceProtocol["SERVICE_PROTOCOL_CHANNEL"] = 2] = "SERVICE_PROTOCOL_CHANNEL";
    /** SERVICE_PROTOCOL_OBJECTS - SERVICE_PROTOCOL_OBJECTS is the module that provides the ability to create new schemas for data on the network. */
    ServiceProtocol[ServiceProtocol["SERVICE_PROTOCOL_OBJECTS"] = 3] = "SERVICE_PROTOCOL_OBJECTS";
    /** SERVICE_PROTOCOL_FUNCTIONS - SERVICE_PROTOCOL_FUNCTIONS is the module that provides the ability to create new functions for data on the network. */
    ServiceProtocol[ServiceProtocol["SERVICE_PROTOCOL_FUNCTIONS"] = 4] = "SERVICE_PROTOCOL_FUNCTIONS";
    ServiceProtocol[ServiceProtocol["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(ServiceProtocol || (ServiceProtocol = {}));
export function serviceProtocolFromJSON(object) {
    switch (object) {
        case 0:
        case "SERVICE_PROTOCOL_UNSPECIFIED":
            return ServiceProtocol.SERVICE_PROTOCOL_UNSPECIFIED;
        case 1:
        case "SERVICE_PROTOCOL_BUCKETS":
            return ServiceProtocol.SERVICE_PROTOCOL_BUCKETS;
        case 2:
        case "SERVICE_PROTOCOL_CHANNEL":
            return ServiceProtocol.SERVICE_PROTOCOL_CHANNEL;
        case 3:
        case "SERVICE_PROTOCOL_OBJECTS":
            return ServiceProtocol.SERVICE_PROTOCOL_OBJECTS;
        case 4:
        case "SERVICE_PROTOCOL_FUNCTIONS":
            return ServiceProtocol.SERVICE_PROTOCOL_FUNCTIONS;
        case -1:
        case "UNRECOGNIZED":
        default:
            return ServiceProtocol.UNRECOGNIZED;
    }
}
export function serviceProtocolToJSON(object) {
    switch (object) {
        case ServiceProtocol.SERVICE_PROTOCOL_UNSPECIFIED:
            return "SERVICE_PROTOCOL_UNSPECIFIED";
        case ServiceProtocol.SERVICE_PROTOCOL_BUCKETS:
            return "SERVICE_PROTOCOL_BUCKETS";
        case ServiceProtocol.SERVICE_PROTOCOL_CHANNEL:
            return "SERVICE_PROTOCOL_CHANNEL";
        case ServiceProtocol.SERVICE_PROTOCOL_OBJECTS:
            return "SERVICE_PROTOCOL_OBJECTS";
        case ServiceProtocol.SERVICE_PROTOCOL_FUNCTIONS:
            return "SERVICE_PROTOCOL_FUNCTIONS";
        default:
            return "UNKNOWN";
    }
}
/** ServiceType is the type of service that is being registered. */
export var ServiceType;
(function (ServiceType) {
    /** SERVICE_TYPE_UNSPECIFIED - SERVICE_TYPE_UNSPECIFIED is the default value. */
    ServiceType[ServiceType["SERVICE_TYPE_UNSPECIFIED"] = 0] = "SERVICE_TYPE_UNSPECIFIED";
    /** SERVICE_TYPE_DID_COMM_MESSAGING - SERVICE_TYPE_APPLICATION is the type of service that is a DApp. */
    ServiceType[ServiceType["SERVICE_TYPE_DID_COMM_MESSAGING"] = 1] = "SERVICE_TYPE_DID_COMM_MESSAGING";
    /** SERVICE_TYPE_LINKED_DOMAINS - SERVICE_TYPE_SERVICE is the type of service that is a service. */
    ServiceType[ServiceType["SERVICE_TYPE_LINKED_DOMAINS"] = 2] = "SERVICE_TYPE_LINKED_DOMAINS";
    /** SERVICE_TYPE_SONR - SERVICE_TYPE_SONR is the type of service that is a DApp. */
    ServiceType[ServiceType["SERVICE_TYPE_SONR"] = 3] = "SERVICE_TYPE_SONR";
    ServiceType[ServiceType["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(ServiceType || (ServiceType = {}));
export function serviceTypeFromJSON(object) {
    switch (object) {
        case 0:
        case "SERVICE_TYPE_UNSPECIFIED":
            return ServiceType.SERVICE_TYPE_UNSPECIFIED;
        case 1:
        case "SERVICE_TYPE_DID_COMM_MESSAGING":
            return ServiceType.SERVICE_TYPE_DID_COMM_MESSAGING;
        case 2:
        case "SERVICE_TYPE_LINKED_DOMAINS":
            return ServiceType.SERVICE_TYPE_LINKED_DOMAINS;
        case 3:
        case "SERVICE_TYPE_SONR":
            return ServiceType.SERVICE_TYPE_SONR;
        case -1:
        case "UNRECOGNIZED":
        default:
            return ServiceType.UNRECOGNIZED;
    }
}
export function serviceTypeToJSON(object) {
    switch (object) {
        case ServiceType.SERVICE_TYPE_UNSPECIFIED:
            return "SERVICE_TYPE_UNSPECIFIED";
        case ServiceType.SERVICE_TYPE_DID_COMM_MESSAGING:
            return "SERVICE_TYPE_DID_COMM_MESSAGING";
        case ServiceType.SERVICE_TYPE_LINKED_DOMAINS:
            return "SERVICE_TYPE_LINKED_DOMAINS";
        case ServiceType.SERVICE_TYPE_SONR:
            return "SERVICE_TYPE_SONR";
        default:
            return "UNKNOWN";
    }
}
export var VerificationMethod_Type;
(function (VerificationMethod_Type) {
    /** TYPE_UNSPECIFIED - TYPE_UNSPECIFIED is the default value. */
    VerificationMethod_Type[VerificationMethod_Type["TYPE_UNSPECIFIED"] = 0] = "TYPE_UNSPECIFIED";
    /** TYPE_ECDSA_SECP256K1 - TYPE_ECDSA_SECP256K1 represents the Ed25519VerificationKey2018 key type. */
    VerificationMethod_Type[VerificationMethod_Type["TYPE_ECDSA_SECP256K1"] = 1] = "TYPE_ECDSA_SECP256K1";
    /** TYPE_X25519 - TYPE_X25519 represents the X25519KeyAgreementKey2019 key type. */
    VerificationMethod_Type[VerificationMethod_Type["TYPE_X25519"] = 2] = "TYPE_X25519";
    /** TYPE_ED25519 - TYPE_ED25519 represents the Ed25519VerificationKey2018 key type. */
    VerificationMethod_Type[VerificationMethod_Type["TYPE_ED25519"] = 3] = "TYPE_ED25519";
    /** TYPE_BLS_12381_G1 - TYPE_BLS_12381_G1 represents the Bls12381G1Key2020 key type */
    VerificationMethod_Type[VerificationMethod_Type["TYPE_BLS_12381_G1"] = 4] = "TYPE_BLS_12381_G1";
    /** TYPE_BLS_12381_G2 - TYPE_BLS_12381_G2 represents the Bls12381G2Key2020 key type */
    VerificationMethod_Type[VerificationMethod_Type["TYPE_BLS_12381_G2"] = 5] = "TYPE_BLS_12381_G2";
    /** TYPE_RSA - TYPE_RSA represents the RsaVerificationKey2018 key type. */
    VerificationMethod_Type[VerificationMethod_Type["TYPE_RSA"] = 6] = "TYPE_RSA";
    /** TYPE_VERIFIABLE_CONDITION - TYPE_VERIFIABLE_CONDITION represents the VerifiableCondition2021 key type. */
    VerificationMethod_Type[VerificationMethod_Type["TYPE_VERIFIABLE_CONDITION"] = 7] = "TYPE_VERIFIABLE_CONDITION";
    VerificationMethod_Type[VerificationMethod_Type["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(VerificationMethod_Type || (VerificationMethod_Type = {}));
export function verificationMethod_TypeFromJSON(object) {
    switch (object) {
        case 0:
        case "TYPE_UNSPECIFIED":
            return VerificationMethod_Type.TYPE_UNSPECIFIED;
        case 1:
        case "TYPE_ECDSA_SECP256K1":
            return VerificationMethod_Type.TYPE_ECDSA_SECP256K1;
        case 2:
        case "TYPE_X25519":
            return VerificationMethod_Type.TYPE_X25519;
        case 3:
        case "TYPE_ED25519":
            return VerificationMethod_Type.TYPE_ED25519;
        case 4:
        case "TYPE_BLS_12381_G1":
            return VerificationMethod_Type.TYPE_BLS_12381_G1;
        case 5:
        case "TYPE_BLS_12381_G2":
            return VerificationMethod_Type.TYPE_BLS_12381_G2;
        case 6:
        case "TYPE_RSA":
            return VerificationMethod_Type.TYPE_RSA;
        case 7:
        case "TYPE_VERIFIABLE_CONDITION":
            return VerificationMethod_Type.TYPE_VERIFIABLE_CONDITION;
        case -1:
        case "UNRECOGNIZED":
        default:
            return VerificationMethod_Type.UNRECOGNIZED;
    }
}
export function verificationMethod_TypeToJSON(object) {
    switch (object) {
        case VerificationMethod_Type.TYPE_UNSPECIFIED:
            return "TYPE_UNSPECIFIED";
        case VerificationMethod_Type.TYPE_ECDSA_SECP256K1:
            return "TYPE_ECDSA_SECP256K1";
        case VerificationMethod_Type.TYPE_X25519:
            return "TYPE_X25519";
        case VerificationMethod_Type.TYPE_ED25519:
            return "TYPE_ED25519";
        case VerificationMethod_Type.TYPE_BLS_12381_G1:
            return "TYPE_BLS_12381_G1";
        case VerificationMethod_Type.TYPE_BLS_12381_G2:
            return "TYPE_BLS_12381_G2";
        case VerificationMethod_Type.TYPE_RSA:
            return "TYPE_RSA";
        case VerificationMethod_Type.TYPE_VERIFIABLE_CONDITION:
            return "TYPE_VERIFIABLE_CONDITION";
        default:
            return "UNKNOWN";
    }
}
const baseDid = {
    method: "",
    network: "",
    id: "",
    paths: "",
    query: "",
    fragment: "",
};
export const Did = {
    encode(message, writer = Writer.create()) {
        if (message.method !== "") {
            writer.uint32(10).string(message.method);
        }
        if (message.network !== "") {
            writer.uint32(18).string(message.network);
        }
        if (message.id !== "") {
            writer.uint32(26).string(message.id);
        }
        for (const v of message.paths) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseDid };
        message.paths = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
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
    fromJSON(object) {
        const message = { ...baseDid };
        message.paths = [];
        if (object.method !== undefined && object.method !== null) {
            message.method = String(object.method);
        }
        else {
            message.method = "";
        }
        if (object.network !== undefined && object.network !== null) {
            message.network = String(object.network);
        }
        else {
            message.network = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.paths !== undefined && object.paths !== null) {
            for (const e of object.paths) {
                message.paths.push(String(e));
            }
        }
        if (object.query !== undefined && object.query !== null) {
            message.query = String(object.query);
        }
        else {
            message.query = "";
        }
        if (object.fragment !== undefined && object.fragment !== null) {
            message.fragment = String(object.fragment);
        }
        else {
            message.fragment = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.method !== undefined && (obj.method = message.method);
        message.network !== undefined && (obj.network = message.network);
        message.id !== undefined && (obj.id = message.id);
        if (message.paths) {
            obj.paths = message.paths.map((e) => e);
        }
        else {
            obj.paths = [];
        }
        message.query !== undefined && (obj.query = message.query);
        message.fragment !== undefined && (obj.fragment = message.fragment);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseDid };
        message.paths = [];
        if (object.method !== undefined && object.method !== null) {
            message.method = object.method;
        }
        else {
            message.method = "";
        }
        if (object.network !== undefined && object.network !== null) {
            message.network = object.network;
        }
        else {
            message.network = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.paths !== undefined && object.paths !== null) {
            for (const e of object.paths) {
                message.paths.push(e);
            }
        }
        if (object.query !== undefined && object.query !== null) {
            message.query = object.query;
        }
        else {
            message.query = "";
        }
        if (object.fragment !== undefined && object.fragment !== null) {
            message.fragment = object.fragment;
        }
        else {
            message.fragment = "";
        }
        return message;
    },
};
const baseDidDocument = {
    context: "",
    id: "",
    controller: "",
    authentication: "",
    assertionMethod: "",
    capabilityInvocation: "",
    capabilityDelegation: "",
    keyAgreement: "",
    alsoKnownAs: "",
};
export const DidDocument = {
    encode(message, writer = Writer.create()) {
        for (const v of message.context) {
            writer.uint32(10).string(v);
        }
        if (message.id !== "") {
            writer.uint32(18).string(message.id);
        }
        for (const v of message.controller) {
            writer.uint32(26).string(v);
        }
        for (const v of message.verificationMethod) {
            VerificationMethod.encode(v, writer.uint32(34).fork()).ldelim();
        }
        for (const v of message.authentication) {
            writer.uint32(42).string(v);
        }
        for (const v of message.assertionMethod) {
            writer.uint32(50).string(v);
        }
        for (const v of message.capabilityInvocation) {
            writer.uint32(58).string(v);
        }
        for (const v of message.capabilityDelegation) {
            writer.uint32(66).string(v);
        }
        for (const v of message.keyAgreement) {
            writer.uint32(74).string(v);
        }
        for (const v of message.service) {
            Service.encode(v, writer.uint32(82).fork()).ldelim();
        }
        for (const v of message.alsoKnownAs) {
            writer.uint32(90).string(v);
        }
        Object.entries(message.metadata).forEach(([key, value]) => {
            DidDocument_MetadataEntry.encode({ key: key, value }, writer.uint32(98).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseDidDocument };
        message.context = [];
        message.controller = [];
        message.verificationMethod = [];
        message.authentication = [];
        message.assertionMethod = [];
        message.capabilityInvocation = [];
        message.capabilityDelegation = [];
        message.keyAgreement = [];
        message.service = [];
        message.alsoKnownAs = [];
        message.metadata = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
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
                    message.verificationMethod.push(VerificationMethod.decode(reader, reader.uint32()));
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
                    message.service.push(Service.decode(reader, reader.uint32()));
                    break;
                case 11:
                    message.alsoKnownAs.push(reader.string());
                    break;
                case 12:
                    const entry12 = DidDocument_MetadataEntry.decode(reader, reader.uint32());
                    if (entry12.value !== undefined) {
                        message.metadata[entry12.key] = entry12.value;
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
        const message = { ...baseDidDocument };
        message.context = [];
        message.controller = [];
        message.verificationMethod = [];
        message.authentication = [];
        message.assertionMethod = [];
        message.capabilityInvocation = [];
        message.capabilityDelegation = [];
        message.keyAgreement = [];
        message.service = [];
        message.alsoKnownAs = [];
        message.metadata = {};
        if (object.context !== undefined && object.context !== null) {
            for (const e of object.context) {
                message.context.push(String(e));
            }
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.controller !== undefined && object.controller !== null) {
            for (const e of object.controller) {
                message.controller.push(String(e));
            }
        }
        if (object.verificationMethod !== undefined &&
            object.verificationMethod !== null) {
            for (const e of object.verificationMethod) {
                message.verificationMethod.push(VerificationMethod.fromJSON(e));
            }
        }
        if (object.authentication !== undefined && object.authentication !== null) {
            for (const e of object.authentication) {
                message.authentication.push(String(e));
            }
        }
        if (object.assertionMethod !== undefined &&
            object.assertionMethod !== null) {
            for (const e of object.assertionMethod) {
                message.assertionMethod.push(String(e));
            }
        }
        if (object.capabilityInvocation !== undefined &&
            object.capabilityInvocation !== null) {
            for (const e of object.capabilityInvocation) {
                message.capabilityInvocation.push(String(e));
            }
        }
        if (object.capabilityDelegation !== undefined &&
            object.capabilityDelegation !== null) {
            for (const e of object.capabilityDelegation) {
                message.capabilityDelegation.push(String(e));
            }
        }
        if (object.keyAgreement !== undefined && object.keyAgreement !== null) {
            for (const e of object.keyAgreement) {
                message.keyAgreement.push(String(e));
            }
        }
        if (object.service !== undefined && object.service !== null) {
            for (const e of object.service) {
                message.service.push(Service.fromJSON(e));
            }
        }
        if (object.alsoKnownAs !== undefined && object.alsoKnownAs !== null) {
            for (const e of object.alsoKnownAs) {
                message.alsoKnownAs.push(String(e));
            }
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
        if (message.context) {
            obj.context = message.context.map((e) => e);
        }
        else {
            obj.context = [];
        }
        message.id !== undefined && (obj.id = message.id);
        if (message.controller) {
            obj.controller = message.controller.map((e) => e);
        }
        else {
            obj.controller = [];
        }
        if (message.verificationMethod) {
            obj.verificationMethod = message.verificationMethod.map((e) => e ? VerificationMethod.toJSON(e) : undefined);
        }
        else {
            obj.verificationMethod = [];
        }
        if (message.authentication) {
            obj.authentication = message.authentication.map((e) => e);
        }
        else {
            obj.authentication = [];
        }
        if (message.assertionMethod) {
            obj.assertionMethod = message.assertionMethod.map((e) => e);
        }
        else {
            obj.assertionMethod = [];
        }
        if (message.capabilityInvocation) {
            obj.capabilityInvocation = message.capabilityInvocation.map((e) => e);
        }
        else {
            obj.capabilityInvocation = [];
        }
        if (message.capabilityDelegation) {
            obj.capabilityDelegation = message.capabilityDelegation.map((e) => e);
        }
        else {
            obj.capabilityDelegation = [];
        }
        if (message.keyAgreement) {
            obj.keyAgreement = message.keyAgreement.map((e) => e);
        }
        else {
            obj.keyAgreement = [];
        }
        if (message.service) {
            obj.service = message.service.map((e) => e ? Service.toJSON(e) : undefined);
        }
        else {
            obj.service = [];
        }
        if (message.alsoKnownAs) {
            obj.alsoKnownAs = message.alsoKnownAs.map((e) => e);
        }
        else {
            obj.alsoKnownAs = [];
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
        const message = { ...baseDidDocument };
        message.context = [];
        message.controller = [];
        message.verificationMethod = [];
        message.authentication = [];
        message.assertionMethod = [];
        message.capabilityInvocation = [];
        message.capabilityDelegation = [];
        message.keyAgreement = [];
        message.service = [];
        message.alsoKnownAs = [];
        message.metadata = {};
        if (object.context !== undefined && object.context !== null) {
            for (const e of object.context) {
                message.context.push(e);
            }
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.controller !== undefined && object.controller !== null) {
            for (const e of object.controller) {
                message.controller.push(e);
            }
        }
        if (object.verificationMethod !== undefined &&
            object.verificationMethod !== null) {
            for (const e of object.verificationMethod) {
                message.verificationMethod.push(VerificationMethod.fromPartial(e));
            }
        }
        if (object.authentication !== undefined && object.authentication !== null) {
            for (const e of object.authentication) {
                message.authentication.push(e);
            }
        }
        if (object.assertionMethod !== undefined &&
            object.assertionMethod !== null) {
            for (const e of object.assertionMethod) {
                message.assertionMethod.push(e);
            }
        }
        if (object.capabilityInvocation !== undefined &&
            object.capabilityInvocation !== null) {
            for (const e of object.capabilityInvocation) {
                message.capabilityInvocation.push(e);
            }
        }
        if (object.capabilityDelegation !== undefined &&
            object.capabilityDelegation !== null) {
            for (const e of object.capabilityDelegation) {
                message.capabilityDelegation.push(e);
            }
        }
        if (object.keyAgreement !== undefined && object.keyAgreement !== null) {
            for (const e of object.keyAgreement) {
                message.keyAgreement.push(e);
            }
        }
        if (object.service !== undefined && object.service !== null) {
            for (const e of object.service) {
                message.service.push(Service.fromPartial(e));
            }
        }
        if (object.alsoKnownAs !== undefined && object.alsoKnownAs !== null) {
            for (const e of object.alsoKnownAs) {
                message.alsoKnownAs.push(e);
            }
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
const baseDidDocument_MetadataEntry = { key: "", value: "" };
export const DidDocument_MetadataEntry = {
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
            ...baseDidDocument_MetadataEntry,
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
            ...baseDidDocument_MetadataEntry,
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
            ...baseDidDocument_MetadataEntry,
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
const baseService = { id: "", type: 0 };
export const Service = {
    encode(message, writer = Writer.create()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.type !== 0) {
            writer.uint32(16).int32(message.type);
        }
        if (message.serviceEndpoint !== undefined) {
            ServiceEndpoint.encode(message.serviceEndpoint, writer.uint32(26).fork()).ldelim();
        }
        Object.entries(message.metadata).forEach(([key, value]) => {
            Service_MetadataEntry.encode({ key: key, value }, writer.uint32(34).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseService };
        message.metadata = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.type = reader.int32();
                    break;
                case 3:
                    message.serviceEndpoint = ServiceEndpoint.decode(reader, reader.uint32());
                    break;
                case 4:
                    const entry4 = Service_MetadataEntry.decode(reader, reader.uint32());
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
        const message = { ...baseService };
        message.metadata = {};
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.type !== undefined && object.type !== null) {
            message.type = serviceTypeFromJSON(object.type);
        }
        else {
            message.type = 0;
        }
        if (object.serviceEndpoint !== undefined &&
            object.serviceEndpoint !== null) {
            message.serviceEndpoint = ServiceEndpoint.fromJSON(object.serviceEndpoint);
        }
        else {
            message.serviceEndpoint = undefined;
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
        message.id !== undefined && (obj.id = message.id);
        message.type !== undefined && (obj.type = serviceTypeToJSON(message.type));
        message.serviceEndpoint !== undefined &&
            (obj.serviceEndpoint = message.serviceEndpoint
                ? ServiceEndpoint.toJSON(message.serviceEndpoint)
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
        const message = { ...baseService };
        message.metadata = {};
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.type !== undefined && object.type !== null) {
            message.type = object.type;
        }
        else {
            message.type = 0;
        }
        if (object.serviceEndpoint !== undefined &&
            object.serviceEndpoint !== null) {
            message.serviceEndpoint = ServiceEndpoint.fromPartial(object.serviceEndpoint);
        }
        else {
            message.serviceEndpoint = undefined;
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
const baseService_MetadataEntry = { key: "", value: "" };
export const Service_MetadataEntry = {
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
        const message = { ...baseService_MetadataEntry };
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
        const message = { ...baseService_MetadataEntry };
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
        const message = { ...baseService_MetadataEntry };
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
const baseServiceEndpoint = {
    transportType: "",
    network: "",
    supportedProtocols: 0,
};
export const ServiceEndpoint = {
    encode(message, writer = Writer.create()) {
        if (message.transportType !== "") {
            writer.uint32(10).string(message.transportType);
        }
        if (message.network !== "") {
            writer.uint32(18).string(message.network);
        }
        writer.uint32(26).fork();
        for (const v of message.supportedProtocols) {
            writer.int32(v);
        }
        writer.ldelim();
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseServiceEndpoint };
        message.supportedProtocols = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.transportType = reader.string();
                    break;
                case 2:
                    message.network = reader.string();
                    break;
                case 3:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.supportedProtocols.push(reader.int32());
                        }
                    }
                    else {
                        message.supportedProtocols.push(reader.int32());
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
        const message = { ...baseServiceEndpoint };
        message.supportedProtocols = [];
        if (object.transportType !== undefined && object.transportType !== null) {
            message.transportType = String(object.transportType);
        }
        else {
            message.transportType = "";
        }
        if (object.network !== undefined && object.network !== null) {
            message.network = String(object.network);
        }
        else {
            message.network = "";
        }
        if (object.supportedProtocols !== undefined &&
            object.supportedProtocols !== null) {
            for (const e of object.supportedProtocols) {
                message.supportedProtocols.push(serviceProtocolFromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.transportType !== undefined &&
            (obj.transportType = message.transportType);
        message.network !== undefined && (obj.network = message.network);
        if (message.supportedProtocols) {
            obj.supportedProtocols = message.supportedProtocols.map((e) => serviceProtocolToJSON(e));
        }
        else {
            obj.supportedProtocols = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseServiceEndpoint };
        message.supportedProtocols = [];
        if (object.transportType !== undefined && object.transportType !== null) {
            message.transportType = object.transportType;
        }
        else {
            message.transportType = "";
        }
        if (object.network !== undefined && object.network !== null) {
            message.network = object.network;
        }
        else {
            message.network = "";
        }
        if (object.supportedProtocols !== undefined &&
            object.supportedProtocols !== null) {
            for (const e of object.supportedProtocols) {
                message.supportedProtocols.push(e);
            }
        }
        return message;
    },
};
const baseVerificationMethod = {
    id: "",
    type: 0,
    controller: "",
    publicKeyHex: "",
    publicKeyBase58: "",
    blockchainAccountId: "",
};
export const VerificationMethod = {
    encode(message, writer = Writer.create()) {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.type !== 0) {
            writer.uint32(16).int32(message.type);
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseVerificationMethod };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = reader.string();
                    break;
                case 2:
                    message.type = reader.int32();
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
    fromJSON(object) {
        const message = { ...baseVerificationMethod };
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.type !== undefined && object.type !== null) {
            message.type = verificationMethod_TypeFromJSON(object.type);
        }
        else {
            message.type = 0;
        }
        if (object.controller !== undefined && object.controller !== null) {
            message.controller = String(object.controller);
        }
        else {
            message.controller = "";
        }
        if (object.publicKeyHex !== undefined && object.publicKeyHex !== null) {
            message.publicKeyHex = String(object.publicKeyHex);
        }
        else {
            message.publicKeyHex = "";
        }
        if (object.publicKeyBase58 !== undefined &&
            object.publicKeyBase58 !== null) {
            message.publicKeyBase58 = String(object.publicKeyBase58);
        }
        else {
            message.publicKeyBase58 = "";
        }
        if (object.blockchainAccountId !== undefined &&
            object.blockchainAccountId !== null) {
            message.blockchainAccountId = String(object.blockchainAccountId);
        }
        else {
            message.blockchainAccountId = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.type !== undefined &&
            (obj.type = verificationMethod_TypeToJSON(message.type));
        message.controller !== undefined && (obj.controller = message.controller);
        message.publicKeyHex !== undefined &&
            (obj.publicKeyHex = message.publicKeyHex);
        message.publicKeyBase58 !== undefined &&
            (obj.publicKeyBase58 = message.publicKeyBase58);
        message.blockchainAccountId !== undefined &&
            (obj.blockchainAccountId = message.blockchainAccountId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseVerificationMethod };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.type !== undefined && object.type !== null) {
            message.type = object.type;
        }
        else {
            message.type = 0;
        }
        if (object.controller !== undefined && object.controller !== null) {
            message.controller = object.controller;
        }
        else {
            message.controller = "";
        }
        if (object.publicKeyHex !== undefined && object.publicKeyHex !== null) {
            message.publicKeyHex = object.publicKeyHex;
        }
        else {
            message.publicKeyHex = "";
        }
        if (object.publicKeyBase58 !== undefined &&
            object.publicKeyBase58 !== null) {
            message.publicKeyBase58 = object.publicKeyBase58;
        }
        else {
            message.publicKeyBase58 = "";
        }
        if (object.blockchainAccountId !== undefined &&
            object.blockchainAccountId !== null) {
            message.blockchainAccountId = object.blockchainAccountId;
        }
        else {
            message.blockchainAccountId = "";
        }
        return message;
    },
};
