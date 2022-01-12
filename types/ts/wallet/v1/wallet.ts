/* eslint-disable */
import { util, configure } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "wallet.v1";

/** TokenType is the type of keychain. */
export enum TokenType {
  TOKEN_TYPE_UNSPECIFIED = 0,
  TOKEN_TYPE_SNR = 1,
  TOKEN_TYPE_ETH = 2,
  TOKEN_TYPE_BTC = 3,
  TOKEN_TYPE_SOL = 4,
  UNRECOGNIZED = -1,
}

export function tokenTypeFromJSON(object: any): TokenType {
  switch (object) {
    case 0:
    case "TOKEN_TYPE_UNSPECIFIED":
      return TokenType.TOKEN_TYPE_UNSPECIFIED;
    case 1:
    case "TOKEN_TYPE_SNR":
      return TokenType.TOKEN_TYPE_SNR;
    case 2:
    case "TOKEN_TYPE_ETH":
      return TokenType.TOKEN_TYPE_ETH;
    case 3:
    case "TOKEN_TYPE_BTC":
      return TokenType.TOKEN_TYPE_BTC;
    case 4:
    case "TOKEN_TYPE_SOL":
      return TokenType.TOKEN_TYPE_SOL;
    case -1:
    case "UNRECOGNIZED":
    default:
      return TokenType.UNRECOGNIZED;
  }
}

export function tokenTypeToJSON(object: TokenType): string {
  switch (object) {
    case TokenType.TOKEN_TYPE_UNSPECIFIED:
      return "TOKEN_TYPE_UNSPECIFIED";
    case TokenType.TOKEN_TYPE_SNR:
      return "TOKEN_TYPE_SNR";
    case TokenType.TOKEN_TYPE_ETH:
      return "TOKEN_TYPE_ETH";
    case TokenType.TOKEN_TYPE_BTC:
      return "TOKEN_TYPE_BTC";
    case TokenType.TOKEN_TYPE_SOL:
      return "TOKEN_TYPE_SOL";
    default:
      return "UNKNOWN";
  }
}

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
