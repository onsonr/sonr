import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "sonrio.sonr.bucket";
export interface MsgCreateBucket {
    creator: string;
    label: string;
    description: string;
    kind: string;
}
export interface MsgCreateBucketResponse {
}
export interface MsgReadBucket {
    creator: string;
    did: string;
}
export interface MsgReadBucketResponse {
}
export interface MsgUpdateBucket {
    creator: string;
    did: string;
    label: string;
    description: string;
}
export interface MsgUpdateBucketResponse {
}
export interface MsgDeactivateBucket {
    creator: string;
    did: string;
    publicKey: string;
}
export interface MsgDeactivateBucketResponse {
}
export declare const MsgCreateBucket: {
    encode(message: MsgCreateBucket, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateBucket;
    fromJSON(object: any): MsgCreateBucket;
    toJSON(message: MsgCreateBucket): unknown;
    fromPartial(object: DeepPartial<MsgCreateBucket>): MsgCreateBucket;
};
export declare const MsgCreateBucketResponse: {
    encode(_: MsgCreateBucketResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateBucketResponse;
    fromJSON(_: any): MsgCreateBucketResponse;
    toJSON(_: MsgCreateBucketResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateBucketResponse>): MsgCreateBucketResponse;
};
export declare const MsgReadBucket: {
    encode(message: MsgReadBucket, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgReadBucket;
    fromJSON(object: any): MsgReadBucket;
    toJSON(message: MsgReadBucket): unknown;
    fromPartial(object: DeepPartial<MsgReadBucket>): MsgReadBucket;
};
export declare const MsgReadBucketResponse: {
    encode(_: MsgReadBucketResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgReadBucketResponse;
    fromJSON(_: any): MsgReadBucketResponse;
    toJSON(_: MsgReadBucketResponse): unknown;
    fromPartial(_: DeepPartial<MsgReadBucketResponse>): MsgReadBucketResponse;
};
export declare const MsgUpdateBucket: {
    encode(message: MsgUpdateBucket, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateBucket;
    fromJSON(object: any): MsgUpdateBucket;
    toJSON(message: MsgUpdateBucket): unknown;
    fromPartial(object: DeepPartial<MsgUpdateBucket>): MsgUpdateBucket;
};
export declare const MsgUpdateBucketResponse: {
    encode(_: MsgUpdateBucketResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateBucketResponse;
    fromJSON(_: any): MsgUpdateBucketResponse;
    toJSON(_: MsgUpdateBucketResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateBucketResponse>): MsgUpdateBucketResponse;
};
export declare const MsgDeactivateBucket: {
    encode(message: MsgDeactivateBucket, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeactivateBucket;
    fromJSON(object: any): MsgDeactivateBucket;
    toJSON(message: MsgDeactivateBucket): unknown;
    fromPartial(object: DeepPartial<MsgDeactivateBucket>): MsgDeactivateBucket;
};
export declare const MsgDeactivateBucketResponse: {
    encode(_: MsgDeactivateBucketResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeactivateBucketResponse;
    fromJSON(_: any): MsgDeactivateBucketResponse;
    toJSON(_: MsgDeactivateBucketResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeactivateBucketResponse>): MsgDeactivateBucketResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateBucket(request: MsgCreateBucket): Promise<MsgCreateBucketResponse>;
    ReadBucket(request: MsgReadBucket): Promise<MsgReadBucketResponse>;
    UpdateBucket(request: MsgUpdateBucket): Promise<MsgUpdateBucketResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeleteBucket(request: MsgDeactivateBucket): Promise<MsgDeactivateBucketResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateBucket(request: MsgCreateBucket): Promise<MsgCreateBucketResponse>;
    ReadBucket(request: MsgReadBucket): Promise<MsgReadBucketResponse>;
    UpdateBucket(request: MsgUpdateBucket): Promise<MsgUpdateBucketResponse>;
    DeleteBucket(request: MsgDeactivateBucket): Promise<MsgDeactivateBucketResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
