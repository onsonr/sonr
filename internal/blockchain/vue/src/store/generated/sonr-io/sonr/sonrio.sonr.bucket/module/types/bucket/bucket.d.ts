import { Did } from "../registry/did";
import { ObjectDoc } from "../object/object";
import { Peer } from "../registry/peer";
import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "sonrio.sonr.bucket";
/** BucketType is the type of a bucket. */
export declare enum BucketType {
    /** BUCKET_TYPE_UNSPECIFIED - BucketTypeUnspecified is the default value. */
    BUCKET_TYPE_UNSPECIFIED = 0,
    /** BUCKET_TYPE_APP - BucketTypeApp is an App specific bucket. For Assets regarding the service. */
    BUCKET_TYPE_APP = 1,
    /**
     * BUCKET_TYPE_USER - BucketTypeUser is a User specific bucket. For any remote user data that is required
     * to be stored in the Network.
     */
    BUCKET_TYPE_USER = 2,
    UNRECOGNIZED = -1
}
export declare function bucketTypeFromJSON(object: any): BucketType;
export declare function bucketTypeToJSON(object: BucketType): string;
/** EventType is the type of event being performed on a Bucket. */
export declare enum BucketEventType {
    /** Bucket_EVENT_TYPE_UNSPECIFIED - EventTypeUnspecified is the default value. */
    Bucket_EVENT_TYPE_UNSPECIFIED = 0,
    /** Bucket_EVENT_TYPE_GET - EventTypeGet is a get event being performed on a Bucket record. */
    Bucket_EVENT_TYPE_GET = 1,
    /** Bucket_EVENT_TYPE_SET - EventTypeSet is a set event on the Bucket store. */
    Bucket_EVENT_TYPE_SET = 2,
    /** Bucket_EVENT_TYPE_DELETE - EventTypeDelete is a delete event on the Bucket store. */
    Bucket_EVENT_TYPE_DELETE = 3,
    UNRECOGNIZED = -1
}
export declare function bucketEventTypeFromJSON(object: any): BucketEventType;
export declare function bucketEventTypeToJSON(object: BucketEventType): string;
/** Bucket is a collection of objects. */
export interface Bucket {
    /** Label is human-readable name of the bucket. */
    label: string;
    /** Description is a human-readable description of the bucket. */
    description: string;
    /** Type is the kind of bucket for either App specific or User specific data. */
    type: BucketType;
    /** Did is the identifier of the bucket. */
    did: Did | undefined;
    /** Objects are stored in a tree structure. */
    objects: ObjectDoc[];
}
/** BucketEvent is the base event type for all Bucket events. */
export interface BucketEvent {
    /** Owner is the peer that originated the event. */
    peer: Peer | undefined;
    /** Type is the type of event being performed on a Bucket. */
    type: BucketEventType;
    /** Object is the entry being modified in the Bucket. */
    object: ObjectDoc | undefined;
    /** Metadata is the metadata associated with the event. */
    metadata: {
        [key: string]: string;
    };
}
export interface BucketEvent_MetadataEntry {
    key: string;
    value: string;
}
export declare const Bucket: {
    encode(message: Bucket, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Bucket;
    fromJSON(object: any): Bucket;
    toJSON(message: Bucket): unknown;
    fromPartial(object: DeepPartial<Bucket>): Bucket;
};
export declare const BucketEvent: {
    encode(message: BucketEvent, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): BucketEvent;
    fromJSON(object: any): BucketEvent;
    toJSON(message: BucketEvent): unknown;
    fromPartial(object: DeepPartial<BucketEvent>): BucketEvent;
};
export declare const BucketEvent_MetadataEntry: {
    encode(message: BucketEvent_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): BucketEvent_MetadataEntry;
    fromJSON(object: any): BucketEvent_MetadataEntry;
    toJSON(message: BucketEvent_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<BucketEvent_MetadataEntry>): BucketEvent_MetadataEntry;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
