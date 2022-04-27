import { Reader, Writer } from "protobufjs/minimal";
import { Did, DidDocument } from "../registry/did";
export declare const protobufPackage = "sonrio.sonr.registry";
export interface MsgRegisterService {
    creator: string;
    serviceName: string;
    publicKey: string;
}
export interface MsgRegisterServiceResponse {
}
/** MsgRegisterName is a request to register a name with the ".snr" name of a peer */
export interface MsgRegisterName {
    creator: string;
    deviceId: string;
    os: string;
    model: string;
    arch: string;
    publicKey: string;
    nameToRegister: string;
}
export interface MsgRegisterNameResponse {
    isSuccess: boolean;
    did: Did | undefined;
    didDocument: DidDocument | undefined;
}
/** MsgAccessName defines the MsgAccessName transaction. */
export interface MsgAccessName {
    /** The account that is accessing the name */
    creator: string;
    /** The name to access */
    name: string;
    /** The Public Key of the peer to access */
    publicKey: string;
    /** The Libp2p peer ID of the peer to access */
    peerId: string;
}
export interface MsgAccessNameResponse {
    name: string;
    publicKey: string;
    peerId: string;
}
export interface MsgUpdateName {
    /** The account that owns the name. */
    creator: string;
    /** The name of the peer to update the name of */
    name: string;
    /** The Updated Metadata */
    metadata: {
        [key: string]: string;
    };
}
export interface MsgUpdateName_MetadataEntry {
    key: string;
    value: string;
}
export interface MsgUpdateNameResponse {
    didDocument: DidDocument | undefined;
    /** optional */
    metadata: {
        [key: string]: string;
    };
}
export interface MsgUpdateNameResponse_MetadataEntry {
    key: string;
    value: string;
}
export interface MsgAccessService {
    /** The account that is accessing the service */
    creator: string;
    /** The name of the service to access */
    did: string;
}
export interface MsgAccessServiceResponse {
    /** Code of the response */
    code: number;
    /** Message of the response */
    message: string;
    /** Data of the response */
    metadata: {
        [key: string]: string;
    };
}
export interface MsgAccessServiceResponse_MetadataEntry {
    key: string;
    value: string;
}
export interface MsgUpdateService {
    /** The account that owns the name. */
    creator: string;
    /** The name of the peer to update the service details of */
    did: string;
    /** The updated configuration for the service */
    configuration: {
        [key: string]: string;
    };
    /** The metadata for any service information required */
    metadata: {
        [key: string]: string;
    };
}
export interface MsgUpdateService_ConfigurationEntry {
    key: string;
    value: string;
}
export interface MsgUpdateService_MetadataEntry {
    key: string;
    value: string;
}
export interface MsgUpdateServiceResponse {
    didDocument: DidDocument | undefined;
    /** The updated configuration for the service */
    configuration: {
        [key: string]: string;
    };
    /** The metadata for any service information required */
    metadata: {
        [key: string]: string;
    };
}
export interface MsgUpdateServiceResponse_ConfigurationEntry {
    key: string;
    value: string;
}
export interface MsgUpdateServiceResponse_MetadataEntry {
    key: string;
    value: string;
}
export declare const MsgRegisterService: {
    encode(message: MsgRegisterService, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRegisterService;
    fromJSON(object: any): MsgRegisterService;
    toJSON(message: MsgRegisterService): unknown;
    fromPartial(object: DeepPartial<MsgRegisterService>): MsgRegisterService;
};
export declare const MsgRegisterServiceResponse: {
    encode(_: MsgRegisterServiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRegisterServiceResponse;
    fromJSON(_: any): MsgRegisterServiceResponse;
    toJSON(_: MsgRegisterServiceResponse): unknown;
    fromPartial(_: DeepPartial<MsgRegisterServiceResponse>): MsgRegisterServiceResponse;
};
export declare const MsgRegisterName: {
    encode(message: MsgRegisterName, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRegisterName;
    fromJSON(object: any): MsgRegisterName;
    toJSON(message: MsgRegisterName): unknown;
    fromPartial(object: DeepPartial<MsgRegisterName>): MsgRegisterName;
};
export declare const MsgRegisterNameResponse: {
    encode(message: MsgRegisterNameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRegisterNameResponse;
    fromJSON(object: any): MsgRegisterNameResponse;
    toJSON(message: MsgRegisterNameResponse): unknown;
    fromPartial(object: DeepPartial<MsgRegisterNameResponse>): MsgRegisterNameResponse;
};
export declare const MsgAccessName: {
    encode(message: MsgAccessName, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAccessName;
    fromJSON(object: any): MsgAccessName;
    toJSON(message: MsgAccessName): unknown;
    fromPartial(object: DeepPartial<MsgAccessName>): MsgAccessName;
};
export declare const MsgAccessNameResponse: {
    encode(message: MsgAccessNameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAccessNameResponse;
    fromJSON(object: any): MsgAccessNameResponse;
    toJSON(message: MsgAccessNameResponse): unknown;
    fromPartial(object: DeepPartial<MsgAccessNameResponse>): MsgAccessNameResponse;
};
export declare const MsgUpdateName: {
    encode(message: MsgUpdateName, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateName;
    fromJSON(object: any): MsgUpdateName;
    toJSON(message: MsgUpdateName): unknown;
    fromPartial(object: DeepPartial<MsgUpdateName>): MsgUpdateName;
};
export declare const MsgUpdateName_MetadataEntry: {
    encode(message: MsgUpdateName_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateName_MetadataEntry;
    fromJSON(object: any): MsgUpdateName_MetadataEntry;
    toJSON(message: MsgUpdateName_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<MsgUpdateName_MetadataEntry>): MsgUpdateName_MetadataEntry;
};
export declare const MsgUpdateNameResponse: {
    encode(message: MsgUpdateNameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateNameResponse;
    fromJSON(object: any): MsgUpdateNameResponse;
    toJSON(message: MsgUpdateNameResponse): unknown;
    fromPartial(object: DeepPartial<MsgUpdateNameResponse>): MsgUpdateNameResponse;
};
export declare const MsgUpdateNameResponse_MetadataEntry: {
    encode(message: MsgUpdateNameResponse_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateNameResponse_MetadataEntry;
    fromJSON(object: any): MsgUpdateNameResponse_MetadataEntry;
    toJSON(message: MsgUpdateNameResponse_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<MsgUpdateNameResponse_MetadataEntry>): MsgUpdateNameResponse_MetadataEntry;
};
export declare const MsgAccessService: {
    encode(message: MsgAccessService, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAccessService;
    fromJSON(object: any): MsgAccessService;
    toJSON(message: MsgAccessService): unknown;
    fromPartial(object: DeepPartial<MsgAccessService>): MsgAccessService;
};
export declare const MsgAccessServiceResponse: {
    encode(message: MsgAccessServiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAccessServiceResponse;
    fromJSON(object: any): MsgAccessServiceResponse;
    toJSON(message: MsgAccessServiceResponse): unknown;
    fromPartial(object: DeepPartial<MsgAccessServiceResponse>): MsgAccessServiceResponse;
};
export declare const MsgAccessServiceResponse_MetadataEntry: {
    encode(message: MsgAccessServiceResponse_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAccessServiceResponse_MetadataEntry;
    fromJSON(object: any): MsgAccessServiceResponse_MetadataEntry;
    toJSON(message: MsgAccessServiceResponse_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<MsgAccessServiceResponse_MetadataEntry>): MsgAccessServiceResponse_MetadataEntry;
};
export declare const MsgUpdateService: {
    encode(message: MsgUpdateService, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateService;
    fromJSON(object: any): MsgUpdateService;
    toJSON(message: MsgUpdateService): unknown;
    fromPartial(object: DeepPartial<MsgUpdateService>): MsgUpdateService;
};
export declare const MsgUpdateService_ConfigurationEntry: {
    encode(message: MsgUpdateService_ConfigurationEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateService_ConfigurationEntry;
    fromJSON(object: any): MsgUpdateService_ConfigurationEntry;
    toJSON(message: MsgUpdateService_ConfigurationEntry): unknown;
    fromPartial(object: DeepPartial<MsgUpdateService_ConfigurationEntry>): MsgUpdateService_ConfigurationEntry;
};
export declare const MsgUpdateService_MetadataEntry: {
    encode(message: MsgUpdateService_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateService_MetadataEntry;
    fromJSON(object: any): MsgUpdateService_MetadataEntry;
    toJSON(message: MsgUpdateService_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<MsgUpdateService_MetadataEntry>): MsgUpdateService_MetadataEntry;
};
export declare const MsgUpdateServiceResponse: {
    encode(message: MsgUpdateServiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateServiceResponse;
    fromJSON(object: any): MsgUpdateServiceResponse;
    toJSON(message: MsgUpdateServiceResponse): unknown;
    fromPartial(object: DeepPartial<MsgUpdateServiceResponse>): MsgUpdateServiceResponse;
};
export declare const MsgUpdateServiceResponse_ConfigurationEntry: {
    encode(message: MsgUpdateServiceResponse_ConfigurationEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateServiceResponse_ConfigurationEntry;
    fromJSON(object: any): MsgUpdateServiceResponse_ConfigurationEntry;
    toJSON(message: MsgUpdateServiceResponse_ConfigurationEntry): unknown;
    fromPartial(object: DeepPartial<MsgUpdateServiceResponse_ConfigurationEntry>): MsgUpdateServiceResponse_ConfigurationEntry;
};
export declare const MsgUpdateServiceResponse_MetadataEntry: {
    encode(message: MsgUpdateServiceResponse_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateServiceResponse_MetadataEntry;
    fromJSON(object: any): MsgUpdateServiceResponse_MetadataEntry;
    toJSON(message: MsgUpdateServiceResponse_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<MsgUpdateServiceResponse_MetadataEntry>): MsgUpdateServiceResponse_MetadataEntry;
};
/** Msg defines the Msg service. */
export interface Msg {
    RegisterService(request: MsgRegisterService): Promise<MsgRegisterServiceResponse>;
    RegisterName(request: MsgRegisterName): Promise<MsgRegisterNameResponse>;
    AccessName(request: MsgAccessName): Promise<MsgAccessNameResponse>;
    UpdateName(request: MsgUpdateName): Promise<MsgUpdateNameResponse>;
    AccessService(request: MsgAccessService): Promise<MsgAccessServiceResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    UpdateService(request: MsgUpdateService): Promise<MsgUpdateServiceResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    RegisterService(request: MsgRegisterService): Promise<MsgRegisterServiceResponse>;
    RegisterName(request: MsgRegisterName): Promise<MsgRegisterNameResponse>;
    AccessName(request: MsgAccessName): Promise<MsgAccessNameResponse>;
    UpdateName(request: MsgUpdateName): Promise<MsgUpdateNameResponse>;
    AccessService(request: MsgAccessService): Promise<MsgAccessServiceResponse>;
    UpdateService(request: MsgUpdateService): Promise<MsgUpdateServiceResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
