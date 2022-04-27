import { Did } from "../registry/did";
import { ObjectDoc } from "../object/object";
import { Peer } from "../registry/peer";
import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "sonrio.sonr.channel";
export interface Channel {
    /** Label is human-readable name of the channel. */
    label: string;
    /** Description is a human-readable description of the channel. */
    description: string;
    /** Did is the identifier of the channel. */
    did: Did | undefined;
    /** RegisterdObject is the object that is registered as the payload for the channel. */
    registeredObject: ObjectDoc | undefined;
}
/** ChannelMessage is a message sent to a channel. */
export interface ChannelMessage {
    /** Owner is the peer that originated the message. */
    peer: Peer | undefined;
    /** Did is the identifier of the channel. */
    did: Did | undefined;
    /** Data is the data being sent. */
    data: ObjectDoc | undefined;
    /** Metadata is the metadata associated with the message. */
    metadata: {
        [key: string]: string;
    };
}
export interface ChannelMessage_MetadataEntry {
    key: string;
    value: string;
}
export declare const Channel: {
    encode(message: Channel, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Channel;
    fromJSON(object: any): Channel;
    toJSON(message: Channel): unknown;
    fromPartial(object: DeepPartial<Channel>): Channel;
};
export declare const ChannelMessage: {
    encode(message: ChannelMessage, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ChannelMessage;
    fromJSON(object: any): ChannelMessage;
    toJSON(message: ChannelMessage): unknown;
    fromPartial(object: DeepPartial<ChannelMessage>): ChannelMessage;
};
export declare const ChannelMessage_MetadataEntry: {
    encode(message: ChannelMessage_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ChannelMessage_MetadataEntry;
    fromJSON(object: any): ChannelMessage_MetadataEntry;
    toJSON(message: ChannelMessage_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<ChannelMessage_MetadataEntry>): ChannelMessage_MetadataEntry;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
