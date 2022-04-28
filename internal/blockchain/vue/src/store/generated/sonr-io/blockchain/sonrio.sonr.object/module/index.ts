// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteWhatIs } from "./types/object/tx";
import { MsgUpdateObject } from "./types/object/tx";
import { MsgDeactivateObject } from "./types/object/tx";
import { MsgCreateWhatIs } from "./types/object/tx";
import { MsgUpdateWhatIs } from "./types/object/tx";
import { MsgCreateObject } from "./types/object/tx";


const types = [
  ["/sonrio.sonr.object.MsgDeleteWhatIs", MsgDeleteWhatIs],
  ["/sonrio.sonr.object.MsgUpdateObject", MsgUpdateObject],
  ["/sonrio.sonr.object.MsgDeactivateObject", MsgDeactivateObject],
  ["/sonrio.sonr.object.MsgCreateWhatIs", MsgCreateWhatIs],
  ["/sonrio.sonr.object.MsgUpdateWhatIs", MsgUpdateWhatIs],
  ["/sonrio.sonr.object.MsgCreateObject", MsgCreateObject],
  
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
    msgDeleteWhatIs: (data: MsgDeleteWhatIs): EncodeObject => ({ typeUrl: "/sonrio.sonr.object.MsgDeleteWhatIs", value: MsgDeleteWhatIs.fromPartial( data ) }),
    msgUpdateObject: (data: MsgUpdateObject): EncodeObject => ({ typeUrl: "/sonrio.sonr.object.MsgUpdateObject", value: MsgUpdateObject.fromPartial( data ) }),
    msgDeactivateObject: (data: MsgDeactivateObject): EncodeObject => ({ typeUrl: "/sonrio.sonr.object.MsgDeactivateObject", value: MsgDeactivateObject.fromPartial( data ) }),
    msgCreateWhatIs: (data: MsgCreateWhatIs): EncodeObject => ({ typeUrl: "/sonrio.sonr.object.MsgCreateWhatIs", value: MsgCreateWhatIs.fromPartial( data ) }),
    msgUpdateWhatIs: (data: MsgUpdateWhatIs): EncodeObject => ({ typeUrl: "/sonrio.sonr.object.MsgUpdateWhatIs", value: MsgUpdateWhatIs.fromPartial( data ) }),
    msgCreateObject: (data: MsgCreateObject): EncodeObject => ({ typeUrl: "/sonrio.sonr.object.MsgCreateObject", value: MsgCreateObject.fromPartial( data ) }),
    
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
