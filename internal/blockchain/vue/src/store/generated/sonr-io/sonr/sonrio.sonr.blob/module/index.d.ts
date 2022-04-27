import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSyncBlob } from "./types/blob/tx";
import { MsgDownloadBlob } from "./types/blob/tx";
import { MsgDeleteBlob } from "./types/blob/tx";
import { MsgUploadBlob } from "./types/blob/tx";
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
    msgSyncBlob: (data: MsgSyncBlob) => EncodeObject;
    msgDownloadBlob: (data: MsgDownloadBlob) => EncodeObject;
    msgDeleteBlob: (data: MsgDeleteBlob) => EncodeObject;
    msgUploadBlob: (data: MsgUploadBlob) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
