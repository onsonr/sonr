import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "sonrio.sonr.registry";
/** Internet Connection Type */
export declare enum Connection {
    CONNECTION_UNSPECIFIED = 0,
    /** CONNECTION_WIFI - ConnectionWifi is used for WiFi connections. */
    CONNECTION_WIFI = 1,
    /** CONNECTION_ETHERNET - ConnectionEthernet is used for Ethernet connections. */
    CONNECTION_ETHERNET = 2,
    /** CONNECTION_MOBILE - ConnectionMobile is used for mobile connections. */
    CONNECTION_MOBILE = 3,
    /** CONNECTION_OFFLINE - CONNECTION_OFFLINE */
    CONNECTION_OFFLINE = 4,
    UNRECOGNIZED = -1
}
export declare function connectionFromJSON(object: any): Connection;
export declare function connectionToJSON(object: Connection): string;
/** Location from GeoIP and OLC information */
export interface Location {
    /** Location Latitude */
    latitude: number;
    /** Location Longitude */
    longitude: number;
    /** Location Placemark Information - Generated */
    placemark: Location_Placemark | undefined;
    /** Last Updated Time */
    lastModified: number;
}
/** Contains detailed placemark information. */
export interface Location_Placemark {
    /** The name associated with the placemark. */
    name: string;
    /** The street associated with the placemark. */
    street: string;
    /**
     * The abbreviated country name, according to the two letter (alpha-2) [ISO
     * standard](https://www.iso.org/iso-3166-country-codes.html).
     */
    isoCountryCode: string;
    /** The name of the country associated with the placemark. */
    country: string;
    /** The postal code associated with the placemark. */
    postalCode: string;
    /** The name of the state or province associated with the placemark. */
    administrativeArea: string;
    /** Additional administrative area information for the placemark. */
    subAdministrativeArea: string;
    /** The name of the city associated with the placemark. */
    locality: string;
    /** Additional city-level information for the placemark. */
    subLocality: string;
    /** The street address associated with the placemark. */
    thoroughfare: string;
    /** Additional street address information for the placemark. */
    subThoroughfare: string;
}
/** Shared Metadata for Messages on all Protocols */
export interface Metadata {
    /** Unix timestamp */
    timestamp: number;
    /** Node ID */
    nodeId: string;
    /** Signature of the message */
    signature: Uint8Array;
    /** Public Key of the message sender */
    publicKey: Uint8Array;
}
/** Basic Info Sent to Peers to Establish Connections */
export interface Peer {
    /** User Sonr Domain */
    sName: string;
    /** Peer Status */
    status: Peer_Status;
    /** Peer Device Info */
    device: Peer_Device | undefined;
    /** Peers General Information */
    profile: Profile | undefined;
    /** Public Key of the Peer */
    publicKey: Uint8Array;
    /** Peer ID */
    peerId: string;
    /** Last Modified Timestamp */
    lastModified: number;
}
/** Peers Active Status */
export declare enum Peer_Status {
    /** STATUS_UNSPECIFIED - Offline - Not Online or Not a Full Node */
    STATUS_UNSPECIFIED = 0,
    /** STATUS_ONLINE - Online - Full Node Available */
    STATUS_ONLINE = 1,
    /** STATUS_AWAY - Away - Not Online, but has a full node */
    STATUS_AWAY = 2,
    /** STATUS_BUSY - Busy - Online, but busy with Transfer */
    STATUS_BUSY = 3,
    UNRECOGNIZED = -1
}
export declare function peer_StatusFromJSON(object: any): Peer_Status;
export declare function peer_StatusToJSON(object: Peer_Status): string;
/** Peer Info for Device */
export interface Peer_Device {
    /** Peer Device ID */
    id: string;
    /** Peer Host Name */
    hostName: string;
    /** Peer Operating System */
    os: string;
    /** Peer Architecture */
    arch: string;
    /** Peers Device Model */
    model: string;
}
/** General Information about Peer passed during Authentication */
export interface Profile {
    /** Sonr Based Username */
    sName: string;
    /** General Info */
    firstName: string;
    /** General Info */
    lastName: string;
    /** Profile Picture */
    picture: Uint8Array;
    /** User Biography */
    bio: string;
    /** Last Modified Timestamp */
    lastModified: number;
}
export declare const Location: {
    encode(message: Location, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Location;
    fromJSON(object: any): Location;
    toJSON(message: Location): unknown;
    fromPartial(object: DeepPartial<Location>): Location;
};
export declare const Location_Placemark: {
    encode(message: Location_Placemark, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Location_Placemark;
    fromJSON(object: any): Location_Placemark;
    toJSON(message: Location_Placemark): unknown;
    fromPartial(object: DeepPartial<Location_Placemark>): Location_Placemark;
};
export declare const Metadata: {
    encode(message: Metadata, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Metadata;
    fromJSON(object: any): Metadata;
    toJSON(message: Metadata): unknown;
    fromPartial(object: DeepPartial<Metadata>): Metadata;
};
export declare const Peer: {
    encode(message: Peer, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Peer;
    fromJSON(object: any): Peer;
    toJSON(message: Peer): unknown;
    fromPartial(object: DeepPartial<Peer>): Peer;
};
export declare const Peer_Device: {
    encode(message: Peer_Device, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Peer_Device;
    fromJSON(object: any): Peer_Device;
    toJSON(message: Peer_Device): unknown;
    fromPartial(object: DeepPartial<Peer_Device>): Peer_Device;
};
export declare const Profile: {
    encode(message: Profile, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Profile;
    fromJSON(object: any): Profile;
    toJSON(message: Profile): unknown;
    fromPartial(object: DeepPartial<Profile>): Profile;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
