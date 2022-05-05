import { Did } from "../registry/did";
import { ObjectDoc } from "../object/object";
import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "sonrio.sonr.registry";
/** ServiceConfig is the configuration for a service. */
export interface ServiceConfig {
    /** Name is the name of the service. */
    name: string;
    /** Description is a human readable description of the service. */
    description: string;
    /** Id is the DID of the service. */
    did: Did | undefined;
    /** Maintainers is the DID of the service maintainers. */
    maintainers: Did[];
    /** Channels is a list of channels the service is registered on. */
    channels: Did[];
    /** Buckets is a list of buckets the service is registered on. */
    buckets: Did[];
    /** Objects is a map of objects the service is registered on. */
    objects: {
        [key: string]: ObjectDoc;
    };
    /** Endpoints is a list of endpoints the service is registered on. */
    endpoints: string[];
    /** Metadata is the metadata associated with the event. */
    metadata: {
        [key: string]: string;
    };
    /** Version is the version of the service. Version must be a semantic version. */
    version: string;
}
export interface ServiceConfig_ObjectsEntry {
    key: string;
    value: ObjectDoc | undefined;
}
export interface ServiceConfig_MetadataEntry {
    key: string;
    value: string;
}
export declare const ServiceConfig: {
    encode(message: ServiceConfig, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ServiceConfig;
    fromJSON(object: any): ServiceConfig;
    toJSON(message: ServiceConfig): unknown;
    fromPartial(object: DeepPartial<ServiceConfig>): ServiceConfig;
};
export declare const ServiceConfig_ObjectsEntry: {
    encode(message: ServiceConfig_ObjectsEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ServiceConfig_ObjectsEntry;
    fromJSON(object: any): ServiceConfig_ObjectsEntry;
    toJSON(message: ServiceConfig_ObjectsEntry): unknown;
    fromPartial(object: DeepPartial<ServiceConfig_ObjectsEntry>): ServiceConfig_ObjectsEntry;
};
export declare const ServiceConfig_MetadataEntry: {
    encode(message: ServiceConfig_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ServiceConfig_MetadataEntry;
    fromJSON(object: any): ServiceConfig_MetadataEntry;
    toJSON(message: ServiceConfig_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<ServiceConfig_MetadataEntry>): ServiceConfig_MetadataEntry;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
