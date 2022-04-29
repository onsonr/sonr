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
  last_modified: number;
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
  iso_country_code: string;
  /** The name of the country associated with the placemark. */
  country: string;
  /** The postal code associated with the placemark. */
  postal_code: string;
  /** The name of the state or province associated with the placemark. */
  administrative_area: string;
  /** Additional administrative area information for the placemark. */
  sub_administrative_area: string;
  /** The name of the city associated with the placemark. */
  locality: string;
  /** Additional city-level information for the placemark. */
  sub_locality: string;
  /** The street address associated with the placemark. */
  thoroughfare: string;
  /** Additional street address information for the placemark. */
  sub_thoroughfare: string;
}

/** Shared Metadata for Messages on all Protocols */
export interface Metadata {
  /** Unix timestamp */
  timestamp: number;
  /** Node ID */
  node_id: string;
  /** Signature of the message */
  signature: Uint8Array;
  /** Public Key of the message sender */
  public_key: Uint8Array;
}

/** Basic Info Sent to Peers to Establish Connections */
export interface Peer {
  /** User Sonr Domain */
  s_name: string;
  /** Peer Status */
  status: Peer_Status;
  /** Peer Device Info */
  device: Peer_Device | undefined;
  /** Peers General Information */
  profile: Profile | undefined;
  /** Public Key of the Peer */
  public_key: Uint8Array;
  /** Peer ID */
  peer_id: string;
  /** Last Modified Timestamp */
  last_modified: number;
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
  host_name: string;
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
  s_name: string;
  /** General Info */
  first_name: string;
  /** General Info */
  last_name: string;
  /** Profile Picture */
  picture: Uint8Array;
  /** User Biography */
  bio: string;
  /** Last Modified Timestamp */
  last_modified: number;
}

const baseLocation: object = { latitude: 0, longitude: 0, last_modified: 0 };

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
    if (message.last_modified !== 0) {
      writer.uint32(32).int64(message.last_modified);
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
          message.last_modified = longToNumber(reader.int64() as Long);
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
    if (object.last_modified !== undefined && object.last_modified !== null) {
      message.last_modified = Number(object.last_modified);
    } else {
      message.last_modified = 0;
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
    message.last_modified !== undefined &&
      (obj.last_modified = message.last_modified);
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
    if (object.last_modified !== undefined && object.last_modified !== null) {
      message.last_modified = object.last_modified;
    } else {
      message.last_modified = 0;
    }
    return message;
  },
};

const baseLocation_Placemark: object = {
  name: "",
  street: "",
  iso_country_code: "",
  country: "",
  postal_code: "",
  administrative_area: "",
  sub_administrative_area: "",
  locality: "",
  sub_locality: "",
  thoroughfare: "",
  sub_thoroughfare: "",
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
    if (message.iso_country_code !== "") {
      writer.uint32(26).string(message.iso_country_code);
    }
    if (message.country !== "") {
      writer.uint32(34).string(message.country);
    }
    if (message.postal_code !== "") {
      writer.uint32(42).string(message.postal_code);
    }
    if (message.administrative_area !== "") {
      writer.uint32(50).string(message.administrative_area);
    }
    if (message.sub_administrative_area !== "") {
      writer.uint32(58).string(message.sub_administrative_area);
    }
    if (message.locality !== "") {
      writer.uint32(66).string(message.locality);
    }
    if (message.sub_locality !== "") {
      writer.uint32(74).string(message.sub_locality);
    }
    if (message.thoroughfare !== "") {
      writer.uint32(82).string(message.thoroughfare);
    }
    if (message.sub_thoroughfare !== "") {
      writer.uint32(90).string(message.sub_thoroughfare);
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
          message.iso_country_code = reader.string();
          break;
        case 4:
          message.country = reader.string();
          break;
        case 5:
          message.postal_code = reader.string();
          break;
        case 6:
          message.administrative_area = reader.string();
          break;
        case 7:
          message.sub_administrative_area = reader.string();
          break;
        case 8:
          message.locality = reader.string();
          break;
        case 9:
          message.sub_locality = reader.string();
          break;
        case 10:
          message.thoroughfare = reader.string();
          break;
        case 11:
          message.sub_thoroughfare = reader.string();
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
    if (
      object.iso_country_code !== undefined &&
      object.iso_country_code !== null
    ) {
      message.iso_country_code = String(object.iso_country_code);
    } else {
      message.iso_country_code = "";
    }
    if (object.country !== undefined && object.country !== null) {
      message.country = String(object.country);
    } else {
      message.country = "";
    }
    if (object.postal_code !== undefined && object.postal_code !== null) {
      message.postal_code = String(object.postal_code);
    } else {
      message.postal_code = "";
    }
    if (
      object.administrative_area !== undefined &&
      object.administrative_area !== null
    ) {
      message.administrative_area = String(object.administrative_area);
    } else {
      message.administrative_area = "";
    }
    if (
      object.sub_administrative_area !== undefined &&
      object.sub_administrative_area !== null
    ) {
      message.sub_administrative_area = String(object.sub_administrative_area);
    } else {
      message.sub_administrative_area = "";
    }
    if (object.locality !== undefined && object.locality !== null) {
      message.locality = String(object.locality);
    } else {
      message.locality = "";
    }
    if (object.sub_locality !== undefined && object.sub_locality !== null) {
      message.sub_locality = String(object.sub_locality);
    } else {
      message.sub_locality = "";
    }
    if (object.thoroughfare !== undefined && object.thoroughfare !== null) {
      message.thoroughfare = String(object.thoroughfare);
    } else {
      message.thoroughfare = "";
    }
    if (
      object.sub_thoroughfare !== undefined &&
      object.sub_thoroughfare !== null
    ) {
      message.sub_thoroughfare = String(object.sub_thoroughfare);
    } else {
      message.sub_thoroughfare = "";
    }
    return message;
  },

  toJSON(message: Location_Placemark): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.street !== undefined && (obj.street = message.street);
    message.iso_country_code !== undefined &&
      (obj.iso_country_code = message.iso_country_code);
    message.country !== undefined && (obj.country = message.country);
    message.postal_code !== undefined &&
      (obj.postal_code = message.postal_code);
    message.administrative_area !== undefined &&
      (obj.administrative_area = message.administrative_area);
    message.sub_administrative_area !== undefined &&
      (obj.sub_administrative_area = message.sub_administrative_area);
    message.locality !== undefined && (obj.locality = message.locality);
    message.sub_locality !== undefined &&
      (obj.sub_locality = message.sub_locality);
    message.thoroughfare !== undefined &&
      (obj.thoroughfare = message.thoroughfare);
    message.sub_thoroughfare !== undefined &&
      (obj.sub_thoroughfare = message.sub_thoroughfare);
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
    if (
      object.iso_country_code !== undefined &&
      object.iso_country_code !== null
    ) {
      message.iso_country_code = object.iso_country_code;
    } else {
      message.iso_country_code = "";
    }
    if (object.country !== undefined && object.country !== null) {
      message.country = object.country;
    } else {
      message.country = "";
    }
    if (object.postal_code !== undefined && object.postal_code !== null) {
      message.postal_code = object.postal_code;
    } else {
      message.postal_code = "";
    }
    if (
      object.administrative_area !== undefined &&
      object.administrative_area !== null
    ) {
      message.administrative_area = object.administrative_area;
    } else {
      message.administrative_area = "";
    }
    if (
      object.sub_administrative_area !== undefined &&
      object.sub_administrative_area !== null
    ) {
      message.sub_administrative_area = object.sub_administrative_area;
    } else {
      message.sub_administrative_area = "";
    }
    if (object.locality !== undefined && object.locality !== null) {
      message.locality = object.locality;
    } else {
      message.locality = "";
    }
    if (object.sub_locality !== undefined && object.sub_locality !== null) {
      message.sub_locality = object.sub_locality;
    } else {
      message.sub_locality = "";
    }
    if (object.thoroughfare !== undefined && object.thoroughfare !== null) {
      message.thoroughfare = object.thoroughfare;
    } else {
      message.thoroughfare = "";
    }
    if (
      object.sub_thoroughfare !== undefined &&
      object.sub_thoroughfare !== null
    ) {
      message.sub_thoroughfare = object.sub_thoroughfare;
    } else {
      message.sub_thoroughfare = "";
    }
    return message;
  },
};

const baseMetadata: object = { timestamp: 0, node_id: "" };

export const Metadata = {
  encode(message: Metadata, writer: Writer = Writer.create()): Writer {
    if (message.timestamp !== 0) {
      writer.uint32(8).int64(message.timestamp);
    }
    if (message.node_id !== "") {
      writer.uint32(18).string(message.node_id);
    }
    if (message.signature.length !== 0) {
      writer.uint32(26).bytes(message.signature);
    }
    if (message.public_key.length !== 0) {
      writer.uint32(34).bytes(message.public_key);
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
          message.node_id = reader.string();
          break;
        case 3:
          message.signature = reader.bytes();
          break;
        case 4:
          message.public_key = reader.bytes();
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
    if (object.node_id !== undefined && object.node_id !== null) {
      message.node_id = String(object.node_id);
    } else {
      message.node_id = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = bytesFromBase64(object.signature);
    }
    if (object.public_key !== undefined && object.public_key !== null) {
      message.public_key = bytesFromBase64(object.public_key);
    }
    return message;
  },

  toJSON(message: Metadata): unknown {
    const obj: any = {};
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    message.node_id !== undefined && (obj.node_id = message.node_id);
    message.signature !== undefined &&
      (obj.signature = base64FromBytes(
        message.signature !== undefined ? message.signature : new Uint8Array()
      ));
    message.public_key !== undefined &&
      (obj.public_key = base64FromBytes(
        message.public_key !== undefined ? message.public_key : new Uint8Array()
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
    if (object.node_id !== undefined && object.node_id !== null) {
      message.node_id = object.node_id;
    } else {
      message.node_id = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = new Uint8Array();
    }
    if (object.public_key !== undefined && object.public_key !== null) {
      message.public_key = object.public_key;
    } else {
      message.public_key = new Uint8Array();
    }
    return message;
  },
};

const basePeer: object = {
  s_name: "",
  status: 0,
  peer_id: "",
  last_modified: 0,
};

export const Peer = {
  encode(message: Peer, writer: Writer = Writer.create()): Writer {
    if (message.s_name !== "") {
      writer.uint32(10).string(message.s_name);
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
    if (message.public_key.length !== 0) {
      writer.uint32(42).bytes(message.public_key);
    }
    if (message.peer_id !== "") {
      writer.uint32(50).string(message.peer_id);
    }
    if (message.last_modified !== 0) {
      writer.uint32(56).int64(message.last_modified);
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
          message.s_name = reader.string();
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
          message.public_key = reader.bytes();
          break;
        case 6:
          message.peer_id = reader.string();
          break;
        case 7:
          message.last_modified = longToNumber(reader.int64() as Long);
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
    if (object.s_name !== undefined && object.s_name !== null) {
      message.s_name = String(object.s_name);
    } else {
      message.s_name = "";
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
    if (object.public_key !== undefined && object.public_key !== null) {
      message.public_key = bytesFromBase64(object.public_key);
    }
    if (object.peer_id !== undefined && object.peer_id !== null) {
      message.peer_id = String(object.peer_id);
    } else {
      message.peer_id = "";
    }
    if (object.last_modified !== undefined && object.last_modified !== null) {
      message.last_modified = Number(object.last_modified);
    } else {
      message.last_modified = 0;
    }
    return message;
  },

  toJSON(message: Peer): unknown {
    const obj: any = {};
    message.s_name !== undefined && (obj.s_name = message.s_name);
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
    message.public_key !== undefined &&
      (obj.public_key = base64FromBytes(
        message.public_key !== undefined ? message.public_key : new Uint8Array()
      ));
    message.peer_id !== undefined && (obj.peer_id = message.peer_id);
    message.last_modified !== undefined &&
      (obj.last_modified = message.last_modified);
    return obj;
  },

  fromPartial(object: DeepPartial<Peer>): Peer {
    const message = { ...basePeer } as Peer;
    if (object.s_name !== undefined && object.s_name !== null) {
      message.s_name = object.s_name;
    } else {
      message.s_name = "";
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
    if (object.public_key !== undefined && object.public_key !== null) {
      message.public_key = object.public_key;
    } else {
      message.public_key = new Uint8Array();
    }
    if (object.peer_id !== undefined && object.peer_id !== null) {
      message.peer_id = object.peer_id;
    } else {
      message.peer_id = "";
    }
    if (object.last_modified !== undefined && object.last_modified !== null) {
      message.last_modified = object.last_modified;
    } else {
      message.last_modified = 0;
    }
    return message;
  },
};

const basePeer_Device: object = {
  id: "",
  host_name: "",
  os: "",
  arch: "",
  model: "",
};

export const Peer_Device = {
  encode(message: Peer_Device, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.host_name !== "") {
      writer.uint32(18).string(message.host_name);
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
          message.host_name = reader.string();
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
    if (object.host_name !== undefined && object.host_name !== null) {
      message.host_name = String(object.host_name);
    } else {
      message.host_name = "";
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
    message.host_name !== undefined && (obj.host_name = message.host_name);
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
    if (object.host_name !== undefined && object.host_name !== null) {
      message.host_name = object.host_name;
    } else {
      message.host_name = "";
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
  s_name: "",
  first_name: "",
  last_name: "",
  bio: "",
  last_modified: 0,
};

export const Profile = {
  encode(message: Profile, writer: Writer = Writer.create()): Writer {
    if (message.s_name !== "") {
      writer.uint32(10).string(message.s_name);
    }
    if (message.first_name !== "") {
      writer.uint32(18).string(message.first_name);
    }
    if (message.last_name !== "") {
      writer.uint32(26).string(message.last_name);
    }
    if (message.picture.length !== 0) {
      writer.uint32(34).bytes(message.picture);
    }
    if (message.bio !== "") {
      writer.uint32(50).string(message.bio);
    }
    if (message.last_modified !== 0) {
      writer.uint32(56).int64(message.last_modified);
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
          message.s_name = reader.string();
          break;
        case 2:
          message.first_name = reader.string();
          break;
        case 3:
          message.last_name = reader.string();
          break;
        case 4:
          message.picture = reader.bytes();
          break;
        case 6:
          message.bio = reader.string();
          break;
        case 7:
          message.last_modified = longToNumber(reader.int64() as Long);
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
    if (object.s_name !== undefined && object.s_name !== null) {
      message.s_name = String(object.s_name);
    } else {
      message.s_name = "";
    }
    if (object.first_name !== undefined && object.first_name !== null) {
      message.first_name = String(object.first_name);
    } else {
      message.first_name = "";
    }
    if (object.last_name !== undefined && object.last_name !== null) {
      message.last_name = String(object.last_name);
    } else {
      message.last_name = "";
    }
    if (object.picture !== undefined && object.picture !== null) {
      message.picture = bytesFromBase64(object.picture);
    }
    if (object.bio !== undefined && object.bio !== null) {
      message.bio = String(object.bio);
    } else {
      message.bio = "";
    }
    if (object.last_modified !== undefined && object.last_modified !== null) {
      message.last_modified = Number(object.last_modified);
    } else {
      message.last_modified = 0;
    }
    return message;
  },

  toJSON(message: Profile): unknown {
    const obj: any = {};
    message.s_name !== undefined && (obj.s_name = message.s_name);
    message.first_name !== undefined && (obj.first_name = message.first_name);
    message.last_name !== undefined && (obj.last_name = message.last_name);
    message.picture !== undefined &&
      (obj.picture = base64FromBytes(
        message.picture !== undefined ? message.picture : new Uint8Array()
      ));
    message.bio !== undefined && (obj.bio = message.bio);
    message.last_modified !== undefined &&
      (obj.last_modified = message.last_modified);
    return obj;
  },

  fromPartial(object: DeepPartial<Profile>): Profile {
    const message = { ...baseProfile } as Profile;
    if (object.s_name !== undefined && object.s_name !== null) {
      message.s_name = object.s_name;
    } else {
      message.s_name = "";
    }
    if (object.first_name !== undefined && object.first_name !== null) {
      message.first_name = object.first_name;
    } else {
      message.first_name = "";
    }
    if (object.last_name !== undefined && object.last_name !== null) {
      message.last_name = object.last_name;
    } else {
      message.last_name = "";
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
    if (object.last_modified !== undefined && object.last_modified !== null) {
      message.last_modified = object.last_modified;
    } else {
      message.last_modified = 0;
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
