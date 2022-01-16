/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "common";

/** Internet Connection Type */
export enum Connection {
  CONNECTION_UNSPECIFIED = "CONNECTION_UNSPECIFIED",
  /** CONNECTION_WIFI - ConnectionWifi is used for WiFi connections. */
  CONNECTION_WIFI = "CONNECTION_WIFI",
  /** CONNECTION_ETHERNET - ConnectionEthernet is used for Ethernet connections. */
  CONNECTION_ETHERNET = "CONNECTION_ETHERNET",
  /** CONNECTION_MOBILE - ConnectionMobile is used for mobile connections. */
  CONNECTION_MOBILE = "CONNECTION_MOBILE",
  /** CONNECTION_OFFLINE - CONNECTION_OFFLINE */
  CONNECTION_OFFLINE = "CONNECTION_OFFLINE",
  UNRECOGNIZED = "UNRECOGNIZED",
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

export function connectionToNumber(object: Connection): number {
  switch (object) {
    case Connection.CONNECTION_UNSPECIFIED:
      return 0;
    case Connection.CONNECTION_WIFI:
      return 1;
    case Connection.CONNECTION_ETHERNET:
      return 2;
    case Connection.CONNECTION_MOBILE:
      return 3;
    case Connection.CONNECTION_OFFLINE:
      return 4;
    default:
      return 0;
  }
}

/** Location from GeoIP and OLC information */
export interface Location {
  /** Location Latitude */
  latitude: number;
  /** Location Longitude */
  longitude: number;
  /** Location Placemark Information - Generated */
  placemark?: Location_Placemark;
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
  signature: Buffer;
  /** Public Key of the message sender */
  publicKey: Buffer;
}

/** Standard MIME with Additional Extensions */
export interface MIME {
  /** Type of File */
  type: MIME_Type;
  /** Extension of File */
  subtype: string;
  /** Type/Subtype i.e. (image/jpeg) */
  value: string;
}

/** File Content Type */
export enum MIME_Type {
  /** TYPE_UNSPECIFIED - Other File Type - If cannot derive from Subtype */
  TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED",
  /** TYPE_AUDIO - Sound, Audio Files */
  TYPE_AUDIO = "TYPE_AUDIO",
  /** TYPE_DOCUMENT - Document Files - PDF, Word, Excel, etc. */
  TYPE_DOCUMENT = "TYPE_DOCUMENT",
  /** TYPE_IMAGE - Image Files */
  TYPE_IMAGE = "TYPE_IMAGE",
  /** TYPE_TEXT - Text Based Files */
  TYPE_TEXT = "TYPE_TEXT",
  /** TYPE_VIDEO - Video Files */
  TYPE_VIDEO = "TYPE_VIDEO",
  /** TYPE_URL - URL Links */
  TYPE_URL = "TYPE_URL",
  /** TYPE_CRYPTO - Crypto Files */
  TYPE_CRYPTO = "TYPE_CRYPTO",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function mIME_TypeFromJSON(object: any): MIME_Type {
  switch (object) {
    case 0:
    case "TYPE_UNSPECIFIED":
      return MIME_Type.TYPE_UNSPECIFIED;
    case 1:
    case "TYPE_AUDIO":
      return MIME_Type.TYPE_AUDIO;
    case 2:
    case "TYPE_DOCUMENT":
      return MIME_Type.TYPE_DOCUMENT;
    case 3:
    case "TYPE_IMAGE":
      return MIME_Type.TYPE_IMAGE;
    case 4:
    case "TYPE_TEXT":
      return MIME_Type.TYPE_TEXT;
    case 5:
    case "TYPE_VIDEO":
      return MIME_Type.TYPE_VIDEO;
    case 6:
    case "TYPE_URL":
      return MIME_Type.TYPE_URL;
    case 7:
    case "TYPE_CRYPTO":
      return MIME_Type.TYPE_CRYPTO;
    case -1:
    case "UNRECOGNIZED":
    default:
      return MIME_Type.UNRECOGNIZED;
  }
}

export function mIME_TypeToJSON(object: MIME_Type): string {
  switch (object) {
    case MIME_Type.TYPE_UNSPECIFIED:
      return "TYPE_UNSPECIFIED";
    case MIME_Type.TYPE_AUDIO:
      return "TYPE_AUDIO";
    case MIME_Type.TYPE_DOCUMENT:
      return "TYPE_DOCUMENT";
    case MIME_Type.TYPE_IMAGE:
      return "TYPE_IMAGE";
    case MIME_Type.TYPE_TEXT:
      return "TYPE_TEXT";
    case MIME_Type.TYPE_VIDEO:
      return "TYPE_VIDEO";
    case MIME_Type.TYPE_URL:
      return "TYPE_URL";
    case MIME_Type.TYPE_CRYPTO:
      return "TYPE_CRYPTO";
    default:
      return "UNKNOWN";
  }
}

export function mIME_TypeToNumber(object: MIME_Type): number {
  switch (object) {
    case MIME_Type.TYPE_UNSPECIFIED:
      return 0;
    case MIME_Type.TYPE_AUDIO:
      return 1;
    case MIME_Type.TYPE_DOCUMENT:
      return 2;
    case MIME_Type.TYPE_IMAGE:
      return 3;
    case MIME_Type.TYPE_TEXT:
      return 4;
    case MIME_Type.TYPE_VIDEO:
      return 5;
    case MIME_Type.TYPE_URL:
      return 6;
    case MIME_Type.TYPE_CRYPTO:
      return 7;
    default:
      return 0;
  }
}

/** Basic Info Sent to Peers to Establish Connections */
export interface Peer {
  /** User Sonr Domain */
  sName: string;
  /** Peer Status */
  status: Peer_Status;
  /** Peer Device Info */
  device?: Peer_Device;
  /** Peers General Information */
  profile?: Profile;
  /** Public Key of the Peer */
  publicKey: Buffer;
  /** Peer ID */
  peerId: string;
  /** Last Modified Timestamp */
  lastModified: number;
}

/** Peers Active Status */
export enum Peer_Status {
  /** STATUS_UNSPECIFIED - Offline - Not Online or Not a Full Node */
  STATUS_UNSPECIFIED = "STATUS_UNSPECIFIED",
  /** STATUS_ONLINE - Online - Full Node Available */
  STATUS_ONLINE = "STATUS_ONLINE",
  /** STATUS_AWAY - Away - Not Online, but has a full node */
  STATUS_AWAY = "STATUS_AWAY",
  /** STATUS_BUSY - Busy - Online, but busy with Transfer */
  STATUS_BUSY = "STATUS_BUSY",
  UNRECOGNIZED = "UNRECOGNIZED",
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

export function peer_StatusToNumber(object: Peer_Status): number {
  switch (object) {
    case Peer_Status.STATUS_UNSPECIFIED:
      return 0;
    case Peer_Status.STATUS_ONLINE:
      return 1;
    case Peer_Status.STATUS_AWAY:
      return 2;
    case Peer_Status.STATUS_BUSY:
      return 3;
    default:
      return 0;
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
  picture: Buffer;
  /** User Biography */
  bio: string;
  /** Last Modified Timestamp */
  lastModified: number;
}

function createBaseLocation(): Location {
  return { latitude: 0, longitude: 0, placemark: undefined, lastModified: 0 };
}

export const Location = {
  encode(
    message: Location,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): Location {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLocation();
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
    return {
      latitude: isSet(object.latitude) ? Number(object.latitude) : 0,
      longitude: isSet(object.longitude) ? Number(object.longitude) : 0,
      placemark: isSet(object.placemark)
        ? Location_Placemark.fromJSON(object.placemark)
        : undefined,
      lastModified: isSet(object.lastModified)
        ? Number(object.lastModified)
        : 0,
    };
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
      (obj.lastModified = Math.round(message.lastModified));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Location>, I>>(object: I): Location {
    const message = createBaseLocation();
    message.latitude = object.latitude ?? 0;
    message.longitude = object.longitude ?? 0;
    message.placemark =
      object.placemark !== undefined && object.placemark !== null
        ? Location_Placemark.fromPartial(object.placemark)
        : undefined;
    message.lastModified = object.lastModified ?? 0;
    return message;
  },
};

function createBaseLocation_Placemark(): Location_Placemark {
  return {
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
}

export const Location_Placemark = {
  encode(
    message: Location_Placemark,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): Location_Placemark {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLocation_Placemark();
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
    return {
      name: isSet(object.name) ? String(object.name) : "",
      street: isSet(object.street) ? String(object.street) : "",
      isoCountryCode: isSet(object.isoCountryCode)
        ? String(object.isoCountryCode)
        : "",
      country: isSet(object.country) ? String(object.country) : "",
      postalCode: isSet(object.postalCode) ? String(object.postalCode) : "",
      administrativeArea: isSet(object.administrativeArea)
        ? String(object.administrativeArea)
        : "",
      subAdministrativeArea: isSet(object.subAdministrativeArea)
        ? String(object.subAdministrativeArea)
        : "",
      locality: isSet(object.locality) ? String(object.locality) : "",
      subLocality: isSet(object.subLocality) ? String(object.subLocality) : "",
      thoroughfare: isSet(object.thoroughfare)
        ? String(object.thoroughfare)
        : "",
      subThoroughfare: isSet(object.subThoroughfare)
        ? String(object.subThoroughfare)
        : "",
    };
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

  fromPartial<I extends Exact<DeepPartial<Location_Placemark>, I>>(
    object: I
  ): Location_Placemark {
    const message = createBaseLocation_Placemark();
    message.name = object.name ?? "";
    message.street = object.street ?? "";
    message.isoCountryCode = object.isoCountryCode ?? "";
    message.country = object.country ?? "";
    message.postalCode = object.postalCode ?? "";
    message.administrativeArea = object.administrativeArea ?? "";
    message.subAdministrativeArea = object.subAdministrativeArea ?? "";
    message.locality = object.locality ?? "";
    message.subLocality = object.subLocality ?? "";
    message.thoroughfare = object.thoroughfare ?? "";
    message.subThoroughfare = object.subThoroughfare ?? "";
    return message;
  },
};

function createBaseMetadata(): Metadata {
  return {
    timestamp: 0,
    nodeId: "",
    signature: Buffer.alloc(0),
    publicKey: Buffer.alloc(0),
  };
}

export const Metadata = {
  encode(
    message: Metadata,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): Metadata {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMetadata();
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
          message.signature = reader.bytes() as Buffer;
          break;
        case 4:
          message.publicKey = reader.bytes() as Buffer;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Metadata {
    return {
      timestamp: isSet(object.timestamp) ? Number(object.timestamp) : 0,
      nodeId: isSet(object.nodeId) ? String(object.nodeId) : "",
      signature: isSet(object.signature)
        ? Buffer.from(bytesFromBase64(object.signature))
        : Buffer.alloc(0),
      publicKey: isSet(object.publicKey)
        ? Buffer.from(bytesFromBase64(object.publicKey))
        : Buffer.alloc(0),
    };
  },

  toJSON(message: Metadata): unknown {
    const obj: any = {};
    message.timestamp !== undefined &&
      (obj.timestamp = Math.round(message.timestamp));
    message.nodeId !== undefined && (obj.nodeId = message.nodeId);
    message.signature !== undefined &&
      (obj.signature = base64FromBytes(
        message.signature !== undefined ? message.signature : Buffer.alloc(0)
      ));
    message.publicKey !== undefined &&
      (obj.publicKey = base64FromBytes(
        message.publicKey !== undefined ? message.publicKey : Buffer.alloc(0)
      ));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Metadata>, I>>(object: I): Metadata {
    const message = createBaseMetadata();
    message.timestamp = object.timestamp ?? 0;
    message.nodeId = object.nodeId ?? "";
    message.signature = object.signature ?? Buffer.alloc(0);
    message.publicKey = object.publicKey ?? Buffer.alloc(0);
    return message;
  },
};

function createBaseMIME(): MIME {
  return { type: MIME_Type.TYPE_UNSPECIFIED, subtype: "", value: "" };
}

export const MIME = {
  encode(message: MIME, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.type !== MIME_Type.TYPE_UNSPECIFIED) {
      writer.uint32(8).int32(mIME_TypeToNumber(message.type));
    }
    if (message.subtype !== "") {
      writer.uint32(18).string(message.subtype);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MIME {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMIME();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = mIME_TypeFromJSON(reader.int32());
          break;
        case 2:
          message.subtype = reader.string();
          break;
        case 3:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MIME {
    return {
      type: isSet(object.type)
        ? mIME_TypeFromJSON(object.type)
        : MIME_Type.TYPE_UNSPECIFIED,
      subtype: isSet(object.subtype) ? String(object.subtype) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: MIME): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = mIME_TypeToJSON(message.type));
    message.subtype !== undefined && (obj.subtype = message.subtype);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MIME>, I>>(object: I): MIME {
    const message = createBaseMIME();
    message.type = object.type ?? MIME_Type.TYPE_UNSPECIFIED;
    message.subtype = object.subtype ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBasePeer(): Peer {
  return {
    sName: "",
    status: Peer_Status.STATUS_UNSPECIFIED,
    device: undefined,
    profile: undefined,
    publicKey: Buffer.alloc(0),
    peerId: "",
    lastModified: 0,
  };
}

export const Peer = {
  encode(message: Peer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sName !== "") {
      writer.uint32(10).string(message.sName);
    }
    if (message.status !== Peer_Status.STATUS_UNSPECIFIED) {
      writer.uint32(16).int32(peer_StatusToNumber(message.status));
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

  decode(input: _m0.Reader | Uint8Array, length?: number): Peer {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePeer();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sName = reader.string();
          break;
        case 2:
          message.status = peer_StatusFromJSON(reader.int32());
          break;
        case 3:
          message.device = Peer_Device.decode(reader, reader.uint32());
          break;
        case 4:
          message.profile = Profile.decode(reader, reader.uint32());
          break;
        case 5:
          message.publicKey = reader.bytes() as Buffer;
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
    return {
      sName: isSet(object.sName) ? String(object.sName) : "",
      status: isSet(object.status)
        ? peer_StatusFromJSON(object.status)
        : Peer_Status.STATUS_UNSPECIFIED,
      device: isSet(object.device)
        ? Peer_Device.fromJSON(object.device)
        : undefined,
      profile: isSet(object.profile)
        ? Profile.fromJSON(object.profile)
        : undefined,
      publicKey: isSet(object.publicKey)
        ? Buffer.from(bytesFromBase64(object.publicKey))
        : Buffer.alloc(0),
      peerId: isSet(object.peerId) ? String(object.peerId) : "",
      lastModified: isSet(object.lastModified)
        ? Number(object.lastModified)
        : 0,
    };
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
        message.publicKey !== undefined ? message.publicKey : Buffer.alloc(0)
      ));
    message.peerId !== undefined && (obj.peerId = message.peerId);
    message.lastModified !== undefined &&
      (obj.lastModified = Math.round(message.lastModified));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Peer>, I>>(object: I): Peer {
    const message = createBasePeer();
    message.sName = object.sName ?? "";
    message.status = object.status ?? Peer_Status.STATUS_UNSPECIFIED;
    message.device =
      object.device !== undefined && object.device !== null
        ? Peer_Device.fromPartial(object.device)
        : undefined;
    message.profile =
      object.profile !== undefined && object.profile !== null
        ? Profile.fromPartial(object.profile)
        : undefined;
    message.publicKey = object.publicKey ?? Buffer.alloc(0);
    message.peerId = object.peerId ?? "";
    message.lastModified = object.lastModified ?? 0;
    return message;
  },
};

function createBasePeer_Device(): Peer_Device {
  return { id: "", hostName: "", os: "", arch: "", model: "" };
}

export const Peer_Device = {
  encode(
    message: Peer_Device,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): Peer_Device {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePeer_Device();
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
    return {
      id: isSet(object.id) ? String(object.id) : "",
      hostName: isSet(object.hostName) ? String(object.hostName) : "",
      os: isSet(object.os) ? String(object.os) : "",
      arch: isSet(object.arch) ? String(object.arch) : "",
      model: isSet(object.model) ? String(object.model) : "",
    };
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

  fromPartial<I extends Exact<DeepPartial<Peer_Device>, I>>(
    object: I
  ): Peer_Device {
    const message = createBasePeer_Device();
    message.id = object.id ?? "";
    message.hostName = object.hostName ?? "";
    message.os = object.os ?? "";
    message.arch = object.arch ?? "";
    message.model = object.model ?? "";
    return message;
  },
};

function createBaseProfile(): Profile {
  return {
    sName: "",
    firstName: "",
    lastName: "",
    picture: Buffer.alloc(0),
    bio: "",
    lastModified: 0,
  };
}

export const Profile = {
  encode(
    message: Profile,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): Profile {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProfile();
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
          message.picture = reader.bytes() as Buffer;
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
    return {
      sName: isSet(object.sName) ? String(object.sName) : "",
      firstName: isSet(object.firstName) ? String(object.firstName) : "",
      lastName: isSet(object.lastName) ? String(object.lastName) : "",
      picture: isSet(object.picture)
        ? Buffer.from(bytesFromBase64(object.picture))
        : Buffer.alloc(0),
      bio: isSet(object.bio) ? String(object.bio) : "",
      lastModified: isSet(object.lastModified)
        ? Number(object.lastModified)
        : 0,
    };
  },

  toJSON(message: Profile): unknown {
    const obj: any = {};
    message.sName !== undefined && (obj.sName = message.sName);
    message.firstName !== undefined && (obj.firstName = message.firstName);
    message.lastName !== undefined && (obj.lastName = message.lastName);
    message.picture !== undefined &&
      (obj.picture = base64FromBytes(
        message.picture !== undefined ? message.picture : Buffer.alloc(0)
      ));
    message.bio !== undefined && (obj.bio = message.bio);
    message.lastModified !== undefined &&
      (obj.lastModified = Math.round(message.lastModified));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Profile>, I>>(object: I): Profile {
    const message = createBaseProfile();
    message.sName = object.sName ?? "";
    message.firstName = object.firstName ?? "";
    message.lastName = object.lastName ?? "";
    message.picture = object.picture ?? Buffer.alloc(0);
    message.bio = object.bio ?? "";
    message.lastModified = object.lastModified ?? 0;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
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
  for (const byte of arr) {
    bin.push(String.fromCharCode(byte));
  }
  return btoa(bin.join(""));
}

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<
        Exclude<keyof I, KeysOfUnion<P>>,
        never
      >;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
