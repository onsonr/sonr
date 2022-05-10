// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgBuyNameAlias } from "./types/registry/v1/tx";
import { MsgTransferNameAlias } from "./types/registry/v1/tx";
import { MsgCreateWhoIs } from "./types/registry/v1/tx";
import { MsgTransferAppAlias } from "./types/registry/v1/tx";
import { MsgDeactivateWhoIs } from "./types/registry/v1/tx";
import { MsgBuyAppAlias } from "./types/registry/v1/tx";
import { MsgUpdateWhoIs } from "./types/registry/v1/tx";


const types = [
  ["/sonrio.sonr.registry.MsgBuyNameAlias", MsgBuyNameAlias],
  ["/sonrio.sonr.registry.MsgTransferNameAlias", MsgTransferNameAlias],
  ["/sonrio.sonr.registry.MsgCreateWhoIs", MsgCreateWhoIs],
  ["/sonrio.sonr.registry.MsgTransferAppAlias", MsgTransferAppAlias],
  ["/sonrio.sonr.registry.MsgDeactivateWhoIs", MsgDeactivateWhoIs],
  ["/sonrio.sonr.registry.MsgBuyAppAlias", MsgBuyAppAlias],
  ["/sonrio.sonr.registry.MsgUpdateWhoIs", MsgUpdateWhoIs],
  
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
    msgBuyNameAlias: (data: MsgBuyNameAlias): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgBuyNameAlias", value: MsgBuyNameAlias.fromPartial( data ) }),
    msgTransferNameAlias: (data: MsgTransferNameAlias): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgTransferNameAlias", value: MsgTransferNameAlias.fromPartial( data ) }),
    msgCreateWhoIs: (data: MsgCreateWhoIs): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgCreateWhoIs", value: MsgCreateWhoIs.fromPartial( data ) }),
    msgTransferAppAlias: (data: MsgTransferAppAlias): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgTransferAppAlias", value: MsgTransferAppAlias.fromPartial( data ) }),
    msgDeactivateWhoIs: (data: MsgDeactivateWhoIs): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgDeactivateWhoIs", value: MsgDeactivateWhoIs.fromPartial( data ) }),
    msgBuyAppAlias: (data: MsgBuyAppAlias): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgBuyAppAlias", value: MsgBuyAppAlias.fromPartial( data ) }),
    msgUpdateWhoIs: (data: MsgUpdateWhoIs): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgUpdateWhoIs", value: MsgUpdateWhoIs.fromPartial( data ) }),
    
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
