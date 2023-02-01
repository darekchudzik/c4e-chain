// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgInitialClaim } from "./types/cfeairdrop/tx";
import { MsgAddClaimRecords } from "./types/cfeairdrop/tx";
import { MsgRemoveAirdropCampaign } from "./types/cfeairdrop/tx";
import { MsgAddMissionToAidropCampaign } from "./types/cfeairdrop/tx";
import { MsgCloseAirdropCampaign } from "./types/cfeairdrop/tx";
import { MsgEditAirdropCampaign } from "./types/cfeairdrop/tx";
import { MsgClaim } from "./types/cfeairdrop/tx";
import { MsgStartAirdropCampaign } from "./types/cfeairdrop/tx";
import { MsgCreateAirdropCampaign } from "./types/cfeairdrop/tx";
import { MsgDeleteClaimRecord } from "./types/cfeairdrop/tx";


const types = [
  ["/chain4energy.c4echain.cfeairdrop.MsgInitialClaim", MsgInitialClaim],
  ["/chain4energy.c4echain.cfeairdrop.MsgAddClaimRecords", MsgAddClaimRecords],
  ["/chain4energy.c4echain.cfeairdrop.MsgRemoveAirdropCampaign", MsgRemoveAirdropCampaign],
  ["/chain4energy.c4echain.cfeairdrop.MsgAddMissionToAidropCampaign", MsgAddMissionToAidropCampaign],
  ["/chain4energy.c4echain.cfeairdrop.MsgCloseAirdropCampaign", MsgCloseAirdropCampaign],
  ["/chain4energy.c4echain.cfeairdrop.MsgEditAirdropCampaign", MsgEditAirdropCampaign],
  ["/chain4energy.c4echain.cfeairdrop.MsgClaim", MsgClaim],
  ["/chain4energy.c4echain.cfeairdrop.MsgStartAirdropCampaign", MsgStartAirdropCampaign],
  ["/chain4energy.c4echain.cfeairdrop.MsgCreateAirdropCampaign", MsgCreateAirdropCampaign],
  ["/chain4energy.c4echain.cfeairdrop.MsgDeleteClaimRecord", MsgDeleteClaimRecord],
  
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
    msgInitialClaim: (data: MsgInitialClaim): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgInitialClaim", value: MsgInitialClaim.fromPartial( data ) }),
    msgAddClaimRecords: (data: MsgAddClaimRecords): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgAddClaimRecords", value: MsgAddClaimRecords.fromPartial( data ) }),
    msgRemoveAirdropCampaign: (data: MsgRemoveAirdropCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgRemoveAirdropCampaign", value: MsgRemoveAirdropCampaign.fromPartial( data ) }),
    msgAddMissionToAidropCampaign: (data: MsgAddMissionToAidropCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgAddMissionToAidropCampaign", value: MsgAddMissionToAidropCampaign.fromPartial( data ) }),
    msgCloseAirdropCampaign: (data: MsgCloseAirdropCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgCloseAirdropCampaign", value: MsgCloseAirdropCampaign.fromPartial( data ) }),
    msgEditAirdropCampaign: (data: MsgEditAirdropCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgEditAirdropCampaign", value: MsgEditAirdropCampaign.fromPartial( data ) }),
    msgClaim: (data: MsgClaim): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgClaim", value: MsgClaim.fromPartial( data ) }),
    msgStartAirdropCampaign: (data: MsgStartAirdropCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgStartAirdropCampaign", value: MsgStartAirdropCampaign.fromPartial( data ) }),
    msgCreateAirdropCampaign: (data: MsgCreateAirdropCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgCreateAirdropCampaign", value: MsgCreateAirdropCampaign.fromPartial( data ) }),
    msgDeleteClaimRecord: (data: MsgDeleteClaimRecord): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgDeleteClaimRecord", value: MsgDeleteClaimRecord.fromPartial( data ) }),
    
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
