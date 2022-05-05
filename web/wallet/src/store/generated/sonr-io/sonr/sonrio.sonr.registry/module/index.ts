// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAccessService } from "./types/registry/tx";
import { MsgRegisterName } from "./types/registry/tx";
import { MsgAccessName } from "./types/registry/tx";
import { MsgRegisterService } from "./types/registry/tx";
import { MsgUpdateService } from "./types/registry/tx";
import { MsgUpdateName } from "./types/registry/tx";


const types = [
  ["/sonrio.sonr.registry.MsgAccessService", MsgAccessService],
  ["/sonrio.sonr.registry.MsgRegisterName", MsgRegisterName],
  ["/sonrio.sonr.registry.MsgAccessName", MsgAccessName],
  ["/sonrio.sonr.registry.MsgRegisterService", MsgRegisterService],
  ["/sonrio.sonr.registry.MsgUpdateService", MsgUpdateService],
  ["/sonrio.sonr.registry.MsgUpdateName", MsgUpdateName],
  
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
    msgAccessService: (data: MsgAccessService): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgAccessService", value: MsgAccessService.fromPartial( data ) }),
    msgRegisterName: (data: MsgRegisterName): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgRegisterName", value: MsgRegisterName.fromPartial( data ) }),
    msgAccessName: (data: MsgAccessName): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgAccessName", value: MsgAccessName.fromPartial( data ) }),
    msgRegisterService: (data: MsgRegisterService): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgRegisterService", value: MsgRegisterService.fromPartial( data ) }),
    msgUpdateService: (data: MsgUpdateService): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgUpdateService", value: MsgUpdateService.fromPartial( data ) }),
    msgUpdateName: (data: MsgUpdateName): EncodeObject => ({ typeUrl: "/sonrio.sonr.registry.MsgUpdateName", value: MsgUpdateName.fromPartial( data ) }),
    
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
