/* eslint-disable */
import Long from "long";
import {
  makeGenericClientConstructor,
  ChannelCredentials,
  ChannelOptions,
  UntypedServiceImplementation,
  handleUnaryCall,
  handleServerStreamingCall,
  Client,
  ClientUnaryCall,
  Metadata,
  CallOptions,
  ClientReadableStream,
  ServiceError,
} from "@grpc/grpc-js";
import _m0 from "protobufjs/minimal";
import {
  AccessNameRequest,
  RegisterNameRequest,
  UpdateNameRequest,
  AccessServiceRequest,
  RegisterServiceRequest,
  UpdateServiceRequest,
  CreateChannelRequest,
  ReadChannelRequest,
  UpdateChannelRequest,
  DeleteChannelRequest,
  ListenChannelRequest,
  CreateBucketRequest,
  ReadBucketRequest,
  UpdateBucketRequest,
  DeleteBucketRequest,
  ListenBucketRequest,
  CreateObjectRequest,
  ReadObjectRequest,
  UpdateObjectRequest,
  DeleteObjectRequest,
  UploadBlobRequest,
  DownloadBlobRequest,
  SyncBlobRequest,
  DeleteBlobRequest,
  ParseDidRequest,
  ResolveDidRequest,
} from "../../../node/highway/v1/request";
import {
  AccessNameResponse,
  RegisterNameResponse,
  UpdateNameResponse,
  AccessServiceResponse,
  RegisterServiceResponse,
  UpdateServiceResponse,
  CreateChannelResponse,
  ReadChannelResponse,
  UpdateChannelResponse,
  DeleteChannelResponse,
  ListenChannelResponse,
  CreateBucketResponse,
  ReadBucketResponse,
  UpdateBucketResponse,
  DeleteBucketResponse,
  ListenBucketResponse,
  CreateObjectResponse,
  ReadObjectResponse,
  UpdateObjectResponse,
  DeleteObjectResponse,
  UploadBlobResponse,
  DownloadBlobResponse,
  SyncBlobResponse,
  DeleteBlobResponse,
  ParseDidResponse,
  ResolveDidResponse,
} from "../../../node/highway/v1/response";

export const protobufPackage = "node.highway.v1";

/** HighwayService is a RPC service for interfacing over the Highway node. */
export const HighwayServiceService = {
  /**
   * AccessName returns details and publicly available information about the Peer given calling node
   * has permission to access. i.e "prad.snr" -> "firstname online profilePic city"
   */
  accessName: {
    path: "/node.highway.v1.HighwayService/AccessName",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: AccessNameRequest) =>
      Buffer.from(AccessNameRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => AccessNameRequest.decode(value),
    responseSerialize: (value: AccessNameResponse) =>
      Buffer.from(AccessNameResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => AccessNameResponse.decode(value),
  },
  /**
   * RegisterName registers a new ".snr" name for the calling node. It is only allowed to be called
   * once per node.
   */
  registerName: {
    path: "/node.highway.v1.HighwayService/RegisterName",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: RegisterNameRequest) =>
      Buffer.from(RegisterNameRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => RegisterNameRequest.decode(value),
    responseSerialize: (value: RegisterNameResponse) =>
      Buffer.from(RegisterNameResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => RegisterNameResponse.decode(value),
  },
  /** UpdateName updates the public information of the calling node. */
  updateName: {
    path: "/node.highway.v1.HighwayService/UpdateName",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: UpdateNameRequest) =>
      Buffer.from(UpdateNameRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => UpdateNameRequest.decode(value),
    responseSerialize: (value: UpdateNameResponse) =>
      Buffer.from(UpdateNameResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => UpdateNameResponse.decode(value),
  },
  /**
   * AccessService creates a new signing key for the calling node in order to be authorized to
   * access the service. It is only allowed to be called once per node.
   */
  accessService: {
    path: "/node.highway.v1.HighwayService/AccessService",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: AccessServiceRequest) =>
      Buffer.from(AccessServiceRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => AccessServiceRequest.decode(value),
    responseSerialize: (value: AccessServiceResponse) =>
      Buffer.from(AccessServiceResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => AccessServiceResponse.decode(value),
  },
  /**
   * RegisterService registers a new service for the calling node. The calling node must have
   * already been enabled for development.
   */
  registerService: {
    path: "/node.highway.v1.HighwayService/RegisterService",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: RegisterServiceRequest) =>
      Buffer.from(RegisterServiceRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => RegisterServiceRequest.decode(value),
    responseSerialize: (value: RegisterServiceResponse) =>
      Buffer.from(RegisterServiceResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) =>
      RegisterServiceResponse.decode(value),
  },
  /** UpdateService updates the details and public configuration of the calling node's service. */
  updateService: {
    path: "/node.highway.v1.HighwayService/UpdateService",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: UpdateServiceRequest) =>
      Buffer.from(UpdateServiceRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => UpdateServiceRequest.decode(value),
    responseSerialize: (value: UpdateServiceResponse) =>
      Buffer.from(UpdateServiceResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => UpdateServiceResponse.decode(value),
  },
  /**
   * CreateChannel creates a new Publish/Subscribe topic channel for the given service.
   * The calling node must have already registered a service for the channel.
   */
  createChannel: {
    path: "/node.highway.v1.HighwayService/CreateChannel",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: CreateChannelRequest) =>
      Buffer.from(CreateChannelRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => CreateChannelRequest.decode(value),
    responseSerialize: (value: CreateChannelResponse) =>
      Buffer.from(CreateChannelResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => CreateChannelResponse.decode(value),
  },
  /**
   * ReadChannel lists all peers subscribed to the given channel, and additional details about
   * the channels configuration.
   */
  readChannel: {
    path: "/node.highway.v1.HighwayService/ReadChannel",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: ReadChannelRequest) =>
      Buffer.from(ReadChannelRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => ReadChannelRequest.decode(value),
    responseSerialize: (value: ReadChannelResponse) =>
      Buffer.from(ReadChannelResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => ReadChannelResponse.decode(value),
  },
  /** UpdateChannel updates the configuration of the given channel. */
  updateChannel: {
    path: "/node.highway.v1.HighwayService/UpdateChannel",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: UpdateChannelRequest) =>
      Buffer.from(UpdateChannelRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => UpdateChannelRequest.decode(value),
    responseSerialize: (value: UpdateChannelResponse) =>
      Buffer.from(UpdateChannelResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => UpdateChannelResponse.decode(value),
  },
  /** DeleteChannel deletes the given channel if the calling node is the owner of the channel. */
  deleteChannel: {
    path: "/node.highway.v1.HighwayService/DeleteChannel",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: DeleteChannelRequest) =>
      Buffer.from(DeleteChannelRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => DeleteChannelRequest.decode(value),
    responseSerialize: (value: DeleteChannelResponse) =>
      Buffer.from(DeleteChannelResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => DeleteChannelResponse.decode(value),
  },
  /**
   * ListenChannel subscribes the calling node to the given channel and returns all publish events
   * as a stream.
   */
  listenChannel: {
    path: "/node.highway.v1.HighwayService/ListenChannel",
    requestStream: false,
    responseStream: true,
    requestSerialize: (value: ListenChannelRequest) =>
      Buffer.from(ListenChannelRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => ListenChannelRequest.decode(value),
    responseSerialize: (value: ListenChannelResponse) =>
      Buffer.from(ListenChannelResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => ListenChannelResponse.decode(value),
  },
  /** CreateBucket creates a new bucket for the calling nodes service. */
  createBucket: {
    path: "/node.highway.v1.HighwayService/CreateBucket",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: CreateBucketRequest) =>
      Buffer.from(CreateBucketRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => CreateBucketRequest.decode(value),
    responseSerialize: (value: CreateBucketResponse) =>
      Buffer.from(CreateBucketResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => CreateBucketResponse.decode(value),
  },
  /**
   * ReadBucket lists all the blobs in the given bucket. The calling node must have access to the
   * bucket.
   */
  readBucket: {
    path: "/node.highway.v1.HighwayService/ReadBucket",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: ReadBucketRequest) =>
      Buffer.from(ReadBucketRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => ReadBucketRequest.decode(value),
    responseSerialize: (value: ReadBucketResponse) =>
      Buffer.from(ReadBucketResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => ReadBucketResponse.decode(value),
  },
  /**
   * UpdateBucket updates the configuration of the given bucket. The calling node must have access
   * to the bucket.
   */
  updateBucket: {
    path: "/node.highway.v1.HighwayService/UpdateBucket",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: UpdateBucketRequest) =>
      Buffer.from(UpdateBucketRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => UpdateBucketRequest.decode(value),
    responseSerialize: (value: UpdateBucketResponse) =>
      Buffer.from(UpdateBucketResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => UpdateBucketResponse.decode(value),
  },
  /** DeleteBucket deletes the given bucket if the calling node is the owner of the bucket. */
  deleteBucket: {
    path: "/node.highway.v1.HighwayService/DeleteBucket",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: DeleteBucketRequest) =>
      Buffer.from(DeleteBucketRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => DeleteBucketRequest.decode(value),
    responseSerialize: (value: DeleteBucketResponse) =>
      Buffer.from(DeleteBucketResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => DeleteBucketResponse.decode(value),
  },
  /**
   * ListenBucket subscribes the calling node to the given bucket and returns all publish events
   * as a stream.
   */
  listenBucket: {
    path: "/node.highway.v1.HighwayService/ListenBucket",
    requestStream: false,
    responseStream: true,
    requestSerialize: (value: ListenBucketRequest) =>
      Buffer.from(ListenBucketRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => ListenBucketRequest.decode(value),
    responseSerialize: (value: ListenBucketResponse) =>
      Buffer.from(ListenBucketResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => ListenBucketResponse.decode(value),
  },
  /**
   * CreateObject defines a new object to be utilized by the calling node's service. The object will
   * be placed in the Highway Service Graph and can be used in channels and other modules.
   */
  createObject: {
    path: "/node.highway.v1.HighwayService/CreateObject",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: CreateObjectRequest) =>
      Buffer.from(CreateObjectRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => CreateObjectRequest.decode(value),
    responseSerialize: (value: CreateObjectResponse) =>
      Buffer.from(CreateObjectResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => CreateObjectResponse.decode(value),
  },
  /** ReadObject returns the details of the given object provided its DID or Label. */
  readObject: {
    path: "/node.highway.v1.HighwayService/ReadObject",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: ReadObjectRequest) =>
      Buffer.from(ReadObjectRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => ReadObjectRequest.decode(value),
    responseSerialize: (value: ReadObjectResponse) =>
      Buffer.from(ReadObjectResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => ReadObjectResponse.decode(value),
  },
  /** UpdateObject modifies the property fields of the given object. */
  updateObject: {
    path: "/node.highway.v1.HighwayService/UpdateObject",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: UpdateObjectRequest) =>
      Buffer.from(UpdateObjectRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => UpdateObjectRequest.decode(value),
    responseSerialize: (value: UpdateObjectResponse) =>
      Buffer.from(UpdateObjectResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => UpdateObjectResponse.decode(value),
  },
  /** DeleteObject deletes the given object if the calling node is the owner of the object. */
  deleteObject: {
    path: "/node.highway.v1.HighwayService/DeleteObject",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: DeleteObjectRequest) =>
      Buffer.from(DeleteObjectRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => DeleteObjectRequest.decode(value),
    responseSerialize: (value: DeleteObjectResponse) =>
      Buffer.from(DeleteObjectResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => DeleteObjectResponse.decode(value),
  },
  /** UploadBlob uploads a file or buffer to the calling node's service IPFS storage. */
  uploadBlob: {
    path: "/node.highway.v1.HighwayService/UploadBlob",
    requestStream: false,
    responseStream: true,
    requestSerialize: (value: UploadBlobRequest) =>
      Buffer.from(UploadBlobRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => UploadBlobRequest.decode(value),
    responseSerialize: (value: UploadBlobResponse) =>
      Buffer.from(UploadBlobResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => UploadBlobResponse.decode(value),
  },
  /** DownloadBlob downloads a file or buffer from the calling node's service IPFS storage. */
  downloadBlob: {
    path: "/node.highway.v1.HighwayService/DownloadBlob",
    requestStream: false,
    responseStream: true,
    requestSerialize: (value: DownloadBlobRequest) =>
      Buffer.from(DownloadBlobRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => DownloadBlobRequest.decode(value),
    responseSerialize: (value: DownloadBlobResponse) =>
      Buffer.from(DownloadBlobResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => DownloadBlobResponse.decode(value),
  },
  /** SyncBlob synchronizes a local file from the calling node to the given service's IPFS storage. */
  syncBlob: {
    path: "/node.highway.v1.HighwayService/SyncBlob",
    requestStream: false,
    responseStream: true,
    requestSerialize: (value: SyncBlobRequest) =>
      Buffer.from(SyncBlobRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => SyncBlobRequest.decode(value),
    responseSerialize: (value: SyncBlobResponse) =>
      Buffer.from(SyncBlobResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => SyncBlobResponse.decode(value),
  },
  /** DeleteBlob deletes the given file from the calling node's service IPFS storage. */
  deleteBlob: {
    path: "/node.highway.v1.HighwayService/DeleteBlob",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: DeleteBlobRequest) =>
      Buffer.from(DeleteBlobRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => DeleteBlobRequest.decode(value),
    responseSerialize: (value: DeleteBlobResponse) =>
      Buffer.from(DeleteBlobResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => DeleteBlobResponse.decode(value),
  },
  /** ParseDid parses a potential DID string into a DID object. */
  parseDid: {
    path: "/node.highway.v1.HighwayService/ParseDid",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: ParseDidRequest) =>
      Buffer.from(ParseDidRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => ParseDidRequest.decode(value),
    responseSerialize: (value: ParseDidResponse) =>
      Buffer.from(ParseDidResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => ParseDidResponse.decode(value),
  },
  /**
   * ResolveDid resolves a DID to its DID document if the DID is valid and the calling node has
   * access to the DID Document.
   */
  resolveDid: {
    path: "/node.highway.v1.HighwayService/ResolveDid",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: ResolveDidRequest) =>
      Buffer.from(ResolveDidRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => ResolveDidRequest.decode(value),
    responseSerialize: (value: ResolveDidResponse) =>
      Buffer.from(ResolveDidResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => ResolveDidResponse.decode(value),
  },
} as const;

export interface HighwayServiceServer extends UntypedServiceImplementation {
  /**
   * AccessName returns details and publicly available information about the Peer given calling node
   * has permission to access. i.e "prad.snr" -> "firstname online profilePic city"
   */
  accessName: handleUnaryCall<AccessNameRequest, AccessNameResponse>;
  /**
   * RegisterName registers a new ".snr" name for the calling node. It is only allowed to be called
   * once per node.
   */
  registerName: handleUnaryCall<RegisterNameRequest, RegisterNameResponse>;
  /** UpdateName updates the public information of the calling node. */
  updateName: handleUnaryCall<UpdateNameRequest, UpdateNameResponse>;
  /**
   * AccessService creates a new signing key for the calling node in order to be authorized to
   * access the service. It is only allowed to be called once per node.
   */
  accessService: handleUnaryCall<AccessServiceRequest, AccessServiceResponse>;
  /**
   * RegisterService registers a new service for the calling node. The calling node must have
   * already been enabled for development.
   */
  registerService: handleUnaryCall<
    RegisterServiceRequest,
    RegisterServiceResponse
  >;
  /** UpdateService updates the details and public configuration of the calling node's service. */
  updateService: handleUnaryCall<UpdateServiceRequest, UpdateServiceResponse>;
  /**
   * CreateChannel creates a new Publish/Subscribe topic channel for the given service.
   * The calling node must have already registered a service for the channel.
   */
  createChannel: handleUnaryCall<CreateChannelRequest, CreateChannelResponse>;
  /**
   * ReadChannel lists all peers subscribed to the given channel, and additional details about
   * the channels configuration.
   */
  readChannel: handleUnaryCall<ReadChannelRequest, ReadChannelResponse>;
  /** UpdateChannel updates the configuration of the given channel. */
  updateChannel: handleUnaryCall<UpdateChannelRequest, UpdateChannelResponse>;
  /** DeleteChannel deletes the given channel if the calling node is the owner of the channel. */
  deleteChannel: handleUnaryCall<DeleteChannelRequest, DeleteChannelResponse>;
  /**
   * ListenChannel subscribes the calling node to the given channel and returns all publish events
   * as a stream.
   */
  listenChannel: handleServerStreamingCall<
    ListenChannelRequest,
    ListenChannelResponse
  >;
  /** CreateBucket creates a new bucket for the calling nodes service. */
  createBucket: handleUnaryCall<CreateBucketRequest, CreateBucketResponse>;
  /**
   * ReadBucket lists all the blobs in the given bucket. The calling node must have access to the
   * bucket.
   */
  readBucket: handleUnaryCall<ReadBucketRequest, ReadBucketResponse>;
  /**
   * UpdateBucket updates the configuration of the given bucket. The calling node must have access
   * to the bucket.
   */
  updateBucket: handleUnaryCall<UpdateBucketRequest, UpdateBucketResponse>;
  /** DeleteBucket deletes the given bucket if the calling node is the owner of the bucket. */
  deleteBucket: handleUnaryCall<DeleteBucketRequest, DeleteBucketResponse>;
  /**
   * ListenBucket subscribes the calling node to the given bucket and returns all publish events
   * as a stream.
   */
  listenBucket: handleServerStreamingCall<
    ListenBucketRequest,
    ListenBucketResponse
  >;
  /**
   * CreateObject defines a new object to be utilized by the calling node's service. The object will
   * be placed in the Highway Service Graph and can be used in channels and other modules.
   */
  createObject: handleUnaryCall<CreateObjectRequest, CreateObjectResponse>;
  /** ReadObject returns the details of the given object provided its DID or Label. */
  readObject: handleUnaryCall<ReadObjectRequest, ReadObjectResponse>;
  /** UpdateObject modifies the property fields of the given object. */
  updateObject: handleUnaryCall<UpdateObjectRequest, UpdateObjectResponse>;
  /** DeleteObject deletes the given object if the calling node is the owner of the object. */
  deleteObject: handleUnaryCall<DeleteObjectRequest, DeleteObjectResponse>;
  /** UploadBlob uploads a file or buffer to the calling node's service IPFS storage. */
  uploadBlob: handleServerStreamingCall<UploadBlobRequest, UploadBlobResponse>;
  /** DownloadBlob downloads a file or buffer from the calling node's service IPFS storage. */
  downloadBlob: handleServerStreamingCall<
    DownloadBlobRequest,
    DownloadBlobResponse
  >;
  /** SyncBlob synchronizes a local file from the calling node to the given service's IPFS storage. */
  syncBlob: handleServerStreamingCall<SyncBlobRequest, SyncBlobResponse>;
  /** DeleteBlob deletes the given file from the calling node's service IPFS storage. */
  deleteBlob: handleUnaryCall<DeleteBlobRequest, DeleteBlobResponse>;
  /** ParseDid parses a potential DID string into a DID object. */
  parseDid: handleUnaryCall<ParseDidRequest, ParseDidResponse>;
  /**
   * ResolveDid resolves a DID to its DID document if the DID is valid and the calling node has
   * access to the DID Document.
   */
  resolveDid: handleUnaryCall<ResolveDidRequest, ResolveDidResponse>;
}

export interface HighwayServiceClient extends Client {
  /**
   * AccessName returns details and publicly available information about the Peer given calling node
   * has permission to access. i.e "prad.snr" -> "firstname online profilePic city"
   */
  accessName(
    request: AccessNameRequest,
    callback: (error: ServiceError | null, response: AccessNameResponse) => void
  ): ClientUnaryCall;
  accessName(
    request: AccessNameRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: AccessNameResponse) => void
  ): ClientUnaryCall;
  accessName(
    request: AccessNameRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: AccessNameResponse) => void
  ): ClientUnaryCall;
  /**
   * RegisterName registers a new ".snr" name for the calling node. It is only allowed to be called
   * once per node.
   */
  registerName(
    request: RegisterNameRequest,
    callback: (
      error: ServiceError | null,
      response: RegisterNameResponse
    ) => void
  ): ClientUnaryCall;
  registerName(
    request: RegisterNameRequest,
    metadata: Metadata,
    callback: (
      error: ServiceError | null,
      response: RegisterNameResponse
    ) => void
  ): ClientUnaryCall;
  registerName(
    request: RegisterNameRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (
      error: ServiceError | null,
      response: RegisterNameResponse
    ) => void
  ): ClientUnaryCall;
  /** UpdateName updates the public information of the calling node. */
  updateName(
    request: UpdateNameRequest,
    callback: (error: ServiceError | null, response: UpdateNameResponse) => void
  ): ClientUnaryCall;
  updateName(
    request: UpdateNameRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: UpdateNameResponse) => void
  ): ClientUnaryCall;
  updateName(
    request: UpdateNameRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: UpdateNameResponse) => void
  ): ClientUnaryCall;
  /**
   * AccessService creates a new signing key for the calling node in order to be authorized to
   * access the service. It is only allowed to be called once per node.
   */
  accessService(
    request: AccessServiceRequest,
    callback: (
      error: ServiceError | null,
      response: AccessServiceResponse
    ) => void
  ): ClientUnaryCall;
  accessService(
    request: AccessServiceRequest,
    metadata: Metadata,
    callback: (
      error: ServiceError | null,
      response: AccessServiceResponse
    ) => void
  ): ClientUnaryCall;
  accessService(
    request: AccessServiceRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (
      error: ServiceError | null,
      response: AccessServiceResponse
    ) => void
  ): ClientUnaryCall;
  /**
   * RegisterService registers a new service for the calling node. The calling node must have
   * already been enabled for development.
   */
  registerService(
    request: RegisterServiceRequest,
    callback: (
      error: ServiceError | null,
      response: RegisterServiceResponse
    ) => void
  ): ClientUnaryCall;
  registerService(
    request: RegisterServiceRequest,
    metadata: Metadata,
    callback: (
      error: ServiceError | null,
      response: RegisterServiceResponse
    ) => void
  ): ClientUnaryCall;
  registerService(
    request: RegisterServiceRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (
      error: ServiceError | null,
      response: RegisterServiceResponse
    ) => void
  ): ClientUnaryCall;
  /** UpdateService updates the details and public configuration of the calling node's service. */
  updateService(
    request: UpdateServiceRequest,
    callback: (
      error: ServiceError | null,
      response: UpdateServiceResponse
    ) => void
  ): ClientUnaryCall;
  updateService(
    request: UpdateServiceRequest,
    metadata: Metadata,
    callback: (
      error: ServiceError | null,
      response: UpdateServiceResponse
    ) => void
  ): ClientUnaryCall;
  updateService(
    request: UpdateServiceRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (
      error: ServiceError | null,
      response: UpdateServiceResponse
    ) => void
  ): ClientUnaryCall;
  /**
   * CreateChannel creates a new Publish/Subscribe topic channel for the given service.
   * The calling node must have already registered a service for the channel.
   */
  createChannel(
    request: CreateChannelRequest,
    callback: (
      error: ServiceError | null,
      response: CreateChannelResponse
    ) => void
  ): ClientUnaryCall;
  createChannel(
    request: CreateChannelRequest,
    metadata: Metadata,
    callback: (
      error: ServiceError | null,
      response: CreateChannelResponse
    ) => void
  ): ClientUnaryCall;
  createChannel(
    request: CreateChannelRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (
      error: ServiceError | null,
      response: CreateChannelResponse
    ) => void
  ): ClientUnaryCall;
  /**
   * ReadChannel lists all peers subscribed to the given channel, and additional details about
   * the channels configuration.
   */
  readChannel(
    request: ReadChannelRequest,
    callback: (
      error: ServiceError | null,
      response: ReadChannelResponse
    ) => void
  ): ClientUnaryCall;
  readChannel(
    request: ReadChannelRequest,
    metadata: Metadata,
    callback: (
      error: ServiceError | null,
      response: ReadChannelResponse
    ) => void
  ): ClientUnaryCall;
  readChannel(
    request: ReadChannelRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (
      error: ServiceError | null,
      response: ReadChannelResponse
    ) => void
  ): ClientUnaryCall;
  /** UpdateChannel updates the configuration of the given channel. */
  updateChannel(
    request: UpdateChannelRequest,
    callback: (
      error: ServiceError | null,
      response: UpdateChannelResponse
    ) => void
  ): ClientUnaryCall;
  updateChannel(
    request: UpdateChannelRequest,
    metadata: Metadata,
    callback: (
      error: ServiceError | null,
      response: UpdateChannelResponse
    ) => void
  ): ClientUnaryCall;
  updateChannel(
    request: UpdateChannelRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (
      error: ServiceError | null,
      response: UpdateChannelResponse
    ) => void
  ): ClientUnaryCall;
  /** DeleteChannel deletes the given channel if the calling node is the owner of the channel. */
  deleteChannel(
    request: DeleteChannelRequest,
    callback: (
      error: ServiceError | null,
      response: DeleteChannelResponse
    ) => void
  ): ClientUnaryCall;
  deleteChannel(
    request: DeleteChannelRequest,
    metadata: Metadata,
    callback: (
      error: ServiceError | null,
      response: DeleteChannelResponse
    ) => void
  ): ClientUnaryCall;
  deleteChannel(
    request: DeleteChannelRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (
      error: ServiceError | null,
      response: DeleteChannelResponse
    ) => void
  ): ClientUnaryCall;
  /**
   * ListenChannel subscribes the calling node to the given channel and returns all publish events
   * as a stream.
   */
  listenChannel(
    request: ListenChannelRequest,
    options?: Partial<CallOptions>
  ): ClientReadableStream<ListenChannelResponse>;
  listenChannel(
    request: ListenChannelRequest,
    metadata?: Metadata,
    options?: Partial<CallOptions>
  ): ClientReadableStream<ListenChannelResponse>;
  /** CreateBucket creates a new bucket for the calling nodes service. */
  createBucket(
    request: CreateBucketRequest,
    callback: (
      error: ServiceError | null,
      response: CreateBucketResponse
    ) => void
  ): ClientUnaryCall;
  createBucket(
    request: CreateBucketRequest,
    metadata: Metadata,
    callback: (
      error: ServiceError | null,
      response: CreateBucketResponse
    ) => void
  ): ClientUnaryCall;
  createBucket(
    request: CreateBucketRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (
      error: ServiceError | null,
      response: CreateBucketResponse
    ) => void
  ): ClientUnaryCall;
  /**
   * ReadBucket lists all the blobs in the given bucket. The calling node must have access to the
   * bucket.
   */
  readBucket(
    request: ReadBucketRequest,
    callback: (error: ServiceError | null, response: ReadBucketResponse) => void
  ): ClientUnaryCall;
  readBucket(
    request: ReadBucketRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: ReadBucketResponse) => void
  ): ClientUnaryCall;
  readBucket(
    request: ReadBucketRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: ReadBucketResponse) => void
  ): ClientUnaryCall;
  /**
   * UpdateBucket updates the configuration of the given bucket. The calling node must have access
   * to the bucket.
   */
  updateBucket(
    request: UpdateBucketRequest,
    callback: (
      error: ServiceError | null,
      response: UpdateBucketResponse
    ) => void
  ): ClientUnaryCall;
  updateBucket(
    request: UpdateBucketRequest,
    metadata: Metadata,
    callback: (
      error: ServiceError | null,
      response: UpdateBucketResponse
    ) => void
  ): ClientUnaryCall;
  updateBucket(
    request: UpdateBucketRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (
      error: ServiceError | null,
      response: UpdateBucketResponse
    ) => void
  ): ClientUnaryCall;
  /** DeleteBucket deletes the given bucket if the calling node is the owner of the bucket. */
  deleteBucket(
    request: DeleteBucketRequest,
    callback: (
      error: ServiceError | null,
      response: DeleteBucketResponse
    ) => void
  ): ClientUnaryCall;
  deleteBucket(
    request: DeleteBucketRequest,
    metadata: Metadata,
    callback: (
      error: ServiceError | null,
      response: DeleteBucketResponse
    ) => void
  ): ClientUnaryCall;
  deleteBucket(
    request: DeleteBucketRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (
      error: ServiceError | null,
      response: DeleteBucketResponse
    ) => void
  ): ClientUnaryCall;
  /**
   * ListenBucket subscribes the calling node to the given bucket and returns all publish events
   * as a stream.
   */
  listenBucket(
    request: ListenBucketRequest,
    options?: Partial<CallOptions>
  ): ClientReadableStream<ListenBucketResponse>;
  listenBucket(
    request: ListenBucketRequest,
    metadata?: Metadata,
    options?: Partial<CallOptions>
  ): ClientReadableStream<ListenBucketResponse>;
  /**
   * CreateObject defines a new object to be utilized by the calling node's service. The object will
   * be placed in the Highway Service Graph and can be used in channels and other modules.
   */
  createObject(
    request: CreateObjectRequest,
    callback: (
      error: ServiceError | null,
      response: CreateObjectResponse
    ) => void
  ): ClientUnaryCall;
  createObject(
    request: CreateObjectRequest,
    metadata: Metadata,
    callback: (
      error: ServiceError | null,
      response: CreateObjectResponse
    ) => void
  ): ClientUnaryCall;
  createObject(
    request: CreateObjectRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (
      error: ServiceError | null,
      response: CreateObjectResponse
    ) => void
  ): ClientUnaryCall;
  /** ReadObject returns the details of the given object provided its DID or Label. */
  readObject(
    request: ReadObjectRequest,
    callback: (error: ServiceError | null, response: ReadObjectResponse) => void
  ): ClientUnaryCall;
  readObject(
    request: ReadObjectRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: ReadObjectResponse) => void
  ): ClientUnaryCall;
  readObject(
    request: ReadObjectRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: ReadObjectResponse) => void
  ): ClientUnaryCall;
  /** UpdateObject modifies the property fields of the given object. */
  updateObject(
    request: UpdateObjectRequest,
    callback: (
      error: ServiceError | null,
      response: UpdateObjectResponse
    ) => void
  ): ClientUnaryCall;
  updateObject(
    request: UpdateObjectRequest,
    metadata: Metadata,
    callback: (
      error: ServiceError | null,
      response: UpdateObjectResponse
    ) => void
  ): ClientUnaryCall;
  updateObject(
    request: UpdateObjectRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (
      error: ServiceError | null,
      response: UpdateObjectResponse
    ) => void
  ): ClientUnaryCall;
  /** DeleteObject deletes the given object if the calling node is the owner of the object. */
  deleteObject(
    request: DeleteObjectRequest,
    callback: (
      error: ServiceError | null,
      response: DeleteObjectResponse
    ) => void
  ): ClientUnaryCall;
  deleteObject(
    request: DeleteObjectRequest,
    metadata: Metadata,
    callback: (
      error: ServiceError | null,
      response: DeleteObjectResponse
    ) => void
  ): ClientUnaryCall;
  deleteObject(
    request: DeleteObjectRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (
      error: ServiceError | null,
      response: DeleteObjectResponse
    ) => void
  ): ClientUnaryCall;
  /** UploadBlob uploads a file or buffer to the calling node's service IPFS storage. */
  uploadBlob(
    request: UploadBlobRequest,
    options?: Partial<CallOptions>
  ): ClientReadableStream<UploadBlobResponse>;
  uploadBlob(
    request: UploadBlobRequest,
    metadata?: Metadata,
    options?: Partial<CallOptions>
  ): ClientReadableStream<UploadBlobResponse>;
  /** DownloadBlob downloads a file or buffer from the calling node's service IPFS storage. */
  downloadBlob(
    request: DownloadBlobRequest,
    options?: Partial<CallOptions>
  ): ClientReadableStream<DownloadBlobResponse>;
  downloadBlob(
    request: DownloadBlobRequest,
    metadata?: Metadata,
    options?: Partial<CallOptions>
  ): ClientReadableStream<DownloadBlobResponse>;
  /** SyncBlob synchronizes a local file from the calling node to the given service's IPFS storage. */
  syncBlob(
    request: SyncBlobRequest,
    options?: Partial<CallOptions>
  ): ClientReadableStream<SyncBlobResponse>;
  syncBlob(
    request: SyncBlobRequest,
    metadata?: Metadata,
    options?: Partial<CallOptions>
  ): ClientReadableStream<SyncBlobResponse>;
  /** DeleteBlob deletes the given file from the calling node's service IPFS storage. */
  deleteBlob(
    request: DeleteBlobRequest,
    callback: (error: ServiceError | null, response: DeleteBlobResponse) => void
  ): ClientUnaryCall;
  deleteBlob(
    request: DeleteBlobRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: DeleteBlobResponse) => void
  ): ClientUnaryCall;
  deleteBlob(
    request: DeleteBlobRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: DeleteBlobResponse) => void
  ): ClientUnaryCall;
  /** ParseDid parses a potential DID string into a DID object. */
  parseDid(
    request: ParseDidRequest,
    callback: (error: ServiceError | null, response: ParseDidResponse) => void
  ): ClientUnaryCall;
  parseDid(
    request: ParseDidRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: ParseDidResponse) => void
  ): ClientUnaryCall;
  parseDid(
    request: ParseDidRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: ParseDidResponse) => void
  ): ClientUnaryCall;
  /**
   * ResolveDid resolves a DID to its DID document if the DID is valid and the calling node has
   * access to the DID Document.
   */
  resolveDid(
    request: ResolveDidRequest,
    callback: (error: ServiceError | null, response: ResolveDidResponse) => void
  ): ClientUnaryCall;
  resolveDid(
    request: ResolveDidRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: ResolveDidResponse) => void
  ): ClientUnaryCall;
  resolveDid(
    request: ResolveDidRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: ResolveDidResponse) => void
  ): ClientUnaryCall;
}

export const HighwayServiceClient = makeGenericClientConstructor(
  HighwayServiceService,
  "node.highway.v1.HighwayService"
) as unknown as {
  new (
    address: string,
    credentials: ChannelCredentials,
    options?: Partial<ChannelOptions>
  ): HighwayServiceClient;
  service: typeof HighwayServiceService;
};

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}
