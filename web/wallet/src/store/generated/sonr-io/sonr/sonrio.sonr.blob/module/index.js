// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSyncBlob } from "./types/blob/tx";
import { MsgDownloadBlob } from "./types/blob/tx";
import { MsgDeleteBlob } from "./types/blob/tx";
import { MsgUploadBlob } from "./types/blob/tx";
const types = [
    ["/sonrio.sonr.blob.MsgSyncBlob", MsgSyncBlob],
    ["/sonrio.sonr.blob.MsgDownloadBlob", MsgDownloadBlob],
    ["/sonrio.sonr.blob.MsgDeleteBlob", MsgDeleteBlob],
    ["/sonrio.sonr.blob.MsgUploadBlob", MsgUploadBlob],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgSyncBlob: (data) => ({ typeUrl: "/sonrio.sonr.blob.MsgSyncBlob", value: MsgSyncBlob.fromPartial(data) }),
        msgDownloadBlob: (data) => ({ typeUrl: "/sonrio.sonr.blob.MsgDownloadBlob", value: MsgDownloadBlob.fromPartial(data) }),
        msgDeleteBlob: (data) => ({ typeUrl: "/sonrio.sonr.blob.MsgDeleteBlob", value: MsgDeleteBlob.fromPartial(data) }),
        msgUploadBlob: (data) => ({ typeUrl: "/sonrio.sonr.blob.MsgUploadBlob", value: MsgUploadBlob.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
