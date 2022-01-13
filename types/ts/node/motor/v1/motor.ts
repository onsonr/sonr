/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
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
import { map } from "rxjs/operators";
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

export const protobufPackage = "node.motor.v1";

/** / This file contains service for the Node RPC Server */

/** RPC Service with Equivalent Methods of a binded Node */
export interface MotorService {
  /**
   * Node Methods
   * Respond Method to an Invite with Decision
   */
  Share(request: ShareRequest): Promise<ShareResponse>;
  /** Respond Method to an Invite with Decision */
  Decide(request: DecideRequest): Promise<DecideResponse>;
  /** Search Method to find a Peer by SName or PeerID */
  Search(request: SearchRequest): Promise<SearchResponse>;
  /**
   * Events Streams
   * Returns a stream of Lobby Refresh Events
   */
  OnLobbyRefresh(
    request: OnLobbyRefreshRequest
  ): Observable<OnLobbyRefreshResponse>;
  /** Returns a stream of Mailbox Message Events */
  OnMailboxMessage(
    request: OnMailboxMessageRequest
  ): Observable<OnMailboxMessageResponse>;
  /** Returns a stream of DecisionEvent's for Accepted Invites */
  OnTransmitDecision(
    request: OnTransmitDecisionRequest
  ): Observable<OnTransmitDecisionResponse>;
  /** Returns a stream of DecisionEvent's for Invites */
  OnTransmitInvite(
    request: OnTransmitInviteRequest
  ): Observable<OnTransmitInviteResponse>;
  /** Returns a stream of ProgressEvent's for Sessions */
  OnTransmitProgress(
    request: OnTransmitProgressRequest
  ): Observable<OnTransmitProgressResponse>;
  /** Returns a stream of Completed Transfers */
  OnTransmitComplete(
    request: OnTransmitCompleteRequest
  ): Observable<OnTransmitCompleteResponse>;
}

export class MotorServiceClientImpl implements MotorService {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Share = this.Share.bind(this);
    this.Decide = this.Decide.bind(this);
    this.Search = this.Search.bind(this);
    this.OnLobbyRefresh = this.OnLobbyRefresh.bind(this);
    this.OnMailboxMessage = this.OnMailboxMessage.bind(this);
    this.OnTransmitDecision = this.OnTransmitDecision.bind(this);
    this.OnTransmitInvite = this.OnTransmitInvite.bind(this);
    this.OnTransmitProgress = this.OnTransmitProgress.bind(this);
    this.OnTransmitComplete = this.OnTransmitComplete.bind(this);
  }
  Share(request: ShareRequest): Promise<ShareResponse> {
    const data = ShareRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.motor.v1.MotorService",
      "Share",
      data
    );
    return promise.then((data) => ShareResponse.decode(new _m0.Reader(data)));
  }

  Decide(request: DecideRequest): Promise<DecideResponse> {
    const data = DecideRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.motor.v1.MotorService",
      "Decide",
      data
    );
    return promise.then((data) => DecideResponse.decode(new _m0.Reader(data)));
  }

  Search(request: SearchRequest): Promise<SearchResponse> {
    const data = SearchRequest.encode(request).finish();
    const promise = this.rpc.request(
      "node.motor.v1.MotorService",
      "Search",
      data
    );
    return promise.then((data) => SearchResponse.decode(new _m0.Reader(data)));
  }

  OnLobbyRefresh(
    request: OnLobbyRefreshRequest
  ): Observable<OnLobbyRefreshResponse> {
    const data = OnLobbyRefreshRequest.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(
      "node.motor.v1.MotorService",
      "OnLobbyRefresh",
      data
    );
    return result.pipe(
      map((data) => OnLobbyRefreshResponse.decode(new _m0.Reader(data)))
    );
  }

  OnMailboxMessage(
    request: OnMailboxMessageRequest
  ): Observable<OnMailboxMessageResponse> {
    const data = OnMailboxMessageRequest.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(
      "node.motor.v1.MotorService",
      "OnMailboxMessage",
      data
    );
    return result.pipe(
      map((data) => OnMailboxMessageResponse.decode(new _m0.Reader(data)))
    );
  }

  OnTransmitDecision(
    request: OnTransmitDecisionRequest
  ): Observable<OnTransmitDecisionResponse> {
    const data = OnTransmitDecisionRequest.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(
      "node.motor.v1.MotorService",
      "OnTransmitDecision",
      data
    );
    return result.pipe(
      map((data) => OnTransmitDecisionResponse.decode(new _m0.Reader(data)))
    );
  }

  OnTransmitInvite(
    request: OnTransmitInviteRequest
  ): Observable<OnTransmitInviteResponse> {
    const data = OnTransmitInviteRequest.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(
      "node.motor.v1.MotorService",
      "OnTransmitInvite",
      data
    );
    return result.pipe(
      map((data) => OnTransmitInviteResponse.decode(new _m0.Reader(data)))
    );
  }

  OnTransmitProgress(
    request: OnTransmitProgressRequest
  ): Observable<OnTransmitProgressResponse> {
    const data = OnTransmitProgressRequest.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(
      "node.motor.v1.MotorService",
      "OnTransmitProgress",
      data
    );
    return result.pipe(
      map((data) => OnTransmitProgressResponse.decode(new _m0.Reader(data)))
    );
  }

  OnTransmitComplete(
    request: OnTransmitCompleteRequest
  ): Observable<OnTransmitCompleteResponse> {
    const data = OnTransmitCompleteRequest.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(
      "node.motor.v1.MotorService",
      "OnTransmitComplete",
      data
    );
    return result.pipe(
      map((data) => OnTransmitCompleteResponse.decode(new _m0.Reader(data)))
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
