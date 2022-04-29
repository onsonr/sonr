// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateWhoIs } from "./types/registry/tx";
import { MsgRegisterApplication } from "./types/registry/tx";
import { MsgUpdateWhoIs } from "./types/registry/tx";
import { MsgUpdateName } from "./types/registry/tx";
import { MsgUpdateApplication } from "./types/registry/tx";
import { MsgAccessApplication } from "./types/registry/tx";
import { MsgDeleteWhoIs } from "./types/registry/tx";
import { MsgRegisterName } from "./types/registry/tx";
import { MsgAccessName } from "./types/registry/tx";


const types = [
  ["/sonrio.sonr.registry.MsgCreateWhoIs", MsgCreateWhoIs],
  ["/sonrio.sonr.registry.MsgRegisterApplication", MsgRegisterApplication],
  ["/sonrio.sonr.registry.MsgUpdateWhoIs", MsgUpdateWhoIs],
  ["/sonrio.sonr.registry.MsgUpdateName", MsgUpdateName],
  ["/sonrio.sonr.registry.MsgUpdateApplication", MsgUpdateApplication],
  ["/sonrio.sonr.registry.MsgAccessApplication", MsgAccessApplication],
  ["/sonrio.sonr.registry.MsgDeleteWhoIs", MsgDeleteWhoIs],
  ["/sonrio.sonr.registry.MsgRegisterName", MsgRegisterName],
  ["/sonrio.sonr.registry.MsgAccessName", MsgAccessName],
  
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
    msgRegisterApplication: (data: MsgRegisterApplication): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgRegisterApplication", value: MsgRegisterApplication.fromPartial( data ) }),
    msgUpdateWhoIs: (data: MsgUpdateWhoIs): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgUpdateWhoIs", value: MsgUpdateWhoIs.fromPartial( data ) }),
    msgUpdateName: (data: MsgUpdateName): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgUpdateName", value: MsgUpdateName.fromPartial( data ) }),
    msgUpdateApplication: (data: MsgUpdateApplication): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgUpdateApplication", value: MsgUpdateApplication.fromPartial( data ) }),
    msgAccessApplication: (data: MsgAccessApplication): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgAccessApplication", value: MsgAccessApplication.fromPartial( data ) }),
    msgDeleteWhoIs: (data: MsgDeleteWhoIs): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgDeleteWhoIs", value: MsgDeleteWhoIs.fromPartial( data ) }),
    msgRegisterName: (data: MsgRegisterName): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgRegisterName", value: MsgRegisterName.fromPartial( data ) }),
    msgAccessName: (data: MsgAccessName): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgAccessName", value: MsgAccessName.fromPartial( data ) }),
    
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
