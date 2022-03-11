// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgWithdrawAllAvailable } from "./types/cfevesting/tx";
import { MsgDelegate } from "./types/cfevesting/tx";
import { MsgVest } from "./types/cfevesting/tx";
import { MsgUndelegate } from "./types/cfevesting/tx";
import { MsgWithdrawDelegatorReward } from "./types/cfevesting/tx";
import { MsgBeginRedelegate } from "./types/cfevesting/tx";


const types = [
  ["/chain4energy.c4echain.cfevesting.MsgWithdrawAllAvailable", MsgWithdrawAllAvailable],
  ["/chain4energy.c4echain.cfevesting.MsgDelegate", MsgDelegate],
  ["/chain4energy.c4echain.cfevesting.MsgVest", MsgVest],
  ["/chain4energy.c4echain.cfevesting.MsgUndelegate", MsgUndelegate],
  ["/chain4energy.c4echain.cfevesting.MsgWithdrawDelegatorReward", MsgWithdrawDelegatorReward],
  ["/chain4energy.c4echain.cfevesting.MsgBeginRedelegate", MsgBeginRedelegate],
  
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
    msgWithdrawAllAvailable: (data: MsgWithdrawAllAvailable): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfevesting.MsgWithdrawAllAvailable", value: MsgWithdrawAllAvailable.fromPartial( data ) }),
    msgDelegate: (data: MsgDelegate): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfevesting.MsgDelegate", value: MsgDelegate.fromPartial( data ) }),
    msgVest: (data: MsgVest): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfevesting.MsgVest", value: MsgVest.fromPartial( data ) }),
    msgUndelegate: (data: MsgUndelegate): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfevesting.MsgUndelegate", value: MsgUndelegate.fromPartial( data ) }),
    msgWithdrawDelegatorReward: (data: MsgWithdrawDelegatorReward): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfevesting.MsgWithdrawDelegatorReward", value: MsgWithdrawDelegatorReward.fromPartial( data ) }),
    msgBeginRedelegate: (data: MsgBeginRedelegate): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfevesting.MsgBeginRedelegate", value: MsgBeginRedelegate.fromPartial( data ) }),
    
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