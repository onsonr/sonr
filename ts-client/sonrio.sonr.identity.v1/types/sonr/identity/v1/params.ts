/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "sonrio.sonr.identity.v1";

/** Params defines the parameters for the module. */
export interface Params {
  didBaseContext: string;
  didImplementationContext: string;
  ipfsGateway: string;
  ipfsApi: string;
}

function createBaseParams(): Params {
  return { didBaseContext: "", didImplementationContext: "", ipfsGateway: "", ipfsApi: "" };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.didBaseContext !== "") {
      writer.uint32(10).string(message.didBaseContext);
    }
    if (message.didImplementationContext !== "") {
      writer.uint32(18).string(message.didImplementationContext);
    }
    if (message.ipfsGateway !== "") {
      writer.uint32(26).string(message.ipfsGateway);
    }
    if (message.ipfsApi !== "") {
      writer.uint32(34).string(message.ipfsApi);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.didBaseContext = reader.string();
          break;
        case 2:
          message.didImplementationContext = reader.string();
          break;
        case 3:
          message.ipfsGateway = reader.string();
          break;
        case 4:
          message.ipfsApi = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    return {
      didBaseContext: isSet(object.didBaseContext) ? String(object.didBaseContext) : "",
      didImplementationContext: isSet(object.didImplementationContext) ? String(object.didImplementationContext) : "",
      ipfsGateway: isSet(object.ipfsGateway) ? String(object.ipfsGateway) : "",
      ipfsApi: isSet(object.ipfsApi) ? String(object.ipfsApi) : "",
    };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.didBaseContext !== undefined && (obj.didBaseContext = message.didBaseContext);
    message.didImplementationContext !== undefined && (obj.didImplementationContext = message.didImplementationContext);
    message.ipfsGateway !== undefined && (obj.ipfsGateway = message.ipfsGateway);
    message.ipfsApi !== undefined && (obj.ipfsApi = message.ipfsApi);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.didBaseContext = object.didBaseContext ?? "";
    message.didImplementationContext = object.didImplementationContext ?? "";
    message.ipfsGateway = object.ipfsGateway ?? "";
    message.ipfsApi = object.ipfsApi ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
