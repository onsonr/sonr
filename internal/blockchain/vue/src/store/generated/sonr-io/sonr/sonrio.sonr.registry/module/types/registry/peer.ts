/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sonrio.sonr.registry";

/** Internet Connection Type */
export enum Connection {
  CONNECTION_UNSPECIFIED = 0,
  /** CONNECTION_WIFI - ConnectionWifi is used for WiFi connections. */
  CONNECTION_WIFI = 1,
  /** CONNECTION_ETHERNET - ConnectionEthernet is used for Ethernet connections. */
  CONNECTION_ETHERNET = 2,
  /** CONNECTION_MOBILE - ConnectionMobile is used for mobile connections. */
  CONNECTION_MOBILE = 3,
  /** CONNECTION_OFFLINE - CONNECTION_OFFLINE */
  CONNECTION_OFFLINE = 4,
  UNRECOGNIZED = -1,
}

export function connectionFromJSON(object: any): Connection {
  switch (object) {
    case 0:
    case "CONNECTION_UNSPECIFIED":
      return Connection.CONNECTION_UNSPECIFIED;
    case 1:
    case "CONNECTION_WIFI":
      return Connection.CONNECTION_WIFI;
    case 2:
    case "CONNECTION_ETHERNET":
      return Connection.CONNECTION_ETHERNET;
    case 3:
    case "CONNECTION_MOBILE":
      return Connection.CONNECTION_MOBILE;
    case 4:
    case "CONNECTION_OFFLINE":
      return Connection.CONNECTION_OFFLINE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Connection.UNRECOGNIZED;
  }
}

export function connectionToJSON(object: Connection): string {
  switch (object) {
    case Connection.CONNECTION_UNSPECIFIED:
      return "CONNECTION_UNSPECIFIED";
    case Connection.CONNECTION_WIFI:
      return "CONNECTION_WIFI";
    case Connection.CONNECTION_ETHERNET:
      return "CONNECTION_ETHERNET";
    case Connection.CONNECTION_MOBILE:
      return "CONNECTION_MOBILE";
    case Connection.CONNECTION_OFFLINE:
      return "CONNECTION_OFFLINE";
    default:
      return "UNKNOWN";
  }
}

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
export enum Peer_Status {
  /** STATUS_UNSPECIFIED - Offline - Not Online or Not a Full Node */
  STATUS_UNSPECIFIED = 0,
  /** STATUS_ONLINE - Online - Full Node Available */
  STATUS_ONLINE = 1,
  /** STATUS_AWAY - Away - Not Online, but has a full node */
  STATUS_AWAY = 2,
  /** STATUS_BUSY - Busy - Online, but busy with Transfer */
  STATUS_BUSY = 3,
  UNRECOGNIZED = -1,
}

export function peer_StatusFromJSON(object: any): Peer_Status {
  switch (object) {
    case 0:
    case "STATUS_UNSPECIFIED":
      return Peer_Status.STATUS_UNSPECIFIED;
    case 1:
    case "STATUS_ONLINE":
      return Peer_Status.STATUS_ONLINE;
    case 2:
    case "STATUS_AWAY":
      return Peer_Status.STATUS_AWAY;
    case 3:
    case "STATUS_BUSY":
      return Peer_Status.STATUS_BUSY;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Peer_Status.UNRECOGNIZED;
  }
}

export function peer_StatusToJSON(object: Peer_Status): string {
  switch (object) {
    case Peer_Status.STATUS_UNSPECIFIED:
      return "STATUS_UNSPECIFIED";
    case Peer_Status.STATUS_ONLINE:
      return "STATUS_ONLINE";
    case Peer_Status.STATUS_AWAY:
      return "STATUS_AWAY";
    case Peer_Status.STATUS_BUSY:
      return "STATUS_BUSY";
    default:
      return "UNKNOWN";
  }
}

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

const baseLocation: object = { latitude: 0, longitude: 0, lastModified: 0 };

export const Location = {
  encode(message: Location, writer: Writer = Writer.create()): Writer {
    if (message.latitude !== 0) {
      writer.uint32(9).double(message.latitude);
    }
    if (message.longitude !== 0) {
      writer.uint32(17).double(message.longitude);
    }
    if (message.placemark !== undefined) {
      Location_Placemark.encode(
        message.placemark,
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.lastModified !== 0) {
      writer.uint32(32).int64(message.lastModified);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Location {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLocation } as Location;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.latitude = reader.double();
          break;
        case 2:
          message.longitude = reader.double();
          break;
        case 3:
          message.placemark = Location_Placemark.decode(
            reader,
            reader.uint32()
          );
          break;
        case 4:
          message.lastModified = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Location {
    const message = { ...baseLocation } as Location;
    if (object.latitude !== undefined && object.latitude !== null) {
      message.latitude = Number(object.latitude);
    } else {
      message.latitude = 0;
    }
    if (object.longitude !== undefined && object.longitude !== null) {
      message.longitude = Number(object.longitude);
    } else {
      message.longitude = 0;
    }
    if (object.placemark !== undefined && object.placemark !== null) {
      message.placemark = Location_Placemark.fromJSON(object.placemark);
    } else {
      message.placemark = undefined;
    }
    if (object.lastModified !== undefined && object.lastModified !== null) {
      message.lastModified = Number(object.lastModified);
    } else {
      message.lastModified = 0;
    }
    return message;
  },

  toJSON(message: Location): unknown {
    const obj: any = {};
    message.latitude !== undefined && (obj.latitude = message.latitude);
    message.longitude !== undefined && (obj.longitude = message.longitude);
    message.placemark !== undefined &&
      (obj.placemark = message.placemark
        ? Location_Placemark.toJSON(message.placemark)
        : undefined);
    message.lastModified !== undefined &&
      (obj.lastModified = message.lastModified);
    return obj;
  },

  fromPartial(object: DeepPartial<Location>): Location {
    const message = { ...baseLocation } as Location;
    if (object.latitude !== undefined && object.latitude !== null) {
      message.latitude = object.latitude;
    } else {
      message.latitude = 0;
    }
    if (object.longitude !== undefined && object.longitude !== null) {
      message.longitude = object.longitude;
    } else {
      message.longitude = 0;
    }
    if (object.placemark !== undefined && object.placemark !== null) {
      message.placemark = Location_Placemark.fromPartial(object.placemark);
    } else {
      message.placemark = undefined;
    }
    if (object.lastModified !== undefined && object.lastModified !== null) {
      message.lastModified = object.lastModified;
    } else {
      message.lastModified = 0;
    }
    return message;
  },
};

const baseLocation_Placemark: object = {
  name: "",
  street: "",
  isoCountryCode: "",
  country: "",
  postalCode: "",
  administrativeArea: "",
  subAdministrativeArea: "",
  locality: "",
  subLocality: "",
  thoroughfare: "",
  subThoroughfare: "",
};

export const Location_Placemark = {
  encode(
    message: Location_Placemark,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.street !== "") {
      writer.uint32(18).string(message.street);
    }
    if (message.isoCountryCode !== "") {
      writer.uint32(26).string(message.isoCountryCode);
    }
    if (message.country !== "") {
      writer.uint32(34).string(message.country);
    }
    if (message.postalCode !== "") {
      writer.uint32(42).string(message.postalCode);
    }
    if (message.administrativeArea !== "") {
      writer.uint32(50).string(message.administrativeArea);
    }
    if (message.subAdministrativeArea !== "") {
      writer.uint32(58).string(message.subAdministrativeArea);
    }
    if (message.locality !== "") {
      writer.uint32(66).string(message.locality);
    }
    if (message.subLocality !== "") {
      writer.uint32(74).string(message.subLocality);
    }
    if (message.thoroughfare !== "") {
      writer.uint32(82).string(message.thoroughfare);
    }
    if (message.subThoroughfare !== "") {
      writer.uint32(90).string(message.subThoroughfare);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Location_Placemark {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLocation_Placemark } as Location_Placemark;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.street = reader.string();
          break;
        case 3:
          message.isoCountryCode = reader.string();
          break;
        case 4:
          message.country = reader.string();
          break;
        case 5:
          message.postalCode = reader.string();
          break;
        case 6:
          message.administrativeArea = reader.string();
          break;
        case 7:
          message.subAdministrativeArea = reader.string();
          break;
        case 8:
          message.locality = reader.string();
          break;
        case 9:
          message.subLocality = reader.string();
          break;
        case 10:
          message.thoroughfare = reader.string();
          break;
        case 11:
          message.subThoroughfare = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Location_Placemark {
    const message = { ...baseLocation_Placemark } as Location_Placemark;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.street !== undefined && object.street !== null) {
      message.street = String(object.street);
    } else {
      message.street = "";
    }
    if (object.isoCountryCode !== undefined && object.isoCountryCode !== null) {
      message.isoCountryCode = String(object.isoCountryCode);
    } else {
      message.isoCountryCode = "";
    }
    if (object.country !== undefined && object.country !== null) {
      message.country = String(object.country);
    } else {
      message.country = "";
    }
    if (object.postalCode !== undefined && object.postalCode !== null) {
      message.postalCode = String(object.postalCode);
    } else {
      message.postalCode = "";
    }
    if (
      object.administrativeArea !== undefined &&
      object.administrativeArea !== null
    ) {
      message.administrativeArea = String(object.administrativeArea);
    } else {
      message.administrativeArea = "";
    }
    if (
      object.subAdministrativeArea !== undefined &&
      object.subAdministrativeArea !== null
    ) {
      message.subAdministrativeArea = String(object.subAdministrativeArea);
    } else {
      message.subAdministrativeArea = "";
    }
    if (object.locality !== undefined && object.locality !== null) {
      message.locality = String(object.locality);
    } else {
      message.locality = "";
    }
    if (object.subLocality !== undefined && object.subLocality !== null) {
      message.subLocality = String(object.subLocality);
    } else {
      message.subLocality = "";
    }
    if (object.thoroughfare !== undefined && object.thoroughfare !== null) {
      message.thoroughfare = String(object.thoroughfare);
    } else {
      message.thoroughfare = "";
    }
    if (
      object.subThoroughfare !== undefined &&
      object.subThoroughfare !== null
    ) {
      message.subThoroughfare = String(object.subThoroughfare);
    } else {
      message.subThoroughfare = "";
    }
    return message;
  },

  toJSON(message: Location_Placemark): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.street !== undefined && (obj.street = message.street);
    message.isoCountryCode !== undefined &&
      (obj.isoCountryCode = message.isoCountryCode);
    message.country !== undefined && (obj.country = message.country);
    message.postalCode !== undefined && (obj.postalCode = message.postalCode);
    message.administrativeArea !== undefined &&
      (obj.administrativeArea = message.administrativeArea);
    message.subAdministrativeArea !== undefined &&
      (obj.subAdministrativeArea = message.subAdministrativeArea);
    message.locality !== undefined && (obj.locality = message.locality);
    message.subLocality !== undefined &&
      (obj.subLocality = message.subLocality);
    message.thoroughfare !== undefined &&
      (obj.thoroughfare = message.thoroughfare);
    message.subThoroughfare !== undefined &&
      (obj.subThoroughfare = message.subThoroughfare);
    return obj;
  },

  fromPartial(object: DeepPartial<Location_Placemark>): Location_Placemark {
    const message = { ...baseLocation_Placemark } as Location_Placemark;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.street !== undefined && object.street !== null) {
      message.street = object.street;
    } else {
      message.street = "";
    }
    if (object.isoCountryCode !== undefined && object.isoCountryCode !== null) {
      message.isoCountryCode = object.isoCountryCode;
    } else {
      message.isoCountryCode = "";
    }
    if (object.country !== undefined && object.country !== null) {
      message.country = object.country;
    } else {
      message.country = "";
    }
    if (object.postalCode !== undefined && object.postalCode !== null) {
      message.postalCode = object.postalCode;
    } else {
      message.postalCode = "";
    }
    if (
      object.administrativeArea !== undefined &&
      object.administrativeArea !== null
    ) {
      message.administrativeArea = object.administrativeArea;
    } else {
      message.administrativeArea = "";
    }
    if (
      object.subAdministrativeArea !== undefined &&
      object.subAdministrativeArea !== null
    ) {
      message.subAdministrativeArea = object.subAdministrativeArea;
    } else {
      message.subAdministrativeArea = "";
    }
    if (object.locality !== undefined && object.locality !== null) {
      message.locality = object.locality;
    } else {
      message.locality = "";
    }
    if (object.subLocality !== undefined && object.subLocality !== null) {
      message.subLocality = object.subLocality;
    } else {
      message.subLocality = "";
    }
    if (object.thoroughfare !== undefined && object.thoroughfare !== null) {
      message.thoroughfare = object.thoroughfare;
    } else {
      message.thoroughfare = "";
    }
    if (
      object.subThoroughfare !== undefined &&
      object.subThoroughfare !== null
    ) {
      message.subThoroughfare = object.subThoroughfare;
    } else {
      message.subThoroughfare = "";
    }
    return message;
  },
};

const baseMetadata: object = { timestamp: 0, nodeId: "" };

export const Metadata = {
  encode(message: Metadata, writer: Writer = Writer.create()): Writer {
    if (message.timestamp !== 0) {
      writer.uint32(8).int64(message.timestamp);
    }
    if (message.nodeId !== "") {
      writer.uint32(18).string(message.nodeId);
    }
    if (message.signature.length !== 0) {
      writer.uint32(26).bytes(message.signature);
    }
    if (message.publicKey.length !== 0) {
      writer.uint32(34).bytes(message.publicKey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Metadata {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMetadata } as Metadata;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.timestamp = longToNumber(reader.int64() as Long);
          break;
        case 2:
          message.nodeId = reader.string();
          break;
        case 3:
          message.signature = reader.bytes();
          break;
        case 4:
          message.publicKey = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Metadata {
    const message = { ...baseMetadata } as Metadata;
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = Number(object.timestamp);
    } else {
      message.timestamp = 0;
    }
    if (object.nodeId !== undefined && object.nodeId !== null) {
      message.nodeId = String(object.nodeId);
    } else {
      message.nodeId = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = bytesFromBase64(object.signature);
    }
    if (object.publicKey !== undefined && object.publicKey !== null) {
      message.publicKey = bytesFromBase64(object.publicKey);
    }
    return message;
  },

  toJSON(message: Metadata): unknown {
    const obj: any = {};
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    message.nodeId !== undefined && (obj.nodeId = message.nodeId);
    message.signature !== undefined &&
      (obj.signature = base64FromBytes(
        message.signature !== undefined ? message.signature : new Uint8Array()
      ));
    message.publicKey !== undefined &&
      (obj.publicKey = base64FromBytes(
        message.publicKey !== undefined ? message.publicKey : new Uint8Array()
      ));
    return obj;
  },

  fromPartial(object: DeepPartial<Metadata>): Metadata {
    const message = { ...baseMetadata } as Metadata;
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = object.timestamp;
    } else {
      message.timestamp = 0;
    }
    if (object.nodeId !== undefined && object.nodeId !== null) {
      message.nodeId = object.nodeId;
    } else {
      message.nodeId = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = new Uint8Array();
    }
    if (object.publicKey !== undefined && object.publicKey !== null) {
      message.publicKey = object.publicKey;
    } else {
      message.publicKey = new Uint8Array();
    }
    return message;
  },
};

const basePeer: object = { sName: "", status: 0, peerId: "", lastModified: 0 };

export const Peer = {
  encode(message: Peer, writer: Writer = Writer.create()): Writer {
    if (message.sName !== "") {
      writer.uint32(10).string(message.sName);
    }
    if (message.status !== 0) {
      writer.uint32(16).int32(message.status);
    }
    if (message.device !== undefined) {
      Peer_Device.encode(message.device, writer.uint32(26).fork()).ldelim();
    }
    if (message.profile !== undefined) {
      Profile.encode(message.profile, writer.uint32(34).fork()).ldelim();
    }
    if (message.publicKey.length !== 0) {
      writer.uint32(42).bytes(message.publicKey);
    }
    if (message.peerId !== "") {
      writer.uint32(50).string(message.peerId);
    }
    if (message.lastModified !== 0) {
      writer.uint32(56).int64(message.lastModified);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Peer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePeer } as Peer;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sName = reader.string();
          break;
        case 2:
          message.status = reader.int32() as any;
          break;
        case 3:
          message.device = Peer_Device.decode(reader, reader.uint32());
          break;
        case 4:
          message.profile = Profile.decode(reader, reader.uint32());
          break;
        case 5:
          message.publicKey = reader.bytes();
          break;
        case 6:
          message.peerId = reader.string();
          break;
        case 7:
          message.lastModified = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Peer {
    const message = { ...basePeer } as Peer;
    if (object.sName !== undefined && object.sName !== null) {
      message.sName = String(object.sName);
    } else {
      message.sName = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = peer_StatusFromJSON(object.status);
    } else {
      message.status = 0;
    }
    if (object.device !== undefined && object.device !== null) {
      message.device = Peer_Device.fromJSON(object.device);
    } else {
      message.device = undefined;
    }
    if (object.profile !== undefined && object.profile !== null) {
      message.profile = Profile.fromJSON(object.profile);
    } else {
      message.profile = undefined;
    }
    if (object.publicKey !== undefined && object.publicKey !== null) {
      message.publicKey = bytesFromBase64(object.publicKey);
    }
    if (object.peerId !== undefined && object.peerId !== null) {
      message.peerId = String(object.peerId);
    } else {
      message.peerId = "";
    }
    if (object.lastModified !== undefined && object.lastModified !== null) {
      message.lastModified = Number(object.lastModified);
    } else {
      message.lastModified = 0;
    }
    return message;
  },

  toJSON(message: Peer): unknown {
    const obj: any = {};
    message.sName !== undefined && (obj.sName = message.sName);
    message.status !== undefined &&
      (obj.status = peer_StatusToJSON(message.status));
    message.device !== undefined &&
      (obj.device = message.device
        ? Peer_Device.toJSON(message.device)
        : undefined);
    message.profile !== undefined &&
      (obj.profile = message.profile
        ? Profile.toJSON(message.profile)
        : undefined);
    message.publicKey !== undefined &&
      (obj.publicKey = base64FromBytes(
        message.publicKey !== undefined ? message.publicKey : new Uint8Array()
      ));
    message.peerId !== undefined && (obj.peerId = message.peerId);
    message.lastModified !== undefined &&
      (obj.lastModified = message.lastModified);
    return obj;
  },

  fromPartial(object: DeepPartial<Peer>): Peer {
    const message = { ...basePeer } as Peer;
    if (object.sName !== undefined && object.sName !== null) {
      message.sName = object.sName;
    } else {
      message.sName = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    if (object.device !== undefined && object.device !== null) {
      message.device = Peer_Device.fromPartial(object.device);
    } else {
      message.device = undefined;
    }
    if (object.profile !== undefined && object.profile !== null) {
      message.profile = Profile.fromPartial(object.profile);
    } else {
      message.profile = undefined;
    }
    if (object.publicKey !== undefined && object.publicKey !== null) {
      message.publicKey = object.publicKey;
    } else {
      message.publicKey = new Uint8Array();
    }
    if (object.peerId !== undefined && object.peerId !== null) {
      message.peerId = object.peerId;
    } else {
      message.peerId = "";
    }
    if (object.lastModified !== undefined && object.lastModified !== null) {
      message.lastModified = object.lastModified;
    } else {
      message.lastModified = 0;
    }
    return message;
  },
};

const basePeer_Device: object = {
  id: "",
  hostName: "",
  os: "",
  arch: "",
  model: "",
};

export const Peer_Device = {
  encode(message: Peer_Device, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.hostName !== "") {
      writer.uint32(18).string(message.hostName);
    }
    if (message.os !== "") {
      writer.uint32(26).string(message.os);
    }
    if (message.arch !== "") {
      writer.uint32(34).string(message.arch);
    }
    if (message.model !== "") {
      writer.uint32(42).string(message.model);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Peer_Device {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePeer_Device } as Peer_Device;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.hostName = reader.string();
          break;
        case 3:
          message.os = reader.string();
          break;
        case 4:
          message.arch = reader.string();
          break;
        case 5:
          message.model = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Peer_Device {
    const message = { ...basePeer_Device } as Peer_Device;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.hostName !== undefined && object.hostName !== null) {
      message.hostName = String(object.hostName);
    } else {
      message.hostName = "";
    }
    if (object.os !== undefined && object.os !== null) {
      message.os = String(object.os);
    } else {
      message.os = "";
    }
    if (object.arch !== undefined && object.arch !== null) {
      message.arch = String(object.arch);
    } else {
      message.arch = "";
    }
    if (object.model !== undefined && object.model !== null) {
      message.model = String(object.model);
    } else {
      message.model = "";
    }
    return message;
  },

  toJSON(message: Peer_Device): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.hostName !== undefined && (obj.hostName = message.hostName);
    message.os !== undefined && (obj.os = message.os);
    message.arch !== undefined && (obj.arch = message.arch);
    message.model !== undefined && (obj.model = message.model);
    return obj;
  },

  fromPartial(object: DeepPartial<Peer_Device>): Peer_Device {
    const message = { ...basePeer_Device } as Peer_Device;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.hostName !== undefined && object.hostName !== null) {
      message.hostName = object.hostName;
    } else {
      message.hostName = "";
    }
    if (object.os !== undefined && object.os !== null) {
      message.os = object.os;
    } else {
      message.os = "";
    }
    if (object.arch !== undefined && object.arch !== null) {
      message.arch = object.arch;
    } else {
      message.arch = "";
    }
    if (object.model !== undefined && object.model !== null) {
      message.model = object.model;
    } else {
      message.model = "";
    }
    return message;
  },
};

const baseProfile: object = {
  sName: "",
  firstName: "",
  lastName: "",
  bio: "",
  lastModified: 0,
};

export const Profile = {
  encode(message: Profile, writer: Writer = Writer.create()): Writer {
    if (message.sName !== "") {
      writer.uint32(10).string(message.sName);
    }
    if (message.firstName !== "") {
      writer.uint32(18).string(message.firstName);
    }
    if (message.lastName !== "") {
      writer.uint32(26).string(message.lastName);
    }
    if (message.picture.length !== 0) {
      writer.uint32(34).bytes(message.picture);
    }
    if (message.bio !== "") {
      writer.uint32(50).string(message.bio);
    }
    if (message.lastModified !== 0) {
      writer.uint32(56).int64(message.lastModified);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Profile {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseProfile } as Profile;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sName = reader.string();
          break;
        case 2:
          message.firstName = reader.string();
          break;
        case 3:
          message.lastName = reader.string();
          break;
        case 4:
          message.picture = reader.bytes();
          break;
        case 6:
          message.bio = reader.string();
          break;
        case 7:
          message.lastModified = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Profile {
    const message = { ...baseProfile } as Profile;
    if (object.sName !== undefined && object.sName !== null) {
      message.sName = String(object.sName);
    } else {
      message.sName = "";
    }
    if (object.firstName !== undefined && object.firstName !== null) {
      message.firstName = String(object.firstName);
    } else {
      message.firstName = "";
    }
    if (object.lastName !== undefined && object.lastName !== null) {
      message.lastName = String(object.lastName);
    } else {
      message.lastName = "";
    }
    if (object.picture !== undefined && object.picture !== null) {
      message.picture = bytesFromBase64(object.picture);
    }
    if (object.bio !== undefined && object.bio !== null) {
      message.bio = String(object.bio);
    } else {
      message.bio = "";
    }
    if (object.lastModified !== undefined && object.lastModified !== null) {
      message.lastModified = Number(object.lastModified);
    } else {
      message.lastModified = 0;
    }
    return message;
  },

  toJSON(message: Profile): unknown {
    const obj: any = {};
    message.sName !== undefined && (obj.sName = message.sName);
    message.firstName !== undefined && (obj.firstName = message.firstName);
    message.lastName !== undefined && (obj.lastName = message.lastName);
    message.picture !== undefined &&
      (obj.picture = base64FromBytes(
        message.picture !== undefined ? message.picture : new Uint8Array()
      ));
    message.bio !== undefined && (obj.bio = message.bio);
    message.lastModified !== undefined &&
      (obj.lastModified = message.lastModified);
    return obj;
  },

  fromPartial(object: DeepPartial<Profile>): Profile {
    const message = { ...baseProfile } as Profile;
    if (object.sName !== undefined && object.sName !== null) {
      message.sName = object.sName;
    } else {
      message.sName = "";
    }
    if (object.firstName !== undefined && object.firstName !== null) {
      message.firstName = object.firstName;
    } else {
      message.firstName = "";
    }
    if (object.lastName !== undefined && object.lastName !== null) {
      message.lastName = object.lastName;
    } else {
      message.lastName = "";
    }
    if (object.picture !== undefined && object.picture !== null) {
      message.picture = object.picture;
    } else {
      message.picture = new Uint8Array();
    }
    if (object.bio !== undefined && object.bio !== null) {
      message.bio = object.bio;
    } else {
      message.bio = "";
    }
    if (object.lastModified !== undefined && object.lastModified !== null) {
      message.lastModified = object.lastModified;
    } else {
      message.lastModified = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
