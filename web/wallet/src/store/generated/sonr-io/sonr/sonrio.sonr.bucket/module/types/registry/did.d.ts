import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "sonrio.sonr.registry";
/** NetworkType is the type of network the DID is on. */
export declare enum NetworkType {
    /** NETWORK_TYPE_UNSPECIFIED - Unspecified is the default value. Gets converted to "did:sonr:". */
    NETWORK_TYPE_UNSPECIFIED = 0,
    /** NETWORK_TYPE_MAINNET - Mainnet is the main network. It prefix is "did:sonr:" or "did:sonr:mainnet:". */
    NETWORK_TYPE_MAINNET = 1,
    /** NETWORK_TYPE_TESTNET - Testnet is the deployed test network. It's prefix is "did:sonr:testnet:". */
    NETWORK_TYPE_TESTNET = 2,
    /** NETWORK_TYPE_DEVNET - Devnet is the localhost test network. It's prefix is "did:sonr:devnet:". */
    NETWORK_TYPE_DEVNET = 3,
    UNRECOGNIZED = -1
}
export declare function networkTypeFromJSON(object: any): NetworkType;
export declare function networkTypeToJSON(object: NetworkType): string;
/** ServiceProtocol are core modules that can be installed on custom services on the Sonr network. */
export declare enum ServiceProtocol {
    /** SERVICE_PROTOCOL_UNSPECIFIED - SERVICE_PROTOCOL_UNSPECIFIED is the default value. */
    SERVICE_PROTOCOL_UNSPECIFIED = 0,
    /** SERVICE_PROTOCOL_BUCKETS - SERVICE_PROTOCOL_BUCKETS is the module that provides the ability to store and retrieve data. */
    SERVICE_PROTOCOL_BUCKETS = 1,
    /** SERVICE_PROTOCOL_CHANNEL - SERVICE_PROTOCOL_CHANNEL is the module that provides the ability to communicate with other services. */
    SERVICE_PROTOCOL_CHANNEL = 2,
    /** SERVICE_PROTOCOL_OBJECTS - SERVICE_PROTOCOL_OBJECTS is the module that provides the ability to create new schemas for data on the network. */
    SERVICE_PROTOCOL_OBJECTS = 3,
    /** SERVICE_PROTOCOL_FUNCTIONS - SERVICE_PROTOCOL_FUNCTIONS is the module that provides the ability to create new functions for data on the network. */
    SERVICE_PROTOCOL_FUNCTIONS = 4,
    UNRECOGNIZED = -1
}
export declare function serviceProtocolFromJSON(object: any): ServiceProtocol;
export declare function serviceProtocolToJSON(object: ServiceProtocol): string;
/** ServiceType is the type of service that is being registered. */
export declare enum ServiceType {
    /** SERVICE_TYPE_UNSPECIFIED - SERVICE_TYPE_UNSPECIFIED is the default value. */
    SERVICE_TYPE_UNSPECIFIED = 0,
    /** SERVICE_TYPE_DID_COMM_MESSAGING - SERVICE_TYPE_APPLICATION is the type of service that is a DApp. */
    SERVICE_TYPE_DID_COMM_MESSAGING = 1,
    /** SERVICE_TYPE_LINKED_DOMAINS - SERVICE_TYPE_SERVICE is the type of service that is a service. */
    SERVICE_TYPE_LINKED_DOMAINS = 2,
    /** SERVICE_TYPE_SONR - SERVICE_TYPE_SONR is the type of service that is a DApp. */
    SERVICE_TYPE_SONR = 3,
    UNRECOGNIZED = -1
}
export declare function serviceTypeFromJSON(object: any): ServiceType;
export declare function serviceTypeToJSON(object: ServiceType): string;
/**
 * Did represents a string that has been parsed and validated as a DID. The parts are stored
 * in the individual fields.
 */
export interface Did {
    /** Method is the method used to create the DID. For the Sonr network it is "sonr". */
    method: string;
    /** Network is the network the DID is on. For testnet it is "testnet". i.e "did:sonr:testnet:". */
    network: string;
    /** id is the trailing identifier after the network. i.e. "did:sonr:testnet:abc123" */
    id: string;
    /** Paths is a list of paths that the DID is valid for. This is used to identify the Service. */
    paths: string[];
    /** Query is the query string that was used to create the DID. This is followed by a '?'. */
    query: string;
    /** Fragment is the fragment string that was used to create the DID. This is followed by a '#'. */
    fragment: string;
}
/** DidDocument is the document that describes a DID. This document is stored on the blockchain. */
export interface DidDocument {
    /** Context is the context of the DID document. This is used to identify the Service. */
    context: string[];
    /** Id is the DID of the document. */
    id: string;
    /** Controller is the DID of the controller of the document. This will be the individual user devices and mailboxes. */
    controller: string[];
    /** VerificationMethod is the list of verification methods for the user. */
    verificationMethod: VerificationMethod[];
    /** Authentication is the list of authentication methods for the user. */
    authentication: string[];
    /** AssertionMethod is the list of assertion methods for the user. */
    assertionMethod: string[];
    /** CapabilityInvocation is the list of capability invocation methods for the user. */
    capabilityInvocation: string[];
    /** CapabilityDelegation is the list of capability delegation methods for the user. */
    capabilityDelegation: string[];
    /** KeyAgreement is the list of key agreement methods for the user. */
    keyAgreement: string[];
    /** Service is the list of services or DApps that the user has access to. */
    service: Service[];
    /** AlsoKnownAs is the list of ".snr" aliases for the user. */
    alsoKnownAs: string[];
    /** Metadata is the metadata of the service. */
    metadata: {
        [key: string]: string;
    };
}
export interface DidDocument_MetadataEntry {
    key: string;
    value: string;
}
/** Service is a Application that runs on the Sonr network. */
export interface Service {
    /** ID is the DID of the service. */
    id: string;
    /** Type is the type of the service. */
    type: ServiceType;
    /** ServiceEndpoint is the endpoint of the service. */
    serviceEndpoint: ServiceEndpoint | undefined;
    /** Metadata is the metadata of the service. */
    metadata: {
        [key: string]: string;
    };
}
export interface Service_MetadataEntry {
    key: string;
    value: string;
}
/** ServiceEndpoint is the endpoint of the service. */
export interface ServiceEndpoint {
    /** TransportType is the type of transport used to connect to the service. */
    transportType: string;
    /** Network is the network the service is on. */
    network: string;
    /**
     * SupportedProtocols is the list of protocols supported by the service.
     * (e.g. "channels", "buckets", "objects", "storage")
     */
    supportedProtocols: ServiceProtocol[];
}
/** VerificationMethod is a method that can be used to verify the DID. */
export interface VerificationMethod {
    /** ID is the DID of the verification method. */
    id: string;
    /** Type is the type of the verification method. */
    type: VerificationMethod_Type;
    /** Controller is the DID of the controller of the verification method. */
    controller: string;
    /** PublicKeyHex is the public key of the verification method in hexidecimal. */
    publicKeyHex: string;
    /** PublicKeyBase58 is the public key of the verification method in base58. */
    publicKeyBase58: string;
    /** BlockchainAccountId is the blockchain account id of the verification method. */
    blockchainAccountId: string;
}
export declare enum VerificationMethod_Type {
    /** TYPE_UNSPECIFIED - TYPE_UNSPECIFIED is the default value. */
    TYPE_UNSPECIFIED = 0,
    /** TYPE_ECDSA_SECP256K1 - TYPE_ECDSA_SECP256K1 represents the Ed25519VerificationKey2018 key type. */
    TYPE_ECDSA_SECP256K1 = 1,
    /** TYPE_X25519 - TYPE_X25519 represents the X25519KeyAgreementKey2019 key type. */
    TYPE_X25519 = 2,
    /** TYPE_ED25519 - TYPE_ED25519 represents the Ed25519VerificationKey2018 key type. */
    TYPE_ED25519 = 3,
    /** TYPE_BLS_12381_G1 - TYPE_BLS_12381_G1 represents the Bls12381G1Key2020 key type */
    TYPE_BLS_12381_G1 = 4,
    /** TYPE_BLS_12381_G2 - TYPE_BLS_12381_G2 represents the Bls12381G2Key2020 key type */
    TYPE_BLS_12381_G2 = 5,
    /** TYPE_RSA - TYPE_RSA represents the RsaVerificationKey2018 key type. */
    TYPE_RSA = 6,
    /** TYPE_VERIFIABLE_CONDITION - TYPE_VERIFIABLE_CONDITION represents the VerifiableCondition2021 key type. */
    TYPE_VERIFIABLE_CONDITION = 7,
    UNRECOGNIZED = -1
}
export declare function verificationMethod_TypeFromJSON(object: any): VerificationMethod_Type;
export declare function verificationMethod_TypeToJSON(object: VerificationMethod_Type): string;
export declare const Did: {
    encode(message: Did, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Did;
    fromJSON(object: any): Did;
    toJSON(message: Did): unknown;
    fromPartial(object: DeepPartial<Did>): Did;
};
export declare const DidDocument: {
    encode(message: DidDocument, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): DidDocument;
    fromJSON(object: any): DidDocument;
    toJSON(message: DidDocument): unknown;
    fromPartial(object: DeepPartial<DidDocument>): DidDocument;
};
export declare const DidDocument_MetadataEntry: {
    encode(message: DidDocument_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): DidDocument_MetadataEntry;
    fromJSON(object: any): DidDocument_MetadataEntry;
    toJSON(message: DidDocument_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<DidDocument_MetadataEntry>): DidDocument_MetadataEntry;
};
export declare const Service: {
    encode(message: Service, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Service;
    fromJSON(object: any): Service;
    toJSON(message: Service): unknown;
    fromPartial(object: DeepPartial<Service>): Service;
};
export declare const Service_MetadataEntry: {
    encode(message: Service_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Service_MetadataEntry;
    fromJSON(object: any): Service_MetadataEntry;
    toJSON(message: Service_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<Service_MetadataEntry>): Service_MetadataEntry;
};
export declare const ServiceEndpoint: {
    encode(message: ServiceEndpoint, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ServiceEndpoint;
    fromJSON(object: any): ServiceEndpoint;
    toJSON(message: ServiceEndpoint): unknown;
    fromPartial(object: DeepPartial<ServiceEndpoint>): ServiceEndpoint;
};
export declare const VerificationMethod: {
    encode(message: VerificationMethod, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): VerificationMethod;
    fromJSON(object: any): VerificationMethod;
    toJSON(message: VerificationMethod): unknown;
    fromPartial(object: DeepPartial<VerificationMethod>): VerificationMethod;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
