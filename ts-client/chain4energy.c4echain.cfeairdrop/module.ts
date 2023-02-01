// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgStartCampaign } from "./types/cfeairdrop/tx";
import { MsgAddMissionToAidropCampaign } from "./types/cfeairdrop/tx";
import { MsgCreateCampaign } from "./types/cfeairdrop/tx";
import { MsgInitialClaim } from "./types/cfeairdrop/tx";
import { MsgEditCampaign } from "./types/cfeairdrop/tx";
import { MsgAddClaimRecords } from "./types/cfeairdrop/tx";
import { MsgDeleteClaimRecord } from "./types/cfeairdrop/tx";
import { MsgCloseCampaign } from "./types/cfeairdrop/tx";
import { MsgClaim } from "./types/cfeairdrop/tx";


export { MsgStartCampaign, MsgAddMissionToAidropCampaign, MsgCreateCampaign, MsgInitialClaim, MsgEditCampaign, MsgAddClaimRecords, MsgDeleteClaimRecord, MsgCloseCampaign, MsgClaim };

type sendMsgStartCampaignParams = {
  value: MsgStartCampaign,
  fee?: StdFee,
  memo?: string
};

type sendMsgAddMissionToAidropCampaignParams = {
  value: MsgAddMissionToAidropCampaign,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateCampaignParams = {
  value: MsgCreateCampaign,
  fee?: StdFee,
  memo?: string
};

type sendMsgInitialClaimParams = {
  value: MsgInitialClaim,
  fee?: StdFee,
  memo?: string
};

type sendMsgEditCampaignParams = {
  value: MsgEditCampaign,
  fee?: StdFee,
  memo?: string
};

type sendMsgAddClaimRecordsParams = {
  value: MsgAddClaimRecords,
  fee?: StdFee,
  memo?: string
};

type sendMsgDeleteClaimRecordParams = {
  value: MsgDeleteClaimRecord,
  fee?: StdFee,
  memo?: string
};

type sendMsgCloseCampaignParams = {
  value: MsgCloseCampaign,
  fee?: StdFee,
  memo?: string
};

type sendMsgClaimParams = {
  value: MsgClaim,
  fee?: StdFee,
  memo?: string
};


type msgStartCampaignParams = {
  value: MsgStartCampaign,
};

type msgAddMissionToAidropCampaignParams = {
  value: MsgAddMissionToAidropCampaign,
};

type msgCreateCampaignParams = {
  value: MsgCreateCampaign,
};

type msgInitialClaimParams = {
  value: MsgInitialClaim,
};

type msgEditCampaignParams = {
  value: MsgEditCampaign,
};

type msgAddClaimRecordsParams = {
  value: MsgAddClaimRecords,
};

type msgDeleteClaimRecordParams = {
  value: MsgDeleteClaimRecord,
};

type msgCloseCampaignParams = {
  value: MsgCloseCampaign,
};

type msgClaimParams = {
  value: MsgClaim,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgStartCampaign({ value, fee, memo }: sendMsgStartCampaignParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgStartCampaign: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgStartCampaign({ value: MsgStartCampaign.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgStartCampaign: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAddMissionToAidropCampaign({ value, fee, memo }: sendMsgAddMissionToAidropCampaignParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAddMissionToAidropCampaign: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAddMissionToAidropCampaign({ value: MsgAddMissionToAidropCampaign.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAddMissionToAidropCampaign: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateCampaign({ value, fee, memo }: sendMsgCreateCampaignParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateCampaign: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateCampaign({ value: MsgCreateCampaign.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateCampaign: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgInitialClaim({ value, fee, memo }: sendMsgInitialClaimParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgInitialClaim: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgInitialClaim({ value: MsgInitialClaim.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgInitialClaim: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgEditCampaign({ value, fee, memo }: sendMsgEditCampaignParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgEditCampaign: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgEditCampaign({ value: MsgEditCampaign.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgEditCampaign: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAddClaimRecords({ value, fee, memo }: sendMsgAddClaimRecordsParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAddClaimRecords: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAddClaimRecords({ value: MsgAddClaimRecords.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAddClaimRecords: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgDeleteClaimRecord({ value, fee, memo }: sendMsgDeleteClaimRecordParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDeleteClaimRecord: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDeleteClaimRecord({ value: MsgDeleteClaimRecord.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDeleteClaimRecord: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCloseCampaign({ value, fee, memo }: sendMsgCloseCampaignParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCloseCampaign: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCloseCampaign({ value: MsgCloseCampaign.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCloseCampaign: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgClaim({ value, fee, memo }: sendMsgClaimParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgClaim: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgClaim({ value: MsgClaim.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgClaim: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgStartCampaign({ value }: msgStartCampaignParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgStartCampaign", value: MsgStartCampaign.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgStartCampaign: Could not create message: ' + e.message)
			}
		},
		
		msgAddMissionToAidropCampaign({ value }: msgAddMissionToAidropCampaignParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgAddMissionToAidropCampaign", value: MsgAddMissionToAidropCampaign.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAddMissionToAidropCampaign: Could not create message: ' + e.message)
			}
		},
		
		msgCreateCampaign({ value }: msgCreateCampaignParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgCreateCampaign", value: MsgCreateCampaign.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateCampaign: Could not create message: ' + e.message)
			}
		},
		
		msgInitialClaim({ value }: msgInitialClaimParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgInitialClaim", value: MsgInitialClaim.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgInitialClaim: Could not create message: ' + e.message)
			}
		},
		
		msgEditCampaign({ value }: msgEditCampaignParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgEditCampaign", value: MsgEditCampaign.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgEditCampaign: Could not create message: ' + e.message)
			}
		},
		
		msgAddClaimRecords({ value }: msgAddClaimRecordsParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgAddClaimRecords", value: MsgAddClaimRecords.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAddClaimRecords: Could not create message: ' + e.message)
			}
		},
		
		msgDeleteClaimRecord({ value }: msgDeleteClaimRecordParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgDeleteClaimRecord", value: MsgDeleteClaimRecord.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDeleteClaimRecord: Could not create message: ' + e.message)
			}
		},
		
		msgCloseCampaign({ value }: msgCloseCampaignParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgCloseCampaign", value: MsgCloseCampaign.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCloseCampaign: Could not create message: ' + e.message)
			}
		},
		
		msgClaim({ value }: msgClaimParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgClaim", value: MsgClaim.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgClaim: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			Chain4EnergyC4EchainCfeairdrop: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;