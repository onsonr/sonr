// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgReadChannel } from "./types/channel/tx";
import { MsgCreateChannel } from "./types/channel/tx";
import { MsgDeactivateChannel } from "./types/channel/tx";
import { MsgUpdateChannel } from "./types/channel/tx";


const types = [
  ["/sonrio.sonr.channel.MsgReadChannel", MsgReadChannel],
  ["/sonrio.sonr.channel.MsgCreateChannel", MsgCreateChannel],
  ["/sonrio.sonr.channel.MsgDeactivateChannel", MsgDeactivateChannel],
  ["/sonrio.sonr.channel.MsgUpdateChannel", MsgUpdateChannel],

];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgReadChannel: (data: MsgReadChannel): EncodeObject => ({ typeUrl: "/sonrio.sonr.channel.MsgReadChannel", value: MsgReadChannel.fromPartial( data ) }),
    msgCreateChannel: (data: MsgCreateChannel): EncodeObject => ({ typeUrl: "/sonrio.sonr.channel.MsgCreateChannel", value: MsgCreateChannel.fromPartial( data ) }),
    MsgDeactivateChannel: (data: MsgDeactivateChannel): EncodeObject => ({ typeUrl: "/sonrio.sonr.channel.MsgDeactivateChannel", value: MsgDeactivateChannel.fromPartial( data ) }),
    msgUpdateChannel: (data: MsgUpdateChannel): EncodeObject => ({ typeUrl: "/sonrio.sonr.channel.MsgUpdateChannel", value: MsgUpdateChannel.fromPartial( data ) }),

  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
