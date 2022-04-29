import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgReadChannel } from "./types/channel/tx";
import { MsgCreateChannel } from "./types/channel/tx";
import { MsgDeactivateChannel } from "./types/channel/tx";
import { MsgUpdateChannel } from "./types/channel/tx";
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
    msgReadChannel: (data: MsgReadChannel) => EncodeObject;
    msgCreateChannel: (data: MsgCreateChannel) => EncodeObject;
    MsgDeactivateChannel: (data: MsgDeactivateChannel) => EncodeObject;
    msgUpdateChannel: (data: MsgUpdateChannel) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
