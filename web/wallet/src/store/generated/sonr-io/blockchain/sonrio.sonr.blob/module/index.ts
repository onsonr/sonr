// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import {
  Registry,
  OfflineSigner,
  EncodeObject,
  DirectSecp256k1HdWallet,
} from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDownloadBlob } from "./types/blob/tx";
import { MsgUploadBlob } from "./types/blob/tx";
import { MsgCreateThereIs } from "./types/blob/tx";
import { MsgUpdateThereIs } from "./types/blob/tx";
import { MsgDeleteBlob } from "./types/blob/tx";
import { MsgSyncBlob } from "./types/blob/tx";
import { MsgDeleteThereIs } from "./types/blob/tx";

const types = [
  ["/sonrio.sonr.blob.MsgDownloadBlob", MsgDownloadBlob],
  ["/sonrio.sonr.blob.MsgUploadBlob", MsgUploadBlob],
  ["/sonrio.sonr.blob.MsgCreateThereIs", MsgCreateThereIs],
  ["/sonrio.sonr.blob.MsgUpdateThereIs", MsgUpdateThereIs],
  ["/sonrio.sonr.blob.MsgDeleteBlob", MsgDeleteBlob],
  ["/sonrio.sonr.blob.MsgSyncBlob", MsgSyncBlob],
  ["/sonrio.sonr.blob.MsgDeleteThereIs", MsgDeleteThereIs],
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string;
}

interface SignAndBroadcastOptions {
  fee: StdFee;
  memo?: string;
}

const txClient = async (
  wallet: OfflineSigner,
  { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }
) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, {
      registry,
    });
  } else {
    client = await SigningStargateClient.offline(wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (
      msgs: EncodeObject[],
      { fee, memo }: SignAndBroadcastOptions = { fee: defaultFee, memo: "" }
    ) => client.signAndBroadcast(address, msgs, fee, memo),
    msgDownloadBlob: (data: MsgDownloadBlob): EncodeObject => ({
      typeUrl: "/sonrio.sonr.blob.MsgDownloadBlob",
      value: MsgDownloadBlob.fromPartial(data),
    }),
    msgUploadBlob: (data: MsgUploadBlob): EncodeObject => ({
      typeUrl: "/sonrio.sonr.blob.MsgUploadBlob",
      value: MsgUploadBlob.fromPartial(data),
    }),
    msgCreateThereIs: (data: MsgCreateThereIs): EncodeObject => ({
      typeUrl: "/sonrio.sonr.blob.MsgCreateThereIs",
      value: MsgCreateThereIs.fromPartial(data),
    }),
    msgUpdateThereIs: (data: MsgUpdateThereIs): EncodeObject => ({
      typeUrl: "/sonrio.sonr.blob.MsgUpdateThereIs",
      value: MsgUpdateThereIs.fromPartial(data),
    }),
    msgDeleteBlob: (data: MsgDeleteBlob): EncodeObject => ({
      typeUrl: "/sonrio.sonr.blob.MsgDeleteBlob",
      value: MsgDeleteBlob.fromPartial(data),
    }),
    msgSyncBlob: (data: MsgSyncBlob): EncodeObject => ({
      typeUrl: "/sonrio.sonr.blob.MsgSyncBlob",
      value: MsgSyncBlob.fromPartial(data),
    }),
    msgDeleteThereIs: (data: MsgDeleteThereIs): EncodeObject => ({
      typeUrl: "/sonrio.sonr.blob.MsgDeleteThereIs",
      value: MsgDeleteThereIs.fromPartial(data),
    }),
  };
};

interface QueryClientOptions {
  addr: string;
}

const queryClient = async (
  { addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }
) => {
  return new Api({ baseUrl: addr });
};

export { txClient, queryClient };
