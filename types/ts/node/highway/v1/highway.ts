/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
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
import { map } from "rxjs/operators";
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

export const protobufPackage = "node.highway.v1";

/** / This file contains service for the Node RPC Server */

/** RPC Service with Equivalent Methods of a binded Node */
export interface HighwayService {
  /**
   * AccessName returns details and publicly available information about the Peer given calling node
   * has permission to access. i.e "prad.snr" -> "firstname online profilePic city"
   */
  AccessName(request: AccessNameRequest): Promise<AccessNameResponse>;
  /**
   * RegisterName registers a new ".snr" name for the calling node. It is only allowed to be called
   * once per node.
   */
  RegisterName(request: RegisterNameRequest): Promise<RegisterNameResponse>;
  /** UpdateName updates the public information of the calling node. */
  UpdateName(request: UpdateNameRequest): Promise<UpdateNameResponse>;
  /**
   * AccessService creates a new signing key for the calling node in order to be authorized to
   * access the service. It is only allowed to be called once per node.
   */
  AccessService(request: AccessServiceRequest): Promise<AccessServiceResponse>;
  /**
   * RegisterService registers a new service for the calling node. The calling node must have
   * already been enabled for development.
   */
  RegisterService(
    request: RegisterServiceRequest
  ): Promise<RegisterServiceResponse>;
  /** UpdateService updates the details and public configuration of the calling node's service. */
  UpdateService(request: UpdateServiceRequest): Promise<UpdateServiceResponse>;
  /**
   * CreateChannel creates a new Publish/Subscribe topic channel for the given service.
   * The calling node must have already registered a service for the channel.
   */
  CreateChannel(request: CreateChannelRequest): Promise<CreateChannelResponse>;
  /**
   * ReadChannel lists all peers subscribed to the given channel, and additional details about
   * the channels configuration.
   */
  ReadChannel(request: ReadChannelRequest): Promise<ReadChannelResponse>;
  /** UpdateChannel updates the configuration of the given channel. */
  UpdateChannel(request: UpdateChannelRequest): Promise<UpdateChannelResponse>;
  /** DeleteChannel deletes the given channel if the calling node is the owner of the channel. */
  DeleteChannel(request: DeleteChannelRequest): Promise<DeleteChannelResponse>;
  /**
   * ListenChannel subscribes the calling node to the given channel and returns all publish events
   * as a stream.
   */
  ListenChannel(
    request: ListenChannelRequest
  ): Observable<ListenChannelResponse>;
  /** CreateBucket creates a new bucket for the calling nodes service. */
  CreateBucket(request: CreateBucketRequest): Promise<CreateBucketResponse>;
  /**
   * ReadBucket lists all the blobs in the given bucket. The calling node must have access to the
   * bucket.
   */
  ReadBucket(request: ReadBucketRequest): Promise<ReadBucketResponse>;
  /**
   * UpdateBucket updates the configuration of the given bucket. The calling node must have access
   * to the bucket.
   */
  UpdateBucket(request: UpdateBucketRequest): Promise<UpdateBucketResponse>;
  /** DeleteBucket deletes the given bucket if the calling node is the owner of the bucket. */
  DeleteBucket(request: DeleteBucketRequest): Promise<DeleteBucketResponse>;
  /**
   * ListenBucket subscribes the calling node to the given bucket and returns all publish events
   * as a stream.
   */
  ListenBucket(request: ListenBucketRequest): Observable<ListenBucketResponse>;
  /**
   * CreateObject defines a new object to be utilized by the calling node's service. The object will
   * be placed in the Highway Service Graph and can be used in channels and other modules.
   */
  CreateObject(request: CreateObjectRequest): Promise<CreateObjectResponse>;
  /** ReadObject returns the details of the given object provided its DID or Label. */
  ReadObject(request: ReadObjectRequest): Promise<ReadObjectResponse>;
  /** UpdateObject modifies the property fields of the given object. */
  UpdateObject(request: UpdateObjectRequest): Promise<UpdateObjectResponse>;
  /** DeleteObject deletes the given object if the calling node is the owner of the object. */
  DeleteObject(request: DeleteObjectRequest): Promise<DeleteObjectResponse>;
  /** UploadBlob uploads a file or buffer to the calling node's service IPFS storage. */
  UploadBlob(request: UploadBlobRequest): Observable<UploadBlobResponse>;
  /** DownloadBlob downloads a file or buffer from the calling node's service IPFS storage. */
  DownloadBlob(request: DownloadBlobRequest): Observable<DownloadBlobResponse>;
  /** SyncBlob synchronizes a local file from the calling node to the given service's IPFS storage. */
  SyncBlob(request: SyncBlobRequest): Observable<SyncBlobResponse>;
  /** DeleteBlob deletes the given file from the calling node's service IPFS storage. */
  DeleteBlob(request: DeleteBlobRequest): Promise<DeleteBlobResponse>;
  /** ParseDid parses a potential DID string into a DID object. */
  ParseDid(request: ParseDidRequest): Promise<ParseDidResponse>;
  /**
   * ResolveDid resolves a DID to its DID document if the DID is valid and the calling node has
   * access to the DID Document.
   */
  ResolveDid(request: ResolveDidRequest): Promise<ResolveDidResponse>;
}

export class HighwayServiceClientImpl implements HighwayService {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.AccessName = this.AccessName.bind(this);
    this.RegisterName = this.RegisterName.bind(this);
    this.UpdateName = this.UpdateName.bind(this);
    this.AccessService = this.AccessService.bind(this);
    this.RegisterService = this.RegisterService.bind(this);
    this.UpdateService = this.UpdateService.bind(this);
    this.CreateChannel = this.CreateChannel.bind(this);
    this.ReadChannel = this.ReadChannel.bind(this);
    this.UpdateChannel = this.UpdateChannel.bind(this);
    this.DeleteChannel = this.DeleteChannel.bind(this);
    this.ListenChannel = this.ListenChannel.bind(this);
    this.CreateBucket = this.CreateBucket.bind(this);
    this.ReadBucket = this.ReadBucket.bind(this);
    this.UpdateBucket = this.UpdateBucket.bind(this);
    this.DeleteBucket = this.DeleteBucket.bind(this);
    this.ListenBucket = this.ListenBucket.bind(this);
    this.CreateObject = this.CreateObject.bind(this);
    this.ReadObject = this.ReadObject.bind(this);
    this.UpdateObject = this.UpdateObject.bind(this);
    this.DeleteObject = this.DeleteObject.bind(this);
    this.UploadBlob = this.UploadBlob.bind(this);
    this.DownloadBlob = this.DownloadBlob.bind(this);
    this.SyncBlob = this.SyncBlob.bind(this);
    this.DeleteBlob = this.DeleteBlob.bind(this);
    this.ParseDid = this.ParseDid.bind(this);
    this.ResolveDid = this.ResolveDid.bind(this);
  }
  AccessName(request: AccessNameRequest): Promise<AccessNameResponse> {
    const data = AccessNameRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "AccessName",
      data
    );
    return promise.then((data) =>
      AccessNameResponse.decode(new _m0.Reader(data))
    );
  }

  RegisterName(request: RegisterNameRequest): Promise<RegisterNameResponse> {
    const data = RegisterNameRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "RegisterName",
      data
    );
    return promise.then((data) =>
      RegisterNameResponse.decode(new _m0.Reader(data))
    );
  }

  UpdateName(request: UpdateNameRequest): Promise<UpdateNameResponse> {
    const data = UpdateNameRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "UpdateName",
      data
    );
    return promise.then((data) =>
      UpdateNameResponse.decode(new _m0.Reader(data))
    );
  }

  AccessService(request: AccessServiceRequest): Promise<AccessServiceResponse> {
    const data = AccessServiceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "AccessService",
      data
    );
    return promise.then((data) =>
      AccessServiceResponse.decode(new _m0.Reader(data))
    );
  }

  RegisterService(
    request: RegisterServiceRequest
  ): Promise<RegisterServiceResponse> {
    const data = RegisterServiceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "RegisterService",
      data
    );
    return promise.then((data) =>
      RegisterServiceResponse.decode(new _m0.Reader(data))
    );
  }

  UpdateService(request: UpdateServiceRequest): Promise<UpdateServiceResponse> {
    const data = UpdateServiceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "UpdateService",
      data
    );
    return promise.then((data) =>
      UpdateServiceResponse.decode(new _m0.Reader(data))
    );
  }

  CreateChannel(request: CreateChannelRequest): Promise<CreateChannelResponse> {
    const data = CreateChannelRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "CreateChannel",
      data
    );
    return promise.then((data) =>
      CreateChannelResponse.decode(new _m0.Reader(data))
    );
  }

  ReadChannel(request: ReadChannelRequest): Promise<ReadChannelResponse> {
    const data = ReadChannelRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "ReadChannel",
      data
    );
    return promise.then((data) =>
      ReadChannelResponse.decode(new _m0.Reader(data))
    );
  }

  UpdateChannel(request: UpdateChannelRequest): Promise<UpdateChannelResponse> {
    const data = UpdateChannelRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "UpdateChannel",
      data
    );
    return promise.then((data) =>
      UpdateChannelResponse.decode(new _m0.Reader(data))
    );
  }

  DeleteChannel(request: DeleteChannelRequest): Promise<DeleteChannelResponse> {
    const data = DeleteChannelRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "DeleteChannel",
      data
    );
    return promise.then((data) =>
      DeleteChannelResponse.decode(new _m0.Reader(data))
    );
  }

  ListenChannel(
    request: ListenChannelRequest
  ): Observable<ListenChannelResponse> {
    const data = ListenChannelRequest.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(
      "node.highway.v1.HighwayService",
      "ListenChannel",
      data
    );
    return result.pipe(
      map((data) => ListenChannelResponse.decode(new _m0.Reader(data)))
    );
  }

  CreateBucket(request: CreateBucketRequest): Promise<CreateBucketResponse> {
    const data = CreateBucketRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "CreateBucket",
      data
    );
    return promise.then((data) =>
      CreateBucketResponse.decode(new _m0.Reader(data))
    );
  }

  ReadBucket(request: ReadBucketRequest): Promise<ReadBucketResponse> {
    const data = ReadBucketRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "ReadBucket",
      data
    );
    return promise.then((data) =>
      ReadBucketResponse.decode(new _m0.Reader(data))
    );
  }

  UpdateBucket(request: UpdateBucketRequest): Promise<UpdateBucketResponse> {
    const data = UpdateBucketRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "UpdateBucket",
      data
    );
    return promise.then((data) =>
      UpdateBucketResponse.decode(new _m0.Reader(data))
    );
  }

  DeleteBucket(request: DeleteBucketRequest): Promise<DeleteBucketResponse> {
    const data = DeleteBucketRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "DeleteBucket",
      data
    );
    return promise.then((data) =>
      DeleteBucketResponse.decode(new _m0.Reader(data))
    );
  }

  ListenBucket(request: ListenBucketRequest): Observable<ListenBucketResponse> {
    const data = ListenBucketRequest.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(
      "node.highway.v1.HighwayService",
      "ListenBucket",
      data
    );
    return result.pipe(
      map((data) => ListenBucketResponse.decode(new _m0.Reader(data)))
    );
  }

  CreateObject(request: CreateObjectRequest): Promise<CreateObjectResponse> {
    const data = CreateObjectRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "CreateObject",
      data
    );
    return promise.then((data) =>
      CreateObjectResponse.decode(new _m0.Reader(data))
    );
  }

  ReadObject(request: ReadObjectRequest): Promise<ReadObjectResponse> {
    const data = ReadObjectRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "ReadObject",
      data
    );
    return promise.then((data) =>
      ReadObjectResponse.decode(new _m0.Reader(data))
    );
  }

  UpdateObject(request: UpdateObjectRequest): Promise<UpdateObjectResponse> {
    const data = UpdateObjectRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "UpdateObject",
      data
    );
    return promise.then((data) =>
      UpdateObjectResponse.decode(new _m0.Reader(data))
    );
  }

  DeleteObject(request: DeleteObjectRequest): Promise<DeleteObjectResponse> {
    const data = DeleteObjectRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "DeleteObject",
      data
    );
    return promise.then((data) =>
      DeleteObjectResponse.decode(new _m0.Reader(data))
    );
  }

  UploadBlob(request: UploadBlobRequest): Observable<UploadBlobResponse> {
    const data = UploadBlobRequest.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(
      "node.highway.v1.HighwayService",
      "UploadBlob",
      data
    );
    return result.pipe(
      map((data) => UploadBlobResponse.decode(new _m0.Reader(data)))
    );
  }

  DownloadBlob(request: DownloadBlobRequest): Observable<DownloadBlobResponse> {
    const data = DownloadBlobRequest.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(
      "node.highway.v1.HighwayService",
      "DownloadBlob",
      data
    );
    return result.pipe(
      map((data) => DownloadBlobResponse.decode(new _m0.Reader(data)))
    );
  }

  SyncBlob(request: SyncBlobRequest): Observable<SyncBlobResponse> {
    const data = SyncBlobRequest.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(
      "node.highway.v1.HighwayService",
      "SyncBlob",
      data
    );
    return result.pipe(
      map((data) => SyncBlobResponse.decode(new _m0.Reader(data)))
    );
  }

  DeleteBlob(request: DeleteBlobRequest): Promise<DeleteBlobResponse> {
    const data = DeleteBlobRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "DeleteBlob",
      data
    );
    return promise.then((data) =>
      DeleteBlobResponse.decode(new _m0.Reader(data))
    );
  }

  ParseDid(request: ParseDidRequest): Promise<ParseDidResponse> {
    const data = ParseDidRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "ParseDid",
      data
    );
    return promise.then((data) =>
      ParseDidResponse.decode(new _m0.Reader(data))
    );
  }

  ResolveDid(request: ResolveDidRequest): Promise<ResolveDidResponse> {
    const data = ResolveDidRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.highway.v1.HighwayService",
      "ResolveDid",
      data
    );
    return promise.then((data) =>
      ResolveDidResponse.decode(new _m0.Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
  clientStreamingRequest(
    service: string,
    method: string,
    data: Observable<Uint8Array>
  ): Promise<Uint8Array>;
  serverStreamingRequest(
    service: string,
    method: string,
    data: Uint8Array
  ): Observable<Uint8Array>;
  bidirectionalStreamingRequest(
    service: string,
    method: string,
    data: Observable<Uint8Array>
  ): Observable<Uint8Array>;
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}
