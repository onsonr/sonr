/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import {
  Location,
  Profile,
  Connection,
  Environment,
  Peer,
  connectionToNumber,
  environmentToNumber,
  connectionFromJSON,
  environmentFromJSON,
  connectionToJSON,
  environmentToJSON,
} from "../../../common/core";
import { SupplyItem } from "../../../common/data";

export const protobufPackage = "node.motor.v1";

/**
 * -----------------------------------------------------------------------------
 * Motor Node API
 * -----------------------------------------------------------------------------
 * (Client) InitializeRequest Message to Establish Sonr Host/API/Room
 */
export interface InitializeRequest {
  /** Current Runtime Location */
  location?: Location;
  /** Users Contact Card */
  profile?: Profile;
  /** Internet Connection Type */
  connection: Connection;
  /** Libp2p Host config */
  hostOptions?: InitializeRequest_HostOptions;
  /** Service Config */
  serviceOptions?: InitializeRequest_ServiceOptions;
  /** File System Config */
  deviceOptions?: InitializeRequest_DeviceOptions;
  /** Environment Config */
  environment: Environment;
  /** Domain TXT Records */
  variables: { [key: string]: string };
  /** Wallet Passphrase */
  walletPassphrase: string;
}

export interface InitializeRequest_VariablesEntry {
  key: string;
  value: string;
}

/** Optional Message to Initialize FileSystem */
export interface InitializeRequest_DeviceOptions {
  /** Device ID */
  id: string;
  homeDir: string;
  supportDir: string;
  tempDir: string;
}

/** Libp2p Host Options */
export interface InitializeRequest_HostOptions {
  /** Enable QUIC Transport */
  quicTransport: boolean;
  /** Enable HTTP Transport */
  httpTransport: boolean;
  /** Enable IPv4 Only */
  ipv4Only: boolean;
  /** List of Listen Addresses (optional) */
  listenAddrs: InitializeRequest_IPAddress[];
}

/** Service Configuration */
export interface InitializeRequest_ServiceOptions {
  /** Enable Textile Client and Threads */
  textile: boolean;
  /** Enable Mailbox */
  mailbox: boolean;
  /** Enable Buckets */
  buckets: boolean;
  /** Auto Update Interval (seconds) - Default 5s */
  interval: number;
}

/** IP Address Interface */
export interface InitializeRequest_IPAddress {
  /** Name of Interface */
  name: string;
  /** IP Address of Interface */
  address: string;
  /** Wether it is a Loopback Interface */
  internal: boolean;
  /** Address Family */
  family: InitializeRequest_IPAddress_Family;
}

export enum InitializeRequest_IPAddress_Family {
  FAMILY_UNSPECIFIED = "FAMILY_UNSPECIFIED",
  /** FAMILY_IPV4 - IPv4 Address */
  FAMILY_IPV4 = "FAMILY_IPV4",
  /** FAMILY_IPV6 - IPv6 Address */
  FAMILY_IPV6 = "FAMILY_IPV6",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function initializeRequest_IPAddress_FamilyFromJSON(
  object: any
): InitializeRequest_IPAddress_Family {
  switch (object) {
    case 0:
    case "FAMILY_UNSPECIFIED":
      return InitializeRequest_IPAddress_Family.FAMILY_UNSPECIFIED;
    case 1:
    case "FAMILY_IPV4":
      return InitializeRequest_IPAddress_Family.FAMILY_IPV4;
    case 2:
    case "FAMILY_IPV6":
      return InitializeRequest_IPAddress_Family.FAMILY_IPV6;
    case -1:
    case "UNRECOGNIZED":
    default:
      return InitializeRequest_IPAddress_Family.UNRECOGNIZED;
  }
}

export function initializeRequest_IPAddress_FamilyToJSON(
  object: InitializeRequest_IPAddress_Family
): string {
  switch (object) {
    case InitializeRequest_IPAddress_Family.FAMILY_UNSPECIFIED:
      return "FAMILY_UNSPECIFIED";
    case InitializeRequest_IPAddress_Family.FAMILY_IPV4:
      return "FAMILY_IPV4";
    case InitializeRequest_IPAddress_Family.FAMILY_IPV6:
      return "FAMILY_IPV6";
    default:
      return "UNKNOWN";
  }
}

export function initializeRequest_IPAddress_FamilyToNumber(
  object: InitializeRequest_IPAddress_Family
): number {
  switch (object) {
    case InitializeRequest_IPAddress_Family.FAMILY_UNSPECIFIED:
      return 0;
    case InitializeRequest_IPAddress_Family.FAMILY_IPV4:
      return 1;
    case InitializeRequest_IPAddress_Family.FAMILY_IPV6:
      return 2;
    default:
      return 0;
  }
}

/** (Client) ShareRequest is request to share supplied files/urls with a peer */
export interface ShareRequest {
  /** Peer to Share with */
  peer?: Peer;
  /** Supply Items to share */
  items: SupplyItem[];
}

/** (Client) DecideRequest is request to respond to a share request */
export interface DecideRequest {
  /** True if Supply is Active */
  decision: boolean;
  /** Peer to Share with */
  peer?: Peer;
}

/** (Client) SearchRequest is Message for Searching for Peer */
export interface SearchRequest {
  /** SName combined with Device ID and Hashed */
  sName: string | undefined;
  /** Peer ID */
  peerId: string | undefined;
}

export interface OnLobbyRefreshRequest {}

export interface OnMailboxMessageRequest {}

export interface OnTransmitDecisionRequest {}

export interface OnTransmitInviteRequest {}

export interface OnTransmitProgressRequest {}

export interface OnTransmitCompleteRequest {}

function createBaseInitializeRequest(): InitializeRequest {
  return {
    location: undefined,
    profile: undefined,
    connection: Connection.CONNECTION_UNSPECIFIED,
    hostOptions: undefined,
    serviceOptions: undefined,
    deviceOptions: undefined,
    environment: Environment.ENVIRONMENT_UNSPECIFIED,
    variables: {},
    walletPassphrase: "",
  };
}

export const InitializeRequest = {
  encode(
    message: InitializeRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.location !== undefined) {
      Location.encode(message.location, writer.uint32(10).fork()).ldelim();
    }
    if (message.profile !== undefined) {
      Profile.encode(message.profile, writer.uint32(18).fork()).ldelim();
    }
    if (message.connection !== Connection.CONNECTION_UNSPECIFIED) {
      writer.uint32(24).int32(connectionToNumber(message.connection));
    }
    if (message.hostOptions !== undefined) {
      InitializeRequest_HostOptions.encode(
        message.hostOptions,
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.serviceOptions !== undefined) {
      InitializeRequest_ServiceOptions.encode(
        message.serviceOptions,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.deviceOptions !== undefined) {
      InitializeRequest_DeviceOptions.encode(
        message.deviceOptions,
        writer.uint32(50).fork()
      ).ldelim();
    }
    if (message.environment !== Environment.ENVIRONMENT_UNSPECIFIED) {
      writer.uint32(56).int32(environmentToNumber(message.environment));
    }
    Object.entries(message.variables).forEach(([key, value]) => {
      InitializeRequest_VariablesEntry.encode(
        { key: key as any, value },
        writer.uint32(66).fork()
      ).ldelim();
    });
    if (message.walletPassphrase !== "") {
      writer.uint32(74).string(message.walletPassphrase);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InitializeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInitializeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.location = Location.decode(reader, reader.uint32());
          break;
        case 2:
          message.profile = Profile.decode(reader, reader.uint32());
          break;
        case 3:
          message.connection = connectionFromJSON(reader.int32());
          break;
        case 4:
          message.hostOptions = InitializeRequest_HostOptions.decode(
            reader,
            reader.uint32()
          );
          break;
        case 5:
          message.serviceOptions = InitializeRequest_ServiceOptions.decode(
            reader,
            reader.uint32()
          );
          break;
        case 6:
          message.deviceOptions = InitializeRequest_DeviceOptions.decode(
            reader,
            reader.uint32()
          );
          break;
        case 7:
          message.environment = environmentFromJSON(reader.int32());
          break;
        case 8:
          const entry8 = InitializeRequest_VariablesEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry8.value !== undefined) {
            message.variables[entry8.key] = entry8.value;
          }
          break;
        case 9:
          message.walletPassphrase = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InitializeRequest {
    return {
      location: isSet(object.location)
        ? Location.fromJSON(object.location)
        : undefined,
      profile: isSet(object.profile)
        ? Profile.fromJSON(object.profile)
        : undefined,
      connection: isSet(object.connection)
        ? connectionFromJSON(object.connection)
        : Connection.CONNECTION_UNSPECIFIED,
      hostOptions: isSet(object.hostOptions)
        ? InitializeRequest_HostOptions.fromJSON(object.hostOptions)
        : undefined,
      serviceOptions: isSet(object.serviceOptions)
        ? InitializeRequest_ServiceOptions.fromJSON(object.serviceOptions)
        : undefined,
      deviceOptions: isSet(object.deviceOptions)
        ? InitializeRequest_DeviceOptions.fromJSON(object.deviceOptions)
        : undefined,
      environment: isSet(object.environment)
        ? environmentFromJSON(object.environment)
        : Environment.ENVIRONMENT_UNSPECIFIED,
      variables: isObject(object.variables)
        ? Object.entries(object.variables).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
      walletPassphrase: isSet(object.walletPassphrase)
        ? String(object.walletPassphrase)
        : "",
    };
  },

  toJSON(message: InitializeRequest): unknown {
    const obj: any = {};
    message.location !== undefined &&
      (obj.location = message.location
        ? Location.toJSON(message.location)
        : undefined);
    message.profile !== undefined &&
      (obj.profile = message.profile
        ? Profile.toJSON(message.profile)
        : undefined);
    message.connection !== undefined &&
      (obj.connection = connectionToJSON(message.connection));
    message.hostOptions !== undefined &&
      (obj.hostOptions = message.hostOptions
        ? InitializeRequest_HostOptions.toJSON(message.hostOptions)
        : undefined);
    message.serviceOptions !== undefined &&
      (obj.serviceOptions = message.serviceOptions
        ? InitializeRequest_ServiceOptions.toJSON(message.serviceOptions)
        : undefined);
    message.deviceOptions !== undefined &&
      (obj.deviceOptions = message.deviceOptions
        ? InitializeRequest_DeviceOptions.toJSON(message.deviceOptions)
        : undefined);
    message.environment !== undefined &&
      (obj.environment = environmentToJSON(message.environment));
    obj.variables = {};
    if (message.variables) {
      Object.entries(message.variables).forEach(([k, v]) => {
        obj.variables[k] = v;
      });
    }
    message.walletPassphrase !== undefined &&
      (obj.walletPassphrase = message.walletPassphrase);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<InitializeRequest>, I>>(
    object: I
  ): InitializeRequest {
    const message = createBaseInitializeRequest();
    message.location =
      object.location !== undefined && object.location !== null
        ? Location.fromPartial(object.location)
        : undefined;
    message.profile =
      object.profile !== undefined && object.profile !== null
        ? Profile.fromPartial(object.profile)
        : undefined;
    message.connection = object.connection ?? Connection.CONNECTION_UNSPECIFIED;
    message.hostOptions =
      object.hostOptions !== undefined && object.hostOptions !== null
        ? InitializeRequest_HostOptions.fromPartial(object.hostOptions)
        : undefined;
    message.serviceOptions =
      object.serviceOptions !== undefined && object.serviceOptions !== null
        ? InitializeRequest_ServiceOptions.fromPartial(object.serviceOptions)
        : undefined;
    message.deviceOptions =
      object.deviceOptions !== undefined && object.deviceOptions !== null
        ? InitializeRequest_DeviceOptions.fromPartial(object.deviceOptions)
        : undefined;
    message.environment =
      object.environment ?? Environment.ENVIRONMENT_UNSPECIFIED;
    message.variables = Object.entries(object.variables ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    message.walletPassphrase = object.walletPassphrase ?? "";
    return message;
  },
};

function createBaseInitializeRequest_VariablesEntry(): InitializeRequest_VariablesEntry {
  return { key: "", value: "" };
}

export const InitializeRequest_VariablesEntry = {
  encode(
    message: InitializeRequest_VariablesEntry,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): InitializeRequest_VariablesEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInitializeRequest_VariablesEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InitializeRequest_VariablesEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: InitializeRequest_VariablesEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<InitializeRequest_VariablesEntry>, I>
  >(object: I): InitializeRequest_VariablesEntry {
    const message = createBaseInitializeRequest_VariablesEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseInitializeRequest_DeviceOptions(): InitializeRequest_DeviceOptions {
  return { id: "", homeDir: "", supportDir: "", tempDir: "" };
}

export const InitializeRequest_DeviceOptions = {
  encode(
    message: InitializeRequest_DeviceOptions,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.homeDir !== "") {
      writer.uint32(18).string(message.homeDir);
    }
    if (message.supportDir !== "") {
      writer.uint32(26).string(message.supportDir);
    }
    if (message.tempDir !== "") {
      writer.uint32(34).string(message.tempDir);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): InitializeRequest_DeviceOptions {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInitializeRequest_DeviceOptions();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.homeDir = reader.string();
          break;
        case 3:
          message.supportDir = reader.string();
          break;
        case 4:
          message.tempDir = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InitializeRequest_DeviceOptions {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      homeDir: isSet(object.homeDir) ? String(object.homeDir) : "",
      supportDir: isSet(object.supportDir) ? String(object.supportDir) : "",
      tempDir: isSet(object.tempDir) ? String(object.tempDir) : "",
    };
  },

  toJSON(message: InitializeRequest_DeviceOptions): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.homeDir !== undefined && (obj.homeDir = message.homeDir);
    message.supportDir !== undefined && (obj.supportDir = message.supportDir);
    message.tempDir !== undefined && (obj.tempDir = message.tempDir);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<InitializeRequest_DeviceOptions>, I>>(
    object: I
  ): InitializeRequest_DeviceOptions {
    const message = createBaseInitializeRequest_DeviceOptions();
    message.id = object.id ?? "";
    message.homeDir = object.homeDir ?? "";
    message.supportDir = object.supportDir ?? "";
    message.tempDir = object.tempDir ?? "";
    return message;
  },
};

function createBaseInitializeRequest_HostOptions(): InitializeRequest_HostOptions {
  return {
    quicTransport: false,
    httpTransport: false,
    ipv4Only: false,
    listenAddrs: [],
  };
}

export const InitializeRequest_HostOptions = {
  encode(
    message: InitializeRequest_HostOptions,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.quicTransport === true) {
      writer.uint32(8).bool(message.quicTransport);
    }
    if (message.httpTransport === true) {
      writer.uint32(16).bool(message.httpTransport);
    }
    if (message.ipv4Only === true) {
      writer.uint32(24).bool(message.ipv4Only);
    }
    for (const v of message.listenAddrs) {
      InitializeRequest_IPAddress.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): InitializeRequest_HostOptions {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInitializeRequest_HostOptions();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.quicTransport = reader.bool();
          break;
        case 2:
          message.httpTransport = reader.bool();
          break;
        case 3:
          message.ipv4Only = reader.bool();
          break;
        case 4:
          message.listenAddrs.push(
            InitializeRequest_IPAddress.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InitializeRequest_HostOptions {
    return {
      quicTransport: isSet(object.quicTransport)
        ? Boolean(object.quicTransport)
        : false,
      httpTransport: isSet(object.httpTransport)
        ? Boolean(object.httpTransport)
        : false,
      ipv4Only: isSet(object.ipv4Only) ? Boolean(object.ipv4Only) : false,
      listenAddrs: Array.isArray(object?.listenAddrs)
        ? object.listenAddrs.map((e: any) =>
            InitializeRequest_IPAddress.fromJSON(e)
          )
        : [],
    };
  },

  toJSON(message: InitializeRequest_HostOptions): unknown {
    const obj: any = {};
    message.quicTransport !== undefined &&
      (obj.quicTransport = message.quicTransport);
    message.httpTransport !== undefined &&
      (obj.httpTransport = message.httpTransport);
    message.ipv4Only !== undefined && (obj.ipv4Only = message.ipv4Only);
    if (message.listenAddrs) {
      obj.listenAddrs = message.listenAddrs.map((e) =>
        e ? InitializeRequest_IPAddress.toJSON(e) : undefined
      );
    } else {
      obj.listenAddrs = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<InitializeRequest_HostOptions>, I>>(
    object: I
  ): InitializeRequest_HostOptions {
    const message = createBaseInitializeRequest_HostOptions();
    message.quicTransport = object.quicTransport ?? false;
    message.httpTransport = object.httpTransport ?? false;
    message.ipv4Only = object.ipv4Only ?? false;
    message.listenAddrs =
      object.listenAddrs?.map((e) =>
        InitializeRequest_IPAddress.fromPartial(e)
      ) || [];
    return message;
  },
};

function createBaseInitializeRequest_ServiceOptions(): InitializeRequest_ServiceOptions {
  return { textile: false, mailbox: false, buckets: false, interval: 0 };
}

export const InitializeRequest_ServiceOptions = {
  encode(
    message: InitializeRequest_ServiceOptions,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.textile === true) {
      writer.uint32(8).bool(message.textile);
    }
    if (message.mailbox === true) {
      writer.uint32(16).bool(message.mailbox);
    }
    if (message.buckets === true) {
      writer.uint32(24).bool(message.buckets);
    }
    if (message.interval !== 0) {
      writer.uint32(32).int32(message.interval);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): InitializeRequest_ServiceOptions {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInitializeRequest_ServiceOptions();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.textile = reader.bool();
          break;
        case 2:
          message.mailbox = reader.bool();
          break;
        case 3:
          message.buckets = reader.bool();
          break;
        case 4:
          message.interval = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InitializeRequest_ServiceOptions {
    return {
      textile: isSet(object.textile) ? Boolean(object.textile) : false,
      mailbox: isSet(object.mailbox) ? Boolean(object.mailbox) : false,
      buckets: isSet(object.buckets) ? Boolean(object.buckets) : false,
      interval: isSet(object.interval) ? Number(object.interval) : 0,
    };
  },

  toJSON(message: InitializeRequest_ServiceOptions): unknown {
    const obj: any = {};
    message.textile !== undefined && (obj.textile = message.textile);
    message.mailbox !== undefined && (obj.mailbox = message.mailbox);
    message.buckets !== undefined && (obj.buckets = message.buckets);
    message.interval !== undefined &&
      (obj.interval = Math.round(message.interval));
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<InitializeRequest_ServiceOptions>, I>
  >(object: I): InitializeRequest_ServiceOptions {
    const message = createBaseInitializeRequest_ServiceOptions();
    message.textile = object.textile ?? false;
    message.mailbox = object.mailbox ?? false;
    message.buckets = object.buckets ?? false;
    message.interval = object.interval ?? 0;
    return message;
  },
};

function createBaseInitializeRequest_IPAddress(): InitializeRequest_IPAddress {
  return {
    name: "",
    address: "",
    internal: false,
    family: InitializeRequest_IPAddress_Family.FAMILY_UNSPECIFIED,
  };
}

export const InitializeRequest_IPAddress = {
  encode(
    message: InitializeRequest_IPAddress,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.internal === true) {
      writer.uint32(24).bool(message.internal);
    }
    if (
      message.family !== InitializeRequest_IPAddress_Family.FAMILY_UNSPECIFIED
    ) {
      writer
        .uint32(32)
        .int32(initializeRequest_IPAddress_FamilyToNumber(message.family));
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): InitializeRequest_IPAddress {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInitializeRequest_IPAddress();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.internal = reader.bool();
          break;
        case 4:
          message.family = initializeRequest_IPAddress_FamilyFromJSON(
            reader.int32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InitializeRequest_IPAddress {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      address: isSet(object.address) ? String(object.address) : "",
      internal: isSet(object.internal) ? Boolean(object.internal) : false,
      family: isSet(object.family)
        ? initializeRequest_IPAddress_FamilyFromJSON(object.family)
        : InitializeRequest_IPAddress_Family.FAMILY_UNSPECIFIED,
    };
  },

  toJSON(message: InitializeRequest_IPAddress): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.address !== undefined && (obj.address = message.address);
    message.internal !== undefined && (obj.internal = message.internal);
    message.family !== undefined &&
      (obj.family = initializeRequest_IPAddress_FamilyToJSON(message.family));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<InitializeRequest_IPAddress>, I>>(
    object: I
  ): InitializeRequest_IPAddress {
    const message = createBaseInitializeRequest_IPAddress();
    message.name = object.name ?? "";
    message.address = object.address ?? "";
    message.internal = object.internal ?? false;
    message.family =
      object.family ?? InitializeRequest_IPAddress_Family.FAMILY_UNSPECIFIED;
    return message;
  },
};

function createBaseShareRequest(): ShareRequest {
  return { peer: undefined, items: [] };
}

export const ShareRequest = {
  encode(
    message: ShareRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.peer !== undefined) {
      Peer.encode(message.peer, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.items) {
      SupplyItem.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ShareRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseShareRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.peer = Peer.decode(reader, reader.uint32());
          break;
        case 2:
          message.items.push(SupplyItem.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ShareRequest {
    return {
      peer: isSet(object.peer) ? Peer.fromJSON(object.peer) : undefined,
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => SupplyItem.fromJSON(e))
        : [],
    };
  },

  toJSON(message: ShareRequest): unknown {
    const obj: any = {};
    message.peer !== undefined &&
      (obj.peer = message.peer ? Peer.toJSON(message.peer) : undefined);
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? SupplyItem.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ShareRequest>, I>>(
    object: I
  ): ShareRequest {
    const message = createBaseShareRequest();
    message.peer =
      object.peer !== undefined && object.peer !== null
        ? Peer.fromPartial(object.peer)
        : undefined;
    message.items = object.items?.map((e) => SupplyItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseDecideRequest(): DecideRequest {
  return { decision: false, peer: undefined };
}

export const DecideRequest = {
  encode(
    message: DecideRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.decision === true) {
      writer.uint32(8).bool(message.decision);
    }
    if (message.peer !== undefined) {
      Peer.encode(message.peer, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DecideRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDecideRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.decision = reader.bool();
          break;
        case 2:
          message.peer = Peer.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DecideRequest {
    return {
      decision: isSet(object.decision) ? Boolean(object.decision) : false,
      peer: isSet(object.peer) ? Peer.fromJSON(object.peer) : undefined,
    };
  },

  toJSON(message: DecideRequest): unknown {
    const obj: any = {};
    message.decision !== undefined && (obj.decision = message.decision);
    message.peer !== undefined &&
      (obj.peer = message.peer ? Peer.toJSON(message.peer) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DecideRequest>, I>>(
    object: I
  ): DecideRequest {
    const message = createBaseDecideRequest();
    message.decision = object.decision ?? false;
    message.peer =
      object.peer !== undefined && object.peer !== null
        ? Peer.fromPartial(object.peer)
        : undefined;
    return message;
  },
};

function createBaseSearchRequest(): SearchRequest {
  return { sName: undefined, peerId: undefined };
}

export const SearchRequest = {
  encode(
    message: SearchRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.sName !== undefined) {
      writer.uint32(10).string(message.sName);
    }
    if (message.peerId !== undefined) {
      writer.uint32(18).string(message.peerId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SearchRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSearchRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sName = reader.string();
          break;
        case 2:
          message.peerId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SearchRequest {
    return {
      sName: isSet(object.sName) ? String(object.sName) : undefined,
      peerId: isSet(object.peerId) ? String(object.peerId) : undefined,
    };
  },

  toJSON(message: SearchRequest): unknown {
    const obj: any = {};
    message.sName !== undefined && (obj.sName = message.sName);
    message.peerId !== undefined && (obj.peerId = message.peerId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SearchRequest>, I>>(
    object: I
  ): SearchRequest {
    const message = createBaseSearchRequest();
    message.sName = object.sName ?? undefined;
    message.peerId = object.peerId ?? undefined;
    return message;
  },
};

function createBaseOnLobbyRefreshRequest(): OnLobbyRefreshRequest {
  return {};
}

export const OnLobbyRefreshRequest = {
  encode(
    _: OnLobbyRefreshRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): OnLobbyRefreshRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOnLobbyRefreshRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): OnLobbyRefreshRequest {
    return {};
  },

  toJSON(_: OnLobbyRefreshRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OnLobbyRefreshRequest>, I>>(
    _: I
  ): OnLobbyRefreshRequest {
    const message = createBaseOnLobbyRefreshRequest();
    return message;
  },
};

function createBaseOnMailboxMessageRequest(): OnMailboxMessageRequest {
  return {};
}

export const OnMailboxMessageRequest = {
  encode(
    _: OnMailboxMessageRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): OnMailboxMessageRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOnMailboxMessageRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): OnMailboxMessageRequest {
    return {};
  },

  toJSON(_: OnMailboxMessageRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OnMailboxMessageRequest>, I>>(
    _: I
  ): OnMailboxMessageRequest {
    const message = createBaseOnMailboxMessageRequest();
    return message;
  },
};

function createBaseOnTransmitDecisionRequest(): OnTransmitDecisionRequest {
  return {};
}

export const OnTransmitDecisionRequest = {
  encode(
    _: OnTransmitDecisionRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): OnTransmitDecisionRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOnTransmitDecisionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): OnTransmitDecisionRequest {
    return {};
  },

  toJSON(_: OnTransmitDecisionRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OnTransmitDecisionRequest>, I>>(
    _: I
  ): OnTransmitDecisionRequest {
    const message = createBaseOnTransmitDecisionRequest();
    return message;
  },
};

function createBaseOnTransmitInviteRequest(): OnTransmitInviteRequest {
  return {};
}

export const OnTransmitInviteRequest = {
  encode(
    _: OnTransmitInviteRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): OnTransmitInviteRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOnTransmitInviteRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): OnTransmitInviteRequest {
    return {};
  },

  toJSON(_: OnTransmitInviteRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OnTransmitInviteRequest>, I>>(
    _: I
  ): OnTransmitInviteRequest {
    const message = createBaseOnTransmitInviteRequest();
    return message;
  },
};

function createBaseOnTransmitProgressRequest(): OnTransmitProgressRequest {
  return {};
}

export const OnTransmitProgressRequest = {
  encode(
    _: OnTransmitProgressRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): OnTransmitProgressRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOnTransmitProgressRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): OnTransmitProgressRequest {
    return {};
  },

  toJSON(_: OnTransmitProgressRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OnTransmitProgressRequest>, I>>(
    _: I
  ): OnTransmitProgressRequest {
    const message = createBaseOnTransmitProgressRequest();
    return message;
  },
};

function createBaseOnTransmitCompleteRequest(): OnTransmitCompleteRequest {
  return {};
}

export const OnTransmitCompleteRequest = {
  encode(
    _: OnTransmitCompleteRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): OnTransmitCompleteRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOnTransmitCompleteRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): OnTransmitCompleteRequest {
    return {};
  },

  toJSON(_: OnTransmitCompleteRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OnTransmitCompleteRequest>, I>>(
    _: I
  ): OnTransmitCompleteRequest {
    const message = createBaseOnTransmitCompleteRequest();
    return message;
  },
};

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

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
