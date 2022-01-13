/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Peer } from "../../../common/v1/core";
import { DidDocument, Did } from "../../../did/v1/did";
import { ObjectDoc } from "../../../common/v1/object";

export const protobufPackage = "node.highway.v1";

/** / This file contains service for the Node RPC Server */

/** AccessNameResponse is a response to a request for a name */
export interface AccessNameResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** Data of the response */
  peer: Peer | undefined;
}

/** RegisterNameResponse is a request to register a name */
export interface RegisterNameResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  didDocument: DidDocument | undefined;
}

/** UpdateNameResponse is a response to a request to update a name */
export interface UpdateNameResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
  /** Data of the updated name record */
  metadata: { [key: string]: string };
}

export interface UpdateNameResponse_MetadataEntry {
  key: string;
  value: string;
}

/** AccessServiceResponse is a response to a request for a service */
export interface AccessServiceResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** Data of the response */
  metadata: { [key: string]: string };
}

export interface AccessServiceResponse_MetadataEntry {
  key: string;
  value: string;
}

/** RegisterServiceResponse is a request to register a name */
export interface RegisterServiceResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  didDocument: DidDocument | undefined;
}

/** UpdateServiceResponse is a response to a request to update a name */
export interface UpdateServiceResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
  /** Data of the updated name record */
  metadata: { [key: string]: string };
}

export interface UpdateServiceResponse_MetadataEntry {
  key: string;
  value: string;
}

/** CreateChannelResponse is a response to a request to create a channel */
export interface CreateChannelResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
}

/** ReadChannelResponse is a response to a request to read a channel */
export interface ReadChannelResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
  /** Subscribers of the channel */
  subscribers: Peer[];
  /** Owners of the channel */
  owners: Peer[];
}

/** UpdateChannelResponse is a response to a request to update a channel */
export interface UpdateChannelResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
  /** Subscribers of the channel */
  subscribers: Peer[];
  /** Owners of the channel */
  owners: Peer[];
}

/** DeleteChannelResponse is a response to a request to delete a channel */
export interface DeleteChannelResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
}

/** ListenChannelResponse is a response of the published data to the channel */
export interface ListenChannelResponse {
  /** Code of the response */
  code: number;
  /** DID of the response */
  did: Did | undefined;
  /** Additional information about the response */
  metadata: { [key: string]: string };
  /** Data of the response */
  message: Uint8Array;
}

export interface ListenChannelResponse_MetadataEntry {
  key: string;
  value: string;
}

/** CreateBucketResponse is a response to a request to create a bucket */
export interface CreateBucketResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
}

/** ReadBucketResponse is a response to a request to read a bucket */
export interface ReadBucketResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
  /** Subscribers of the bucket */
  objects: ObjectDoc[];
  /** Owners of the bucket */
  owners: Peer[];
}

/** UpdateBucketResponse is a response to a request to update a bucket */
export interface UpdateBucketResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
  /** Subscribers of the bucket */
  objects: ObjectDoc[];
  /** Owners of the bucket */
  owners: Peer[];
}

/** DeleteBucketResponse is a response to a request to delete a bucket */
export interface DeleteBucketResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
}

/** ListenBucketResponse is a response of the published data to the bucket */
export interface ListenBucketResponse {
  /** Code of the response */
  code: number;
  /** DID of the response */
  did: Did | undefined;
  /** Additional information about the response */
  metadata: { [key: string]: string };
  /** Stream of objects in the bucket */
  objects: ObjectDoc[];
}

export interface ListenBucketResponse_MetadataEntry {
  key: string;
  value: string;
}

/** CreateObjectResponse is a response to a request to create an object */
export interface CreateObjectResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
}

/** ReadObjectResponse is a response to a request to read an object */
export interface ReadObjectResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
  /** Data of the response */
  object: ObjectDoc | undefined;
}

/** UpdateObjectResponse is a response to a request to update an object */
export interface UpdateObjectResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
  /** Data of the response */
  object: ObjectDoc | undefined;
  /** Metadata is additional metadata of the response */
  metadata: { [key: string]: string };
}

export interface UpdateObjectResponse_MetadataEntry {
  key: string;
  value: string;
}

/** DeleteObjectResponse is a response to a request to delete an object */
export interface DeleteObjectResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
}

/** UploadBlobResponse is a response to a request to upload a blob */
export interface UploadBlobResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
  /** Pinned is true if the blob is pinned to IPFS */
  pinned: boolean;
}

/** DownloadBlobResponse is a response to a request to download a blob */
export interface DownloadBlobResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
  /** Path of downloaded blob */
  path: string;
}

/** SyncBlobResponse is a response to a request to sync a blob */
export interface SyncBlobResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
}

/** DeleteBlobResponse is a response to a request to delete a blob */
export interface DeleteBlobResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
}

/** ParseDidResponse is a response to a request to parse a DID */
export interface ParseDidResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  did: Did | undefined;
}

/** ResolveDidResponse is a response to a request to resolve a DID */
export interface ResolveDidResponse {
  /** Code of the response */
  code: number;
  /** Message of the response */
  message: string;
  /** DID of the response */
  didDocument: DidDocument | undefined;
}

function createBaseAccessNameResponse(): AccessNameResponse {
  return { code: 0, message: "", peer: undefined };
}

export const AccessNameResponse = {
  encode(
    message: AccessNameResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.peer !== undefined) {
      Peer.encode(message.peer, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AccessNameResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAccessNameResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.peer = Peer.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AccessNameResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      peer: isSet(object.peer) ? Peer.fromJSON(object.peer) : undefined,
    };
  },

  toJSON(message: AccessNameResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.peer !== undefined &&
      (obj.peer = message.peer ? Peer.toJSON(message.peer) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AccessNameResponse>, I>>(
    object: I
  ): AccessNameResponse {
    const message = createBaseAccessNameResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.peer =
      object.peer !== undefined && object.peer !== null
        ? Peer.fromPartial(object.peer)
        : undefined;
    return message;
  },
};

function createBaseRegisterNameResponse(): RegisterNameResponse {
  return { code: 0, message: "", didDocument: undefined };
}

export const RegisterNameResponse = {
  encode(
    message: RegisterNameResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.didDocument !== undefined) {
      DidDocument.encode(
        message.didDocument,
        writer.uint32(26).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): RegisterNameResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterNameResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.didDocument = DidDocument.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegisterNameResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      didDocument: isSet(object.didDocument)
        ? DidDocument.fromJSON(object.didDocument)
        : undefined,
    };
  },

  toJSON(message: RegisterNameResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.didDocument !== undefined &&
      (obj.didDocument = message.didDocument
        ? DidDocument.toJSON(message.didDocument)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterNameResponse>, I>>(
    object: I
  ): RegisterNameResponse {
    const message = createBaseRegisterNameResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.didDocument =
      object.didDocument !== undefined && object.didDocument !== null
        ? DidDocument.fromPartial(object.didDocument)
        : undefined;
    return message;
  },
};

function createBaseUpdateNameResponse(): UpdateNameResponse {
  return { code: 0, message: "", did: undefined, metadata: {} };
}

export const UpdateNameResponse = {
  encode(
    message: UpdateNameResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      UpdateNameResponse_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateNameResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateNameResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        case 4:
          const entry4 = UpdateNameResponse_MetadataEntry.decode(
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

  fromJSON(object: any): UpdateNameResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
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

  toJSON(message: UpdateNameResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateNameResponse>, I>>(
    object: I
  ): UpdateNameResponse {
    const message = createBaseUpdateNameResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
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

function createBaseUpdateNameResponse_MetadataEntry(): UpdateNameResponse_MetadataEntry {
  return { key: "", value: "" };
}

export const UpdateNameResponse_MetadataEntry = {
  encode(
    message: UpdateNameResponse_MetadataEntry,
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
  ): UpdateNameResponse_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateNameResponse_MetadataEntry();
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

  fromJSON(object: any): UpdateNameResponse_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: UpdateNameResponse_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<UpdateNameResponse_MetadataEntry>, I>
  >(object: I): UpdateNameResponse_MetadataEntry {
    const message = createBaseUpdateNameResponse_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseAccessServiceResponse(): AccessServiceResponse {
  return { code: 0, message: "", metadata: {} };
}

export const AccessServiceResponse = {
  encode(
    message: AccessServiceResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      AccessServiceResponse_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(26).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): AccessServiceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAccessServiceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          const entry3 = AccessServiceResponse_MetadataEntry.decode(
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

  fromJSON(object: any): AccessServiceResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
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

  toJSON(message: AccessServiceResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AccessServiceResponse>, I>>(
    object: I
  ): AccessServiceResponse {
    const message = createBaseAccessServiceResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
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

function createBaseAccessServiceResponse_MetadataEntry(): AccessServiceResponse_MetadataEntry {
  return { key: "", value: "" };
}

export const AccessServiceResponse_MetadataEntry = {
  encode(
    message: AccessServiceResponse_MetadataEntry,
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
  ): AccessServiceResponse_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAccessServiceResponse_MetadataEntry();
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

  fromJSON(object: any): AccessServiceResponse_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: AccessServiceResponse_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<AccessServiceResponse_MetadataEntry>, I>
  >(object: I): AccessServiceResponse_MetadataEntry {
    const message = createBaseAccessServiceResponse_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseRegisterServiceResponse(): RegisterServiceResponse {
  return { code: 0, message: "", didDocument: undefined };
}

export const RegisterServiceResponse = {
  encode(
    message: RegisterServiceResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.didDocument !== undefined) {
      DidDocument.encode(
        message.didDocument,
        writer.uint32(26).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): RegisterServiceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterServiceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.didDocument = DidDocument.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegisterServiceResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      didDocument: isSet(object.didDocument)
        ? DidDocument.fromJSON(object.didDocument)
        : undefined,
    };
  },

  toJSON(message: RegisterServiceResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.didDocument !== undefined &&
      (obj.didDocument = message.didDocument
        ? DidDocument.toJSON(message.didDocument)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterServiceResponse>, I>>(
    object: I
  ): RegisterServiceResponse {
    const message = createBaseRegisterServiceResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.didDocument =
      object.didDocument !== undefined && object.didDocument !== null
        ? DidDocument.fromPartial(object.didDocument)
        : undefined;
    return message;
  },
};

function createBaseUpdateServiceResponse(): UpdateServiceResponse {
  return { code: 0, message: "", did: undefined, metadata: {} };
}

export const UpdateServiceResponse = {
  encode(
    message: UpdateServiceResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      UpdateServiceResponse_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): UpdateServiceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateServiceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        case 4:
          const entry4 = UpdateServiceResponse_MetadataEntry.decode(
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

  fromJSON(object: any): UpdateServiceResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
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

  toJSON(message: UpdateServiceResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateServiceResponse>, I>>(
    object: I
  ): UpdateServiceResponse {
    const message = createBaseUpdateServiceResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
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

function createBaseUpdateServiceResponse_MetadataEntry(): UpdateServiceResponse_MetadataEntry {
  return { key: "", value: "" };
}

export const UpdateServiceResponse_MetadataEntry = {
  encode(
    message: UpdateServiceResponse_MetadataEntry,
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
  ): UpdateServiceResponse_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateServiceResponse_MetadataEntry();
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

  fromJSON(object: any): UpdateServiceResponse_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: UpdateServiceResponse_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<UpdateServiceResponse_MetadataEntry>, I>
  >(object: I): UpdateServiceResponse_MetadataEntry {
    const message = createBaseUpdateServiceResponse_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseCreateChannelResponse(): CreateChannelResponse {
  return { code: 0, message: "", did: undefined };
}

export const CreateChannelResponse = {
  encode(
    message: CreateChannelResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): CreateChannelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateChannelResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateChannelResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
    };
  },

  toJSON(message: CreateChannelResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateChannelResponse>, I>>(
    object: I
  ): CreateChannelResponse {
    const message = createBaseCreateChannelResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    return message;
  },
};

function createBaseReadChannelResponse(): ReadChannelResponse {
  return { code: 0, message: "", did: undefined, subscribers: [], owners: [] };
}

export const ReadChannelResponse = {
  encode(
    message: ReadChannelResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.subscribers) {
      Peer.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.owners) {
      Peer.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadChannelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadChannelResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        case 4:
          message.subscribers.push(Peer.decode(reader, reader.uint32()));
          break;
        case 5:
          message.owners.push(Peer.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ReadChannelResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
      subscribers: Array.isArray(object?.subscribers)
        ? object.subscribers.map((e: any) => Peer.fromJSON(e))
        : [],
      owners: Array.isArray(object?.owners)
        ? object.owners.map((e: any) => Peer.fromJSON(e))
        : [],
    };
  },

  toJSON(message: ReadChannelResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    if (message.subscribers) {
      obj.subscribers = message.subscribers.map((e) =>
        e ? Peer.toJSON(e) : undefined
      );
    } else {
      obj.subscribers = [];
    }
    if (message.owners) {
      obj.owners = message.owners.map((e) => (e ? Peer.toJSON(e) : undefined));
    } else {
      obj.owners = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ReadChannelResponse>, I>>(
    object: I
  ): ReadChannelResponse {
    const message = createBaseReadChannelResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    message.subscribers =
      object.subscribers?.map((e) => Peer.fromPartial(e)) || [];
    message.owners = object.owners?.map((e) => Peer.fromPartial(e)) || [];
    return message;
  },
};

function createBaseUpdateChannelResponse(): UpdateChannelResponse {
  return { code: 0, message: "", did: undefined, subscribers: [], owners: [] };
}

export const UpdateChannelResponse = {
  encode(
    message: UpdateChannelResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.subscribers) {
      Peer.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.owners) {
      Peer.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): UpdateChannelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateChannelResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        case 4:
          message.subscribers.push(Peer.decode(reader, reader.uint32()));
          break;
        case 5:
          message.owners.push(Peer.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateChannelResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
      subscribers: Array.isArray(object?.subscribers)
        ? object.subscribers.map((e: any) => Peer.fromJSON(e))
        : [],
      owners: Array.isArray(object?.owners)
        ? object.owners.map((e: any) => Peer.fromJSON(e))
        : [],
    };
  },

  toJSON(message: UpdateChannelResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    if (message.subscribers) {
      obj.subscribers = message.subscribers.map((e) =>
        e ? Peer.toJSON(e) : undefined
      );
    } else {
      obj.subscribers = [];
    }
    if (message.owners) {
      obj.owners = message.owners.map((e) => (e ? Peer.toJSON(e) : undefined));
    } else {
      obj.owners = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateChannelResponse>, I>>(
    object: I
  ): UpdateChannelResponse {
    const message = createBaseUpdateChannelResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    message.subscribers =
      object.subscribers?.map((e) => Peer.fromPartial(e)) || [];
    message.owners = object.owners?.map((e) => Peer.fromPartial(e)) || [];
    return message;
  },
};

function createBaseDeleteChannelResponse(): DeleteChannelResponse {
  return { code: 0, message: "", did: undefined };
}

export const DeleteChannelResponse = {
  encode(
    message: DeleteChannelResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): DeleteChannelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteChannelResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteChannelResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
    };
  },

  toJSON(message: DeleteChannelResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteChannelResponse>, I>>(
    object: I
  ): DeleteChannelResponse {
    const message = createBaseDeleteChannelResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    return message;
  },
};

function createBaseListenChannelResponse(): ListenChannelResponse {
  return { code: 0, did: undefined, metadata: {}, message: new Uint8Array() };
}

export const ListenChannelResponse = {
  encode(
    message: ListenChannelResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(18).fork()).ldelim();
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ListenChannelResponse_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(26).fork()
      ).ldelim();
    });
    if (message.message.length !== 0) {
      writer.uint32(34).bytes(message.message);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ListenChannelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListenChannelResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.did = Did.decode(reader, reader.uint32());
          break;
        case 3:
          const entry3 = ListenChannelResponse_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry3.value !== undefined) {
            message.metadata[entry3.key] = entry3.value;
          }
          break;
        case 4:
          message.message = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListenChannelResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
      message: isSet(object.message)
        ? bytesFromBase64(object.message)
        : new Uint8Array(),
    };
  },

  toJSON(message: ListenChannelResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    message.message !== undefined &&
      (obj.message = base64FromBytes(
        message.message !== undefined ? message.message : new Uint8Array()
      ));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListenChannelResponse>, I>>(
    object: I
  ): ListenChannelResponse {
    const message = createBaseListenChannelResponse();
    message.code = object.code ?? 0;
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    message.message = object.message ?? new Uint8Array();
    return message;
  },
};

function createBaseListenChannelResponse_MetadataEntry(): ListenChannelResponse_MetadataEntry {
  return { key: "", value: "" };
}

export const ListenChannelResponse_MetadataEntry = {
  encode(
    message: ListenChannelResponse_MetadataEntry,
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
  ): ListenChannelResponse_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListenChannelResponse_MetadataEntry();
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

  fromJSON(object: any): ListenChannelResponse_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: ListenChannelResponse_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<ListenChannelResponse_MetadataEntry>, I>
  >(object: I): ListenChannelResponse_MetadataEntry {
    const message = createBaseListenChannelResponse_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseCreateBucketResponse(): CreateBucketResponse {
  return { code: 0, message: "", did: undefined };
}

export const CreateBucketResponse = {
  encode(
    message: CreateBucketResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): CreateBucketResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBucketResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateBucketResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
    };
  },

  toJSON(message: CreateBucketResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateBucketResponse>, I>>(
    object: I
  ): CreateBucketResponse {
    const message = createBaseCreateBucketResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    return message;
  },
};

function createBaseReadBucketResponse(): ReadBucketResponse {
  return { code: 0, message: "", did: undefined, objects: [], owners: [] };
}

export const ReadBucketResponse = {
  encode(
    message: ReadBucketResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.objects) {
      ObjectDoc.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.owners) {
      Peer.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadBucketResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadBucketResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        case 4:
          message.objects.push(ObjectDoc.decode(reader, reader.uint32()));
          break;
        case 5:
          message.owners.push(Peer.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ReadBucketResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
      objects: Array.isArray(object?.objects)
        ? object.objects.map((e: any) => ObjectDoc.fromJSON(e))
        : [],
      owners: Array.isArray(object?.owners)
        ? object.owners.map((e: any) => Peer.fromJSON(e))
        : [],
    };
  },

  toJSON(message: ReadBucketResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    if (message.objects) {
      obj.objects = message.objects.map((e) =>
        e ? ObjectDoc.toJSON(e) : undefined
      );
    } else {
      obj.objects = [];
    }
    if (message.owners) {
      obj.owners = message.owners.map((e) => (e ? Peer.toJSON(e) : undefined));
    } else {
      obj.owners = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ReadBucketResponse>, I>>(
    object: I
  ): ReadBucketResponse {
    const message = createBaseReadBucketResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    message.objects =
      object.objects?.map((e) => ObjectDoc.fromPartial(e)) || [];
    message.owners = object.owners?.map((e) => Peer.fromPartial(e)) || [];
    return message;
  },
};

function createBaseUpdateBucketResponse(): UpdateBucketResponse {
  return { code: 0, message: "", did: undefined, objects: [], owners: [] };
}

export const UpdateBucketResponse = {
  encode(
    message: UpdateBucketResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.objects) {
      ObjectDoc.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.owners) {
      Peer.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): UpdateBucketResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateBucketResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        case 4:
          message.objects.push(ObjectDoc.decode(reader, reader.uint32()));
          break;
        case 5:
          message.owners.push(Peer.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateBucketResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
      objects: Array.isArray(object?.objects)
        ? object.objects.map((e: any) => ObjectDoc.fromJSON(e))
        : [],
      owners: Array.isArray(object?.owners)
        ? object.owners.map((e: any) => Peer.fromJSON(e))
        : [],
    };
  },

  toJSON(message: UpdateBucketResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    if (message.objects) {
      obj.objects = message.objects.map((e) =>
        e ? ObjectDoc.toJSON(e) : undefined
      );
    } else {
      obj.objects = [];
    }
    if (message.owners) {
      obj.owners = message.owners.map((e) => (e ? Peer.toJSON(e) : undefined));
    } else {
      obj.owners = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateBucketResponse>, I>>(
    object: I
  ): UpdateBucketResponse {
    const message = createBaseUpdateBucketResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    message.objects =
      object.objects?.map((e) => ObjectDoc.fromPartial(e)) || [];
    message.owners = object.owners?.map((e) => Peer.fromPartial(e)) || [];
    return message;
  },
};

function createBaseDeleteBucketResponse(): DeleteBucketResponse {
  return { code: 0, message: "", did: undefined };
}

export const DeleteBucketResponse = {
  encode(
    message: DeleteBucketResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): DeleteBucketResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteBucketResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteBucketResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
    };
  },

  toJSON(message: DeleteBucketResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteBucketResponse>, I>>(
    object: I
  ): DeleteBucketResponse {
    const message = createBaseDeleteBucketResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    return message;
  },
};

function createBaseListenBucketResponse(): ListenBucketResponse {
  return { code: 0, did: undefined, metadata: {}, objects: [] };
}

export const ListenBucketResponse = {
  encode(
    message: ListenBucketResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(18).fork()).ldelim();
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ListenBucketResponse_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(26).fork()
      ).ldelim();
    });
    for (const v of message.objects) {
      ObjectDoc.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ListenBucketResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListenBucketResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.did = Did.decode(reader, reader.uint32());
          break;
        case 3:
          const entry3 = ListenBucketResponse_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry3.value !== undefined) {
            message.metadata[entry3.key] = entry3.value;
          }
          break;
        case 4:
          message.objects.push(ObjectDoc.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListenBucketResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
              acc[key] = String(value);
              return acc;
            },
            {}
          )
        : {},
      objects: Array.isArray(object?.objects)
        ? object.objects.map((e: any) => ObjectDoc.fromJSON(e))
        : [],
    };
  },

  toJSON(message: ListenBucketResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    if (message.objects) {
      obj.objects = message.objects.map((e) =>
        e ? ObjectDoc.toJSON(e) : undefined
      );
    } else {
      obj.objects = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListenBucketResponse>, I>>(
    object: I
  ): ListenBucketResponse {
    const message = createBaseListenBucketResponse();
    message.code = object.code ?? 0;
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{
      [key: string]: string;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    message.objects =
      object.objects?.map((e) => ObjectDoc.fromPartial(e)) || [];
    return message;
  },
};

function createBaseListenBucketResponse_MetadataEntry(): ListenBucketResponse_MetadataEntry {
  return { key: "", value: "" };
}

export const ListenBucketResponse_MetadataEntry = {
  encode(
    message: ListenBucketResponse_MetadataEntry,
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
  ): ListenBucketResponse_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListenBucketResponse_MetadataEntry();
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

  fromJSON(object: any): ListenBucketResponse_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: ListenBucketResponse_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<ListenBucketResponse_MetadataEntry>, I>
  >(object: I): ListenBucketResponse_MetadataEntry {
    const message = createBaseListenBucketResponse_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseCreateObjectResponse(): CreateObjectResponse {
  return { code: 0, message: "", did: undefined };
}

export const CreateObjectResponse = {
  encode(
    message: CreateObjectResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): CreateObjectResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateObjectResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateObjectResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
    };
  },

  toJSON(message: CreateObjectResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateObjectResponse>, I>>(
    object: I
  ): CreateObjectResponse {
    const message = createBaseCreateObjectResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    return message;
  },
};

function createBaseReadObjectResponse(): ReadObjectResponse {
  return { code: 0, message: "", did: undefined, object: undefined };
}

export const ReadObjectResponse = {
  encode(
    message: ReadObjectResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    if (message.object !== undefined) {
      ObjectDoc.encode(message.object, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadObjectResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadObjectResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        case 4:
          message.object = ObjectDoc.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ReadObjectResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
      object: isSet(object.object)
        ? ObjectDoc.fromJSON(object.object)
        : undefined,
    };
  },

  toJSON(message: ReadObjectResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    message.object !== undefined &&
      (obj.object = message.object
        ? ObjectDoc.toJSON(message.object)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ReadObjectResponse>, I>>(
    object: I
  ): ReadObjectResponse {
    const message = createBaseReadObjectResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    message.object =
      object.object !== undefined && object.object !== null
        ? ObjectDoc.fromPartial(object.object)
        : undefined;
    return message;
  },
};

function createBaseUpdateObjectResponse(): UpdateObjectResponse {
  return {
    code: 0,
    message: "",
    did: undefined,
    object: undefined,
    metadata: {},
  };
}

export const UpdateObjectResponse = {
  encode(
    message: UpdateObjectResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    if (message.object !== undefined) {
      ObjectDoc.encode(message.object, writer.uint32(34).fork()).ldelim();
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      UpdateObjectResponse_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(42).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): UpdateObjectResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateObjectResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        case 4:
          message.object = ObjectDoc.decode(reader, reader.uint32());
          break;
        case 5:
          const entry5 = UpdateObjectResponse_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry5.value !== undefined) {
            message.metadata[entry5.key] = entry5.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateObjectResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
      object: isSet(object.object)
        ? ObjectDoc.fromJSON(object.object)
        : undefined,
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

  toJSON(message: UpdateObjectResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    message.object !== undefined &&
      (obj.object = message.object
        ? ObjectDoc.toJSON(message.object)
        : undefined);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateObjectResponse>, I>>(
    object: I
  ): UpdateObjectResponse {
    const message = createBaseUpdateObjectResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    message.object =
      object.object !== undefined && object.object !== null
        ? ObjectDoc.fromPartial(object.object)
        : undefined;
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

function createBaseUpdateObjectResponse_MetadataEntry(): UpdateObjectResponse_MetadataEntry {
  return { key: "", value: "" };
}

export const UpdateObjectResponse_MetadataEntry = {
  encode(
    message: UpdateObjectResponse_MetadataEntry,
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
  ): UpdateObjectResponse_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateObjectResponse_MetadataEntry();
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

  fromJSON(object: any): UpdateObjectResponse_MetadataEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: UpdateObjectResponse_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<UpdateObjectResponse_MetadataEntry>, I>
  >(object: I): UpdateObjectResponse_MetadataEntry {
    const message = createBaseUpdateObjectResponse_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseDeleteObjectResponse(): DeleteObjectResponse {
  return { code: 0, message: "", did: undefined };
}

export const DeleteObjectResponse = {
  encode(
    message: DeleteObjectResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): DeleteObjectResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteObjectResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteObjectResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
    };
  },

  toJSON(message: DeleteObjectResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteObjectResponse>, I>>(
    object: I
  ): DeleteObjectResponse {
    const message = createBaseDeleteObjectResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    return message;
  },
};

function createBaseUploadBlobResponse(): UploadBlobResponse {
  return { code: 0, message: "", did: undefined, pinned: false };
}

export const UploadBlobResponse = {
  encode(
    message: UploadBlobResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    if (message.pinned === true) {
      writer.uint32(32).bool(message.pinned);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UploadBlobResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUploadBlobResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        case 4:
          message.pinned = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UploadBlobResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
      pinned: isSet(object.pinned) ? Boolean(object.pinned) : false,
    };
  },

  toJSON(message: UploadBlobResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    message.pinned !== undefined && (obj.pinned = message.pinned);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UploadBlobResponse>, I>>(
    object: I
  ): UploadBlobResponse {
    const message = createBaseUploadBlobResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    message.pinned = object.pinned ?? false;
    return message;
  },
};

function createBaseDownloadBlobResponse(): DownloadBlobResponse {
  return { code: 0, message: "", did: undefined, path: "" };
}

export const DownloadBlobResponse = {
  encode(
    message: DownloadBlobResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    if (message.path !== "") {
      writer.uint32(34).string(message.path);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): DownloadBlobResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDownloadBlobResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        case 4:
          message.path = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DownloadBlobResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
      path: isSet(object.path) ? String(object.path) : "",
    };
  },

  toJSON(message: DownloadBlobResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    message.path !== undefined && (obj.path = message.path);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DownloadBlobResponse>, I>>(
    object: I
  ): DownloadBlobResponse {
    const message = createBaseDownloadBlobResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    message.path = object.path ?? "";
    return message;
  },
};

function createBaseSyncBlobResponse(): SyncBlobResponse {
  return { code: 0, message: "", did: undefined };
}

export const SyncBlobResponse = {
  encode(
    message: SyncBlobResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SyncBlobResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSyncBlobResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SyncBlobResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
    };
  },

  toJSON(message: SyncBlobResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SyncBlobResponse>, I>>(
    object: I
  ): SyncBlobResponse {
    const message = createBaseSyncBlobResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    return message;
  },
};

function createBaseDeleteBlobResponse(): DeleteBlobResponse {
  return { code: 0, message: "", did: undefined };
}

export const DeleteBlobResponse = {
  encode(
    message: DeleteBlobResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteBlobResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteBlobResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteBlobResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
    };
  },

  toJSON(message: DeleteBlobResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteBlobResponse>, I>>(
    object: I
  ): DeleteBlobResponse {
    const message = createBaseDeleteBlobResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    return message;
  },
};

function createBaseParseDidResponse(): ParseDidResponse {
  return { code: 0, message: "", did: undefined };
}

export const ParseDidResponse = {
  encode(
    message: ParseDidResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ParseDidResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParseDidResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.did = Did.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ParseDidResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      did: isSet(object.did) ? Did.fromJSON(object.did) : undefined,
    };
  },

  toJSON(message: ParseDidResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ParseDidResponse>, I>>(
    object: I
  ): ParseDidResponse {
    const message = createBaseParseDidResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.did =
      object.did !== undefined && object.did !== null
        ? Did.fromPartial(object.did)
        : undefined;
    return message;
  },
};

function createBaseResolveDidResponse(): ResolveDidResponse {
  return { code: 0, message: "", didDocument: undefined };
}

export const ResolveDidResponse = {
  encode(
    message: ResolveDidResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.didDocument !== undefined) {
      DidDocument.encode(
        message.didDocument,
        writer.uint32(26).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ResolveDidResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseResolveDidResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.didDocument = DidDocument.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ResolveDidResponse {
    return {
      code: isSet(object.code) ? Number(object.code) : 0,
      message: isSet(object.message) ? String(object.message) : "",
      didDocument: isSet(object.didDocument)
        ? DidDocument.fromJSON(object.didDocument)
        : undefined,
    };
  },

  toJSON(message: ResolveDidResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = Math.round(message.code));
    message.message !== undefined && (obj.message = message.message);
    message.didDocument !== undefined &&
      (obj.didDocument = message.didDocument
        ? DidDocument.toJSON(message.didDocument)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ResolveDidResponse>, I>>(
    object: I
  ): ResolveDidResponse {
    const message = createBaseResolveDidResponse();
    message.code = object.code ?? 0;
    message.message = object.message ?? "";
    message.didDocument =
      object.didDocument !== undefined && object.didDocument !== null
        ? DidDocument.fromPartial(object.didDocument)
        : undefined;
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
