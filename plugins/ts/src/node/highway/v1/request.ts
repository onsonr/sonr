/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import {
  BucketType,
  ObjectField,
  bucketTypeToNumber,
  bucketTypeFromJSON,
  bucketTypeToJSON,
} from "../../../common/object";

export const protobufPackage = "node.highway.v1";

/** / This file contains service for the Node RPC Server */

/** AccessNameRequest is a request to get details from the ".snr" name of a peer */
export interface AccessNameRequest {
  /** The name of the peer to get the details from */
  name: string;
  /** The public key of the peer to get the details from */
  publicKey: string;
}

/** RegisterNameRequest is a request to register a name with the ".snr" name of a peer */
export interface RegisterNameRequest {
  /** The name of the peer to register the name with */
  deviceId: string;
  /** The Operating System of the peer to register the name with */
  os: string;
  /** The model of the peer to register the name with */
  model: string;
  /** The architecture of the peer to register the name with */
  arch: string;
  /** The public key of the peer to register the name with */
  publicKey: string;
  /** The name to register */
  nameToRegister: string;
}

/** UpdateNameRequest is a request to update the ".snr" name of a peer */
export interface UpdateNameRequest {
  /** The name of the peer to update the name of */
  name: string;
  /** The Updated Metadata */
  metadata: { [key: string]: string };
}

export interface UpdateNameRequest_MetadataEntry {
  key: string;
  value: string;
}

/** AccessServiceRequest is a request to get the service details of a peer */
export interface AccessServiceRequest {
  /** The name of the peer to get the service details of */
  did: string;
  /** The metadata for any service information required */
  metadata: { [key: string]: string };
}

export interface AccessServiceRequest_MetadataEntry {
  key: string;
  value: string;
}

/** RegisterServiceRequest is a request to register a service with a peer */
export interface RegisterServiceRequest {
  /** The name of the peer to register the service with */
  serviceName: string;
  /** The configuration for the service */
  configuration: { [key: string]: string };
  /** The public key of the peer to register the service with */
  publicKey: string;
}

export interface RegisterServiceRequest_ConfigurationEntry {
  key: string;
  value: string;
}

/** UpdateServiceRequest is a request to update the service details of a peer */
export interface UpdateServiceRequest {
  /** The name of the peer to update the service details of */
  did: string;
  /** The updated configuration for the service */
  configuration: { [key: string]: string };
  /** The metadata for any service information required */
  metadata: { [key: string]: string };
}

export interface UpdateServiceRequest_ConfigurationEntry {
  key: string;
  value: string;
}

export interface UpdateServiceRequest_MetadataEntry {
  key: string;
  value: string;
}

/** CreateChannelRequest is the request to create a new channel */
export interface CreateChannelRequest {
  /** Name is the name of the channel */
  name: string;
  /** Description is the description of the channel */
  description: string;
  /** Owners is the list of provisioned nodes for the channel */
  owners: string[];
}

/** ReadChannelRequest is the request to read a channel */
export interface ReadChannelRequest {
  /** DID is the identifier of the channel */
  did: string;
}

/** UpdateChannelRequest is the request to update a channel */
export interface UpdateChannelRequest {
  /** Did is the DID of the channel */
  did: string;
  /** Metadata is the metadata of the channel thats being updated */
  metadata: { [key: string]: string };
}

export interface UpdateChannelRequest_MetadataEntry {
  key: string;
  value: string;
}

/** DeleteChannelRequest is the request to delete a channel */
export interface DeleteChannelRequest {
  /** Did is the DID of the channel */
  did: string;
  /** Metadata is the metadata of the channel thats being deleted */
  metadata: { [key: string]: string };
  /** Public key of the node that is deleting the channel */
  publicKey: string;
}

export interface DeleteChannelRequest_MetadataEntry {
  key: string;
  value: string;
}

/** ListenChannelRequest is the request to subscribe to a channel */
export interface ListenChannelRequest {
  /** Name is the name of the channel */
  did: string;
  /** Metadata is additional metadata for the channel */
  metadata: { [key: string]: string };
}

export interface ListenChannelRequest_MetadataEntry {
  key: string;
  value: string;
}

/** CreateBucketRequest is the request to create a new bucket */
export interface CreateBucketRequest {
  /** Label is the human-readable name of the bucket */
  label: string;
  /** Description is the description of the bucket */
  description: string;
  /** Owners is the list of provisioned nodes for the bucket */
  bucketType: BucketType;
  /** Metadata is the metadata of the bucket thats being created */
  metadata: { [key: string]: string };
}

export interface CreateBucketRequest_MetadataEntry {
  key: string;
  value: string;
}

/** ReadBucketRequest is the request to read a bucket */
export interface ReadBucketRequest {
  /** DID is the identifier of the bucket */
  did: string;
  /** Metadata is the metadata of the bucket thats being read */
  metadata: { [key: string]: string };
}

export interface ReadBucketRequest_MetadataEntry {
  key: string;
  value: string;
}

/** UpdateBucketRequest is the request to update a bucket */
export interface UpdateBucketRequest {
  /** DID is the DID of the bucket */
  did: string;
  /** Label is the human-readable name of the bucket */
  label: string;
  /** Description is the description of the bucket */
  description: string;
  /** Metadata is the metadata of the bucket thats being updated */
  metadata: { [key: string]: string };
}

export interface UpdateBucketRequest_MetadataEntry {
  key: string;
  value: string;
}

/** DeleteBucketRequest is the request to delete a bucket */
export interface DeleteBucketRequest {
  /** DID is the DID of the bucket */
  did: string;
  /** Metadata is the metadata of the bucket thats being deleted */
  metadata: { [key: string]: string };
  /** Public key of the node that is deleting the bucket */
  publicKey: string;
}

export interface DeleteBucketRequest_MetadataEntry {
  key: string;
  value: string;
}

/** ListenBucketRequest is the request to subscribe to a bucket */
export interface ListenBucketRequest {
  /** DID is the DID of the bucket */
  did: string;
  /** Metadata is the metadata of the bucket thats being listened to */
  metadata: { [key: string]: string };
}

export interface ListenBucketRequest_MetadataEntry {
  key: string;
  value: string;
}

/** CreateObjectRequest is the request to create a new object */
export interface CreateObjectRequest {
  /** Label is the label of the object */
  label: string;
  /** Name is the name of the object */
  name: string;
  /** Fields is the fields of the object */
  fields: ObjectField[];
}

/** ReadObjectRequest is the request to read an object */
export interface ReadObjectRequest {
  /** DID is the identifier of the object */
  did: string;
}

/** UpdateObjectRequest is the request to update an object */
export interface UpdateObjectRequest {
  /** DID is the identifier of the object */
  did: string;
  /** Fields is the fields of the object */
  fields: ObjectField[];
}

/** DeleteObjectRequest is the request to delete an object */
export interface DeleteObjectRequest {
  /** DID is the identifier of the object */
  did: string;
  /** Metadata is the metadata of the object thats being deleted */
  metadata: { [key: string]: string };
  /** Public key of the node that is deleting the object */
  publicKey: string;
}

export interface DeleteObjectRequest_MetadataEntry {
  key: string;
  value: string;
}

/** UploadBlobRequest is the request to upload a blob */
export interface UploadBlobRequest {
  /** Label is the label of the blob */
  label: string;
  /** Path is the path of the blob */
  path: string;
  /** Bucket is the bucket of the blob */
  bucketDid: string;
  /** Size is the size of the blob */
  size: number;
  /** LastModified is the last modified time of the blob */
  lastModified: number;
}

/** DownloadBlobRequest is the request to download a blob */
export interface DownloadBlobRequest {
  /** DID is the identifier of the blob */
  did: string;
  /** Out Path is the download path of the blob */
  outPath: string;
}

/** SyncBlobRequest is the request to sync a blob */
export interface SyncBlobRequest {
  /** DID is the identifier of the blob */
  did: string;
  /** Destination DID is the identifier of the destination service storage */
  destinationDid: string;
  /** Path is the location of the blob */
  path: string;
}

/** DeleteBlobRequest is the request to delete a blob */
export interface DeleteBlobRequest {
  /** DID is the identifier of the blob */
  did: string;
  /** Metadata is the metadata of the blob thats being deleted */
  metadata: { [key: string]: string };
  /** Public key of the node that is deleting the blob */
  publicKey: string;
}

export interface DeleteBlobRequest_MetadataEntry {
  key: string;
  value: string;
}

/** ParseDidRequest is the request to convert a string to a DID object */
export interface ParseDidRequest {
  /** DID is the DID of the DID */
  did: string;
  /** Metadata is the metadata of the blob thats being deleted */
  metadata: { [key: string]: string };
}

export interface ParseDidRequest_MetadataEntry {
  key: string;
  value: string;
}

/** ResolveDidRequest is the request to resolve a DID */
export interface ResolveDidRequest {
  /** DID is the DID of the DID */
  did: string;
  /** Metadata is the metadata of the blob thats being deleted */
  metadata: { [key: string]: string };
}

export interface ResolveDidRequest_MetadataEntry {
  key: string;
  value: string;
}

function createBaseAccessNameRequest(): AccessNameRequest {
  return { name: "", publicKey: "" };
}

export const AccessNameRequest = {
  encode(
    message: AccessNameRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.publicKey !== "") {
      writer.uint32(18).string(message.publicKey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AccessNameRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAccessNameRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.publicKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AccessNameRequest {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      publicKey: isSet(object.publicKey) ? String(object.publicKey) : "",
    };
  },

  toJSON(message: AccessNameRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.publicKey !== undefined && (obj.publicKey = message.publicKey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AccessNameRequest>, I>>(
    object: I
  ): AccessNameRequest {
    const message = createBaseAccessNameRequest();
    message.name = object.name ?? "";
    message.publicKey = object.publicKey ?? "";
    return message;
  },
};

function createBaseRegisterNameRequest(): RegisterNameRequest {
  return {
    deviceId: "",
    os: "",
    model: "",
    arch: "",
    publicKey: "",
    nameToRegister: "",
  };
}

export const RegisterNameRequest = {
  encode(
    message: RegisterNameRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.deviceId !== "") {
      writer.uint32(10).string(message.deviceId);
    }
    if (message.os !== "") {
      writer.uint32(18).string(message.os);
    }
    if (message.model !== "") {
      writer.uint32(26).string(message.model);
    }
    if (message.arch !== "") {
      writer.uint32(34).string(message.arch);
    }
    if (message.publicKey !== "") {
      writer.uint32(42).string(message.publicKey);
    }
    if (message.nameToRegister !== "") {
      writer.uint32(50).string(message.nameToRegister);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterNameRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterNameRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.deviceId = reader.string();
          break;
        case 2:
          message.os = reader.string();
          break;
        case 3:
          message.model = reader.string();
          break;
        case 4:
          message.arch = reader.string();
          break;
        case 5:
          message.publicKey = reader.string();
          break;
        case 6:
          message.nameToRegister = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegisterNameRequest {
    return {
      deviceId: isSet(object.deviceId) ? String(object.deviceId) : "",
      os: isSet(object.os) ? String(object.os) : "",
      model: isSet(object.model) ? String(object.model) : "",
      arch: isSet(object.arch) ? String(object.arch) : "",
      publicKey: isSet(object.publicKey) ? String(object.publicKey) : "",
      nameToRegister: isSet(object.nameToRegister)
        ? String(object.nameToRegister)
        : "",
    };
  },

  toJSON(message: RegisterNameRequest): unknown {
    const obj: any = {};
    message.deviceId !== undefined && (obj.deviceId = message.deviceId);
    message.os !== undefined && (obj.os = message.os);
    message.model !== undefined && (obj.model = message.model);
    message.arch !== undefined && (obj.arch = message.arch);
    message.publicKey !== undefined && (obj.publicKey = message.publicKey);
    message.nameToRegister !== undefined &&
      (obj.nameToRegister = message.nameToRegister);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterNameRequest>, I>>(
    object: I
  ): RegisterNameRequest {
    const message = createBaseRegisterNameRequest();
    message.deviceId = object.deviceId ?? "";
    message.os = object.os ?? "";
    message.model = object.model ?? "";
    message.arch = object.arch ?? "";
    message.publicKey = object.publicKey ?? "";
    message.nameToRegister = object.nameToRegister ?? "";
    return message;
  },
};

function createBaseUpdateNameRequest(): UpdateNameRequest {
  return { name: "", metadata: {} };
}

export const UpdateNameRequest = {
  encode(
    message: UpdateNameRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      UpdateNameRequest_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(18).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateNameRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateNameRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          const entry2 = UpdateNameRequest_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry2.value !== undefined) {
            message.metadata[entry2.key] = entry2.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateNameRequest {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
    };
  },

  toJSON(message: UpdateNameRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateNameRequest>, I>>(
    object: I
  ): UpdateNameRequest {
    const message = createBaseUpdateNameRequest();
    message.name = object.name ?? "";
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseUpdateNameRequest_MetadataEntry(): UpdateNameRequest_MetadataEntry {
  return { key: "", value: "" };
}

export const UpdateNameRequest_MetadataEntry = {
  encode(
    message: UpdateNameRequest_MetadataEntry,
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
  ): UpdateNameRequest_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateNameRequest_MetadataEntry();
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

  fromJSON(object: any): UpdateNameRequest_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: UpdateNameRequest_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateNameRequest_MetadataEntry>, I>>(
    object: I
  ): UpdateNameRequest_MetadataEntry {
    const message = createBaseUpdateNameRequest_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseAccessServiceRequest(): AccessServiceRequest {
  return { did: "", metadata: {} };
}

export const AccessServiceRequest = {
  encode(
    message: AccessServiceRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      AccessServiceRequest_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(18).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): AccessServiceRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAccessServiceRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          const entry2 = AccessServiceRequest_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry2.value !== undefined) {
            message.metadata[entry2.key] = entry2.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AccessServiceRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
    };
  },

  toJSON(message: AccessServiceRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AccessServiceRequest>, I>>(
    object: I
  ): AccessServiceRequest {
    const message = createBaseAccessServiceRequest();
    message.did = object.did ?? "";
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseAccessServiceRequest_MetadataEntry(): AccessServiceRequest_MetadataEntry {
  return { key: "", value: "" };
}

export const AccessServiceRequest_MetadataEntry = {
  encode(
    message: AccessServiceRequest_MetadataEntry,
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
  ): AccessServiceRequest_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAccessServiceRequest_MetadataEntry();
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

  fromJSON(object: any): AccessServiceRequest_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: AccessServiceRequest_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<AccessServiceRequest_MetadataEntry>, I>
  >(object: I): AccessServiceRequest_MetadataEntry {
    const message = createBaseAccessServiceRequest_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseRegisterServiceRequest(): RegisterServiceRequest {
  return { serviceName: "", configuration: {}, publicKey: "" };
}

export const RegisterServiceRequest = {
  encode(
    message: RegisterServiceRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.serviceName !== "") {
      writer.uint32(10).string(message.serviceName);
    }
    Object.entries(message.configuration).forEach(([key, value]) => {
      RegisterServiceRequest_ConfigurationEntry.encode(
        { key: key as any, value },
        writer.uint32(18).fork()
      ).ldelim();
    });
    if (message.publicKey !== "") {
      writer.uint32(26).string(message.publicKey);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): RegisterServiceRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterServiceRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.serviceName = reader.string();
          break;
        case 2:
          const entry2 = RegisterServiceRequest_ConfigurationEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry2.value !== undefined) {
            message.configuration[entry2.key] = entry2.value;
          }
          break;
        case 3:
          message.publicKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegisterServiceRequest {
    return {
      serviceName: isSet(object.serviceName) ? String(object.serviceName) : "",
      configuration: isObject(object.configuration)
        ? Object.entries(object.configuration).reduce<{
            [key: string]: string;
          }>((acc, [key, value]) => {
            acc[key] = String(value);
            return acc;
          }, {})
        : {},
      publicKey: isSet(object.publicKey) ? String(object.publicKey) : "",
    };
  },

  toJSON(message: RegisterServiceRequest): unknown {
    const obj: any = {};
    message.serviceName !== undefined &&
      (obj.serviceName = message.serviceName);
    obj.configuration = {};
    if (message.configuration) {
      Object.entries(message.configuration).forEach(([k, v]) => {
        obj.configuration[k] = v;
      });
    }
    message.publicKey !== undefined && (obj.publicKey = message.publicKey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterServiceRequest>, I>>(
    object: I
  ): RegisterServiceRequest {
    const message = createBaseRegisterServiceRequest();
    message.serviceName = object.serviceName ?? "";
    message.configuration = Object.entries(object.configuration ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    message.publicKey = object.publicKey ?? "";
    return message;
  },
};

function createBaseRegisterServiceRequest_ConfigurationEntry(): RegisterServiceRequest_ConfigurationEntry {
  return { key: "", value: "" };
}

export const RegisterServiceRequest_ConfigurationEntry = {
  encode(
    message: RegisterServiceRequest_ConfigurationEntry,
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
  ): RegisterServiceRequest_ConfigurationEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterServiceRequest_ConfigurationEntry();
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

  fromJSON(object: any): RegisterServiceRequest_ConfigurationEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: RegisterServiceRequest_ConfigurationEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<RegisterServiceRequest_ConfigurationEntry>, I>
  >(object: I): RegisterServiceRequest_ConfigurationEntry {
    const message = createBaseRegisterServiceRequest_ConfigurationEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseUpdateServiceRequest(): UpdateServiceRequest {
  return { did: "", configuration: {}, metadata: {} };
}

export const UpdateServiceRequest = {
  encode(
    message: UpdateServiceRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    Object.entries(message.configuration).forEach(([key, value]) => {
      UpdateServiceRequest_ConfigurationEntry.encode(
        { key: key as any, value },
        writer.uint32(18).fork()
      ).ldelim();
    });
    Object.entries(message.metadata).forEach(([key, value]) => {
      UpdateServiceRequest_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(26).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): UpdateServiceRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateServiceRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          const entry2 = UpdateServiceRequest_ConfigurationEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry2.value !== undefined) {
            message.configuration[entry2.key] = entry2.value;
          }
          break;
        case 3:
          const entry3 = UpdateServiceRequest_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry3.value !== undefined) {
            message.metadata[entry3.key] = entry3.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateServiceRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      configuration: isObject(object.configuration)
        ? Object.entries(object.configuration).reduce<{
            [key: string]: string;
          }>((acc, [key, value]) => {
            acc[key] = String(value);
            return acc;
          }, {})
        : {},
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
    };
  },

  toJSON(message: UpdateServiceRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    obj.configuration = {};
    if (message.configuration) {
      Object.entries(message.configuration).forEach(([k, v]) => {
        obj.configuration[k] = v;
      });
    }
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateServiceRequest>, I>>(
    object: I
  ): UpdateServiceRequest {
    const message = createBaseUpdateServiceRequest();
    message.did = object.did ?? "";
    message.configuration = Object.entries(object.configuration ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseUpdateServiceRequest_ConfigurationEntry(): UpdateServiceRequest_ConfigurationEntry {
  return { key: "", value: "" };
}

export const UpdateServiceRequest_ConfigurationEntry = {
  encode(
    message: UpdateServiceRequest_ConfigurationEntry,
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
  ): UpdateServiceRequest_ConfigurationEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateServiceRequest_ConfigurationEntry();
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

  fromJSON(object: any): UpdateServiceRequest_ConfigurationEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: UpdateServiceRequest_ConfigurationEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<UpdateServiceRequest_ConfigurationEntry>, I>
  >(object: I): UpdateServiceRequest_ConfigurationEntry {
    const message = createBaseUpdateServiceRequest_ConfigurationEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseUpdateServiceRequest_MetadataEntry(): UpdateServiceRequest_MetadataEntry {
  return { key: "", value: "" };
}

export const UpdateServiceRequest_MetadataEntry = {
  encode(
    message: UpdateServiceRequest_MetadataEntry,
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
  ): UpdateServiceRequest_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateServiceRequest_MetadataEntry();
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

  fromJSON(object: any): UpdateServiceRequest_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: UpdateServiceRequest_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<UpdateServiceRequest_MetadataEntry>, I>
  >(object: I): UpdateServiceRequest_MetadataEntry {
    const message = createBaseUpdateServiceRequest_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseCreateChannelRequest(): CreateChannelRequest {
  return { name: "", description: "", owners: [] };
}

export const CreateChannelRequest = {
  encode(
    message: CreateChannelRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    for (const v of message.owners) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): CreateChannelRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateChannelRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.owners.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateChannelRequest {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      owners: Array.isArray(object?.owners)
        ? object.owners.map((e: any) => String(e))
        : [],
    };
  },

  toJSON(message: CreateChannelRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined &&
      (obj.description = message.description);
    if (message.owners) {
      obj.owners = message.owners.map((e) => e);
    } else {
      obj.owners = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateChannelRequest>, I>>(
    object: I
  ): CreateChannelRequest {
    const message = createBaseCreateChannelRequest();
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.owners = object.owners?.map((e) => e) || [];
    return message;
  },
};

function createBaseReadChannelRequest(): ReadChannelRequest {
  return { did: "" };
}

export const ReadChannelRequest = {
  encode(
    message: ReadChannelRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadChannelRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadChannelRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ReadChannelRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
    };
  },

  toJSON(message: ReadChannelRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ReadChannelRequest>, I>>(
    object: I
  ): ReadChannelRequest {
    const message = createBaseReadChannelRequest();
    message.did = object.did ?? "";
    return message;
  },
};

function createBaseUpdateChannelRequest(): UpdateChannelRequest {
  return { did: "", metadata: {} };
}

export const UpdateChannelRequest = {
  encode(
    message: UpdateChannelRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      UpdateChannelRequest_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(18).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): UpdateChannelRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateChannelRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          const entry2 = UpdateChannelRequest_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry2.value !== undefined) {
            message.metadata[entry2.key] = entry2.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateChannelRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
    };
  },

  toJSON(message: UpdateChannelRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateChannelRequest>, I>>(
    object: I
  ): UpdateChannelRequest {
    const message = createBaseUpdateChannelRequest();
    message.did = object.did ?? "";
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseUpdateChannelRequest_MetadataEntry(): UpdateChannelRequest_MetadataEntry {
  return { key: "", value: "" };
}

export const UpdateChannelRequest_MetadataEntry = {
  encode(
    message: UpdateChannelRequest_MetadataEntry,
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
  ): UpdateChannelRequest_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateChannelRequest_MetadataEntry();
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

  fromJSON(object: any): UpdateChannelRequest_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: UpdateChannelRequest_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<UpdateChannelRequest_MetadataEntry>, I>
  >(object: I): UpdateChannelRequest_MetadataEntry {
    const message = createBaseUpdateChannelRequest_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseDeleteChannelRequest(): DeleteChannelRequest {
  return { did: "", metadata: {}, publicKey: "" };
}

export const DeleteChannelRequest = {
  encode(
    message: DeleteChannelRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      DeleteChannelRequest_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(18).fork()
      ).ldelim();
    });
    if (message.publicKey !== "") {
      writer.uint32(26).string(message.publicKey);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): DeleteChannelRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteChannelRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          const entry2 = DeleteChannelRequest_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry2.value !== undefined) {
            message.metadata[entry2.key] = entry2.value;
          }
          break;
        case 3:
          message.publicKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteChannelRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
      publicKey: isSet(object.publicKey) ? String(object.publicKey) : "",
    };
  },

  toJSON(message: DeleteChannelRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    message.publicKey !== undefined && (obj.publicKey = message.publicKey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteChannelRequest>, I>>(
    object: I
  ): DeleteChannelRequest {
    const message = createBaseDeleteChannelRequest();
    message.did = object.did ?? "";
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    message.publicKey = object.publicKey ?? "";
    return message;
  },
};

function createBaseDeleteChannelRequest_MetadataEntry(): DeleteChannelRequest_MetadataEntry {
  return { key: "", value: "" };
}

export const DeleteChannelRequest_MetadataEntry = {
  encode(
    message: DeleteChannelRequest_MetadataEntry,
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
  ): DeleteChannelRequest_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteChannelRequest_MetadataEntry();
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

  fromJSON(object: any): DeleteChannelRequest_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: DeleteChannelRequest_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<DeleteChannelRequest_MetadataEntry>, I>
  >(object: I): DeleteChannelRequest_MetadataEntry {
    const message = createBaseDeleteChannelRequest_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseListenChannelRequest(): ListenChannelRequest {
  return { did: "", metadata: {} };
}

export const ListenChannelRequest = {
  encode(
    message: ListenChannelRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ListenChannelRequest_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(18).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ListenChannelRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListenChannelRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          const entry2 = ListenChannelRequest_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry2.value !== undefined) {
            message.metadata[entry2.key] = entry2.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListenChannelRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
    };
  },

  toJSON(message: ListenChannelRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListenChannelRequest>, I>>(
    object: I
  ): ListenChannelRequest {
    const message = createBaseListenChannelRequest();
    message.did = object.did ?? "";
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseListenChannelRequest_MetadataEntry(): ListenChannelRequest_MetadataEntry {
  return { key: "", value: "" };
}

export const ListenChannelRequest_MetadataEntry = {
  encode(
    message: ListenChannelRequest_MetadataEntry,
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
  ): ListenChannelRequest_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListenChannelRequest_MetadataEntry();
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

  fromJSON(object: any): ListenChannelRequest_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: ListenChannelRequest_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<ListenChannelRequest_MetadataEntry>, I>
  >(object: I): ListenChannelRequest_MetadataEntry {
    const message = createBaseListenChannelRequest_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseCreateBucketRequest(): CreateBucketRequest {
  return {
    label: "",
    description: "",
    bucketType: BucketType.BUCKET_TYPE_UNSPECIFIED,
    metadata: {},
  };
}

export const CreateBucketRequest = {
  encode(
    message: CreateBucketRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.bucketType !== BucketType.BUCKET_TYPE_UNSPECIFIED) {
      writer.uint32(24).int32(bucketTypeToNumber(message.bucketType));
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      CreateBucketRequest_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateBucketRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBucketRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.bucketType = bucketTypeFromJSON(reader.int32());
          break;
        case 4:
          const entry4 = CreateBucketRequest_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry4.value !== undefined) {
            message.metadata[entry4.key] = entry4.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateBucketRequest {
    return {
      label: isSet(object.label) ? String(object.label) : "",
      description: isSet(object.description) ? String(object.description) : "",
      bucketType: isSet(object.bucketType)
        ? bucketTypeFromJSON(object.bucketType)
        : BucketType.BUCKET_TYPE_UNSPECIFIED,
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
    };
  },

  toJSON(message: CreateBucketRequest): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.description !== undefined &&
      (obj.description = message.description);
    message.bucketType !== undefined &&
      (obj.bucketType = bucketTypeToJSON(message.bucketType));
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateBucketRequest>, I>>(
    object: I
  ): CreateBucketRequest {
    const message = createBaseCreateBucketRequest();
    message.label = object.label ?? "";
    message.description = object.description ?? "";
    message.bucketType =
      object.bucketType ?? BucketType.BUCKET_TYPE_UNSPECIFIED;
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseCreateBucketRequest_MetadataEntry(): CreateBucketRequest_MetadataEntry {
  return { key: "", value: "" };
}

export const CreateBucketRequest_MetadataEntry = {
  encode(
    message: CreateBucketRequest_MetadataEntry,
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
  ): CreateBucketRequest_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBucketRequest_MetadataEntry();
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

  fromJSON(object: any): CreateBucketRequest_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: CreateBucketRequest_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<CreateBucketRequest_MetadataEntry>, I>
  >(object: I): CreateBucketRequest_MetadataEntry {
    const message = createBaseCreateBucketRequest_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseReadBucketRequest(): ReadBucketRequest {
  return { did: "", metadata: {} };
}

export const ReadBucketRequest = {
  encode(
    message: ReadBucketRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ReadBucketRequest_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(18).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadBucketRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadBucketRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          const entry2 = ReadBucketRequest_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry2.value !== undefined) {
            message.metadata[entry2.key] = entry2.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ReadBucketRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
    };
  },

  toJSON(message: ReadBucketRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ReadBucketRequest>, I>>(
    object: I
  ): ReadBucketRequest {
    const message = createBaseReadBucketRequest();
    message.did = object.did ?? "";
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseReadBucketRequest_MetadataEntry(): ReadBucketRequest_MetadataEntry {
  return { key: "", value: "" };
}

export const ReadBucketRequest_MetadataEntry = {
  encode(
    message: ReadBucketRequest_MetadataEntry,
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
  ): ReadBucketRequest_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadBucketRequest_MetadataEntry();
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

  fromJSON(object: any): ReadBucketRequest_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: ReadBucketRequest_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ReadBucketRequest_MetadataEntry>, I>>(
    object: I
  ): ReadBucketRequest_MetadataEntry {
    const message = createBaseReadBucketRequest_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseUpdateBucketRequest(): UpdateBucketRequest {
  return { did: "", label: "", description: "", metadata: {} };
}

export const UpdateBucketRequest = {
  encode(
    message: UpdateBucketRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    if (message.label !== "") {
      writer.uint32(18).string(message.label);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      UpdateBucketRequest_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateBucketRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateBucketRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          message.label = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          const entry4 = UpdateBucketRequest_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry4.value !== undefined) {
            message.metadata[entry4.key] = entry4.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateBucketRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      label: isSet(object.label) ? String(object.label) : "",
      description: isSet(object.description) ? String(object.description) : "",
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
    };
  },

  toJSON(message: UpdateBucketRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    message.label !== undefined && (obj.label = message.label);
    message.description !== undefined &&
      (obj.description = message.description);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateBucketRequest>, I>>(
    object: I
  ): UpdateBucketRequest {
    const message = createBaseUpdateBucketRequest();
    message.did = object.did ?? "";
    message.label = object.label ?? "";
    message.description = object.description ?? "";
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseUpdateBucketRequest_MetadataEntry(): UpdateBucketRequest_MetadataEntry {
  return { key: "", value: "" };
}

export const UpdateBucketRequest_MetadataEntry = {
  encode(
    message: UpdateBucketRequest_MetadataEntry,
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
  ): UpdateBucketRequest_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateBucketRequest_MetadataEntry();
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

  fromJSON(object: any): UpdateBucketRequest_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: UpdateBucketRequest_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<UpdateBucketRequest_MetadataEntry>, I>
  >(object: I): UpdateBucketRequest_MetadataEntry {
    const message = createBaseUpdateBucketRequest_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseDeleteBucketRequest(): DeleteBucketRequest {
  return { did: "", metadata: {}, publicKey: "" };
}

export const DeleteBucketRequest = {
  encode(
    message: DeleteBucketRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      DeleteBucketRequest_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(18).fork()
      ).ldelim();
    });
    if (message.publicKey !== "") {
      writer.uint32(26).string(message.publicKey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteBucketRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteBucketRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          const entry2 = DeleteBucketRequest_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry2.value !== undefined) {
            message.metadata[entry2.key] = entry2.value;
          }
          break;
        case 3:
          message.publicKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteBucketRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
      publicKey: isSet(object.publicKey) ? String(object.publicKey) : "",
    };
  },

  toJSON(message: DeleteBucketRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    message.publicKey !== undefined && (obj.publicKey = message.publicKey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteBucketRequest>, I>>(
    object: I
  ): DeleteBucketRequest {
    const message = createBaseDeleteBucketRequest();
    message.did = object.did ?? "";
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    message.publicKey = object.publicKey ?? "";
    return message;
  },
};

function createBaseDeleteBucketRequest_MetadataEntry(): DeleteBucketRequest_MetadataEntry {
  return { key: "", value: "" };
}

export const DeleteBucketRequest_MetadataEntry = {
  encode(
    message: DeleteBucketRequest_MetadataEntry,
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
  ): DeleteBucketRequest_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteBucketRequest_MetadataEntry();
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

  fromJSON(object: any): DeleteBucketRequest_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: DeleteBucketRequest_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<DeleteBucketRequest_MetadataEntry>, I>
  >(object: I): DeleteBucketRequest_MetadataEntry {
    const message = createBaseDeleteBucketRequest_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseListenBucketRequest(): ListenBucketRequest {
  return { did: "", metadata: {} };
}

export const ListenBucketRequest = {
  encode(
    message: ListenBucketRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ListenBucketRequest_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(18).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListenBucketRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListenBucketRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          const entry2 = ListenBucketRequest_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry2.value !== undefined) {
            message.metadata[entry2.key] = entry2.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListenBucketRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
    };
  },

  toJSON(message: ListenBucketRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListenBucketRequest>, I>>(
    object: I
  ): ListenBucketRequest {
    const message = createBaseListenBucketRequest();
    message.did = object.did ?? "";
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseListenBucketRequest_MetadataEntry(): ListenBucketRequest_MetadataEntry {
  return { key: "", value: "" };
}

export const ListenBucketRequest_MetadataEntry = {
  encode(
    message: ListenBucketRequest_MetadataEntry,
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
  ): ListenBucketRequest_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListenBucketRequest_MetadataEntry();
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

  fromJSON(object: any): ListenBucketRequest_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: ListenBucketRequest_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<ListenBucketRequest_MetadataEntry>, I>
  >(object: I): ListenBucketRequest_MetadataEntry {
    const message = createBaseListenBucketRequest_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseCreateObjectRequest(): CreateObjectRequest {
  return { label: "", name: "", fields: [] };
}

export const CreateObjectRequest = {
  encode(
    message: CreateObjectRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    for (const v of message.fields) {
      ObjectField.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateObjectRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateObjectRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.fields.push(ObjectField.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateObjectRequest {
    return {
      label: isSet(object.label) ? String(object.label) : "",
      name: isSet(object.name) ? String(object.name) : "",
      fields: Array.isArray(object?.fields)
        ? object.fields.map((e: any) => ObjectField.fromJSON(e))
        : [],
    };
  },

  toJSON(message: CreateObjectRequest): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.name !== undefined && (obj.name = message.name);
    if (message.fields) {
      obj.fields = message.fields.map((e) =>
        e ? ObjectField.toJSON(e) : undefined
      );
    } else {
      obj.fields = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateObjectRequest>, I>>(
    object: I
  ): CreateObjectRequest {
    const message = createBaseCreateObjectRequest();
    message.label = object.label ?? "";
    message.name = object.name ?? "";
    message.fields =
      object.fields?.map((e) => ObjectField.fromPartial(e)) || [];
    return message;
  },
};

function createBaseReadObjectRequest(): ReadObjectRequest {
  return { did: "" };
}

export const ReadObjectRequest = {
  encode(
    message: ReadObjectRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadObjectRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadObjectRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ReadObjectRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
    };
  },

  toJSON(message: ReadObjectRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ReadObjectRequest>, I>>(
    object: I
  ): ReadObjectRequest {
    const message = createBaseReadObjectRequest();
    message.did = object.did ?? "";
    return message;
  },
};

function createBaseUpdateObjectRequest(): UpdateObjectRequest {
  return { did: "", fields: [] };
}

export const UpdateObjectRequest = {
  encode(
    message: UpdateObjectRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    for (const v of message.fields) {
      ObjectField.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateObjectRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateObjectRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          message.fields.push(ObjectField.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateObjectRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      fields: Array.isArray(object?.fields)
        ? object.fields.map((e: any) => ObjectField.fromJSON(e))
        : [],
    };
  },

  toJSON(message: UpdateObjectRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    if (message.fields) {
      obj.fields = message.fields.map((e) =>
        e ? ObjectField.toJSON(e) : undefined
      );
    } else {
      obj.fields = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateObjectRequest>, I>>(
    object: I
  ): UpdateObjectRequest {
    const message = createBaseUpdateObjectRequest();
    message.did = object.did ?? "";
    message.fields =
      object.fields?.map((e) => ObjectField.fromPartial(e)) || [];
    return message;
  },
};

function createBaseDeleteObjectRequest(): DeleteObjectRequest {
  return { did: "", metadata: {}, publicKey: "" };
}

export const DeleteObjectRequest = {
  encode(
    message: DeleteObjectRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      DeleteObjectRequest_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(18).fork()
      ).ldelim();
    });
    if (message.publicKey !== "") {
      writer.uint32(26).string(message.publicKey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteObjectRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteObjectRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          const entry2 = DeleteObjectRequest_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry2.value !== undefined) {
            message.metadata[entry2.key] = entry2.value;
          }
          break;
        case 3:
          message.publicKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteObjectRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
      publicKey: isSet(object.publicKey) ? String(object.publicKey) : "",
    };
  },

  toJSON(message: DeleteObjectRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    message.publicKey !== undefined && (obj.publicKey = message.publicKey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteObjectRequest>, I>>(
    object: I
  ): DeleteObjectRequest {
    const message = createBaseDeleteObjectRequest();
    message.did = object.did ?? "";
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    message.publicKey = object.publicKey ?? "";
    return message;
  },
};

function createBaseDeleteObjectRequest_MetadataEntry(): DeleteObjectRequest_MetadataEntry {
  return { key: "", value: "" };
}

export const DeleteObjectRequest_MetadataEntry = {
  encode(
    message: DeleteObjectRequest_MetadataEntry,
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
  ): DeleteObjectRequest_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteObjectRequest_MetadataEntry();
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

  fromJSON(object: any): DeleteObjectRequest_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: DeleteObjectRequest_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<DeleteObjectRequest_MetadataEntry>, I>
  >(object: I): DeleteObjectRequest_MetadataEntry {
    const message = createBaseDeleteObjectRequest_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseUploadBlobRequest(): UploadBlobRequest {
  return { label: "", path: "", bucketDid: "", size: 0, lastModified: 0 };
}

export const UploadBlobRequest = {
  encode(
    message: UploadBlobRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.path !== "") {
      writer.uint32(18).string(message.path);
    }
    if (message.bucketDid !== "") {
      writer.uint32(26).string(message.bucketDid);
    }
    if (message.size !== 0) {
      writer.uint32(32).int64(message.size);
    }
    if (message.lastModified !== 0) {
      writer.uint32(40).int64(message.lastModified);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UploadBlobRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUploadBlobRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.path = reader.string();
          break;
        case 3:
          message.bucketDid = reader.string();
          break;
        case 4:
          message.size = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.lastModified = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UploadBlobRequest {
    return {
      label: isSet(object.label) ? String(object.label) : "",
      path: isSet(object.path) ? String(object.path) : "",
      bucketDid: isSet(object.bucketDid) ? String(object.bucketDid) : "",
      size: isSet(object.size) ? Number(object.size) : 0,
      lastModified: isSet(object.lastModified)
        ? Number(object.lastModified)
        : 0,
    };
  },

  toJSON(message: UploadBlobRequest): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.path !== undefined && (obj.path = message.path);
    message.bucketDid !== undefined && (obj.bucketDid = message.bucketDid);
    message.size !== undefined && (obj.size = Math.round(message.size));
    message.lastModified !== undefined &&
      (obj.lastModified = Math.round(message.lastModified));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UploadBlobRequest>, I>>(
    object: I
  ): UploadBlobRequest {
    const message = createBaseUploadBlobRequest();
    message.label = object.label ?? "";
    message.path = object.path ?? "";
    message.bucketDid = object.bucketDid ?? "";
    message.size = object.size ?? 0;
    message.lastModified = object.lastModified ?? 0;
    return message;
  },
};

function createBaseDownloadBlobRequest(): DownloadBlobRequest {
  return { did: "", outPath: "" };
}

export const DownloadBlobRequest = {
  encode(
    message: DownloadBlobRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    if (message.outPath !== "") {
      writer.uint32(18).string(message.outPath);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DownloadBlobRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDownloadBlobRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          message.outPath = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DownloadBlobRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      outPath: isSet(object.outPath) ? String(object.outPath) : "",
    };
  },

  toJSON(message: DownloadBlobRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    message.outPath !== undefined && (obj.outPath = message.outPath);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DownloadBlobRequest>, I>>(
    object: I
  ): DownloadBlobRequest {
    const message = createBaseDownloadBlobRequest();
    message.did = object.did ?? "";
    message.outPath = object.outPath ?? "";
    return message;
  },
};

function createBaseSyncBlobRequest(): SyncBlobRequest {
  return { did: "", destinationDid: "", path: "" };
}

export const SyncBlobRequest = {
  encode(
    message: SyncBlobRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    if (message.destinationDid !== "") {
      writer.uint32(18).string(message.destinationDid);
    }
    if (message.path !== "") {
      writer.uint32(26).string(message.path);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SyncBlobRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSyncBlobRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          message.destinationDid = reader.string();
          break;
        case 3:
          message.path = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SyncBlobRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      destinationDid: isSet(object.destinationDid)
        ? String(object.destinationDid)
        : "",
      path: isSet(object.path) ? String(object.path) : "",
    };
  },

  toJSON(message: SyncBlobRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    message.destinationDid !== undefined &&
      (obj.destinationDid = message.destinationDid);
    message.path !== undefined && (obj.path = message.path);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SyncBlobRequest>, I>>(
    object: I
  ): SyncBlobRequest {
    const message = createBaseSyncBlobRequest();
    message.did = object.did ?? "";
    message.destinationDid = object.destinationDid ?? "";
    message.path = object.path ?? "";
    return message;
  },
};

function createBaseDeleteBlobRequest(): DeleteBlobRequest {
  return { did: "", metadata: {}, publicKey: "" };
}

export const DeleteBlobRequest = {
  encode(
    message: DeleteBlobRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      DeleteBlobRequest_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(18).fork()
      ).ldelim();
    });
    if (message.publicKey !== "") {
      writer.uint32(26).string(message.publicKey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteBlobRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteBlobRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          const entry2 = DeleteBlobRequest_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry2.value !== undefined) {
            message.metadata[entry2.key] = entry2.value;
          }
          break;
        case 3:
          message.publicKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteBlobRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
      publicKey: isSet(object.publicKey) ? String(object.publicKey) : "",
    };
  },

  toJSON(message: DeleteBlobRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    message.publicKey !== undefined && (obj.publicKey = message.publicKey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteBlobRequest>, I>>(
    object: I
  ): DeleteBlobRequest {
    const message = createBaseDeleteBlobRequest();
    message.did = object.did ?? "";
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    message.publicKey = object.publicKey ?? "";
    return message;
  },
};

function createBaseDeleteBlobRequest_MetadataEntry(): DeleteBlobRequest_MetadataEntry {
  return { key: "", value: "" };
}

export const DeleteBlobRequest_MetadataEntry = {
  encode(
    message: DeleteBlobRequest_MetadataEntry,
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
  ): DeleteBlobRequest_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteBlobRequest_MetadataEntry();
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

  fromJSON(object: any): DeleteBlobRequest_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: DeleteBlobRequest_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteBlobRequest_MetadataEntry>, I>>(
    object: I
  ): DeleteBlobRequest_MetadataEntry {
    const message = createBaseDeleteBlobRequest_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseParseDidRequest(): ParseDidRequest {
  return { did: "", metadata: {} };
}

export const ParseDidRequest = {
  encode(
    message: ParseDidRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ParseDidRequest_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(18).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ParseDidRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParseDidRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          const entry2 = ParseDidRequest_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry2.value !== undefined) {
            message.metadata[entry2.key] = entry2.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ParseDidRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
    };
  },

  toJSON(message: ParseDidRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ParseDidRequest>, I>>(
    object: I
  ): ParseDidRequest {
    const message = createBaseParseDidRequest();
    message.did = object.did ?? "";
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseParseDidRequest_MetadataEntry(): ParseDidRequest_MetadataEntry {
  return { key: "", value: "" };
}

export const ParseDidRequest_MetadataEntry = {
  encode(
    message: ParseDidRequest_MetadataEntry,
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
  ): ParseDidRequest_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParseDidRequest_MetadataEntry();
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

  fromJSON(object: any): ParseDidRequest_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: ParseDidRequest_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ParseDidRequest_MetadataEntry>, I>>(
    object: I
  ): ParseDidRequest_MetadataEntry {
    const message = createBaseParseDidRequest_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseResolveDidRequest(): ResolveDidRequest {
  return { did: "", metadata: {} };
}

export const ResolveDidRequest = {
  encode(
    message: ResolveDidRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.did !== "") {
      writer.uint32(10).string(message.did);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ResolveDidRequest_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(18).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ResolveDidRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseResolveDidRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.did = reader.string();
          break;
        case 2:
          const entry2 = ResolveDidRequest_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry2.value !== undefined) {
            message.metadata[entry2.key] = entry2.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ResolveDidRequest {
    return {
      did: isSet(object.did) ? String(object.did) : "",
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
    };
  },

  toJSON(message: ResolveDidRequest): unknown {
    const obj: any = {};
    message.did !== undefined && (obj.did = message.did);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ResolveDidRequest>, I>>(
    object: I
  ): ResolveDidRequest {
    const message = createBaseResolveDidRequest();
    message.did = object.did ?? "";
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseResolveDidRequest_MetadataEntry(): ResolveDidRequest_MetadataEntry {
  return { key: "", value: "" };
}

export const ResolveDidRequest_MetadataEntry = {
  encode(
    message: ResolveDidRequest_MetadataEntry,
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
  ): ResolveDidRequest_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseResolveDidRequest_MetadataEntry();
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

  fromJSON(object: any): ResolveDidRequest_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: ResolveDidRequest_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ResolveDidRequest_MetadataEntry>, I>>(
    object: I
  ): ResolveDidRequest_MetadataEntry {
    const message = createBaseResolveDidRequest_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
