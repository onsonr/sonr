import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeactivateObject } from "./types/object/tx";
import { MsgReadObject } from "./types/object/tx";
import { MsgCreateObject } from "./types/object/tx";
import { MsgUpdateObject } from "./types/object/tx";
export declare const MissingWalletError: Error;
export declare const registry: Registry;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => any;
    MsgDeactivateObject: (data: MsgDeactivateObject) => EncodeObject;
    msgReadObject: (data: MsgReadObject) => EncodeObject;
    msgCreateObject: (data: MsgCreateObject) => EncodeObject;
    msgUpdateObject: (data: MsgUpdateObject) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
