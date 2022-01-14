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
  ShareRequest,
  DecideRequest,
  SearchRequest,
  OnLobbyRefreshRequest,
  OnMailboxMessageRequest,
  OnTransmitDecisionRequest,
  OnTransmitInviteRequest,
  OnTransmitProgressRequest,
  OnTransmitCompleteRequest,
} from "../../../node/motor/v1/request";
import {
  ShareResponse,
  DecideResponse,
  SearchResponse,
  OnLobbyRefreshResponse,
  OnMailboxMessageResponse,
  OnTransmitDecisionResponse,
  OnTransmitInviteResponse,
  OnTransmitProgressResponse,
  OnTransmitCompleteResponse,
} from "../../../node/motor/v1/response";

export const protobufPackage = "node.motor.v1";

/** MotorService is a RPC service for interfacing over the Motor node. */
export const MotorServiceService = {
  /**
   * Node Methods
   * Respond Method to an Invite with Decision
   */
  share: {
    path: "/node.motor.v1.MotorService/Share",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: ShareRequest) =>
      Buffer.from(ShareRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => ShareRequest.decode(value),
    responseSerialize: (value: ShareResponse) =>
      Buffer.from(ShareResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => ShareResponse.decode(value),
  },
  /** Respond Method to an Invite with Decision */
  decide: {
    path: "/node.motor.v1.MotorService/Decide",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: DecideRequest) =>
      Buffer.from(DecideRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => DecideRequest.decode(value),
    responseSerialize: (value: DecideResponse) =>
      Buffer.from(DecideResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => DecideResponse.decode(value),
  },
  /** Search Method to find a Peer by SName or PeerID */
  search: {
    path: "/node.motor.v1.MotorService/Search",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: SearchRequest) =>
      Buffer.from(SearchRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => SearchRequest.decode(value),
    responseSerialize: (value: SearchResponse) =>
      Buffer.from(SearchResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => SearchResponse.decode(value),
  },
  /**
   * Events Streams
   * Returns a stream of Lobby Refresh Events
   */
  onLobbyRefresh: {
    path: "/node.motor.v1.MotorService/OnLobbyRefresh",
    requestStream: false,
    responseStream: true,
    requestSerialize: (value: OnLobbyRefreshRequest) =>
      Buffer.from(OnLobbyRefreshRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => OnLobbyRefreshRequest.decode(value),
    responseSerialize: (value: OnLobbyRefreshResponse) =>
      Buffer.from(OnLobbyRefreshResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) =>
      OnLobbyRefreshResponse.decode(value),
  },
  /** Returns a stream of Mailbox Message Events */
  onMailboxMessage: {
    path: "/node.motor.v1.MotorService/OnMailboxMessage",
    requestStream: false,
    responseStream: true,
    requestSerialize: (value: OnMailboxMessageRequest) =>
      Buffer.from(OnMailboxMessageRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) =>
      OnMailboxMessageRequest.decode(value),
    responseSerialize: (value: OnMailboxMessageResponse) =>
      Buffer.from(OnMailboxMessageResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) =>
      OnMailboxMessageResponse.decode(value),
  },
  /** Returns a stream of DecisionEvent's for Accepted Invites */
  onTransmitDecision: {
    path: "/node.motor.v1.MotorService/OnTransmitDecision",
    requestStream: false,
    responseStream: true,
    requestSerialize: (value: OnTransmitDecisionRequest) =>
      Buffer.from(OnTransmitDecisionRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) =>
      OnTransmitDecisionRequest.decode(value),
    responseSerialize: (value: OnTransmitDecisionResponse) =>
      Buffer.from(OnTransmitDecisionResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) =>
      OnTransmitDecisionResponse.decode(value),
  },
  /** Returns a stream of DecisionEvent's for Invites */
  onTransmitInvite: {
    path: "/node.motor.v1.MotorService/OnTransmitInvite",
    requestStream: false,
    responseStream: true,
    requestSerialize: (value: OnTransmitInviteRequest) =>
      Buffer.from(OnTransmitInviteRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) =>
      OnTransmitInviteRequest.decode(value),
    responseSerialize: (value: OnTransmitInviteResponse) =>
      Buffer.from(OnTransmitInviteResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) =>
      OnTransmitInviteResponse.decode(value),
  },
  /** Returns a stream of ProgressEvent's for Sessions */
  onTransmitProgress: {
    path: "/node.motor.v1.MotorService/OnTransmitProgress",
    requestStream: false,
    responseStream: true,
    requestSerialize: (value: OnTransmitProgressRequest) =>
      Buffer.from(OnTransmitProgressRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) =>
      OnTransmitProgressRequest.decode(value),
    responseSerialize: (value: OnTransmitProgressResponse) =>
      Buffer.from(OnTransmitProgressResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) =>
      OnTransmitProgressResponse.decode(value),
  },
  /** Returns a stream of Completed Transfers */
  onTransmitComplete: {
    path: "/node.motor.v1.MotorService/OnTransmitComplete",
    requestStream: false,
    responseStream: true,
    requestSerialize: (value: OnTransmitCompleteRequest) =>
      Buffer.from(OnTransmitCompleteRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) =>
      OnTransmitCompleteRequest.decode(value),
    responseSerialize: (value: OnTransmitCompleteResponse) =>
      Buffer.from(OnTransmitCompleteResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) =>
      OnTransmitCompleteResponse.decode(value),
  },
} as const;

export interface MotorServiceServer extends UntypedServiceImplementation {
  /**
   * Node Methods
   * Respond Method to an Invite with Decision
   */
  share: handleUnaryCall<ShareRequest, ShareResponse>;
  /** Respond Method to an Invite with Decision */
  decide: handleUnaryCall<DecideRequest, DecideResponse>;
  /** Search Method to find a Peer by SName or PeerID */
  search: handleUnaryCall<SearchRequest, SearchResponse>;
  /**
   * Events Streams
   * Returns a stream of Lobby Refresh Events
   */
  onLobbyRefresh: handleServerStreamingCall<
    OnLobbyRefreshRequest,
    OnLobbyRefreshResponse
  >;
  /** Returns a stream of Mailbox Message Events */
  onMailboxMessage: handleServerStreamingCall<
    OnMailboxMessageRequest,
    OnMailboxMessageResponse
  >;
  /** Returns a stream of DecisionEvent's for Accepted Invites */
  onTransmitDecision: handleServerStreamingCall<
    OnTransmitDecisionRequest,
    OnTransmitDecisionResponse
  >;
  /** Returns a stream of DecisionEvent's for Invites */
  onTransmitInvite: handleServerStreamingCall<
    OnTransmitInviteRequest,
    OnTransmitInviteResponse
  >;
  /** Returns a stream of ProgressEvent's for Sessions */
  onTransmitProgress: handleServerStreamingCall<
    OnTransmitProgressRequest,
    OnTransmitProgressResponse
  >;
  /** Returns a stream of Completed Transfers */
  onTransmitComplete: handleServerStreamingCall<
    OnTransmitCompleteRequest,
    OnTransmitCompleteResponse
  >;
}

export interface MotorServiceClient extends Client {
  /**
   * Node Methods
   * Respond Method to an Invite with Decision
   */
  share(
    request: ShareRequest,
    callback: (error: ServiceError | null, response: ShareResponse) => void
  ): ClientUnaryCall;
  share(
    request: ShareRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: ShareResponse) => void
  ): ClientUnaryCall;
  share(
    request: ShareRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: ShareResponse) => void
  ): ClientUnaryCall;
  /** Respond Method to an Invite with Decision */
  decide(
    request: DecideRequest,
    callback: (error: ServiceError | null, response: DecideResponse) => void
  ): ClientUnaryCall;
  decide(
    request: DecideRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: DecideResponse) => void
  ): ClientUnaryCall;
  decide(
    request: DecideRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: DecideResponse) => void
  ): ClientUnaryCall;
  /** Search Method to find a Peer by SName or PeerID */
  search(
    request: SearchRequest,
    callback: (error: ServiceError | null, response: SearchResponse) => void
  ): ClientUnaryCall;
  search(
    request: SearchRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: SearchResponse) => void
  ): ClientUnaryCall;
  search(
    request: SearchRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: SearchResponse) => void
  ): ClientUnaryCall;
  /**
   * Events Streams
   * Returns a stream of Lobby Refresh Events
   */
  onLobbyRefresh(
    request: OnLobbyRefreshRequest,
    options?: Partial<CallOptions>
  ): ClientReadableStream<OnLobbyRefreshResponse>;
  onLobbyRefresh(
    request: OnLobbyRefreshRequest,
    metadata?: Metadata,
    options?: Partial<CallOptions>
  ): ClientReadableStream<OnLobbyRefreshResponse>;
  /** Returns a stream of Mailbox Message Events */
  onMailboxMessage(
    request: OnMailboxMessageRequest,
    options?: Partial<CallOptions>
  ): ClientReadableStream<OnMailboxMessageResponse>;
  onMailboxMessage(
    request: OnMailboxMessageRequest,
    metadata?: Metadata,
    options?: Partial<CallOptions>
  ): ClientReadableStream<OnMailboxMessageResponse>;
  /** Returns a stream of DecisionEvent's for Accepted Invites */
  onTransmitDecision(
    request: OnTransmitDecisionRequest,
    options?: Partial<CallOptions>
  ): ClientReadableStream<OnTransmitDecisionResponse>;
  onTransmitDecision(
    request: OnTransmitDecisionRequest,
    metadata?: Metadata,
    options?: Partial<CallOptions>
  ): ClientReadableStream<OnTransmitDecisionResponse>;
  /** Returns a stream of DecisionEvent's for Invites */
  onTransmitInvite(
    request: OnTransmitInviteRequest,
    options?: Partial<CallOptions>
  ): ClientReadableStream<OnTransmitInviteResponse>;
  onTransmitInvite(
    request: OnTransmitInviteRequest,
    metadata?: Metadata,
    options?: Partial<CallOptions>
  ): ClientReadableStream<OnTransmitInviteResponse>;
  /** Returns a stream of ProgressEvent's for Sessions */
  onTransmitProgress(
    request: OnTransmitProgressRequest,
    options?: Partial<CallOptions>
  ): ClientReadableStream<OnTransmitProgressResponse>;
  onTransmitProgress(
    request: OnTransmitProgressRequest,
    metadata?: Metadata,
    options?: Partial<CallOptions>
  ): ClientReadableStream<OnTransmitProgressResponse>;
  /** Returns a stream of Completed Transfers */
  onTransmitComplete(
    request: OnTransmitCompleteRequest,
    options?: Partial<CallOptions>
  ): ClientReadableStream<OnTransmitCompleteResponse>;
  onTransmitComplete(
    request: OnTransmitCompleteRequest,
    metadata?: Metadata,
    options?: Partial<CallOptions>
  ): ClientReadableStream<OnTransmitCompleteResponse>;
}

export const MotorServiceClient = makeGenericClientConstructor(
  MotorServiceService,
  "node.motor.v1.MotorService"
) as unknown as {
  new (
    address: string,
    credentials: ChannelCredentials,
    options?: Partial<ChannelOptions>
  ): MotorServiceClient;
  service: typeof MotorServiceService;
};

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}
