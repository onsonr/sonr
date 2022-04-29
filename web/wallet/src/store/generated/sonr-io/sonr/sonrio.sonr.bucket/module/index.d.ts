import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeactivateBucket } from "./types/bucket/tx";
import { MsgCreateBucket } from "./types/bucket/tx";
import { MsgReadBucket } from "./types/bucket/tx";
import { MsgUpdateBucket } from "./types/bucket/tx";
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
    MsgDeactivateBucket: (data: MsgDeactivateBucket) => EncodeObject;
    msgCreateBucket: (data: MsgCreateBucket) => EncodeObject;
    msgReadBucket: (data: MsgReadBucket) => EncodeObject;
    msgUpdateBucket: (data: MsgUpdateBucket) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
