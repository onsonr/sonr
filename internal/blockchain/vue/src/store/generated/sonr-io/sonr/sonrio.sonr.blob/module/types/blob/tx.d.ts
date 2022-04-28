import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "sonrio.sonr.blob";
export interface MsgUploadBlob {
    creator: string;
    label: string;
    path: string;
    refDid: string;
    size: number;
    lastModified: number;
}
export interface MsgUploadBlobResponse {
}
export interface MsgDownloadBlob {
    creator: string;
    did: string;
    path: string;
    timeout: number;
}
export interface MsgDownloadBlobResponse {
}
export interface MsgSyncBlob {
    creator: string;
    did: string;
    path: string;
    timeout: number;
}
export interface MsgSyncBlobResponse {
}
export interface MsgDeleteBlob {
    creator: string;
    did: string;
    publicKey: string;
}
export interface MsgDeleteBlobResponse {
}
export declare const MsgUploadBlob: {
    encode(message: MsgUploadBlob, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUploadBlob;
    fromJSON(object: any): MsgUploadBlob;
    toJSON(message: MsgUploadBlob): unknown;
    fromPartial(object: DeepPartial<MsgUploadBlob>): MsgUploadBlob;
};
export declare const MsgUploadBlobResponse: {
    encode(_: MsgUploadBlobResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUploadBlobResponse;
    fromJSON(_: any): MsgUploadBlobResponse;
    toJSON(_: MsgUploadBlobResponse): unknown;
    fromPartial(_: DeepPartial<MsgUploadBlobResponse>): MsgUploadBlobResponse;
};
export declare const MsgDownloadBlob: {
    encode(message: MsgDownloadBlob, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDownloadBlob;
    fromJSON(object: any): MsgDownloadBlob;
    toJSON(message: MsgDownloadBlob): unknown;
    fromPartial(object: DeepPartial<MsgDownloadBlob>): MsgDownloadBlob;
};
export declare const MsgDownloadBlobResponse: {
    encode(_: MsgDownloadBlobResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDownloadBlobResponse;
    fromJSON(_: any): MsgDownloadBlobResponse;
    toJSON(_: MsgDownloadBlobResponse): unknown;
    fromPartial(_: DeepPartial<MsgDownloadBlobResponse>): MsgDownloadBlobResponse;
};
export declare const MsgSyncBlob: {
    encode(message: MsgSyncBlob, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSyncBlob;
    fromJSON(object: any): MsgSyncBlob;
    toJSON(message: MsgSyncBlob): unknown;
    fromPartial(object: DeepPartial<MsgSyncBlob>): MsgSyncBlob;
};
export declare const MsgSyncBlobResponse: {
    encode(_: MsgSyncBlobResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSyncBlobResponse;
    fromJSON(_: any): MsgSyncBlobResponse;
    toJSON(_: MsgSyncBlobResponse): unknown;
    fromPartial(_: DeepPartial<MsgSyncBlobResponse>): MsgSyncBlobResponse;
};
export declare const MsgDeleteBlob: {
    encode(message: MsgDeleteBlob, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteBlob;
    fromJSON(object: any): MsgDeleteBlob;
    toJSON(message: MsgDeleteBlob): unknown;
    fromPartial(object: DeepPartial<MsgDeleteBlob>): MsgDeleteBlob;
};
export declare const MsgDeleteBlobResponse: {
    encode(_: MsgDeleteBlobResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteBlobResponse;
    fromJSON(_: any): MsgDeleteBlobResponse;
    toJSON(_: MsgDeleteBlobResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteBlobResponse>): MsgDeleteBlobResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    UploadBlob(request: MsgUploadBlob): Promise<MsgUploadBlobResponse>;
    DownloadBlob(request: MsgDownloadBlob): Promise<MsgDownloadBlobResponse>;
    SyncBlob(request: MsgSyncBlob): Promise<MsgSyncBlobResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeleteBlob(request: MsgDeleteBlob): Promise<MsgDeleteBlobResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    UploadBlob(request: MsgUploadBlob): Promise<MsgUploadBlobResponse>;
    DownloadBlob(request: MsgDownloadBlob): Promise<MsgDownloadBlobResponse>;
    SyncBlob(request: MsgSyncBlob): Promise<MsgSyncBlobResponse>;
    DeleteBlob(request: MsgDeleteBlob): Promise<MsgDeleteBlobResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
