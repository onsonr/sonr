// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeactivateObject } from "./types/object/tx";
import { MsgReadObject } from "./types/object/tx";
import { MsgCreateObject } from "./types/object/tx";
import { MsgUpdateObject } from "./types/object/tx";
const types = [
    ["/sonrio.sonr.object.MsgDeactivateObject", MsgDeactivateObject],
    ["/sonrio.sonr.object.MsgReadObject", MsgReadObject],
    ["/sonrio.sonr.object.MsgCreateObject", MsgCreateObject],
    ["/sonrio.sonr.object.MsgUpdateObject", MsgUpdateObject],
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
        MsgDeactivateObject: (data) => ({ typeUrl: "/sonrio.sonr.object.MsgDeactivateObject", value: MsgDeactivateObject.fromPartial(data) }),
        msgReadObject: (data) => ({ typeUrl: "/sonrio.sonr.object.MsgReadObject", value: MsgReadObject.fromPartial(data) }),
        msgCreateObject: (data) => ({ typeUrl: "/sonrio.sonr.object.MsgCreateObject", value: MsgCreateObject.fromPartial(data) }),
        msgUpdateObject: (data) => ({ typeUrl: "/sonrio.sonr.object.MsgUpdateObject", value: MsgUpdateObject.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
