// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateWhoIs } from "./types/registry/v1/tx";
import { MsgUpdateWhoIs } from "./types/registry/v1/tx";
import { MsgTransferAlias } from "./types/registry/v1/tx";
import { MsgSellAlias } from "./types/registry/v1/tx";
import { MsgBuyAlias } from "./types/registry/v1/tx";
import { MsgDeactivateWhoIs } from "./types/registry/v1/tx";


const types = [
  ["/sonrio.sonr.registry.MsgCreateWhoIs", MsgCreateWhoIs],
  ["/sonrio.sonr.registry.MsgUpdateWhoIs", MsgUpdateWhoIs],
  ["/sonrio.sonr.registry.MsgTransferAlias", MsgTransferAlias],
  ["/sonrio.sonr.registry.MsgSellAlias", MsgSellAlias],
  ["/sonrio.sonr.registry.MsgBuyAlias", MsgBuyAlias],
  ["/sonrio.sonr.registry.MsgDeactivateWhoIs", MsgDeactivateWhoIs],
  
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
    msgCreateWhoIs: (data: MsgCreateWhoIs): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgCreateWhoIs", value: MsgCreateWhoIs.fromPartial( data ) }),
    msgUpdateWhoIs: (data: MsgUpdateWhoIs): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgUpdateWhoIs", value: MsgUpdateWhoIs.fromPartial( data ) }),
    msgTransferAlias: (data: MsgTransferAlias): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgTransferAlias", value: MsgTransferAlias.fromPartial( data ) }),
    msgSellAlias: (data: MsgSellAlias): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgSellAlias", value: MsgSellAlias.fromPartial( data ) }),
    msgBuyAlias: (data: MsgBuyAlias): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgBuyAlias", value: MsgBuyAlias.fromPartial( data ) }),
    msgDeactivateWhoIs: (data: MsgDeactivateWhoIs): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgDeactivateWhoIs", value: MsgDeactivateWhoIs.fromPartial( data ) }),
    
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
