import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAccessService } from "./types/registry/tx";
import { MsgRegisterName } from "./types/registry/tx";
import { MsgAccessName } from "./types/registry/tx";
import { MsgRegisterService } from "./types/registry/tx";
import { MsgUpdateService } from "./types/registry/tx";
import { MsgUpdateName } from "./types/registry/tx";
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
    msgAccessService: (data: MsgAccessService) => EncodeObject;
    msgRegisterName: (data: MsgRegisterName) => EncodeObject;
    msgAccessName: (data: MsgAccessName) => EncodeObject;
    msgRegisterService: (data: MsgRegisterService) => EncodeObject;
    msgUpdateService: (data: MsgUpdateService) => EncodeObject;
    msgUpdateName: (data: MsgUpdateName) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
